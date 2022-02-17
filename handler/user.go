package handler

import (
	"cloviel-api/helper"
	"cloviel-api/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	// mapping input JSON to struct Input
	var input user.RegisterInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errorFormatter := helper.ErrorValidationFormat(err)

		errorMessage := gin.H{"errors": errorFormatter}

		response := helper.APIResponse("Register account failed", "error", http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// send struct Input to service
	newUser, err := h.userService.Register(input)
	if err != nil {
		response := helper.APIResponse("Register account failed", "error", http.StatusBadRequest, nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// generate JWT token
	token := "123"

	// return final response to client
	userFormatter := user.FormatUser(newUser, token)
	response := helper.APIResponse("Account has been registered", "success", http.StatusOK, userFormatter)

	c.JSON(http.StatusOK, response)
}

func (h *userHandler) LoginUser(c *gin.Context)  {
	// mapping input JSON to LoginInput
	var input user.LoginInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errorFormatter := helper.ErrorValidationFormat(err)

		errorMessage := gin.H{"errors": errorFormatter}

		response := helper.APIResponse("Register account failed", "error", http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// passing LoginInput to service
	loggedInUser, err := h.userService.Login(input)
	if err != nil {
		if err != nil {
			errorMessage := gin.H{"errors": err.Error()}
	
			response := helper.APIResponse("Login failed", "error", http.StatusUnprocessableEntity, errorMessage)
			c.JSON(http.StatusUnprocessableEntity, response)
			return
		}
	}

	// generate JWT token

	// return final response to client
	userFormatter := user.FormatUser(loggedInUser, "123")
	response := helper.APIResponse("Successfully loggedin", "success", http.StatusOK, userFormatter)

	c.JSON(http.StatusOK, response)
}
