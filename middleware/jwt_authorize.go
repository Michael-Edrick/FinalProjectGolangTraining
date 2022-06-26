package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"
)

type Claims struct {
	Email string `json:"email"`
}

var mySecretKey = []byte("secretkey")

func IsAuthorized ()func(handler http.Handler) http.Handler {
	return func(handler http.Handler) http.Handler {
		return http.HandlerFunc (func(w http.ResponseWriter, r *http.Request){
			header:= r.Header.Get("Authorization")
			splitToken := strings.Split(header, "Bearer ")
			header = splitToken[1]
			fmt.Println(header)
			handler.ServeHTTP(w, r)

			token, err := jwt.Parse(header, func(token *jwt.Token) (interface{}, error) {
				// Don't forget to validate the alg is what you expect:
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
				}
			// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
			return mySecretKey, nil
			})
			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				fmt.Println(claims["foo"], claims["nbf"])
			} else {
				fmt.Println(err)
			}	
		})
	}
}

