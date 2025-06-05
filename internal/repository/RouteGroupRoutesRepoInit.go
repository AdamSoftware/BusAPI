package repository 


import (
  "Bus-Backend/internal/models"
	"Bus-Backend/internal/Logging"
  "gorm.io/gorm"
  "fmt"
)


type RouteGroupRoutesRepoInit struct {
  generic *GenericRepoInit[*models.RouteGroupRoutes]
}


func NewRouteGroupRoutesRepo(db *gorm.DB) (*RouteGroupRoutesRepoInit, error) {
  genericRepo, err := NewGenericRepo[*models.RouteGroupRoutes](db)
  if err != nil {
    return nil, fmt.Errorf("failed to create RouteGroupRoutesRepo: %w", err)
  }

  return &RouteGroupRoutesRepoInit{generic: genericRepo}, nil
}


// using the CRUD operations from the generic repo 
// FindbyId CRUD operation receiving the entity id as an argument
func (r *RouteGroupRoutesRepoInit) FindById(Id int) (*models.RouteGroupRoutes, error) {
  return r.generic.FindById(Id)
}

// insert CRUD operation receiving the entity as an argument
func (r *RouteGroupRoutesRepoInit) Insert(RouteGroupRoutes *models.RouteGroupRoutes) (*models.RouteGroupRoutes, error) {
  return r.generic.Insert(RouteGroupRoutes)
}

// Update CRUD operation receiving the entity as an argument
func (r *RouteGroupRoutesRepoInit) Update(RouteGroupRoutes *models.RouteGroupRoutes) (*models.RouteGroupRoutes, error) {
  return r.generic.Update(RouteGroupRoutes)
}

// Delete CRUD operation receiving the entity id as an argument
func (r *RouteGroupRoutesRepoInit) Delete(Id int) error {
  return r.generic.Delete(Id)
}

// Get CRUD operation returning all entities
func (r *RouteGroupRoutesRepoInit) Get() ([]*models.RouteGroupRoutes, error) {
  return r.generic.Get()
}

// custom methods

// GetByRouteGroupId retrieves all RouteGroupRoutes associated with a specific route group ID
func (r *RouteGroupRoutesRepoInit) GetByRouteGroupId(routeGroupId int) ([]*models.RouteGroupRoutes, error) {
	if routeGroupId <= 0 {
		Logging.Logs.Warn("routeGroupId cannot be zero")
		return nil, fmt.Errorf("routeGroupId cannot be nil")
	}

	var routes []*models.RouteGroupRoutes
	if err := r.generic.db.Where("route_group_id = ?", routeGroupId).Find(&routes).Error; err != nil {
		return nil, fmt.Errorf("failed to get routes by route group ID: %w", err)
	}
	return routes, nil
}


// GetByRouteId retrieves all RouteGroupRoutes associated with a specific route ID
func (r *RouteGroupRoutesRepoInit) GetByRouteId(routeId int) ([]*models.RouteGroupRoutes, error) {
	if routeId <= 0 {
		Logging.Logs.Warn("routeId cannot be zero")
		return nil, fmt.Errorf("routeId cannot be nil")
	}

	var routes []*models.RouteGroupRoutes
	if err := r.generic.db.Where("route_id = ?", routeId).Find(&routes).Error; err != nil {
		return nil, fmt.Errorf("failed to get routes by route ID: %w", err)
	}
	return routes, nil
}

// DeleteByRouteGroupAndRoute deletes a RouteGroupRoutes by routeGroupId and routeId
func (r *RouteGroupRoutesRepoInit) DeleteByRouteGroupAndRoute(routeGroupId int, routeId int) error {
	if routeGroupId <= 0 || routeId <= 0 {
		Logging.Logs.Warn("routeGroupId and routeId cannot be zero")
		return fmt.Errorf("routeGroupId and routeId cannot be nil")
	}

	if err := r.generic.db.Where("route_group_id = ? AND route_id = ?", routeGroupId, routeId).Delete(&models.RouteGroupRoutes{}).Error; err != nil {
		return fmt.Errorf("failed to delete RouteGroupRoutes by routeGroupId and routeId: %w", err)
	}
	return nil
}
