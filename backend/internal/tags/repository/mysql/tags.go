package mysql

import (
	"database/sql"
	"github.com/JokeTrue/my-arts/internal/models"
	"github.com/JokeTrue/my-arts/internal/tags"
	"github.com/jmoiron/sqlx"
)

type TagsRepository struct {
	writeDB *sqlx.DB
	readDB  *sqlx.DB
}

func NewTagsRepository(writeDB, readDB *sqlx.DB) *TagsRepository {
	return &TagsRepository{writeDB: writeDB, readDB: readDB}
}

func (r *TagsRepository) Delete(id int) error {
	res, err := r.writeDB.Exec(QueryDeleteTag, id)
	if err != nil {
		return tags.ErrTagQuery
	}

	affect, err := res.RowsAffected()
	if err != nil {
		return tags.ErrTagQuery
	}
	if affect == 0 {
		return tags.ErrTagNotFound
	}

	return nil
}

func (r *TagsRepository) GetTags() ([]*models.ProductTag, error) {
	list := []*models.ProductTag{}
	if err := r.readDB.Select(&list, QueryGetTags); err != nil {
		return nil, tags.ErrTagQuery
	}
	return list, nil
}

func (r *TagsRepository) Create(tag models.ProductTag) (int, error) {
	result, err := r.writeDB.Exec(QueryCreateTag, tag.ProductID, tag.Title)
	if err != nil {
		return 0, tags.ErrTagQuery
	}

	lastID, err := result.LastInsertId()
	if err != nil {
		return 0, tags.ErrTagQuery
	}

	return int(lastID), nil
}

func (r *TagsRepository) GetTag(id int) (*models.ProductTag, error) {
	var tag models.ProductTag
	if err := r.readDB.Get(&tag, QueryGetTagByID, id); err != nil {
		if err == sql.ErrNoRows {
			return nil, tags.ErrTagNotFound
		}
		return nil, tags.ErrTagQuery
	}

	return &tag, nil
}

func (r *TagsRepository) Update(tag models.ProductTag) (*models.ProductTag, error) {
	res, err := r.writeDB.Exec(QueryUpdateTag, tag.Title, tag.ID)
	if err != nil {
		return nil, tags.ErrTagQuery
	}

	affect, err := res.RowsAffected()
	if err != nil {
		return nil, tags.ErrTagQuery
	}
	if affect == 0 {
		return nil, tags.ErrTagNotFound
	}

	return &tag, nil
}

func (r *TagsRepository) GetProductTags(productId int) ([]*models.ProductTag, error) {
	list := []*models.ProductTag{}
	if err := r.readDB.Select(&list, QueryGetProductTags, productId); err != nil {
		return nil, tags.ErrTagQuery
	}
	return list, nil
}
