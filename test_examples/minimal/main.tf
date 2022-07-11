module "rds_postgres" {
  #source               = "git::https://github.com/Datatamer/terraform-rds-postgres.git?ref=x.y.z"
  source = "../../examples/minimal"

  name_prefix         = var.name_prefix
  ingress_cidr_blocks = var.ingress_cidr_blocks
  tags                = var.tags
  vpc_id              = module.vpc.vpc_id
  subnet_ids          = module.vpc.database_subnets
}

module "vpc" {
  source  = "terraform-aws-modules/vpc/aws"
  version = "3.1.0"

  name = "${var.name_prefix}-terratest-vpc"
  cidr = var.vpc_cidr

  azs              = data.aws_availability_zones.available.names
  database_subnets = var.database_subnets # ["172.18.6.0/24", "172.18.7.0/24"]

  // Remove these comments if you want public access to your RTB
  # public_subnets = [ "172.18.0.0/24", "172.18.1.0/24" ]
  # create_database_internet_gateway_route = true

  create_database_subnet_group       = false
  create_database_subnet_route_table = true

  enable_nat_gateway = false
  enable_vpn_gateway = false

  tags = var.tags
}

data "aws_region" "current" {}

data "aws_availability_zones" "available" {
  state = "available"
}
