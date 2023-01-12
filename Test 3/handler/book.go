package handler

import (
	"gits/book"
	"gits/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

type bookHandler struct {
	service book.Service
}

func NewBookHandler(bookService book.Service) *bookHandler {
	return &bookHandler{bookService}
}

func (h *bookHandler) RegisterBook(c *gin.Context) {
	var input book.BookDataInput

	err := c.ShouldBindJSON(&input)
	if err != nil {

		errors := helper.FormatValidationError(err)

		errorMessage := gin.H{"errors": errors}

		response := helper.ApiResponse("Failed to create book", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return

	}

	newBook, err := h.service.RegisterBook(input)

	if err != nil {
		response := helper.ApiResponse("Failed to create book", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.ApiResponse("Succes  create  book", http.StatusCreated, "success", book.FormatBook(newBook))
	c.JSON(http.StatusCreated, response)

}

func (h *bookHandler) GetBookById(c *gin.Context) {

	var inputID book.BookParamInput
	err := c.ShouldBindUri(&inputID)

	if err != nil {
		response := helper.ApiResponse("Failed to update book", http.StatusBadRequest, "error", nil)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	bookData, err := h.service.GetBookByID(inputID)

	if bookData.ID == 0 {
		response := helper.ApiResponse("book data not found", http.StatusBadRequest, "error", nil)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.ApiResponse("Succes  get book data", http.StatusOK, "success", book.FormatBookDetail(bookData))
	c.JSON(http.StatusOK, response)

}

func (h *bookHandler) UpdateBookById(c *gin.Context) {

	var inputID book.BookParamInput
	err := c.ShouldBindUri(&inputID)

	if err != nil {
		response := helper.ApiResponse("Failed to update book", http.StatusBadRequest, "error", nil)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	bookData, err := h.service.GetBookByID(inputID)

	if bookData.ID == 0 {
		response := helper.ApiResponse("book data not found", http.StatusBadRequest, "error", nil)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	var input book.BookDataInput

	err = c.ShouldBindJSON(&input)

	if err != nil {

		errors := helper.FormatValidationError(err)

		errorMessage := gin.H{"errors": errors}

		response := helper.ApiResponse("Failed to update book", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return

	}

	newBook, err := h.service.UpdateBookData(inputID, input)

	if err != nil {
		response := helper.ApiResponse("Failed to update book", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.ApiResponse("Succes  update book data", http.StatusOK, "success", book.FormatBook(newBook))
	c.JSON(http.StatusOK, response)

}

func (h *bookHandler) DestroyBookById(c *gin.Context) {

	var inputID book.BookParamInput
	err := c.ShouldBindUri(&inputID)

	if err != nil {
		response := helper.ApiResponse("Failed to destroy book", http.StatusBadRequest, "error", nil)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	bookData, err := h.service.GetBookByID(inputID)

	if bookData.ID == 0 {
		response := helper.ApiResponse("book data not found", http.StatusBadRequest, "error", nil)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	err = h.service.DestroyBook(inputID)

	if err != nil {
		response := helper.ApiResponse("Failed to delete book", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.ApiResponse("Succes  delete  book", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)
}
