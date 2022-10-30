package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var hashObjectCmd = &cobra.Command{
	Use:   "hash-object",
	Short: "hash-object",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("git-go hashobject run")
		return nil
	},
}
