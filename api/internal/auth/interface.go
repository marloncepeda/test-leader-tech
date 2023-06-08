package auth

type AuthPortRepositories interface {
	GetLogin(auth AuthModel) (AuthModelOut, error)
}
