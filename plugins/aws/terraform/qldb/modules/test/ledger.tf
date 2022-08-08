
resource "aws_qldb_ledger" "cq-ledger" {
  name             = "${var.prefix}-ledger"
  permissions_mode = "STANDARD"
  deletion_protection = false
  tags = var.tags
}