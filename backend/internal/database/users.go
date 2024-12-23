package database

import (
	"database/sql"
)

type UserDAO interface {
	GetUserByID(id string) (*User, error)
	CreateUser(user *UserInsert) error
	GetUserByAuth0(sub string) (*User, error)
}

type userDAO struct {
	db *sql.DB
}

type User struct {
	ID          string  `json:"id"`
	Username    string  `json:"username"`
	Email       string  `json:"email"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
	Region      *string `json:"region"`
	DateOfBirth *string `json:"date_of_birth"`
	IsActive    bool    `json:"is_active"`
	Auth0Sub    string  `json:"auth0_sub"`
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
	const query = `INSERT INTO users (username, email, region, auth0_sub, date_of_birth) VALUES ($1, $2, $3, $4, $5)`
	_, err := u.db.Exec(query, user.Username, user.Email, user.Region, user.Auth0Sub, user.DateOfBirth)
	if err != nil {
		return err
	}
	return nil
}

func (u *userDAO) GetUserByAuth0(sub string) (*User, error) {
	const query = `SELECT id, username, email, created_at, updated_at, region, date_of_birth, is_active, auth0_sub FROM users WHERE auth0_sub = $1`
	row := u.db.QueryRow(query, sub)
	user := &User{}
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.CreatedAt, &user.UpdatedAt, &user.Region, &user.DateOfBirth, &user.IsActive, &user.Auth0Sub)
	if err != nil {
		return nil, err
	}
	return user, nil
}

type UserInsert struct {
	Username    string  `json:"username"`
	Email       string  `json:"email"`
	Region      *string `json:"region"`
	Auth0Sub    string  `json:"auth0_sub"`
	DateOfBirth *string `json:"date_of_birth"`
}
