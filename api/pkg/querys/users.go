package querys

const (
	GetUsers    = "SELECT id,name,lastname,username,email FROM users WHERE id = $1 limit 1;"
	PostUsers   = "INSERT INTO users(name,lastname,username,password,type,email,salt) VALUES ($1, $2, $3, $4, $5, $6, $7);"
	UpdateUsers = "UPDATE users SET name = $1,lastname = $2, password = $3, username = $4, email = $5  WHERE id = $6;"
	DeleteUsers = "DELETE FROM users WHERE id = $1;"
)
