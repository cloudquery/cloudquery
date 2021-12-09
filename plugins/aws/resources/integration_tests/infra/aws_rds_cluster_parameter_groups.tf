resource "aws_rds_cluster_parameter_group" "rds_cluster_param_group" {
  name        = "rds-cluster-pg-${var.test_prefix}-${var.test_suffix}"
  family      = "aurora-mysql8.0"
  description = "Test RDS cluster parameter group"

  parameter {
    name  = "character_set_server"
    value = "utf8"
  }

  parameter {
    name  = "character_set_client"
    value = "utf8"
  }
}
