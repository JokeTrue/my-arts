package mysql

import (
	"database/sql"
	"github.com/JokeTrue/my-arts/internal/categories"
	"github.com/JokeTrue/my-arts/internal/models"
	"github.com/jmoiron/sqlx"
)

type CategoriesRepository struct {
	writeDB *sqlx.DB
	readDB  *sqlx.DB
}

func NewCategoriesRepository(writeDB, readDB *sqlx.DB) *CategoriesRepository {
	return &CategoriesRepository{writeDB: writeDB, readDB: readDB}
}

func (r *CategoriesRepository) Delete(id int) error {
	res, err := r.writeDB.Exec(QueryDeleteCategory, id)
	if err != nil {
		return categories.ErrCategoryQuery
	}

	affect, err := res.RowsAffected()
	if err != nil {
		return categories.ErrCategoryQuery
	}
	if affect == 0 {
		return categories.ErrCategoryNotFound
	}

	return nil
}

func (r *CategoriesRepository) GetCategories() ([]*models.Category, error) {
	list := []*models.Category{}
	if err := r.readDB.Select(&list, QueryGetCategories); err != nil {
		return nil, categories.ErrCategoryQuery
	}
	return list, nil
}

func (r *CategoriesRepository) GetCategory(id int) (*models.Category, error) {
	var category models.Category
	if err := r.readDB.Get(&category, QueryGetCategoryByID, id); err != nil {
		if err == sql.ErrNoRows {
			return nil, categories.ErrCategoryNotFound
		}
		return nil, categories.ErrCategoryQuery
	}

	return &category, nil
}

func (r *CategoriesRepository) Create(category models.Category) (int, error) {
	result, err := r.writeDB.Exec(QueryCreateCategory, category.Title)
	if err != nil {
		return 0, categories.ErrCategoryQuery
	}

	lastID, err := result.LastInsertId()
	if err != nil {
		return 0, categories.ErrCategoryQuery
	}

	return int(lastID), nil
}

func (r *CategoriesRepository) Update(category models.Category) (*models.Category, error) {
	res, err := r.writeDB.Exec(QueryUpdateCategory, category.Title, category.ID)
	if err != nil {
		return nil, categories.ErrCategoryQuery
	}

	affect, err := res.RowsAffected()
	if err != nil {
		return nil, categories.ErrCategoryQuery
	}
	if affect == 0 {
		return nil, categories.ErrCategoryNotFound
	}

	return &category, nil
}
