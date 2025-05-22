package repository

// repository is like the data access layer in c# like entityFramework

import (
	"Bus-Backend/internal/models"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"Bus-Backend/internal/Logging"
	"gorm.io/gorm"
)

// this is for the course options of the different repos
type GenericRepoInit[T models.GenericModel] struct {
	db  *gorm.DB
	log *logrus.Logger
}

func NewGenericRepo[T models.GenericModel](db *gorm.DB, logger *logrus.Logger) (*GenericRepoInit[T], error) {

	// if the database was never found then return an error
	if db == nil {
		err := errors.New("The database was never initialized")
		return nil, fmt.Errorf("Failed to create GenericRepo: %w", err)
	}

	// if the logger was never initialized then create a new one
	if logger == nil {
		return nil, errors.New("Logger was never initialized")
	}

	logger.SetFormatter(&logrus.TextFormatter{FullTimestamp: true})
	logger.SetLevel(logrus.InfoLevel)

	return &GenericRepoInit[T]{db: db, log: logger}, nil
}

// FindById returns a single entity by id
func (r *GenericRepoInit[T]) FindById(id int) (T, error) {
	var entity T
	err := r.db.First(&entity, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		Logging.Logs.Warnf("Entity with id %d not found", id)
		var zero T
		return zero, fmt.Errorf("Entity with id %d not found: %w", id, err)
	}

	return entity, nil
}

// Get returns all entities
func (r *GenericRepoInit[T]) Get() ([]T, error) {
	var entities []T
	err := r.db.Find(&entities).Error
	if err != nil {
		Logging.Logs.Errorf("Error retrieving entities: %v", err)
		return nil, fmt.Errorf("Error retrieving entities: %v", err)
	}

	return entities, nil
}

// Insert adds a new entity to the database

func (r *GenericRepoInit[T]) Insert(entity T) (T, error) {
    err := r.db.Create(entity).Error // Don't pass a pointer because entity is already a pointer 
    if err != nil {
				Logging.Logs.Errorf("Error inserting entity: %v", err)
        var zero T
        return zero, fmt.Errorf("Error inserting entity: %v", err)
    }

    return entity, nil
}



// Update modifies an existing entity in the database
func (r *GenericRepoInit[T]) Update(entity T) (T, error) {
    Logging.Logs.Infof("Got to the Update: %+v", entity)

    // Log the state before saving
    Logging.Logs.Infof("Before save: %+v", entity)

    // Perform the update
    if err := r.db.Save(entity).Error; err != nil {
        var zero T
				Logging.Logs.Errorf("Error updating entity: %v", err)
        return zero, err
    }

    // Log the state after saving
    Logging.Logs.Infof("After save: %+v", entity)

    // Return the updated entity
    return entity, nil
}

// Delete removes an entity from the database by id
func (r *GenericRepoInit[T]) Delete(id int) error {
	var entity T
	err := r.db.First(&entity, id).Error
	// check if the entity was found if not throw error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		Logging.Logs.Warnf("Entity with id %d not found for deletion", id)
		return fmt.Errorf("Entity with id %d not found: %w", id, err)
	} else if err != nil {
		// check to see if there was any other error found
		Logging.Logs.Errorf("Something happened: %v", err)
		return fmt.Errorf("Something happened: %v", err)
	}

	// if the database never delete the entity return error
	err = r.db.Delete(&entity).Error
	if err != nil {
		Logging.Logs.Errorf("Error deleting entity: %v", err)
		return fmt.Errorf("Error deleting entity: %v", err)
	}

	return nil
}
