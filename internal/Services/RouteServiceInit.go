package services

import (
	"Bus-Backend/internal/models"
	"Bus-Backend/internal/repository"
)

type RouteServiceInit struct {
	genericService GenericService[*models.Route]
	repo           *repository.GenericRepoInit[*models.Route]
}

func NewRouteService(repo *repository.GenericRepoInit[*models.Route]) RouteService {
	return &RouteServiceInit{
		genericService: NewGenericServiceInit(repo),
		repo:           repo,
	}
}

func (s *RouteServiceInit) Get() ([]*models.Route, error) {
	return s.genericService.Get()
}

// Find user by ID
func (s *RouteServiceInit) FindById(id int) (*models.Route, error) {
	return s.genericService.FindById(id)
}

// Insert a new user
func (s *RouteServiceInit) Insert(user *models.Route) (*models.Route, error) {
	return s.genericService.Insert(user)
}

// Update an existing user
func (s *RouteServiceInit) Update(user *models.Route) (*models.Route, error) {
	return s.genericService.Update(user)
}

// Delete a user by ID
func (s *RouteServiceInit) Delete(id int) error {
	return s.genericService.Delete(id)
}




