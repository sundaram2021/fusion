package mongodb

// Database represents the overall database configuration.
type Database struct {
    Name      string     `json:"name"`
    Engine    string     `json:"engine"`
    Collections []Collection `json:"collections"`
    Options   Options     `json:"options,omitempty"`
}

// Collection represents a collection within the database.
type Collection struct {
    Name      string       `json:"name"`
    Documents []Document   `json:"documents"`
    Indexes   []Index      `json:"indexes,omitempty"`
    ForeignKeys []ForeignKey  `json:"foreign_keys,omitempty"`
}

// Document represents a document structure in a collection.
type Document struct {
    Field  string      `json:"field"`
    Type   string      `json:"type"`
    Required bool      `json:"required,omitempty"`
    Default interface{} `json:"default,omitempty"` // Can be any type
}

// Index represents an index on a collection.
type Index struct {
    Name   string   `json:"name"`
    Fields []string `json:"fields"`
    Unique bool     `json:"unique,omitempty"`
}

// ForeignKey represents a foreign key constraint.
type ForeignKey struct {
    Name       string   `json:"name"`
    Field      string   `json:"field"`
    References Reference `json:"references"`
    OnDelete   string   `json:"on_delete,omitempty"`
    OnUpdate   string   `json:"on_update,omitempty"`
}

// Reference represents the referenced collection and field for foreign keys.
type Reference struct {
    Collection string `json:"collection"`
    Field      string `json:"field"`
}

// Options contains global database options.
type Options struct {
    MaxConnections     int    `json:"max_connections,omitempty"` // Maximum allowed connections
    ReplicaSetName     string `json:"replica_set_name,omitempty"` // Name of the replica set if applicable (optional)
    WriteConcernLevel  string `json:"write_concern_level,omitempty"` // Write concern level (optional)
    ReadConcernLevel   string `json:"read_concern_level,omitempty"` // Read concern level (optional)
    AuthEnabled        bool   `json:"auth_enabled,omitempty"` // Authentication setting (optional)
    SSLEnabled         bool   `json:"ssl_enabled,omitempty"`  // SSL setting (optional)
}