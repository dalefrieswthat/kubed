package layout

// Query helpers answer agent-style questions from a Layout so the right information
// is passed to the agent/LLM. Tests assert these return expected results for fixture data.

// ResourcesByKind returns all resources of the given kind (e.g. "Deployment", "Service").
func (l *Layout) ResourcesByKind(kind string) []Resource {
	var out []Resource
	for _, r := range l.Resources {
		if r.Kind == kind {
			out = append(out, r)
		}
	}
	return out
}

// ResourceByRef returns the resource matching ref, or nil.
func (l *Layout) ResourceByRef(ref Ref) *Resource {
	for i := range l.Resources {
		r := &l.Resources[i]
		if r.Kind == ref.Kind && r.Name == ref.Name && r.Namespace == ref.Namespace {
			return r
		}
	}
	return nil
}

// RelationshipsFrom returns relationships where from matches ref (e.g. "what does this service select?").
func (l *Layout) RelationshipsFrom(ref Ref) []Relationship {
	var out []Relationship
	for _, rel := range l.Relationships {
		if ref.Kind == rel.From.Kind && ref.Name == rel.From.Name && ref.Namespace == rel.From.Namespace {
			out = append(out, rel)
		}
	}
	return out
}

// RelationshipsTo returns relationships where to matches ref (e.g. "what uses this configmap?").
func (l *Layout) RelationshipsTo(ref Ref) []Relationship {
	var out []Relationship
	for _, rel := range l.Relationships {
		if ref.Kind == rel.To.Kind && ref.Name == rel.To.Name && ref.Namespace == rel.To.Namespace {
			out = append(out, rel)
		}
	}
	return out
}

// DeploymentSelectors returns refs of Deployments selected by the given Service ref (by relationship).
func (l *Layout) DeploymentSelectors(serviceRef Ref) []Ref {
	var refs []Ref
	for _, rel := range l.RelationshipsFrom(serviceRef) {
		if rel.Kind == "service-selects-deployment" {
			refs = append(refs, rel.To)
		}
	}
	return refs
}

// DeploymentDependencies returns refs of ConfigMaps and Secrets used by the given Deployment ref.
func (l *Layout) DeploymentDependencies(depRef Ref) []Ref {
	var refs []Ref
	for _, rel := range l.RelationshipsFrom(depRef) {
		if rel.Kind == "deployment-uses-configmap" || rel.Kind == "deployment-uses-secret" {
			refs = append(refs, rel.To)
		}
	}
	return refs
}
