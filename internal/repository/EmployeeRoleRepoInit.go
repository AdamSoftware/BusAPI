package repository

import (
  "Bus-Backend/internal/models"
  "gorm.io/gorm"
  "fmt"
)


type EmployeeRoleRepoInit struct {
  generic *GenericRepoInit[*models.EmployeeRole]
}


func NewEmployeeRoleRepo(db *gorm.DB) (*EmployeeRoleRepoInit, error) {
  genericRepo, err := NewGenericRepo[*models.EmployeeRole](db)
  if err != nil {
    return nil, fmt.Errorf("failed to create EmployeeRoleRepo: %w", err)
  }

  return &EmployeeRoleRepoInit{generic: genericRepo}, nil
}


// using the CRUD operations from the generic repo 
// FindbyId CRUD operation receiving the entity id as an argument
func (r *EmployeeRoleRepoInit) FindById(EmployeeRoleId int) (*models.EmployeeRole, error) {
  return r.generic.FindById(EmployeeRoleId)
}

// insert CRUD operation receiving the entity as an argument
func (r *EmployeeRoleRepoInit) Insert(EmployeeRole *models.EmployeeRole) (*models.EmployeeRole, error) {
  return r.generic.Insert(EmployeeRole)
}

// Update CRUD operation receiving the entity as an argument
func (r *EmployeeRoleRepoInit) Update(EmployeeRole *models.EmployeeRole) (*models.EmployeeRole, error) {
  return r.generic.Update(EmployeeRole)
}

// Delete CRUD operation receiving the entity id as an argument
func (r *EmployeeRoleRepoInit) Delete(EmployeeRoleId int) error {
  return r.generic.Delete(EmployeeRoleId)
}

// Get CRUD operation returning all entities
func (r *EmployeeRoleRepoInit) Get() ([]*models.EmployeeRole, error) {
  return r.generic.Get()
}
