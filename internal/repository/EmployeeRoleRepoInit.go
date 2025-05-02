package repository

import (
  "Bus-Backend/internal/models"
  "github.com/sirupsen/logrus"
  "gorm.io/gorm"
  "fmt"
)


type EmployeeRoleRepoInit struct {
  generic *GenericRepoInit[*models.EmployeeRole]
}


func NewEmployeeRoleRepo(db *gorm.DB, logger *logrus.Logger) (*EmployeeRoleRepoInit, error) {
  genericRepo, err := NewGenericRepo[*models.EmployeeRole](db, logger)
  if err != nil {
    return nil, fmt.Errorf("failed to create EmployeeRoleRepo: %w", err)
  }

  return &EmployeeRoleRepoInit{generic: genericRepo}, nil
}


// using the CRUD operations from the generic repo 
// FindbyId CRUD operation receiving the entity id as an argument
func (r *EmployeeRoleRepoInit) FindById(userId int) (*models.EmployeeRole, error) {
  return r.generic.FindById(userId)
}

// insert CRUD operation receiving the entity as an argument
func (r *EmployeeRoleRepoInit) Insert(user *models.EmployeeRole) (*models.EmployeeRole, error) {
  return r.generic.Insert(user)
}

// Update CRUD operation receiving the entity as an argument
func (r *EmployeeRoleRepoInit) Update(user *models.EmployeeRole) (*models.EmployeeRole, error) {
  return r.generic.Update(user)
}

// Delete CRUD operation receiving the entity id as an argument
func (r *EmployeeRoleRepoInit) Delete(userId int) error {
  return r.generic.Delete(userId)
}

// Get CRUD operation returning all entities
func (r *EmployeeRoleRepoInit) Get() ([]*models.EmployeeRole, error) {
  return r.generic.Get()
}
