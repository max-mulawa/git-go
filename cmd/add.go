package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "addialize git repo",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("git-go add run")
		return nil
	},
}
