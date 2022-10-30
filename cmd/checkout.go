package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var checkoutCmd = &cobra.Command{
	Use:   "checkout",
	Short: "checkoutialize git repo",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("git-go checkout run")
		return nil
	},
}
