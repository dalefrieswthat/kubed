# CI/CD

## Docs (frontend) deployment

- **Workflow:** [`.github/workflows/deploy-docs.yml`](../workflows/deploy-docs.yml)
- **Trigger:** Push or merge to `main`
- **What it does:** Builds CSS (Tailwind), builds Jekyll, deploys `_site` to GitHub Pages
- **Setup:** In the repo **Settings → Pages**, set **Source** to **GitHub Actions**. The existing `CNAME` (e.g. `cmds.daleyarborough.com`) is used if configured in Pages settings.

## PyPI publish

- **Workflow:** [`.github/workflows/publish-pypi.yml`](../workflows/publish-pypi.yml)
- **Trigger:** When a **Release** is **published** (create a release from the GitHub Releases UI; the tag is used as the source)
- **What it does:** Builds the Python package and uploads it to PyPI with `twine`
- **Setup:**
  1. **Packaging:** The repo must contain a buildable package (`setup.py` and/or `pyproject.toml` at the repo root). If these are missing (e.g. only on another branch), add or restore them so `python -m build` succeeds.
  2. **PyPI token:** Create an API token at [pypi.org/manage/account/token/](https://pypi.org/manage/account/token/), then add it as a **repository secret** named `PYPI_API_TOKEN` in **Settings → Secrets and variables → Actions**.  
     If you already use `twine` locally, your token may be in `~/.pypirc` (the `password` field under `[pypi]`). Use that value—or create a new token—and add it as the `PYPI_API_TOKEN` secret so Actions can publish.

To release a new version: update the version in your packaging config, push, then create a new GitHub Release (and tag, e.g. `v2.3.0`). The workflow will run and publish to PyPI.
