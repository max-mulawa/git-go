package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "git-go",
	// Short: "git-go is a git client written in go",
	// Long: `git-go is a git client with only basic functionality supported such us init, commit, add etc.
	// 		It's based on Python tutorial \"Write yourself a Git!\", check it on https://wyag.thb.lt/`,
	// Run: func(cmd *cobra.Command, args []string) {
	// 	fmt.Println("git-go v0.1")
	// },
	Args: cobra.MinimumNArgs(1),
}

func init() {
	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(catFileCmd)
	rootCmd.AddCommand(checkoutCmd)
	rootCmd.AddCommand(commitCmd)
	rootCmd.AddCommand(hashObjectCmd)
	rootCmd.AddCommand(logCmd)
	rootCmd.AddCommand(lsTreeCmd)
	rootCmd.AddCommand(mergeCmd)
	rootCmd.AddCommand(rebaseCmd)
	rootCmd.AddCommand(revParseCmd)
	rootCmd.AddCommand(rmCmd)
	rootCmd.AddCommand(showRefCmd)
	rootCmd.AddCommand(tagCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
