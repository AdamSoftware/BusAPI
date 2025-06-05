package services

import (
	"Bus-Backend/internal/models"
	"Bus-Backend/internal/dto"
)


type RouteStopsService interface {
	// crud methods
  FindById(Id int) (*models.RouteStops, error)
  Insert(RouteStops *models.RouteStops) (*models.RouteStops, error)
  Update(RouteStops *models.RouteStops) (*models.RouteStops, error)
  Delete(Id int) error
  Get() ([]*models.RouteStops, error)

	// custom methods
	FindByRouteAndStop(routeId int, stopId int) ([]*models.RouteStops, error)
	DeleteByRouteAndStop(routeId int, stopId int) error
	// will be displayed asc order
  GetByRouteId(routeId int) ([]*models.RouteStops, error)
	UpdateStopOrder(routeId int, stopId int, newOrder int) error

	// getting the ordered stops with details about the routes like name, id, etc.
	GetOrderedStopsWithDetails(routeId int) ([]dto.OrderStopsInfo, error)
}
