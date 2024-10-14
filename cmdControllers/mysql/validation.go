package mysql

import (
	"errors"
)

// ValidateDatabase checks if the database structure is valid.
func ValidateDatabase(db Database) error {
	if db.Name == "" {
		return errors.New("database name is required")
	}

	if db.Charset == "" {
		return errors.New("charset is required")
	}

	if db.Collation == "" {
		return errors.New("collation is required")
	}

	for _, table := range db.Tables {
		if err := ValidateTable(table); err != nil {
			return err
		}
	}

	return nil
}

// ValidateTable checks if the table structure is valid.
func ValidateTable(table Table) error {
	if table.Name == "" {
		return errors.New("table name is required")
	}

	if len(table.Columns) == 0 {
		return errors.New("at least one column is required in table " + table.Name)
	}

	primaryKeyCount := 0
	for _, column := range table.Columns {
		if column.PrimaryKey {
			primaryKeyCount++
		}

		if column.Name == "" || column.Type == "" {
			return errors.New("column name and type are required in table " + table.Name)
		}
	}

	if primaryKeyCount == 0 {
		return errors.New("at least one primary key is required in table " + table.Name)
	}

	for _, fk := range table.ForeignKeys {
		if err := ValidateForeignKey(fk, table.Name); err != nil {
			return err
		}
	}

	return nil
}

// ValidateForeignKey checks if the foreign key structure is valid.
func ValidateForeignKey(fk ForeignKey, tableName string) error {
	if fk.Name == "" || fk.Column == "" || fk.References.Table == "" || fk.References.Column == "" {
		return errors.New("foreign key attributes must not be empty in table " + tableName)
	}
	return nil
}
