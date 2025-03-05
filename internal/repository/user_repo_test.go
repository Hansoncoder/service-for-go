package repository_test

import (
	"log"
	"testing"
	configs "veo/internal/config"
	"veo/internal/database"
	"veo/internal/repository"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

// setupTestDB initializes the test database.
func setupTestDB(t *testing.T) *gorm.DB {
	// Show the current path
	// cwd, _ := os.Getwd()
	// fmt.Println("Current working directory:", cwd)
	// Expected current path: ./internal/repository

	// Load the configuration file
	// Config path: ../../config/config.yaml
	cfg, err := configs.Load("../../config/config.yaml")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Initialize the database
	if err := database.Init(cfg.Database); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	return database.GetDB()
}

// TestUpdatePassword verifies that a user's password can be successfully updated.
func TestUpdatePassword(t *testing.T) {
	db := setupTestDB(t)
	userID := 1
	newPassword := "123456"

	// Create a user repository instance
	repo := repository.NewUserRepository(db)

	// Update the password
	err := repo.UpdatePassword(userID, newPassword)
	assert.NoError(t, err)

	// Retrieve the user and verify the updated password
	user, err := repo.GetUserByUserID(userID)
	assert.NoError(t, err)
	result := user.CheckPassword(newPassword)
	assert.True(t, result, "Password verification successful")
}

// TestCheckPassword verifies that the stored password can be correctly validated.
func TestCheckPassword(t *testing.T) {
	db := setupTestDB(t)
	userID := 2
	password := "123456"

	// Create a user repository instance
	repo := repository.NewUserRepository(db)

	// Retrieve the user and check password validation
	user, err := repo.GetUserByUserID(userID)
	assert.NoError(t, err)
	result := user.CheckPassword(password)
	assert.True(t, result, "Password verification successful")
}
