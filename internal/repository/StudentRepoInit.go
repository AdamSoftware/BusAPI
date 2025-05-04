package repository

import (
	"Bus-Backend/internal/models"
	"fmt"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// star is a pointer to the model need to use this so it doesn't break
type StudentRepoInit struct {
	generic *GenericRepoInit[*models.Student]
}

func NewStudentRepo(db *gorm.DB, logger *logrus.Logger) (*StudentRepoInit, error) {
	// Create a new instance of GenericRepoInit for Student model
	genericRepo, err := NewGenericRepo[*models.Student](db, logger)
	if err != nil {
		return nil, fmt.Errorf("failed to create StudentRepo: %w", err)
	}

	return &StudentRepoInit{generic: genericRepo}, nil
}

// using the CRUD operations from the generic repo
// FindbyId CRUD operation receiving the entity id as an argument
func (r *StudentRepoInit) FindById(StudentId int) (*models.Student, error) {
	return r.generic.FindById(StudentId)
}

// insert CRUD operation receiving the entity as an argument
func (r *StudentRepoInit) Insert(Student *models.Student) (*models.Student, error) {
	return r.generic.Insert(Student)
}

// Update CRUD operation receiving the entity as an argument
func (r *StudentRepoInit) Update(Student *models.Student) (*models.Student, error) {
	return r.generic.Update(Student)
}

// Delete CRUD operation receiving the entity id as an argument
func (r *StudentRepoInit) Delete(StudentId int) error {
	return r.generic.Delete(StudentId)
}

// Get CRUD operation returning all entities
func (r *StudentRepoInit) Get() ([]*models.Student, error) {
	return r.generic.Get()
}

// FindBySchoolId retrieves all the students associated with a specific school ID
func (r *StudentRepoInit) FindBySchoolId(schoolId int) ([]*models.Student, error) {
	if schoolId <= 0 {
		return nil, fmt.Errorf("invalid school ID")
	}

	var students []*models.Student
	result := r.generic.db.Where("school_id = ?", schoolId).Find(&students)

	if result.Error != nil {
		r.generic.log.WithError(result.Error).Error("failed to find students by school ID")
		return nil, fmt.Errorf("failed to find students by school ID: %w", result.Error)
	}

	return students, nil
}
