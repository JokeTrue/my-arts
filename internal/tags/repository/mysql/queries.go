package mysql

const (
	QueryDeleteTag = `DELETE FROM product_tags WHERE id = ?`

	QueryGetTags = `SELECT * FROM product_tags ORDER BY title`

	QueryGetProductTags = `SELECT * FROM product_tags WHERE product_id = ? ORDER BY title`

	QueryGetTagByID = `SELECT * FROM product_tags WHERE id = ?`

	QueryCreateTag = `INSERT INTO product_tags (product_id, title) VALUES (?, ?)`

	QueryUpdateTag = `UPDATE product_tags SET title = ? WHERE id = ?`
)

