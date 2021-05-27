package cmd

import (
	"fmt"
	"log"

	"github.com/hauntarl/task-manager/db"
	"github.com/spf13/cobra"
)

// listCmd displays all the todos to the user
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all of your tasks.",
	Run: func(_ *cobra.Command, args []string) {
		tasks, err := db.ReadTasks()
		if err != nil {
			log.Fatal(err)
		}
		if len(tasks) == 0 {
			fmt.Println("You have no tasks to complete, why not take a vacation? ⛵")
			return
		}

		fmt.Println("Your tasks...")
		for i, task := range tasks {
			fmt.Printf("%d. %s\n", i+1, task.Val)
		}
	},
}

// register list command at our root to create a new cmd-line flag
func init() { rootCmd.AddCommand(listCmd) }
