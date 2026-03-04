package layout

// Layout is the root schema for .kubed/layout.json (v1).
type Layout struct {
	Version      string         `json:"version"`
	GeneratedAt  string         `json:"generated_at"`
	Context      string         `json:"context"`
	Namespaces   []string       `json:"namespaces"`
	Resources    []Resource     `json:"resources"`
	Relationships []Relationship `json:"relationships"`
}

// Resource represents a single indexed resource (Deployment, Service, ConfigMap, Secret).
type Resource struct {
	Kind      string   `json:"kind"`
	Name      string   `json:"name"`
	Namespace string   `json:"namespace"`
	Replicas  *int32   `json:"replicas,omitempty"` // Deployment only
	Images    []string `json:"images,omitempty"`   // Deployment only
}

// Ref is a reference to a resource (kind/name/namespace).
type Ref struct {
	Kind      string `json:"kind"`
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
}

// Relationship describes a link between two resources (e.g. service-selects-deployment).
type Relationship struct {
	From Ref    `json:"from"`
	To   Ref    `json:"to"`
	Kind string `json:"kind"` // e.g. "service-selects-deployment", "deployment-uses-configmap"
}
