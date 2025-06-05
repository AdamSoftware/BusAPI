package repository

import (
  "Bus-Backend/internal/models"
)

type RouteRepo interface {
  FindById(RouteId int) (*models.Route, error)
  Insert(Route *models.Route) (*models.Route, error)
  Update(Route *models.Route) (*models.Route, error)
  Delete(RouteId int) error
  Get() ([]*models.Route, error)
	GetAllRoutesByStopId(stopId int) ([]*models.Route, error)
}
