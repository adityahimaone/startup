package campaign

import "gorm.io/gorm"

//API Contract
type Repository interface {
	FindAll() ([]Campaign, error) // slice -> membalikan banyak data
	FindByUserID(userID int) ([]Campaign, error)
}

//Private struct
type repository struct {
	db *gorm.DB
}

//membuat instance struct repository
func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Campaign, error) {
	var campaigns []Campaign
	err := r.db.Preload("CampaignImages", "campaign_images.is_primary = 1").Find(&campaigns).Error
	if err != nil {
		return campaigns, err
	}
	return campaigns, nil
}

func (r *repository) FindByUserID(userID int) ([]Campaign, error) {
	var campaigns []Campaign
	err := r.db.Where("user_id = ?", userID).Preload("CampaignImages", "campaign_images.is_primary = 1").Find(&campaigns).Error
	if err != nil {
		return campaigns, err
	}
	return campaigns, nil
}
