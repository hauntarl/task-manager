package cmd

import (
	"fmt"
	"log"

	"github.com/hauntarl/task-manager/db"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
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

func init() {
	rootCmd.AddCommand(listCmd)
}
