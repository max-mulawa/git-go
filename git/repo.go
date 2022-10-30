package git

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/ini.v1"
)

type Repo struct {
	worktree string
	gitdir   string
}

func NewRepo(path string, force bool) (*Repo, error) {
	worktree := path
	gitdir := filepath.Join(worktree, ".git")

	gitDirInfo, err := os.Stat(gitdir)
	if err == nil {
		if !gitDirInfo.IsDir() {
			return nil, fmt.Errorf("Not a git repository %s", path)
		}
	} else {
		if !(err == os.ErrNotExist && force) {
			return nil, fmt.Errorf("Not a git repository %s", path)
		}
	}

	repo := &Repo{
		worktree: worktree,
		gitdir:   gitdir,
	}

	cfPath := repo.RepoFile("config")
	cfFile, err := os.Stat(cfPath)
	if err != nil {
		if !(err == os.ErrNotExist && force) {
			return nil, fmt.Errorf("configuration file missing or not accessible")
		}
	} else if cfFile.IsDir() {
		return nil, fmt.Errorf("%s is a directory", cfPath)
	}

	iniFile, err := ini.Load(cfPath)
	if err != nil {
		return nil, fmt.Errorf("failed to load configuration file %s: %w", cfPath, err)
	}

	if !force {
		section, err := iniFile.GetSection("core")
		if err != nil {
			return nil, fmt.Errorf("malformed %s: %w", cfPath, err)
		}
		key, err := section.GetKey("repositoryformatversion")
		if err != nil {
			return nil, fmt.Errorf("malformed %s, couldn't find repositoryformatversion: %w", cfPath, err)
		}

		version, err := key.Int()
		if err != nil {
			return nil, fmt.Errorf("malformed %s, repositoryformatversion is not an integer: %w", cfPath, err)
		}

		if version != 0 {
			return nil, fmt.Errorf("unsupported repositoryformatversion %s", key.String())
		}
	}

	return repo, nil
}

func (r *Repo) RepoPath(path string) string {
	return filepath.Join(r.gitdir, path)
}

func (r *Repo) RepoFile(path string) {

}

func (r *Repo) EnsureFile(path ...string) string {
	//os.Dir
	p := r.RepoFile(path)
	return
}
