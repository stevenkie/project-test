package util

import (
	"math/rand"
	"time"

	"github.com/dgrijalva/jwt-go"
)

//GenerateToken for login process, this will be used as jwt token for protected endpoints (use in header Authorization)
func GenerateToken(secret string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"random":     rand.Int(),
		"created_at": time.Now().UTC(),
	})
	return token.SignedString([]byte(secret))
}
