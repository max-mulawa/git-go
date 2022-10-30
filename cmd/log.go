package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var logCmd = &cobra.Command{
	Use:   "log",
	Short: "logialize git repo",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("git-go log run")
		return nil
	},
}
