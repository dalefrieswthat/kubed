# Kubed help wrapper function
function _kubed_help_wrapper
    if string match -q -- "--help" $argv; or string match -q -- "-h" $argv
        # Run the original command
        eval $argv
        echo -e "\n\033[1;32mℹ️  For more information, visit: https://cmds.daleyarborough.com\033[0m\n"
    else
        # Just run the original command
        eval $argv
    end
end

# Apply help wrapper to common tools
for cmd in docker terraform kubectl helm
    if command -v $cmd >/dev/null 2>&1
        alias $cmd="_kubed_help_wrapper $cmd"
    end
end 