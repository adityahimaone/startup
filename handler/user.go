package handler

import (
	"fmt"
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

//user memasukan input
//input ditangkap handler
//mapping dari input user ke input struct
//input struct parsing ke service
//service mencari dg bantuan repository user dengan email x
//jika ketemu mencocokkan password

func (h *userHandler) Login(c *gin.Context) {
	var input user.LoginInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Login Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	loggedinUser, err := h.userService.Login(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Login Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	formatter := user.FormatUser(loggedinUser, "tokenx")
	response := helper.APIResponse("Login Success", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

// ada input email dari user
// input email di mapping ke struct input -> handler
// struct input di parsing di service
// service akan memanngil reposirtory apakah email sudah ada atau belum
// repository akan mengquery ke DB

func (h *userHandler) CheckEmailAvailable(c *gin.Context) {
	var input user.CheckEmailInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Checking Email Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	isEmailAvailable, err := h.userService.IsEmailAvailable(input)
	if err != nil {
		errorMessage := gin.H{"errors": "Server Error"}
		response := helper.APIResponse("Checking Email Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	data := gin.H{
		"is_available": isEmailAvailable,
	}

	var metaMessage string

	if isEmailAvailable {
		metaMessage = "Email is Available"
	} else {
		metaMessage = "Email has been registered"
	}

	response := helper.APIResponse(metaMessage, http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)
}

// tangkap input dari user Form Body
// simpan gambarnya di folder "images//"
// service -> panggil repository -> user yg mengakses
// JWT (sementara hardcode)
// repository -> mengambil data user yg ID nya tertentu
// Repo update Data User simapn lokasi file
func (h *userHandler) UploudAvatar(c *gin.Context) {
	//field avatar
	file, err := c.FormFile("avatar")
	if err != nil {
		data := gin.H{
			"is_uplouded": false,
		}
		response := helper.APIResponse("Failed to uploud avatar image", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	userID := 4
	//path := "images/avatar/" + file.Filename
	path := fmt.Sprintf("images/avatar/%d-%s", userID, file.Filename)
	err = c.SaveUploadedFile(file, path)
	if err != nil {
		data := gin.H{
			"is_uplouded": false,
		}
		response := helper.APIResponse("Failed to uploud avatar image", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	//panggil service
	//hardcode dulu
	_, err = h.userService.SaveAvatar(userID, path)
	if err != nil {
		data := gin.H{
			"is_uplouded": false,
		}
		response := helper.APIResponse("Failed to uploud avatar image", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	data := gin.H{
		"is_uplouded": true,
	}
	response := helper.APIResponse("Avatar Successfully Update", http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)
}
