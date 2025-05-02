package repository

import (
  "Bus-Backend/internal/models"
  "fmt"
  "gorm.io/gorm"
  "github.com/sirupsen/logrus"
)

type EmployeeRepoInit struct {
  generic *GenericRepoInit[*models.Employee]
}


func NewEmployeeRepo(db *gorm.DB, logger *logrus.Logger) (*EmployeeRepoInit, error) {
  genericRepo, err := NewGenericRepo[*models.Employee](db, logger)
  if err != nil {
    return nil, fmt.Errorf("failed to create EmployeeRepo: %w", err)
  }

  return &EmployeeRepoInit{generic: genericRepo}, nil
}


// using the CRUD operations from the generic repo 
// FindbyId CRUD operation receiving the entity id as an argument
func (r *EmployeeRepoInit) FindById(userId int) (*models.Employee, error) {
  return r.generic.FindById(userId)
}

// insert CRUD operation receiving the entity as an argument
func (r *EmployeeRepoInit) Insert(user *models.Employee) (*models.Employee, error) {
  return r.generic.Insert(user)
}

// Update CRUD operation receiving the entity as an argument
func (r *EmployeeRepoInit) Update(user *models.Employee) (*models.Employee, error) {
  return r.generic.Update(user)
}

// Delete CRUD operation receiving the entity id as an argument
func (r *EmployeeRepoInit) Delete(userId int) error {
  return r.generic.Delete(userId)
}

// Get CRUD operation returning all entities
func (r *EmployeeRepoInit) Get() ([]*models.Employee, error) {
  return r.generic.Get()
}
