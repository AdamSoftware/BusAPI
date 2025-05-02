package repository

import (
  "Bus-Backend/internal/models"
)


type BusRepo interface {
  FindById(BusId int) (*models.Bus, error)
  Insert(Bus *models.Bus) (*models.Bus, error)
  Update(Bus *models.Bus) (*models.Bus, error)
  Delete(BusId int) error
  Get() ([]*models.Bus, error)
}

