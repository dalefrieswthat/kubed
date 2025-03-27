from setuptools import setup

setup(
    name="kubed",
    version="1.3.0",
    description="CLI productivity tool with autocompletion for Docker, Terraform, Helm, and Kubernetes",
    long_description=open("README.md").read(),
    long_description_content_type="text/markdown",
    author="Dale Yarborough",
    author_email="info@daleyarborough.com",
    url="https://daleyarborough.com",
    packages=["kubed", "kubed.aliases", "kubed.completions.bash", "kubed.completions.zsh"],
    include_package_data=True,
    python_requires=">=3.6",
    install_requires=[
        "click>=7.0",
        "requests",
    ],
    entry_points={
        "console_scripts": [
            "kubed-setup=kubed.cli:setup_command",
            "kubed-completions-path=kubed.cli:completions_path_command",
            "kubed-aliases-path=kubed.cli:aliases_path_command",
        ],
    },
    classifiers=[
        "Development Status :: 3 - Alpha",
        "Intended Audience :: Developers",
        "Intended Audience :: System Administrators",
        "Programming Language :: Python :: 3",
        "Programming Language :: Python :: 3.6",
        "Programming Language :: Python :: 3.7",
        "Programming Language :: Python :: 3.8",
        "Programming Language :: Python :: 3.9",
        "Topic :: Software Development :: Libraries :: Python Modules",
        "Topic :: System :: Systems Administration",
    ],
    license="MIT"
)