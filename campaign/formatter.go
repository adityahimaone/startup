package campaign

type ResponseCampaign struct {
	ID               int    `json:"id"`
	UserID           int    `json:"user_id"`
	Name             string `json:"name"`
	ShortDescription string `json:"short_description"`
	ImageURL         string `json:"image_url"`
	GoalAmount       int    `json:"goal_amount"`
	CurrentAmount    int    `json:"current_amount"`
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
		ImageURL:         imageURL,
	}
}

func NewResponseCampaignArray(entityCampaign []Campaign) []ResponseCampaign {
	var result []ResponseCampaign
	for _, v := range entityCampaign {
		result = append(result, NewResponseCampaign(v))
	}
	return result
}
