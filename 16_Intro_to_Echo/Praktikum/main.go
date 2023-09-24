package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type User struct {
	Id       int    `json:"id" form:"id"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

var users []User

// -------------------- controller --------------------

// get all users
func GetUsersController(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get all users",
		"users":    users,
	})
}

// get user by id
func GetUserController(c echo.Context) error {
	// Get the user ID from the URL parameter
	idParam := c.Param("id")
	userID, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid user ID"})
	}

	// Find the user with the specified ID
	var foundUser User
	userFound := false
	for _, user := range users {
		if user.Id == userID {
			foundUser = user
			userFound = true
			break
		}
	}

	if userFound {
		// Return the found user as a JSON response
		return c.JSON(http.StatusOK, foundUser)
	} else {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "User not found"})
	}

}

// delete user by id
func DeleteUserController(c echo.Context) error {
	// Get the user ID from the URL parameter
	idParam := c.Param("id")
	userID, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid user ID"})
	}

	// Find the index of the user with the specified ID
	userIndex := -1
	for i, user := range users {
		if user.Id == userID {
			userIndex = i
			break
		}
	}

	if userIndex == -1 {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "User not found"})
	}

	// Remove the user from the users slice
	users = append(users[:userIndex], users[userIndex+1:]...)

	return c.JSON(http.StatusOK, map[string]string{"message": "User deleted successfully"})
}

// update user by id
func UpdateUserController(c echo.Context) error {
	// Get the user ID from the URL parameter
	idParam := c.Param("id")
	userID, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid user ID"})
	}

	// Parse the request body into a User struct for updates
	var updatedUser User
	if err := c.Bind(&updatedUser); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	// Find the index of the user with the specified ID
	userIndex := -1
	for i, user := range users {
		if user.Id == userID {
			userIndex = i
			break
		}
	}

	if userIndex == -1 {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "User not found"})
	}

	// Update the user
	users[userIndex] = updatedUser

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "User updated successfully",
		"user":    updatedUser,
	})
}

// create new user
func CreateUserController(c echo.Context) error {
	// binding data
	user := User{}
	c.Bind(&user)

	if len(users) == 0 {
		user.Id = 1
	} else {
		newId := users[len(users)-1].Id + 1
		user.Id = newId
	}
	users = append(users, user)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success create user",
		"user":     user,
	})
}

// ---------------------------------------------------
func main() {
	users = []User{
		{Id: 1, Name: "User1", Email: "user1@example.com", Password: "password1"},
		{Id: 2, Name: "User2", Email: "user2@example.com", Password: "password2"},
		{Id: 3, Name: "User3", Email: "user3@example.com", Password: "password3"},
	}
	e := echo.New()
	// routing with query parameter
	e.GET("/users", GetUsersController)
	e.GET("/users/:id", GetUserController)
	e.POST("/users", CreateUserController)

	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8000"))
}
