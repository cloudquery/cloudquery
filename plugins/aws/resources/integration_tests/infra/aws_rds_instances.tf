data "aws_rds_certificate" "rds_ca" {
  latest_valid_till = true
}

resource "aws_db_instance" "rds_db_instance" {
  allocated_storage    = 10
  engine               = "mysql"
  engine_version       = "5.7"
  instance_class       = "db.t3.micro"
  name                 = "rds${var.test_suffix}"
  username             = "foo"
  password             = "foobarbaz"
  parameter_group_name = "default.mysql5.7"
  skip_final_snapshot  = true
  ca_cert_identifier   = data.aws_rds_certificate.rds_ca.id

  tags = {
    Name = "rds-${var.test_prefix}${var.test_suffix}"
  }
}