package handler

import (
	"cloviel-api/event"
	"cloviel-api/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

type eventHandler struct {
	eventService event.Service
}

func NewEventHandler(eventService event.Service) *eventHandler {
	return &eventHandler{eventService}
}

func (h *eventHandler) CreateNewCompany(c *gin.Context) {

	// mapping input JSON ke struct Companyinput
	var input event.CompanyInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errorFormatter := helper.ErrorValidationFormat(err)
		errorMessage := gin.H{"errors": errorFormatter}

		response := helper.APIResponse("Failed to create company", "error", http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// passing CompanyInput to service
	newCompany, err := h.eventService.CreateNewCompany(input)
	if err != nil {
		errorFormatter := helper.ErrorValidationFormat(err)
		errorMessage := gin.H{"errors": errorFormatter}

		response := helper.APIResponse("Failed to create company", "error", http.StatusInternalServerError, errorMessage)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	responseFormatter := event.FormatCompany(newCompany)

	// return response to client
	response := helper.APIResponse("Successfully to create company", "success", http.StatusOK, responseFormatter)
	c.JSON(http.StatusOK, response)
}
