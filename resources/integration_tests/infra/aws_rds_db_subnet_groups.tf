resource "aws_db_subnet_group" "rds_db_subnet" {
  name       = "rds_db_subnet${var.test_prefix}${var.test_suffix}"
  subnet_ids = [aws_subnet.aws_vpc_subnet.id, aws_subnet.aws_vpc_subnet2.id]

  tags = {
    Name = "rds_db_subnet${var.test_prefix}${var.test_suffix}"
  }
}