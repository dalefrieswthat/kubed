# Kubed

**Reduce AI agent token usage with two small files.**

CLI and context tooling for Kubernetes, Docker, Terraform, and Helm. Kubed gives AI agents two lightweight files instead of expensive discovery:

| File | What it stores | Size |
|------|----------------|------|
| `.kubed/layout.json` | Infra layout: Dockerfiles, Terraform, Helm, K8s, project structure, shared infra | ~1,500–3,000 tokens |
| `.kubed/learned.json` | Accumulated knowledge: architecture facts, important paths, dependencies, patterns | ~500 tokens |

**Without Kubed:** Agents run `find`, `ls -laR`, `kubectl get`, `grep -r` → 50,000+ tokens, 10+ tool calls, repeated every session.

**With Kubed:** Agents read two files → ~2,000–3,500 tokens, 1–2 file reads, knowledge persists across sessions.

- **Docs:** [cmds.daleyarborough.com](https://cmds.daleyarborough.com)
- **Layout index:** `kubed layout capture` writes `.kubed/layout.json`. Schema: [docs/LAYOUT_SCHEMA.md](docs/LAYOUT_SCHEMA.md).
- **Learned cache:** `kubed learned add-fact "..."` persists discoveries. Agents read it first to avoid re-discovering.
- **Reducing tokens:** [docs/reducing-token-usage.md](docs/reducing-token-usage.md)

## Installation (Python package)

```bash
pip install kubed
kubed-setup
```

Non-interactive: `kubed-setup --force-yes`. See [cmds.daleyarborough.com](https://cmds.daleyarborough.com) for full docs.

## Layout commands (Go binary)

Build the binary (requires Go 1.21+):

```bash
make build   # outputs build/kubed
# or: go build -o build/kubed ./cmd/kubed
```

- **`kubed layout capture`** — Index infra paths, project structure, shared infra, and (optionally) K8s resources; write `.kubed/layout.json`. Optional: `--all-namespaces`.
- **`kubed layout show`** — Print `layout.json` to stdout.

### Learned cache (accumulated knowledge)

- **`kubed learned show`** — Print `.kubed/learned.json` (facts, paths, deps, patterns accumulated across sessions).
- **`kubed learned summary`** — Quick summary: "3 facts, 2 paths, 5 deps, 1 patterns".
- **`kubed learned add-fact "fact" [--category=X] [--source=Y]`** — Add a discovered fact.
- **`kubed learned add-path "/path" "description" [--tags=a,b]`** — Add an important path.
- **`kubed learned add-dep "name" [--kind=X] [--version=Y]`** — Add a dependency.

Agents read `learned.json` first to avoid re-discovering the same things. When the agent learns something useful (architecture, dependencies, patterns), it persists it so future sessions benefit.

*This repo does not use Kubernetes; we validate layout with `make test` (fixture-based agent-query tests).*

## Development: test before pushing

Run **`make verify`** before pushing to main (same idea as Kubernetes: only working, quality code).

```bash
make verify   # runs go vet + tests (unit + agent-query tests)
make test     # tests only
make build    # build binary
make help     # list targets
```

Tests include **agent-query tests**: given a layout fixture, they assert that the data passed to an agent/LLM answers questions like "what workloads exist?", "what does service X select?", "what configmaps does deployment Y use?". See `internal/layout/query_test.go`.

---

# Kubed Documentation (site)

This is the documentation website for the Kubed CLI tool, built with Jekyll and Tailwind CSS.

## Setup

### Prerequisites

- Ruby 2.7+ with Bundler
- Node.js 14+ with npm

### Installation

1. Install Ruby dependencies:

```bash
bundle install
```

2. Install Node.js dependencies:

```bash
npm install
```

### Development

To run the site locally with live reload:

```bash
npm run dev
```

You can then access the site at http://localhost:4000

### Building for Production

```bash
npm run build:css
bundle exec jekyll build
```

The site will be generated in the `_site` directory.

### Content

- `index.md`: The home page
- `/docker/`, `/kubernetes/`, `/terraform/`, `/helm/`: Tool-specific documentation
- `installation.md`: Installation instructions

---

## License

MIT. See the LICENSE file for details.

## Contributing

Contributions are welcome. Please open an issue or submit a pull request.

## Contact

[cmds.daleyarborough.com](https://cmds.daleyarborough.com) · daleyarborough@gmail.com
