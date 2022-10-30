package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var tagCmd = &cobra.Command{
	Use:   "tag",
	Short: "tagialize git repo",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("git-go tag run")
		return nil
	},
}
