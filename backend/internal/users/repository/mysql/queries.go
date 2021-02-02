package mysql

const (
	QueryGetUsers = "SELECT * FROM users"

	QueryGetUserByID = "SELECT * FROM users WHERE id = ?"

	QueryGetUserByEmail = "SELECT * FROM users WHERE email = ?"

	QueryDeleteUser = "DELETE from users where id = ?"

	QueryCreateUser = "INSERT INTO users (email, password, first_name, last_name, age, gender, location, biography) VALUES (?, ?, ?, ?, ?, ?, ?, ?)"

	QueryUpdateUser = "UPDATE users SET first_name=?, last_name=?, age=?, gender=?, location=?, biography=? WHERE id=?"

	QuerySearchUsers = "SELECT * FROM users %s LIMIT 100"

	QueryGetUserFriends = `
	WITH friendship AS (
		SELECT user_1, user_2
		FROM friendships
		WHERE user_1 = ?
		UNION ALL
		SELECT user_2, user_1
		FROM friendships
		WHERE user_2 = ?
	)
	
	SELECT u.*
	FROM friendship f
	LEFT JOIN users AS U ON f.user_2 = u.id`
)
