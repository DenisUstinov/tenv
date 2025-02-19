This project is based on envconfig. See NOTES.md for license details.

---

# tenv - Library Installation and Usage

## Installation

To use the `tenv` library for managing environment variable-based configuration in Go, follow these steps:

1. **Install the `tenv` library** by running the following command:
   ```bash
   go get github.com/DenisUstinov/tenv
   ```

2. **Import the `tenv` library** in your Go code:
   ```go
   import "github.com/DenisUstinov/tenv"
   ```

## Usage

The `tenv` library allows you to load environment variables into your Go struct fields. Here's a step-by-step guide to using it:

1. **Define your configuration struct**: Create a struct that holds the configuration fields, each field mapped to an environment variable using `tenv` tags.

   Example:
   ```go
   type Config struct {
       Host     string `tenv:"HOST"`
       Port     int    `tenv:"PORT"`
       Debug    bool   `tenv:"DEBUG"`
       LogLevel string `tenv:"LOG_LEVEL"`
   }
   ```

2. **Set environment variables**: Before running the application, make sure to set the necessary environment variables that correspond to the struct fields.

   Example:
   ```bash
   export HOST="localhost"
   export PORT="8080"
   export DEBUG="true"
   export LOG_LEVEL="info"
   ```

3. **Populate the struct with environment variables**: Use `tenv.PopulateFromEnv` to automatically fill the struct fields with the corresponding values from the environment.

   Example code:
   ```go
   var cfg Config
   err := tenv.PopulateFromEnv(&cfg)
   if err != nil {
       log.Fatalf("Error processing configuration: %v", err)
   }
   ```

4. **Access the configuration**: Once the struct is populated, you can access the values in your Go code.

   Example:
   ```go
   fmt.Printf("Host: %s\n", cfg.Host)
   fmt.Printf("Port: %d\n", cfg.Port)
   fmt.Printf("Debug: %v\n", cfg.Debug)
   fmt.Printf("LogLevel: %s\n", cfg.LogLevel)
   ```

5. **Handling the configuration**: You can use the populated values to control application behavior.

   Example:
   ```go
   if !cfg.Debug {
       fmt.Println("Debug mode is off.")
   } else {
       fmt.Println("Debug mode is on.")
   }
   ```

## Example Code

Here is a complete example of how to use the `tenv` library to load environment variables into a struct and print the values:

```go
package main

import (
    "fmt"
    "log"
    "os"

    tenv "github.com/DenisUstinov/tenv" // Import the library
)

type Config struct {
    Host     string `tenv:"HOST"`
    Port     int    `tenv:"PORT"`
    Debug    bool   `tenv:"DEBUG"`
    LogLevel string `tenv:"LOG_LEVEL"`
}

func main() {
    // Set environment variables (for testing)
    os.Setenv("HOST", "localhost")
    os.Setenv("PORT", "8080")
    os.Setenv("DEBUG", "true")
    os.Setenv("LOG_LEVEL", "info")

    // Initialize the config structure
    var cfg Config

    // Populate the struct with environment variables
    err := tenv.PopulateFromEnv(&cfg)
    if err != nil {
        log.Fatalf("Error processing configuration: %v", err)
    }

    // Print the received values
    fmt.Printf("Host: %s\n", cfg.Host)
    fmt.Printf("Port: %d\n", cfg.Port)
    fmt.Printf("Debug: %v\n", cfg.Debug)
    fmt.Printf("LogLevel: %s\n", cfg.LogLevel)

    // Example logic based on the config
    if !cfg.Debug {
        fmt.Println("Debug mode is off.")
    } else {
        fmt.Println("Debug mode is on.")
    }
}
```

---

## Summary

- Install the library using `go get`.
- Define a struct with `tenv` tags for your configuration fields.
- Use `tenv.PopulateFromEnv` to load environment variables into the struct.
- Access the struct fields to control your application's behavior.

By using the `tenv` library, you can manage application configurations in a clean and simple way through environment variables.

