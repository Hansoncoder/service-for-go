package main

import (
	v1 "veo/internal/api/v1"
	"veo/internal/configs"
	"veo/internal/database"
	"veo/internal/repository"
	"veo/internal/service"
	"veo/internal/utils"

	"github.com/gin-gonic/gin"
)

var logger = utils.GetLogger()

func main() {
	// Load configuration from the YAML file
	cfg, err := configs.Load("config/config.yaml")
	if err != nil {
		logger.Errorf("Failed to load config: %v", err)
	}

	// Initialize the database connection
	if err := database.Init(cfg.Database); err != nil {
		logger.Errorf("Failed to initialize database: %v", err)
	}
	defer database.Close() // Ensure the database connection is closed when the application exits

	// Initialize the repository layer (Data Access Layer)
	userRepo := repository.NewUserRepository(database.GetDB())

	// Initialize the service layer (Business Logic Layer)
	userService := service.NewUserService(userRepo)

	// Initialize the API layer (Controller Layer)
	accountAPI := v1.NewAccountAPI(userService)
	userAPI := v1.NewUserAPI(userService)

	// Start the HTTP server using the Gin framework
	router := gin.Default()

	// Set up API routes
	v1.SetupAccountRouter(router, accountAPI)
	v1.SetupUserRouter(router, userAPI)

	// Run the server on port 8080
	if err := router.Run(":8080"); err != nil {
		logger.Errorf("Failed to start server: %v", err)
	}
}
