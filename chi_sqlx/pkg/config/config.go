package config

import (
	"log"

	"github.com/spf13/viper"
)

// Config holds the configuration values
type Config struct {
	DBUser     string
	DBPassword string
	DBName     string
	DBHost     string
	DBPort     string
}

// LoadConfig loads environment variables using Viper
func LoadConfig() Config {
	// Set the file name of the configuration file without the extension
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	// Set the path to look for the configuration file
	viper.AddConfigPath(".") // or specify a specific directory, e.g., "/path/to/env/"

	// Enable viper to read environment variables
	viper.AutomaticEnv()

	// Read in the configuration file
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Return the configuration struct with values from Viper
	return Config{
		DBUser:     viper.GetString("DB_USER"),
		DBPassword: viper.GetString("DB_PASSWORD"),
		DBName:     viper.GetString("DB_NAME"),
		DBHost:     viper.GetString("DB_HOST"),
		DBPort:     viper.GetString("DB_PORT"),
	}
}
