package gql

import (
	"encoding/json"
	"fmt"
	"net/http"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/opentracing/opentracing-go"
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
            credentials: 'omit',
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

type loginUser struct {
	Username string
	Password string
}
type loginResponse struct {
	Token string `json:"token"`
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	fmt.Println("LOGIN HIT")

	var u loginUser
	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	fmt.Println(u.Username, u.Password)
	token, err := getToken()
	if err != nil {
		http.Error(w, err.Error(), 400)
	}
	loginRes := &loginResponse{Token: token}
	res, err := json.Marshal(loginRes)

	http.SetCookie(w, &http.Cookie{Name: "Flavor", Value: "Chocolate Chip", HttpOnly: true, Path: "/"})

	w.Write(res)
}

func NewRouter(schema *graphql.Schema, tracer opentracing.Tracer) *TracedServeMux {
	// mux := http.NewServeMux()
	mux := NewServeMux(tracer)

	mux.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write(page)
	}))
	mux.Handle("/query", &relay.Handler{Schema: schema})

	mux.Handle("/auth/login", http.HandlerFunc(handleLogin))
	// TODO: Add more routes here for other endpoints.
	// TODO: Add authentication endpoints or serving up regular assets?

	return mux
}
