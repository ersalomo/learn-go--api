package handler

import (
	"github.com/gin-gonic/gin"
	"go-api/db/service"
	"net/http"
)

type bookHandler struct {
	service service.BookService
}

func BookHandler(service service.BookService) *bookHandler {
	return &bookHandler{service}
}

func (h *bookHandler) GetBooks(c *gin.Context) {
	books, _ := h.service.FindAll()
	lenBooks := len(books)
	if lenBooks < 1 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "fail",
			"message": "No Data Found!",
			"data":    books,
		})
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"data":   books,
	})
}
