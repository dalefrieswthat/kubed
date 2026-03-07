package layout

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// LearnedPath returns the path to .kubed/learned.json.
// Uses same logic as LayoutPath: repo root if in git, else ~/.kubed/.
func LearnedFilePath() (string, error) {
	root, err := RepoRoot()
	if err != nil {
		return "", err
	}
	if root != "" {
		return filepath.Join(root, ".kubed", "learned.json"), nil
	}
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, ".kubed", "learned.json"), nil
}

// LoadLearned reads .kubed/learned.json. Returns empty cache if file doesn't exist.
func LoadLearned() (*LearnedCache, error) {
	path, err := LearnedFilePath()
	if err != nil {
		return nil, err
	}
	data, err := os.ReadFile(path)
	if os.IsNotExist(err) {
		root, _ := RepoRoot()
		return &LearnedCache{
			Version:  "v1",
			RepoRoot: root,
			Facts:    []LearnedFact{},
			Paths:    []LearnedPath{},
			Deps:     []LearnedDep{},
			Patterns: []LearnedPattern{},
		}, nil
	}
	if err != nil {
		return nil, err
	}
	var cache LearnedCache
	if err := json.Unmarshal(data, &cache); err != nil {
		return nil, err
	}
	return &cache, nil
}

// SaveLearned writes the cache to .kubed/learned.json.
func SaveLearned(cache *LearnedCache) error {
	path, err := LearnedFilePath()
	if err != nil {
		return err
	}
	cache.UpdatedAt = time.Now().UTC().Format(time.RFC3339)
	if cache.Version == "" {
		cache.Version = "v1"
	}
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("create dir %s: %w", dir, err)
	}
	data, err := json.MarshalIndent(cache, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0644)
}

// AddFact appends a fact to the cache (dedupes by fact text).
func (c *LearnedCache) AddFact(fact, category, source string, tags []string) {
	// Check for duplicate
	for _, f := range c.Facts {
		if f.Fact == fact {
			return
		}
	}
	id := fmt.Sprintf("fact-%03d", len(c.Facts)+1)
	c.Facts = append(c.Facts, LearnedFact{
		ID:        id,
		Fact:      fact,
		Category:  category,
		Tags:      tags,
		Source:    source,
		LearnedAt: time.Now().UTC().Format(time.RFC3339),
	})
}

// AddPath adds an important path with description (dedupes by path).
func (c *LearnedCache) AddPath(path, description string, tags []string) {
	for i, p := range c.Paths {
		if p.Path == path {
			// Update description if already exists
			c.Paths[i].Description = description
			c.Paths[i].Tags = tags
			return
		}
	}
	c.Paths = append(c.Paths, LearnedPath{
		Path:        path,
		Description: description,
		Tags:        tags,
		LearnedAt:   time.Now().UTC().Format(time.RFC3339),
	})
}

// AddDep adds a dependency (dedupes by name+kind).
func (c *LearnedCache) AddDep(name, kind, version, usedBy, source string) {
	for i, d := range c.Deps {
		if d.Name == name && d.Kind == kind {
			// Update if exists
			c.Deps[i].Version = version
			c.Deps[i].UsedBy = usedBy
			return
		}
	}
	c.Deps = append(c.Deps, LearnedDep{
		Name:      name,
		Kind:      kind,
		Version:   version,
		UsedBy:    usedBy,
		Source:    source,
		LearnedAt: time.Now().UTC().Format(time.RFC3339),
	})
}

// AddPattern adds a code pattern (dedupes by name).
func (c *LearnedCache) AddPattern(name, pattern, example string) {
	for i, p := range c.Patterns {
		if p.Name == name {
			c.Patterns[i].Pattern = pattern
			c.Patterns[i].Example = example
			return
		}
	}
	c.Patterns = append(c.Patterns, LearnedPattern{
		Name:      name,
		Pattern:   pattern,
		Example:   example,
		LearnedAt: time.Now().UTC().Format(time.RFC3339),
	})
}

// Summary returns a short summary of the cache for quick loading.
func (c *LearnedCache) Summary() string {
	return fmt.Sprintf("%d facts, %d paths, %d deps, %d patterns",
		len(c.Facts), len(c.Paths), len(c.Deps), len(c.Patterns))
}
