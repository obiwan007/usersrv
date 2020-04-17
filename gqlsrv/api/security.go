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

func getToken() (string, error) {
	// Create the Claims
	claims := &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Minute * 1).Unix(),
		Issuer:    "test",
	}
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
			token, err := jwt.Parse(r.Header["Authorization"][0], func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("There was an error")
				}
				return mySigningKey, nil
			})

			if err != nil {
				log.Println("ERROR: Invalid token:", err.Error())
				// fmt.Fprintf(w, err.Error())
			}

			claims, ok := token.Claims.(jwt.MapClaims)
			if ok != true {
				log.Println("No claim in token")
			}
			if token.Valid != true {
				log.Println("Claims invalid", claims)
				fmt.Println(claims["exp"])
				fmt.Println(token.Valid)
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
