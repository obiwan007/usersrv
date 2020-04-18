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

func getToken() (string, error) {
	// Create the Claims
	claims := MyCustomClaims{
		"bar",
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 1).Unix(),
			Issuer:    "test",
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
		log.Println("Handling AUthorization")
		ctx := r.Context()
		if r.Header["Authorization"] != nil {
			log.Println("Token is provided")
			token, err := jwt.ParseWithClaims(r.Header["Authorization"][0], &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("There was an error")
				}
				return mySigningKey, nil
			})

			if err != nil {
				log.Println("ERROR: Invalid token:", err.Error())
				// fmt.Fprintf(w, err.Error())
				token = nil
			}

			// claims, ok := token.Claims.(*MyCustomClaims)

			if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
				log.Printf("%v %v", claims.Email, claims.StandardClaims.ExpiresAt)
			} else {
				log.Println(err)
			}

			if token.Valid != true {
				log.Println(token.Valid)
				token = nil
			}

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
