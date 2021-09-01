resource "aws_redshift_parameter_group" "redshift_parameter_group" {
  name   = "redshift-pg-${var.test_prefix}${var.test_suffix}"
  family = "redshift-1.0"

  parameter {
    name  = "require_ssl"
    value = "true"
  }

  parameter {
    name  = "query_group"
    value = "example"
  }

  parameter {
    name  = "enable_user_activity_logging"
    value = "true"
  }
}

resource "aws_redshift_cluster" "redshift_cluster" {
  cluster_identifier           = "redshift-cluster${var.test_prefix}${var.test_suffix}"
  database_name                = "mydb"
  master_username              = "foo"
  master_password              = "Mustbe8characters"
  node_type                    = "dc2.large"
  cluster_type                 = "single-node"
  cluster_parameter_group_name = aws_redshift_parameter_group.redshift_parameter_group.name
  skip_final_snapshot          = true
}
