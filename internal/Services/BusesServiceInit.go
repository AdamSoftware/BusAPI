package services

import (
	"Bus-Backend/internal/models"
	"Bus-Backend/internal/repository"
)

type BusesServiceInit struct {
	genericService GenericService[*models.Bus]
	repo           *repository.GenericRepoInit[*models.Bus]
}

func NewBusesService(repo *repository.GenericRepoInit[*models.Bus]) BusesService {
	return &BusesServiceInit{
		genericService: NewGenericServiceInit(repo),
		repo:           repo,
	}
}

func (s *BusesServiceInit) Get() ([]*models.Bus, error) {
	return s.genericService.Get()
}

// Find user by ID
func (s *BusesServiceInit) FindById(id int) (*models.Bus, error) {
	return s.genericService.FindById(id)
}

// Insert a new user
func (s *BusesServiceInit) Insert(user *models.Bus) (*models.Bus, error) {
	return s.genericService.Insert(user)
}

// Update an existing user
func (s *BusesServiceInit) Update(user *models.Bus) (*models.Bus, error) {
	return s.genericService.Update(user)
}

// Delete a user by ID
func (s *BusesServiceInit) Delete(id int) error {
	return s.genericService.Delete(id)
}




