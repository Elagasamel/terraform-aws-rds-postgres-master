This deploys:
- a VPC, using the AWS VPC Module. (Can change to tamr-networking later)
- the examples/minimal resources.

Go code (Terratest) runs this and makes sure our examples are working.
The validations are:
- Address
- Port
- Security Group
- username
- dbname

<!-- BEGINNING OF PRE-COMMIT-TERRAFORM DOCS HOOK -->
## Requirements

No requirements.

## Providers

No provider.

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| identifier\_prefix | Identifier prefix for the resources | `string` | n/a | yes |
| parameter\_group\_name | Name of the parameter group | `string` | n/a | yes |
| pg\_password | Password for postgres | `string` | n/a | yes |
| pg\_username | Username for postgres | `string` | n/a | yes |
| postgres\_db\_name | Name of the postgres db | `string` | n/a | yes |

## Outputs

No output.

<!-- END OF PRE-COMMIT-TERRAFORM DOCS HOOK -->
