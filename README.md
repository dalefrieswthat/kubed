# kubed

**kubed** is a CLI productivity tool with autocompletion for Docker, Terraform, Helm, and Kubernetes.

## About

`kubed` enhances your terminal experience by providing powerful autocompletion and useful aliases for cloud-native development tools. It aims to streamline your workflow by reducing the need to remember complex command syntax and improving efficiency with quick access to frequently used commands.

### Key Features

- **Autocompletion**: Provides comprehensive autocompletion for Docker, Terraform, Helm, and Kubernetes (kubectl) commands.
- **Aliases**: Includes useful aliases like `k` for `kubectl`, `d` for `docker`, etc., to speed up command execution.
- **Cross-Platform Support**: Works seamlessly on macOS and Linux.
- **Enhanced Help Output**: Offers improved help output with links to documentation for better understanding and usage.

## Installation

To install the `kubed` package, you can use `pip`, the Python package installer. Make sure you have Python 3.6 or later installed on your system.

```bash
pip install kubed
```

## Usage

Once installed, you can use the following commands:

- **Setup Command:**
  ```bash
  kubed-setup
  ```

- **Get Completions Path:**
  ```bash
  kubed-completions-path
  ```

- **Get Aliases Path:**
  ```bash
  kubed-aliases-path
  ```

## Requirements

- Python 3.6 or later
- The following Python packages:
  - `click>=7.0`
  - `requests`

## License

This project is licensed under the MIT License. See the LICENSE file for more details.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request for any improvements or bug fixes.

## Contact

For more information, visit [Dale Yarborough's website](https://daleyarborough.com) or contact via email at info@daleyarborough.com. 