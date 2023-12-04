package middleware

import (
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

type jwtCustomClaims struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	jwt.RegisteredClaims
}

func GenerateJWT(id uint, name string) string {
	secretKey := os.Getenv("JWT_SIGNING_KEY")

	var payloadData jwtCustomClaims

	payloadData.ID = id
	payloadData.Name = name
	payloadData.ExpiresAt = jwt.NewNumericDate(time.Now().Add(time.Minute * 60))

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payloadData)
	t, _ := token.SignedString([]byte(secretKey))
	return t
}
