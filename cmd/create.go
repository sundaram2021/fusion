package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)



// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create is used to create fusions",
	Long: `The 'create' command creates fusions with the specified database configuration.
You can use this command to create a fusion with a specified database type and Docker image, along with other optional configurations.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Proceed with the fusion creation logic
		fmt.Printf("Creating a fusion with the following details:\n")
		fmt.Printf("Database Type: %s\n", dbType)
		fmt.Printf("Docker Image: %s\n", dbImage)
		if dbUser != "" {
			fmt.Printf("Database User: %s\n", dbUser)
		}
		if dbPassword != "" {
			fmt.Printf("Database Password: %s\n", dbPassword)
		}
		if dbHost != "" {
			fmt.Printf("Database Host: %s\n", dbHost)
		}
		if dbPort != "" {
			fmt.Printf("Database Port: %s\n", dbPort)
		}
		if dbName != "" {
			fmt.Printf("Database Name: %s\n", dbName)
		}
		if tableName != ""{
			fmt.Printf("Database Name: %s\n", dbName)
		}
		// Add your logic to create the fusion here
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	// Add the --db flag for specifying the database type
	createCmd.Flags().StringVar(&dbType, "db", "", "Specify the database type (e.g., postgres, mysql, sqlite)")

	// Add the --image flag for specifying the Docker image
	createCmd.Flags().StringVar(&dbImage, "image", "", "Specify the Docker image for the database")

	createCmd.MarkFlagRequired("db")
	createCmd.MarkFlagFilename("image")

	// Optional flags for database configuration
	createCmd.Flags().StringVar(&dbUser, "user", "", "Database username")
	createCmd.Flags().StringVar(&dbPassword, "password", "", "Database password")
	createCmd.Flags().StringVar(&dbHost, "host", "localhost", "Database host (default is localhost)")
	createCmd.Flags().StringVar(&dbPort, "port", "5432", "Database port (default is 5432 for PostgreSQL)")
	createCmd.Flags().StringVar(&dbName, "dbname", "", "Database name")
	createCmd.Flags().StringVar(&tableName, "tbname", "", "Table name")
}
