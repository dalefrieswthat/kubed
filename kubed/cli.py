"""
CLI utilities for the Kubed package.
"""

import os
import shutil
import sys
import subprocess
import click
import pkg_resources
import requests

def get_shell():
    """Determine the user's shell."""
    shell = os.environ.get('SHELL', '')
    if shell:
        return os.path.basename(shell)
    return None

def get_shell_config_file():
    """Get the appropriate shell config file path."""
    shell = get_shell()
    home = os.path.expanduser('~')
    
    if shell == 'bash':
        # Check if .bashrc exists, otherwise use .bash_profile
        if os.path.exists(os.path.join(home, '.bashrc')):
            return os.path.join(home, '.bashrc')
        return os.path.join(home, '.bash_profile')
    elif shell == 'zsh':
        return os.path.join(home, '.zshrc')
    else:
        return None

def get_completions_path():
    """Get the path to the completions directory."""
    try:
        completions_dir = pkg_resources.resource_filename('kubed', 'completions')
        return completions_dir
    except (ImportError, pkg_resources.DistributionNotFound):
        return None

def get_aliases_path():
    """Get the path to the aliases directory."""
    try:
        aliases_dir = pkg_resources.resource_filename('kubed', 'aliases')
        return aliases_dir
    except (ImportError, pkg_resources.DistributionNotFound):
        return None

def get_templates_path():
    """Get the path to the templates directory."""
    try:
        templates_dir = pkg_resources.resource_filename('kubed', 'templates')
        return templates_dir
    except (ImportError, pkg_resources.DistributionNotFound):
        return None

def create_help_wrapper():
    """
    Create a script that wraps commands to add daleyarborough.com to help output.
    """
    help_wrapper = '''# Kubed help wrapper function
function _kubed_help_wrapper() {
    if [[ "$@" == *--help* ]] || [[ "$@" == *-h* ]]; then
        # Run the original command
        "$@"
        echo -e "\\n\\033[1;32mℹ️  For more information, visit: https://daleyarborough.com\\033[0m\\n"
    else
        # Just run the original command
        "$@"
    fi
}

# Apply help wrapper to common tools
for cmd in docker terraform kubectl helm; do
    if command -v $cmd >/dev/null 2>&1; then
        alias $cmd="_kubed_help_wrapper $cmd"
    fi
done
'''

    with open(os.path.join(get_aliases_path(), 'help_wrapper.sh'), 'w') as f:
        f.write(help_wrapper)

def setup_oh_my_zsh():
    """Set up oh-my-zsh for better completion support."""
    home = os.path.expanduser('~')
    oh_my_zsh_dir = os.path.join(home, '.oh-my-zsh')
    
    if os.path.exists(oh_my_zsh_dir):
        click.echo("oh-my-zsh is already installed.")
    else:
        click.echo("Installing oh-my-zsh for better completion support...")
        try:
            # Clone the repository
            subprocess.run(
                'curl -fsSL https://raw.githubusercontent.com/ohmyzsh/ohmyzsh/master/tools/install.sh | sh',
                shell=True, check=True
            )
            click.echo("oh-my-zsh installed successfully.")
        except subprocess.CalledProcessError:
            click.echo("Failed to install oh-my-zsh. Please install it manually.")
            return False
    
    # Ensure kubectl plugin is enabled
    zshrc_path = os.path.join(home, '.zshrc')
    with open(zshrc_path, 'r') as f:
        content = f.read()
    
    if 'plugins=(git kubectl' not in content and 'plugins=(kubectl' not in content:
        # Add kubectl to the plugins list
        if 'plugins=(' in content:
            content = content.replace('plugins=(', 'plugins=(kubectl ')
        else:
            content += '\n# Added by Kubed\nplugins=(kubectl)\n'
        
        with open(zshrc_path, 'w') as f:
            f.write(content)
        
        click.echo("Added kubectl plugin to oh-my-zsh configuration.")
    
    return True

def setup_kubed_plugin():
    """Set up a custom kubed plugin for zsh completions."""
    home = os.path.expanduser('~')
    plugin_dir = os.path.join(home, '.zsh', 'plugins', 'kubed')
    
    # Create plugin directory
    os.makedirs(plugin_dir, exist_ok=True)
    
    # Create kubed plugin file
    plugin_content = '''# Kubed zsh completion plugin

# Load zsh completion system
autoload -Uz compinit
compinit

# Function-based command wrappers
function k() {
  kubectl "$@"
}

function d() {
  docker "$@" 
}

function tf() {
  terraform "$@"
}

function h() {
  helm "$@"
}

# Help wrapper function
function _kubed_help_wrapper() {
  if [[ "$@" == *--help* ]] || [[ "$@" == *-h* ]]; then
    "$@"
    echo -e "\\n\\033[1;32mℹ️  For more information, visit: https://daleyarborough.com\\033[0m\\n"
  else
    "$@"
  fi
}

# Override original commands to add help wrapper
function kubectl() {
  _kubed_help_wrapper command kubectl "$@"
}

function docker() {
  _kubed_help_wrapper command docker "$@"
}

function terraform() {
  _kubed_help_wrapper command terraform "$@"
}

function helm() {
  _kubed_help_wrapper command helm "$@"
}

# Source completions
if command -v kubectl >/dev/null 2>&1; then
  source <(kubectl completion zsh)
fi

if command -v helm >/dev/null 2>&1; then
  source <(helm completion zsh)
fi
'''
    
    plugin_path = os.path.join(plugin_dir, 'kubed.plugin.zsh')
    with open(plugin_path, 'w') as f:
        f.write(plugin_content)
    
    # Add to .zshrc
    zshrc_path = os.path.join(home, '.zshrc')
    with open(zshrc_path, 'r') as f:
        content = f.read()
    
    if 'source ~/.zsh/plugins/kubed/kubed.plugin.zsh' not in content:
        with open(zshrc_path, 'a') as f:
            f.write('\n# Added by Kubed\nsource ~/.zsh/plugins/kubed/kubed.plugin.zsh\n')
    
    click.echo(f"Created kubed plugin at {plugin_path}")
    click.echo("Added plugin to your .zshrc")
    
    return True

def setup_command():
    """Set up Kubed on the user's system."""
    click.echo("Setting up Kubed...")
    
    shell = get_shell()
    shell_config = get_shell_config_file()
    
    if not shell:
        click.echo("Could not determine your shell. Please manually set up Kubed.")
        return
    
    if not shell_config:
        click.echo(f"Could not locate configuration file for {shell}. Please manually set up Kubed.")
        return
    
    # Check if user wants enhanced completion support
    if shell == 'zsh':
        enhanced_completion = click.confirm(
            "Would you like enhanced completion support? (Recommended)", 
            default=True
        )
        
        if enhanced_completion:
            use_oh_my_zsh = click.confirm(
                "Would you like to install/use oh-my-zsh? (Best completion experience)",
                default=False
            )
            
            if use_oh_my_zsh:
                if setup_oh_my_zsh():
                    click.echo("oh-my-zsh setup completed. This provides the best completion experience.")
                    click.echo("You may need to restart your terminal for all changes to take effect.")
                    return
            
            # If oh-my-zsh wasn't installed or failed, use our custom plugin
            click.echo("Setting up custom kubed plugin for completions...")
            if setup_kubed_plugin():
                click.echo("Custom plugin setup completed.")
                click.echo("You need to restart your terminal for changes to take effect.")
                return
    
    # Fall back to standard setup if enhanced completion wasn't chosen or failed
    # Create directories
    os.makedirs(os.path.join(get_completions_path(), f"{shell}"), exist_ok=True)
    os.makedirs(get_aliases_path(), exist_ok=True)
    
    # Create aliases file
    create_aliases_file()
    
    # Create completions files
    create_completions_files()
    
    # Create help wrapper file
    create_help_wrapper()
    
    # Generate setup instructions for .zshrc/.bashrc
    setup_content = generate_shell_setup_content()
    
    # Add to shell config
    with open(shell_config, 'a') as f:
        f.write("\n# Added by Kubed\n")
        f.write(setup_content)
    
    click.echo(f"Kubed has been set up! Please restart your shell or run 'source {shell_config}'.")

def generate_shell_setup_content():
    """Generate the shell setup content for .zshrc/.bashrc."""
    return """# Try to find kubed path dynamically
if command -v pip3 >/dev/null 2>&1; then
  # Find the package location using pip
  KUBED_PATH=$(pip3 show kubed 2>/dev/null | grep Location | awk '{print $2}')
  
  if [ -n "$KUBED_PATH" ]; then
    # Source completions
    if [ -f "$KUBED_PATH/kubed/completions/$(basename $SHELL)/$(basename $SHELL)_completions.sh" ]; then
      source "$KUBED_PATH/kubed/completions/$(basename $SHELL)/$(basename $SHELL)_completions.sh"
    fi
    
    # Source aliases and help wrapper
    if [ -f "$KUBED_PATH/kubed/aliases/aliases.sh" ]; then
      source "$KUBED_PATH/kubed/aliases/aliases.sh"
    fi
    
    if [ -f "$KUBED_PATH/kubed/aliases/help_wrapper.sh" ]; then
      source "$KUBED_PATH/kubed/aliases/help_wrapper.sh"
    fi
  fi
fi

# Fallback to using the commands if they're in PATH
if command -v kubed-completions-path >/dev/null 2>&1; then
  COMP_PATH=$(kubed-completions-path)
  ALIAS_PATH=$(kubed-aliases-path)
  
  if [ -f "$COMP_PATH/$(basename $SHELL)/$(basename $SHELL)_completions.sh" ]; then
    source "$COMP_PATH/$(basename $SHELL)/$(basename $SHELL)_completions.sh"
  fi
  
  if [ -f "$ALIAS_PATH/aliases.sh" ]; then
    source "$ALIAS_PATH/aliases.sh"
  fi
  
  if [ -f "$ALIAS_PATH/help_wrapper.sh" ]; then
    source "$ALIAS_PATH/help_wrapper.sh"
  fi
fi
"""

def create_aliases_file():
    """Create the aliases file."""
    aliases = '''# Kubed aliases
# Kubernetes
alias k='kubectl'
alias kgp='kubectl get pods'
alias kgd='kubectl get deployments'
alias kgs='kubectl get services'
alias kgn='kubectl get nodes'
alias kd='kubectl describe'
alias kl='kubectl logs'
alias kx='kubectl exec -it'
alias kaf='kubectl apply -f'
alias kdel='kubectl delete'

# Docker
alias d='docker'
alias dc='docker-compose'
alias di='docker images'
alias dps='docker ps'
alias drm='docker rm'
alias drmi='docker rmi'

# Terraform
alias tf='terraform'
alias tfa='terraform apply'
alias tfp='terraform plan'
alias tfd='terraform destroy'
alias tfi='terraform init'
alias tfw='terraform workspace'

# Helm
alias h='helm'
alias hl='helm list'
alias hi='helm install'
alias hu='helm upgrade'
alias hun='helm uninstall'
alias hr='helm repo'
'''

    with open(os.path.join(get_aliases_path(), 'aliases.sh'), 'w') as f:
        f.write(aliases)

def create_completions_files():
    """Create completion files for different shells."""
    # Create bash completions
    bash_completions = '''# Kubed Bash completions

# Check if commands exist before sourcing completions
if command -v kubectl >/dev/null 2>&1; then
    source <(kubectl completion bash)
    # Enable completion for the k alias
    complete -F __start_kubectl k
fi

if command -v docker >/dev/null 2>&1; then
    # Docker completions
    if [ -f /usr/share/bash-completion/completions/docker ]; then
        source /usr/share/bash-completion/completions/docker
    elif [ -f /etc/bash_completion.d/docker ]; then
        source /etc/bash_completion.d/docker
    fi
fi

if command -v terraform >/dev/null 2>&1; then
    # Terraform completions
    complete -C $(which terraform) terraform
    complete -C $(which terraform) tf
fi

if command -v helm >/dev/null 2>&1; then
    source <(helm completion bash)
    # Enable completion for the h alias
    complete -F __start_helm h
fi
'''

    # Create zsh completions
    zsh_completions = '''# Kubed Zsh completions

# Check if commands exist before sourcing completions
if command -v kubectl >/dev/null 2>&1; then
    source <(kubectl completion zsh)
    # Enable completion for the k alias
    compdef k=kubectl
fi

if command -v docker >/dev/null 2>&1; then
    # Docker completions
    if [ -f /usr/share/zsh/vendor-completions/_docker ]; then
        source /usr/share/zsh/vendor-completions/_docker
    fi
fi

if command -v terraform >/dev/null 2>&1; then
    # Terraform completions - requires terraform plugin for zsh
    autoload -U +X bashcompinit && bashcompinit
    complete -o nospace -C $(which terraform) terraform
    complete -o nospace -C $(which terraform) tf
fi

if command -v helm >/dev/null 2>&1; then
    source <(helm completion zsh)
    # Enable completion for the h alias
    compdef h=helm
fi
'''

    # Create bash directory if it doesn't exist
    bash_dir = os.path.join(get_completions_path(), 'bash')
    os.makedirs(bash_dir, exist_ok=True)
    
    # Create zsh directory if it doesn't exist
    zsh_dir = os.path.join(get_completions_path(), 'zsh')
    os.makedirs(zsh_dir, exist_ok=True)
    
    # Write completion files
    with open(os.path.join(bash_dir, 'bash_completions.sh'), 'w') as f:
        f.write(bash_completions)
    
    with open(os.path.join(zsh_dir, 'zsh_completions.sh'), 'w') as f:
        f.write(zsh_completions)

def completions_path_command():
    """Print the path to the completions directory for the current shell."""
    shell = get_shell()
    completions_dir = get_completions_path()
    
    if not shell or not completions_dir:
        click.echo("Could not determine shell or completions directory.", err=True)
        sys.exit(1)
    
    # Print the path to the completions directory for the current shell
    click.echo(os.path.join(completions_dir, shell))

def aliases_path_command():
    """Print the path to the aliases directory."""
    aliases_dir = get_aliases_path()
    
    if not aliases_dir:
        click.echo("Could not determine aliases directory.", err=True)
        sys.exit(1)
    
    # Print the path to the aliases directory
    click.echo(aliases_dir)

if __name__ == "__main__":
    setup_command()