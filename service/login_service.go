package service

import (
	"database/sql"
	"errors"
	"fmt"
)

// LoginService interface untuk login, logout, dan pengecekan status login
type LoginService interface {
	Login(username, password string) error
	Logout()
	IsLoggedIn() bool
}

// Struct loginService dengan atribut tambahan untuk database
type loginService struct {
	db       *sql.DB
	loggedIn bool
}

// NewLoginService membuat instance loginService dengan koneksi database
func NewLoginService(db *sql.DB) LoginService {
	return &loginService{db: db, loggedIn: false}
}

// Login memeriksa kredensial pengguna di database
func (s *loginService) Login(username, password string) error {
	var dbPassword string
	query := "SELECT password FROM users WHERE username = $1"

	// Jalankan query untuk mendapatkan password dari username
	err := s.db.QueryRow(query, username).Scan(&dbPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("username not found")
		}
		return fmt.Errorf("database error: %v", err)
	}

	// Verifikasi kecocokan password
	if password == dbPassword {
		s.loggedIn = true
		return nil
	}
	return errors.New("invalid credentials")
}

// Logout mengubah status login menjadi false
func (s *loginService) Logout() {
	s.loggedIn = false
}

// IsLoggedIn mengembalikan status login
func (s *loginService) IsLoggedIn() bool {
	return s.loggedIn
}
