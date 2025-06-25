package repository

import (
	"Bus-Backend/internal/models"
)


type RouteGroupRepo interface {
  FindById(EmployeeId int) (*models.Employee, error)
  Insert(Employee *models.Employee) (*models.Employee, error)
  Update(Employee *models.Employee) (*models.Employee, error)
  Delete(EmployeeId int) error
  Get() ([]*models.Employee, error)
}
