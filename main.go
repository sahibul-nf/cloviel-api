package main

import (
	"cloviel-api/auth"
	"cloviel-api/company"
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

	// company endpoint
	companyRepo    = company.NewRepository(db)
	companyService = company.NewService(companyRepo)
	companyHandler = handler.NewCompanyHandler(companyService)
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
		apiV1.DELETE("/events/:id", middleware.AuthMiddleware(authService, userService), eventHandler.DeleteEvent)

		// company
		apiV1.POST("/companies", middleware.AuthMiddleware(authService, userService), companyHandler.CreateNewCompany)
		apiV1.POST("/companies/logo", middleware.AuthMiddleware(authService, userService), companyHandler.UploadCompanyLogo)
	}

	server.Run()
}
