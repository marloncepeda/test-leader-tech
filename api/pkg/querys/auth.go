package querys

const (
	GetUsersAuth = "SELECT id,username,password,salt FROM users WHERE username = $1 limit 1;"
)
