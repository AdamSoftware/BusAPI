package services

import (
	"Bus-Backend/internal/models"
	"Bus-Backend/internal/repository"
)

type SchoolServiceInit struct {
	genericService GenericService[*models.School]
	repo           *repository.GenericRepoInit[*models.School]
}

func NewSchoolService(repo *repository.GenericRepoInit[*models.School]) SchoolService {
	return &SchoolServiceInit{
		genericService: NewGenericServiceInit(repo),
		repo:           repo,
	}
}

func (s *SchoolServiceInit) Get() ([]*models.School, error) {
	return s.genericService.Get()
}

// Find user by ID
func (s *SchoolServiceInit) FindById(id int) (*models.School, error) {
	return s.genericService.FindById(id)
}

// Insert a new user
func (s *SchoolServiceInit) Insert(user *models.School) (*models.School, error) {
	return s.genericService.Insert(user)
}

// Update an existing user
func (s *SchoolServiceInit) Update(user *models.School) (*models.School, error) {
	return s.genericService.Update(user)
}

// Delete a user by ID
func (s *SchoolServiceInit) Delete(id int) error {
	return s.genericService.Delete(id)
}




