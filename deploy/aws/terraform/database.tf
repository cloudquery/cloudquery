resource "aws_rds_cluster" "cloudquery" {
  cluster_identifier = "cloudquery"
  engine             = "aurora-postgresql"
  engine_version     = "10.12"
  engine_mode        = "serverless"
  master_username    = "cloudquery"
  database_name      = "cloudquery"
  master_password    = random_password.password.result

  db_subnet_group_name = aws_db_subnet_group.rds_subnet_group.name

  vpc_security_group_ids = [aws_security_group.allow_postgresql.id]

  skip_final_snapshot = true
  apply_immediately = true

  scaling_configuration {
    auto_pause               = true
    max_capacity             = 4
    min_capacity             = 2
    seconds_until_auto_pause = 300
    timeout_action           = "ForceApplyCapacityChange"
  }
}

resource "random_password" "password" {
  length           = 16
  special          = true
  override_special = "_%@"
}

resource "aws_db_subnet_group" "rds_subnet_group" {
  name       = "rds_subnet_group"
  subnet_ids = [aws_subnet.rds_subnet_a.id, aws_subnet.rds_subnet_b.id]

  tags = {
    Name = "Cloudquery RDS Subnet group"
  }
}

resource "aws_ssm_parameter" "cloudquery_master_password" {
  name        = "/cloudquery/database/password/master"
  description = "Master password for cloudquery aurora database"
  type        = "SecureString"
  value       = random_password.password.result
}