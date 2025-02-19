package main

import (
	"fmt"
	"log"
	"os"

	tenv "github.com/DenisUstinov/tenv" // Import the 'tenv' library you created
)

// Config struct represents the application configuration,
// with environment variables mapped to struct fields using 'tenv' tags.
type Config struct {
	Host     string `tenv:"HOST"`      // The 'Host' field will be populated from the 'HOST' environment variable
	Port     int    `tenv:"PORT"`      // The 'Port' field will be populated from the 'PORT' environment variable
	Debug    bool   `tenv:"DEBUG"`     // The 'Debug' field will be populated from the 'DEBUG' environment variable
	LogLevel string `tenv:"LOG_LEVEL"` // The 'LogLevel' field will be populated from the 'LOG_LEVEL' environment variable
}

func main() {
	// Set environment variables for testing purposes
	// These environment variables will be used to populate the 'Config' struct fields
	os.Setenv("HOST", "localhost") // Set the 'HOST' variable to "localhost"
	os.Setenv("PORT", "8080")      // Set the 'PORT' variable to "8080"
	os.Setenv("DEBUG", "true")     // Set the 'DEBUG' variable to "true"
	os.Setenv("LOG_LEVEL", "info") // Set the 'LOG_LEVEL' variable to "info"

	// Initialize the configuration structure (Config) to hold the environment variable values
	var cfg Config

	// Populate the struct with the values from environment variables
	// This function will use the tags in the struct fields to look for the corresponding environment variables
	err := tenv.PopulateFromEnv(&cfg)
	if err != nil {
		// If there is an error in the process of loading environment variables, log and exit
		log.Fatalf("Error processing configuration: %v", err)
	}

	// Print the populated values to verify that the struct has been properly filled
	fmt.Printf("Host: %s\n", cfg.Host)         // Print the 'Host' field value (should be "localhost")
	fmt.Printf("Port: %d\n", cfg.Port)         // Print the 'Port' field value (should be 8080)
	fmt.Printf("Debug: %v\n", cfg.Debug)       // Print the 'Debug' field value (should be true)
	fmt.Printf("LogLevel: %s\n", cfg.LogLevel) // Print the 'LogLevel' field value (should be "info")

	// Example logic based on the configuration values
	// The logic demonstrates how you can use the populated configuration for different behaviors
	if !cfg.Debug {
		// If Debug mode is off, print this message
		fmt.Println("Debug mode is off.")
	} else {
		// If Debug mode is on, print this message
		fmt.Println("Debug mode is on.")
	}
}
