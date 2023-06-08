package querys

const (
	GetUsersAuth = "SELECT id,username,password,salt FROM users WHERE username = ? limit 1;"
)
