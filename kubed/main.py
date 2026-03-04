@click.group()
def cli():
    """Kubed: A CLI tool for Docker, Kubernetes, Terraform, and Helm."""
    pass

@cli.command()
@click.option('--zsh', is_flag=True, help='Use ZSH configuration.')
@click.option('--force-yes', is_flag=True, help='Automatically answer yes to all installation prompts.')
def setup(zsh, force_yes):
    """Set up the kubed command-line tool."""
    from kubed.cli import setup_command
    setup_command(zsh=zsh, force_yes=force_yes)

@cli.command()
@click.argument('command', nargs=-1)
def docker(command):
    """Run docker commands."""
    from kubed.docker import docker_command
    docker_command(command)

@cli.command()
@click.argument('command', nargs=-1)
def k8s(command):
    """Run kubernetes (kubectl) commands."""
    from kubed.k8s import k8s_command
    k8s_command(command)

@cli.command()
@click.argument('command', nargs=-1)
def terraform(command):
    """Run terraform commands."""
    from kubed.terraform import terraform_command
    terraform_command(command)

@cli.command()
@click.argument('command', nargs=-1)
def helm(command):
    """Run helm commands."""
    from kubed.helm import helm_command
    helm_command(command)

if __name__ == '__main__':
    cli() 