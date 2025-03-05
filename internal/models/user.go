package models

import (
	"golang.org/x/crypto/bcrypt"
)

// User represents the database model for a user.
type User struct {
	ID       int    `gorm:"primaryKey"` // Unique user ID (primary key)
	Username string `gorm:"unique"`     // Unique username
	Password string // Hashed password
}

// UserDTO is a data transfer object (DTO) for user data.
// It is used to return user information without sensitive fields.
type UserDTO struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
}

// Sanitize removes sensitive information (e.g., password) and returns a UserDTO.
func (u *User) Sanitize() UserDTO {
	return UserDTO{
		ID:       u.ID,
		Username: u.Username,
	}
}

// GetHashedPassword hashes a given password using bcrypt.
func GetHashedPassword(password string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedBytes), nil
}

// CheckPassword verifies whether the given password matches the stored hashed password.
func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
