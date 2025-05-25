package services

import "Bus-Backend/internal/models"

type BusStudentService interface {
	Get() ([]*models.BusStudent, error)
	FindById(id int) (*models.BusStudent, error)
	Insert(user *models.BusStudent) (*models.BusStudent, error)
	Update(user *models.BusStudent) (*models.BusStudent, error)
	Delete(id int) error
}
