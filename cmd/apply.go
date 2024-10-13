package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)


// applyCmd represents the apply command
var applyCmd = &cobra.Command{
	Use:   "apply",
	Short: "Apply a configuration from a YAML file",
	Long: `The 'apply' command allows you to apply a configuration from a YAML file. 
		For example:fusion apply -f config.yaml`,
	Run: func(cmd *cobra.Command, args []string) {
		// Check if the -f flag (filename) is provided
		if filename == "" {
			fmt.Println("Error: you must specify a YAML file using the -f flag")
			os.Exit(1) // Exit if no file is provided
		}

		// Check if the file exists
		if _, err := os.Stat(filename); os.IsNotExist(err) {
			fmt.Printf("Error: file '%s' does not exist\n", filename)
			os.Exit(1) // Exit if the file doesn't exist
		}

		// Proceed to process the YAML file
		fmt.Printf("Applying configuration from file: %s\n", filename)

		// Add your logic to process the YAML file here
	},
}

func init() {
	rootCmd.AddCommand(applyCmd)

	// Add the local flag `-f` for specifying the YAML file
	applyCmd.Flags().StringVarP(&filename, "file", "f", "", "Path to the YAML configuration file")

	// Mark the `-f` flag as required
	applyCmd.MarkFlagRequired("file")
}
