package handler

import (
	"cloviel-api/event"
	"cloviel-api/helper"
	"cloviel-api/user"
	"fmt"
	"net/http"
	"path/filepath"
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

	responseFormatter := event.FormatDetailEvent(newEvent)

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

	responseFormatter := event.FormatDetailEvent(updatedEvent)

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

func (h *eventHandler) GetEventDetails(c *gin.Context) {
	var inputID event.GetEventDetailInput

	err := c.ShouldBindUri(&inputID)
	if err != nil {
		response := helper.APIResponse("Failed to get event details", "error", http.StatusBadRequest, nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	eventDetail, statusCode, err := h.eventService.GetEvent(inputID.ID)
	if err != nil {

		response := helper.APIResponse(err.Error(), "error", statusCode, nil)
		c.JSON(statusCode, response)
		return
	}

	responseFormatter := event.FormatDetailEvent(eventDetail)

	response := helper.APIResponse("Successfuly get event details", "success", http.StatusOK, responseFormatter)
	c.JSON(http.StatusOK, response)
}

func (h *eventHandler) UploadEventThumbnail(c *gin.Context) {
	fileExtSupport := helper.FileExtSupport()

	var input event.SaveEventImageInput

	err := c.ShouldBind(&input)
	if err != nil {
		errorFormatter := helper.ErrorValidationFormat(err)

		errorMessage := gin.H{
			"is_uploaded": false,
			"errors":      errorFormatter,
		}

		response := helper.APIResponse("Failed to upload the thumbnail image of event", "error", http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// get input form data
	file, err := c.FormFile("thumbnail")
	if err != nil {
		data := gin.H{
			"is_uploaded": false,
			"errors":      err.Error(),
		}

		response := helper.APIResponse("Failed to upload the thumbnail image of event", "error", http.StatusBadRequest, data)
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

		response := helper.APIResponse("Failed to upload the thumbnail image of event", "error", http.StatusUnsupportedMediaType, data)
		c.JSON(http.StatusUnsupportedMediaType, response)
		return
	}

	event, statusCode, err := h.eventService.GetEvent(input.EventID)
	if err != nil {
		data := gin.H{
			"is_uploaded": false,
			"errors":      err.Error(),
		}

		response := helper.APIResponse("Failed to upload the thumbnail image of event", "error", statusCode, data)
		c.JSON(statusCode, response)
		return
	}

	// get userID from currentUser (middleware)
	currentUser := c.MustGet("currentUser").(user.User)
	userID := currentUser.ID

	// make path location for save avatar
	path := fmt.Sprintf("assets/event-images/thumbnail/e%du%d-%s", input.EventID, userID, file.Filename)

	// upload avatar ke server
	err = c.SaveUploadedFile(file, path)
	if err != nil {
		data := gin.H{
			"is_uploaded": false,
			"errors":      err.Error(),
		}

		response := helper.APIResponse("Failed to upload the thumbnail image of event", "error", http.StatusInternalServerError, data)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	// passing path to service for save to db
	_, statusCode, err = h.eventService.SaveEventThumbnail(event, path)
	if err != nil {
		data := gin.H{
			"is_uploaded": false,
			"errors":      err.Error(),
		}

		response := helper.APIResponse("Failed to upload the thumbnail image of event", "error", statusCode, data)
		c.JSON(statusCode, response)
		return
	}

	data := gin.H{"is_uploaded": true}

	response := helper.APIResponse("Event thumbnail successfuly uploaded", "success", http.StatusOK, data)
	c.JSON(http.StatusOK, response)
}

func (h *eventHandler) UploadEventSignature(c *gin.Context) {
	fileExtSupport := helper.FileExtSupport()

	var input event.SaveEventImageInput

	err := c.ShouldBind(&input)
	if err != nil {
		errorFormatter := helper.ErrorValidationFormat(err)

		errorMessage := gin.H{
			"is_uploaded": false,
			"errors":      errorFormatter,
		}

		response := helper.APIResponse("Failed to upload the signature image of event", "error", http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// get input form data
	file, err := c.FormFile("signature")
	if err != nil {
		data := gin.H{
			"is_uploaded": false,
			"errors":      err.Error(),
		}

		response := helper.APIResponse("Failed to upload the signature image of event", "error", http.StatusBadRequest, data)
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

		response := helper.APIResponse("Failed to upload the signature image of event", "error", http.StatusUnsupportedMediaType, data)
		c.JSON(http.StatusUnsupportedMediaType, response)
		return
	}

	event, statusCode, err := h.eventService.GetEvent(input.EventID)
	if err != nil {
		data := gin.H{
			"is_uploaded": false,
			"errors":      err.Error(),
		}

		response := helper.APIResponse("Failed to upload the signature image of event", "error", statusCode, data)
		c.JSON(statusCode, response)
		return
	}

	// get userID from currentUser (middleware)
	currentUser := c.MustGet("currentUser").(user.User)
	userID := currentUser.ID

	// make path location for save avatar
	path := fmt.Sprintf("assets/event-images/signature/e%du%d-%s", input.EventID, userID, file.Filename)

	// upload avatar ke server
	err = c.SaveUploadedFile(file, path)
	if err != nil {
		data := gin.H{
			"is_uploaded": false,
			"errors":      err.Error(),
		}

		response := helper.APIResponse("Failed to upload the signature image of event", "error", http.StatusInternalServerError, data)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	// passing path to service for save to db
	_, err = h.eventService.SaveEventSignature(event, path)
	if err != nil {
		data := gin.H{
			"is_uploaded": false,
			"errors":      err.Error(),
		}

		response := helper.APIResponse("Failed to upload the signature image of event", "error", http.StatusInternalServerError, data)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	data := gin.H{"is_uploaded": true}

	response := helper.APIResponse("Event signature successfuly uploaded", "success", http.StatusOK, data)
	c.JSON(http.StatusOK, response)
}
