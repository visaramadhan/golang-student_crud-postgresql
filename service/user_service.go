package service

import (
	"fmt"

	"github.com/go-embed-go-web/models"
	"github.com/go-embed-go-web/repository"
)

// Interface UserService tanpa CreateService
type UserService interface {
	RegisterNewUser(payload models.User) error
	LoginService(payload models.User) (*models.User, error)
}

// Struct userService yang mengimplementasikan UserService
type userService struct {
	repo repository.UserRepository
}

// Implementasi RegisterNewUser pada userService
func (u *userService) RegisterNewUser(payload models.User) error {
	if payload.ID == 0 || payload.Username == "" || payload.Password == "" || payload.Role == "" {
		return fmt.Errorf("all payload is required")
	}

	err := u.repo.Create(payload)
	if err != nil {
		return fmt.Errorf("failed to create user: %s", err)
	}
	return nil
}

// Implementasi LoginService pada userService
func (u *userService) LoginService(payload models.User) (*models.User, error) {
	if payload.Username == "" || payload.Password == "" {
		return nil, fmt.Errorf("username and password are required")
	}
	users, err := u.repo.GetUserLogin(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to find user: %s", err)
	}
	return users, nil
}

// Fungsi pembuat untuk userService
func CreateService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}
