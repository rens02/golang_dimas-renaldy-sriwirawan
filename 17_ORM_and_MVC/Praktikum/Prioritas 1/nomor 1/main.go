package main

import (
	"Prioritas_1/models"
	"fmt"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

package main


import (
"fmt"
"net/http"
"strconv"


"github.com/jinzhu/gorm"
_ "github.com/jinzhu/gorm/dialects/mysql"
"github.com/labstack/echo"
)


var (
	DB *gorm.DB
)


func init() {
	InitDB()
	InitialMigration()
}


type Config struct {
	DB_Username string
	DB_Password string
	DB_Port     string
	DB_Host     string
	DB_Name     string
}


func InitDB() {


	config := Config{
		DB_Username: "root",
		DB_Password: "root123",
		DB_Port:     "3306",
		DB_Host:     "localhost",
		DB_Name:     "crud_go",
	}


	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.DB_Username,
		config.DB_Password,
		config.DB_Host,
		config.DB_Port,
		config.DB_Name,
	)


	var err error
	DB, err = gorm.Open("mysql", connectionString)
	if err != nil {
		panic(err)
	}
}


type User struct {
	gorm.Model
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}


func InitialMigration() {
	DB.AutoMigrate(&User{})
}


// get all users
func GetUsersController(c echo.Context) error {
	var users []User


	if err := DB.Find(&users).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all users",
		"users":   users,
	})
}


// get user by id
func GetUserController(c echo.Context) error {
	// your solution here
	userID, _ := strconv.Atoi(c.Param("id"))

	// Fetch the complaint by its ID from the database using GORM
	var foundID models.User
	result := DB.First(&foundID, userID)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				Status:  false,
				Message: "ID Not Found",
				Data:    nil,
			})
		}
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			Status:  false,
			Message: "Failed to fetch ID",
			Data:    nil,
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		Status:  true,
		Message: "Success",
		Data:    foundID,
	})
}


// create new user
func CreateUserController(c echo.Context) error {
	user := User{}
	c.Bind(&user)


	if err := DB.Save(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new user",
		"user":    user,
	})
}


func DeleteUserController(c echo.Context) error {
	userID, _ := strconv.Atoi(c.Param("id"))

	// Fetch the complaint by its ID from the database using GORM
	var foundID models.User
	result := DB.Delete(&foundID, userID)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				Status:  false,
				Message: "ID Not Found",
				Data:    nil,
			})
		}
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			Status:  false,
			Message: "Failed to fetch ID",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		Status:  true,
		Message: "User Deleted",
		Data:    nil,
	})
}

// update user by id
func UpdateUserController(c echo.Context) error {
	var updatedUser models.User

	if err := c.Bind(&updatedUser); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			Status:  false,
			Message: "Invalid request body",
			Data:    nil,
		})
	}

	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			Status:  false,
			Message: "Invalid user ID",
			Data:    nil,
		})
	}

	// Fetch the existing user
	existingUser := models.User{}
	err = DB.First(&existingUser, userID).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
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
	err = DB.Save(&existingUser).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			Status:  false,
			Message: "Failed to update user",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		Status:  true,
		Message: "User updated successfully",
		Data:    existingUser,
	})
}


func main() {
	// create a new echo instance
	e := echo.New()
	// Route / to handler function
	e.GET("/users", GetUsersController)
	e.GET("/users/:id", GetUserController)
	e.POST("/users", CreateUserController)
	e.DELETE("/users/:id", DeleteUserController)
	e.PUT("/users/:id", UpdateUserController)


	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8000"))
}