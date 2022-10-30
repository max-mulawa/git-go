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
	cfg      *ini.File
}

func NewRepo(path string, force bool) (*Repo, error) {
	worktree, err := filepath.Abs(path)
	if err != nil {
		return nil, fmt.Errorf("worktree directory %s path parsing failed: %w", worktree, err)
	}
	gitdir := filepath.Join(worktree, ".git")

	gitDirInfo, err := os.Stat(gitdir)
	if err == nil {
		if !gitDirInfo.IsDir() {
			return nil, fmt.Errorf("not a git repository %s", path)
		}
	} else {
		if !(os.IsNotExist(err) && force) {
			return nil, fmt.Errorf("not a git repository %s", path)
		}
	}

	repo := &Repo{
		worktree: worktree,
		gitdir:   gitdir,
	}

	skipCfgLoad := false
	cfPath, err := repo.RepoFile(false, []string{"config"})
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}
	cfFile, err := os.Stat(cfPath)
	if err != nil {
		if !(os.IsNotExist(err) && force) {
			return nil, fmt.Errorf("configuration file missing or not accessible")
		} else {
			skipCfgLoad = true
		}
	} else if cfFile.IsDir() {
		return nil, fmt.Errorf("%s is a directory", cfPath)
	}

	if !skipCfgLoad {
		cfg, err := ini.Load(cfPath)
		if err != nil {
			return nil, fmt.Errorf("failed to load configuration file %s: %w", cfPath, err)
		}
		repo.cfg = cfg
		section, err := cfg.GetSection("core")
		if err != nil {
			return nil, fmt.Errorf("malformed %s: %w", cfPath, err)
		}
		key, err := section.GetKey("repositoryformatversion")
		if err != nil {
			return nil, fmt.Errorf("malformed %s, couldn't find repositoryformatversion: %w", cfPath, err)
		}

		if !force {
			version, err := key.Int()
			if err != nil {
				return nil, fmt.Errorf("malformed %s, repositoryformatversion is not an integer: %w", cfPath, err)
			}

			if version != 0 {
				return nil, fmt.Errorf("unsupported repositoryformatversion %s", key.String())
			}
		}
	}

	return repo, nil
}

func FindRepo(dirPath string) (*Repo, error) {
	repoPath, err := filepath.Abs(dirPath)
	if err != nil {
		return nil, fmt.Errorf("failed find repo: %w", err)
	}

	gitExists := false
	for {
		if gitExists, err = gitDirExists(repoPath); err != nil {
			return nil, err
		}

		if gitExists {
			return NewRepo(repoPath, false)
		}

		parentPath := filepath.Dir(repoPath)
		if repoPath == "/" && parentPath == repoPath {
			break
		}
		repoPath = parentPath
	}

	return nil, fmt.Errorf("repo not found from %s up to the parents", dirPath)
}

func gitDirExists(dirPath string) (bool, error) {
	gitDir, err := os.Stat(filepath.Join(dirPath, ".git"))
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		} else {
			return false, fmt.Errorf("failed to check directory %s: %w", dirPath, err)
		}
	}

	if !gitDir.IsDir() {
		return false, fmt.Errorf("%s not a directory", dirPath)
	}

	return true, nil
}

func (r *Repo) Create() error {
	err := ensureDir(r.worktree)
	if err != nil {
		return err
	}

	err = ensureDir(r.gitdir)
	if err != nil {
		return err
	}

	if _, err := r.RepoDir(true, []string{"branches"}); err != nil {
		return err
	}
	if _, err := r.RepoDir(true, []string{"objects"}); err != nil {
		return err
	}
	if _, err := r.RepoDir(true, []string{"refs", "tags"}); err != nil {
		return err
	}
	if _, err := r.RepoDir(true, []string{"refs", "heads"}); err != nil {
		return err
	}

	cfgPath, err := r.RepoFile(false, []string{"config"})
	if err != nil {
		return err
	}
	err = saveDefaultCfg(cfgPath)
	if err != nil {
		return err
	}

	headPath, err := r.RepoFile(false, []string{"HEAD"})
	if err != nil {
		return err
	}
	os.WriteFile(headPath, []byte("ref: refs/heads/master\n"), os.ModePerm)

	descPath, err := r.RepoFile(false, []string{"description"})
	if err != nil {
		return err
	}
	os.WriteFile(descPath, []byte("Unnamed repository; edit this file 'description' to name the repository.\n"), os.ModePerm)

	return nil
}

func ensureDir(dirPath string) error {
	wt, err := os.Stat(dirPath)
	if err != nil {
		if os.IsNotExist(err) {
			if mkdirErr := os.Mkdir(dirPath, os.ModePerm); mkdirErr != nil {
				return fmt.Errorf("failed to create directory %s: %w", dirPath, mkdirErr)
			}
			wt, err = os.Stat(dirPath)
			if err != nil {
				return fmt.Errorf("failed to check directory %s: %w", dirPath, err)
			}
		} else {
			return fmt.Errorf("failed to check directory %s: %w", dirPath, err)
		}
	}

	if !wt.IsDir() {
		return fmt.Errorf("%s not a directory", dirPath)
	}

	return nil
}

func saveDefaultCfg(cfgPath string) error {
	cfg := ini.Empty()
	coreSection, err := cfg.NewSection("core")
	if err != nil {
		return fmt.Errorf("config core section creation failed: %w", err)
	}

	coreSection.Key("repositoryformatversion").SetValue("0")
	coreSection.Key("filemode").SetValue("false")
	coreSection.Key("bare").SetValue("false")
	cfg.SaveTo(cfgPath)
	return nil
}

func (r *Repo) RepoPath(path []string) string {
	dirPath := r.gitdir
	for _, p := range path {
		dirPath = filepath.Join(dirPath, p)
	}
	return dirPath
}

func (r *Repo) RepoFile(mkdir bool, path []string) (string, error) {
	_, err := r.RepoDir(mkdir, path[:(len(path)-1)])
	if err != nil {
		return "", fmt.Errorf("failed to read file: %w", err)

	}
	return r.RepoPath(path), nil
}

func (r *Repo) RepoDir(mkdir bool, path []string) (string, error) {
	// for i := 1; i <= len(path); i++ {
	// 	//or
	// 	dir := r.RepoPath(path[:i])
	// 	fsDir, err := os.Stat(dir)
	// 	if err == os.ErrNotExist {
	// 		if mkdir {
	// 			if mkdirErr := os.Mkdir(dir, os.ModePerm); mkdirErr != nil {
	// 				return "", fmt.Errorf("failed to create directory %s: %w", dir, mkdirErr)
	// 			}
	// 		}
	// 	} else if err == nil {
	// 		if !fsDir.IsDir() {
	// 			return "", fmt.Errorf("%s is not a directory", dir)
	// 		}
	// 	}
	// }
	dirPath := r.RepoPath(path[:])
	if mkdir {
		if mkdirErr := os.MkdirAll(dirPath, os.ModePerm); mkdirErr != nil {
			return "", fmt.Errorf("failed to create directory %s: %w", dirPath, mkdirErr)
		}
	}

	return dirPath, nil
}
