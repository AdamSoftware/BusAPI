package services


import (
	"Bus-Backend/internal/models"
	"Bus-Backend/internal/repository"
	"fmt"
	"Bus-Backend/internal/Logging"
)

type StopsServiceInit struct {
	genericService GenericService[*models.Stops]
	repo           *repository.GenericRepoInit[*models.Stops]
	StopsRepo      *repository.StopsRepoInit
}

func NewStopsService(repo *repository.GenericRepoInit[*models.Stops], StopsRepo *repository.StopsRepoInit) StopsService {
	return &StopsServiceInit{
		genericService: NewGenericServiceInit(repo),
		repo:           repo,
		StopsRepo:      StopsRepo,
	}
}

func (s *StopsServiceInit) Get() ([]*models.Stops, error) {
	return s.genericService.Get()
}

// Find user by ID
func (s *StopsServiceInit) FindById(id int) (*models.Stops, error) {
	return s.genericService.FindById(id)
}

// Insert a new user
func (s *StopsServiceInit) Insert(user *models.Stops) (*models.Stops, error) {
	return s.genericService.Insert(user)
}

// Update an existing user
func (s *StopsServiceInit) Update(user *models.Stops) (*models.Stops, error) {
	return s.genericService.Update(user)
}

// Delete a user by ID
func (s *StopsServiceInit) Delete(id int) error {
	return s.genericService.Delete(id)
}

// DeleteIfUnused checks if a stop is not associated with any routes and deletes it if so
func (s *StopsServiceInit) DeleteIfUnused(stopId int) error {
	if stopId <= 0 {
		Logging.Logs.Warn("stopId cannot be zero")
		return fmt.Errorf("stopId cannot be nil")
	}

	// Check if the stop is associated with any routes
	isUsed, err := s.StopsRepo.IsStopUsed(stopId)
	if err != nil {
		Logging.Logs.Warn("error checking if stop is used", err)
		return fmt.Errorf("error checking if stop is used: %w", err)
	}

	if isUsed {
		Logging.Logs.Warn("stop is still in use and cannot be deleted")
		return fmt.Errorf("stop is still in use and cannot be deleted")
	}

	return s.genericService.Delete(stopId)
}





