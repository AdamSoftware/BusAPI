package services

import (
	"Bus-Backend/internal/models"
	"Bus-Backend/internal/repository"
)

type BusStudentServiceInit struct {
	genericService GenericService[*models.BusStudent]
	repo           *repository.GenericRepoInit[*models.BusStudent]
}

func NewBusStudentService(repo *repository.GenericRepoInit[*models.BusStudent]) BusStudentService {
	return &BusStudentServiceInit{
		genericService: NewGenericServiceInit(repo),
		repo:           repo,
	}
}

func (s *BusStudentServiceInit) Get() ([]*models.BusStudent, error) {
	return s.genericService.Get()
}

// Find user by ID
func (s *BusStudentServiceInit) FindById(id int) (*models.BusStudent, error) {
	return s.genericService.FindById(id)
}

// Insert a new user
func (s *BusStudentServiceInit) Insert(user *models.BusStudent) (*models.BusStudent, error) {
	return s.genericService.Insert(user)
}

// Update an existing user
func (s *BusStudentServiceInit) Update(user *models.BusStudent) (*models.BusStudent, error) {
	return s.genericService.Update(user)
}

// Delete a user by ID
func (s *BusStudentServiceInit) Delete(id int) error {
	return s.genericService.Delete(id)
}




