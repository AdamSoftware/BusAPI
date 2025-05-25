package services

import (
	"Bus-Backend/internal/models"
	"Bus-Backend/internal/repository"
)

type StudentServiceInit struct {
	genericService GenericService[*models.Student]
	repo           *repository.GenericRepoInit[*models.Student]
	Studentrepo 	 *repository.StudentRepoInit
}

func NewStudentService(repo *repository.GenericRepoInit[*models.Student]) StudentService {
	return &StudentServiceInit{
		genericService: NewGenericServiceInit(repo),
		repo:           repo,
	}
}

func (s *StudentServiceInit) Get() ([]*models.Student, error) {
	return s.genericService.Get()
}

// Find user by School ID
func (s *StudentServiceInit) FindBySchoolId(schoolId int) ([]*models.Student, error) {
	students, err := s.Studentrepo.FindBySchoolId(schoolId) 
	if err != nil {
		return nil, err
	}
	return students, nil
}

// Find user by ID
func (s *StudentServiceInit) FindById(id int) (*models.Student, error) {
	return s.genericService.FindById(id)
}

// Insert a new user
func (s *StudentServiceInit) Insert(user *models.Student) (*models.Student, error) {
	return s.genericService.Insert(user)
}

// Update an existing user
func (s *StudentServiceInit) Update(user *models.Student) (*models.Student, error) {
	return s.genericService.Update(user)
}

// Delete a user by ID
func (s *StudentServiceInit) Delete(id int) error {
	return s.genericService.Delete(id)
}




