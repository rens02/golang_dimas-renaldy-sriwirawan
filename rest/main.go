package main

import (
	"net/http"
	"strconv"
	"github.com/labstack/echo"
)

type User struct {
	Id    int
	Age   int
	Email string
	name  string
}

func main() {
	e := echo.New()

	//routing
	e.GET("/hello", HelloController)
	e.GET("/user", UserController)
	e.GET("/user/:id/:age", UserIDController)

	e.Start(":8000")
}

// controller UserController single
func UserController(e echo.Context) error {

	user := User{Email: "dimas@alta.com", name: "Dimas Renaldy Sriwirawan"}
	return e.JSON(http.StatusOK, user)
}

// controller UserIDController multi
func UserIDController(e echo.Context) error {

	id, _ := strconv.Atoi(e.Param("id"))
	age, _ := strconv.Atoi(e.Param("age"))

	user := User{Id: id, Age: age, Email: "dimas@alta.com", name: "Dimas Renaldy Sriwirawan"}
	return e.JSON(http.StatusOK, user)
}

// controller HelloController
func HelloController(e echo.Context) error {
	return e.JSON(http.StatusOK, "Hello World")
}

// localhost:8000/user

// controller getUser
