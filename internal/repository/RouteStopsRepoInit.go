package repository

import (
  "Bus-Backend/internal/models"
	"Bus-Backend/internal/Logging"
  "gorm.io/gorm"
  "fmt"
	"Bus-Backend/internal/dto"
)


type RouteStopsRepoInit struct {
  generic *GenericRepoInit[*models.RouteStops]
}


func NewRouteStopsRepo(db *gorm.DB) (*RouteStopsRepoInit, error) {
  genericRepo, err := NewGenericRepo[*models.RouteStops](db)
  if err != nil {
    return nil, fmt.Errorf("failed to create RouteStopsRepo: %w", err)
  }

  return &RouteStopsRepoInit{generic: genericRepo}, nil
}


// using the CRUD operations from the generic repo 
// FindbyId CRUD operation receiving the entity id as an argument
func (r *RouteStopsRepoInit) FindById(Id int) (*models.RouteStops, error) {
  return r.generic.FindById(Id)
}

// insert CRUD operation receiving the entity as an argument
func (r *RouteStopsRepoInit) Insert(RouteStops *models.RouteStops) (*models.RouteStops, error) {
  return r.generic.Insert(RouteStops)
}

// Update CRUD operation receiving the entity as an argument
func (r *RouteStopsRepoInit) Update(RouteStops *models.RouteStops) (*models.RouteStops, error) {
  return r.generic.Update(RouteStops)
}

// Delete CRUD operation receiving the entity id as an argument
func (r *RouteStopsRepoInit) Delete(Id int) error {
  return r.generic.Delete(Id)
}

// Get CRUD operation returning all entities
func (r *RouteStopsRepoInit) Get() ([]*models.RouteStops, error) {
  return r.generic.Get()
}

// custom methods

// FindByRouteAndStop retrieves all RouteStops associated with a specific route ID
func (r *RouteStopsRepoInit) FindByRouteAndStop(routeId, stopId int) ([]*models.RouteStops, error) {
	if routeId <= 0 || stopId <= 0 {
		Logging.Logs.Warn("routeId and stopId cannot be zero")
		return nil, fmt.Errorf("routeId and stopId cannot be nil")
	}

	var routeStops []*models.RouteStops
	err := r.generic.db.
		Where("RouteId = ? AND StopId = ?", routeId, stopId).
		Find(&routeStops).Error

	if err != nil {
		return nil, fmt.Errorf("failed to find RouteStops by routeId and stopId: %w", err)
	}

	return routeStops, nil
}

// DeleteByRouteAndStop deletes RouteStops associated with a specific route ID and stop ID
func (r *RouteStopsRepoInit) DeleteByRouteAndStop(routeId, stopId int) error {
	if routeId <= 0 || stopId <= 0 {
		Logging.Logs.Warn("routeId and stopId cannot be zero")
		return fmt.Errorf("routeId and stopId cannot be nil")
	}

	err := r.generic.db.
		Where("RouteId = ? AND StopId = ?", routeId, stopId).
		Delete(&models.RouteStops{}).Error

	if err != nil {
		return fmt.Errorf("failed to delete RouteStops by routeId and stopId: %w", err)
	}

	return nil
}

// GetByRouteId retrieves all RouteStops associated with a specific route ID
func (r *RouteStopsRepoInit) GetByRouteId(routeId int) ([]*models.RouteStops, error) {
	if routeId == 0 {
		Logging.Logs.Warn("routeId cannot be zero")
		return nil, fmt.Errorf("routeId cannot be nil")
	}

	var routeStops []*models.RouteStops
	err := r.generic.db.
		Where("RouteId = ?", routeId).
		Find(&routeStops).Error

	if err != nil {
		return nil, fmt.Errorf("failed to get RouteStops by routeId: %w", err)
	}

	return routeStops, nil
}

//UpdateStopOrder updates the order of stops in a route
func (r *RouteStopsRepoInit) UpdateStopOrder(routeId, stopId, newOrder int) error {
	if routeId <= 0 || stopId <= 0 {
		Logging.Logs.Warn("routeId and stopId cannot be zero")
		return fmt.Errorf("routeId and stopId cannot be nil")
	}

	err := r.generic.db.
		Model(&models.RouteStops{}).
		Where("RouteId = ? AND StopId = ?", routeId, stopId).
		Update("StopOrder", newOrder).Error

	if err != nil {
		return fmt.Errorf("failed to update stop order: %w", err)
	}

	return nil
}

// GetOrderStopsWithDetails retrieves all RouteStops with their associated stop details for a specific route ID
func (r *RouteStopsRepoInit) GetOrderedStopsWithDetails(routeId int) ([]dto.OrderStopsInfo, error) {
	if routeId == 0 {
		Logging.Logs.Warn("routeId cannot be zero")
		return nil, fmt.Errorf("routeId cannot be nil")
	}

	var orderedStops []dto.OrderStopsInfo
	err := r.generic.db.
		Table("RouteStops").
		Select("RouteStops.*, Stops.StopName as StopName, Stops.Id as StopId").
		Joins("JOIN Stops ON RouteStops.StopId = Stops.Id").
		Where("RouteStops.RouteId = ?", routeId).
		Order("RouteStops.StopOrder ASC").
		Scan(&orderedStops).Error

	if err != nil {
		return nil, fmt.Errorf("failed to get ordered stops with details: %w", err)
	}

	return orderedStops, nil
}




