module "rds_postgres" {
  # source               = "git::https://github.com/Datatamer/terraform-aws-rds-postgres.git?ref=3.0.0"
  source = "../.."

  identifier_prefix    = "${var.name_prefix}-example-rds-pg-"
  postgres_name        = "example0"
  parameter_group_name = "${var.name_prefix}-example-rds-postgres-pg"
  username             = "exampleUsername"
  password             = "examplePassword" #tfsec:ignore:GEN003

  vpc_id            = var.vpc_id
  subnet_group_name = "${var.name_prefix}_example_subnet_group"
  # Network requirement: DB subnet group needs a subnet in at least two Availability Zones
  rds_subnet_ids     = var.subnet_ids
  security_group_ids = module.rds-postgres-sg.security_group_ids
  tags               = var.tags
}

module "sg-ports" {
  # source               = "git::https://github.com/Datatamer/terraform-aws-rds-postgres.git//modules/rds-postgres-ports?ref=3.0.0"
  source = "../../modules/rds-postgres-ports"
}

module "rds-postgres-sg" {
  source              = "git::git@github.com:Datatamer/terraform-aws-security-groups.git?ref=1.0.0"
  vpc_id              = var.vpc_id
  ingress_cidr_blocks = var.ingress_cidr_blocks
  egress_cidr_blocks  = var.egress_cidr_blocks
  ingress_ports       = module.sg-ports.ingress_ports
  sg_name_prefix      = var.name_prefix
  egress_protocol     = "all"
  ingress_protocol    = "tcp"
  tags                = var.tags
}
