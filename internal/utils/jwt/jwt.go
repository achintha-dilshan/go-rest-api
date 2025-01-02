package jwt

import (
	"time"

	"github.com/achintha-dilshan/go-rest-api/config"
	"github.com/golang-jwt/jwt/v5"
)

// generate token
func GenerateToken(id int64) (string, error) {
	exp := time.Now().Add(15 * time.Minute).Unix()
	secret := []byte(config.Env.JWTSecret)

	// token claims
	claims := jwt.MapClaims{
		"id":  id,
		"exp": exp,
	}

	// create the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// sign the toke with secret key
	return token.SignedString(secret)
}
