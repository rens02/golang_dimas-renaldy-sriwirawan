package controllers_test

import (
	"bytes"
	"code_structure/config"
	"code_structure/controllers"
	"code_structure/models"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)
// get all data user
func TestGetUsersController(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Perform test
	if assert.NoError(t, controllers.GetUsersController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		var response struct {
			Status  int             `json:"status"`
			Message string          `json:"message"`
			Users   []models.User   `json:"users"`
		}
		json.Unmarshal(rec.Body.Bytes(), &response)

		assert.Equal(t, http.StatusOK, response.Status)
		assert.Equal(t, "success get all users", response.Message)
		assert.NotNil(t, response.Users)
	}
}

// get user by id
func TestGetUserControllerValid(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/users/:id")
	c.SetParamNames("id")
	c.SetParamValues("38")

	if assert.NoError(t, controllers.GetUserController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		var response struct {
			Status  int           `json:"status"`
			Message string        `json:"message"`
			User    models.User   `json:"user"`
		}
		json.Unmarshal(rec.Body.Bytes(), &response)

		assert.Equal(t, http.StatusOK, response.Status)
		assert.Equal(t, "success get user by id", response.Message)
		assert.NotNil(t, response.User)
	}
}
func TestGetUserControllerInvalid(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/users/:id")
	c.SetParamNames("id")
	c.SetParamValues("2")

	// Assertions
	if assert.Error(t, controllers.GetUserController(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}
}

func TestCreateUserControllerValid(t *testing.T) {
	// Setup
	e := echo.New()
	payload := map[string]interface{}{
		"name":     "dimas",
		"email":    "dimas@gmail.com",
		"password": "12345",
	}
	body, _ := json.Marshal(payload)
	req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Perform test
	if assert.NoError(t, controllers.CreateUserController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		var response struct {
			Status  int           `json:"status"`
			Message string        `json:"message"`
			User    models.User   `json:"user"`
		}
		json.Unmarshal(rec.Body.Bytes(), &response)

		assert.Equal(t, http.StatusOK, response.Status)
		assert.Equal(t, "success create new user", response.Message)
		assert.NotNil(t, response.User)

		// Clean up
		config.DB.Unscoped().Delete(&response.User)
	}
}

func TestDeleteUserControllerValid(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/users/46", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/users/:id")
	c.SetParamNames("id")
	c.SetParamValues("46")

	// Assertions
	if assert.NoError(t, controllers.DeleteUserController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}
func TestDeleteUserControllerInvalid(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/users/1000", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/users/:id")
	c.SetParamNames("id")
	c.SetParamValues("1000")

	// Assertions
	if assert.Error(t, controllers.DeleteUserController(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}
}

func TestUpdateUserControllerValid(t *testing.T) {
	// Setup
	e := echo.New()
	user := models.User{Name: "saya", Email: "saya@gmail.com", Password: "123"}
	config.DB.Save(&user)
	user.Name = "saya"
	reqBody, err := json.Marshal(user)
	assert.NoError(t, err)
	req := httptest.NewRequest(http.MethodPut, "/users/48", bytes.NewReader(reqBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/users/:id")
	c.SetParamNames("id")
	c.SetParamValues("48")

	// Assertions
	if assert.NoError(t, controllers.UpdateUserController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.JSONEq(t, `{"status":200,"message":"success update user by id","user":{"id":48,"name":"saya","email":"saya@gmail.com","password":"123"}}`, rec.Body.String())
	}
}
func TestUpdateUserControllerInvalid(t *testing.T) {
	// Setup
	e := echo.New()
	reqBody, err := json.Marshal(map[string]string{"name": "cobaa update"})
	assert.NoError(t, err)
	req := httptest.NewRequest(http.MethodPut, "/users/1000", bytes.NewReader(reqBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/users/:id")
	c.SetParamNames("id")
	c.SetParamValues("1000")

	// Assertions
	if assert.Error(t, controllers.UpdateUserController(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}
}

func TestLoginUserControllerValid(t *testing.T) {
	// Setup
	e := echo.New()
	reqBody, err := json.Marshal(map[string]string{"email": "sai@gmail.com", "password": "123"})
	assert.NoError(t, err)
	req := httptest.NewRequest(http.MethodPost, "/users/login", bytes.NewReader(reqBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, controllers.LoginUserController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		// assert.JSONEq(t, `{"status":200,"message":"success login","user":{"id":2,"name":"ahmad","email":"ahmad@gmail.com","token":""}}`, rec.Body.String())
	}
}
func TestLoginUserControllerInvalid(t *testing.T)  {
	// Setup
	e := echo.New()
	reqBody, err := json.Marshal(map[string]string{"email": "emailAsal@gmail.com", "password": "123"})
	assert.NoError(t, err)
	req := httptest.NewRequest(http.MethodPost, "/users/login", bytes.NewReader(reqBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, controllers.LoginUserController(c)) {
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		assert.JSONEq(t, `{"message":"fail login","error":"record not found"}`, rec.Body.String())
	}

}
func TestLoginUserControllerInvalidPassword(t *testing.T)  {
	// Setup
	e := echo.New()
	reqBody, err := json.Marshal(map[string]string{"email": "sai@gmail.com", "password": "9856"})
	assert.NoError(t, err)
	req := httptest.NewRequest(http.MethodPost, "/users/login", bytes.NewReader(reqBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, controllers.LoginUserController(c)) {
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		assert.JSONEq(t, `{"message":"fail login","error":"record not found"}`, rec.Body.String())
	}

}