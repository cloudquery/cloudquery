resource "aws_redshift_subnet_group" "redshift_subnet_group" {
  name        = "redshift-sg-${var.test_prefix}${var.test_suffix}"
  subnet_ids  = [aws_subnet.aws_vpc_subnet.id, aws_subnet.aws_vpc_subnet2.id]
  description = "my test description"
}
