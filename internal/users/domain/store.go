package domain

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

// Store interface for User persistence layer
type Store interface {
	delete(id int) error
	getAll() ([]User, error)
	getByID(id int) (*User, error)
	create(user User) (int, error)
	update(user User) (*User, error)
	getByEmail(email string) (*User, error)
}

// UserStore for persistence
type UserStore struct {
	db *sqlx.DB
}

// UserStore creates new UserStore for Users
func NewStore(db *sqlx.DB) *UserStore {
	return &UserStore{db: db}
}

func (s *UserStore) delete(id int) error {
	res, err := s.db.Exec("DELETE from users where id = ?", id)
	if err != nil {
		return ErrUserQuery
	}

	affect, err := res.RowsAffected()
	if err != nil {
		return ErrUserQuery
	}
	if affect == 0 {
		return ErrUserNotFound
	}

	return nil
}

func (s *UserStore) getAll() ([]User, error) {
	users := []User{}
	if err := s.db.Select(&users, "SELECT * FROM users WHERE id > 1"); err != nil {
		return nil, ErrUserQuery
	}
	return users, nil
}

func (s *UserStore) getByID(id int) (*User, error) {
	var user User
	if err := s.db.Get(&user, "SELECT * FROM users WHERE id = ?", id); err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrUserNotFound
		}
		return nil, ErrUserQuery
	}
	return &user, nil
}

func (s *UserStore) getByEmail(email string) (*User, error) {
	var user User
	if err := s.db.Get(&user, "SELECT * FROM users WHERE email = ?", email); err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrUserNotFound
		}
		return nil, ErrUserQuery
	}
	return &user, nil
}

func (s *UserStore) create(user User) (int, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}

	query := `INSERT INTO users 
    		  (email, password, first_name, last_name, age, gender, location, biography) 
              VALUES (?, ?, ?, ?, ?, ?, ?, ?)`
	result, err := s.db.Exec(
		query,
		user.Email, string(hashedPassword), user.FirstName, user.LastName, user.Age, user.Gender, user.Location, user.Biography,
	)
	if err != nil {
		return 0, ErrUserQuery
	}

	lastID, err := result.LastInsertId()
	if err != nil {
		return 0, ErrUserQuery
	}

	return int(lastID), nil
}

func (s *UserStore) update(user User) (*User, error) {
	query := "UPDATE users SET first_name=?, last_name=?, age=?, gender=?, city=?, biography=? WHERE id=?"
	res, err := s.db.Exec(
		query,
		user.FirstName, user.LastName, user.Age, user.Gender, user.Location, user.Biography, user.ID,
	)
	if err != nil {
		return nil, ErrUserQuery
	}

	affect, err := res.RowsAffected()
	if err != nil {
		return nil, ErrUserQuery
	}
	if affect == 0 {
		return nil, ErrUserNotFound
	}

	return &user, nil
}
