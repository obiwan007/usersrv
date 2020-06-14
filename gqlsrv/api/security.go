package gql

import (
	"context"
	"fmt"

	"github.com/leodotcloud/log"

	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	claims "github.com/obiwan007/usersrv/pkg/claims"
	"github.com/pkg/errors"
)

var SigningKey = []byte("thiswillbeoverwrittenlater")

func homePage(w http.ResponseWriter, r *http.Request) {
	log.Infof("Endpoint Hit: homePage")
	fmt.Fprintf(w, "Hello World")

}

func getToken(username, picture, mandant, email string) (string, error) {
	// Create the Claims
	claims := claims.MyCustomClaims{
		Name:    username,
		Picture: picture,
		Mandant: mandant,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
			Issuer:    "gqlsrv",
			Subject:   email,
		},
	}
	// claims := &jwt.StandardClaims{
	// 	ExpiresAt: time.Now().Add(time.Minute * 1).Unix(),
	// 	Issuer:    "test",
	// }

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(SigningKey)
	log.Infof("Generated Token")
	return ss, err
}

func isAuthorized(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Infof("Handling Authorization %s", r.URL.String())
		ctx := r.Context()
		var foundAuth *string = nil
		for _, cookie := range r.Cookies() {
			if cookie.Name == "Auth" {
				foundAuth = &cookie.Value
				log.Infof("Authorization Cookie given")
				break
			}
		}
		if r.Header["Authorization"] != nil {
			foundAuth = &r.Header["Authorization"][0]
			log.Infof("Authorization Header given")
		}

		if foundAuth != nil {
			// log.Println("Auth", *foundAuth)
			token, err := jwt.ParseWithClaims(*foundAuth, &claims.MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("There was an error")
				}
				return SigningKey, nil
			})

			if err != nil {
				log.Errorf("ERROR: Invalid token: %v", err.Error())
				// fmt.Fprintf(w, err.Error())
				token = nil
				next.ServeHTTP(w, r.WithContext(ctx))
				return
			}
			if token.Valid != true {
				log.Infof("Valid Token %v", token.Valid)
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
			log.Infof("No Token provided")
			// fmt.Fprintf(w, "Not Authorized")
			next.ServeHTTP(w, r.WithContext(ctx))
		}
	})
}

func validateToken(ctx context.Context) (*jwt.Token, error) {
	t := ctx.Value("jwt")

	token, ok := t.(*jwt.Token)
	if !ok || !token.Valid {
		return nil, errors.Errorf("Unauthorized")
	}
	return token, nil
}
