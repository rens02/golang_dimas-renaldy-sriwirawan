package handler

import (
	"net/http"

	"capston-lms/internal/adapters/http/middleware"
	"capston-lms/internal/application/service"
	"capston-lms/internal/application/usecase"

	"capston-lms/internal/entity"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct {
	Usecase usecase.UserUseCase
}

func (handler AuthHandler) Register() echo.HandlerFunc {
	return func(e echo.Context) error {
		var user entity.User
		if err := e.Bind(&user); err != nil {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"status code": http.StatusBadRequest,
				"message": err.Error(),
			})
		}

		// Validasi input menggunakan package validator
		validate := validator.New()
		if err := validate.Struct(user); err != nil {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"status code": http.StatusBadRequest,
				"message": err.Error(),
			})
		}

		// Validasi email unik
		if err := handler.Usecase.UniqueEmail(user.Email); err != nil {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"status code": http.StatusBadRequest,
				"message": err.Error(),
			})
		}

		hashedPassword, err := service.Encrypt(user.Password)
		if err != nil {
			return e.JSON(http.StatusInternalServerError, map[string]interface{}{
				"status code": http.StatusInternalServerError,
				"message": err.Error(),
			})
		}
		user.Password = string(hashedPassword)

		err = handler.Usecase.CreateUser(user)
		if err != nil {
			return e.JSON(http.StatusInternalServerError, map[string]interface{}{
				"status code": http.StatusInternalServerError,
				"message": err.Error(),
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"status code": http.StatusOK,
			"message": "register success",
			"data":   user,
		})
	}
}

func (handler AuthHandler) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		// Bind request body to user struct
		var user entity.User
		if err := c.Bind(&user); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"status code": http.StatusBadRequest,
				"message": "Invalid request body",
			})
		}

		// Get user by email
		dbUser, err := handler.Usecase.GetUserByEmail(user.Email)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"status code": http.StatusBadRequest,
				"message": "Invalid email or password",
			})
		}

		// Check password
		if err := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password)); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"status code": http.StatusBadRequest,
				"message": "Invalid email or password",
			})
		}

		t, err := middleware.CreateToken(int(dbUser.ID), dbUser.Email)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"status code": http.StatusBadRequest,
				"message": "Failed to create token",
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"status code": http.StatusOK,
			"message": "login success",
			"token": t,
		})
	}
}
