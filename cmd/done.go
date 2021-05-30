package cmd

import (
	"fmt"
	"log"

	"github.com/hauntarl/task-manager/db"
	"github.com/spf13/cobra"
)

// rootCmd gives basic information about app's command-line flags
var doneCmd = &cobra.Command{
	Use:   "done",
	Short: "List all completed tasks",
	Run: func(_ *cobra.Command, args []string) {
		tasks, err := db.ReadTasks(db.Bdone)
		if err != nil {
			log.Fatal(err)
		}
		if len(tasks) == 0 {
			fmt.Println("You haven't completed any tasks yet, why not do them instead? 🙃")
			return
		}

		fmt.Println("Completed tasks...")
		for i, task := range tasks {
			fmt.Printf("%d. %s\n", i+1, task.Val)
		}
	},
}

func init() { rootCmd.AddCommand(doneCmd) }
