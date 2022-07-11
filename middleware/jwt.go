package middleware

import (
	"time"

	"github.com/golang-jwt/jwt"
)

func CreateToken(id int, name, secret string) (string, error) {
	claims := jwt.MapClaims{}
	claims["exp"] = time.Now().Add((168 * time.Hour)).Unix()
	claims["iat"] = time.Now().Unix()
	// claims["nbf"] = time.Now().Add((25 * time.Second)).Unix()
	claims["name"] = name
	claims["id"] = id

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}