package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var catFileCmd = &cobra.Command{
	Use:   "cat-file",
	Short: "cat-fileialize git repo",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("git-go cat-file run")
		return nil
	},
}
