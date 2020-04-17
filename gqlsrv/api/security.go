package gql

import (
	"context"
	"fmt"
	"log"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
)

var mySigningKey = []byte("captainjacksparrowsayshi")

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
	fmt.Println("Endpoint Hit: homePage")

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
			next.ServeHTTP(w, r.WithContext(context.WithValue(ctx, "jwt", token)))
			// if token.Valid {
			// 	next.ServeHTTP(w, r.WithContext(context.WithValue(ctx, "jwt", token)))
			// }
		} else {
			log.Fatalln("No Token provided")
			// fmt.Fprintf(w, "Not Authorized")
			next.ServeHTTP(w, r)
		}
	})
}
