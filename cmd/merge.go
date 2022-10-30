package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var mergeCmd = &cobra.Command{
	Use:   "merge",
	Short: "mergeialize git repo",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("git-go merge run")
		return nil
	},
}
