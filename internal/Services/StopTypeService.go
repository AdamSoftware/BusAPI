package services

import (
	"Bus-Backend/internal/models"
)


type StopTypeService interface {
	FindById(id int) (*models.StopType, error)
	Insert(stop *models.StopType) (*models.StopType, error)
	Update(stop *models.StopType) (*models.StopType, error)
	Delete(id int) error

	GetByTypeName(TypeName string) (*models.StopType, error)
}
