# Reducing token usage in Cursor

Cursor usage (e.g. on the Usage tab) shows token counts per request. Requests can reach 1M+ tokens when the agent pulls in large context: full directory listings, many search results, or repeated discovery. This page describes how to **actually** reduce those numbers when working with infrastructure and project layout.

## 1. Use the layout file instead of discovery

- **Run** `kubed layout capture` at the repo root so `.kubed/layout.json` exists.
- **For infra/structure questions** (“what services do we have?”, “where is the Terraform?”, “what’s the project layout?”), the agent should read `.kubed/layout.json` instead of running `find`, `ls -laR`, `grep -r`, or `kubectl get ...`. Discovery output is often 50K–500K+ tokens; the layout file is on the order of 1–2K tokens.
- **Cursor rule:** This repo includes a rule (`.cursor/rules/kubed-infra.mdc`) that tells the agent to use the layout first and avoid discovery. It applies when you’re in infra-related files or `.kubed/`. In other repos, add a similar rule or @-mention `.kubed/layout.json` when asking infra questions.

## 2. Use the learned cache (accumulated knowledge)

`.kubed/learned.json` stores facts, important paths, dependencies, and patterns discovered during previous sessions. The agent reads this file first to avoid re-discovering the same things.

**View the cache:**
```bash
kubed learned show      # full JSON
kubed learned summary   # quick count: "3 facts, 2 paths, 5 deps, 1 patterns"
```

**Add to the cache** (manually or via agent):
```bash
kubed learned add-fact "API uses PostgreSQL" --category=architecture --source=docker-compose.yml
kubed learned add-path "src/api/" "Main API service" --tags=api,backend
kubed learned add-dep "PostgreSQL" --kind=database --version=15 --used-by=api
```

**What to persist:**
- Architecture decisions (how services connect, what databases are used)
- Important paths with descriptions
- Tech stack (databases, frameworks, libraries)
- Code patterns (where tests live, config file conventions)

When the agent discovers something useful, it should persist it so the next session doesn't re-discover.

## 3. Read only the section you need

The layout is section-based. If the question is only about “where are the Dockerfiles?”, the agent only needs the `infra_paths` section. Reading that section (or a small script that extracts it) keeps context minimal instead of loading the whole file and all other sections.

| Question type | Section `id` to use |
|---------------|---------------------|
| Infra paths (Dockerfiles, Terraform, Helm) | `infra_paths` |
| Top-level project structure | `project_structure` |
| Shared/cross-repo infra | `shared_infra` |
| K8s resources and relationships | `k8s_layout` |
| Project overview, goals | context sections (e.g. `overview`, `goals`) |

## 4. Narrow context in the chat

- **@-mention specific files** instead of the whole repo or huge folders. For “where is X?”, @-mention `.kubed/layout.json` so the agent uses it as the primary context.
- **New chat for new topics** so previous long context doesn’t keep getting included.
- **Avoid** “search the whole codebase” or “list everything” unless you really need it; prefer “read the layout and tell me …”.

## 5. Model and plan choices

- **Thinking / long-reasoning models** use many output tokens. Use them when you need deep reasoning; for “what exists where?” use a non-thinking or smaller model.
- **Scope the ask:** “What Terraform do we have in this repo?” is better than “Analyze the entire repo” when you only need a high-level list.

## 6. Keep the layout and learned cache up to date

Run `kubed layout capture` after significant repo or cluster changes. A stale layout can cause the agent to run discovery anyway to “double-check”, which increases tokens again.

---

**Summary:** Generate the layout (`kubed layout capture`), persist discoveries to the learned cache (`kubed learned add-*`), point the agent at these files for infra/structure questions (via rule or @-mention), and prefer a single section when possible. That replaces large discovery context with small, fixed snapshots and reduces token usage per request.
