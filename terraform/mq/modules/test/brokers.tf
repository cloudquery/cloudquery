module "vpc" {
  source  = "terraform-aws-modules/vpc/aws"
  version = "~> 3.0"

  name = "${var.prefix}-mq"
  cidr = "10.99.0.0/18"

  azs             = ["us-east-1a", "us-east-1b"]
  // public_subnets   = ["10.99.0.0/24", "10.99.1.0/24", "10.99.2.0/24"]
  private_subnets  = ["10.99.3.0/24", "10.99.4.0/24"]
  database_subnets = ["10.99.7.0/24", "10.99.8.0/24"]

  create_database_subnet_group       = true
  create_database_subnet_route_table = true

  enable_nat_gateway     = false
  create_egress_only_igw = false
}

module "security_group" {
  source  = "terraform-aws-modules/security-group/aws"
  version = "~> 4.0"

  name        = "${var.prefix}-mq"
  description = "${var.prefix}-mq security group"
  vpc_id      = module.vpc.vpc_id

  tags = var.tags
}

module "mq_broker" {
  source  = "cloudposse/mq-broker/aws"
  version = "0.15.0"

  namespace                  = "${var.prefix}-mq"
  stage                      = "test"
  name                       = "${var.prefix}-mq"
  apply_immediately          = true
  auto_minor_version_upgrade = true
  deployment_mode            = "ACTIVE_STANDBY_MULTI_AZ"
  engine_type                = "ActiveMQ"
  engine_version             = "5.15.15"
  host_instance_type         = "mq.t3.micro"
  publicly_accessible        = false
  general_log_enabled        = true
  audit_log_enabled          = true
  encryption_enabled         = true
  use_aws_owned_key          = true
  vpc_id                     = module.vpc.vpc_id
  subnet_ids                 = module.vpc.private_subnets
  // security_groups            = module.deny_all_sg.security_group_id
}
