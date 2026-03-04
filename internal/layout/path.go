package layout

import (
	"os"
	"path/filepath"
)

// LayoutPath returns the path to write/read layout.json.
// If cwd is inside a git repo, returns ".kubed/layout.json" (relative to cwd).
// Otherwise returns "~/.kubed/layout.json" (expanded).
func LayoutPath() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	if isGitRepo(dir) {
		return filepath.Join(dir, ".kubed", "layout.json"), nil
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
