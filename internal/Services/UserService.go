package services

import "Bus-Backend/internal/models"

type UserService interface {
	Get() ([]*models.User, error)
	FindById(id int) (*models.User, error)
	FindByRole(roleId int) ([]*models.User, error)
	Insert(user *models.User) (*models.User, error)
	Update(user *models.User) (*models.User, error)
	Delete(id int) error
	Login(username, password string) (*models.User, error)
}
