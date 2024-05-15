package utils

import "github.com/dgrijalva/jwt-go"

var Secret_Key = "secret"

func GenerateToken(claims *jwt.StandardClaims) (string, error) {
	tokens := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	jwt, err := tokens.SignedString([]byte(Secret_Key))
	if err != nil {
		return "", err
	}

	return jwt, nil
}
