package mongodb

import (
	"errors"
	"regexp"
)

// ValidateDatabase checks if the database structure is valid.
func ValidateDatabase(db Database) error {
	if db.Name == "" {
		return errors.New("database name is required")
	}

	if db.Engine == "" {
		return errors.New("engine type is required")
	}

	for _, collection := range db.Collections {
		if err := ValidateCollection(collection); err != nil {
			return err
		}
	}

	return nil
}

// ValidateCollection checks if the collection structure is valid.
func ValidateCollection(collection Collection) error {
	if collection.Name == "" {
		return errors.New("collection name is required")
	}

	if len(collection.Documents) == 0 {
		return errors.New("at least one document is required in collection " + collection.Name)
	}

	for _, doc := range collection.Documents {
		if err := ValidateDocument(doc); err != nil {
			return err
		}
	}

	return nil
}

// ValidateDocument checks if the document structure is valid.
func ValidateDocument(doc Document) error {
	if doc.Field == "" || doc.Type == "" {
		return errors.New("document field and type are required")
	}
	return nil
}

// ValidateForeignKey checks if the foreign key structure is valid.
func ValidateForeignKey(fk ForeignKey) error {
	if fk.Name == "" || fk.Field == "" || fk.References.Collection == "" || fk.References.Field == "" {
		return errors.New("foreign key attributes must not be empty")
	}
	return nil
}

// ValidateEmailFormat checks if the email format is valid using regex.
func ValidateEmailFormat(email string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}