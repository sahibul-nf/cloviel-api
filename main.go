package main

import (
	"cloviel-api/auth"
	"cloviel-api/company"
	"cloviel-api/config"
	"cloviel-api/event"
	"cloviel-api/handler"
	"cloviel-api/middleware"
	"cloviel-api/presenter"
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

	// presenter endpoint
	presenterRepo    = presenter.NewRepository(db)
	presenterService = presenter.NewService(presenterRepo)
	presenterHandler = handler.NewPresenterHandler(presenterService, eventService)
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
		apiV1.GET("/events", eventHandler.GetAllEvent)
		apiV1.GET("/events/:id", eventHandler.GetEventDetails)
		apiV1.POST("/events", middleware.AuthMiddleware(authService, userService), eventHandler.CreateNewEvent)
		apiV1.POST("/events/thumbnail", middleware.AuthMiddleware(authService, userService), eventHandler.UploadEventThumbnail)
		apiV1.POST("/events/signature", middleware.AuthMiddleware(authService, userService), eventHandler.UploadEventSignature)
		apiV1.PUT("/events/:id", middleware.AuthMiddleware(authService, userService), eventHandler.UpdateEvent)
		apiV1.DELETE("/events/:id", middleware.AuthMiddleware(authService, userService), eventHandler.DeleteEvent)

		// company of event
		apiV1.POST("/events/companies", middleware.AuthMiddleware(authService, userService), companyHandler.CreateNewCompany)
		apiV1.POST("/events/companies/logo", middleware.AuthMiddleware(authService, userService), companyHandler.UploadCompanyLogo)

		// presenter of event
		apiV1.POST("/events/presenters", middleware.AuthMiddleware(authService, userService), presenterHandler.CreateNewPresenter)
	}

	server.Run()
}
