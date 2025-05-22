package services

type Entity interface{GetId() int}


type GenericService[T any] interface {
  FindById(id int) (*T, error)
  Get() ([]*T, error)
  Insert(entity *T) (*T, error)
  Update(entity *T) (*T, error)
  Delete(id int) error
}
