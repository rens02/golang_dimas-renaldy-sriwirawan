package main

import (
	"beritaalta/config"
	"beritaalta/controllers"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	config.Init()
	e.GET("/news", controllers.GetNewsController)
}
