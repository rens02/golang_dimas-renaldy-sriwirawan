package controllers

import (
	"beritaalta/config"
	"beritaalta/model/news"
	"github.com/labstack/echo/v4"
	"net/http"
)

func GetNewsController(e echo.Context) error {
	var news []news.News
	result := config.Db.Find(&news)
	if result.Error != nil {
		return e.JSON(http.StatusInternalServerError, nil)
	}
	return e.JSON(http.StatusOK, news)
}
