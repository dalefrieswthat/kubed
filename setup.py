from setuptools import setup, find_packages
import os
import sys
import subprocess

def post_install():
    # Create a very noticeable terminal message
    terminal_width = 80
    try:
        # Try to get actual terminal width if possible
        terminal_width = os.get_terminal_size().columns
    except:
        pass
    
    border = "!" * terminal_width
    
    print("\n\n")
    print(border)
    print(border)
    print("")
    print(" 🚨 IMPORTANT: RESTART YOUR TERMINAL OR RUN THE FOLLOWING COMMAND:".center(terminal_width))
    print("")
    print(" source ~/.zshrc".center(terminal_width))
    print("")
    print(" You must do this for kubed to work properly!".center(terminal_width))
    print("")
    print(border)
    print(border)
    print("\n\n")

if __name__ == "__main__":
    setup(
        name="kubed",
        version="2.2.0",
        packages=find_packages(),
        include_package_data=True,
        install_requires=[
            "setuptools>=42.0.0",
            "click",
            "rich",
            "pyyaml",
            "requests",
            "docker",
            "kubernetes",
            "helm",
        ],
        entry_points={
            "console_scripts": [
                "kubed-setup=kubed.cli:setup_command",
                "kubed=kubed.cli:cli",
            ],
        },
        author="Dale Yarborough",
        author_email="daleyarborough@gmail.com",
        description="A CLI tool for managing Kubernetes, Docker, Terraform, and Helm workflows",
        long_description=open("README.md").read(),
        long_description_content_type="text/markdown",
        url="https://cmds.daleyarborough.com",
        classifiers=[
            "Development Status :: 5 - Production/Stable",
            "Intended Audience :: Developers",
            "License :: OSI Approved :: MIT License",
            "Operating System :: OS Independent",
            "Programming Language :: Python :: 3",
            "Programming Language :: Python :: 3.6",
            "Programming Language :: Python :: 3.7",
            "Programming Language :: Python :: 3.8",
            "Programming Language :: Python :: 3.9",
            "Programming Language :: Python :: 3.10",
            "Programming Language :: Python :: 3.11",
            "Topic :: Software Development :: Libraries :: Python Modules",
        ],
        python_requires=">=3.6",
    )
    # Run post-install message
    if "install" in sys.argv or "develop" in sys.argv:
        post_install()