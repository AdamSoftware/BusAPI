package services

import "Bus-Backend/internal/models"

type StudentService interface {
	Get() ([]*models.Student, error)
	FindById(id int) (*models.Student, error)
	FindBySchoolId(schoolId int) ([]*models.Student, error)
	Insert(user *models.Student) (*models.Student, error)
	Update(user *models.Student) (*models.Student, error)
	Delete(id int) error
}
