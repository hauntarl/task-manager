package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/hauntarl/task-manager/db"
	"github.com/spf13/cobra"
)

// addCmd creates a new task in database
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a task to your task list.",
	Run: func(_ *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		if _, err := db.AddTask(task); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Added '%s' to your task list.\n", task)
	},
}

// register add command at our root to create a new cmd-line flag
func init() { rootCmd.AddCommand(addCmd) }
