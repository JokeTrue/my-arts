package mysql

import (
	"database/sql"
	"fmt"
	"github.com/JokeTrue/my-arts/internal/models"
	"github.com/JokeTrue/my-arts/internal/products"
	"github.com/jmoiron/sqlx"
	"strconv"
	"strings"
)

type ProductsRepository struct {
	db *sqlx.DB
}

func NewProductsRepository(db *sqlx.DB) *ProductsRepository {
	return &ProductsRepository{db: db}
}

func (r *ProductsRepository) Create(product models.Product) (int, error) {
	panic("implement me")
}

func (r *ProductsRepository) GetProduct(id int) (*models.Product, error) {
	var product models.Product
	if err := r.db.Get(&product, QueryGetProductByID, id); err != nil {
		if err == sql.ErrNoRows {
			return nil, products.ErrProductQuery
		}
		return nil, products.ErrProductQuery
	}

	if err := r.setProductsPhotos(&product); err != nil {
		return nil, err
	}

	if err := r.setProductsTags(&product); err != nil {
		return nil, err
	}

	return &product, nil
}

func (r *ProductsRepository) Update(product models.Product) (int, error) {
	panic("implement me")
}

func (r *ProductsRepository) Delete(id int) error {
	queries := []string{
		QueryDeleteProductPhotos,
		QueryDeleteProductTags,
		QueryDeleteProduct,
	}

	tx, err := r.db.Beginx()
	if err != nil {
		return err
	}

	for _, query := range queries {
		if _, err := tx.Exec(query, id); err != nil {
			_ = tx.Rollback()
			return err
		}
	}

	return tx.Commit()
}

func (r *ProductsRepository) GetProducts(states []string) ([]*models.Product, error) {
	if states == nil {
		states = models.AllProductStates
	}

	list := []*models.Product{}

	query := fmt.Sprintf(QueryGetProducts, strings.Join(states, "','"))
	if err := r.db.Select(&list, query); err != nil {
		return nil, products.ErrProductQuery
	}

	if err := r.setProductsPhotos(list...); err != nil {
		return nil, err
	}

	if err := r.setProductsTags(list...); err != nil {
		return nil, err
	}

	return list, nil
}

func (r *ProductsRepository) GetUserProducts(userId int, states []string) ([]*models.Product, error) {
	if states == nil {
		states = models.AllProductStates
	}

	list := []*models.Product{}
	query := fmt.Sprintf(QueryGetUserProducts, strings.Join(states, "','"))
	if err := r.db.Select(&list, query, userId); err != nil {
		return nil, products.ErrProductQuery
	}

	if err := r.setProductsPhotos(list...); err != nil {
		return nil, err
	}

	if err := r.setProductsTags(list...); err != nil {
		return nil, err
	}

	return list, nil
}

func (r *ProductsRepository) setProductsPhotos(productsList ...*models.Product) error {
	if len(productsList) == 0 {
		return nil
	}

	productsMap := make(map[int]*models.Product)
	for _, product := range productsList {
		productsMap[product.ID] = product
	}

	var productIDs []string
	for productId := range productsMap {
		productIDs = append(productIDs, strconv.Itoa(productId))
	}

	var photos []*models.ProductPhoto
	query := fmt.Sprintf(QueryGetProductPhotos, strings.Join(productIDs, ","))
	if err := r.db.Select(&photos, query); err != nil && err != sql.ErrNoRows {
		return products.ErrProductQuery
	}

	for _, photo := range photos {
		if product, ok := productsMap[photo.ProductID]; ok {
			product.Photos = append(product.Photos, photo)
		}
	}

	return nil
}

func (r *ProductsRepository) setProductsTags(productsList ...*models.Product) error {
	if len(productsList) == 0 {
		return nil
	}

	productsMap := make(map[int]*models.Product)
	for _, product := range productsList {
		productsMap[product.ID] = product
	}

	var productIDs []string
	for productId := range productsMap {
		productIDs = append(productIDs, strconv.Itoa(productId))
	}

	var tags []*models.ProductTag
	query := fmt.Sprintf(QueryGetProductTags, strings.Join(productIDs, ","))
	if err := r.db.Select(&tags, query); err != nil && err != sql.ErrNoRows {
		return products.ErrProductQuery
	}

	for _, tag := range tags {
		if product, ok := productsMap[tag.ProductID]; ok {
			product.Tags = append(product.Tags, tag)
		}
	}

	return nil
}
