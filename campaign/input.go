package campaign

import (
	"fmt"
	"github.com/gosimple/slug"
	"startup/user"
)

type RequestCampaignDetail struct {
	ID int `uri:"id" binding:"required"`
}

type RequestCreateCampaign struct {
	Name             string `json:"name"`
	ShortDescription string `json:"short_description"`
	Description      string `json:"description"`
	GoalAmount       int    `json:"goal_amount"`
	Perks            string `json:"perks"`
	User             user.User
}

//mapping req create campaign to model Campaign
func (req *RequestCreateCampaign) toModel() Campaign {
	//proses generate slug from third party
	combineSlug := fmt.Sprintf("%s %s", req.Name, req.User.ID)
	slugGenerate := slug.Make(combineSlug)
	return Campaign{
		Name:             req.Name,
		ShortDescription: req.ShortDescription,
		Description:      req.Description,
		Perks:            req.Perks,
		GoalAmount:       req.GoalAmount,
		UserId:           req.User.ID,
		Slug:             slugGenerate,
	}
}
