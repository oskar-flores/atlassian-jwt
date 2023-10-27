package main

import (
	"flag"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

func createToken(issuer string, signingKey string) (string, error) {
	// Create a new token object, specifying signing method and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss": issuer,
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	})

	// Sign and get the complete encoded token as a string
	tokenString, err := token.SignedString([]byte(signingKey))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func main() {
	clientKey := flag.String("clientKey", "", "Issuer")
	sharedSecret := flag.String("secret", "", "Signing Key")

	flag.Parse()

	if *clientKey == "" || *sharedSecret == "" {
		fmt.Println("Usage: ./atlassian-jwt -clientKey=<clientKey> -key=<sharedSecret>")
		return
	}

	tokenString, err := createToken(*clientKey, *sharedSecret)
	if err != nil {
		fmt.Println("Error creating token: ", err)
		return
	}
	fmt.Println(tokenString)
}
