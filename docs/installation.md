---
layout: default
title: Installation Guide
nav_order: 2
description: "Installation instructions for Kubed CLI tool"
permalink: /installation
---

# Installation Guide

Kubed is a powerful CLI tool that enhances your terminal experience when working with Docker, Kubernetes, Terraform, and Helm. This guide covers all installation options and requirements.

## Requirements

Before installing Kubed, ensure you have the following prerequisites:

- **Python 3.6 or later**
- **Pip** (Python package installer)
- **Optional tools** that will be installed automatically if missing:
  - Docker
  - Kubernetes (kubectl)
  - Terraform
  - Helm

## Basic Installation

The simplest way to install Kubed is using pip:

```bash
pip install kubed
```

After installation, you'll need to run the setup command to configure your environment:

```bash
kubed-setup
```

This will:
1. Check for required tools and offer to install any that are missing
2. Set up shell completions and aliases
3. Configure your shell environment
4. For zsh users, install oh-my-zsh and Powerlevel10k theme for optimal experience
5. For fish users, install Starship prompt for optimal experience

## Non-interactive Installation

For automated deployments or scripts, you can use the `--force-yes` flag to automatically accept all installation prompts:

```bash
kubed-setup --force-yes
```

This will automatically:
- Install any missing tools without prompting
- Set up oh-my-zsh and Powerlevel10k for zsh users
- Set up Starship for fish users
- Configure your shell environment

## Installation Options

The `kubed-setup` command accepts the following options:

| Option | Description |
|--------|-------------|
| `--zsh` | Force using ZSH configuration even if not detected |
| `--force-yes` | Automatically answer yes to all installation prompts |

Example:
```bash
kubed-setup --zsh --force-yes
```

## Post-Installation

**IMPORTANT:** After installation, you MUST either:
1. Restart your terminal, or
2. Source your shell configuration file:
   ```bash
   source ~/.zshrc    # for zsh users
   source ~/.bashrc   # for bash users
   source ~/.config/fish/config.fish  # for fish users
   ```

Kubed will offer to run the source command for you during setup.

## Manual Component Installation

If you prefer to install the tools manually, you can use the following commands:

### Docker
```bash
brew install docker    # macOS
```

### Kubernetes (kubectl)
```bash
brew install kubectl   # macOS
```

### Terraform
```bash
brew install terraform # macOS
```

### Helm
```bash
brew install helm      # macOS
```

### Fish Shell (Recommended)
```bash
brew install fish      # macOS
```

### Starship Prompt (for Fish)
```bash
curl -sS https://starship.rs/install.sh | sh
```

## Troubleshooting

### Missing setuptools
If you encounter errors related to `pkg_resources` or `setuptools`, install or upgrade setuptools:

```bash
pip install --upgrade setuptools
```

### Shell Issues
If your shell isn't properly configured after installation:

1. Check if the kubed setup block exists in your shell config:
   ```bash
   grep kubed ~/.zshrc  # or ~/.bashrc or ~/.config/fish/config.fish
   ```
2. Manually source the configuration:
   ```bash
   source ~/.zshrc      # or ~/.bashrc or ~/.config/fish/config.fish
   ```
3. Run setup again:
   ```bash
   kubed-setup
   ```

### Contact Support
For additional help, please [open an issue](https://github.com/dalefrieswthat/kubed/issues) or contact support at daleyarborough@gmail.com. 