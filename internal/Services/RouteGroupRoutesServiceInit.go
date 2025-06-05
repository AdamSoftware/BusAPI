package services


import (
	"Bus-Backend/internal/models"
	"Bus-Backend/internal/repository"
	"fmt"
	"Bus-Backend/internal/Logging"
)

type RouteGroupRoutesServiceInit struct {
	genericService GenericService[*models.RouteGroupRoutes]
	repo           *repository.GenericRepoInit[*models.RouteGroupRoutes]
	RouteGroupRoutesRepo      *repository.RouteGroupRoutesRepoInit
}

func NewRouteGroupRoutesService(repo *repository.GenericRepoInit[*models.RouteGroupRoutes], RouteGroupRoutesRepo *repository.RouteGroupRoutesRepoInit) RouteGroupRoutesService {
	return &RouteGroupRoutesServiceInit{
		genericService: NewGenericServiceInit(repo),
		repo:           repo,
		RouteGroupRoutesRepo:      RouteGroupRoutesRepo,
	}
}

func (s *RouteGroupRoutesServiceInit) Get() ([]*models.RouteGroupRoutes, error) {
	return s.genericService.Get()
}

// Find user by ID
func (s *RouteGroupRoutesServiceInit) FindById(id int) (*models.RouteGroupRoutes, error) {
	return s.genericService.FindById(id)
}

// Insert a new user
func (s *RouteGroupRoutesServiceInit) Insert(user *models.RouteGroupRoutes) (*models.RouteGroupRoutes, error) {
	return s.genericService.Insert(user)
}

// Update an existing user
func (s *RouteGroupRoutesServiceInit) Update(user *models.RouteGroupRoutes) (*models.RouteGroupRoutes, error) {
	return s.genericService.Update(user)
}

// Delete a user by ID
func (s *RouteGroupRoutesServiceInit) Delete(id int) error {
	return s.genericService.Delete(id)
}

// Custom Methods

// GetByRouteGroupId retrieves RouteGroupRoutes by RouteGroupId
func (s *RouteGroupRoutesServiceInit) GetByRouteGroupId(routeGroupId int) ([]*models.RouteGroupRoutes, error) {
	routes, err := s.RouteGroupRoutesRepo.GetByRouteGroupId(routeGroupId)
	if err != nil {
		Logging.Logs.Warnf("Error getting RouteGroupRoutes by RouteGroupId %d: %v", routeGroupId, err)
		fmt.Printf("Error getting RouteGroupRoutes by RouteGroupId %d: %v\n", routeGroupId, err)
		return nil, err
	}
	return routes, nil
}

// GetByRouteId retrieves RouteGroupRoutes by RouteId
func (s *RouteGroupRoutesServiceInit) GetByRouteId(routeId int) ([]*models.RouteGroupRoutes, error) {
	routes, err := s.RouteGroupRoutesRepo.GetByRouteId(routeId)
	if err != nil {
		Logging.Logs.Warnf("Error getting RouteGroupRoutes by RouteId %d: %v", routeId, err)
		fmt.Printf("Error getting RouteGroupRoutes by RouteId %d: %v\n", routeId, err)
		return nil, err
	}
	return routes, nil
}

// DeleteByRouteGroupAndRoute deletes RouteGroupRoutes by RouteGroupId and RouteId
func (s *RouteGroupRoutesServiceInit) DeleteByRouteGroupAndRoute(routeGroupId int, routeId int) error {
	err := s.RouteGroupRoutesRepo.DeleteByRouteGroupAndRoute(routeGroupId, routeId)
	if err != nil {
		Logging.Logs.Warnf("Error deleting RouteGroupRoutes by RouteGroupId %d and RouteId %d: %v", routeGroupId, routeId, err)
		fmt.Printf("Error deleting RouteGroupRoutes by RouteGroupId %d and RouteId %d: %v\n", routeGroupId, routeId, err)
		return err
	}
	return nil
}
