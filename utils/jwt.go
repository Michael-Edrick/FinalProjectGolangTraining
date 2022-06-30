package utils

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
)

var mySecretKey = []byte("secretkey")

func GenerateJWT(id int) (string, error) {
	config := InitConfig()
	var mySigningKey = []byte(config.SecretKey)
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["userid"] = id
	claims["exp"] = time.Now().Add(24 * time.Hour).Unix()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		fmt.Printf("Something Went Wrong: %s", err.Error())
		return "", err
	}
	return tokenString, nil
}

func ParseJWT(header string) (jwt.MapClaims, error) {
	splitToken := strings.Split(header, "Bearer ")
	if len(splitToken) < 2 {
		return nil, errors.New("token error")
	}
	header = splitToken[1]
	fmt.Printf("token:%v\n", header)
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
		return claims, nil
	} else {
		fmt.Println(err)
		return nil, errors.New("token parse error")
	}
}
