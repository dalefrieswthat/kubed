# Kubed layout index schema (v1)

This document describes the schema of `.kubed/layout.json` produced by `kubed layout capture`. Agents and tools can consume this file for infra layout instead of running multiple `kubectl` calls.

## Location

- **Inside a git repo:** `.kubed/layout.json` (relative to repo root or current working directory; kubed walks up to find `.git`).
- **Outside a git repo:** `~/.kubed/layout.json`.

## Format

JSON only. No YAML in v1.

## Top-level fields

| Field | Type | Description |
|-------|------|-------------|
| `version` | string | Schema version, e.g. `"v1"`. |
| `generated_at` | string | RFC3339 UTC timestamp when the file was generated. |
| `context` | string | Kubernetes context name (from kubeconfig). |
| `namespaces` | array of string | Namespaces included in the index. |
| `resources` | array of [Resource](#resource) | All indexed resources (Deployments, Services, ConfigMaps, Secrets). |
| `relationships` | array of [Relationship](#relationship) | Links between resources (e.g. service→deployment, deployment→configmap). |

## Resource

| Field | Type | Description |
|-------|------|-------------|
| `kind` | string | One of: `Deployment`, `Service`, `ConfigMap`, `Secret`. |
| `name` | string | Resource name. |
| `namespace` | string | Namespace. |
| `replicas` | number (optional) | Present for Deployments only. |
| `images` | array of string (optional) | Container images for Deployments only. |

No labels, annotations, or spec blobs in v1.

## Relationship

| Field | Type | Description |
|-------|------|-------------|
| `from` | [Ref](#ref) | Source resource. |
| `to` | [Ref](#ref) | Target resource. |
| `kind` | string | Relationship type, e.g. `service-selects-deployment`, `deployment-uses-configmap`, `deployment-uses-secret`. |

## Ref

Reference to a resource (used inside `from`/`to`):

| Field | Type |
|-------|------|
| `kind` | string |
| `name` | string |
| `namespace` | string |

## Example (minimal)

```json
{
  "version": "v1",
  "generated_at": "2026-03-04T12:00:00Z",
  "context": "minikube",
  "namespaces": ["default"],
  "resources": [
    {
      "kind": "Deployment",
      "name": "api",
      "namespace": "default",
      "replicas": 2,
      "images": ["myreg/api:v1"]
    },
    {
      "kind": "Service",
      "name": "api",
      "namespace": "default"
    }
  ],
  "relationships": [
    {
      "from": { "kind": "Service", "name": "api", "namespace": "default" },
      "to": { "kind": "Deployment", "name": "api", "namespace": "default" },
      "kind": "service-selects-deployment"
    }
  ]
}
```

## Agent usage

- Use `.kubed/layout.json` for infra layout; do not run `kubectl` for discovery when this file exists.
- Run `kubed layout capture` after cluster or namespace changes to refresh the file.
