package service_test

import (
	"log"
	"testing"

	"veo/internal/configs"
	"veo/internal/database"
	"veo/internal/repository"
	"veo/internal/service"

	"github.com/stretchr/testify/assert"
)

// Initializes the test database and returns a UserService instance.
func setupTestUserService(t *testing.T) service.UserService {
	// Show current path for debugging
	// cwd, _ := os.Getwd()
	// fmt.Println("Current working directory:", cwd)
	// Expected path: ./internal/repository

	// Load configuration file
	// Configuration path: ../../config/config.yaml
	cfg, err := configs.Load("../../../config/config.yaml")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Initialize the database connection
	if err := database.Init(cfg.Database); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// Initialize the repository layer
	userRepo := repository.NewUserRepository(database.GetDB())

	// Initialize the service layer
	return service.NewUserService(userRepo)
}

// Test user registration, login, password update, and deletion.
func TestRegister(t *testing.T) {
	service := setupTestUserService(t)
	username := "testregist"
	password := "123456"
	newPassword := "abc123"

	// Register a new user
	user, err := service.Register(username, password)
	assert.NoError(t, err)
	assert.True(t, user.ID > 1, "User registration failure")

	// Register a new user
	_, err10 := service.Register(username, password)
	assert.True(t, err10 != nil, "User repeat registration")

	// Verify that the user exists
	getUser, _ := service.GetUserByUsername(username)
	assert.NotNil(t, getUser)
	assert.True(t, getUser.Username == username, "User registration failed")

	// Attempt to log in
	loginUser, _ := service.Login(username, password)
	assert.NotNil(t, loginUser)
	assert.True(t, loginUser.Username == username, "Login failed")

	// Update the user's password
	err3 := service.UpdatePassword(getUser.ID, password, newPassword)
	assert.NoError(t, err3)

	// Delete the user
	err4 := service.DeleteUser(loginUser.ID)
	assert.NoError(t, err4)
}
