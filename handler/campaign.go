package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"startup/campaign"
	"startup/helper"
	"startup/user"
	"strconv"
)

// tangkap parameter di handler
// handler -> service (service menentukan apakah repository mana yg di panggil /method mana yg di call)
// repository (Get All, Get by User ID) -> manggi DB

// untuk mewakili handler -> apa yg dibutuhkan
type campaignHandler struct {
	service campaign.Service
}

//membuat objek/struct dr campaignHandler yg akan dipanggil di main.go
func NewCampaignHandler(service campaign.Service) *campaignHandler {
	//return objek baru dari campaign hadnler dan parsing service
	return &campaignHandler{service: service}
}

//function untuk map ke api
func (handler *campaignHandler) GetCampaigns(c *gin.Context) {
	//tangkap request
	userID, _ := strconv.Atoi(c.Query("user_id"))

	campaigns, err := handler.service.GetCampaigns(userID)
	if err != nil {
		errorMessage := gin.H{"errors": "Error 404"}
		response := helper.APIResponse("Error get campaigns", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := campaign.NewResponseCampaignArray(campaigns)
	response := helper.APIResponse("List of campaigns", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

// api/v1/campaigns/{id}
//handler : mapping id yg di url ke struct input => service, call formatter
// service : struct input ubtuk menangkap id di url -> manggil repo
// repository : get campaign by id
func (handler *campaignHandler) GetCampaign(c *gin.Context) {
	var req campaign.RequestCampaignDetail
	err := c.ShouldBindUri(&req)
	if err != nil {
		response := helper.APIResponse("Failed to get detail campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	campaignDetail, err := handler.service.GetCampaignByID(req)
	if err != nil {
		response := helper.APIResponse("Failed to get detail campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := campaign.NewResponseCampaignDetail(campaignDetail)
	response := helper.APIResponse("Campaign Detail", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
	return
}

/*
tangkap parameter dari user ke input(dto) struct
ambil current user dari jwt/hanler
panggil service, parameter input struct yg sudah di mapping dan slug
panggil repository untuk simpan data dari service
*/

func (h *campaignHandler) CreateCampaign(c *gin.Context) {
	var req campaign.RequestCreateCampaign
	err := c.ShouldBindJSON(&req)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Create Campaign Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	//get user from jwt
	currentUser := c.MustGet("currentUser").(user.User)
	req.User = currentUser

	newCampaign, err := h.service.CreateCampaign(req)
	if err != nil {
		response := helper.APIResponse("Create Campaign Failed", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	response := helper.APIResponse("Success Create Campaign", http.StatusOK, "succes", campaign.NewResponseCampaign(newCampaign))
	c.JSON(http.StatusOK, response)
	return
}

/*
- user masukan input
- handler
- mapping dari input ke input struct(dto) ada (2)
- parameter dari user dan di uri -> parsing ke service
- service
- repository update data campaign
*/

func (h *campaignHandler) UpdateCampaign(c *gin.Context) {
	var reqID campaign.RequestCampaignDetail
	err := c.ShouldBindUri(&reqID)
	if err != nil {
		response := helper.APIResponse("Failed to update campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var reqBody campaign.RequestCreateCampaign
	err = c.ShouldBindJSON(&reqBody)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Update Campaign Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	updatedCampaign, err := h.service.UpdateCampaign(reqID, reqBody)
	if err != nil {
		response := helper.APIResponse("Failed to update campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("Success Update campaign", http.StatusOK, "Success", updatedCampaign)
	c.JSON(http.StatusOK, response)
	return
}
