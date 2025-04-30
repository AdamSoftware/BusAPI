package services

import (
  "Bus-Backend/internal/models"
)



type UserService interface {
  RegeristerUser(user *models.User) (*models.User, error)
  GetUserById(userId int) (*models.User, error)
  GetUserByRole(employeeRole int) ([]*models.User, error)
  InsertUser(user *models.User) (*models.User, error)
  UpdateUser(user *models.User) (*models.User, error)
  DeleteUser(UserId int) (*models.User, error)  
  GetAllUsers() error
}





