package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// doCmd represents the do command
var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Marks a task as completed.",
	Run: func(cmd *cobra.Command, args []string) {
		ids := make([]int, 0, len(args))
		for _, arg := range args {
			if id, err := strconv.Atoi(arg); err != nil {
				fmt.Printf("'%s' is not a valid ID\n", arg)
			} else {
				ids = append(ids, id)
			}
		}
		fmt.Println(ids)
	},
}

func init() {
	rootCmd.AddCommand(doCmd)
}
