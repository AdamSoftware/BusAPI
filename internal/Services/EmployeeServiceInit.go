package services

import (
	"Bus-Backend/internal/models"
	"Bus-Backend/internal/repository"
)

type EmployeeServiceInit struct {
	genericService GenericService[*models.Employee]
	repo           *repository.GenericRepoInit[*models.Employee]
}

func NewEmployeeService(repo *repository.GenericRepoInit[*models.Employee]) EmployeeService {
	return &EmployeeServiceInit{
		genericService: NewGenericServiceInit(repo),
		repo:           repo,
	}
}

func (s *EmployeeServiceInit) Get() ([]*models.Employee, error) {
	return s.genericService.Get()
}

// Find user by ID
func (s *EmployeeServiceInit) FindById(id int) (*models.Employee, error) {
	return s.genericService.FindById(id)
}

// Insert a new user
func (s *EmployeeServiceInit) Insert(user *models.Employee) (*models.Employee, error) {
	return s.genericService.Insert(user)
}

// Update an existing user
func (s *EmployeeServiceInit) Update(user *models.Employee) (*models.Employee, error) {
	return s.genericService.Update(user)
}

// Delete a user by ID
func (s *EmployeeServiceInit) Delete(id int) error {
	return s.genericService.Delete(id)
}




