package main

import (
	"cloviel-api/config"
	"cloviel-api/handler"
	"cloviel-api/middleware"
	"cloviel-api/user"

	"github.com/gin-gonic/gin"
)

var (
	db = config.DBConnection()

	// user endpoint
	userRepo    = user.NewRepository(db)
	userService = user.NewService(userRepo)
	userHandler = handler.NewUserHandler(userService)
)

func main() {
	defer config.CloseDatabaseConnection(db)

	api := gin.Default()
	api.Use(middleware.CORSMiddleware())

	userEndpoint := api.Group("/api/v1")
	{
		userEndpoint.POST("/users", userHandler.RegisterUser)
		userEndpoint.POST("/users/login", userHandler.LoginUser)
	}

	api.Run()
}
