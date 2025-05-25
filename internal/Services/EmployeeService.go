package services

import "Bus-Backend/internal/models"

type EmployeeService interface {
	Get() ([]*models.Employee, error)
	FindById(id int) (*models.Employee, error)
	Insert(user *models.Employee) (*models.Employee, error)
	Update(user *models.Employee) (*models.Employee, error)
	Delete(id int) error
}
