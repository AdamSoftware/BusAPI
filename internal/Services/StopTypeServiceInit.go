package services


import (
	"Bus-Backend/internal/models"
	"Bus-Backend/internal/repository"
	"fmt"
	"Bus-Backend/internal/Logging"
)

type StopTypeServiceInit struct {
	genericService GenericService[*models.StopType]
	repo           *repository.GenericRepoInit[*models.StopType]
	StopTypeRepo   *repository.StopTypeRepoInit
}

func NewStopTypeService(
	repo *repository.GenericRepoInit[*models.StopType],
	StopTypeRepo *repository.StopTypeRepoInit,
	)StopTypeService {
		return &StopTypeServiceInit{
		genericService: NewGenericServiceInit(repo),
		repo:           repo,
		StopTypeRepo: StopTypeRepo,
	}
}

func (s *StopTypeServiceInit) Get() ([]*models.StopType, error) {
	return s.genericService.Get()
}

// Find user by ID
func (s *StopTypeServiceInit) FindById(id int) (*models.StopType, error) {
	return s.genericService.FindById(id)
}

// Insert a new user
func (s *StopTypeServiceInit) Insert(user *models.StopType) (*models.StopType, error) {
	return s.genericService.Insert(user)
}

// Update an existing user
func (s *StopTypeServiceInit) Update(user *models.StopType) (*models.StopType, error) {
	return s.genericService.Update(user)
}

// Delete a user by ID
func (s *StopTypeServiceInit) Delete(id int) error {
	return s.genericService.Delete(id)
}

// GetByTypeName retrieves a StopType by its TypeName

func (s *StopTypeServiceInit) GetByTypeName(TypeName string) (*models.StopType, error) {
	if TypeName == "" {
		Logging.Logs.Warn("TypeName cannot be empty")
		return nil, fmt.Errorf("TypeName cannot be empty")
	}

	stopType, err := s.StopTypeRepo.GetByTypeName(TypeName)
	if err != nil {
		Logging.Logs.Warn(err, "Error retrieving StopType by TypeName")
		return nil, err
	}

	if stopType == nil {
		Logging.Logs.Warnf("StopType with TypeName '%s' not found", TypeName)
		return nil, fmt.Errorf("StopType with TypeName '%s' not found", TypeName)
	}

	return stopType, nil
}
