package repository

import (
  "Bus-Backend/internal/models"
  "fmt"
  "gorm.io/gorm"
  "github.com/sirupsen/logrus"
)


type BusRepoInit struct {
  generic *GenericRepoInit[*models.Bus]
}


func NewBusRepo(db *gorm.DB, logger *logrus.Logger) (*BusRepoInit, error) {
  // Create a new instance of GenericRepoInit for Bus model
  genericRepo, err := NewGenericRepo[*models.Bus](db)
  if err != nil {
    return nil, fmt.Errorf("failed to create BusRepo: %w", err)
  }

  return &BusRepoInit{generic: genericRepo}, nil
}


// using the CRUD operations from the generic repo 
// FindbyId CRUD operation receiving the entity id as an argument
func (r *BusRepoInit) FindById(BusId int) (*models.Bus, error) {
  return r.generic.FindById(BusId)
}

// insert CRUD operation receiving the entity as an argument
func (r *BusRepoInit) Insert(Bus *models.Bus) (*models.Bus, error) {
  return r.generic.Insert(Bus)
}

// Update CRUD operation receiving the entity as an argument
func (r *BusRepoInit) Update(Bus *models.Bus) (*models.Bus, error) {
  return r.generic.Update(Bus)
}

// Delete CRUD operation receiving the entity id as an argument
func (r *BusRepoInit) Delete(BusId int) error {
  return r.generic.Delete(BusId)
}

// Get CRUD operation returning all entities
func (r *BusRepoInit) Get() ([]*models.Bus, error) {
  return r.generic.Get()
}
