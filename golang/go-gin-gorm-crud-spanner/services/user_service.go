package services

import (
	"github.com/gitish/polyglot_training/models"
	"github.com/gitish/polyglot_training/repositories"
)

// UserService defines the interface for user-related operations.
type UserService interface {
	CreateUser(user models.User) (models.User, error)
	GetAllUsers() ([]models.User, error)
	GetUserByID(id uint) (models.User, error)
	UpdateUser(user models.User) (models.User, error)
	DeleteUser(id uint) error
}

// userServiceImpl implements the UserService interface.
type userServiceImpl struct {
	repo repositories.UserRepository
}

// NewUserService creates a new instance of userServiceImpl.
func NewUserService(repo repositories.UserRepository) UserService {
	return &userServiceImpl{repo: repo}
}

// CreateUser creates a new user.
func (s *userServiceImpl) CreateUser(user models.User) (models.User, error) {
	return s.repo.CreateUser(user)
}

// GetAllUsers retrieves all users.
func (s *userServiceImpl) GetAllUsers() ([]models.User, error) {
	return s.repo.GetAllUsers()
}

// GetUserByID retrieves a user by their ID.
func (s *userServiceImpl) GetUserByID(id uint) (models.User, error) {
	return s.repo.GetUserByID(id)
}

// UpdateUser updates an existing user.
func (s *userServiceImpl) UpdateUser(user models.User) (models.User, error) {
	return s.repo.UpdateUser(user)
}

// DeleteUser deletes a user by their ID.
func (s *userServiceImpl) DeleteUser(id uint) error {
	return s.repo.DeleteUser(id)
}
