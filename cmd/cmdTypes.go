package cmd

var (
	dbType      string // Flag for the database type (e.g., postgres, mysql, sqlite)
	dbImage     string // Flag for the Docker image
	dbUser      string // Flag for the database user
	dbPassword  string // Flag for the database password
	dbHost      string // Flag for the database host
	dbPort      string // Flag for the database port
	dbName      string // Flag for the database name
	filename 	string // yaml file configurations
	tableName  string
)