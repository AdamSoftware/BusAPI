package services

import (
	"Bus-Backend/internal/models"
	"Bus-Backend/internal/repository"
)

type EmployeeRoleServiceInit struct {
	genericService GenericService[*models.EmployeeRole]
	repo           *repository.GenericRepoInit[*models.EmployeeRole]
}

func NewEmployeeRoleService(repo *repository.GenericRepoInit[*models.EmployeeRole]) EmployeeRoleService {
	return &EmployeeRoleServiceInit{
		genericService: NewGenericServiceInit(repo),
		repo:           repo,
	}
}

func (s *EmployeeRoleServiceInit) Get() ([]*models.EmployeeRole, error) {
	return s.genericService.Get()
}

// Find user by ID
func (s *EmployeeRoleServiceInit) FindById(id int) (*models.EmployeeRole, error) {
	return s.genericService.FindById(id)
}

// Insert a new user
func (s *EmployeeRoleServiceInit) Insert(user *models.EmployeeRole) (*models.EmployeeRole, error) {
	return s.genericService.Insert(user)
}

// Update an existing user
func (s *EmployeeRoleServiceInit) Update(user *models.EmployeeRole) (*models.EmployeeRole, error) {
	return s.genericService.Update(user)
}

// Delete a user by ID
func (s *EmployeeRoleServiceInit) Delete(id int) error {
	return s.genericService.Delete(id)
}




