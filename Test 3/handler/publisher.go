package handler

import (
	"gits/helper"
	"gits/publisher"
	"net/http"

	"github.com/gin-gonic/gin"
)

type publisherHandler struct {
	service publisher.Service
}

func NewPublisherHandler(publisherService publisher.Service) *publisherHandler {
	return &publisherHandler{publisherService}
}

func (h *publisherHandler) RegisterPublisher(c *gin.Context) {
	var input publisher.PublisherDataInput

	err := c.ShouldBindJSON(&input)

	if err != nil {

		errors := helper.FormatValidationError(err)

		errorMessage := gin.H{"errors": errors}

		response := helper.ApiResponse("Failed to create publisher", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return

	}

	newPublisher, err := h.service.RegisterPublisher(input)

	if err != nil {
		response := helper.ApiResponse("Failed to create publisher", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.ApiResponse("Succes  create  publisher", http.StatusCreated, "success", publisher.FormatPublisher(newPublisher))
	c.JSON(http.StatusCreated, response)

}

func (h *publisherHandler) GetPublisherById(c *gin.Context) {

	var inputID publisher.PublisherParamInput
	err := c.ShouldBindUri(&inputID)

	if err != nil {
		response := helper.ApiResponse("Failed to update publisher", http.StatusBadRequest, "error", nil)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	publisherData, err := h.service.GetPublisherByID(inputID)

	if publisherData.ID == 0 {
		response := helper.ApiResponse("publisher data not found", http.StatusBadRequest, "error", nil)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.ApiResponse("Succes  get publisher data", http.StatusOK, "success", publisher.FormatPublisher(publisherData))
	c.JSON(http.StatusOK, response)

}

func (h *publisherHandler) UpdatePublisherById(c *gin.Context) {

	var inputID publisher.PublisherParamInput
	err := c.ShouldBindUri(&inputID)

	if err != nil {
		response := helper.ApiResponse("Failed to update publisher", http.StatusBadRequest, "error", nil)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	publisherData, err := h.service.GetPublisherByID(inputID)

	if publisherData.ID == 0 {
		response := helper.ApiResponse("publisher data not found", http.StatusBadRequest, "error", nil)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	var input publisher.PublisherDataInput

	err = c.ShouldBindJSON(&input)

	if err != nil {

		errors := helper.FormatValidationError(err)

		errorMessage := gin.H{"errors": errors}

		response := helper.ApiResponse("Failed to update publisher", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return

	}

	newPublisher, err := h.service.UpdatePublisherData(inputID, input)

	if err != nil {
		response := helper.ApiResponse("Failed to update publisher", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.ApiResponse("Succes  update publisher data", http.StatusOK, "success", publisher.FormatPublisher(newPublisher))
	c.JSON(http.StatusOK, response)

}
func (h *publisherHandler) DestroyPublisherById(c *gin.Context) {

	var inputID publisher.PublisherParamInput
	err := c.ShouldBindUri(&inputID)

	if err != nil {
		response := helper.ApiResponse("Failed to destroy publisher", http.StatusBadRequest, "error", nil)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	publisherData, err := h.service.GetPublisherByID(inputID)

	if publisherData.ID == 0 {
		response := helper.ApiResponse("publisher data not found", http.StatusBadRequest, "error", nil)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	err = h.service.DestroyPublisher(inputID)

	if err != nil {
		response := helper.ApiResponse("Failed to delete publisher", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.ApiResponse("Succes  delete  publisher", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)
}
