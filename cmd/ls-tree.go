package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var lsTreeCmd = &cobra.Command{
	Use:   "ls-tree",
	Short: "ls-tree",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("git-go ls-tree run")
		return nil
	},
}
