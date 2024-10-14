package mongodb

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
				Name:      "my_mongodb_db",
				Engine:    "MongoDB",
				Collections: []Collection{
					{
						Name: "customers",
						Documents: []Document{
							{Field: "customer_id", Type: "ObjectId"},
							{Field: "name", Type: "String", Required: true},
							{Field: "email", Type: "String", Required: true, Default: ""},
							{Field: "created_at", Type: "Date", Default: "CURRENT_TIMESTAMP"},
						},
					},
				},
			},
			expected: "",
		},
		{
			name: "missing name",
			db: Database{
				Engine:    "MongoDB",
				Collections: []Collection{{Name: "customers"}},
			},
			expected: "database name is required",
		},
        {
            name: "missing engine",
            db: Database{
                Name:      "my_mongodb_db",
                Collections: []Collection{{Name: "customers"}},
            },
            expected: "engine type is required",
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

func TestValidateCollection(t *testing.T) {
	tests := []struct {
	    name     string
	    collection Collection
	    expected string // Expected error message or empty for no errors expected
    }{
        {
            name:  "valid collection",
            collection: Collection{Name: "customers", Documents: []Document{
                {Field: "customer_id", Type: "ObjectId"},
                {Field: "name", Type: "String", Required: true},
            }},
            expected: "",
        },
        {
            name:  "missing collection name",
            collection: Collection{Documents: []Document{{Field:"customer_id", Type:"ObjectId"}}},
            expected: "collection name is required",
        },
        {
            name:  "no documents",
            collection: Collection{Name:"customers"},
            expected:"at least one document is required in collection customers",
        },
    }

	for _, tt := range tests {
	    t.Run(tt.name, func(t *testing.T) {
	        err := ValidateCollection(tt.collection)
	        if (err != nil && err.Error() != tt.expected) || (err == nil && tt.expected != "") {
	            t.Errorf("expected error %q, got %q", tt.expected, err)
	        }
	    })
    }
}

func TestValidateDocument(t *testing.T) {
	tests := []struct {
	    name     string
	    doc      Document
	    expected string // Expected error message or empty for no errors expected
    }{
        {
            name:"valid document",
            doc : Document{Field:"name", Type:"String", Required:true},
            expected:"",
        },
        {
            name:"missing field name",
            doc : Document{Type:"String"},
            expected:"document field and type are required",
        },
        {
            name:"missing type",
            doc : Document{Field:"name"},
            expected:"document field and type are required",
        },
    }

	for _, tt := range tests {
	    t.Run(tt.name, func(t *testing.T) {
	        err := ValidateDocument(tt.doc)
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
	    expected string // Expected error message or empty for no errors expected
    }{
        {
            name:"valid foreign key",
            fk : ForeignKey{Name:"fk_customer_order", Field:"customer_id", References : Reference{Collection:"customers", Field:"customer_id"}},
            expected:"",
        },
        {
            name:"missing foreign key name",
            fk : ForeignKey{Field:"customer_id", References : Reference{Collection:"customers", Field:"customer_id"}},
            expected:"foreign key attributes must not be empty",
        },
        {
            name:"missing foreign key field",
            fk : ForeignKey{Name:"fk_customer_order", References : Reference{Collection:"customers", Field:"customer_id"}},
            expected:"foreign key attributes must not be empty",
        },
        {
            name:"missing referenced collection",
            fk : ForeignKey{Name:"fk_customer_order", Field:"customer_id", References : Reference{Field:"customer_id"}},
            expected:"foreign key attributes must not be empty",
        },
        {
            name:"missing referenced field",
            fk : ForeignKey{Name:"fk_customer_order", Field:"customer_id", References : Reference{Collection:"customers"}},
            expected:"foreign key attributes must not be empty",
        },
    }

	for _, tt := range tests {
	    t.Run(tt.name, func(t *testing.T) {
	        err := ValidateForeignKey(tt.fk)
	        if (err != nil && err.Error() != tt.expected) || (err == nil && tt.expected != "") {
	            t.Errorf("expected error %q, got %q", tt.expected, err)
	        }
	    })
    }
}