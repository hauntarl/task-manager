package cmd

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hauntarl/task-manager/db"
	"github.com/spf13/cobra"
)

// doCmd marks a task as complete
var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Marks a task as completed.",
	Run: func(_ *cobra.Command, args []string) {
		// fetch tasks to retrieve internal ids associated with each one
		tasks, err := db.ReadTasks()
		if err != nil {
			log.Fatal(err)
		}
		if len(tasks) == 0 {
			fmt.Println("You have no tasks to complete, why not take a vacation? ⛵")
			return
		}

		ids := make(map[int]struct{}, len(tasks)) // remove duplicate entries
		for _, arg := range args {
			if id, err := strconv.Atoi(arg); err != nil {
				fmt.Printf("'%s' is not a valid ID\n", arg)
			} else if id > len(tasks) || id < 1 {
				fmt.Printf("id: '%d' out of bounds: '1-%d'\n", id, len(tasks))
			} else {
				ids[id-1] = struct{}{} // convert to zero-based indexing
			}
		}

		for id := range ids {
			val := tasks[id].Val
			if err := db.DeleteTask(tasks[id].Key); err != nil {
				fmt.Printf("Failed to delete task: '%s'\n", val)
				continue
			}
			fmt.Printf("You have completed the '%s' task.\n", val)
		}
	},
}

// register do command at our root to create a new cmd-line flag
func init() { rootCmd.AddCommand(doCmd) }
