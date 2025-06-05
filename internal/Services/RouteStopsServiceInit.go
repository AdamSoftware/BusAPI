package services 

import (
	"Bus-Backend/internal/models"
	"Bus-Backend/internal/repository"
	"fmt"
	"Bus-Backend/internal/Logging"
	"Bus-Backend/internal/dto"
)

type RouteStopsServiceInit struct {
	genericService GenericService[*models.RouteStops]
	repo           *repository.GenericRepoInit[*models.RouteStops]
	RouteStopsRepo *repository.RouteStopsRepoInit
	StopsService   StopsService
}

func NewRouteStopsService(
	repo *repository.GenericRepoInit[*models.RouteStops],
	RouteStopsRepo *repository.RouteStopsRepoInit,
	StopsService StopsService) RouteStopsService {
	return &RouteStopsServiceInit{
		genericService: NewGenericServiceInit(repo),
		repo:           repo,
		RouteStopsRepo: RouteStopsRepo,
		StopsService:   StopsService,
	}
}

func (s *RouteStopsServiceInit) Get() ([]*models.RouteStops, error) {
	return s.genericService.Get()
}

// Find user by ID
func (s *RouteStopsServiceInit) FindById(id int) (*models.RouteStops, error) {
	return s.genericService.FindById(id)
}

// Insert a new user
func (s *RouteStopsServiceInit) Insert(user *models.RouteStops) (*models.RouteStops, error) {
	return s.genericService.Insert(user)
}

// Update an existing user
func (s *RouteStopsServiceInit) Update(user *models.RouteStops) (*models.RouteStops, error) {
	return s.genericService.Update(user)
}

// Delete a user by ID
func (s *RouteStopsServiceInit) Delete(id int) error {
	return s.genericService.Delete(id)
}


// Custom Methods

// FindByRouteAndStopIds retrieves a RouteStops by routeId and stopId
func (s *RouteStopsServiceInit) FindByRouteAndStop(routeId, stopId int) ([]*models.RouteStops, error) {
	if routeId <= 0 || stopId <= 0 {
		Logging.Logs.Warn("routeId and stopId cannot be zero")
		return nil, fmt.Errorf("routeId and stopId cannot be nil")
	}
	return s.RouteStopsRepo.FindByRouteAndStop(routeId, stopId)
}

// DeleteByRouteAndStopIds deletes a RouteStops by routeId and stopId

func (s *RouteStopsServiceInit) DeleteByRouteAndStop(routeId, stopId int) error {
	if routeId <= 0 || stopId <= 0 {
		Logging.Logs.Warn("routeId and stopId cannot be zero")
		return fmt.Errorf("routeId and stopId cannot be nil")
	}

	err := s.RouteStopsRepo.DeleteByRouteAndStop(routeId, stopId)
	if err != nil {
		Logging.Logs.Error("error deleting RouteStops by routeId and stopId: ", err)
		return fmt.Errorf("error deleting RouteStops by routeId and stopId: %w", err)
	}

	if s.StopsService != nil {
		if err := s.StopsService.DeleteIfUnused(stopId); err != nil {
			Logging.Logs.Error("error checking if stop is used: ", err)
			return fmt.Errorf("error checking if stop is used: %w", err)
		}
	}
	return nil
}


// GetByRouteId retrieves all RouteStops associated with a specific route ID
func (s *RouteStopsServiceInit) GetByRouteId(routeId int) ([]*models.RouteStops, error) {
	if routeId <= 0 {
		Logging.Logs.Warn("routeId cannot be zero")
		return nil, fmt.Errorf("routeId cannot be nil")
	}
	return s.RouteStopsRepo.GetByRouteId(routeId)
}

// UpdateStopOrder updates the order of stops in a route
func (s *RouteStopsServiceInit) UpdateStopOrder(routeId, stopId, newOrder int) error {
	if routeId <= 0 || stopId <= 0 {
		Logging.Logs.Warn("routeId and stopId cannot be zero")
		return fmt.Errorf("routeId and stopId cannot be nil")
	}
	return s.RouteStopsRepo.UpdateStopOrder(routeId, stopId, newOrder)
}

// GetOrderedStopsWithDetails retrieves ordered stops with details about the routes
func (s *RouteStopsServiceInit) GetOrderedStopsWithDetails(routeId int) ([]dto.OrderStopsInfo, error) {
	if routeId <= 0 {
		Logging.Logs.Warn("routeId cannot be zero")
		return nil, fmt.Errorf("routeId cannot be nil")
	}
	return s.RouteStopsRepo.GetOrderedStopsWithDetails(routeId)
}
