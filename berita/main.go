package main

import (
	"beritaalta/controllers"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	e := echo.New()
	e.GET("/news", controllers.GetNewsController())
}
