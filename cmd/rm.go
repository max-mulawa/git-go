package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "rmialize git repo",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("git-go rm run")
		return nil
	},
}
