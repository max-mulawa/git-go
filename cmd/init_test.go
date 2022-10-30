package cmd_test

import (
	"os"
	"testing"

	"github.com/max-mulawa/git-go/git"
	"github.com/stretchr/testify/require"
)

func TestInit(t *testing.T) {
	repoTempDir, err := os.MkdirTemp(os.TempDir(), "repos")
	require.NoError(t, err)
	defer os.RemoveAll(repoTempDir)

	repo, err := git.NewRepo(repoTempDir, true)
	require.NoError(t, err)

	err = repo.Create()
	require.NoError(t, err)
}
