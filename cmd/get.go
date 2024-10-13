package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	dbSize     string
	numRows    int
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get is used to retrieve data from a fusion",
	Long: `The 'get' command retrieves data from a specified database with the given configuration.
You can specify the number of rows or size of the data to be retrieved.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Proceed with the data retrieval logic
		fmt.Printf("Getting data from the following details:\n")
		fmt.Printf("Database Name: %s\n", dbName)
		fmt.Printf("Table Name: %s\n", tableName)
		fmt.Printf("Database User: %s\n", dbUser)

		if dbSize != "" {
			fmt.Printf("Requested Size: %s\n", dbSize)
		}
		if numRows > 0 {
			fmt.Printf("Requested Rows: %d\n", numRows)
		}

		// Add your logic to retrieve the data here
		// Example: Fetch rows from a database or retrieve data by size
		if dbSize != "" {
			fmt.Printf("Retrieving up to %s of data...\n", dbSize)
		}
		if numRows > 0 {
			fmt.Printf("Retrieving %d rows of data...\n", numRows)
		}
	},
}

func init() {
	rootCmd.AddCommand(getCmd)

	// Mandatory flags for database configuration
	getCmd.Flags().StringVar(&dbName, "dbname", "", "Specify the name of the database (required)")
	getCmd.Flags().StringVar(&tableName, "tablename", "", "Specify the table name in the database (required)")
	getCmd.Flags().StringVar(&dbUser, "user", "", "Specify the database username (required)")

	// Optional flags for size and number of rows
	getCmd.Flags().StringVar(&dbSize, "size", "", "Specify the size of data to retrieve (e.g., 20MB)")
	getCmd.Flags().IntVarP(&numRows, "number", "n", 0, "Specify the number of rows to retrieve")

	// Mark dbname, tablename, and user as required
	getCmd.MarkFlagRequired("dbname")
	getCmd.MarkFlagRequired("tablename")
	getCmd.MarkFlagRequired("user")
}
