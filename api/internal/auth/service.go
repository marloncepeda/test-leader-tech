package auth

import (
	"api/internal/utils"
	"errors"
	"log"
)

type AuthService struct {
	repo AuthPortRepositories
}

func NewAuthService(repo AuthPortRepositories) *AuthService {
	return &AuthService{repo: repo}
}

func (as *AuthService) Login(auth AuthModel) (string, error) {

	db, err := as.repo.GetLogin(auth)
	if err != nil {
		return "", err
	}

	c := utils.CheckPassword(auth.Password, db.Password, db.Salt)
	if !c {
		return "", errors.New("invalid password")
	}

	jwt, err := utils.GenerateJWT(db.Id, db.Username)
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	return jwt, nil
}

//func (as *AuthService) Logout()

//func (as *AuthService) Validate()
