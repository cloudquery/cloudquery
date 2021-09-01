resource "aws_vpc_peering_connection" "aws_ec2_vpc_peering_connections" {
  peer_vpc_id = aws_vpc.aws_vpc.id
  vpc_id      = aws_vpc.aws_ec2_vpc_peering_vpc.id

  accepter {
    allow_remote_vpc_dns_resolution = true
  }

  requester {
    allow_remote_vpc_dns_resolution = true
  }
  auto_accept = true

}

resource "aws_vpc" "aws_ec2_vpc_peering_vpc" {
  cidr_block = "10.1.0.0/16"
  tags = {
    Name = "vpc-peering-${var.test_prefix}-${var.test_suffix}"
  }
  enable_dns_hostnames = true
}

resource "aws_vpc_endpoint" "aws_ec2_vpc_endpoint" {
  vpc_id       = aws_vpc.aws_ec2_vpc_peering_vpc.id
  service_name = "com.amazonaws.${data.aws_region.current.name}.s3"

  tags = {
    Environment = "test"
  }
}

resource "aws_ec2_transit_gateway" "aws_ec2_vpc_endpoint_tg" {
  description = "description ${var.test_prefix}-${var.test_suffix}"
}


resource "aws_security_group" "aws_ec2_vpc_security_group" {
  name   = "ec2-sg-${var.test_prefix}${var.test_suffix}"
  vpc_id = aws_vpc.aws_vpc.id

  ingress {
    protocol  = "tcp"
    from_port = "80"
    to_port   = "80"
  }

  egress {
    from_port = 0
    to_port   = 0
    protocol  = "-1"
    cidr_blocks = [
    "0.0.0.0/0"]
    ipv6_cidr_blocks = [
    "::/0"]
  }
}

resource "aws_network_acl" "aws_ec2_vpc_network_acl" {
  vpc_id = aws_vpc.aws_vpc.id

  egress = [
    {
      protocol        = "tcp"
      rule_no         = 200
      action          = "allow"
      cidr_block      = "10.3.0.0/18"
      from_port       = 443
      to_port         = 443
      icmp_code       = 0
      icmp_type       = 0
      ipv6_cidr_block = ""
    }
  ]

  ingress = [
    {
      protocol        = "tcp"
      rule_no         = 100
      action          = "allow"
      cidr_block      = "10.3.0.0/18"
      from_port       = 80
      to_port         = 80
      icmp_code       = 0
      icmp_type       = 0
      ipv6_cidr_block = ""
    }
  ]

  tags = {
    Name = "ec2-acl-${var.test_prefix}-${var.test_suffix}"
  }
}


resource "aws_eip" "aws_ec2_vpc_eip" {
  vpc = true

  lifecycle {
    create_before_destroy = true
  }
}

resource "aws_nat_gateway" "aws_ec2_vpc_nat_gateway" {
  allocation_id = aws_eip.aws_ec2_vpc_eip.id
  subnet_id = aws_subnet.aws_vpc_subnet.id
  tags = {
    Name = "ec2-nat-${var.test_prefix}-${var.test_suffix}"
  }
}

