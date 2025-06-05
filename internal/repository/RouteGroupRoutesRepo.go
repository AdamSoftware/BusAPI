package repository

import (
	"Bus-Backend/internal/models"
)

type RouteGroupRoutesRepo struct {	
	FindById func(Id int) (*models.RouteGroupRoutes, error)
	Insert   func(RouteGroupRoutes *models.RouteGroupRoutes) (*models.RouteGroupRoutes, error)
	Update   func(RouteGroupRoutes *models.RouteGroupRoutes) (*models.RouteGroupRoutes, error)
	Delete   func(Id int) error

	// Custom Methods

	// fetch all routes for a given route group
	GetByRouteGroupId func(routeGroupId int) ([]*models.RouteGroupRoutes, error)
	// fetch all route groups for a given route
	GetByRouteId func(routeId int) ([]*models.RouteGroupRoutes, error)
	// find a route group route by routeGroupId and routeId
	DeleteByRouteGroupAndRoute func(routeGroupId int, routeId int) error
}
