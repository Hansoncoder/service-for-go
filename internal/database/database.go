package database

import (
	"fmt"
	"sync"
	"veo/internal/configs"
	"veo/internal/utils"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db   *gorm.DB  // Global database connection instance
	once sync.Once // Ensures database initialization happens only once
)

var logger = utils.GetLogger()

// Init initializes the database connection using the provided configuration.
func Init(config configs.DBConfig) error {
	var initErr error
	once.Do(func() {
		// Construct the Data Source Name (DSN) for MySQL connection
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s",
			config.Username,
			config.Password,
			config.Host,
			config.Port,
			config.DBName,
			config.Charset,
		)

		var err error
		// Open a connection to the database
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			logger.Errorf("failed to connect to database: %v", err)
			return
		}

		// Retrieve the underlying *sql.DB object
		sqlDB, err := db.DB()
		if err != nil {
			logger.Errorf("failed to get DB object: %v", err)
			return
		}

		// Configure database connection pooling
		sqlDB.SetMaxIdleConns(10)   // Set the maximum number of idle connections
		sqlDB.SetMaxOpenConns(100)  // Set the maximum number of open connections
		sqlDB.SetConnMaxLifetime(0) // Disable connection timeout
		logger.Info("Database connected successfully")
	})
	return initErr
}

// GetDB returns the database connection instance.
func GetDB() *gorm.DB {
	return db
}

// Close closes the database connection.
func Close() {
	if db == nil {
		return
	}

	sqlDB, err := db.DB()
	if err != nil {
		logger.Errorf("Failed to get DB object: %v", err)
		return
	}

	if err := sqlDB.Close(); err != nil {
		logger.Errorf("Failed to close database connection: %v", err)
	}
}
