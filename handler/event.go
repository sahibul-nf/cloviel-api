package handler

import (
	"cloviel-api/event"
	"cloviel-api/helper"
	"cloviel-api/user"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type eventHandler struct {
	eventService event.Service
}

func NewEventHandler(eventService event.Service) *eventHandler {
	return &eventHandler{eventService}
}

func (h *eventHandler) CreateNewEvent(c *gin.Context) {
	// mapping input JSON to Eventinput struct
	var input event.EventInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errorFormatter := helper.ErrorValidationFormat(err)
		errorMessage := gin.H{"errors": errorFormatter}

		response := helper.APIResponse("Failed to create event", "error", http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	input.User = currentUser

	// passing EventInput to service
	newEvent, err := h.eventService.CreateNewEvent(input)
	if err != nil {
		errorFormatter := helper.ErrorValidationFormat(err)
		errorMessage := gin.H{"errors": errorFormatter}

		response := helper.APIResponse("Failed to create event", "error", http.StatusInternalServerError, errorMessage)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	responseFormatter := event.FormatEvent(newEvent)

	// return response to client
	response := helper.APIResponse("Successfully to create company", "success", http.StatusOK, responseFormatter)
	c.JSON(http.StatusOK, response)
}

func (h *eventHandler) UpdateEvent(c *gin.Context) {

	var inputID event.GetEventDetailInput

	err := c.ShouldBindUri(&inputID)
	if err != nil {
		response := helper.APIResponse("Failed to update event", "error", http.StatusBadRequest, nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var inputData event.EventInput
	err = c.ShouldBindJSON(&inputData)
	if err != nil {
		errorFormatter := helper.ErrorValidationFormat(err)
		errorMessage := gin.H{"errors": errorFormatter}

		response := helper.APIResponse("Failed to update event", "error", http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	inputData.User = currentUser

	updatedEvent, err := h.eventService.UpdateEvent(inputID, inputData)
	if err != nil {
		errorFormatter := helper.ErrorValidationFormat(err)
		errorMessage := gin.H{"errors": errorFormatter}

		response := helper.APIResponse("Failed to update event", "error", http.StatusInternalServerError, errorMessage)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	responseFormatter := event.FormatEvent(updatedEvent)

	// return response to client
	response := helper.APIResponse("Successfully to update event", "success", http.StatusOK, responseFormatter)
	c.JSON(http.StatusOK, response)
}

func (h *eventHandler) DeleteEvent(c *gin.Context) {
	var inputID event.GetEventDetailInput

	err := c.ShouldBindUri(&inputID)
	if err != nil {
		errorMessage := gin.H{
			"is_deleted": false,
			"errors":     err.Error(),
		}

		response := helper.APIResponse("Failed to delete event", "error", http.StatusBadRequest, errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	userID := currentUser.ID

	isDeleted, statusCode, errMsg := h.eventService.DeleteEvent(inputID.ID, userID)
	if errMsg != "" {
		errorMessage := gin.H{
			"is_deleted": isDeleted,
			"errors":     errMsg,
		}

		response := helper.APIResponse("Failed to delete event", "error", statusCode, errorMessage)
		c.JSON(statusCode, response)
		return
	}

	responseFormatter := gin.H{
		"is_deleted": isDeleted,
	}

	// return response to client
	response := helper.APIResponse("Successfully to delete event", "success", statusCode, responseFormatter)
	c.JSON(statusCode, response)
}

func (h *eventHandler) GetAllEvent(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Query("user_id"))

	events, statusCode, err := h.eventService.GetAllEvent(userID)
	if err != nil {
		eventsFormatter := event.FormatEvents(events)

		response := helper.APIResponse("Failed to get list event", "error", statusCode, eventsFormatter)
		c.JSON(statusCode, response)
		return
	}

	eventsFormatter := event.FormatEvents(events)

	response := helper.APIResponse("Successfuly get list of events", "success", http.StatusOK, eventsFormatter)
	c.JSON(http.StatusOK, response)
}
