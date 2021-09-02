package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"startup/helper"
	"startup/user"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	// tangkap input dari user -> map input dari user ke struct RegisterUserInput ->
	// struct di atas kita parsing sebagai parameter service

	//objek input akan di mapping
	var input user.RegisterUserInput
	//mapping
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("register Account Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	//masukan ke db
	newUser, err := h.userService.RegisterUser(input)

	if err != nil {
		response := helper.APIResponse("register Account Failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := user.FormatUser(newUser, "tokenx")
	response := helper.APIResponse("Account has been resgistered", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}
