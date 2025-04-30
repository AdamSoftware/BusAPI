package services

import (
)


type GenericRepo[T any] interface {
  FindById(id int) (*T, error)
  FindAll() ([]*T, error)
  Insert(entity *T) (*T, error)
  Update(entity *T) (*T, error)
  Delete(id int) error
}
