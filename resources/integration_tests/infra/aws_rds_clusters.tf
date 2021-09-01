resource "aws_rds_cluster_instance" "cluster_instances" {
  count              = 1
  identifier         = "rdsclusterdb${var.test_suffix}"
  cluster_identifier = aws_rds_cluster.rds_cluster.id
  instance_class     = "db.t3.small"
  engine             = aws_rds_cluster.rds_cluster.engine
  engine_version     = aws_rds_cluster.rds_cluster.engine_version
}

resource "aws_rds_cluster" "rds_cluster" {
  cluster_identifier      = "rdscluster${var.test_suffix}"
  database_name           = "rdsclusterdb${var.test_suffix}"
  master_username         = "foo"
  master_password         = "bar123foo456"
  backup_retention_period = 5
  preferred_backup_window = "07:00-09:00"
}