package services

import "Bus-Backend/internal/models"

type BusesService interface {
	Get() ([]*models.Bus, error)
	FindById(id int) (*models.Bus, error)
	Insert(user *models.Bus) (*models.Bus, error)
	Update(user *models.Bus) (*models.Bus, error)
	Delete(id int) error
}
