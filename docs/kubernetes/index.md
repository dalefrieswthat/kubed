---
layout: default
title: Kubernetes Commands
nav_order: 3
description: "Common kubectl commands and best practices"
permalink: /kubernetes/
---

# Kubernetes Commands

This section provides a comprehensive reference for kubectl commands, including common use cases and best practices.

## Basic Commands

### Resource Management
```bash
# Get resources
kubectl get pods
kubectl get services
kubectl get deployments
kubectl get nodes

# Get resources in all namespaces
kubectl get all --all-namespaces

# Get detailed information about a resource
kubectl describe pod POD_NAME
kubectl describe service SERVICE_NAME
kubectl describe node NODE_NAME

# Delete resources
kubectl delete pod POD_NAME
kubectl delete service SERVICE_NAME
kubectl delete deployment DEPLOYMENT_NAME
```

### Resource Creation
```bash
# Create resources from file
kubectl apply -f manifest.yaml

# Create resources from URL
kubectl apply -f https://raw.githubusercontent.com/kubernetes/website/main/content/en/examples/application/deployment.yaml

# Create resources from stdin
cat manifest.yaml | kubectl apply -f -
```

## Common Use Cases

### Working with Pods
```bash
# Get pod logs
kubectl logs POD_NAME

# Follow pod logs
kubectl logs -f POD_NAME

# Execute command in pod
kubectl exec -it POD_NAME -- /bin/bash

# Port forwarding
kubectl port-forward POD_NAME 8080:80
```

### Managing Deployments
```bash
# Scale deployment
kubectl scale deployment DEPLOYMENT_NAME --replicas=3

# Update deployment
kubectl set image deployment/DEPLOYMENT_NAME CONTAINER_NAME=IMAGE_NAME

# Rollback deployment
kubectl rollout undo deployment/DEPLOYMENT_NAME

# View deployment status
kubectl rollout status deployment/DEPLOYMENT_NAME
```

### Working with ConfigMaps and Secrets
```bash
# Create ConfigMap from file
kubectl create configmap CONFIG_NAME --from-file=config.yaml

# Create Secret from file
kubectl create secret generic SECRET_NAME --from-file=secret.yaml

# List ConfigMaps and Secrets
kubectl get configmaps
kubectl get secrets
```

## Best Practices

1. **Use Namespaces**: Organize resources logically
   ```bash
   # Create namespace
   kubectl create namespace my-namespace
   
   # Set default namespace
   kubectl config set-context --current --namespace=my-namespace
   ```

2. **Resource Limits**: Always set resource limits
   ```yaml
   resources:
     requests:
       memory: "64Mi"
       cpu: "250m"
     limits:
       memory: "128Mi"
       cpu: "500m"
   ```

3. **Health Checks**: Implement readiness and liveness probes
   ```yaml
   readinessProbe:
     httpGet:
       path: /health
       port: 8080
     initialDelaySeconds: 5
     periodSeconds: 10
   ```

## Troubleshooting

### Common Issues

1. **Pod Won't Start**
   ```bash
   # Check pod events
   kubectl describe pod POD_NAME
   
   # Check pod logs
   kubectl logs POD_NAME
   ```

2. **Service Not Accessible**
   ```bash
   # Check service endpoints
   kubectl get endpoints SERVICE_NAME
   
   # Check service details
   kubectl describe service SERVICE_NAME
   ```

3. **Node Issues**
   ```bash
   # Check node status
   kubectl get nodes
   
   # Check node details
   kubectl describe node NODE_NAME
   ```

## Additional Resources

- [Kubernetes Official Documentation](https://kubernetes.io/docs/)
- [kubectl Cheat Sheet](https://kubernetes.io/docs/reference/kubectl/cheatsheet/)
- [Kubernetes Best Practices](https://kubernetes.io/docs/concepts/configuration/overview/) 