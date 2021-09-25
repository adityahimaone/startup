package campaign

type Service interface {
	GetCampaigns(useerID int) ([]Campaign, error)
	GetCampaignByID(req RequestCampaignDetail) (Campaign, error)
	CreateCampaign(req RequestCreateCampaign) (Campaign, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetCampaigns(userID int) ([]Campaign, error) {
	if userID != 0 {
		campaigns, err := s.repository.FindByUserID(userID)
		if err != nil {
			return campaigns, err
		}
		return campaigns, nil
	}
	campaigns, err := s.repository.FindAll()
	if err != nil {
		return campaigns, err
	}
	return campaigns, nil
}

func (s *service) GetCampaignByID(req RequestCampaignDetail) (Campaign, error) {
	campaign, err := s.repository.FindByID(req.ID)
	if err != nil {
		return campaign, err
	}
	return campaign, nil
}

func (s *service) CreateCampaign(req RequestCreateCampaign) (Campaign, error) {
	newCampaign, err := s.repository.Save(req.toModel())
	if err != nil {
		return newCampaign, err
	}
	return newCampaign, nil
}
