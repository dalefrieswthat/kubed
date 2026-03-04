# Kubed layout index schema

This document describes the schema of `.kubed/layout.json` produced by `kubed layout capture`. Agents and tools can consume this file for infra layout, project structure, and context without running multiple discovery commands.

## Location

- **Inside a git repo:** `.kubed/layout.json` at the repo root (kubed walks up to find `.git`).
- **Outside a git repo:** `~/.kubed/layout.json`.

## Format

JSON only. Current output is **v2** (section-based). Legacy v1 (k8s-only) is documented below for reference.

---

## Schema v2 (current)

`kubed layout capture` writes a **section-based index** so agents can pull only the sections they need (by `id` or `tags`), reducing token burn and latency.

### Top-level fields (v2)

| Field | Type | Description |
|-------|------|-------------|
| `version` | string | `"v2"`. |
| `generated_at` | string | RFC3339 UTC timestamp. |
| `repo_root` | string | Absolute path to the repo root. |
| `sections` | array of Section | Indexable sections. |

### Section (v2)

Each section has a stable `id`, optional `tags`, a one-line `summary`, and a `payload`. Use `id` or `tags` to select only relevant sections when answering a query.

| Field | Type | Description |
|-------|------|-------------|
| `id` | string | Stable identifier, e.g. `infra_paths`, `project_structure`, `overview`, `k8s_layout`. |
| `tags` | array of string | Tags for filter-by-topic. |
| `summary` | string | One-line summary for previews. |
| `payload` | object | Section data (shape depends on `id`). |

### Section IDs and payloads

- **infra_paths** — Payload: `paths[]` with `path` and `type` (terraform, docker, helm, dockerfile, compose, kubed).
- **project_structure** — Payload: `entries[]` with `path` and `kind` (dir, file).
- **overview, goals, repo-layout, etc.** — Payload: `content` and `source`. From `cursor-context/` files in `context.yaml`, split by `##` headings.
- **k8s_layout** — Present only when kubeconfig is available. Payload: same as v1 (context, namespaces, resources, relationships).

### Agent usage (v2)

- **Token optimization:** Return only sections matching the query (by `id` or `tags`).
- Run `kubed layout capture` after repo or cluster changes.

---

## Schema v1 (legacy / k8s_layout payload)

When a section has `id: "k8s_layout"`, its `payload` has this shape:

| Field | Type | Description |
|-------|------|-------------|
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
