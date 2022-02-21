package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new word",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("added %s to your task list.\n")
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
