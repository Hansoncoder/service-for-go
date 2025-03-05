package service

import (
	"veo/internal/models"
	"veo/internal/repository"
	"veo/internal/utils"
	CustomError "veo/pkg/errors"
)

var logger = utils.GetLogger()

// userService implements the UserService interface.
type userService struct {
	userRepo repository.UserRepository
}

// NewUserService creates a new instance of userService.
func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

// UserService defines the interface for user-related operations.
type UserService interface {
	// Register creates a new user with the given username and password.
	Register(username, password string) error

	// Login authenticates a user by username and password, returning the user object on success.
	Login(username, password string) (*models.User, error)

	// GetUserByUsername retrieves a user by their username.
	GetUserByUsername(username string) (*models.User, error)

	// GetUserByID retrieves a user by their unique ID.
	GetUserByID(id int) (*models.User, error)

	// UpdatePassword changes a user's password after verifying the old password.
	UpdatePassword(id int, oldPassword, newPassword string) error

	// DeleteUser removes a user account by ID.
	DeleteUser(id int) error
}

// GetUserByID retrieves a user by their unique ID.
func (s *userService) GetUserByID(id int) (*models.User, error) {
	return s.userRepo.GetUserByUserID(id)
}

// GetUserByUsername retrieves a user by their username.
func (s *userService) GetUserByUsername(username string) (*models.User, error) {
	return s.userRepo.GetUserByUsername(username)
}

// Register creates a new user with the given username and password.
func (s *userService) Register(username, password string) error {
	if username == "" || password == "" {
		return CustomError.NewError(CustomError.CodeAuthFailed, "username or password is nil")
	}

	// Check if the user already exists
	user, _ := s.userRepo.GetUserByUsername(username)
	if user != nil {
		logger.Info("user already exists")
		return CustomError.NewError(CustomError.CodeAuthFailed, "user already exists")
	}

	// Hash the password before saving
	savePassword, err1 := models.GetHashedPassword(password)
	if err1 != nil {
		logger.Error(err1.Error())
		return err1
	}

	user1 := models.User{
		Username: username,
		Password: savePassword,
	}
	return s.userRepo.UpdateUser(&user1)
}

// Login authenticates a user using their username and password.
func (s *userService) Login(username, password string) (*models.User, error) {
	user, err := s.userRepo.GetUserByUsername(username)
	if err != nil {
		return nil, err
	}

	// Verify the password
	if !user.CheckPassword(password) {
		return nil, CustomError.NewError(CustomError.CodeAuthFailed, "password fail")
	}

	return user, nil
}

// UpdatePassword updates a user's password after verifying the old password.
func (s *userService) UpdatePassword(id int, oldPassword, newPassword string) error {
	user, err := s.userRepo.GetUserByUserID(id)
	if err != nil {
		return err
	}

	// Verify the old password before updating
	if !user.CheckPassword(oldPassword) {
		return CustomError.NewError(CustomError.CodeAuthFailed, "password fail")
	}
	return s.userRepo.UpdatePassword(id, newPassword)
}

// DeleteUser removes a user account by ID.
func (s *userService) DeleteUser(id int) error {
	return s.userRepo.DeleteUser(id)
}
