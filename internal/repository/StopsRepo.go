package repository

import (
	"Bus-Backend/internal/models"
)


type StopsRepo interface {
  FindById(Id int) (*models.Stops, error)
  Insert(Stops *models.Stops) (*models.Stops, error)
  Update(Stops *models.Stops) (*models.Stops, error)
  Delete(Id int) error
  Get() ([]*models.Stops, error)
	IsStopUsed(stopId int) (bool, error)
}
