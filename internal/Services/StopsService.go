package services

import (
	"Bus-Backend/internal/models"
)


type StopsService interface {
	FindById(id int) (*models.Stops, error)
	Insert(stop *models.Stops) (*models.Stops, error)
	Update(stop *models.Stops) (*models.Stops, error)
	Delete(id int) error
	Get() ([]*models.Stops, error)
	DeleteIfUnused(stopId int) error
}
