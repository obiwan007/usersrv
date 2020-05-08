package gql

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/opentracing/opentracing-go"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var page = []byte(`
<!DOCTYPE html>
<html>
  <head>
    <style>
      body {
        height: 100%;
        margin: 0;
        width: 100%;
        overflow: hidden;
      }

      #graphiql {
        height: 100vh;
      }
    </style>

    <!--
      This GraphiQL example depends on Promise and fetch, which are available in
      modern browsers, but can be "polyfilled" for older browsers.
      GraphiQL itself depends on React DOM.
      If you do not want to rely on a CDN, you can host these files locally or
      include them directly in your favored resource bunder.
    -->
    <script
      crossorigin
      src="https://unpkg.com/react@16/umd/react.development.js"
    ></script>
    <script
      crossorigin
      src="https://unpkg.com/react-dom@16/umd/react-dom.development.js"
    ></script>

    <!--
      These two files can be found in the npm module, however you may wish to
      copy them directly into your environment, or perhaps include them in your
      favored resource bundler.
     -->
    <link rel="stylesheet" href="https://unpkg.com/graphiql/graphiql.min.css" />
  </head>

  <body>
    <div id="graphiql">Loading...</div>
    <script
      src="https://unpkg.com/graphiql/graphiql.min.js"
      type="application/javascript"
    ></script>
    <script src="/renderExample.js" type="application/javascript"></script>
    <script>
      function graphQLFetcher(graphQLParams) {
        return fetch(
          '/query',
          {
            method: 'post',
            headers: {
              Accept: 'application/json',
              'Content-Type': 'application/json',
            },
            body: JSON.stringify(graphQLParams),
            credentials: 'include',
          },
        ).then(function (response) {
          return response.json().catch(function () {
            return response.text();
          });
        });
      }

      ReactDOM.render(
        React.createElement(GraphiQL, {
          fetcher: graphQLFetcher,
          defaultVariableEditorOpen: true,
        }),
        document.getElementById('graphiql'),
      );
    </script>
  </body>
</html>
`)

var (
	googleOauthConfig *oauth2.Config
	// TODO: randomize it
	oauthStateString = "pseudo-random"
)

type loginUser struct {
	Username string
	Password string
}
type loginResponse struct {
	Token string `json:"token"`
}

func handleRefresh(w http.ResponseWriter, r *http.Request) {
	log.Println("Refresh HIT")

	t := r.Context().Value("jwt")

	existingToken, ok := t.(*jwt.Token)
	if !ok || !existingToken.Valid {
		http.Error(w, "Unauthorized, no valid token provided", 401)
		return
	}

	claims := existingToken.Claims.(*MyCustomClaims)
	log.Println("Refresh Subject:", claims.Subject)
	token, err := getToken(claims.Name, claims.Picture, claims.Subject)
	if err != nil {
		http.Error(w, err.Error(), 400)
	}
	loginRes := &loginResponse{Token: token}
	res, err := json.Marshal(loginRes)

	http.SetCookie(w, &http.Cookie{Name: "Auth", Value: token, HttpOnly: true, Path: "/", Expires: time.Now().Add(time.Hour * 1)})

	w.Write(res)
}

// func handleLogin(w http.ResponseWriter, r *http.Request) {
// 	log.Println("Login HIT")

// 	var u loginUser
// 	if r.Body == nil {
// 		http.Error(w, "Please send a request body", 400)
// 		return
// 	}

// 	err := json.NewDecoder(r.Body).Decode(&u)
// 	if err != nil {
// 		http.Error(w, err.Error(), 400)
// 		return
// 	}
// 	fmt.Println(u.Username, u.Password)
// 	token, err := getToken(u.Username, u.)
// 	if err != nil {
// 		http.Error(w, err.Error(), 400)
// 	}
// 	loginRes := &loginResponse{Token: token}
// 	res, err := json.Marshal(loginRes)

// 	http.SetCookie(w, &http.Cookie{Name: "Auth", Value: token, HttpOnly: true, Path: "/", Expires: time.Now().Add(time.Hour * 1)})

// 	w.Write(res)
// }

func handleGoogleLogin(w http.ResponseWriter, r *http.Request) {
	url := googleOauthConfig.AuthCodeURL(oauthStateString)
	log.Println("Handling Google Login Redirect", googleOauthConfig.ClientID)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func handleGoogleCallback(w http.ResponseWriter, r *http.Request) {
	content, err := getUserInfo(r.FormValue("state"), r.FormValue("code"))
	if err != nil {
		fmt.Println(err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	log.Println("Content", content)
	log.Println("Name", content.Name)
	token, err := getToken(content.Name, content.Picture, content.Email)
	if err != nil {
		http.Error(w, err.Error(), 400)
	}
	// loginRes := &loginResponse{Token: token}
	// res, err := json.Marshal(loginRes)

	http.SetCookie(w, &http.Cookie{Name: "Auth", Value: token, HttpOnly: true, Path: "/", Expires: time.Now().Add(time.Hour * 1)})
	// w.Write(res)
	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
	// log.Fprintf(w, "Content: %s\n", content)
}

type UserInfo struct {
	Sub           string `json:"sub"`
	Name          string `json:"name"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Profile       string `json:"profile"`
	Picture       string `json:"picture"`
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
	Gender        string `json:"gender"`
}

func getUserInfo(state string, code string) (*UserInfo, error) {
	if state != oauthStateString {
		return nil, fmt.Errorf("invalid oauth state")
	}

	token, err := googleOauthConfig.Exchange(oauth2.NoContext, code)
	if err != nil {
		return nil, fmt.Errorf("code exchange failed: %s", err.Error())
	}

	response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		return nil, fmt.Errorf("failed getting user info: %s", err.Error())
	}

	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed reading response body: %s", err.Error())
	}

	var result UserInfo
	if err := json.Unmarshal(contents, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func NewRouter(schema *graphql.Schema, tracer opentracing.Tracer, gClientID, gClientSecr, redirectUrl *string) *TracedServeMux {
	// mux := http.NewServeMux()

	googleOauthConfig = &oauth2.Config{
		ClientID:     *gClientID,   //
		ClientSecret: *gClientSecr, //
		RedirectURL:  *redirectUrl,
		Scopes: []string{"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile"},
		Endpoint: google.Endpoint,
	}
	log.Printf("OAuth %s %s %s", (*gClientID)[:6], (*gClientSecr)[:6], *redirectUrl)
	mux := NewServeMux(tracer)

	mux.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write(page)
	}))
	mux.Handle("/query", &relay.Handler{Schema: schema})

	mux.Handle("/auth/login", http.HandlerFunc(handleGoogleLogin))
	mux.Handle("/auth/callback", http.HandlerFunc(handleGoogleCallback))
	mux.Handle("/auth/refresh", http.HandlerFunc(handleRefresh))
	// TODO: Add more routes here for other endpoints.
	// TODO: Add authentication endpoints or serving up regular assets?

	return mux
}
