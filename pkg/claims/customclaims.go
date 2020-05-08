package customclaims

import "github.com/dgrijalva/jwt-go"

type MyCustomClaims struct {
	Name    string `json:"name"`
	Picture string `json:"picture"`
	jwt.StandardClaims
}
