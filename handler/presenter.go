package handler

import (
	"cloviel-api/event"
	"cloviel-api/helper"
	"cloviel-api/presenter"
	"cloviel-api/user"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type presenterHandler struct {
	presenterService presenter.Service
	eventService     event.Service
}

func NewPresenterHandler(presenterService presenter.Service, eventService event.Service) *presenterHandler {
	return &presenterHandler{presenterService, eventService}
}

func (h *presenterHandler) CreateNewPresenter(c *gin.Context) {
	var input presenter.PresenterInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errorFormatter := helper.ErrorValidationFormat(err)
		errorMessage := gin.H{"errors": errorFormatter}

		response := helper.APIResponse("Failed to create presenter of event", "error", http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	_, statusCode, err := h.eventService.GetEvent(input.EventID)
	if err != nil {

		response := helper.APIResponse(err.Error(), "error", statusCode, nil)
		c.JSON(statusCode, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	input.User = currentUser

	newPresenter, err := h.presenterService.CreateNewPresenter(input)
	if err != nil {
		errorFormatter := helper.ErrorValidationFormat(err)
		errorMessage := gin.H{"errors": errorFormatter}

		response := helper.APIResponse("Failed to create presenter of event", "error", http.StatusInternalServerError, errorMessage)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	responseFormatter := presenter.FormatPresenter(newPresenter)

	// return response to client
	response := helper.APIResponse("Successfully to create presenter of event", "success", http.StatusOK, responseFormatter)
	c.JSON(http.StatusOK, response)
}

func (h *presenterHandler) UpdatePresenter(c *gin.Context) {
	var inputID presenter.PresenterDetailInput

	err := c.ShouldBindUri(&inputID)
	if err != nil {
		response := helper.APIResponse("Failed to update presenter of event", "error", http.StatusBadRequest, nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var inputData presenter.PresenterInput

	err = c.ShouldBindJSON(&inputData)
	if err != nil {
		errorFormatter := helper.ErrorValidationFormat(err)
		errorMessage := gin.H{"errors": errorFormatter}

		response := helper.APIResponse("Failed to update presenter of event", "error", http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	event, statusCode, err := h.eventService.GetEvent(inputData.EventID)
	if err != nil {
		response := helper.APIResponse(err.Error(), "error", statusCode, nil)
		c.JSON(statusCode, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	inputData.User = currentUser

	if event.UserID != inputData.User.ID {
		response := helper.APIResponse(errors.New("not an owner of the event").Error(), "error", statusCode, nil)
		c.JSON(statusCode, response)
		return
	}

	updatedPresenter, statusCode, err := h.presenterService.UpdatePresenter(inputID.ID, inputData)
	if err != nil {
		response := helper.APIResponse(err.Error(), "error", statusCode, nil)
		c.JSON(statusCode, response)
		return
	}

	responseFormatter := presenter.FormatPresenter(updatedPresenter)

	// return response to client
	response := helper.APIResponse("Successfully to update presenter of event", "success", http.StatusOK, responseFormatter)
	c.JSON(http.StatusOK, response)
}
