package layout

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

// loadFixture loads testdata/layout.json for agent-query tests.
func loadFixture(t *testing.T) *Layout {
	t.Helper()
	data, err := os.ReadFile(filepath.Join("testdata", "layout.json"))
	if err != nil {
		t.Fatalf("read fixture: %v", err)
	}
	var l Layout
	if err := json.Unmarshal(data, &l); err != nil {
		t.Fatalf("unmarshal fixture: %v", err)
	}
	return &l
}

// Agent-style query tests: ensure the right information is available for questions
// a user might ask an agent. These guard that layout capture + query helpers pass
// the correct data to the agent/LLM.

func TestAgentQuery_WhatWorkloadsExist(t *testing.T) {
	l := loadFixture(t)
	// "What deployments exist?"
	deps := l.ResourcesByKind("Deployment")
	if len(deps) != 2 {
		t.Fatalf("expected 2 deployments, got %d", len(deps))
	}
	names := map[string]bool{}
	for _, d := range deps {
		names[d.Name] = true
	}
	if !names["api"] || !names["worker"] {
		t.Errorf("expected deployments api and worker, got %v", names)
	}
	// "What services exist?"
	svcs := l.ResourcesByKind("Service")
	if len(svcs) != 2 {
		t.Fatalf("expected 2 services, got %d", len(svcs))
	}
}

func TestAgentQuery_WhatDoesServiceSelect(t *testing.T) {
	l := loadFixture(t)
	// "What deployment does service 'api' select?"
	refs := l.DeploymentSelectors(Ref{Kind: "Service", Name: "api", Namespace: "default"})
	if len(refs) != 1 {
		t.Fatalf("expected 1 deployment for service api, got %d", len(refs))
	}
	if refs[0].Name != "api" || refs[0].Kind != "Deployment" {
		t.Errorf("expected deployment api, got %+v", refs[0])
	}
	// "What deployment does service 'worker' select?"
	refs = l.DeploymentSelectors(Ref{Kind: "Service", Name: "worker", Namespace: "app"})
	if len(refs) != 1 || refs[0].Name != "worker" {
		t.Errorf("expected deployment worker, got %+v", refs)
	}
}

func TestAgentQuery_WhatDoesDeploymentUse(t *testing.T) {
	l := loadFixture(t)
	// "What configmaps and secrets does deployment 'api' use?"
	deps := l.DeploymentDependencies(Ref{Kind: "Deployment", Name: "api", Namespace: "default"})
	if len(deps) != 2 {
		t.Fatalf("expected 2 dependencies (configmap + secret), got %d", len(deps))
	}
	names := map[string]bool{}
	for _, r := range deps {
		names[r.Name] = true
	}
	if !names["api-config"] || !names["api-secret"] {
		t.Errorf("expected api-config and api-secret, got %v", names)
	}
	// "What does deployment 'worker' use?" (none in fixture)
	deps = l.DeploymentDependencies(Ref{Kind: "Deployment", Name: "worker", Namespace: "app"})
	if len(deps) != 0 {
		t.Errorf("expected 0 dependencies for worker, got %d", len(deps))
	}
}

func TestAgentQuery_WhatUsesConfigMap(t *testing.T) {
	l := loadFixture(t)
	// "What uses configmap 'api-config'?"
	rels := l.RelationshipsTo(Ref{Kind: "ConfigMap", Name: "api-config", Namespace: "default"})
	if len(rels) != 1 {
		t.Fatalf("expected 1 relationship to api-config, got %d", len(rels))
	}
	if rels[0].From.Kind != "Deployment" || rels[0].From.Name != "api" {
		t.Errorf("expected from Deployment api, got %+v", rels[0].From)
	}
}

func TestAgentQuery_ResourceByRef(t *testing.T) {
	l := loadFixture(t)
	// "Get deployment 'api' details"
	r := l.ResourceByRef(Ref{Kind: "Deployment", Name: "api", Namespace: "default"})
	if r == nil {
		t.Fatal("expected deployment api")
	}
	if r.Replicas == nil || *r.Replicas != 2 {
		t.Errorf("expected replicas 2, got %v", r.Replicas)
	}
	if len(r.Images) != 1 || r.Images[0] != "myreg/api:v1" {
		t.Errorf("expected image myreg/api:v1, got %v", r.Images)
	}
	// Missing resource
	r = l.ResourceByRef(Ref{Kind: "Deployment", Name: "nonexist", Namespace: "default"})
	if r != nil {
		t.Errorf("expected nil for nonexist, got %+v", r)
	}
}
