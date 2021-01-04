package mysql

const (
	QueryGetProductByID = `
	SELECT
		   products.*,
		   c.id as 'category.id',
		   c.title as 'category.title',
		   c.created_at as 'category.created_at'
	FROM products
	LEFT JOIN categories c on c.id = products.category_id
	WHERE products.id = ?`

	QueryGetProductTags   = `SELECT * FROM product_tags WHERE product_id IN (%s)`
	QueryGetProductPhotos = `SELECT * FROM product_photos WHERE product_id IN (%s)`

	QueryDeleteProduct       = `DELETE FROM products WHERE id = ?`
	QueryDeleteProductPhotos = `DELETE FROM product_photos WHERE product_id = ?`
	QueryDeleteProductTags   = `DELETE FROM product_tags WHERE product_id = ?`

	QueryGetProducts = `
	SELECT
		   products.*,
		   c.id as 'category.id',
		   c.title as 'category.title',
		   c.created_at as 'category.created_at'
	FROM products
	LEFT JOIN categories c on c.id
	WHERE products.state IN ('%s')`

	QueryGetUserProducts = `
	SELECT
		   products.*,
		   c.id as 'category.id',
		   c.title as 'category.title',
		   c.created_at as 'category.created_at'
	FROM products
	LEFT JOIN categories c on c.id
	WHERE products.user_id = ? AND products.state IN ('%s')`
)
