package repository

import (
  "Bus-Backend/internal/models"
  "github.com/sirupsen/logrus"
  "gorm.io/gorm"
  "fmt"
)


type RouteRepoInit struct {
  generic *GenericRepoInit[*models.Route]
}

func NewRouteRepo(db *gorm.DB, logger *logrus.Logger) (*RouteRepoInit, error) {
  genericRepo, err := NewGenericRepo[*models.Route](db, logger)
  if err != nil {
    return nil, fmt.Errorf("failed to create RouteRepo: %w", err)
  }

  return &RouteRepoInit{generic: genericRepo}, nil
}

// using the CRUD operations from the generic repo 
// FindbyId CRUD operation receiving the entity id as an argument
func (r *RouteRepoInit) FindById(userId int) (*models.Route, error) {
  return r.generic.FindById(userId)
}

// insert CRUD operation receiving the entity as an argument
func (r *RouteRepoInit) Insert(user *models.Route) (*models.Route, error) {
  return r.generic.Insert(user)
}

// Update CRUD operation receiving the entity as an argument
func (r *RouteRepoInit) Update(user *models.Route) (*models.Route, error) {
  return r.generic.Update(user)
}

// Delete CRUD operation receiving the entity id as an argument
func (r *RouteRepoInit) Delete(userId int) error {
  return r.generic.Delete(userId)
}

// Get CRUD operation returning all entities
func (r *RouteRepoInit) Get() ([]*models.Route, error) {
  return r.generic.Get()
}
