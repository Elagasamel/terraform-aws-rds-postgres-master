variable "name_prefix" {
  type        = string
  description = "Identifier prefix for the resources"
}

variable "vpc_cidr" {
  type        = string
  description = "CIDR for VPC to be created"
}

variable "database_subnets" {
  description = "CIDR blocks to attach to use in subnets"
  type        = list(string)
}

variable "egress_cidr_blocks" {
  description = "CIDR blocks to attach to security groups for egress"
  type        = list(string)
}

variable "ingress_cidr_blocks" {
  description = "CIDR blocks to attach to security groups for egress"
  type        = list(string)
}

variable "tags" {
  type = map(string)
  default = {
    Terraform   = "true"
    Terratest   = "true"
    Environment = "test"
  }
}
