
package postgresql

import (
	"testing"
)

func TestValidateDatabase(t *testing.T) {
	tests := []struct {
		name     string
		db       Database
		expected string // Expected error message or empty for no errors expected
	}{
		{
			name: "valid database",
			db: Database{
				Name:      "my_postgresql_db",
				Charset:   "UTF8",
				Collation: "en_US.UTF-8",
				Tables: []Table{
					{
						Name: "customers",
						Columns: []Column{
							{Name: "customer_id", Type: "SERIAL", PrimaryKey: true},
						},
					},
				},
			},
			expected: "",
		},
		{
			name: "missing name",
			db: Database{
				Charset:   "UTF8",
				Collation: "en_US.UTF-8",
				Tables:    []Table{{Name: "customers"}},
			},
			expected: "database name is required",
		},
		{
			name: "missing charset",
			db: Database{
				Name:      "my_postgresql_db",
				Collation: "en_US.UTF-8",
				Tables:    []Table{{Name: "customers"}},
			},
			expected: "charset is required",
		},
		{
			name: "missing collation",
			db: Database{
				Name:    "my_postgresql_db",
				Charset: "UTF8",
				Tables:  []Table{{Name: "customers"}},
			},
			expected: "collation is required",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateDatabase(tt.db)
			if (err != nil && err.Error() != tt.expected) || (err == nil && tt.expected != "") {
				t.Errorf("expected error %q, got %q", tt.expected, err)
			}
		})
	}
}

func TestValidateTable(t *testing.T) {
	tests := []struct {
		name     string
		table    Table
		expected string // Expected error message or empty for no errors expected
	}{
		{
			name:  "valid table",
			table: Table{Name: "customers", Columns: []Column{{Name: "customer_id", Type: "SERIAL", PrimaryKey: true}}},
			expected: "",
		},
        {
            name:  "missing table name",
            table: Table{Columns: []Column{{Name: "customer_id", Type: "SERIAL", PrimaryKey: true}}},
            expected: "table name is required",
        },
        {
            name:  "no columns",
            table: Table{Name: "customers"},
            expected: "at least one column is required in table customers",
        },
        {
            name:  "no primary key",
            table: Table{Name: "customers", Columns: []Column{{Name: "customer_id", Type: "SERIAL"}}},
            expected: "at least one primary key is required in table customers",
        },
    }

	for _, tt := range tests {
	    t.Run(tt.name, func(t *testing.T) {
	        err := ValidateTable(tt.table)
	        if (err != nil && err.Error() != tt.expected) || (err == nil && tt.expected != "") {
	            t.Errorf("expected error %q, got %q", tt.expected, err)
	        }
	    })
    }
}

func TestValidateForeignKey(t *testing.T) {
	tests := []struct {
	    name     string
	    fk       ForeignKey
	    tableName string
	    expected string // Expected error message or empty for no errors expected
    }{
        {
            name:     "valid foreign key",
            fk:       ForeignKey{Name:"fk_customer_order", Column:"customer_id", References : Reference{Table:"customers", Column:"customer_id"}},
            tableName:"orders",
            expected:"",
        },
        {
            name:"missing foreign key name",
            fk : ForeignKey{Column:"customer_id", References : Reference{Table:"customers", Column:"customer_id"}},
            tableName:"orders",
            expected:"foreign key attributes must not be empty in table orders",
        },
        {
            name:"missing foreign key column",
            fk : ForeignKey{Name:"fk_customer_order", References : Reference{Table:"customers", Column:"customer_id"}},
            tableName:"orders",
            expected:"foreign key attributes must not be empty in table orders",
        },
        {
            name:"missing referenced table",
            fk : ForeignKey{Name:"fk_customer_order", Column:"customer_id", References : Reference{Column:"customer_id"}},
            tableName:"orders",
            expected:"foreign key attributes must not be empty in table orders",
        },
        {
            name:"missing referenced column",
            fk : ForeignKey{Name:"fk_customer_order", Column:"customer_id", References : Reference{Table:"customers"}},
            tableName:"orders",
            expected:"foreign key attributes must not be empty in table orders",
        },
    }

	for _, tt := range tests {
	    t.Run(tt.name, func(t *testing.T) {
	        err := ValidateForeignKey(tt.fk, tt.tableName)
	        if (err != nil && err.Error() != tt.expected) || (err == nil && tt.expected != "") {
	            t.Errorf("expected error %q, got %q", tt.expected, err)
	        }
	    })
    }
}