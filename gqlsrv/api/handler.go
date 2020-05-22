package gql

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	claims "github.com/obiwan007/usersrv/pkg/claims"
	pb "github.com/obiwan007/usersrv/proto"
	"github.com/opentracing/opentracing-go"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"golang.org/x/oauth2/microsoft"
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
	googleOauthConfig    *oauth2.Config
	appleOauthConfig     *oauth2.Config
	microsoftOauthConfig *oauth2.Config
	// TODO: randomize it
	oauthStateString = "pseudo-random"
	UserSrvClient    pb.UserServiceClient
)

type loginUser struct {
	Username string
	Password string
}
type loginResponse struct {
	Token string `json:"token"`
}

func handleRefresh(w http.ResponseWriter, r *http.Request) {
	t := r.Context().Value("jwt")
	existingToken, ok := t.(*jwt.Token)
	if !ok || !existingToken.Valid {
		http.Error(w, "Unauthorized, no valid token provided", 401)
		return
	}

	claims := existingToken.Claims.(*claims.MyCustomClaims)
	log.Println("Refresh Subject:", claims.Subject)
	token, err := getToken(claims.Name, claims.Picture, claims.Mandant, claims.Subject)
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

func handleAppleLogin(w http.ResponseWriter, r *http.Request) {
	url := appleOauthConfig.AuthCodeURL(oauthStateString)
	log.Println("Handling Apple Login Redirect", appleOauthConfig.ClientID)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func handleMicrosoftLogin(w http.ResponseWriter, r *http.Request) {
	url := microsoftOauthConfig.AuthCodeURL(oauthStateString)
	log.Println("Handling Microsoft Login Redirect", microsoftOauthConfig.ClientID)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func handleGoogleLogin(w http.ResponseWriter, r *http.Request) {

	url := googleOauthConfig.AuthCodeURL(oauthStateString)
	log.Println("Handling Google Login Redirect", googleOauthConfig.ClientID)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
	return
}

func handleDefaultLogin(w http.ResponseWriter, r *http.Request) {
	keys, ok := r.URL.Query()["provider"]

	if !ok || len(keys[0]) < 1 {
		log.Println("Url Param 'provider' is missing")
		handleGoogleLogin(w, r)
		return
	}
	key := keys[0]
	log.Println("Url Param 'provider' is: " + string(key))
	switch key {
	case "apple":
		handleAppleLogin(w, r)
		break

	case "microsoft":
		handleMicrosoftLogin(w, r)
		break
	}
}

func handleGoogleCallback(w http.ResponseWriter, r *http.Request) {
	content, err := getUserInfoGoogle(r.FormValue("state"), r.FormValue("code"))
	if err != nil {
		fmt.Println(err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	log.Println("Content", content)
	log.Println("Name", content.Name)
	token, err := getToken(content.Name, content.Picture, "0", content.Email)
	if err != nil {
		http.Error(w, err.Error(), 400)
	}
	UserSrvClient.AddUser(r.Context(), &pb.User{Jwt: token, Name: content.Name, Email: content.Email})
	// loginRes := &loginResponse{Token: token}
	// res, err := json.Marshal(loginRes)

	http.SetCookie(w, &http.Cookie{Name: "Auth", Value: token, HttpOnly: true, Path: "/", Expires: time.Now().Add(time.Hour * 1)})
	// w.Write(res)
	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
	// log.Fprintf(w, "Content: %s\n", content)
}

func handleMicrosoftCallback(w http.ResponseWriter, r *http.Request) {
	authorizationCode := r.URL.Query().Get("code")
	log.Println("Code:", authorizationCode)
	content, err := getUserInfoMicrosoft(r.FormValue("state"), r.FormValue("code"))
	if err != nil {
		fmt.Println(err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	log.Println("Name", content.Name)
	log.Println("Email", content.Email)
	token, err := getToken(content.Name, content.Picture, "0", content.Email)
	if err != nil {
		http.Error(w, err.Error(), 400)
	}
	UserSrvClient.AddUser(r.Context(), &pb.User{Jwt: token, Name: content.Name, Email: content.Email})
	// loginRes := &loginResponse{Token: token}
	// res, err := json.Marshal(loginRes)

	http.SetCookie(w, &http.Cookie{Name: "Auth", Value: token, HttpOnly: true, Path: "/", Expires: time.Now().Add(time.Hour * 1)})
	// w.Write(res)
	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
	// log.Fprintf(w, "Content: %s\n", content)
}

type UserInfoMicrosoft struct {
	Sub           string `json:"sub"`
	Name          string `json:"displayName"`
	GivenName     string `json:"givenName"`
	FamilyName    string `json:"surname"`
	Profile       string `json:"profile"`
	Picture       string `json:"picture"`
	Email         string `json:"userPrincipalName"`
	EmailVerified bool   `json:"email_verified"`
	Gender        string `json:"gender"`
}

func getUserInfoMicrosoft(state string, code string) (*UserInfoMicrosoft, error) {
	if state != oauthStateString {
		return nil, fmt.Errorf("invalid oauth state")
	}
	log.Println("Code:", code, microsoftOauthConfig.RedirectURL)

	// opts:=[]oauth2.AuthCodeOption{"wl.email",},
	token, err := microsoftOauthConfig.Exchange(context.Background(), code)
	if err != nil {
		return nil, fmt.Errorf("code exchange failed: %s", err.Error())
	}
	log.Println("Access token retrieved from Azure", token.AccessToken)
	url := "https://graph.microsoft.com/v1.0/me"
	// Create a Bearer string by appending string access token
	var bearer = "Bearer " + token.AccessToken

	// Create a new request using http
	req, err := http.NewRequest("GET", url, nil)

	// add authorization header to the req
	req.Header.Add("Authorization", bearer)
	req.Header.Add("Accept", "application/json")
	// Send req using http Client
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed getting user info: %s", err.Error())
	}

	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed reading response body: %s", err.Error())
	}
	log.Println("Content", string(contents))

	var result UserInfoMicrosoft
	if err := json.Unmarshal(contents, &result); err != nil {
		return nil, err
	}

	// urlPhoto := "https://graph.microsoft.com/beta/users/" + result.Email + "/photos/48x48/$value"
	// reqPhoto, err := http.NewRequest("GET", urlPhoto, nil)
	// reqPhoto.Header.Add("Authorization", bearer)
	// client = &http.Client{}
	// responsePhoto, err := client.Do(reqPhoto)
	// if responsePhoto.StatusCode != 200 {
	// 	log.Println("No photo found", responsePhoto.StatusCode)
	// 	return &result, nil
	// }
	// if err != nil {
	// 	return nil, fmt.Errorf("failed getting user Photo: %s", err.Error())
	// }

	// defer responsePhoto.Body.Close()
	// contentsPhoto, err := ioutil.ReadAll(responsePhoto.Body)
	// if err != nil {
	// 	return nil, fmt.Errorf("failed reading response body: %s", err.Error())
	// }

	// imgBase64Str := base64.StdEncoding.EncodeToString(contentsPhoto)
	// log.Println("Content Photo:", urlPhoto, len(imgBase64Str))

	// result.Picture = "data:image/jpeg;base64," + imgBase64Str

	return &result, nil
}

func handleAppleCallback(w http.ResponseWriter, r *http.Request) {
	content, err := getUserInfoGoogle(r.FormValue("state"), r.FormValue("code"))
	if err != nil {
		fmt.Println(err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	log.Println("Content", content)
	log.Println("Name", content.Name)
	token, err := getToken(content.Name, content.Picture, "0", content.Email)
	if err != nil {
		http.Error(w, err.Error(), 400)
	}
	UserSrvClient.AddUser(r.Context(), &pb.User{Jwt: token, Name: content.Name, Email: content.Email})
	// loginRes := &loginResponse{Token: token}
	// res, err := json.Marshal(loginRes)

	http.SetCookie(w, &http.Cookie{Name: "Auth", Value: token, HttpOnly: true, Path: "/", Expires: time.Now().Add(time.Hour * 1)})
	// w.Write(res)
	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
	// log.Fprintf(w, "Content: %s\n", content)
}

type UserInfoGoogle struct {
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

func getUserInfoGoogle(state string, code string) (*UserInfoGoogle, error) {
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

	var result UserInfoGoogle
	if err := json.Unmarshal(contents, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type AuthSecret struct {
	ClientID, ClientSecr, RedirectUrl, Tenant *string
}

func NewRouter(schema *graphql.Schema, tracer opentracing.Tracer, gSecrets, appleSecrets, microsoftSecrets *AuthSecret) *TracedServeMux {
	// mux := http.NewServeMux()

	googleOauthConfig = &oauth2.Config{
		ClientID:     *gSecrets.ClientID,   //
		ClientSecret: *gSecrets.ClientSecr, //
		RedirectURL:  *gSecrets.RedirectUrl,
		Scopes: []string{"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile"},
		Endpoint: google.Endpoint,
	}

	appleOauthConfig = &oauth2.Config{
		ClientID:     *appleSecrets.ClientID,   //
		ClientSecret: *appleSecrets.ClientSecr, //
		RedirectURL:  *appleSecrets.RedirectUrl,
		Scopes:       []string{"email", "name"},
		Endpoint: oauth2.Endpoint{AuthURL: "https://appleid.apple.com/auth/authorize",
			TokenURL: "https://appleid.apple.com/auth/token"},
	}

	microsoftOauthConfig = &oauth2.Config{
		ClientID:     *microsoftSecrets.ClientID,   //
		ClientSecret: *microsoftSecrets.ClientSecr, //
		RedirectURL:  *microsoftSecrets.RedirectUrl,
		Scopes:       []string{"user.read", "profile", "openid"},
		Endpoint:     microsoft.AzureADEndpoint(*microsoftSecrets.Tenant),
	}

	// log.Printf("OAuth %s %s %s", (*gClientID)[:6], (*gClientSecr)[:6], *redirectUrl)
	mux := NewServeMux(tracer)

	mux.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write(page)
	}))
	mux.Handle("/query", &relay.Handler{Schema: schema})

	mux.Handle("/auth/login", http.HandlerFunc(handleDefaultLogin))
	mux.Handle("/auth/logina", http.HandlerFunc(handleAppleLogin))
	mux.Handle("/auth/callback", http.HandlerFunc(handleGoogleCallback))
	mux.Handle("/auth/callbackA", http.HandlerFunc(handleAppleCallback))
	mux.Handle("/auth/callbackM", http.HandlerFunc(handleMicrosoftCallback))
	mux.Handle("/auth/refresh", http.HandlerFunc(handleRefresh))
	// TODO: Add more routes here for other endpoints.
	// TODO: Add authentication endpoints or serving up regular assets?

	return mux
}
