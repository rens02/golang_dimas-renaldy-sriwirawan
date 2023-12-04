package controllers

import (
	"Prioritas2/config"
	"Prioritas2/models"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

// get all users
func GetUsersController(c echo.Context) error {
	var users []models.User

	result := config.DB.Find(&users)

	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, models.BaseResponse{
			Status:  false,
			Message: "Failed to get all users",
			Data:    nil,
		})
	}
	return c.JSON(http.StatusOK, models.BaseResponse{
		Status:  true,
		Message: "Success get all users",
		Data:    users,
	})
}

// get user by id
func GetUserController(c echo.Context) error {
	userID, _ := strconv.Atoi(c.Param("id"))

	// Fetch the user by their ID from the database using GORM
	var user models.User
	result := config.DB.Preload("Blogs").First(&user, userID)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, models.BaseResponse{
				Status:  false,
				Message: "ID Not Found",
				Data:    nil,
			})
		}
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Status:  false,
			Message: "Failed to fetch ID",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, models.BaseResponse{
		Status:  true,
		Message: "Success",
		Data:    user,
	})
}

// create new user
func CreateUserController(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)

	if err := config.DB.Save(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, models.BaseResponse{
		Status:  true,
		Message: "Success",
		Data:    user,
	})
}

// delete user by id
func DeleteUserController(c echo.Context) error {
	userID, _ := strconv.Atoi(c.Param("id"))

	// Fetch the complaint by its ID from the database using GORM
	var foundID models.User
	result := config.DB.Delete(&foundID, userID)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, models.BaseResponse{
				Status:  false,
				Message: "ID Not Found",
				Data:    nil,
			})
		}
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Status:  false,
			Message: "Failed to fetch ID",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, models.BaseResponse{
		Status:  true,
		Message: "User Deleted",
		Data:    nil,
	})
}

// update user by id
func UpdateUserController(c echo.Context) error {
	var updatedUser models.User

	if err := c.Bind(&updatedUser); err != nil {
		return c.JSON(http.StatusBadRequest, models.BaseResponse{
			Status:  false,
			Message: "Invalid request body",
			Data:    nil,
		})
	}

	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.BaseResponse{
			Status:  false,
			Message: "Invalid user ID",
			Data:    nil,
		})
	}

	// Fetch the existing user
	existingUser := models.User{}
	err = config.DB.First(&existingUser, userID).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Status:  false,
			Message: "Failed to fetch user",
			Data:    nil,
		})
	}

	// Update user information
	existingUser.Email = updatedUser.Email
	existingUser.Name = updatedUser.Name
	existingUser.Password = updatedUser.Password

	// Save the updated user
	err = config.DB.Save(&existingUser).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Status:  false,
			Message: "Failed to update user",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, models.BaseResponse{
		Status:  true,
		Message: "User updated successfully",
		Data:    existingUser,
	})
}
