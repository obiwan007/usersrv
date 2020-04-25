package gql

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var mySigningKey = []byte("captainjacksparrowsayshi")

func homePage(w http.ResponseWriter, r *http.Request) {
	log.Println("Endpoint Hit: homePage")
	fmt.Fprintf(w, "Hello World")

}

type MyCustomClaims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

func getToken(subject string) (string, error) {
	// Create the Claims
	claims := MyCustomClaims{
		"bar",
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
			Issuer:    "gqlsrv",
			Subject:   subject,
		},
	}
	// claims := &jwt.StandardClaims{
	// 	ExpiresAt: time.Now().Add(time.Minute * 1).Unix(),
	// 	Issuer:    "test",
	// }

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	log.Println("Generated Token", ss)
	return ss, err
}

func isAuthorized(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Handling AUthorization", r.URL.String())
		ctx := r.Context()
		var foundAuth *string = nil
		for _, cookie := range r.Cookies() {
			if cookie.Name == "Auth" {
				foundAuth = &cookie.Value
				log.Println("Authorization Cookie given")
				break
			}
		}
		if r.Header["Authorization"] != nil {
			foundAuth = &r.Header["Authorization"][0]
			log.Println("Authorization Header given")
		}

		if foundAuth != nil {
			log.Println("Auth", *foundAuth)
			token, err := jwt.ParseWithClaims(*foundAuth, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("There was an error")
				}
				return mySigningKey, nil
			})

			if err != nil {
				log.Println("ERROR: Invalid token:", err.Error())
				// fmt.Fprintf(w, err.Error())
				token = nil
				next.ServeHTTP(w, r.WithContext(ctx))
				return
			}
			if token.Valid != true {
				log.Println(token.Valid)
				next.ServeHTTP(w, r.WithContext(ctx))
				return
			}
			// claims, ok := token.Claims.(*MyCustomClaims)

			// if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
			// 	log.Printf("%v %v", claims.StandardClaims.Subject, claims.StandardClaims.ExpiresAt)
			// } else {
			// 	log.Println(err)
			// }

			next.ServeHTTP(w, r.WithContext(context.WithValue(ctx, "jwt", token)))

			// if token.Valid {
			// 	next.ServeHTTP(w, r.WithContext(context.WithValue(ctx, "jwt", token)))
			// }
		} else {
			log.Println("No Token provided")
			// fmt.Fprintf(w, "Not Authorized")
			next.ServeHTTP(w, r.WithContext(ctx))
		}
	})
}
