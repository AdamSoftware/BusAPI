package services

import "Bus-Backend/internal/models"

type EmployeeRoleService interface {
	Get() ([]*models.EmployeeRole, error)
	FindById(id int) (*models.EmployeeRole, error)
	Insert(user *models.EmployeeRole) (*models.EmployeeRole, error)
	Update(user *models.EmployeeRole) (*models.EmployeeRole, error)
	Delete(id int) error
}
