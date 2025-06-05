package repository

import (
  "Bus-Backend/internal/models"
	"Bus-Backend/internal/Logging"
  "gorm.io/gorm"
  "fmt"
)


type StopTypeRepoInit struct {
  generic *GenericRepoInit[*models.StopType]
}


func NewStopTypeRepo(db *gorm.DB) (*StopTypeRepoInit, error) {
  genericRepo, err := NewGenericRepo[*models.StopType](db)
  if err != nil {
    return nil, fmt.Errorf("failed to create StopTypeRepo: %w", err)
  }

  return &StopTypeRepoInit{generic: genericRepo}, nil
}


// using the CRUD operations from the generic repo 
// FindbyId CRUD operation receiving the entity id as an argument
func (r *StopTypeRepoInit) FindById(Id int) (*models.StopType, error) {
  return r.generic.FindById(Id)
}

// insert CRUD operation receiving the entity as an argument
func (r *StopTypeRepoInit) Insert(StopType *models.StopType) (*models.StopType, error) {
  return r.generic.Insert(StopType)
}

// Update CRUD operation receiving the entity as an argument
func (r *StopTypeRepoInit) Update(StopType *models.StopType) (*models.StopType, error) {
  return r.generic.Update(StopType)
}

// Delete CRUD operation receiving the entity id as an argument
func (r *StopTypeRepoInit) Delete(Id int) error {
  return r.generic.Delete(Id)
}

// Get CRUD operation returning all entities
func (r *StopTypeRepoInit) Get() ([]*models.StopType, error) {
  return r.generic.Get()
}

// GetByTypeName retrieves a StopType by its type name
func (r *StopTypeRepoInit) GetByTypeName(typeName string) (*models.StopType, error) {
	if typeName == "" {
		Logging.Logs.Warn("typeName cannot be empty")
		return nil, fmt.Errorf("typeName cannot be empty")
	}

	var stopType models.StopType
	err := r.generic.db.
		Where("TypeName = ?", typeName).
		First(&stopType).Error

	if err != nil {
		Logging.Logs.Warnf("failed to find StopType by TypeName: %v", err)
		return nil, err
	}

	return &stopType, nil
}


