package services

import (
	"new-beego-api/models"
)

type ObjectService struct{}

func NewObjectService() *ObjectService {
	return &ObjectService{}
}

func (s *ObjectService) GetAll() map[string]*models.Object {
	return models.Objects
}

func (s *ObjectService) Get(objectId string) (*models.Object, error) {
	return models.GetOne(objectId)
}

func (s *ObjectService) Delete(objectId string) {
	models.Delete(objectId)
}
