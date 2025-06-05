package repository

import (
  "Bus-Backend/internal/models"
	"Bus-Backend/internal/Logging"
  "gorm.io/gorm"
  "fmt"
)


type RouteRepoInit struct {
  generic *GenericRepoInit[*models.Route]
}

func NewRouteRepo(db *gorm.DB) (*RouteRepoInit, error) {
  genericRepo, err := NewGenericRepo[*models.Route](db)
  if err != nil {
    return nil, fmt.Errorf("failed to create RouteRepo: %w", err)
  }

  return &RouteRepoInit{generic: genericRepo}, nil
}

// using the CRUD operations from the generic repo 
// FindbyId CRUD operation receiving the entity id as an argument
func (r *RouteRepoInit) FindById(RouteId int) (*models.Route, error) {
  return r.generic.FindById(RouteId)
}

// insert CRUD operation receiving the entity as an argument
func (r *RouteRepoInit) Insert(Route *models.Route) (*models.Route, error) {
  return r.generic.Insert(Route)
}

// Update CRUD operation receiving the entity as an argument
func (r *RouteRepoInit) Update(Route *models.Route) (*models.Route, error) {
  return r.generic.Update(Route)
}

// Delete CRUD operation receiving the entity id as an argument
func (r *RouteRepoInit) Delete(RouteId int) error {
  return r.generic.Delete(RouteId)
}

// Get CRUD operation returning all entities
func (r *RouteRepoInit) Get() ([]*models.Route, error) {
  return r.generic.Get()
}


// GetAllRoutesByStopId retrieves all routes associated with a specific stop ID
func (r *RouteRepoInit) GetAllRoutesByStopId(stopId int) ([]*models.Route, error) {
	var routes []*models.Route
	// using a join to get all the routes under the routeStops table 
	// where the stopId matches the one getting passed into it
	err := r.generic.db. 
	  Table("Route").
		Joins("JOIN RouteStops ON RouteStops.RouteId = Route.Id").
		Where("RouteStops.StopId = ?", stopId).
		Find(&routes).Error
	
	if err != nil {
		Logging.Logs.Warnf("failed to get routes by stop ID %d: %w", stopId, err)
		return nil, fmt.Errorf("failed to get routes by stop ID %d: %w", stopId, err) 
	}

	return routes, nil
}
