resource "aws_db_parameter_group" "db_parameter_group" {
  name   = "rds-db-pg-${var.test_prefix}-${var.test_suffix}"
  family = "mysql8.0"
  description = "Test RDS DB parameter group"

  parameter {
    name  = "character_set_server"
    value = "utf8"
  }

  parameter {
    name  = "character_set_client"
    value = "utf8"
  }
}
