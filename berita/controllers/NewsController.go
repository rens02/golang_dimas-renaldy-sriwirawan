package controllers

import (
	"beritaalta/config"
	"beritaalta/model/berita"
	"github.com/labstack/echo/v4"
	"net/http"
)

func GetNewsController(e echo.Context) error {
	var news []berita.News
	result := config.Db.Find(&news)
	if result.Error != nil {
		return e.JSON(http.StatusInternalServerError, nil)
	}
	return e.JSON(http.StatusOK, news)
}

func PostNewsController(e echo.Context) error {
	// Parse the request body into a News struct
	var news berita.News
	if err := e.Bind(&news); err != nil {
		return e.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	// Save the news article to the database
	result := config.Db.Create(&news)
	if result.Error != nil {
		return e.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create news article"})
	}

	// Return the created news article as JSON response
	return e.JSON(http.StatusCreated, news)
}
