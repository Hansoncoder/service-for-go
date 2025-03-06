package repository

import (
	"veo/internal/models"
	"veo/internal/utils"
	"veo/pkg/errors"

	"gorm.io/gorm"
)

var logger = utils.GetLogger()
var NewUserNotFound = errors.NewUserNotFound

// UserRepository handles database operations for users
type UserRepository struct {
	db *gorm.DB
}

// NewUserRepository creates a new instance of UserRepository
func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

// CreateUser creates a new user in the database
func (r *UserRepository) CreateUser(user *models.User) error {
	var existingUser models.User
	if err := r.db.Where("username = ?", user.Username).First(&existingUser).Error; err == nil {
		return errors.NewUserExists("user exists '" + user.Username + "'")
	} else if err != gorm.ErrRecordNotFound {
		return err
	}

	if err := r.db.Create(user).Error; err != nil {
		logger.Error(err.Error())
		return err
	}

	return nil
}

// GetUserByID retrieves a user by their ID
func (r *UserRepository) GetUserByID(id int) (*models.User, error) {
	var user models.User
	err := r.db.First(&user, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, NewUserNotFound("User not found")
		}
		return nil, err
	}
	return &user, nil
}

// GetUserByUsername retrieves a user by their username
func (r *UserRepository) GetUserByUsername(username string) (*models.User, error) {
	var user *models.User
	err := r.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, NewUserNotFound("User not found")
		}
		return nil, err
	}
	return user, nil
}

// UpdatePassword updates a user's password
func (r *UserRepository) UpdatePassword(userID int, hashedPassword string) error {
	return r.db.Model(&models.User{}).Where("id = ?", userID).Update("password", hashedPassword).Error
}

// DeleteUser removes a user from the database
func (r *UserRepository) DeleteUser(id int) error {
	return r.db.Delete(&models.User{}, id).Error
}
