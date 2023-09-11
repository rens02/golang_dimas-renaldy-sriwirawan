package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	//routing
	e.GET("/hello", HelloController)

	e.Start(":8000")
}

// controller HelloController
func HelloController(e echo.Context) error {
	return e.String(http.StatusOK, "Hello World")
}

// localhost:8000/user

// controller getUser
