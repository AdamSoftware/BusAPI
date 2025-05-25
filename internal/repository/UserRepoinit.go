package repository

import (
	"Bus-Backend/internal/models"
	"errors"
	"fmt"
	"Bus-Backend/internal/Logging"
	"gorm.io/gorm"
)

// star is a pointer to the model need to use this so it doesn't break
type UserRepoInit struct {
	generic *GenericRepoInit[*models.User]
}

func NewUserRepo(db *gorm.DB) (*UserRepoInit, error) {

	if db == nil{
		return nil, fmt.Errorf("database or logger is nil")
	}

	// Create a new instance of GenericRepoInit for User model
	genericRepo, err := NewGenericRepo[*models.User](db)
	if err != nil {
		return nil, fmt.Errorf("failed to create UserRepo: %w", err)
	}

	return &UserRepoInit{generic: genericRepo}, nil
}

// using the CRUD operations from the generic repo
// FindbyId CRUD operation receiving the entity id as an argument
func (r *UserRepoInit) FindById(userId int) (*models.User, error) {
	return r.generic.FindById(userId)
}

// insert CRUD operation receiving the entity as an argument
func (r *UserRepoInit) Insert(user *models.User) (*models.User, error) {
	return r.generic.Insert(user)
}

// Update CRUD operation receiving the entity as an argument
func (r *UserRepoInit) Update(user *models.User) (*models.User, error) {
	return r.generic.Update(user)
}

// Delete CRUD operation receiving the entity id as an argument
func (r *UserRepoInit) Delete(userId int) error {
	return r.generic.Delete(userId)
}

// Get CRUD operation returning all entities
func (r *UserRepoInit) Get() ([]*models.User, error) {
	return r.generic.Get()
}

// FindByRole retrives all the users with that role
// this is not in the generic specific to the user

func (r *UserRepoInit) FindByRole(employeeRole int) ([]*models.User, error) {
	var users []*models.User
	err := r.generic.db.Where("EmployeeRoleId = ?", employeeRole).Find(&users).Error
	// check to see if there was anythign found
	if err != nil {
		// check to see if it can find that role that was requested
		if errors.Is(err, gorm.ErrRecordNotFound) {
			Logging.Logs.Warnf("no users found with role %d: %v", employeeRole, err)
			return nil, fmt.Errorf("no users found with role %d: %w", employeeRole, err)
		}
		Logging.Logs.Warnf("error retrieving users with role %d: %v", employeeRole, err)
		// error if it cound't find anything about that role or that person
		return nil, fmt.Errorf("error retrieving users with role %d: %w", employeeRole, err)
	}

	return users, nil
}
