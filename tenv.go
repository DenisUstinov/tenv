package envconfig

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

// ErrInvalidEnvVarSpecification is returned when the passed parameter is not a pointer to a struct.
var ErrInvalidEnvVarSpecification = errors.New("specification must be a struct pointer")

// EnvVarInfo holds information about an environment variable,
// such as the field name, an alternative name for the environment variable,
// the key used to look up the environment variable, the reference to the field itself, and its tags.
type EnvVarInfo struct {
	Name  string            // Field name in the struct
	Alt   string            // Alternative name for the environment variable (if specified in tags)
	Key   string            // Key used to search for the environment variable
	Field reflect.Value     // Reference to the field where the value will be set
	Tags  reflect.StructTag // Struct tags containing metadata
}

// CollectFieldInfo extracts information about the fields of a struct that corresponds to environment variables.
// This method analyzes the provided pointer to a struct and creates a list of environment variables that map to the struct fields.
func CollectFieldInfo(spec interface{}) ([]EnvVarInfo, error) {
	// Check if the provided parameter is a pointer to a struct
	s := reflect.ValueOf(spec)
	if s.Kind() != reflect.Ptr {
		return nil, ErrInvalidEnvVarSpecification
	}
	// Dereference the pointer to get the struct itself
	s = s.Elem()
	if s.Kind() != reflect.Struct {
		return nil, ErrInvalidEnvVarSpecification
	}

	// Get the type information of the struct for further analysis
	typeOfSpec := s.Type()
	infos := make([]EnvVarInfo, 0, s.NumField()) // Slice to store information about each environment variable

	// Iterate over all fields in the struct
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)              // Get the field value
		ftype := typeOfSpec.Field(i) // Get the type information of the field
		if !f.CanSet() {
			continue // Skip the field if it cannot be set
		}

		// Handle possible pointers to structs and initialize them if necessary
		for f.Kind() == reflect.Ptr {
			if f.IsNil() {
				if f.Type().Elem().Kind() != reflect.Struct {
					break
				}
				f.Set(reflect.New(f.Type().Elem())) // Create a new instance of the struct if it's nil
			}
			f = f.Elem() // Move to the next level of dereferencing the pointer
		}

		// Create an EnvVarInfo object for this field
		info := EnvVarInfo{
			Name:  ftype.Name,                             // Field name
			Field: f,                                      // Reference to the field
			Tags:  ftype.Tag,                              // Tags of the field
			Alt:   strings.ToUpper(ftype.Tag.Get("tenv")), // Alternative environment variable name
		}
		// Set the key for looking up the environment variable
		info.Key = info.Name
		if info.Alt != "" {
			info.Key = info.Alt // Use the alternative name if provided
		}
		info.Key = strings.ToUpper(info.Key) // The key should always be in uppercase
		infos = append(infos, info)          // Add the environment variable info to the slice
	}

	return infos, nil
}

// PopulateFromEnv populates a struct with values from environment variables.
// For each field in the struct, it looks for the corresponding environment variable, and if found,
// the value is converted and assigned to the field.
func PopulateFromEnv(spec interface{}) error {
	// First, extract the field information
	infos, err := CollectFieldInfo(spec)
	if err != nil {
		return err
	}

	// Iterate over the environment variable information
	for _, info := range infos {
		// Look for the environment variable by its key
		value, ok := os.LookupEnv(info.Key)
		if !ok && info.Alt != "" {
			// If the main key is not found, try the alternative name
			value, ok = os.LookupEnv(info.Alt)
		}
		if !ok || value == "" {
			continue // If the environment variable is not found, skip this field
		}

		// Convert and assign the value to the field
		err = ConvertAndSetField(value, info.Field)
		if err != nil {
			return fmt.Errorf("error processing field %s: %w", info.Name, err)
		}
	}

	return nil
}

// ConvertAndSetField converts the environment variable string value to the appropriate field type and assigns it to the field.
func ConvertAndSetField(value string, field reflect.Value) error {
	// Check the field type and convert the value accordingly
	switch field.Kind() {
	case reflect.String:
		field.SetString(value) // Set the value as a string
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		val, err := strconv.ParseInt(value, 0, field.Type().Bits()) // Convert to an integer
		if err != nil {
			return fmt.Errorf("invalid value for int: %s", value)
		}
		field.SetInt(val)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		val, err := strconv.ParseUint(value, 0, field.Type().Bits()) // Convert to an unsigned integer
		if err != nil {
			return fmt.Errorf("invalid value for uint: %s", value)
		}
		field.SetUint(val)
	case reflect.Bool:
		val, err := strconv.ParseBool(value) // Convert to a boolean
		if err != nil {
			return fmt.Errorf("invalid value for bool: %s", value)
		}
		field.SetBool(val)
	case reflect.Float32, reflect.Float64:
		val, err := strconv.ParseFloat(value, field.Type().Bits()) // Convert to a floating-point number
		if err != nil {
			return fmt.Errorf("invalid value for float: %s", value)
		}
		field.SetFloat(val)
	}

	return nil
}
