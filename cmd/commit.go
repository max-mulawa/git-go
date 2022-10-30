package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var commitCmd = &cobra.Command{
	Use:   "commit",
	Short: "commitialize git repo",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("git-go commit run")
		return nil
	},
}
