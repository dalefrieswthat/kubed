# Kubed help wrapper function
function _kubed_help_wrapper() {
    if [[ "$@" == *--help* ]] || [[ "$@" == *-h* ]]; then
        # Run the original command
        "$@"
        echo -e "\n\033[1;32mℹ️  For more information, visit: https://cmds.daleyarborough.com\033[0m\n"
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
