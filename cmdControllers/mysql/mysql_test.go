package mysql

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
				Name:      "my-mysql-db",
				Charset:   "utf8mb4",
				Collation: "utf8mb4_unicode_ci",
				Tables:    []Table{{Name: "customers", Columns: []Column{{Name: "customer_id", Type: "INT", PrimaryKey: true}}}},
			},
			expected: "",
		},
		{
			name: "missing name",
			db: Database{
				Charset:   "utf8mb4",
				Collation: "utf8mb4_unicode_ci",
				Tables:    []Table{{Name: "customers"}},
			},
			expected: "database name is required",
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
