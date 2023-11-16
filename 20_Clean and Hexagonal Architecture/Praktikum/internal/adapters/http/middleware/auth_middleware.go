package middleware

import (
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
)

func CreateToken(userID int, email string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = userID
	claims["email"] = email
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	secret := viper.GetString("SECRET_JWT")
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func AuthMiddleware() echo.MiddlewareFunc {
	return middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(viper.GetString("SECRET_JWT")),
		ErrorHandler: func(err error) error {
			if _, ok := err.(*jwt.ValidationError); ok {
				return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token")
			}
			return err
		},
		Claims: &jwt.MapClaims{},
	})
}
