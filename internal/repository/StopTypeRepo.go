package repository

import (
	"Bus-Backend/internal/models"
)

type StopTypeRepo interface {
	FindById(Id int) (*models.StopType, error)
	Insert(stopType *models.StopType) (*models.StopType, error)
	Update(stopType *models.StopType) (*models.StopType, error)
	Delete(stopTypeId int) error
	Get() ([]*models.StopType, error)

	GetByTypeName(TypeName string) (*models.StopType, error)
}
