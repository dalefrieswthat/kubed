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

// =============================================================================
// Knowledge Cache schema (.kubed/learned.json)
// Accumulated knowledge from agent exploration, persists across sessions.
// =============================================================================

// LearnedCache is the root schema for .kubed/learned.json.
// Agents append to this as they discover facts; future queries read it first.
type LearnedCache struct {
	Version   string         `json:"version"`    // "v1"
	UpdatedAt string         `json:"updated_at"` // RFC3339
	RepoRoot  string         `json:"repo_root,omitempty"`
	Facts     []LearnedFact  `json:"facts"`      // Discovered facts
	Paths     []LearnedPath  `json:"paths"`      // Important paths with descriptions
	Deps      []LearnedDep   `json:"deps"`       // Dependencies/tech stack
	Patterns  []LearnedPattern `json:"patterns"` // Code patterns and conventions
}

// LearnedFact is a single piece of discovered knowledge.
type LearnedFact struct {
	ID        string   `json:"id"`                  // unique, e.g. "fact-001"
	Fact      string   `json:"fact"`                // the knowledge, e.g. "API uses PostgreSQL for persistence"
	Category  string   `json:"category,omitempty"`  // e.g. "architecture", "config", "deployment"
	Tags      []string `json:"tags,omitempty"`      // for filtering
	Source    string   `json:"source,omitempty"`    // where this was learned, e.g. "docker-compose.yml"
	LearnedAt string   `json:"learned_at"`          // RFC3339
}

// LearnedPath is an important file/directory with a description.
type LearnedPath struct {
	Path        string   `json:"path"`
	Description string   `json:"description"`        // what this path is for
	Tags        []string `json:"tags,omitempty"`
	LearnedAt   string   `json:"learned_at"`
}

// LearnedDep is a discovered dependency or tech stack component.
type LearnedDep struct {
	Name      string `json:"name"`               // e.g. "PostgreSQL", "Redis", "React"
	Kind      string `json:"kind"`               // "database", "cache", "framework", "library", "service"
	Version   string `json:"version,omitempty"`
	UsedBy    string `json:"used_by,omitempty"`  // which service/component uses it
	Source    string `json:"source,omitempty"`   // where discovered
	LearnedAt string `json:"learned_at"`
}

// LearnedPattern is a code convention or pattern.
type LearnedPattern struct {
	Name        string `json:"name"`        // e.g. "test file location"
	Pattern     string `json:"pattern"`     // e.g. "tests are in __tests__/ folders"
	Example     string `json:"example,omitempty"`
	LearnedAt   string `json:"learned_at"`
}
