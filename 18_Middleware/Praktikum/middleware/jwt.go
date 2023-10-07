package middleware

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type jwtCustomClaims struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	jwt.RegisteredClaims
}

func GenerateJWT(id uint, name string) string {

	var payloadData jwtCustomClaims

	payloadData.ID = id
	payloadData.Name = name
	payloadData.ExpiresAt = jwt.NewNumericDate(time.Now().Add(time.Minute * 60))

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payloadData)
	t, _ := token.SignedString([]byte("1234"))
	return t
}
