package services

import "Bus-Backend/internal/models"

type RouteService interface {
	Get() ([]*models.Route, error)
	FindById(id int) (*models.Route, error)
	Insert(user *models.Route) (*models.Route, error)
	Update(user *models.Route) (*models.Route, error)
	Delete(id int) error
}
