package repository

import (
  "Bus-Backend/internal/models"
	"Bus-Backend/internal/Logging"
  "gorm.io/gorm"
  "fmt"
)


type StopsRepoInit struct {
  generic *GenericRepoInit[*models.Stops]
}


func NewStopsRepo(db *gorm.DB) (*StopsRepoInit, error) {
  genericRepo, err := NewGenericRepo[*models.Stops](db)
  if err != nil {
    return nil, fmt.Errorf("failed to create StopsRepo: %w", err)
  }

  return &StopsRepoInit{generic: genericRepo}, nil
}


// using the CRUD operations from the generic repo 
// FindbyId CRUD operation receiving the entity id as an argument
func (r *StopsRepoInit) FindById(Id int) (*models.Stops, error) {
  return r.generic.FindById(Id)
}

// insert CRUD operation receiving the entity as an argument
func (r *StopsRepoInit) Insert(Stops *models.Stops) (*models.Stops, error) {
  return r.generic.Insert(Stops)
}

// Update CRUD operation receiving the entity as an argument
func (r *StopsRepoInit) Update(Stops *models.Stops) (*models.Stops, error) {
  return r.generic.Update(Stops)
}

// Delete CRUD operation receiving the entity id as an argument
func (r *StopsRepoInit) Delete(Id int) error {
  return r.generic.Delete(Id)
}

// Get CRUD operation returning all entities
func (r *StopsRepoInit) Get() ([]*models.Stops, error) {
  return r.generic.Get()
}

// IsStopUsed checks if a stop is associated with any routes
func (r *StopsRepoInit) IsStopUsed(stopId int) (bool, error) {
	var count int64
	err := r.generic.db.
		Model(&models.RouteStops{}).
		Where("StopId = ?", stopId).
		Count(&count).Error

	if err != nil {
		Logging.Logs.Error("error checking if stop is used: ", err)
		return false, err
	}

	return count > 0, nil
}







