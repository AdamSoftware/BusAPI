package services



import (
  "Errors"
  "Bus-Backend/internal/models"
  "Bus-Backend/internal/repository"
)


type userService struct {
  repo repository.UserRepository
}


func NewUserService(repo repository.UserRepository) *userService {
  return &userService{repo: repo,}
}


func (s *userService) RegeristerUser(user *models.User) (*models.User, error) {

  existingUser, err := s.repo.FindById(user.UserId)

  // check to see if the user already exists
  if err != nil && existingUser != nil {
    return nil, errors.New("User is already herek") 
  }
  return s.repo.Insert(user) 

}


func (s *userService) GetUserById(userId int) (*models.User, error) {
  user, err := s.repo.FindById(userId)
  if err != nil {
    return nil, errors.New("User not found")
  }
  return user, nil
}
