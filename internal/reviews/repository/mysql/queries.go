package mysql

const (
	QueryGetReviewByID = `
	SELECT reviews.*,
		   reviewer.id         as 'reviewer.id',
		   reviewer.first_name as 'reviewer.first_name',
		   reviewer.last_name  as 'reviewer.last_name'
	FROM reviews
	LEFT JOIN users reviewer on reviewer.id = reviews.reviewer_id
	WHERE reviews.id = ?`

	QueryDeleteReview   = `DELETE FROM reviews WHERE id = ?`

	QueryGetUserReviews = `
	SELECT reviews.*,
		   reviewer.id         as 'reviewer.id',
		   reviewer.first_name as 'reviewer.first_name',
		   reviewer.last_name  as 'reviewer.last_name'
	FROM reviews
	LEFT JOIN users reviewer on reviewer.id = reviews.reviewer_id
	WHERE user_id = ?`

	QueryCreateReview = `
	INSERT INTO reviews (user_id, reviewer_id, comment, delivery_rating, communication_rating, accuracy_rating)
	VALUES (?, ?, ?, ?, ?, ?)`

	QueryUpdateReview = `
	UPDATE reviews 
	SET comment=?, delivery_rating=?, communication_rating=?, accuracy_rating=?, edited_at=NOW() WHERE id = ?`
)
