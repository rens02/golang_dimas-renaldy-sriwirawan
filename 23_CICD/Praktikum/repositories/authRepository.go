package repositories

import (
	"Prioritas2/config"
	"Prioritas2/models"
)

func Login(password string, email string) (bool, error) {
	var user models.User
	// SELECT * FROM users WHERE password = "parameter" & email = "parameter";
	result := config.DB.First(&user, "password = ? AND email = ?", password, email)
	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}
