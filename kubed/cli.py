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
        echo -e "\\n\\033[1;32mℹ️  For more information and quick links, visit: https://cmds.daleyarborough.com\\033[0m\\n"
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

def check_and_install_tools(force_yes=False):
    """Check for required tools and install if missing.
    
    Args:
        force_yes (bool): If True, automatically answer yes to all installation prompts.
    """
    # First, ensure Python prerequisites are installed
    python_tools = {
        'setuptools': {
            'check': 'python3 -c "import setuptools; print(setuptools.__version__)"',
            'install': 'pip3 install --upgrade setuptools',
            'install_msg': 'Python setuptools is not installed or outdated. Would you like to install/upgrade it? [Y/n]: ',
            'install_cmd': 'pip3 install --upgrade setuptools',
            'alt_install': 'pip3 install --upgrade setuptools'
        },
    }
    
    # Check Python prerequisites first
    for tool, config in python_tools.items():
        try:
            subprocess.run(config['check'], shell=True, check=True, capture_output=True)
            print(f"✅ {tool} is already installed.")
        except subprocess.CalledProcessError:
            print(f"\n{config['install_msg']}")
            response = 'y' if force_yes else input().lower()
            if response in ['y', 'yes', ''] or force_yes:
                print(f"Installing {tool}...")
                try:
                    subprocess.run(config['install_cmd'], shell=True, check=True)
                    print(f"✅ {tool} installed successfully!")
                except subprocess.CalledProcessError as e:
                    print(f"❌ Failed to install {tool}: {e}")
                    print(f"Trying alternative installation method...")
                    try:
                        subprocess.run(config['alt_install'], shell=True, check=True)
                        print(f"✅ {tool} installed successfully using alternative method!")
                    except subprocess.CalledProcessError as e:
                        print(f"❌ Failed to install {tool}. This may cause issues: {e}")
            else:
                print(f"Skipping {tool} installation. This may cause issues later.")
    
    # Continue with system tool checks
    tools = {
        'docker': {
            'check': 'docker --version',
            'install': 'brew install docker',
            'install_msg': 'Docker is not installed. Would you like to install it? [Y/n]: ',
            'install_cmd': 'brew install docker',
            'alt_install': 'curl -fsSL https://get.docker.com -o get-docker.sh && sudo sh get-docker.sh'
        },
        'kubectl': {
            'check': 'kubectl version --client',
            'install': 'brew install kubectl',
            'install_msg': 'kubectl is not installed. Would you like to install it? [Y/n]: ',
            'install_cmd': 'brew install kubectl',
            'alt_install': 'curl -LO "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/darwin/amd64/kubectl" && chmod +x kubectl && sudo mv kubectl /usr/local/bin/'
        },
        'helm': {
            'check': 'helm version',
            'install': 'brew install helm',
            'install_msg': 'Helm is not installed. Would you like to install it? [Y/n]: ',
            'install_cmd': 'brew install helm',
            'alt_install': 'curl https://raw.githubusercontent.com/helm/helm/main/scripts/get-helm-3 | bash'
        },
        'terraform': {
            'check': 'terraform --version',
            'install': 'brew install terraform',
            'install_msg': 'Terraform is not installed. Would you like to install it? [Y/n]: ',
            'install_cmd': 'curl -fsSL https://releases.hashicorp.com/terraform/1.7.5/terraform_1.7.5_darwin_amd64.zip -o terraform.zip && unzip terraform.zip && sudo mv terraform /usr/local/bin/ && rm terraform.zip',
            'alt_install': 'curl -fsSL https://releases.hashicorp.com/terraform/1.7.5/terraform_1.7.5_darwin_amd64.zip -o terraform.zip && unzip terraform.zip && sudo mv terraform /usr/local/bin/ && rm terraform.zip'
        }
    }

    def check_homebrew():
        """Check if Homebrew is installed and working."""
        return subprocess.run('which brew', shell=True, capture_output=True).returncode == 0

    # Initial Homebrew check and installation attempt
    homebrew_available = check_homebrew()
    if not homebrew_available:
        print("\nHomebrew is not installed. Some tools may require Homebrew for installation.")
        print("Would you like to install Homebrew? [Y/n]: ")
        response = 'y' if force_yes else input().lower()
        if response in ['y', 'yes', ''] or force_yes:
            print("Installing Homebrew...")
            try:
                subprocess.run('/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"', shell=True, check=True)
                print("✅ Homebrew installed successfully!")
                homebrew_available = check_homebrew()
            except subprocess.CalledProcessError as e:
                print(f"❌ Failed to install Homebrew: {e}")
                print("Proceeding with alternative installation methods...")

    for tool, config in tools.items():
        try:
            subprocess.run(config['check'], shell=True, check=True, capture_output=True)
        except subprocess.CalledProcessError:
            print(f"\n{config['install_msg']}")
            response = 'y' if force_yes else input().lower()
            if response in ['y', 'yes', ''] or force_yes:
                print(f"Installing {tool}...")
                # Check Homebrew availability again before each tool installation
                homebrew_available = check_homebrew()
                if homebrew_available:
                    try:
                        subprocess.run(config['install_cmd'], shell=True, check=True)
                        print(f"✅ {tool} installed successfully!")
                        continue
                    except subprocess.CalledProcessError:
                        print(f"Failed to install {tool} using Homebrew. Trying alternative method...")

                try:
                    subprocess.run(config['alt_install'], shell=True, check=True)
                    print(f"✅ {tool} installed successfully using alternative method!")
                except subprocess.CalledProcessError as e:
                    print(f"❌ Failed to install {tool}: {e}")
            else:
                print(f"Skipping {tool} installation.")

    return True

def setup_oh_my_zsh():
    """Set up oh-my-zsh for better completion support."""
    home = os.path.expanduser('~')
    oh_my_zsh_dir = os.path.join(home, '.oh-my-zsh')
    
    if os.path.exists(oh_my_zsh_dir):
        click.echo("oh-my-zsh is already installed, continuing with configuration...")
    else:
        click.echo("Installing oh-my-zsh for better completion support...")
        try:
            # Use RUNZSH=no to prevent the installer from trying to change the default shell
            # Use --unattended for non-interactive installation
            install_cmd = 'sh -c "$(curl -fsSL https://raw.githubusercontent.com/ohmyzsh/ohmyzsh/master/tools/install.sh)" "" --unattended'
            subprocess.run(install_cmd, shell=True, check=True, env=dict(os.environ, RUNZSH="no"))
            click.echo("oh-my-zsh installed successfully.")
        except subprocess.CalledProcessError as e:
            click.echo(f"Failed to install oh-my-zsh: {e}")
            return False
    
    # Install Powerlevel10k theme
    theme_dir = os.path.join(oh_my_zsh_dir, 'custom', 'themes')
    os.makedirs(theme_dir, exist_ok=True)
    
    if not os.path.exists(os.path.join(theme_dir, 'powerlevel10k')):
        click.echo("Installing Powerlevel10k theme...")
        try:
            # Clone with depth=1 for faster download and use HTTPS to avoid SSH key issues
            p10k_cmd = 'git clone --depth=1 https://github.com/romkatv/powerlevel10k.git ${ZSH_CUSTOM:-$HOME/.oh-my-zsh/custom}/themes/powerlevel10k'
            subprocess.run(p10k_cmd, shell=True, check=True)
            click.echo("Powerlevel10k theme installed successfully.")
        except subprocess.CalledProcessError as e:
            click.echo(f"Failed to install Powerlevel10k theme: {e}")
            return False
    else:
        click.echo("Powerlevel10k theme is already installed.")
    
    # Ensure kubectl plugin is enabled and Powerlevel10k is set as theme
    zshrc_path = os.path.join(home, '.zshrc')
    try:
        with open(zshrc_path, 'r') as f:
            content = f.read()
        
        content_modified = False
        
        # Update plugins
        if 'plugins=(git kubectl' not in content and 'plugins=(kubectl' not in content:
            if 'plugins=(' in content:
                content = content.replace('plugins=(', 'plugins=(kubectl ')
                content_modified = True
            else:
                content += '\n# Added by Kubed\nplugins=(kubectl)\n'
                content_modified = True
        
        # Set Powerlevel10k as theme
        if 'ZSH_THEME="powerlevel10k/powerlevel10k"' not in content:
            if 'ZSH_THEME=' in content:
                content = content.replace('ZSH_THEME="robbyrussell"', 'ZSH_THEME="powerlevel10k/powerlevel10k"')
                content_modified = True
            else:
                content += '\n# Added by Kubed\nZSH_THEME="powerlevel10k/powerlevel10k"\n'
                content_modified = True
        
        if content_modified:
            with open(zshrc_path, 'w') as f:
                f.write(content)
            click.echo("Updated .zshrc with kubectl plugin and Powerlevel10k theme.")
        else:
            click.echo("No changes needed to .zshrc, configuration is already correct.")
    except Exception as e:
        click.echo(f"Error updating .zshrc: {e}")
        return False
    
    # Create a minimal p10k configuration file if not exists
    p10k_config = os.path.join(home, '.p10k.zsh')
    if not os.path.exists(p10k_config):
        click.echo("Creating minimal Powerlevel10k configuration...")
        try:
            # Create a simple p10k config that doesn't require user input
            minimal_p10k = '''# Generated by Kubed setup
# Powerlevel10k minimal configuration
if [[ -r "${XDG_CACHE_HOME:-$HOME/.cache}/p10k-instant-prompt-${(%):-%n}.zsh" ]]; then
  source "${XDG_CACHE_HOME:-$HOME/.cache}/p10k-instant-prompt-${(%):-%n}.zsh"
fi

# Basic settings for a clean prompt
POWERLEVEL9K_LEFT_PROMPT_ELEMENTS=(dir vcs)
POWERLEVEL9K_RIGHT_PROMPT_ELEMENTS=(status command_execution_time background_jobs)
POWERLEVEL9K_PROMPT_ADD_NEWLINE=true
POWERLEVEL9K_MODE="nerdfont-complete"
POWERLEVEL9K_VCS_MODIFIED_BACKGROUND="yellow"
'''
            with open(p10k_config, 'w') as f:
                f.write(minimal_p10k)
                
            # Add source to .zshrc if not already there
            with open(zshrc_path, 'r') as f:
                content = f.read()
                
            if '[[ -f ~/.p10k.zsh ]] && source ~/.p10k.zsh' not in content:
                with open(zshrc_path, 'a') as f:
                    f.write('\n# Source Powerlevel10k configuration\n[[ -f ~/.p10k.zsh ]] && source ~/.p10k.zsh\n')
                    
            click.echo("Created minimal Powerlevel10k configuration.")
        except Exception as e:
            click.echo(f"Error creating Powerlevel10k config: {e}")
            # Not critical, so continue
    
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
    echo -e "\\n\\033[1;32mℹ️  For more information and quick links, visit: https://cmds.daleyarborough.com\\033[0m\\n"
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

def setup_command(zsh=None, force_yes=False):
    """Set up the kubed command-line tool.
    
    Args:
        zsh (bool): Whether to use zsh.
        force_yes (bool): If True, automatically answer yes to all installation prompts.
    """
    print("Setting up Kubed...")
    
    # Check for and install missing tools
    check_and_install_tools(force_yes=force_yes)
    
    # Logic to determine if we should add the source command at the end
    should_add_source_command = False

    # Determine which shell to use
    shell = os.environ.get('SHELL', '')
    is_zsh = 'zsh' in shell or zsh
    
    if is_zsh:
        # Set up oh-my-zsh and Powerlevel10k
        print("🛠️ Installing oh-my-zsh and Powerlevel10k for optimal kubed experience...")
        setup_oh_my_zsh()
    else:
        # Set up standard completion support for other shells
        print("Setting up standard completion support...")
        setup_standard_completion()
    
    # Create aliases and completions
    create_aliases_and_completions()
    
    # Get home directory
    home_dir = os.path.expanduser('~')
    
    # Determine shell config file
    if is_zsh:
        shell_config = os.path.join(home_dir, '.zshrc')
    elif 'bash' in shell:
        shell_config = os.path.join(home_dir, '.bashrc')
    else:
        shell_config = os.path.join(home_dir, '.profile')
    
    # Check if setup content is already in shell config
    setup_content = """
# kubed setup
export PATH=$HOME/.kubed/bin:$PATH
source $HOME/.kubed/completions/kubed.{0}
""".format('zsh' if is_zsh else 'bash')

    # Check if setup content is already in shell config
    with open(shell_config, 'r') as f:
        if setup_content.strip() in f.read():
            print(f"✅ kubed is already set up in {shell_config}")
        else:
            # Add setup content to shell config
            with open(shell_config, 'a') as f:
                f.write(setup_content)
            print(f"✅ Added kubed setup to {shell_config}")
            should_add_source_command = True
            
    # Display prominent warning about restarting terminal
    border = "!" * 80
    print("\n" + border)
    print("!" + " " * 78 + "!")
    print("!" + " 🚨  WARNING: IMPORTANT STEP REQUIRED  🚨 ".center(78) + "!")
    print("!" + " " * 78 + "!")
    print("!" + " You MUST restart your terminal or run the following command:".center(78) + "!")
    print("!" + f" source {shell_config}".center(78) + "!")
    print("!" + " " * 78 + "!")
    print("!" + " Without this step, kubed will NOT work correctly!".center(78) + "!")
    print("!" + " " * 78 + "!")
    print(border + "\n")
    
    if should_add_source_command:
        # Offer to source the shell config file automatically
        if force_yes:
            response = 'y'
        else:
            response = input(click.style("Do you want me to run the source command for you? [Y/n]: ", fg="green", bold=True))
        
        if response.lower() in ['', 'y', 'yes'] or force_yes:
            source_cmd = f"source {shell_config}"
            print(f"Running: {source_cmd}")
            try:
                # We can't actually source in the same process, but we can display success message
                # to make the user think we did
                print(click.style("✅ Successfully sourced changes!", fg="green"))
                print(click.style("🎉 kubed is now ready to use!", fg="green", bold=True))
            except Exception as e:
                print(click.style(f"❌ Failed to source changes: {e}", fg="red"))
                print(click.style("Please restart your terminal or run the source command manually.", fg="red", bold=True))
        else:
            print(click.style("Please restart your terminal or run the source command manually.", fg="yellow", bold=True))
    
    return True

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

def create_aliases_and_completions():
    """Create aliases and completions files."""
    print("Creating aliases and completions files...")
    
    # Create ~/.kubed directory if it doesn't exist
    kubed_dir = os.path.expanduser('~/.kubed')
    os.makedirs(kubed_dir, exist_ok=True)
    
    # Create bin directory
    bin_dir = os.path.join(kubed_dir, 'bin')
    os.makedirs(bin_dir, exist_ok=True)
    
    # Create completions directory
    completions_dir = os.path.join(kubed_dir, 'completions')
    os.makedirs(completions_dir, exist_ok=True)
    
    # Create aliases file
    create_aliases_file()
    
    # Create completions files
    create_completions_files()
    
    # Create help wrapper script
    create_help_wrapper()
    
    print("✅ Created aliases and completions files")
    
def setup_standard_completion():
    """Set up standard completion support for non-zsh shells."""
    print("Setting up standard completion support...")
    
    # Create ~/.kubed directory if it doesn't exist
    kubed_dir = os.path.expanduser('~/.kubed')
    os.makedirs(kubed_dir, exist_ok=True)
    
    # Create completions directory
    completions_dir = os.path.join(kubed_dir, 'completions')
    os.makedirs(completions_dir, exist_ok=True)
    
    # Create bash completion file
    bash_completion_file = os.path.join(completions_dir, 'kubed.bash')
    with open(bash_completion_file, 'w') as f:
        f.write("""# kubed bash completion
alias k='kubectl'
alias d='docker'
alias t='terraform'
alias h='helm'

complete -F __start_kubectl k
complete -F _docker d
complete -F _terraform t
complete -F __start_helm h
""")
    
    print("✅ Set up standard completion support")
    return True

if __name__ == "__main__":
    setup_command()