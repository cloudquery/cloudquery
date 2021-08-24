resource "aws_vpc" "aws_vpc" {
  cidr_block = "10.0.0.0/16"
  tags ={
    Name = "vpc${var.test_prefix}-${var.test_suffix}"
  }
}

resource "aws_subnet" "aws_vpc_subnet" {
  vpc_id = aws_vpc.aws_vpc.id
  cidr_block = "10.0.1.0/24"
  availability_zone = "us-east-1e"
  tags = {
    Name = "vpc-subnet${var.test_prefix}-${var.test_suffix}"
  }
}

resource "aws_subnet" "aws_vpc_subnet2" {
  vpc_id = aws_vpc.aws_vpc.id
  cidr_block = "10.0.2.0/24"
  availability_zone = "us-east-1f"
  tags = {
    Name = "vpc-subnet2${var.test_prefix}-${var.test_suffix}"
  }
}

resource "aws_subnet" "aws_vpc_subnet3" {
  vpc_id = aws_vpc.aws_vpc.id
  cidr_block = "10.0.3.0/24"
  availability_zone = "us-east-1a"
  tags = {
    Name = "vpc-subnet3${var.test_prefix}-${var.test_suffix}"
  }
}

resource "aws_route_table" "aws_vpc_rt" {
  vpc_id = aws_vpc.aws_vpc.id

  tags = {
    Name = "vpc-routetable${var.test_prefix}-${var.test_suffix}"
  }
}

resource "aws_route" "aws_vpc_public_internet_gateway" {
  route_table_id         = aws_route_table.aws_vpc_rt.id
  destination_cidr_block = "0.0.0.0/0"
  gateway_id             = aws_internet_gateway.aws_vpc_igw.id

  timeouts {
    create = "5m"
  }
}

resource "aws_route_table_association" "aws_vpc_rt_assoc_s2" {
  subnet_id = aws_subnet.aws_vpc_subnet2.id
  route_table_id = aws_route_table.aws_vpc_rt.id
}

resource "aws_route_table_association" "aws_vpc_rt_assoc_s3" {
  subnet_id = aws_subnet.aws_vpc_subnet3.id
  route_table_id = aws_route_table.aws_vpc_rt.id
}

resource "aws_internet_gateway" "aws_vpc_igw" {
  vpc_id = aws_vpc.aws_vpc.id
}

resource "aws_network_interface" "aws_vpc_eni_s2" {
  subnet_id = aws_subnet.aws_vpc_subnet2.id
  private_ips = ["10.0.2.100"]

  tags = {
    Name = "primary_network_interface"
  }
}

resource "aws_network_interface" "aws_vpc_eni_s3" {
  subnet_id = aws_subnet.aws_vpc_subnet3.id
  private_ips = ["10.0.3.100"]

  tags = {
    Name = "primary_network_interface"
  }
}