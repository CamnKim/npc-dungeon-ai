package database

import (
	"database/sql"
)

type UserDAO interface {
	GetUserByID(id string) (*User, error)
	CreateUser(user *UserInsert) error
}

type userDAO struct {
	db *sql.DB
}

type User struct {
	ID        string  `json:"id"`
	Username  string  `json:"username"`
	Email     string  `json:"email"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
	Region    *string `json:"region"`
}

func (s *service) User() UserDAO {
	return &userDAO{
		db: s.db,
	}
}

func (u *userDAO) GetUserByID(id string) (*User, error) {
	const query = `SELECT id, username, email, created_at, updated_at, region FROM users WHERE id = $1`
	row := u.db.QueryRow(query, id)
	user := &User{}
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.CreatedAt, &user.UpdatedAt, &user.Region)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userDAO) CreateUser(user *UserInsert) error {
	const query = `INSERT INTO users (username, email, region) VALUES ($1, $2, $3)`
	_, err := u.db.Exec(query, user.Username, user.Email, user.Region)
	if err != nil {
		return err
	}
	return nil
}

type UserInsert struct {
	Username string  `json:"username"`
	Email    string  `json:"email"`
	Region   *string `json:"region"`
}
