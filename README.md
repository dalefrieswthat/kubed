# Kubed Documentation

This is the documentation website for the Kubed CLI tool, built with Jekyll and Tailwind CSS.

## Setup

### Prerequisites

- Ruby 2.7+ with Bundler
- Node.js 14+ with npm

### Installation

1. Install Ruby dependencies:

```bash
bundle install
```

2. Install Node.js dependencies:

```bash
npm install
```

### Development

To run the site locally with live reload:

```bash
npm run dev
```

This will:
- Build the Tailwind CSS
- Watch for CSS changes
- Run the Jekyll server

You can then access the site at http://localhost:4000

### Building for Production

To build the site for production:

```bash
npm run build:css
bundle exec jekyll build
```

The site will be generated in the `_site` directory.

## Customization

### Tailwind Configuration

The Tailwind CSS configuration is located in `tailwind.config.js`. You can modify this file to customize the theme.

### Layout Files

The main layout templates are in the `_layouts` directory. The `default.html` layout is the main template.

### Content Files

- `index.md`: The home page
- `/docker/`, `/kubernetes/`, `/terraform/`, `/helm/`: Tool-specific documentation
- `installation.md`: Installation instructions

## Implementation Notes

This site uses:

- **Jekyll** as the static site generator
- **Tailwind CSS** for styling
- **PostCSS** for processing CSS
- **Font Awesome** for icons

The site is designed to be responsive and follows modern web development practices. 