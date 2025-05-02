package repository

import (
  "Bus-Backend/internal/models"
)


type SchoolRepo interface {
  FindById(SchoolId int) (*models.School, error)
  Insert(School *models.School) (*models.School, error)
  Update(School *models.School) (*models.School, error)
  Delete(SchoolId int) error
  Get() ([]*models.School, error)
}
