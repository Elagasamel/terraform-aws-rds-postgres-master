# Tamr Terraform Template Repo

## v3.1.2 - May 11th 2022
* Adds idle_in_transaction_session_timeout to parameter group

## v3.1.1 - February 16th 2022
* Updates version file to prevent the major upgrade to the AWS provider version 4.0.
* Adds new extensions to be ignored in the ".gitignore" file.

## v3.1.0 - November 3rd 2021
* Adds variables `enabled_cloudwatch_logs_exports`, `param_log_min_duration_statement` and `param_log_statement` configuring the logs and its publishing to CloudWatch

## v3.0.0 - July 27th 2021
* Updates default engine version to not specify a minor version
* Adds new variable `auto_minor_version_upgrade` to allow AWS to upgrade RDS minor versions

## v2.1.0 - July 12nd 2021
* Adds tags for RDS Subnet Group.
* Adds new variable `tags` to set tags for all resources
* Deprecates `additional_tags` in favor of `tags`

## v2.0.0 - June 30th 2021
* Accepts a list of security groups
* Returns a list of ports used by RDS
* Removes ability for the creation of security groups

## v1.0.0 - April 12th 2021
* Updates minimum Terraform version to 13
* Updates minimum AWS provider version to 3.36.0
* Adds explicit type to `additional_tags` variable

## v0.4.1 - Nov 6th 2020
* Adds input `db_port` to configure port that postgres database accepts connections on
* Fixes example so it accepts `ingress_sg_ids` as input

## v0.4.0 - Oct 27th 2020
* Consolidates inputs `tamr_vm_sg_id` and `spark_cluster_sg_ids` into one input, `ingress_sg_ids`

## v0.3.1 - Sep 10th 2020
* Adds outputs, `rds_username` and `rds_dbname`

## v0.3.0 - Sep 10th 2020
* Adds creation of aws_db_subnet_group resource
  * Adds variable rds_subnet_ids to specify subnet IDs in subnet group
* Renames variable subnet_name to subnet_group_name

## v0.2.1 - Jun 22nd 2020
* Adds variable "engine_version" to set the postgres version
* Formatting and documentation updates

## v0.1.0 - Feb 25th 2020
* Initing project
