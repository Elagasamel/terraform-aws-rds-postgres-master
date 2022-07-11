<!-- BEGINNING OF PRE-COMMIT-TERRAFORM DOCS HOOK -->
## Requirements

No requirements.

## Providers

No provider.

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| ingress\_cidr\_blocks | CIDR blocks to attach to security groups for ingress | `list(string)` | n/a | yes |
| name\_prefix | A string to prepend to names of resources created by this example | `any` | n/a | yes |
| subnet\_ids | List of at least 2 subnets in different AZs for DB subnet group | `list(string)` | n/a | yes |
| vpc\_id | VPC ID of network. | `string` | n/a | yes |
| egress\_cidr\_blocks | CIDR blocks to attach to security groups for egress | `list(string)` | <pre>[<br>  "0.0.0.0/0"<br>]</pre> | no |
| tags | A map of tags to add to all resources created by this example. | `map(string)` | <pre>{<br>  "Author": "Tamr",<br>  "Environment": "Example"<br>}</pre> | no |

## Outputs

| Name | Description |
|------|-------------|
| rds | n/a |

<!-- END OF PRE-COMMIT-TERRAFORM DOCS HOOK -->
