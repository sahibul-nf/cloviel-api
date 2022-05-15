package main

import (
	"cloviel-api/auth"
	"cloviel-api/config"
	"cloviel-api/event"
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
	authService = auth.NewService()
	userHandler = handler.NewUserHandler(userService, authService)

	// event endpoint
	eventRepo    = event.NewRepository(db)
	eventService = event.NewService(eventRepo)
	eventHandler = handler.NewEventHandler(eventService)
)

func main() {
	defer config.CloseDatabaseConnection(db)

	server := gin.Default()
	server.Use(middleware.CORSMiddleware())

	apiV1 := server.Group("/api/v1")
	{
		// user
		apiV1.POST("/users", userHandler.RegisterUser)
		apiV1.POST("/users/login", userHandler.LoginUser)
		apiV1.POST("/users/avatar", middleware.AuthMiddleware(authService, userService), userHandler.UploadAvatar)

		// event
		apiV1.POST("/events", middleware.AuthMiddleware(authService, userService), eventHandler.CreateNewEvent)
		apiV1.PUT("/events/:id", middleware.AuthMiddleware(authService, userService), eventHandler.UpdateEvent)
		apiV1.POST("/events/company", middleware.AuthMiddleware(authService, userService), eventHandler.CreateNewCompany)
		apiV1.POST("/events/company/logo", middleware.AuthMiddleware(authService, userService), eventHandler.UploadCompanyLogo)
	}

	server.Run()
}
