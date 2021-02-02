package mysql

const (
	QueryGetReviewByID = `
	SELECT * FROM friendship_requests WHERE id = ?
`

	QueryGetUserFriendshipRequests = `
	SELECT 
		u.first_name as 'user.first_name', 
		u.last_name as 'user.last_name', 
		u.age as 'user.age', 
		u.location as 'user.location', 
		fr.id,
		fr.created_at 
	FROM friendship_requests fr
	LEFT JOIN users AS u ON fr.actor_id = u.id 
	WHERE fr.friend_id = ?`

	QueryCreateFriendshipRequest = "INSERT INTO friendship_requests (actor_id, friend_id) VALUES (?, ?)"

	QueryDeleteFriendshipRequest = `DELETE FROM friendship_requests WHERE id = ?`

	QueryAcceptFriendshipRequest = `INSERT INTO friendships (user_1, user_2) VALUES (?, ?)`
)
