package querys

const (
	GetUsers    = "SELECT id,name,lastname,username,email FROM users WHERE id = ? limit 1;"
	PostUsers   = "INSERT INTO users(name,lastname,password,username,email,salt) VALUES (?,?,?,?,?,?);"
	UpdateUsers = "UPDATE users SET name = ?,lastname = ?, password = ?, username = ?, email = ?  WHERE id = ?;"
	DeleteUsers = "DELETE FROM users WHERE id = ?;"
)
