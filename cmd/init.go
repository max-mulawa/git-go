package cmd

import (
	"fmt"

	"github.com/max-mulawa/git-go/git"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "initialize git repo",
	Args:  cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("git-go init running")

		worktree := "."
		if len(args) >= 1 {
			worktree = args[0]
		}

		repo, err := git.NewRepo(worktree, true)
		if err != nil {
			return fmt.Errorf("repo object creation failed: %w", err)
		} else {
			fmt.Printf("repo struct: %+v", repo)
		}
		err = repo.Create()
		if err != nil {
			return fmt.Errorf("repo init failed: %w", err)
		}

		return nil
	},
}
