package mysql

const (
	QueryDeleteCategory = `DELETE FROM categories WHERE id = ?`

	QueryGetCategories = `SELECT * FROM categories ORDER BY title`

	QueryGetCategoryByID = `SELECT * FROM categories WHERE id = ?`

	QueryCreateCategory = `INSERT INTO categories (title) VALUES (?)`

	QueryUpdateCategory = `UPDATE categories SET title = ? WHERE id = ?`
)
