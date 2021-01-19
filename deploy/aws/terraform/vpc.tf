resource "aws_vpc" "rds_vpc" {
  cidr_block = "10.1.0.0/16"
}

resource "aws_subnet" "rds_subnet_a" {
  vpc_id     = aws_vpc.rds_vpc.id
  cidr_block = "10.1.1.0/24"
  availability_zone = data.aws_availability_zones.available.names[0]


  tags = {
    Name = "Private RDS Subnet A"
  }
}

data "aws_availability_zones" "available" {
  state = "available"
}

resource "aws_subnet" "rds_subnet_b" {
  vpc_id     = aws_vpc.rds_vpc.id
  cidr_block = "10.1.2.0/24"
  availability_zone = data.aws_availability_zones.available.names[1]

  tags = {
    Name = "Private RDS Subnet B"
  }
}

resource "aws_security_group" "allow_mysql" {
  name        = "allow_mysql"
  description = "Allow inbound mysql connections"
  vpc_id      = aws_vpc.rds_vpc.id

  ingress {
    description = "Mysql from security group"
    from_port   = 3306
    to_port     = 3306
    protocol    = "tcp"
    self        = true
  }

  egress {
    from_port   = 3306
    to_port     = 3306
    protocol    = "tcp"
    self        = true
  }

  tags = {
    Name = "allow_tls"
  }
}