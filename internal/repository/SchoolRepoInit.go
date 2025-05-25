package repository

import (
  "Bus-Backend/internal/models"
  "gorm.io/gorm"
  "fmt"
)


type SchoolRepoInit struct {
  generic *GenericRepoInit[*models.School]
}


func NewSchoolRepo(db *gorm.DB) (*SchoolRepoInit, error) {
  genericRepo, err := NewGenericRepo[*models.School](db)
  if err != nil {
    return nil, fmt.Errorf("failed to create SchoolRepo: %w", err)
  }

  return &SchoolRepoInit{generic: genericRepo}, nil
}


// using the CRUD operations from the generic repo 
// FindbyId CRUD operation receiving the entity id as an argument
func (r *SchoolRepoInit) FindById(SchoolId int) (*models.School, error) {
  return r.generic.FindById(SchoolId)
}

// insert CRUD operation receiving the entity as an argument
func (r *SchoolRepoInit) Insert(School *models.School) (*models.School, error) {
  return r.generic.Insert(School)
}

// Update CRUD operation receiving the entity as an argument
func (r *SchoolRepoInit) Update(School *models.School) (*models.School, error) {
  return r.generic.Update(School)
}

// Delete CRUD operation receiving the entity id as an argument
func (r *SchoolRepoInit) Delete(SchoolId int) error {
  return r.generic.Delete(SchoolId)
}

// Get CRUD operation returning all entities
func (r *SchoolRepoInit) Get() ([]*models.School, error) {
  return r.generic.Get()
}

