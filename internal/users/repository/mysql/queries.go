package mysql

const (
	QueryGetUsers       = "SELECT * FROM users WHERE id > 1"
	QueryGetUserByID    = "SELECT * FROM users WHERE id = ?"
	QueryGetUserByEmail = "SELECT * FROM users WHERE email = ?"
	QueryDeleteUser     = "DELETE from users where id = ?"
	QueryCreateUser     = "INSERT INTO users (email, password, first_name, last_name, age, gender, location, biography) VALUES (?, ?, ?, ?, ?, ?, ?, ?)"
	QueryUpdateUser     = "UPDATE users SET first_name=?, last_name=?, age=?, gender=?, location=?, biography=? WHERE id=?"
)
