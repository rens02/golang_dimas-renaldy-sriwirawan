package main

import (
	"Prioritas2/controllers"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"os"
)

func main() {
	secretKey := os.Getenv("JWT_SECRET_KEY")
	// create a new echo instance
	e := echo.New()
	// Route / to handler function

	// Create new user (not Authenticated)
	e.POST("/users", controllers.CreateUserController)
	// Login
	e.POST("/users/login", controllers.LoginController)

	//JWT Authenticated
	eAuth := e.Group("")
	eAuth.Use(echojwt.JWT([]byte(secretKey)))
	// Users Authenticated
	eAuth.GET("/users", controllers.GetUsersController)
	eAuth.GET("/users/:id", controllers.GetUserController)
	eAuth.DELETE("/users/:id", controllers.DeleteUserController)
	eAuth.PUT("/users/:id", controllers.UpdateUserController)
	// Books Authenticated
	eAuth.GET("/books", controllers.GetBooksController)
	eAuth.GET("/books/:id", controllers.GetBookController)
	eAuth.POST("/books", controllers.CreateBookController)
	eAuth.DELETE("/books/:id", controllers.DeleteBookController)
	eAuth.PUT("/books/:id", controllers.UpdateBookController)

	e.GET("/blogs", controllers.GetAllBlog)
	e.GET("/blogs/:id", controllers.GetBlogByID)
	e.POST("/blogs", controllers.CreateBlog)
	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8000"))
}
