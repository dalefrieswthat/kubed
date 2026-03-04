#!/bin/bash
# Build script for Kubed

set -e

echo "Building Kubed package..."

# Check if virtual environment exists
if [ ! -d "venv" ]; then
    echo "Creating virtual environment..."
    python3 -m venv venv
fi

# Activate virtual environment
source venv/bin/activate

# Install requirements
echo "Installing dependencies..."
pip install -r requirements.txt

# Install the package in development mode
echo "Installing package in development mode..."
pip install -e .

# Create required directories if they don't exist
mkdir -p kubed/completions/bash kubed/completions/zsh kubed/aliases

echo "Kubed package built successfully!"
echo "You can install it with: pip install ."
echo "Or activate the development environment with: source venv/bin/activate"
echo "And run the setup command with: kubed-setup"

# Deactivate virtual environment
deactivate
