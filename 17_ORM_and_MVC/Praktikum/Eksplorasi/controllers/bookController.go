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
func GetBooksController(c echo.Context) error {
	var book []models.Books

	result := config.DB.Find(&book)

	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, models.BaseResponse{
			Status:  false,
			Message: "Failed to get all book",
			Data:    nil,
		})
	}
	return c.JSON(http.StatusOK, models.BaseResponse{
		Status:  true,
		Message: "Success get all book",
		Data:    book,
	})
}

// get user by id
func GetBookController(c echo.Context) error {
	BookID, _ := strconv.Atoi(c.Param("id"))

	// Fetch the complaint by its ID from the database using GORM
	var foundID models.Books
	result := config.DB.First(&foundID, BookID)
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
		Data:    foundID,
	})
}

// create new Book
func CreateBookController(c echo.Context) error {
	Book := models.Books{}
	c.Bind(&Book)

	if err := config.DB.Save(&Book).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, models.BaseResponse{
		Status:  true,
		Message: "Success",
		Data:    Book,
	})
}

// delete Book by id
func DeleteBookController(c echo.Context) error {
	BookID, _ := strconv.Atoi(c.Param("id"))

	// Fetch the complaint by its ID from the database using GORM
	var foundID models.Books
	result := config.DB.Delete(&foundID, BookID)
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
		Message: "Book Deleted",
		Data:    nil,
	})
}

// update Book by id
func UpdateBookController(c echo.Context) error {
	var updatedBook models.Books

	if err := c.Bind(&updatedBook); err != nil {
		return c.JSON(http.StatusBadRequest, models.BaseResponse{
			Status:  false,
			Message: "Invalid request body",
			Data:    nil,
		})
	}

	BookID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.BaseResponse{
			Status:  false,
			Message: "Invalid Book ID",
			Data:    nil,
		})
	}

	// Fetch the existing Book
	existingBook := models.Books{}
	err = config.DB.First(&existingBook, BookID).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Status:  false,
			Message: "Failed to fetch Book",
			Data:    nil,
		})
	}

	// Update Book information
	existingBook.Judul = updatedBook.Judul
	existingBook.Penerbit = updatedBook.Penerbit
	existingBook.Penulis = updatedBook.Penulis

	// Save the updated Book
	err = config.DB.Save(&existingBook).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Status:  false,
			Message: "Failed to update Book",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, models.BaseResponse{
		Status:  true,
		Message: "Book updated successfully",
		Data:    existingBook,
	})
}
