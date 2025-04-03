# Kubed Fish completions

# Function to check if a command exists
function _kubed_command_exists
    command -v $argv[1] >/dev/null 2>&1
end

# Function to source kubectl completion
function _kubed_source_kubectl
    if _kubed_command_exists kubectl
        kubectl completion fish | source
        complete -c k -w kubectl
    end
end

# Function to source docker completion
function _kubed_source_docker
    if _kubed_command_exists docker
        docker completion fish | source
        complete -c d -w docker
    end
end

# Function to source terraform completion
function _kubed_source_terraform
    if _kubed_command_exists terraform
        terraform -generate-autocomplete
        complete -c tf -w terraform
    end
end

# Function to source helm completion
function _kubed_source_helm
    if _kubed_command_exists helm
        helm completion fish | source
        complete -c h -w helm
    end
end

# Main function to initialize all completions
function _kubed_init_completions
    _kubed_source_kubectl
    _kubed_source_docker
    _kubed_source_terraform
    _kubed_source_helm
end

# Initialize completions
_kubed_init_completions 