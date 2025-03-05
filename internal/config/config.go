package configs

import (
	"github.com/spf13/viper"
)

// Config represents the main application configuration structure.
type Config struct {
	Database DBConfig // Database configuration
}

// DBConfig holds the database connection details.
type DBConfig struct {
	Username string // Database username
	Password string // Database password
	Host     string // Database host address
	Port     int    // Database port number
	DBName   string // Database name
	Charset  string // Character set for the database
}

// Load reads the configuration file from the specified path and unmarshals it into the Config struct.
func Load(configPath string) (*Config, error) {
	viper.SetConfigFile(configPath) // Set the path of the configuration file

	// Read the configuration file
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var cfg Config
	// Unmarshal the configuration file into the Config struct
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
