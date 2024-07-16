# openssd
Openssd stands for Open (source) Super Simple Diary. Openssd is a plug-n-run project to host a version of yourself on the webternet.

## Features

- Markdown-based articles stored in a Git repository
- Responsive design using Tailwind CSS
- Dynamic content loading with HTMX
- Docker-based deployment for easy setup

## Prerequisites

- Git
- Docker and Docker Compose (installed automatically by the installation script)

## Installation

### Linux

1. Open a terminal
2. Run the following command:

```bash
curl -sSL https://raw.githubusercontent.com/yourusername/blog-project/main/scripts/install_linux.sh | bash
```

### macOS

1. Open a terminal
2. Run the following command:

```bash
curl -sSL https://raw.githubusercontent.com/yourusername/blog-project/main/scripts/install_macos.sh | bash
```

## Usage

After installation, your blog will be accessible at `http://localhost:8080`.

To add or modify articles:

1. Clone your articles repository (the one specified in the `GIT_REPO_URL` environment variable)
2. Add or edit Markdown files in the `articles` directory
3. Commit and push your changes
4. The blog will automatically update with the new content

## Customization

### Changing the theme

To change the blog's appearance:

1. Edit the HTML templates in the `ui/templates` directory
2. Modify the Tailwind CSS classes to adjust the styling

### Adding new pages

To add new pages to your blog:

1. Create a new HTML template in the `ui/templates` directory
2. Add a new route handler in the `go/client/main.go` file
3. Update the navigation menu in the templates to include the new page

## Troubleshooting

If you encounter any issues, please check the following:

1. Ensure Docker is running
2. Check if the containers are up with `docker-compose ps`
3. View the logs with `docker-compose logs`

For further assistance, please open an issue on the GitHub repository.

## License

This project is licensed under the MIT License - see the LICENSE file for details.