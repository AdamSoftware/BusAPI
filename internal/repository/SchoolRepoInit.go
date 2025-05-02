package repository

import (
  "Bus-Backend/internal/models"
  "github.com/sirupsen/logrus"
  "gorm.io/gorm"
  "fmt"
)


type SchoolRepoInit struct {
  generic *GenericRepoInit[*models.School]
}


func NewSchoolRepo(db *gorm.DB, logger *logrus.Logger) (*SchoolRepoInit, error) {
  genericRepo, err := NewGenericRepo[*models.School](db, logger)
  if err != nil {
    return nil, fmt.Errorf("failed to create SchoolRepo: %w", err)
  }

  return &SchoolRepoInit{generic: genericRepo}, nil
}


// using the CRUD operations from the generic repo 
// FindbyId CRUD operation receiving the entity id as an argument
func (r *SchoolRepoInit) FindById(userId int) (*models.School, error) {
  return r.generic.FindById(userId)
}

// insert CRUD operation receiving the entity as an argument
func (r *SchoolRepoInit) Insert(user *models.School) (*models.School, error) {
  return r.generic.Insert(user)
}

// Update CRUD operation receiving the entity as an argument
func (r *SchoolRepoInit) Update(user *models.School) (*models.School, error) {
  return r.generic.Update(user)
}

// Delete CRUD operation receiving the entity id as an argument
func (r *SchoolRepoInit) Delete(userId int) error {
  return r.generic.Delete(userId)
}

// Get CRUD operation returning all entities
func (r *SchoolRepoInit) Get() ([]*models.School, error) {
  return r.generic.Get()
}
