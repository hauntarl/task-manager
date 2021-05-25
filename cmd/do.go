package cmd

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hauntarl/task-manager/db"
	"github.com/spf13/cobra"
)

// doCmd represents the do command
var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Marks a task as completed.",
	Run: func(_ *cobra.Command, args []string) {
		tasks, err := db.ReadTasks()
		if err != nil {
			log.Fatal(err)
		}
		if len(tasks) == 0 {
			fmt.Println("You have no tasks to complete, why not take a vacation? ⛵")
			return
		}

		ids := make(map[int]struct{}, len(tasks))
		for _, arg := range args {
			if id, err := strconv.Atoi(arg); err != nil {
				fmt.Printf("'%s' is not a valid ID\n", arg)
			} else if id > len(tasks) || id < 1 {
				fmt.Printf("id: '%d' out of bounds: '1-%d'\n", id, len(tasks))
			} else {
				ids[id-1] = struct{}{}
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

func init() {
	rootCmd.AddCommand(doCmd)
}
