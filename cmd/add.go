package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/hauntarl/task-manager/db"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a task to your task list.",
	Run: func(_ *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		_, err := db.AddTask(task)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Added '%s' to your task list.\n", task)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
