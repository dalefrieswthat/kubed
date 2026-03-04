---
layout: default
title: Helm Commands
nav_order: 5
description: "Common Helm commands and best practices"
permalink: /helm/
---

# Helm Commands

This section provides a comprehensive reference for Helm commands, including common use cases and best practices.

## Basic Commands

### Repository Management
```bash
# Add a repository
helm repo add bitnami https://charts.bitnami.com/bitnami

# List repositories
helm repo list

# Update repositories
helm repo update

# Remove repository
helm repo remove bitnami
```

### Chart Management
```bash
# Search for charts
helm search repo nginx

# Show chart information
helm show all bitnami/nginx

# Show chart values
helm show values bitnami/nginx

# Install chart
helm install my-release bitnami/nginx

# Upgrade release
helm upgrade my-release bitnami/nginx

# Uninstall release
helm uninstall my-release
```

## Common Use Cases

### Working with Values
```bash
# Install with custom values
helm install my-release bitnami/nginx -f values.yaml

# Install with set values
helm install my-release bitnami/nginx --set service.type=NodePort

# Install with multiple value files
helm install my-release bitnami/nginx -f values.yaml -f prod-values.yaml
```

### Release Management
```bash
# List releases
helm list

# List all releases (including failed)
helm list --all

# Get release status
helm status my-release

# Rollback release
helm rollback my-release 1
```

### Chart Development
```bash
# Create new chart
helm create my-chart

# Package chart
helm package my-chart

# Lint chart
helm lint my-chart

# Test chart
helm test my-release
```

## Best Practices

1. **Version Management**: Always specify chart versions
   ```bash
   # Install specific version
   helm install my-release bitnami/nginx --version 13.2.23
   ```

2. **Value Organization**: Use separate value files
   ```yaml
   # values.yaml
   service:
     type: ClusterIP
     port: 80

   # prod-values.yaml
   service:
     type: LoadBalancer
     port: 443
   ```

3. **Release Naming**: Use meaningful release names
   ```bash
   # Good example
   helm install prod-nginx bitnami/nginx

   # Bad example
   helm install release1 bitnami/nginx
   ```

## Troubleshooting

### Common Issues

1. **Release Failed**
   ```bash
   # Check release status
   helm status my-release
   
   # Check release history
   helm history my-release
   ```

2. **Chart Dependencies**
   ```bash
   # Update dependencies
   helm dependency update my-chart
   
   # Build dependencies
   helm dependency build my-chart
   ```

3. **Repository Issues**
   ```bash
   # Update repository cache
   helm repo update
   
   # Check repository status
   helm repo list
   ```

## Additional Resources

- [Helm Official Documentation](https://helm.sh/docs/)
- [Helm Hub](https://hub.helm.sh/)
- [Helm Best Practices](https://helm.sh/docs/chart_best_practices/) 