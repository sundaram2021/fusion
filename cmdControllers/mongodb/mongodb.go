package mongodb

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
	"io/ioutil"
)

// LoadAndValidateConfig loads the YAML configuration and validates it.
func LoadAndValidateConfig(filePath string) error {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
	    return fmt.Errorf("error reading file: %w", err)
	}

	var db Database
	err = yaml.Unmarshal(data, &db)
	if err != nil {
	    return fmt.Errorf("error parsing YAML: %w", err)
	}

	if err = ValidateDatabase(db); err != nil {
	    return fmt.Errorf("validation error: %w", err)
	}

	fmt.Println("Database structure is valid!")
	return nil
}

// Main function to execute loading and validating configuration from YAML file.
func Main() {
	if len(os.Args) < 3 || os.Args[1] != "apply" || os.Args[2] == "" {
	    fmt.Println("Usage: fusion apply -f <file.yaml>")
	    return
	}

	filePath := os.Args[2]
	if err := LoadAndValidateConfig(filePath); err != nil {
	    fmt.Println(err)
	    return
	}
	
	fmt.Println("Proceeding with creating the MongoDB database...")
}