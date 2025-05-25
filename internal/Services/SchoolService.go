package services

import "Bus-Backend/internal/models"

type SchoolService interface {
	Get() ([]*models.School, error)
	FindById(id int) (*models.School, error)
	Insert(user *models.School) (*models.School, error)
	Update(user *models.School) (*models.School, error)
	Delete(id int) error
}
