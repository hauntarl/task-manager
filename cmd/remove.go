package cmd

import (
	"fmt"
	"log"

	"github.com/hauntarl/task-manager/db"
	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove the task from list.",
	Run: func(_ *cobra.Command, args []string) {
		// fetch tasks to retrieve internal ids associated with each one
		tasks, err := db.ReadTasks(db.Btodo)
		if err != nil {
			log.Fatal(err)
		}
		if len(tasks) == 0 {
			fmt.Println("You have no tasks to remove, start adding some... 📜")
			return
		}

		for id := range parseIDs(args, len(tasks)) {
			val := tasks[id].Val
			if err := db.DeleteTask(tasks[id].Key); err != nil {
				fmt.Printf("Failed to delete the task: '%s'\n", val)
				continue
			}
			fmt.Printf("Task '%s' has been deleted successfully.\n", val)
		}
	},
}

func init() { rootCmd.AddCommand(removeCmd) }
