# Kubed

CLI productivity tool for Docker, Kubernetes, Terraform, and Helm: completions, aliases, and a **layout index** for agent-friendly infra views.

- **Docs site:** [cmds.daleyarborough.com](https://cmds.daleyarborough.com) — built with Jekyll + Tailwind (see [Setup](#setup) below).
- **Layout index (V1):** `kubed layout capture` writes `.kubed/layout.json` from your current kube context; `kubed layout show` prints it. Agents can read this file instead of running `kubectl` for discovery. Schema: [docs/LAYOUT_SCHEMA.md](docs/LAYOUT_SCHEMA.md).

## Layout commands (Go binary)

Build the binary (requires Go 1.21+):

```bash
make build   # outputs build/kubed
# or: go build -o build/kubed ./cmd/kubed
```

- **`kubed layout capture`** — Connect to current kube context; index Deployments, Services, ConfigMaps, Secrets and their relationships; write `.kubed/layout.json` (or `~/.kubed/layout.json` if not in a git repo). Optional: `--all-namespaces`.
- **`kubed layout show`** — Print `layout.json` to stdout. If missing, prints "run kubed layout capture" and exits 1.

Use `.kubed/layout.json` for infra layout; run `kubed layout capture` after cluster changes.

## Development: test before pushing

Run **`make verify`** before pushing to main (same idea as Kubernetes: only working, quality code).

```bash
make verify   # runs go vet + tests (unit + agent-query tests)
make test     # tests only
make build    # build binary
make help     # list targets
```

Tests include **agent-query tests**: given a layout fixture, they assert that the data passed to an agent/LLM answers questions like “what workloads exist?”, “what does service X select?”, “what configmaps does deployment Y use?”. See `internal/layout/query_test.go`.

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

This will:
- Build the Tailwind CSS
- Watch for CSS changes
- Run the Jekyll server

You can then access the site at http://localhost:4000

### Building for Production

To build the site for production:

```bash
npm run build:css
bundle exec jekyll build
```

The site will be generated in the `_site` directory.

## Customization

### Tailwind Configuration

The Tailwind CSS configuration is located in `tailwind.config.js`. You can modify this file to customize the theme.

### Layout Files

The main layout templates are in the `_layouts` directory. The `default.html` layout is the main template.

### Content Files

- `index.md`: The home page
- `/docker/`, `/kubernetes/`, `/terraform/`, `/helm/`: Tool-specific documentation
- `installation.md`: Installation instructions

## Implementation Notes

This site uses:

- **Jekyll** as the static site generator
- **Tailwind CSS** for styling
- **PostCSS** for processing CSS
- **Font Awesome** for icons

The site is designed to be responsive and follows modern web development practices. 