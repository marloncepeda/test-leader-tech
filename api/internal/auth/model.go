package auth

type AuthModel struct {
	Id       int
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthModelOut struct {
	Id       int
	Username string `json:"username"`
	Password []byte `json:"password"`
	Salt     []byte
}
