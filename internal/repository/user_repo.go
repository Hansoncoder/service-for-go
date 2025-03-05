package repository

import (
	"errors"
	"veo/internal/models"
	CustomError "veo/pkg/errors"

	"gorm.io/gorm"
)

// UserRepository defines methods for user data operations.
type UserRepository interface {
	GetUserByUserID(id int) (*models.User, error)            // Retrieve user by ID
	GetUserByUsername(username string) (*models.User, error) // Retrieve user by username
	UpdateUser(user *models.User) error                      // Update user information
	UpdatePassword(id int, password string) error            // Update user password
	DeleteUser(id int) error                                 // Delete user by ID
}

// userRepo is an implementation of UserRepository.
type userRepo struct {
	db *gorm.DB // Database instance,
}

// NewUserRepository creates a new instance of UserRepository.
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepo{db: db}
}

// GetUserByUserID retrieves a user by their ID.
func (r *userRepo) GetUserByUserID(id int) (*models.User, error) {
	var user models.User
	err := r.db.Where("ID = ?", id).First(&user).Error
	if err != nil {
		// Handle case when the user is not found
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, CustomError.NewError(CustomError.CodeUserNotFound, "User does not exist")
		}
		return nil, err
	}
	return &user, nil
}

// GetUserByUsername retrieves a user by their username.
func (r *userRepo) GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	err := r.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		// Handle case when the user is not found
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, CustomError.NewError(CustomError.CodeUserNotFound, "User does not exist")
		}
		return nil, err
	}
	return &user, nil
}

// User represents the database schema for users.
type User struct {
	ID       int    `gorm:"primaryKey;autoIncrement"`
	Username string `gorm:"unique"` // Unique constraint on username
	Password string // Hashed password
}

// UpdateUser updates the user information in the database.
func (r *userRepo) UpdateUser(user *models.User) error {
	return r.db.Save(user).Error
}

// DeleteUser removes a user from the database based on their ID.
func (r *userRepo) DeleteUser(id int) error {
	// Use GORM's Delete method to remove the user
	result := r.db.Delete(&User{}, id)
	if result.Error != nil {
		return result.Error
	}

	// Check if a record was actually deleted
	if result.RowsAffected == 0 {
		return CustomError.NewError(CustomError.CodeUserNotFound, "User does not exist")
	}

	return nil
}

// UpdatePassword updates a user's password in the database.
func (r *userRepo) UpdatePassword(id int, password string) error {
	// Hash the password before storing
	hashedPassword, err := models.GetHashedPassword(password)
	if err != nil {
		return err
	}

	// Execute the database update operation
	result := r.db.Model(&models.User{}).Where("id = ?", id).Update("password", string(hashedPassword))
	if result.Error != nil {
		return result.Error
	}

	// Check if any record was updated
	if result.RowsAffected == 0 {
		return CustomError.NewError(CustomError.CodeUserNotFound, "User not found or password not updated")
	}

	return nil
}
