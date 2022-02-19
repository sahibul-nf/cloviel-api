package handler

import (
	"cloviel-api/auth"
	"cloviel-api/helper"
	"cloviel-api/user"
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
	authService auth.Service
}

func NewUserHandler(userService user.Service, authService auth.Service) *userHandler {
	return &userHandler{userService, authService}
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
		response := helper.APIResponse("Register account failed", "error", http.StatusInternalServerError, nil)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	// generate JWT token
	token := "123"

	// return final response to client
	userFormatter := user.FormatUser(newUser, token)
	response := helper.APIResponse("Account has been registered", "success", http.StatusOK, userFormatter)

	c.JSON(http.StatusOK, response)
}

func (h *userHandler) LoginUser(c *gin.Context) {
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

			response := helper.APIResponse("Login failed", "error", http.StatusBadRequest, errorMessage)
			c.JSON(http.StatusBadRequest, response)
			return
		}
	}

	// generate JWT token
	token, err := h.authService.GenerateToken(int(loggedInUser.ID))
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Login failed", "error", http.StatusInternalServerError, errorMessage)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	// return final response to client
	userFormatter := user.FormatUser(loggedInUser, token)
	response := helper.APIResponse("Successfully loggedin", "success", http.StatusOK, userFormatter)

	c.JSON(http.StatusOK, response)
}

func (h *userHandler) UploadAvatar(c *gin.Context) {

	fileExtSupport := []string{
		".png",
		".jpeg",
		".jpg",
	}

	// get input form data
	file, err := c.FormFile("avatar")
	if err != nil {
		data := gin.H{
			"is_uploaded": false,
			"errors":      err.Error(),
		}

		response := helper.APIResponse("Failed to upload avatar image", "error", http.StatusBadRequest, data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// cek file extention
	fileExtension := filepath.Ext(file.Filename)
	if fileExtension != fileExtSupport[0] && fileExtension != fileExtSupport[1] && fileExtension != fileExtSupport[2] {
		data := gin.H{
			"is_uploaded": false,
			"errors":      "The provided file format is not allowed. Please upload a JPEG or PNG image",
		}

		response := helper.APIResponse("Failed to upload avatar image", "error", http.StatusBadGateway, data)
		c.JSON(http.StatusBadGateway, response)
		return
	}

	// get userID from currentUser (middleware)
	// currentUser := c.MustGet("currentUser").(user.User)
	userID := 1

	// make path location for save avatar
	path := fmt.Sprintf("assets/user-avatars/%d-%s", userID, file.Filename)

	// upload avatar ke server
	err = c.SaveUploadedFile(file, path)
	if err != nil {
		data := gin.H{
			"is_uploaded": false,
			"errors":      err.Error(),
		}

		response := helper.APIResponse("Failed to upload avatar image", "error", http.StatusBadGateway, data)
		c.JSON(http.StatusBadGateway, response)
		return
	}

	// passing path to service for save to db
	_, err = h.userService.SaveAvatar(int(userID), path)
	if err != nil {
		data := gin.H{
			"is_uploaded": false,
			"errors":      err.Error(),
		}

		response := helper.APIResponse("Failed to upload avatar image", "error", http.StatusInternalServerError, data)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	data := gin.H{"is_uploaded": true, "errors": nil}

	response := helper.APIResponse("Avatar successfuly uploaded", "success", http.StatusOK, data)
	c.JSON(http.StatusOK, response)
}
