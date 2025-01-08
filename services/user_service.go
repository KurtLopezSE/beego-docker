package services

import (
	"new-beego-api/models"
)

// UserService provides methods for user operations
type UserService struct{}

// NewUserService creates a new UserService
func NewUserService() *UserService {
	return &UserService{}
}

// AddUser adds a new user
func (s *UserService) AddUser(user *models.User) (int64, error) {
	return models.AddUser(user)
}

// GetUser retrieves a user by ID
func (s *UserService) GetUser(id int64) (*models.User, error) {
	return models.GetUser(id)
}

// GetAllUsers retrieves all users
func (s *UserService) GetAllUsers() ([]models.User, error) {
	return models.GetAllUsers()
}

// UpdateUser updates an existing user
func (s *UserService) UpdateUser(user *models.User) error {
	return models.UpdateUser(user)
}

// DeleteUser deletes a user by ID
func (s *UserService) DeleteUser(id int64) error {
	return models.DeleteUser(id)
}
