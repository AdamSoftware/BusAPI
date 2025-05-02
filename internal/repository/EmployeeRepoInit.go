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
func (r *EmployeeRepoInit) FindById(EmployeeId int) (*models.Employee, error) {
  return r.generic.FindById(EmployeeId)
}

// insert CRUD operation receiving the entity as an argument
func (r *EmployeeRepoInit) Insert(Employee *models.Employee) (*models.Employee, error) {
  return r.generic.Insert(Employee)
}

// Update CRUD operation receiving the entity as an argument
func (r *EmployeeRepoInit) Update(Employee *models.Employee) (*models.Employee, error) {
  return r.generic.Update(Employee)
}

// Delete CRUD operation receiving the entity id as an argument
func (r *EmployeeRepoInit) Delete(EmployeeId int) error {
  return r.generic.Delete(EmployeeId)
}

// Get CRUD operation returning all entities
func (r *EmployeeRepoInit) Get() ([]*models.Employee, error) {
  return r.generic.Get()
}
