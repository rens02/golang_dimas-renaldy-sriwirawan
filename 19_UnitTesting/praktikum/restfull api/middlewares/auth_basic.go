package middlewares

import (
	"code_structure/config"
	"code_structure/models"

	"github.com/labstack/echo"
)

func BasicAuthDB(username, password string, c echo.Context) (bool, error){
	var user models.User
	err := config.DB.Where("email = ? AND password = ?", username, password).First(&user).Error
	if err != nil {
		return false, err
	}
	return true, nil
}