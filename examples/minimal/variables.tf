variable "vpc_id" {
  type        = string
  description = "VPC ID of network."
}

variable "subnet_ids" {
  type        = list(string)
  description = "List of at least 2 subnets in different AZs for DB subnet group"
}

variable "name_prefix" {
  description = "A string to prepend to names of resources created by this example"
}

variable "ingress_cidr_blocks" {
  description = "CIDR blocks to attach to security groups for ingress"
  type        = list(string)
}

variable "egress_cidr_blocks" {
  description = "CIDR blocks to attach to security groups for egress"
  type        = list(string)
  default     = ["0.0.0.0/0"]
}

variable "tags" {
  type        = map(string)
  description = "A map of tags to add to all resources created by this example."
  default = {
    Author      = "Tamr"
    Environment = "Example"
  }
}
