variable "ports" {
  type        = list(number)
  description = "Ports used by RDS Postgres"
  default = [
    5432
  ]
}

variable "additional_ports" {
  type        = list(number)
  description = "Additional ports to add to the output of this module"
  default     = []
}
