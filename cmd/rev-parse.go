package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var revParseCmd = &cobra.Command{
	Use:   "rev-parse",
	Short: "rev-parse",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("git-go rev-parse run")
		return nil
	},
}
