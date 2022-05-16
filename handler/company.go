package handler

import (
	"cloviel-api/company"
	"cloviel-api/event"
	"cloviel-api/helper"
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

type companyHandler struct {
	companyService company.Service
}

func NewCompanyHandler(companyService company.Service) *companyHandler {
	return &companyHandler{companyService}
}

func (h *companyHandler) CreateNewCompany(c *gin.Context) {

	// mapping input JSON ke struct Companyinput
	var input company.CompanyInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errorFormatter := helper.ErrorValidationFormat(err)
		errorMessage := gin.H{"errors": errorFormatter}

		response := helper.APIResponse("Failed to create company", "error", http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// passing CompanyInput to service
	newCompany, err := h.companyService.CreateNewCompany(input)
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

func (h *companyHandler) UploadCompanyLogo(c *gin.Context) {

	// mapping input form
	var input company.CompanyLogoInput
	err := c.ShouldBind(&input)
	if err != nil {
		errorFormatter := helper.ErrorValidationFormat(err)

		errorMessage := gin.H{
			"is_uploaded": false,
			"errors":      errorFormatter,
		}

		response := helper.APIResponse("Failed to upload logo", "error", http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	file, err := c.FormFile("logo")
	if err != nil {
		data := gin.H{
			"is_uploaded": false,
			"errors":      err.Error(),
		}

		response := helper.APIResponse("Failed to upload logo", "error", http.StatusBadRequest, data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	fileExtSupport := []string{
		".png",
		".jpeg",
		".jpg",
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

	// make path location for save avatar
	path := fmt.Sprintf("assets/company-logo/%d-%s", input.CompanyID, file.Filename)

	// save file to server
	err = c.SaveUploadedFile(file, path)
	if err != nil {
		data := gin.H{
			"is_uploaded": false,
			"errors":      err.Error(),
		}

		response := helper.APIResponse("Failed to upload logo", "error", http.StatusBadRequest, data)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	// update file path to db
	_, err = h.companyService.SaveCompanyLogo(input.CompanyID, path)
	if err != nil {
		data := gin.H{
			"is_uploaded": false,
			"errors":      err.Error(),
		}

		response := helper.APIResponse("Failed to upload logo", "error", http.StatusBadGateway, data)

		c.JSON(http.StatusBadGateway, response)
		return
	}

	data := gin.H{"is_uploaded": true, "errors": nil}

	response := helper.APIResponse("Successfuly to upload company logo", "success", http.StatusOK, data)
	c.JSON(http.StatusOK, response)
}
