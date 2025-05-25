package repository

import (
  "Bus-Backend/internal/models"
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
