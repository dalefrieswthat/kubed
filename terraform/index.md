---
layout: default
title: Terraform Commands
nav_order: 4
description: "Common Terraform commands and best practices"
permalink: /terraform/
---

# Terraform Commands

This section provides a comprehensive reference for Terraform commands, including common use cases and best practices.

## Basic Commands

### Initialization and Planning
```bash
# Initialize Terraform working directory
terraform init

# Format Terraform files
terraform fmt

# Validate Terraform files
terraform validate

# Show execution plan
terraform plan

# Show execution plan with output file
terraform plan -out=tfplan
```

### Resource Management
```bash
# Apply changes
terraform apply

# Apply changes from plan file
terraform apply tfplan

# Destroy infrastructure
terraform destroy

# Show current state
terraform show

# List resources
terraform state list
```

## Common Use Cases

### Working with State
```bash
# Move resource in state
terraform state mv SOURCE_ADDRESS DESTINATION_ADDRESS

# Remove resource from state
terraform state rm ADDRESS

# Import existing resource
terraform import ADDRESS ID
```

### Working with Variables
```bash
# Set variable value
terraform apply -var="instance_type=t2.micro"

# Set variable file
terraform apply -var-file="prod.tfvars"

# Override variable
terraform apply -var="region=us-west-2" -var="environment=staging"
```

### Output Management
```bash
# Show outputs
terraform output

# Show specific output
terraform output instance_ip

# Show outputs in JSON format
terraform output -json
```

## Best Practices

1. **State Management**: Use remote state storage
   ```hcl
   terraform {
     backend "s3" {
       bucket = "my-terraform-state"
       key    = "prod/terraform.tfstate"
       region = "us-west-2"
     }
   }
   ```

2. **Variable Organization**: Use separate variable files
   ```hcl
   # variables.tf
   variable "environment" {
     description = "Environment name"
     type        = string
   }

   # prod.tfvars
   environment = "production"
   ```

3. **Module Usage**: Organize code into modules
   ```hcl
   module "vpc" {
     source = "./modules/vpc"
     
     environment = var.environment
     cidr_block  = var.vpc_cidr
   }
   ```

## Troubleshooting

### Common Issues

1. **State Lock Issues**
   ```bash
   # Force unlock state
   terraform force-unlock LOCK_ID
   ```

2. **Provider Issues**
   ```bash
   # Reinitialize with specific provider
   terraform init -upgrade -plugin-dir=/path/to/provider
   ```

3. **Resource Dependencies**
   ```bash
   # Show dependency graph
   terraform graph
   ```

## Additional Resources

- [Terraform Official Documentation](https://www.terraform.io/docs/)
- [Terraform Best Practices](https://www.terraform.io/docs/cloud/guides/recommended-practices/)
- [Terraform Registry](https://registry.terraform.io/) 