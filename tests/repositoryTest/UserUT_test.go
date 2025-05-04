
package repositorytest

import (
	"Bus-Backend/internal/models"
	"Bus-Backend/internal/repository"
	"testing"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB
var logger *logrus.Logger

func GetUserId(u models.User) int {
	return u.UserId
}

func initTestDependencies() {
	logger = logrus.New()
	logger.SetLevel(logrus.DebugLevel)

	var err error
	db, err = gorm.Open(sqlite.Open("../../Database/bus_database.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to initialize database: " + err.Error())
	}
}

func TestUserUpdate(t *testing.T) {
	initTestDependencies()

	repo, err := repository.NewUserRepo(db, logger)
	if err != nil {
		t.Fatalf("Failed to create UserRepo: %v", err)
	}

	userID := 1
	entity, err := repo.FindById(userID)
	if err != nil {
		t.Fatalf("Failed to find user by ID: %v", err)
	}
	if entity == nil {
		t.Fatalf("User entity should not be nil")
	}

	t.Logf("Before update: %v", entity)

	updateFunc := func(u *models.User) {
		u.Username = "updateduser"
	}

	updateFunc(entity)

	_, err = repo.Update(entity)
	if err != nil {
		t.Fatalf("Failed to update user: %v", err)
	}

	updatedUser, err := repo.FindById(userID)
	if err != nil {
		t.Fatalf("Failed to find updated user by ID: %v", err)
	}

	if updatedUser.Username != "updateduser" {
		t.Errorf("Expected updated username to be 'updateduser', but got '%s'", updatedUser.Username)
	}

	if updatedUser.UserId != entity.UserId {
		t.Errorf("Expected UserId to remain the same: %d, but got: %d", entity.UserId, updatedUser.UserId)
	}
	if updatedUser.Password != entity.Password {
		t.Errorf("Expected Password to remain the same: %s, but got: %s", entity.Password, updatedUser.Password)
	}
}

