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

resource "aws_subnet" "rds_subnet_pub" {
  vpc_id     = aws_vpc.rds_vpc.id
  cidr_block = "10.1.3.0/24"


  tags = {
    Name = "Public RDS Subnet"
  }
}

resource "aws_internet_gateway" "igw" {
  vpc_id = aws_vpc.rds_vpc.id
}

resource "aws_eip" "nat_eip" {
  vpc      = true
}

resource "aws_nat_gateway" "nat_gw" {
  allocation_id = aws_eip.nat_eip.id
  subnet_id     = aws_subnet.rds_subnet_pub.id
}

resource "aws_security_group" "allow_postgresql" {
  name        = "allow_postgresql"
  description = "Allow inbound postresql connections"
  vpc_id      = aws_vpc.rds_vpc.id

  ingress {
    description = "Postgresql from security group"
    from_port   = 5432
    to_port     = 5432
    protocol    = "tcp"
    self        = true
  }

  egress {
    from_port   = 5432
    to_port     = 5432
    protocol    = "tcp"
    self        = true
  }

  tags = {
    Name = "Allow Postgresql"
  }
}

resource "aws_security_group" "allow_egress" {
  name        = "allow_egress"
  description = "Allow outbound connections"
  vpc_id      = aws_vpc.rds_vpc.id

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }

  tags = {
    Name = "allow_outbound"
  }
}

resource "aws_route_table" "public_egress" {
  vpc_id = aws_vpc.rds_vpc.id

  route {
    cidr_block = "0.0.0.0/0"
    gateway_id = aws_internet_gateway.igw.id
  }

  tags = {
    Name = "Public egress"
  }
}

resource "aws_route_table_association" "public" {
  subnet_id      = aws_subnet.rds_subnet_pub.id
  route_table_id = aws_route_table.public_egress.id
}

resource "aws_route_table" "private_egress" {
  vpc_id = aws_vpc.rds_vpc.id

  route {
    cidr_block = "0.0.0.0/0"
    nat_gateway_id = aws_nat_gateway.nat_gw.id
  }

  tags = {
    Name = "Private egress"
  }
}

resource "aws_route_table_association" "private_a" {
  subnet_id      = aws_subnet.rds_subnet_a.id
  route_table_id = aws_route_table.private_egress.id
}

resource "aws_route_table_association" "private_b" {
  subnet_id      = aws_subnet.rds_subnet_b.id
  route_table_id = aws_route_table.private_egress.id
}