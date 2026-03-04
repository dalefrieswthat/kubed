package layout

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLayoutPath_InGitRepo(t *testing.T) {
	// Run from repo root (kubed has .git); path should be .kubed/layout.json under cwd.
	orig, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	// Walk up to find a dir with .git (our repo root)
	dir := orig
	for {
		if _, err := os.Stat(filepath.Join(dir, ".git")); err == nil {
			break
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			t.Skip("not in a git repo")
			return
		}
		dir = parent
	}
	// Change to repo root and get path
	if err := os.Chdir(dir); err != nil {
		t.Fatal(err)
	}
	defer os.Chdir(orig)

	path, err := LayoutPath()
	if err != nil {
		t.Fatal(err)
	}
	want := filepath.Join(dir, ".kubed", "layout.json")
	if path != want {
		t.Errorf("LayoutPath() = %s, want %s", path, want)
	}
}

func TestLayoutPath_OutsideGitRepo(t *testing.T) {
	// Use a temp dir without .git; path should be ~/.kubed/layout.json
	tmp := t.TempDir()
	orig, _ := os.Getwd()
	os.Chdir(tmp)
	defer os.Chdir(orig)

	path, err := LayoutPath()
	if err != nil {
		t.Fatal(err)
	}
	home, _ := os.UserHomeDir()
	want := filepath.Join(home, ".kubed", "layout.json")
	if path != want {
		t.Errorf("LayoutPath() = %s, want %s", path, want)
	}
}

func TestIsGitRepo(t *testing.T) {
	tmp := t.TempDir()
	if isGitRepo(tmp) {
		t.Error("temp dir should not be a git repo")
	}
	if err := os.MkdirAll(filepath.Join(tmp, ".git"), 0755); err != nil {
		t.Fatal(err)
	}
	if !isGitRepo(tmp) {
		t.Error("dir with .git should be a git repo")
	}
	// Parent of tmp should not be considered (we only check tmp)
	sub := filepath.Join(tmp, "sub")
	os.MkdirAll(sub, 0755)
	if !isGitRepo(sub) {
		t.Error("subdir of git repo should be a git repo")
	}
}
