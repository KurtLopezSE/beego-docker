package services

import (
	"new-beego-api/models"
)

type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
}

func (s *UserService) GetAll() map[string]*models.User {
	return models.GetAllUsers()
}

func (s *UserService) Get(uid string) (*models.User, error) {
	return models.GetUser(uid)
}

func (s *UserService) Add(user models.User) string {
	return models.AddUser(user)
}

func (s *UserService) Update(uid string, user *models.User) (*models.User, error) {
	return models.UpdateUser(uid, user)
}

func (s *UserService) Delete(uid string) {
	models.DeleteUser(uid)
}

func (s *UserService) Login(username, password string) bool {
	return models.Login(username, password)
}
