#!/usr/bin/env python3
"""
Kubed - Command-line tools autocompletion and productivity enhancer.
"""

import click
from kubed.cli import setup_command, completions_path_command, aliases_path_command

@click.group()
def cli():
    """Kubed - Command-line tools autocompletion and productivity enhancer."""
    pass

@cli.command('setup')
def setup():
    """Set up Kubed on your system."""
    setup_command()

@cli.command('completions-path')
def completions_path():
    """Get the path to the shell completions."""
    completions_path_command()

@cli.command('aliases-path')
def aliases_path():
    """Get the path to the shell aliases."""
    aliases_path_command()

if __name__ == '__main__':
    cli()
