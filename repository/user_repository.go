package repository

import (
	"database/sql"
	"fmt"

	"github.com/go-embed-go-web/models"
)

type UserRepository interface {
	Create(payload models.User) error
	GetUserLogin(payload models.User) (*models.User, error)
	GetUserByID(email string) (*models.User, error)
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

func (u *userRepository) Create(payload models.User) error {
	query := "INSERT INTO User (username, password, role) VALUES($1, $2,$3)"
	_, err := u.db.Exec(query, payload.Username, payload.Password, payload.Role)
	if err != nil {
		return err
	}
	return nil
}

func (r *userRepository) GetUserLogin(payload models.User) (*models.User, error) {
	query := `SELECT user_id, password , role FROM users WHERE email = ? AND password = ?`
	row := r.db.QueryRow(query, payload.ID, payload.Password)

	var user models.User
	err := row.Scan(&user.ID, &user.Password, &user.Role)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("user not found")
	} else if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) GetUserByID(email string) (*models.User, error) {
	query := `SELECT user_id,  password FROM users WHERE email = ?`
	row := r.db.QueryRow(query, email)

	var user models.User
	err := row.Scan(&user.ID, &user.Password)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("user not found")
	} else if err != nil {
		return nil, err
	}

	return &user, nil
}
