
package repository

import (
  "Bus-Backend/internal/models"
)


type EmployeeRoleRepo interface {
  FindById(EmployeeRoleId int) (*models.EmployeeRole, error)
  Insert(EmployeeRole *models.EmployeeRole) (*models.EmployeeRole, error)
  Update(EmployeeRole *models.EmployeeRole) (*models.EmployeeRole, error)
  Delete(EmployeeRoleId int) error
  Get() ([]*models.EmployeeRole, error)
}
