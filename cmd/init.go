package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "initialize git repo",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("git-go init run")
		return nil
	},
}
