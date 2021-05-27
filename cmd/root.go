package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

// rootCmd gives basic information about app's command-line flags
var rootCmd = &cobra.Command{
	Use:   "task",
	Short: "Task is a CLI Task Manager",
}

// Execute accepts incoming commands and maps them appropriately
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
