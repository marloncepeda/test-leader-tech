package user

import (
	"api/internal/utils"
)

type UserService struct {
	repo UserPortRepositories
}

func NewUserService(repo UserPortRepositories) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetUser(id string) ([]UserModel, error) {

	db, err := s.repo.GetUser(id)
	if err != nil {
		return nil, err
	}
	return db, err
}

func (s *UserService) CreateUser(user UserModel) ([]UserModel, error) {
	pwhash, salt := utils.HashPassword(user.Password)

	user.hash = pwhash
	user.salt = salt

	db, err := s.repo.CreateUser(user)
	if err != nil {
		return nil, err
	}
	return db, err
}

func (s *UserService) UpdateUser(params UserModel) ([]UserModel, error) {
	db, err := s.repo.UpdateUser(params)
	if err != nil {
		return nil, err
	}
	return db, err
}

func (s *UserService) DeleteUser(id string) (string, error) {

	db, err := s.repo.DeleteUser(id)
	if err != nil {
		return "", err
	}
	return db, err
}

//this code can improve with a factory and call the database here and not in the http controller
