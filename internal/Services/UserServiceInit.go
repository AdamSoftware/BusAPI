package services

import (
	"golang.org/x/crypto/bcrypt"
	"Bus-Backend/internal/Logging"
	"Bus-Backend/internal/models"
	"Bus-Backend/internal/repository"
	"errors"
)

type UserServiceInit struct {
	genericService GenericService[*models.User]
	repo           *repository.GenericRepoInit[*models.User]
	Userrepo       *repository.UserRepoInit
}

func NewUserService(repo *repository.GenericRepoInit[*models.User]) UserService {
	return &UserServiceInit{
		genericService: NewGenericServiceInit(repo),
		repo:           repo,
	}
}

// Unhashed password check

func HashPassword(password string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedBytes), nil
}



// Login validates a user by username and password (you can add hash check here)
func (s *UserServiceInit) Login(username, password string) (*models.User, error) {
	users, err := s.genericService.Get()
	if err != nil {
		Logging.Logs.Errorf("Login failed: could not fetch users: %v", err)
		return nil, err
	}

	for _, user := range users {
		if user.Username == username {
			err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
			if err == nil {
				return user, nil // Password matches
			}
			break // Username matched but password did not
		}
	}

	Logging.Logs.Warnf("Login failed: user not found or incorrect password for username: %s", username)
	return nil, errors.New("invalid username or password")
}

// Get all users
func (s *UserServiceInit) Get() ([]*models.User, error) {
	return s.genericService.Get()
}

// Find user by ID
func (s *UserServiceInit) FindById(id int) (*models.User, error) {
	return s.genericService.FindById(id)
}

// Get users by role ID
func (s *UserServiceInit) FindByRole(roleId int) ([]*models.User, error) {
	users, err := s.Userrepo.FindByRole(roleId)
	if err != nil {
		Logging.Logs.Errorf("Error thrown on the service layer from repo: %v", err)
		return nil, err
	}
	return users, nil
}

// Insert a new user
func (s *UserServiceInit) Insert(user *models.User) (*models.User, error) {
	return s.genericService.Insert(user)
}

// Update an existing user
func (s *UserServiceInit) Update(user *models.User) (*models.User, error) {
	return s.genericService.Update(user)
}

// Delete a user by ID
func (s *UserServiceInit) Delete(id int) error {
	return s.genericService.Delete(id)
}


// will set up token stuff for reset password later this is local testing for the logic 
// will also setup the user getting an email or sms to reset there password will be a number thing that will let them in to reset of there phone 
// ResetPassword updates the password for a user
func (s *UserServiceInit) ResetPassword(userId int, newPassword string) error {
	// Hash the new password
	hashedPassword, err := HashPassword(newPassword)
	if err != nil {
		Logging.Logs.Errorf("Error hashing password: %v", err)
		return err
	}

	// Find the user by ID
	user, err := s.FindById(userId)
	if err != nil {
		Logging.Logs.Errorf("Error finding user by ID: %v", err)
		return err
	}

	// Update the user's password
	user.Password = hashedPassword

	// Save the updated user
	_, err = s.Update(user)
	if err != nil {
		Logging.Logs.Errorf("Error updating user password: %v", err)
		return err
	}

	return nil
}
