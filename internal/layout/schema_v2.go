package layout

// IndexV2 is the root schema for .kubed/layout.json (v2).
// Sections are indexable by id or tags so agents/LLMs can request only relevant parts.
type IndexV2 struct {
	Version     string    `json:"version"` // "v2"
	GeneratedAt string    `json:"generated_at"`
	RepoRoot    string    `json:"repo_root,omitempty"`
	Sections    []Section `json:"sections"`
}

// Section is a named, taggable block of data for token-efficient lookup.
// Agents can request e.g. "section id=infra_paths" or "sections with tag=terraform".
type Section struct {
	ID      string      `json:"id"`
	Tags    []string    `json:"tags"`
	Summary string      `json:"summary"`
	Payload interface{} `json:"payload"`
}

// K8sLayoutPayload is the payload for section id "k8s_layout" (optional; only when kube context exists).
type K8sLayoutPayload struct {
	Context      string         `json:"context"`
	Namespaces   []string       `json:"namespaces"`
	Resources    []Resource     `json:"resources"`
	Relationships []Relationship `json:"relationships"`
}

// InfraPathEntry describes one infra-related file or directory.
type InfraPathEntry struct {
	Path string `json:"path"`
	Type string `json:"type"` // e.g. "terraform", "docker", "helm", "dockerfile", "compose"
}

// InfraPathsPayload is the payload for section id "infra_paths".
type InfraPathsPayload struct {
	Paths []InfraPathEntry `json:"paths"`
}

// ProjectEntry is one entry in the project structure tree.
type ProjectEntry struct {
	Path     string   `json:"path"`
	Kind     string   `json:"kind"` // "dir" or "file"
	Children []string `json:"children,omitempty"`
}

// ProjectStructurePayload is the payload for section id "project_structure".
type ProjectStructurePayload struct {
	Entries []ProjectEntry `json:"entries"`
}

// ContextSectionPayload is the payload for context sections (e.g. id "overview", "goals").
type ContextSectionPayload struct {
	Content string `json:"content"`
	Source  string `json:"source,omitempty"` // path to file
}

// SharedInfraRef describes a shared infra directory found in sibling/parent paths.
type SharedInfraRef struct {
	Path     string   `json:"path"`      // absolute or relative path
	Type     string   `json:"type"`      // "terraform", "helm", "docker", etc.
	Relation string   `json:"relation"`  // "sibling", "parent", "workspace"
	Patterns []string `json:"patterns"`  // files/dirs found (e.g. ["main.tf", "modules/"])
}

// SharedInfraPayload is the payload for section id "shared_infra".
type SharedInfraPayload struct {
	Refs []SharedInfraRef `json:"refs"`
}
