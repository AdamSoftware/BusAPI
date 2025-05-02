package repository

import (
  "Bus-Backend/internal/models"
)


type StudentRepo interface { 
  FindById(StudentId int) (*models.Student, error)
  FindBySchoolId(SchoolId int) ([]*models.Student, error)
  Insert(Student *models.Student) (*models.Student, error)
  Update(Student *models.Student) (*models.Student, error)
  Delete(StudentId int) error
  Get() ([]*models.Student, error)
}




