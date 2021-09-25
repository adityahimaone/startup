package campaign

import (
	"fmt"
	"github.com/gosimple/slug"
	"startup/user"
	"strconv"
)

type RequestCampaignDetail struct {
	ID int `uri:"id" binding:"required"`
}

type RequestCreateCampaign struct {
	Name             string `json:"name" binding:"required"`
	ShortDescription string `json:"short_description" binding:"required"`
	Description      string `json:"description" binding:"required"`
	GoalAmount       int    `json:"goal_amount" binding:"required"`
	Perks            string `json:"perks" binding:"required"`
	User             user.User
}

//mapping req create campaign to model Campaign
func (req *RequestCreateCampaign) toModel() Campaign {
	//proses generate slug from third party
	combineSlug := fmt.Sprintf("%s %s", req.Name, strconv.Itoa(req.User.ID))
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
