package controllers

import (
	"Prioritas2/middleware"
	"Prioritas2/models"
	"Prioritas2/repositories"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
)

func LoginController(c echo.Context) error {
	var user models.User
	c.Bind(&user)
	_, err := repositories.Login(user.Password, user.Email)
	if err == gorm.ErrRecordNotFound {
		c.JSON(http.StatusBadRequest, models.BaseResponse{
			Status:  false,
			Message: "Email/Password not Found",
			Data:    nil,
		})
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Status:  false,
			Message: "Failed to Authenticate",
			Data:    nil,
		})
	}

	tokenResult := middleware.GenerateJWT(user.ID, user.Name)

	var response models.UserAuthResponse
	response.ID = user.ID
	response.Email = user.Email
	response.Name = user.Name
	response.Token = tokenResult

	return c.JSON(http.StatusOK, models.BaseResponse{
		Status:  true,
		Message: "Success Login",
		Data:    response,
	})
}
