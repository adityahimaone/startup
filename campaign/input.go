package campaign

type RequestCampaignDetail struct {
	ID int `uri:"id" binding:"required"`
}
