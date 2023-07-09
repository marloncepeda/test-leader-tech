package shipment

import "api/internal/utils"

type ShipmentService struct {
	repo ShipmentPortRepositories
}

func NewShipmentService(repo ShipmentPortRepositories) *ShipmentService {
	return &ShipmentService{repo: repo}
}

func (s *ShipmentService) GetUser(id string) ([]ShipmentModel, error) {

	db, err := s.repo.GetShipment(id)
	if err != nil {
		return nil, err
	}
	return db, err
}

func (s *ShipmentService) CreateUser(user ShipmentModel) ([]ShipmentModel, error) {
	pwhash, salt := utils.HashPassword(user.Password)

	user.hash = pwhash
	user.salt = salt

	db, err := s.repo.CreateShipment(user)
	if err != nil {
		return nil, err
	}
	return db, err
}

func (s *ShipmentService) UpdateUser(params ShipmentModel) ([]ShipmentModel, error) {
	db, err := s.repo.UpdateShipment(params)
	if err != nil {
		return nil, err
	}
	return db, err
}
