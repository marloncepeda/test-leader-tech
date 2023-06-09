package querys

const (
	GetUsers    = "SELECT id,name,lastname,username,email FROM users WHERE id = ? limit 1;"
	PostUsers   = "INSERT INTO users(name,lastname,username,password,type,email,salt) VALUES ($1, $2, $3, $4, $5, $6, $7);"
	UpdateUsers = "UPDATE users SET name = ?,lastname = ?, password = ?, username = ?, email = ?  WHERE id = ?;"
	DeleteUsers = "DELETE FROM users WHERE id = ?;"
)
