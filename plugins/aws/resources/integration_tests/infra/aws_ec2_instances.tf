data "aws_ami" "aws_ec2_instances_ami_ubuntu" {
  most_recent = true

  filter {
    name = "name"
    values = [
    "ubuntu/images/hvm-ssd/ubuntu-focal-20.04-amd64-server-*"]
  }

  filter {
    name = "virtualization-type"
    values = [
    "hvm"]
  }

  owners = [
  "099720109477"]
}

resource "aws_instance" "aws_ec2_instances_ec2_instance" {
  ami                  = data.aws_ami.aws_ec2_instances_ami_ubuntu.id
  subnet_id            = aws_subnet.aws_vpc_subnet.id
  instance_type        = "t2.nano"
  iam_instance_profile = aws_iam_instance_profile.aws_ec2_instances_ec2-instance-profile.name
  vpc_security_group_ids = [
  aws_security_group.aws_ec2_instances_security_group.id]

  ebs_optimized     = "false"
  source_dest_check = "false"
  user_data         = data.template_file.aws_ec2_instances_user_data.rendered
  root_block_device {
    volume_type           = "gp2"
    volume_size           = "30"
    delete_on_termination = "true"
  }

  lifecycle {
    ignore_changes = [
      ami,
      user_data,
      subnet_id,
      key_name,
      ebs_optimized,
    private_ip]
  }

  tags = {
    "Name" = "ec2_instance${var.test_suffix}"
  }
}

resource "aws_iam_role" "aws_ec2_instances_ec2_iam_role" {
  name = "ec2_iam_role_${var.test_prefix}${var.test_suffix}"

  # Terraform's "jsonencode" function converts a
  # Terraform expression result to valid JSON syntax.
  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Sid    = ""
        Principal = {
          Service = "ec2.amazonaws.com"
        }
      },
    ]
  })
}

resource "aws_iam_instance_profile" "aws_ec2_instances_ec2-instance-profile" {
  name = "ec2_instance_profile_${var.test_prefix}${var.test_suffix}"
  path = "/"
  role = aws_iam_role.aws_ec2_instances_ec2_iam_role.id
}

data "template_file" "aws_ec2_instances_user_data" {
  template = <<EOF
#!/bin/bash

# Update all packages

sudo yum update -y
sudo yum install -y ecs-init
sudo service docker start
sudo start ecs

#Adding cluster name in ecs config
echo ECS_CLUSTER=openapi-devl-cluster >> /etc/ecs/ecs.config
cat /etc/ecs/ecs.config | grep "ECS_CLUSTER"
EOF
}

resource "aws_security_group" "aws_ec2_instances_security_group" {
  name = "aws_ec2_instances_sg_${var.test_prefix}${var.test_suffix}"

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

