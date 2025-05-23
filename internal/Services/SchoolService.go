package services

import "Bus-Backend/internal/models"

type SchoolService interface {
	Get() ([]*models.User, error)
	FindById(id int) (*models.User, error)
	Insert(user *models.User) (*models.User, error)
	Update(user *models.User) (*models.User, error)
	Delete(id int) error
}
