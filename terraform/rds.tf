# https://github.com/terraform-aws-modules/terraform-aws-rds/tree/master/examples/complete-postgres

################################################################################
# Supporting Resources
################################################################################

module "rds_sg" {
  source  = "terraform-aws-modules/security-group/aws"
  version = "~> 4"

  name        = "cq-provider-aws-rds-sg"
  description = "cq-provider-aws RDS Security Group"
  vpc_id      = module.vpc.vpc_id

}

################################################################################
# RDS Module
################################################################################


module "rds" {
  source  = "terraform-aws-modules/rds/aws"

  identifier = "cq-provider-aws-test"

  create_db_option_group    = false
  create_db_parameter_group = false

  # All available versions: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_PostgreSQL.html#PostgreSQL.Concepts
  engine               = "postgres"
  engine_version       = "13.4"
  family               = "postgres11" # DB parameter group
  major_engine_version = "13"         # DB option group
  instance_class       = "db.t3.micro"

  allocated_storage = 20

  # NOTE: Do NOT use 'user' as the value for 'username' as it throws:
  # "Error creating DB Instance: InvalidParameterValue: MasterUsername
  # user cannot be used as it is a reserved word used by the engine"
  name                   = "completePostgresql"
  username               = "complete_postgresql"
  create_random_password = true
  random_password_length = 12
  port                   = 5432

  subnet_ids             = module.vpc.private_subnets
  vpc_security_group_ids = [module.rds_sg.security_group_id]

  maintenance_window = "Mon:00:00-Mon:03:00"
  backup_window      = "03:00-06:00"

  backup_retention_period = 0
}

