package services

import (
	"Bus-Backend/internal/models"
	"Bus-Backend/internal/repository"
	"fmt"
	"Bus-Backend/internal/Logging"
)

type RouteServiceInit struct {
	genericService GenericService[*models.Route]
	repo           *repository.GenericRepoInit[*models.Route]
	RouteRepo      repository.RouteRepo
}

func NewRouteService(repo *repository.GenericRepoInit[*models.Route], RouteRepo *repository.RouteRepoInit) RouteService {
	return &RouteServiceInit{
		genericService: NewGenericServiceInit(repo),
		repo:           repo,
		RouteRepo:      RouteRepo,
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

// GetAllRoutesByStopId retrieves all routes associated with a specific stop ID
func (s *RouteServiceInit) GetAllRoutesByStopId(stopId int) ([]*models.Route, error) {
	// double checking if there was a stopId passed into the function 
	// will check this again during the repo and Handler layers
	if stopId == 0 {
		Logging.Logs.Warn("stopId cannot be zero")
		return nil, fmt.Errorf("stopId cannot be nil")
	}
	return s.RouteRepo.GetAllRoutesByStopId(stopId)
}




