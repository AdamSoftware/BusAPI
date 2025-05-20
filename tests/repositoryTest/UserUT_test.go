
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

	logger = logrus.New()
	logger.SetLevel(logrus.DebugLevel)

	var Err error
	db, Err = gorm.Open(sqlite.Open("../../Database/bus_database.db"), &gorm.Config{})

	if Err != nil {
		panic("Failed to initialize database: " + Err.Error())
	}

	repo, err := repository.NewUserRepo(db, logger)
	if err != nil {
		t.Fatalf("Failed to create UserRepo: %v", err)
	}

	// Start a new transaction to isolate changes
	tx := db.Begin()
	if tx.Error != nil {
		t.Fatalf("Failed to start transaction: %v", tx.Error)
	}

	// Make sure to rollback the transaction at the end of the test to undo any changes
	defer func() {
		if err := tx.Rollback().Error; err != nil {
			t.Fatalf("Failed to rollback transaction: %v", err)
		}
	}()

	// Find the existing user by ID (we're assuming a user with ID 1 exists in the database)
	UserID := 1
	entity, err := repo.FindById(UserID)
	if err != nil {
		t.Fatalf("Failed to find user by ID: %v", err)
	}

	t.Logf("Before update: %v", entity)

	// Update the entity's username
	updateFunc := func(u *models.User) {
		u.Username = "updateduser"
	}
	updateFunc(entity)

	// Perform the update within the transaction
	_, err = repo.Update(entity)
	if err != nil {
		t.Fatalf("Failed to update user: %v", err)
	}

	// Fetch the updated user from the database
	updatedUser, err := repo.FindById(UserID)
	if err != nil {
		t.Fatalf("Failed to find updated user by ID: %v", err)
	}

	// Verify that the username was updated correctly
	if updatedUser.Username != "updateduser" {
		t.Errorf("Expected updated username to be 'updateduser', but got '%s'", updatedUser.Username)
	}

	// Verify that other fields (like UserId and Password) remain the same
	if updatedUser.UserId != entity.UserId {
		t.Errorf("Expected UserId to remain the same: %d, but got: %d", entity.UserId, updatedUser.UserId)
	}
	if updatedUser.Password != entity.Password {
		t.Errorf("Expected Password to remain the same: %s, but got: %s", entity.Password, updatedUser.Password)
	}

	// At the end of the test, the changes are rolled back, keeping the database in its original state
}






func TestUserInsert(t *testing.T) {
	initTestDependencies()

	repo, err := repository.NewUserRepo(db, logger)
	if err != nil {
		t.Fatalf("Failed to create UserRepo: %v", err)
	}

	// Start a new transaction to isolate changes
	tx := db.Begin()
	if tx.Error != nil {
		t.Fatalf("Failed to start transaction: %v", tx.Error)
	}

	// Make sure to rollback the transaction at the end of the test to undo any changes
	defer func() {
		if err := tx.Rollback().Error; err != nil {
			t.Fatalf("Failed to rollback transaction: %v", err)
		}
	}()

	// Creating a new user entity to insert
	newUser := &models.User{
		Username:    "newuser",
		Password:    "newpassword",
		EmployeeID:  1,
		EmployeeRoles: 2,
	}

	// Insert the user
	insertedUser, err := repo.Insert(newUser)
	if err != nil {
		t.Fatalf("Failed to insert user: %v", err)
	}

	// Fetch the user by ID from the transaction
	fetchedUser, err := repo.FindById(insertedUser.UserId)
	if err != nil {
		t.Fatalf("Failed to find user by ID: %v", err)
	}

	// Assert that the user is correctly inserted
	if fetchedUser.Username != newUser.Username {
		t.Errorf("Expected username to be 'newuser', but got '%s'", fetchedUser.Username)
	}
	if fetchedUser.Password != newUser.Password {
		t.Errorf("Expected password to be 'newpassword', but got '%s'", fetchedUser.Password)
	}

	// At the end of the test, we ensure to rollback the changes, thus keeping the database in a clean state
}



// test for deleteing from the repo

func TestUserDelete(t *testing.T) {
    initTestDependencies()

    repo, err := repository.NewUserRepo(db, logger)
    if err != nil {
        t.Fatalf("Failed to create UserRepo: %v", err)
    }

    // Start a new transaction to isolate changes
    tx := db.Begin()
    if tx.Error != nil {
        t.Fatalf("Failed to start transaction: %v", tx.Error)
    }

    // Make sure to rollback the transaction at the end of the test to undo any changes
    defer func() {
        if err := tx.Rollback().Error; err != nil {
            t.Fatalf("Failed to rollback transaction: %v", err)
        }
    }()

    // Find an existing user to delete (we're assuming a user with ID 2 exists from a previous Insert test)
    
		users, err := repo.Get()
    if err != nil {
        t.Fatalf("Failed to get users: %v", err)
    }
    if len(users) == 0 {
        t.Fatal("No users found to delete")
    }

    var lastUser models.User
    for _, u := range users {
        if u.UserId > lastUser.UserId {
            lastUser = *u 
        }
    }

    t.Logf("Deleting user with ID: %d", lastUser.UserId)

    // Perform delete within the transaction
    err = repo.Delete(lastUser.UserId)
    if err != nil {
        t.Fatalf("Failed to delete user: %v", err)
    }

	
    // Try to find the deleted user by ID, should return an error
    _, err = repo.FindById(lastUser.UserId)
    if err == nil {
        t.Errorf("Expected error when finding deleted user, but no error occurred")
    }

    // Assert the deletion has occurred correctly
    remainingUsers, err := repo.Get()
    if err != nil {
        t.Fatalf("Failed to get users after deletion: %v", err)
    }

    // We expect one fewer user now
    if len(remainingUsers) != len(users)-1 {
        t.Errorf("Expected %d users after deletion, but got %d", len(users)-1, len(remainingUsers))
    }
}

// testing the findbyid 

func TestFindById(t *testing.T) {
	initTestDependencies()

	repo, err := repository.NewUserRepo(db, logger)
	if err != nil {
		t.Fatalf("Failed to create UserRepo: %v", err)
	}

	UserID := 1 
	user, err := repo.FindById(UserID)
	if err != nil {
		t.Fatalf("Failed to find user by ID: %v", err)
	}

	// Assert that the fetched user's ID matches the expected ID
	if user.UserId != UserID {
		t.Errorf("Expected UserID to be %d, but got %d", UserID, user.UserId)
	}
}


// testing the get method

func TestFindByRole(t *testing.T) {
	initTestDependencies()

	repo, err := repository.NewUserRepo(db, logger)
	if err != nil {
		t.Fatalf("Failed to create UserRepo: %v", err)
	}

	role := 1 
	users, err := repo.FindByRole(role)
	if err != nil {
		t.Fatalf("Failed to find users by role: %v", err)
	}

	// Assert that the returned users have the expected role
	for _, user := range users {
		if user.EmployeeRoles != role {
			t.Errorf("Expected role to be %d, but got %d", role, user.EmployeeRoles)
		}
	}
}


func TestGetAll(t *testing.T) {
    initTestDependencies()

    repo, err := repository.NewUserRepo(db, logger)
    if err != nil {
        t.Fatalf("Failed to create UserRepo: %v", err)
    }

    // Get all users from the repository using the Get method
    users, err := repo.Get()
    if err != nil {
        t.Fatalf("Failed to get users: %v", err)
    }

    expectedUserCount := 3 // Change this to the actual number of users you expect in the database

    // Ensure the number of users matches the expected count
    if len(users) != expectedUserCount {
        t.Errorf("Expected %d users, but got %d", expectedUserCount, len(users))
    }

}

