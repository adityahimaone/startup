package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"startup/campaign"
	"startup/helper"
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
