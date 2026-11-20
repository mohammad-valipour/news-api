package services

import (
	"database/sql"
	"errors"
	"news/models"
	"news/repository"
	"strings"
)

type UserService struct {
	UserRepo *repository.UserRepository
}

func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{
		UserRepo: userRepo,
	}
}

// Create creates a new user
func (s *UserService) Create(user *models.User) error {
	user.FirstName = strings.TrimSpace(user.FirstName)
	user.LastName = strings.TrimSpace(user.LastName)
	user.Email = strings.TrimSpace(user.Email)

	if user.FirstName == "" {
		return errors.New("first name is required")
	}

	if user.LastName == "" {
		return errors.New("last name is required")
	}

	if user.Email == "" {
		return errors.New("email is required")
	}

	if user.PasswordHash == "" {
		return errors.New("password is required")
	}

	if user.Role != "admin" && user.Role != "author" {
		return errors.New("role is invalid")
	}

	_, err := s.UserRepo.GetByEmail(user.Email)
	if err == nil {
		return errors.New("email already exists")
	}

	if !errors.Is(err, sql.ErrNoRows) {
		return err
	}

	return s.UserRepo.Create(user)
}

// GetByID returns a user by ID
func (s *UserService) GetByID(id int64) (*models.User, error) {
	if id <= 0 {
		return nil, errors.New("invalid ID")
	}

	return s.UserRepo.GetByID(id)
}

// GetByEmail returns a user by email
func (s *UserService) GetByEmail(email string) (*models.User, error) {
	email = strings.TrimSpace(email)

	if email == "" {
		return nil, errors.New("invalid email")
	}

	return s.UserRepo.GetByEmail(email)
}

// GetAll returns all users
func (s *UserService) GetAll() ([]models.User, error) {
	return s.UserRepo.GetAll()
}
