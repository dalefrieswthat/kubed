# Kubed Zsh completions

# Function to check if a command exists
_kubed_command_exists() {
    command -v "$1" >/dev/null 2>&1
}

# Function to source kubectl completion
_kubed_source_kubectl() {
    if _kubed_command_exists kubectl; then
        source <(kubectl completion zsh)
        compdef k=kubectl
    fi
}

# Function to source docker completion
_kubed_source_docker() {
    if _kubed_command_exists docker; then
        if [ -f /usr/share/zsh/vendor-completions/_docker ]; then
            source /usr/share/zsh/vendor-completions/_docker
        fi
    fi
}

# Function to source terraform completion
_kubed_source_terraform() {
    if _kubed_command_exists terraform; then
        autoload -U +X bashcompinit && bashcompinit
        complete -o nospace -C $(which terraform) terraform
        complete -o nospace -C $(which terraform) tf
    fi
}

# Function to source helm completion
_kubed_source_helm() {
    if _kubed_command_exists helm; then
        source <(helm completion zsh)
        compdef h=helm
    fi
}

# Main function to initialize all completions
_kubed_init_completions() {
    _kubed_source_kubectl
    _kubed_source_docker
    _kubed_source_terraform
    _kubed_source_helm
}

# Initialize completions
_kubed_init_completions
