package http

import (
	db "capston-lms/internal/adapters/db/mysql"
	handler "capston-lms/internal/adapters/http/handler"
	middlewares "capston-lms/internal/adapters/http/middleware"
	repository "capston-lms/internal/adapters/repository"
	usecase "capston-lms/internal/application/usecase"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var (
	// user management
	userRepo    repository.UserRepository
	userHandler handler.UserHandler
	userUsecase usecase.UserUseCase
	// auth
	AuthHandler handler.AuthHandler
)

func declare() {
	userRepo = repository.UserRepository{DB: db.DbMysql}
	userUsecase = usecase.UserUseCase{Repo: userRepo}
	userHandler = handler.UserHandler{UserUsecase: userUsecase}
	// auth
	AuthHandler = handler.AuthHandler{Usecase: userUsecase}

}

func InitRoutes() *echo.Echo {
	db.Init()
	declare()

	e := echo.New()
	e.POST("/login", AuthHandler.Login())
	e.POST("/register", AuthHandler.Register())

	// montor group
	admin := e.Group("/admin")
	admin.Use(middleware.Logger())
	admin.Use(middlewares.AuthMiddleware())

	admin.GET("/users", userHandler.GetAllUsers())
	admin.GET("/users/:id", userHandler.GetUser())
	admin.POST("/users", userHandler.CreateUser())
	admin.DELETE("/users/:id", userHandler.DeleteUser())
	return e
}
