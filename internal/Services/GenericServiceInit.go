package services

import (
	"Bus-Backend/internal/Logging"
	"Bus-Backend/internal/models"
	"Bus-Backend/internal/repository"
	"fmt"
)

type GenericServiceInit[T models.GenericModel] struct {
	repo *repository.GenericRepoInit[T]
}

func NewGenericServiceInit[T models.GenericModel](repo *repository.GenericRepoInit[T]) *GenericServiceInit[T] {
	return &GenericServiceInit[T]{repo: repo}
}

// FindById retrieves an entity by its ID
func (s *GenericServiceInit[T]) FindById(id int) (T, error) {
	entity, err := s.repo.FindById(id)
	if err != nil {
		var zero T
		return zero, fmt.Errorf("error finding entity with ID %d: %v", id, err)
	}
	return entity, nil
}

// Get retrieves all entities
func (s *GenericServiceInit[T]) Get() ([]T, error) {
	entities, err := s.repo.Get()
	if err != nil {
		Logging.Logs.Error(fmt.Sprintf("Error retrieving entities: %v", err))
		return nil, err
	}

	return entities, nil
}

// Insert adds a new entity
func (s *GenericServiceInit[T]) Insert(entity T) (T, error) {
	createdEntity, err := s.repo.Insert(entity)
	if err != nil {
		Logging.Logs.Error(fmt.Sprintf("Error inserting entity: %v", err))
		var zero T
		return zero, err
	}
	return createdEntity, nil
}

// Update modifies an existing entity
func (s *GenericServiceInit[T]) Update(entity T) (T, error) {
	updatedEntity, err := s.repo.Update(entity)
	if err != nil {
		Logging.Logs.Error(fmt.Sprintf("Error updating entity: %v", err))
		return updatedEntity, err
	}
	return updatedEntity, nil
}

// Delete removes an entity by its ID
func (s *GenericServiceInit[T]) Delete(id int) error {
	err := s.repo.Delete(id)
	if err != nil {
		Logging.Logs.Error(fmt.Sprintf("Error deleting entity with ID %d: %v", id, err))
		return err
	}
	return nil
}
