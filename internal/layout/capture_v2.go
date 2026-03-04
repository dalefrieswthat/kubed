package layout

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"gopkg.in/yaml.v3"
)

// contextYAML is the structure of cursor-context/context.yaml (optional).
type contextYAML struct {
	Files []string `yaml:"files"`
}

// RunCaptureV2 builds a v2 index with sections: infra_paths, project_structure, context sections, and optionally k8s_layout.
func RunCaptureV2(allNamespaces bool) (*IndexV2, error) {
	root, err := RepoRoot()
	if err != nil {
		return nil, err
	}
	repoRoot := root
	if repoRoot == "" {
		repoRoot, _ = os.Getwd()
	}

	idx := &IndexV2{
		Version:     "v2",
		GeneratedAt: time.Now().UTC().Format(time.RFC3339),
		RepoRoot:    repoRoot,
		Sections:    nil,
	}

	// 1) Infra paths (always)
	infra := buildInfraPaths(repoRoot)
	idx.Sections = append(idx.Sections, Section{
		ID:      "infra_paths",
		Tags:    []string{"infra", "terraform", "docker", "helm"},
		Summary: "Infra-related files and directories",
		Payload: infra,
	})

	// 2) Project structure (always)
	proj := buildProjectStructure(repoRoot)
	idx.Sections = append(idx.Sections, Section{
		ID:      "project_structure",
		Tags:    []string{"structure", "tree", "architecture"},
		Summary: "Project directory structure",
		Payload: proj,
	})

	// 3) Context sections from cursor-context/ (if present)
	ctxSections := buildContextSections(repoRoot)
	for _, s := range ctxSections {
		idx.Sections = append(idx.Sections, s)
	}

	// 4) Shared infra (sibling/parent directories with infra patterns)
	sharedInfra := buildSharedInfra(repoRoot)
	if len(sharedInfra.Refs) > 0 {
		idx.Sections = append(idx.Sections, Section{
			ID:      "shared_infra",
			Tags:    []string{"infra", "shared", "cross-repo"},
			Summary: "Shared infrastructure in sibling/parent directories",
			Payload: sharedInfra,
		})
	}

	// 5) K8s layout (optional; only if kube config works)
	k8sPayload, err := tryCaptureK8s(allNamespaces)
	if err == nil && k8sPayload != nil {
		idx.Sections = append(idx.Sections, Section{
			ID:      "k8s_layout",
			Tags:    []string{"kubernetes", "layout", "kubectl"},
			Summary: "Kubernetes resources and relationships",
			Payload: k8sPayload,
		})
	}
	// If k8s fails (no config, etc.), we simply omit the section

	return idx, nil
}

// buildInfraPaths scans repo for infra-related paths.
func buildInfraPaths(root string) InfraPathsPayload {
	var paths []InfraPathEntry
	exclude := map[string]bool{
		".git": true, "node_modules": true, "venv": true, ".venv": true,
		"build": true, "dist": true, "_site": true,
		"test_env": true, "build_env": true,
		"site-packages": true, "__pycache__": true,
	}

	// Top-level and one level down for common patterns
	walk := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}
		rel, _ := filepath.Rel(root, path)
		if rel == "." {
			return nil
		}
		parts := strings.Split(filepath.ToSlash(rel), "/")
		// Exclude if any path component is in the exclude set
		for _, part := range parts {
			if exclude[part] {
				if info.IsDir() {
					return filepath.SkipDir
				}
				return nil
			}
		}

		name := info.Name()
		if info.IsDir() {
			switch name {
			case "terraform", "tf":
				paths = append(paths, InfraPathEntry{Path: rel, Type: "terraform"})
				return filepath.SkipDir
			case "helm", "charts":
				paths = append(paths, InfraPathEntry{Path: rel, Type: "helm"})
				return filepath.SkipDir
			case "docker":
				paths = append(paths, InfraPathEntry{Path: rel, Type: "docker"})
				return filepath.SkipDir
			case ".kubed":
				paths = append(paths, InfraPathEntry{Path: rel, Type: "kubed"})
				return filepath.SkipDir
			}
			return nil
		}

		switch {
		case name == "Dockerfile" || strings.HasPrefix(name, "Dockerfile."):
			paths = append(paths, InfraPathEntry{Path: rel, Type: "dockerfile"})
		case name == "docker-compose.yml" || name == "docker-compose.yaml" || strings.HasPrefix(name, "docker-compose."):
			paths = append(paths, InfraPathEntry{Path: rel, Type: "compose"})
		case strings.HasSuffix(name, ".tf"):
			paths = append(paths, InfraPathEntry{Path: rel, Type: "terraform"})
		case name == "Chart.yaml":
			dir := filepath.Dir(rel)
			paths = append(paths, InfraPathEntry{Path: dir, Type: "helm"})
		}
		return nil
	}

	filepath.Walk(root, walk)

	// Dedupe by path and drop noise (symlinks like ~, merge backups)
	seen := make(map[string]bool)
	var out []InfraPathEntry
	for _, p := range paths {
		norm := filepath.Clean(p.Path)
		if seen[norm] || strings.HasPrefix(norm, "~") || strings.Contains(norm, "~") {
			continue
		}
		seen[norm] = true
		out = append(out, InfraPathEntry{Path: norm, Type: p.Type})
	}

	return InfraPathsPayload{Paths: out}
}

// buildProjectStructure returns top-level dirs and key files.
func buildProjectStructure(root string) ProjectStructurePayload {
	var entries []ProjectEntry
	exclude := map[string]bool{
		".git": true, "node_modules": true, "venv": true, ".venv": true,
		"build": true, "dist": true, "_site": true, ".kubed": true,
		"test_env": true, "build_env": true, ".cursor": true,
		"kubed.egg-info": true, ".DS_Store": true,
	}

	entriesMap := make(map[string]*ProjectEntry)

	f, err := os.Open(root)
	if err != nil {
		return ProjectStructurePayload{}
	}
	names, _ := f.Readdirnames(-1)
	f.Close()

	for _, name := range names {
		if exclude[name] {
			continue
		}
		if strings.HasPrefix(name, "terraform~") {
			continue
		}
		full := filepath.Join(root, name)
		info, err := os.Stat(full)
		if err != nil {
			continue
		}
		if info.IsDir() {
			entriesMap[name+"/"] = &ProjectEntry{Path: name + "/", Kind: "dir"}
		} else {
			entriesMap[name] = &ProjectEntry{Path: name, Kind: "file"}
		}
	}

	for _, e := range entriesMap {
		entries = append(entries, *e)
	}
	return ProjectStructurePayload{Entries: entries}
}

// buildSharedInfra looks for shared infra directories in sibling and parent paths.
// Patterns: *-infra, infra-*, shared-*, *-terraform, *-helm
func buildSharedInfra(repoRoot string) SharedInfraPayload {
	var refs []SharedInfraRef
	repoName := filepath.Base(repoRoot)
	parentDir := filepath.Dir(repoRoot)

	// Patterns that indicate shared infra
	infraPatterns := []string{
		"-infra", "infra-", "-terraform", "-helm", "-k8s",
		"shared-", "common-", "-shared", "-common",
	}

	isInfraDir := func(name string) bool {
		lower := strings.ToLower(name)
		for _, pat := range infraPatterns {
			if strings.Contains(lower, pat) {
				return true
			}
		}
		return false
	}

	// Scan for infra patterns in a directory, return list of key files/dirs
	scanInfraPatterns := func(dir string) ([]string, string) {
		var patterns []string
		var infraType string
		entries, err := os.ReadDir(dir)
		if err != nil {
			return nil, ""
		}
		for _, e := range entries {
			name := e.Name()
			switch {
			case name == "main.tf" || name == "variables.tf" || name == "outputs.tf":
				patterns = append(patterns, name)
				infraType = "terraform"
			case strings.HasSuffix(name, ".tf"):
				patterns = append(patterns, name)
				infraType = "terraform"
			case name == "Chart.yaml":
				patterns = append(patterns, name)
				infraType = "helm"
			case name == "Dockerfile" || strings.HasPrefix(name, "Dockerfile."):
				patterns = append(patterns, name)
				if infraType == "" {
					infraType = "docker"
				}
			case name == "docker-compose.yml" || name == "docker-compose.yaml":
				patterns = append(patterns, name)
				if infraType == "" {
					infraType = "compose"
				}
			case strings.HasSuffix(name, ".yaml") || strings.HasSuffix(name, ".yml") || strings.HasSuffix(name, ".json"):
				// CloudFormation templates
				if strings.Contains(strings.ToLower(dir), "cloudformation") {
					patterns = append(patterns, name)
					if infraType == "" {
						infraType = "cloudformation"
					}
				}
			case e.IsDir() && (name == "terraform" || name == "helm" || name == "modules" || name == "cloudformation"):
				patterns = append(patterns, name+"/")
				if infraType == "" {
					switch name {
					case "terraform", "modules":
						infraType = "terraform"
					case "cloudformation":
						infraType = "cloudformation"
					default:
						infraType = name
					}
				}
			}
		}
		if infraType == "" && len(patterns) > 0 {
			infraType = "mixed"
		}
		return patterns, infraType
	}

	// Check sibling directories
	siblings, err := os.ReadDir(parentDir)
	if err == nil {
		for _, sib := range siblings {
			if !sib.IsDir() {
				continue
			}
			name := sib.Name()
			if name == repoName || name == ".git" {
				continue
			}
			if isInfraDir(name) {
				sibPath := filepath.Join(parentDir, name)
				patterns, infraType := scanInfraPatterns(sibPath)
				if len(patterns) > 0 {
					refs = append(refs, SharedInfraRef{
						Path:     sibPath,
						Type:     infraType,
						Relation: "sibling",
						Patterns: patterns,
					})
				}
			}
		}
	}

	// Check for infra subdirectories in sibling repos (e.g. sibling/terraform/, sibling/helm/)
	infraDirNames := map[string]string{
		"terraform":      "terraform",
		"helm":           "helm",
		"charts":         "helm",
		"infra":          "mixed",
		"cloudformation": "cloudformation",
		"cdk":            "cdk",
		"pulumi":         "pulumi",
	}
	if err == nil {
		for _, entry := range siblings {
			if !entry.IsDir() {
				continue
			}
			name := entry.Name()
			if name == repoName || name == ".git" {
				continue
			}
			entryPath := filepath.Join(parentDir, name)
			subEntries, err := os.ReadDir(entryPath)
			if err != nil {
				continue
			}
			for _, sub := range subEntries {
				if !sub.IsDir() {
					continue
				}
				subName := strings.ToLower(sub.Name())
				if infraType, ok := infraDirNames[subName]; ok {
					subPath := filepath.Join(entryPath, sub.Name())
					patterns, detectedType := scanInfraPatterns(subPath)
					if detectedType == "" {
						detectedType = infraType
					}
					if len(patterns) > 0 || infraType != "" {
						if len(patterns) == 0 {
							patterns = []string{sub.Name() + "/"}
						}
						refs = append(refs, SharedInfraRef{
							Path:     subPath,
							Type:     detectedType,
							Relation: "sibling",
							Patterns: patterns,
						})
					}
				} else if isInfraDir(sub.Name()) {
					subPath := filepath.Join(entryPath, sub.Name())
					patterns, infraType := scanInfraPatterns(subPath)
					if len(patterns) > 0 {
						refs = append(refs, SharedInfraRef{
							Path:     subPath,
							Type:     infraType,
							Relation: "sibling",
							Patterns: patterns,
						})
					}
				}
			}
		}
	}

	return SharedInfraPayload{Refs: refs}
}

// buildContextSections reads cursor-context/context.yaml and listed files, splitting by ## headings.
func buildContextSections(root string) []Section {
	ctxDir := filepath.Join(root, "cursor-context")
	contextYamlPath := filepath.Join(ctxDir, "context.yaml")
	data, err := os.ReadFile(contextYamlPath)
	if err != nil {
		return nil
	}

	var cfg contextYAML
	if err := yaml.Unmarshal(data, &cfg); err != nil || len(cfg.Files) == 0 {
		return nil
	}

	var sections []Section
	sectionRe := regexp.MustCompile(`^##\s+(.+)$`)

	for _, file := range cfg.Files {
		fpath := filepath.Join(ctxDir, file)
		content, err := os.ReadFile(fpath)
		if err != nil {
			continue
		}
		scanner := bufio.NewScanner(strings.NewReader(string(content)))
		var currentID, currentContent strings.Builder
		currentTags := []string{"context", "architecture"}

		for scanner.Scan() {
			line := scanner.Text()
			if m := sectionRe.FindStringSubmatch(line); m != nil {
				if currentID.Len() > 0 {
					sections = append(sections, Section{
						ID:      strings.TrimSpace(currentID.String()),
						Tags:    currentTags,
						Summary: firstLine(currentContent.String(), 80),
						Payload: ContextSectionPayload{
							Content: strings.TrimSpace(currentContent.String()),
							Source:  filepath.Join("cursor-context", file),
						},
					})
				}
				currentID.Reset()
				currentContent.Reset()
				currentTags = []string{"context", "architecture"}
				// section id = first token after ## (lowercase, no spaces)
				id := strings.ToLower(strings.TrimSpace(m[1]))
				id = regexp.MustCompile(`\s+`).ReplaceAllString(id, "-")
				currentID.WriteString(id)
				continue
			}
			if currentID.Len() > 0 {
				currentContent.WriteString(line)
				currentContent.WriteString("\n")
			}
		}
		if currentID.Len() > 0 {
			sections = append(sections, Section{
				ID:      strings.TrimSpace(currentID.String()),
				Tags:    currentTags,
				Summary: firstLine(currentContent.String(), 80),
				Payload: ContextSectionPayload{
					Content: strings.TrimSpace(currentContent.String()),
					Source:  filepath.Join("cursor-context", file),
				},
			})
		}
	}
	return sections
}

func firstLine(s string, maxLen int) string {
	s = strings.TrimSpace(s)
	idx := strings.Index(s, "\n")
	if idx > 0 {
		s = s[:idx]
	}
	s = strings.TrimSpace(s)
	if len(s) > maxLen {
		return s[:maxLen] + "..."
	}
	return s
}

// tryCaptureK8s runs k8s capture; returns payload or nil on any error (e.g. no kubeconfig).
func tryCaptureK8s(allNamespaces bool) (*K8sLayoutPayload, error) {
	loadingRules := clientcmd.NewDefaultClientConfigLoadingRules()
	configOverrides := &clientcmd.ConfigOverrides{}
	kubeConfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(loadingRules, configOverrides)

	config, err := kubeConfig.ClientConfig()
	if err != nil {
		return nil, err
	}
	config.QPS = 50
	config.Burst = 100

	rawConfig, err := kubeConfig.RawConfig()
	if err != nil {
		return nil, err
	}
	contextName := rawConfig.CurrentContext

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	ctx := context.Background()
	layout, err := buildLayout(ctx, clientset, contextName, allNamespaces)
	if err != nil {
		return nil, err
	}

	return &K8sLayoutPayload{
		Context:      layout.Context,
		Namespaces:   layout.Namespaces,
		Resources:    layout.Resources,
		Relationships: layout.Relationships,
	}, nil
}

// WriteIndexV2 writes the v2 index to layout.json path.
func WriteIndexV2(idx *IndexV2) (string, error) {
	layoutPath, err := LayoutPath()
	if err != nil {
		return "", err
	}
	dir := filepath.Dir(layoutPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return "", fmt.Errorf("create dir %s: %w", dir, err)
	}
	out, err := json.MarshalIndent(idx, "", "  ")
	if err != nil {
		return "", err
	}
	if err := os.WriteFile(layoutPath, out, 0644); err != nil {
		return "", fmt.Errorf("write %s: %w", layoutPath, err)
	}
	return layoutPath, nil
}
