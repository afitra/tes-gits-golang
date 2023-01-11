package handler

import (
	"gits/author"
	"gits/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

type authorHandler struct {
	service author.Service
}

func NewAuthorHandler(authorService author.Service) *authorHandler {
	return &authorHandler{authorService}
}

func (h *authorHandler) RegisterAuthor(c *gin.Context) {
	var input author.AuthorDataInput

	err := c.ShouldBindJSON(&input)

	if err != nil {

		errors := helper.FormatValidationError(err)

		errorMessage := gin.H{"errors": errors}

		response := helper.ApiResponse("Failed to create author", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return

	}

	newAuthor, err := h.service.RegisterAuthor(input)

	if err != nil {
		response := helper.ApiResponse("Failed to create author", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.ApiResponse("Succes  create  author", http.StatusCreated, "success", author.FormatAuthor(newAuthor))
	c.JSON(http.StatusCreated, response)

}

func (h *authorHandler) GetAuthorById(c *gin.Context) {

	var inputID author.AuthorParamInput
	err := c.ShouldBindUri(&inputID)

	if err != nil {
		response := helper.ApiResponse("Failed to update author", http.StatusBadRequest, "error", nil)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	authorData, err := h.service.GetAuthorByID(inputID)

	if authorData.ID == 0 {
		response := helper.ApiResponse("author data not found", http.StatusBadRequest, "error", nil)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.ApiResponse("Succes  get author data", http.StatusOK, "success", author.FormatAuthor(authorData))
	c.JSON(http.StatusOK, response)

}
func (h *authorHandler) UpdateAuthorById(c *gin.Context) {

	var inputID author.AuthorParamInput
	err := c.ShouldBindUri(&inputID)

	if err != nil {
		response := helper.ApiResponse("Failed to update author", http.StatusBadRequest, "error", nil)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	authorData, err := h.service.GetAuthorByID(inputID)

	if authorData.ID == 0 {
		response := helper.ApiResponse("author data not found", http.StatusBadRequest, "error", nil)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	var input author.AuthorDataInput

	err = c.ShouldBindJSON(&input)

	if err != nil {

		errors := helper.FormatValidationError(err)

		errorMessage := gin.H{"errors": errors}

		response := helper.ApiResponse("Failed to update author", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return

	}

	newAuthor, err := h.service.UpdateAuthorData(inputID, input)

	if err != nil {
		response := helper.ApiResponse("Failed to update author", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.ApiResponse("Succes  update author data", http.StatusOK, "success", author.FormatAuthor(newAuthor))
	c.JSON(http.StatusOK, response)

}
func (h *authorHandler) DestroyAuthorById(c *gin.Context) {

	var inputID author.AuthorParamInput
	err := c.ShouldBindUri(&inputID)

	if err != nil {
		response := helper.ApiResponse("Failed to destroy author", http.StatusBadRequest, "error", nil)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	authorData, err := h.service.GetAuthorByID(inputID)

	if authorData.ID == 0 {
		response := helper.ApiResponse("author data not found", http.StatusBadRequest, "error", nil)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	err = h.service.DestroyAuthor(inputID)

	if err != nil {
		response := helper.ApiResponse("Failed to delete author", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.ApiResponse("Succes  delete  author", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)
}
