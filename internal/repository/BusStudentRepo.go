package repository

import (
  "Bus-Backend/internal/models"
)


type BusStudentRepo interface {
  FindById(BusStudentId int) (*models.BusStudent, error)
  Update(BusStudent *models.BusStudent) (*models.BusStudent, error)
  Delete(BusStudentId int) error
  Get() ([]*models.BusStudent, error)
  Insert(BusStudent *models.BusStudent) (*models.BusStudent, error)
}
