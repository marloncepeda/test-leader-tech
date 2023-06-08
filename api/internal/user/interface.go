package user

type UserPortRepositories interface {
	GetUser(id string) ([]UserModel, error)
	CreateUser(user UserModel) ([]UserModel, error)
	UpdateUser(params UserModel) ([]UserModel, error)
	DeleteUser(id string) (string, error)
}
