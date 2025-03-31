---
layout: default
title: Docker Commands
nav_order: 2
description: "Common Docker commands and best practices"
permalink: /docker/
---

# Docker Commands

This section provides a comprehensive reference for Docker commands, including common use cases and best practices.

## Basic Commands

### Container Management
```bash
# Run a container
docker run [OPTIONS] IMAGE [COMMAND]

# List running containers
docker ps

# List all containers (including stopped)
docker ps -a

# Stop a container
docker stop CONTAINER_ID

# Remove a container
docker rm CONTAINER_ID

# Remove a running container (force)
docker rm -f CONTAINER_ID
```

### Image Management
```bash
# List images
docker images

# Pull an image
docker pull IMAGE_NAME

# Build an image
docker build -t IMAGE_NAME .

# Remove an image
docker rmi IMAGE_NAME
```

## Common Use Cases

### Running a Web Application
```bash
# Run a web application with port mapping
docker run -d -p 8080:80 nginx

# Run with environment variables
docker run -d -e DB_HOST=db -e DB_USER=user myapp
```

### Working with Volumes
```bash
# Create a named volume
docker volume create my_volume

# Mount a volume
docker run -v my_volume:/data myapp

# Mount a host directory
docker run -v /host/path:/container/path myapp
```

### Networking
```bash
# Create a network
docker network create my_network

# Connect a container to a network
docker network connect my_network container_id

# List networks
docker network ls
```

## Best Practices

1. **Use Specific Tags**: Always use specific version tags instead of `latest`
   ```bash
   docker pull nginx:1.25.3
   ```

2. **Clean Up Resources**: Regularly remove unused containers, images, and volumes
   ```bash
   # Remove all stopped containers
   docker container prune
   
   # Remove all unused images
   docker image prune
   
   # Remove all unused volumes
   docker volume prune
   ```

3. **Use Multi-stage Builds**: For smaller production images
   ```dockerfile
   # Build stage
   FROM node:16 AS builder
   WORKDIR /app
   COPY package*.json ./
   RUN npm install
   COPY . .
   RUN npm run build

   # Production stage
   FROM nginx:alpine
   COPY --from=builder /app/build /usr/share/nginx/html
   EXPOSE 80
   CMD ["nginx", "-g", "daemon off;"]
   ```

## Troubleshooting

### Common Issues

1. **Container Won't Start**
   ```bash
   # Check container logs
   docker logs CONTAINER_ID
   
   # Check container details
   docker inspect CONTAINER_ID
   ```

2. **Port Conflicts**
   ```bash
   # Check which process is using a port
   lsof -i :8080
   ```

3. **Permission Issues**
   ```bash
   # Run container with correct user
   docker run -u $(id -u):$(id -g) myapp
   ```

## Additional Resources

- [Docker Official Documentation](https://docs.docker.com/)
- [Docker Hub](https://hub.docker.com/)
- [Docker Compose Documentation](https://docs.docker.com/compose/) 