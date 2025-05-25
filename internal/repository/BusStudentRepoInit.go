package repository


import (
  "Bus-Backend/internal/models"
  "gorm.io/gorm"
  "fmt"
)


type BusStudentRepoInit struct {
  generic *GenericRepoInit[*models.BusStudent]
}


func NewBusStudentRepo(db *gorm.DB) (*BusStudentRepoInit, error) {
  genericRepo, err := NewGenericRepo[*models.BusStudent](db)
  if err != nil {
    return nil, fmt.Errorf("failed to create BusStudentRepo: %w", err)
  }

  return &BusStudentRepoInit{generic: genericRepo}, nil
}


// using the CRUD operations from the generic repo 
// FindbyId CRUD operation receiving the entity id as an argument
func (r *BusStudentRepoInit) FindById(BusStudentId int) (*models.BusStudent, error) {
  return r.generic.FindById(BusStudentId)
}

// insert CRUD operation receiving the entity as an argument
func (r *BusStudentRepoInit) Insert(BusStudent *models.BusStudent) (*models.BusStudent, error) {
  return r.generic.Insert(BusStudent)
}

// Update CRUD operation receiving the entity as an argument
func (r *BusStudentRepoInit) Update(BusStudent *models.BusStudent) (*models.BusStudent, error) {
  return r.generic.Update(BusStudent)
}

// Delete CRUD operation receiving the entity id as an argument
func (r *BusStudentRepoInit) Delete(BusStudentId int) error {
  return r.generic.Delete(BusStudentId)
}

// Get CRUD operation returning all entities
func (r *BusStudentRepoInit) Get() ([]*models.BusStudent, error) {
  return r.generic.Get()
}
