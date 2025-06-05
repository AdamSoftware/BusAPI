package services

import (
	"Bus-Backend/internal/models"
)

type RouteGroupRoutesService interface {
	Get() ([]*models.RouteGroupRoutes, error)
	FindById(id int) (*models.RouteGroupRoutes, error)
	Insert(user *models.RouteGroupRoutes) (*models.RouteGroupRoutes, error)
	Update(user *models.RouteGroupRoutes) (*models.RouteGroupRoutes, error)
	Delete(id int) error

	// Custom Methods
	GetByRouteGroupId(routeGroupId int) ([]*models.RouteGroupRoutes, error)
	GetByRouteId(routeId int) ([]*models.RouteGroupRoutes, error)
	DeleteByRouteGroupAndRoute(routeGroupId int, routeId int) error
}

