from setuptools import setup, find_packages

setup(
    name="kubed",
    version="2.1.6",
    packages=find_packages(),
    include_package_data=True,
    install_requires=[
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