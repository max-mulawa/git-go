package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rebaseCmd = &cobra.Command{
	Use:   "rebase",
	Short: "rebaseialize git repo",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("git-go rebase run")
		return nil
	},
}
