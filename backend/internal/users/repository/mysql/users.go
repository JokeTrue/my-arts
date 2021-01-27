package mysql

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/JokeTrue/my-arts/internal/models"
	"github.com/JokeTrue/my-arts/internal/users"
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

type UsersRepository struct {
	db *sqlx.DB
}

func NewUsersRepository(db *sqlx.DB) *UsersRepository {
	return &UsersRepository{db: db}
}

func (r *UsersRepository) Delete(id int) error {
	res, err := r.db.Exec(QueryDeleteUser, id)
	if err != nil {
		return users.ErrUserQuery
	}

	affect, err := res.RowsAffected()
	if err != nil {
		return users.ErrUserQuery
	}
	if affect == 0 {
		return users.ErrUserNotFound
	}

	return nil
}

func (r *UsersRepository) GetUsers() ([]models.User, error) {
	list := []models.User{}
	if err := r.db.Select(&list, QueryGetUsers); err != nil {
		return nil, users.ErrUserQuery
	}
	return list, nil
}

func (r *UsersRepository) Create(user models.User) (int, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}

	result, err := r.db.Exec(
		QueryCreateUser,
		user.Email, string(hashedPassword), user.FirstName, user.LastName, user.Age, user.Gender, user.Location, user.Biography,
	)
	if err != nil {
		return 0, users.ErrUserQuery
	}

	lastID, err := result.LastInsertId()
	if err != nil {
		return 0, users.ErrUserQuery
	}

	return int(lastID), nil
}

func (r *UsersRepository) GetUserByID(id int) (*models.User, error) {
	var user models.User
	if err := r.db.Get(&user, QueryGetUserByID, id); err != nil {
		if err == sql.ErrNoRows {
			return nil, users.ErrUserNotFound
		}
		return nil, users.ErrUserQuery
	}
	return &user, nil
}

func (r *UsersRepository) Update(user models.User) (*models.User, error) {
	res, err := r.db.Exec(
		QueryUpdateUser,
		user.FirstName, user.LastName, user.Age, user.Gender, user.Location, user.Biography, user.ID,
	)
	if err != nil {
		return nil, users.ErrUserQuery
	}

	affect, err := res.RowsAffected()
	if err != nil {
		return nil, users.ErrUserQuery
	}
	if affect == 0 {
		return nil, users.ErrUserNotFound
	}

	return &user, nil
}

func (r *UsersRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	if err := r.db.Get(&user, QueryGetUserByEmail, email); err != nil {
		if err == sql.ErrNoRows {
			return nil, users.ErrUserNotFound
		}
		return nil, users.ErrUserQuery
	}
	return &user, nil
}

func (r *UsersRepository) SearchUsers(query string) ([]*models.User, error) {
	list := []*models.User{}

	var filterExpression string
	if query != "" {
		filters := map[string]string{
			"first_name": query,
			"last_name":  query,
		}
		expressions := make([]string, 0, len(filters))
		for field, filter := range filters {
			expression := fmt.Sprintf("LOWER(%s) LIKE '%%%s%%'", field, filter)
			expressions = append(expressions, expression)
		}

		filterExpression = " WHERE " + strings.Join(expressions, " OR ")
	}

	searchQuery := fmt.Sprintf(QuerySearchUsers, filterExpression)
	if err := r.db.Select(&list, searchQuery); err != nil {
		return nil, users.ErrUserQuery
	}
	return list, nil
}

func (r *UsersRepository) GetUserFriends(id int) ([]*models.User, error) {
	list := []*models.User{}
	if err := r.db.Select(&list, QueryGetUserFriends, id, id); err != nil {
		return nil, users.ErrUserQuery
	}
	return list, nil
}
