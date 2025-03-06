package repository_test

import (
	"log"
	"testing"
	"veo/internal/configs"
	"veo/internal/database"
	"veo/internal/models"
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
	cfg, err := configs.Load("../../../config/config.yaml")
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
	user, err := repo.GetUserByID(userID)
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
	user, err := repo.GetUserByID(userID)
	assert.NoError(t, err)
	result := user.CheckPassword(password)
	assert.True(t, result, "Password verification successful")
}

func TestUserRepository_CreateUser(t *testing.T) {
	// Setup test database
	db := setupTestDB(t)
	userRepo := repository.NewUserRepository(db)

	tests := []struct {
		name          string
		user          *models.User
		expectedError bool
	}{
		{
			name: "successful creation",
			user: &models.User{
				Username: "testuser",
				Password: "hashedpassword",
			},
			expectedError: false,
		},
		{
			name: "duplicate username",
			user: &models.User{
				Username: "testuser", // Same username as above
				Password: "hashedpassword",
			},
			expectedError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := userRepo.CreateUser(tt.user)
			if tt.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				// Verify user was created
				user, err := userRepo.GetUserByUsername(tt.user.Username)
				assert.NoError(t, err)
				assert.NotNil(t, user)
				assert.Equal(t, tt.user.Username, user.Username)
			}
		})
	}
}

func TestUserRepository_GetUserByUsername(t *testing.T) {
	// Setup test database
	db := setupTestDB(t)
	userRepo := repository.NewUserRepository(db)

	// Create test user
	testUser := &models.User{
		Username: "testuser",
		Password: "hashedpassword",
	}
	err := userRepo.CreateUser(testUser)
	assert.NoError(t, err)

	tests := []struct {
		name          string
		username      string
		expectedError bool
	}{
		{
			name:          "existing user",
			username:      "testuser",
			expectedError: false,
		},
		{
			name:          "non-existing user",
			username:      "nonexistent",
			expectedError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user, err := userRepo.GetUserByUsername(tt.username)
			if tt.expectedError {
				assert.Error(t, err)
				assert.Nil(t, user)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, user)
				assert.Equal(t, tt.username, user.Username)
			}
		})
	}
}
