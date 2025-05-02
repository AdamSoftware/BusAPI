package repository


import (
  "Bus-Backend/internal/models"
  "github.com/sirupsen/logrus"
  "gorm.io/gorm"
  "fmt"
)


type BusStudentRepoInit struct {
  generic *GenericRepoInit[*models.BusStudent]
}


func NewBusStudentRepo(db *gorm.DB, logger *logrus.Logger) (*BusStudentRepoInit, error) {
  genericRepo, err := NewGenericRepo[*models.BusStudent](db, logger)
  if err != nil {
    return nil, fmt.Errorf("failed to create BusStudentRepo: %w", err)
  }

  return &BusStudentRepoInit{generic: genericRepo}, nil
}


// using the CRUD operations from the generic repo 
// FindbyId CRUD operation receiving the entity id as an argument
func (r *BusStudentRepoInit) FindById(userId int) (*models.BusStudent, error) {
  return r.generic.FindById(userId)
}

// insert CRUD operation receiving the entity as an argument
func (r *BusStudentRepoInit) Insert(user *models.BusStudent) (*models.BusStudent, error) {
  return r.generic.Insert(user)
}

// Update CRUD operation receiving the entity as an argument
func (r *BusStudentRepoInit) Update(user *models.BusStudent) (*models.BusStudent, error) {
  return r.generic.Update(user)
}

// Delete CRUD operation receiving the entity id as an argument
func (r *BusStudentRepoInit) Delete(userId int) error {
  return r.generic.Delete(userId)
}

// Get CRUD operation returning all entities
func (r *BusStudentRepoInit) Get() ([]*models.BusStudent, error) {
  return r.generic.Get()
}
