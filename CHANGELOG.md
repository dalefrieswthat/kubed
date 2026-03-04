# Changelog

All notable changes to this project will be documented in this file.

## [2.2.0] - 2025-03-31

### Added
- Setuptools as a prerequisite in setup.py
- Added setuptools check in the CLI to ensure it's properly installed
- Added --force-yes flag to setup command for non-interactive installations
- Added __main__.py for running kubed as a module
- Added a detailed installation guide to the documentation
- Comprehensive documentation website with improved UI

### Changed
- Enhanced terminal restart notifications and user feedback during setup
- Automated installation of oh-my-zsh and Powerlevel10k
- Improved error handling during tool installation
- Updated documentation with streamlined installation guide

### Fixed
- Installation issues related to missing setuptools
- Fixed redundant titles in documentation
- Improved CSS styling for better user experience

## [2.1.7] - 2025-03-29

### Added
- Update Terraform to 1.7.5
- Added Terraform autocomplete for kubed
- Custom installation of Terraform using curl
- Homebrew installation for other components

### Changed
- Improved check_and_install_tools for better error handling
- Enhanced error handling for installation
- Support for alternative installation methods when Homebrew fails

### Fixed
- Fixed an issue where Terraform installation might fail
- Fixed autocompletion for various commands
- Improved error handling for missing tools 