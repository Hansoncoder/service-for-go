package service

import (
	"veo/internal/models"
	"veo/internal/repository"
	"veo/pkg/errors"

	"golang.org/x/crypto/bcrypt"
)

// import func on errors
var (
	CodeError       = errors.CodeError
	NewError        = errors.New
	NewUserExists   = errors.NewUserExists
	NewUserNotFound = errors.NewUserNotFound
	NewAuthFailed   = errors.NewAuthFailed
)

// UserService handles user-related business logic
type UserService struct {
	userRepo *repository.UserRepository
}

// NewUserService creates a new instance of UserService
func NewUserService(userRepo *repository.UserRepository) UserService {
	return UserService{userRepo: userRepo}
}

// Register creates a new user account
func (s *UserService) Register(username, password string) (*models.User, error) {
	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, NewError(CodeError, "Failed to hash password")
	}

	// Create new user
	user := &models.User{
		Username: username,
		Password: string(hashedPassword),
	}

	return user, s.userRepo.CreateUser(user)
}

// Login authenticates a user and returns user information
func (s *UserService) Login(username, password string) (*models.User, error) {
	// Get user by username
	user, err := s.userRepo.GetUserByUsername(username)
	if err != nil {
		return nil, NewUserNotFound("User does not exist")
	}

	// Verify password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, NewAuthFailed("Invalid password")
	}

	return user, nil
}

// GetUserByID retrieves a user by their ID
func (s *UserService) GetUserByID(id int) (*models.User, error) {
	return s.userRepo.GetUserByID(id)
}

// GetUserByUsername retrieves a user by their username
func (s *UserService) GetUserByUsername(username string) (*models.User, error) {
	return s.userRepo.GetUserByUsername(username)
}

// UpdatePassword changes a user's password
func (s *UserService) UpdatePassword(userID int, oldPassword, newPassword string) error {
	// Get user by ID
	user, err := s.userRepo.GetUserByID(userID)
	if err != nil {
		return NewUserNotFound("User not found")
	}

	// Verify old password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(oldPassword))
	if err != nil {
		return NewAuthFailed("Invalid old password")
	}

	// Hash new password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return NewError(CodeError, "Failed to hash new password")
	}

	// Update password
	return s.userRepo.UpdatePassword(userID, string(hashedPassword))
}

// DeleteUser removes a user account by ID.
func (s *UserService) DeleteUser(id int) error {
	return s.userRepo.DeleteUser(id)
}
