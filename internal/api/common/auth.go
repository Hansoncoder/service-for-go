package common

import (
	"time"
	"veo/pkg/errors"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// Secret key used for signing JWT tokens
var jwtKey = []byte("com.hanson.test.jwt.secret.key")

// Define the JWT expiration duration as 5 hours
const (
	JWTExpirationDuration = 5 * time.Hour
)

// UserClaims defines the JWT claims structure
type UserClaims struct {
	ID       int    `json:"userId"`   // User ID
	Username string `json:"username"` // Username
	jwt.StandardClaims
}

// GenerateJWT generates a JWT token for a given user ID and username.
func GenerateJWT(ID int, username string) (string, error) {
	expirationTime := time.Now().Add(JWTExpirationDuration) // Set expiration time
	claims := &UserClaims{
		ID:       ID,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(), // Token expiration timestamp
		},
	}

	// Create a new JWT token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token using the secret key
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// AuthMiddleware is a JWT authentication middleware.
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Retrieve the JWT token from the Authorization header
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			// Return an error if no token is provided
			RespondError(c, errors.NewAuthFailed("Authorization fail"))
			c.Abort()
			return
		}

		// Parse the token and extract claims
		claims := &UserClaims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		// Check if the token is valid
		if err != nil || !token.Valid {
			// Return an error if the token is expired or invalid
			RespondError(c, errors.NewTokenExpired("Authorization fail"))
			c.Abort()
			return
		}

		// Store user information in the request context
		c.Set("userId", claims.ID)
		c.Set("username", claims.Username)

		// Continue with the request
		c.Next()
	}
}
