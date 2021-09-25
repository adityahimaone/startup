package campaign

import (
	"log"
	"strings"
)

type ResponseCampaign struct {
	ID               int    `json:"id"`
	UserID           int    `json:"user_id"`
	Name             string `json:"name"`
	ShortDescription string `json:"short_description"`
	ImageURL         string `json:"image_url"`
	GoalAmount       int    `json:"goal_amount"`
	CurrentAmount    int    `json:"current_amount"`
	Slug             string `json:"slug"`
}

func NewResponseCampaign(entityCampaign Campaign) ResponseCampaign {
	var imageURL string
	if len(entityCampaign.CampaignImages) > 0 {
		imageURL = entityCampaign.CampaignImages[0].FileName
	}
	return ResponseCampaign{
		ID:               entityCampaign.ID,
		UserID:           entityCampaign.UserId,
		Name:             entityCampaign.Name,
		ShortDescription: entityCampaign.ShortDescription,
		GoalAmount:       entityCampaign.GoalAmount,
		CurrentAmount:    entityCampaign.CurrentAmount,
		Slug:             entityCampaign.Slug,
		ImageURL:         imageURL,
	}
}

func NewResponseCampaignArray(entityCampaign []Campaign) []ResponseCampaign {
	result := []ResponseCampaign{}

	for _, v := range entityCampaign {
		result = append(result, NewResponseCampaign(v))
	}
	return result
}

type ResponseCampaignDetail struct {
	ID               int                            `json:"id"`
	Name             string                         `json:"name"`
	ShortDescription string                         `json:"short_description"`
	Description      string                         `json:"description"`
	ImageURL         string                         `json:"image_url"`
	GoalAmount       int                            `json:"goal_amount"`
	CurrentAmount    int                            `json:"current_amount"`
	UserID           int                            `json:"user_id"`
	Slug             string                         `json:"slug"`
	Perks            []string                       `json:"perks"`
	User             ResponseCampaignUserDetail     `json:"user"`
	Images           []ResponseCampaignImagesDetail `json:"images"`
}

type ResponseCampaignUserDetail struct {
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}

type ResponseCampaignImagesDetail struct {
	ImageURL  string `json:"image_url"`
	IsPrimary bool   `json:"is_primary"`
}

func NewResponseCampaignDetail(entityCampaign Campaign) ResponseCampaignDetail {
	// get image url
	var imageURL string
	if len(entityCampaign.CampaignImages) > 0 {
		imageURL = entityCampaign.CampaignImages[0].FileName
	}
	//split perks into slice perks
	var perks []string
	for _, perk := range strings.Split(entityCampaign.Perks, ",") {
		perks = append(perks, strings.TrimSpace(perk))
	}
	//nested obj User to Response Detail
	user := entityCampaign.User
	campaignUserDetail := ResponseCampaignUserDetail{
		Name:     user.Name,
		ImageURL: user.AvatarFileName,
	}
	//nested obj Images to Response Detail
	images := entityCampaign.CampaignImages
	campaignImagesDetail := []ResponseCampaignImagesDetail{}
	for _, image := range images {
		convPrimary := false
		log.Println(image.isPrimary)
		log.Println(image.FileName)
		if image.isPrimary == 1 {
			convPrimary = true
		}
		imagesDetail := ResponseCampaignImagesDetail{
			ImageURL:  image.FileName,
			IsPrimary: convPrimary,
		}
		campaignImagesDetail = append(campaignImagesDetail, imagesDetail)
	}

	return ResponseCampaignDetail{
		ID:               entityCampaign.ID,
		Name:             entityCampaign.Name,
		ShortDescription: entityCampaign.ShortDescription,
		Description:      entityCampaign.Description,
		GoalAmount:       entityCampaign.GoalAmount,
		CurrentAmount:    entityCampaign.CurrentAmount,
		UserID:           entityCampaign.UserId,
		Slug:             entityCampaign.Slug,
		ImageURL:         imageURL,
		Perks:            perks,
		User:             campaignUserDetail,
		Images:           campaignImagesDetail,
	}
}
