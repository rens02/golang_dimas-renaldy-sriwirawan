package controllers

import (
	"code_structure/config"
	"code_structure/middlewares"
	"code_structure/models"
	"net/http"

	"github.com/labstack/echo"
)

// get all users
func GetUsersController(c echo.Context) error {
	var users []models.User

	err := config.DB.Find(&users).Error;
	if  err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": http.StatusOK,
		"message": "success get all users",
		"users":   users,
	})
}
// get user by id
func GetUserController(c echo.Context) error {
	user := models.User{}
	userID := c.Param("id")
	err := config.DB.First(&user, userID).Error;
	if  err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": http.StatusOK,
		"message": "success get user by id",
		"user":    user,
	})
}
// create new user
func CreateUserController(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)


	if err := config.DB.Save(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": http.StatusOK,
		"message": "success create new user",
		"user":    user,
	})
}
// delete user by id
func DeleteUserController(c echo.Context) error {
	user := models.User{}
	userID := c.Param("id")

	if err := config.DB.First(&user, userID).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := config.DB.Delete(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": http.StatusOK,
		"message": "success delete user by id",
	})
}
// update user by id
func UpdateUserController(c echo.Context) error {
	user := models.User{}
	userID := c.Param("id")

	if err := config.DB.First(&user, userID).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.Bind(&user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := config.DB.Save(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": http.StatusOK,
		"message": "success update user by id",
		"user":    user,
	})
}
func LoginUserController(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)

	err := config.DB.Where("email = ? AND password = ?",user.Email, user.Password).First(&user).Error
	if  err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message" : "fail login",
			"error" : err.Error(),
		})
	}

	token, err := middlewares.CreateToken(user.ID, user.Name)
	if err!=nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message" : "fail login",
			"error" : err.Error(),
		})
	}
	userResponse := models.UserResponse{user.ID, user.Name, user.Email, token}
    return c.JSON(http.StatusOK, map[string]interface{}{
        "status":  http.StatusOK,
        "message": "success login",
        "user":    userResponse,
    })
}