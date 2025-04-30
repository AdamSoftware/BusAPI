package repository

import (
  "Bus-Backend/internal/models"
)


type UserRepository interface {
  FindById(UserId int) (*models.User, error)
  FindByRole(EmployeeRole int) ([]*models.User, error)
  Insert(User *models.User) (*models.User, error)
  Update(User *models.User) (*models.User, error)
  Delete(UserId int) error
  FindAll() ([]*models.User, error)
}
