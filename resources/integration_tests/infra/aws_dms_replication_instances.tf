resource "aws_dms_replication_instance" "example" {
  replication_instance_class = "dms.t2.micro"
  replication_instance_id    = "dms-replication-instance-${var.test_prefix}-${var.test_suffix}"
}
