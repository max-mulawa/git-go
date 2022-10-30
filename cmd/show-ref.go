package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var showRefCmd = &cobra.Command{
	Use:   "show-ref",
	Short: "show-ref",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("git-go show-ref run")
		return nil
	},
}
