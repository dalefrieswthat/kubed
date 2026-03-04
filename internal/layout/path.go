package layout

import (
	"os"
	"path/filepath"
)

// RepoRoot returns the repository root directory (where .git is), or empty string if not in a git repo.
func RepoRoot() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	d := dir
	for {
		gitDir := filepath.Join(d, ".git")
		info, err := os.Stat(gitDir)
		if err == nil && info.IsDir() {
			return d, nil
		}
		parent := filepath.Dir(d)
		if parent == d {
			break
		}
		d = parent
	}
	return "", nil
}

// LayoutPath returns the path to write/read layout.json.
// If cwd is inside a git repo, returns "<repo_root>/.kubed/layout.json".
// Otherwise returns "~/.kubed/layout.json" (expanded).
func LayoutPath() (string, error) {
	root, err := RepoRoot()
	if err != nil {
		return "", err
	}
	if root != "" {
		return filepath.Join(root, ".kubed", "layout.json"), nil
	}
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, ".kubed", "layout.json"), nil
}

func isGitRepo(dir string) bool {
	d := dir
	for {
		gitDir := filepath.Join(d, ".git")
		info, err := os.Stat(gitDir)
		if err == nil && info.IsDir() {
			return true
		}
		parent := filepath.Dir(d)
		if parent == d {
			break
		}
		d = parent
	}
	return false
}
