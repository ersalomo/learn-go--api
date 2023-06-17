package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go-api/db/model"
	"go-api/db/service"
	"net/http"
	"strconv"
	"strings"
)

type userHandler struct {
	userService service.Service
}

func UserHandler(userService service.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) GetUser(c *gin.Context) {
	id := c.Param("id")
	idx, _ := strconv.Atoi(id)
	user, err := h.userService.FindById(idx)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status": "fail",
			"msg":    err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   user,
	})
}
func (h *userHandler) GetUsers(c *gin.Context) {
	//email := c.Query("email")
	users, _ := h.userService.FindAll()
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   users,
	})
}
func (h *userHandler) StoreUser(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		var errors []map[string]any
		for _, errorField := range err.(validator.ValidationErrors) {
			key := strings.ToLower(errorField.Field())
			value := strings.ToLower(errorField.ActualTag())
			msg := fmt.Sprintf("Field %v is %v", key, value)
			errors = append(errors, map[string]any{
				key: msg,
			})
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "fail",
			"errors": errors,
		})
		return
	}

	user, err := h.userService.Create(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "fail",
			"errors": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"user":   user,
	})
}
