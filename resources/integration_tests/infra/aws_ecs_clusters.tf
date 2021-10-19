resource "aws_iam_role" "aws_ecs_clusters_ec2_iam_role" {
  name = "ec2_ec2_iam_role_${var.test_prefix}${var.test_suffix}"

  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Sid = ""
        Principal = {
          Service = "ec2.amazonaws.com"
        }
      },
    ]
  })
}

resource "aws_iam_instance_profile" "aws_ecs_clusters_ec2-instance-profile" {
  name = "ecs_ec2_instance_profile_${var.test_prefix}${var.test_suffix}"
  path = "/"
  role = aws_iam_role.aws_ecs_clusters_ec2_iam_role.id
}

resource "aws_security_group" "aws_ecs_clusters_security_group" {
  name = "ecs_clusters_sg${var.test_prefix}${var.test_suffix}"

  vpc_id = aws_vpc.aws_vpc.id

  ingress {
    protocol = "tcp"
    from_port = "80"
    to_port = "80"
  }

  egress {
    from_port = 0
    to_port = 0
    protocol = "-1"
    cidr_blocks = [
      "0.0.0.0/0"]
    ipv6_cidr_blocks = [
      "::/0"]
  }
}

resource "aws_instance" "aws_ecs_clusters_ec2_instance" {
  ami = data.aws_ami.aws_ecs_clusters_ami_ubuntu.id
  subnet_id = aws_subnet.aws_vpc_subnet.id
  instance_type = "t2.nano"
  iam_instance_profile = aws_iam_instance_profile.aws_ecs_clusters_ec2-instance-profile.name
  vpc_security_group_ids = [
    aws_security_group.aws_ecs_clusters_security_group.id]
  ebs_optimized = "false"
  source_dest_check = "false"
  user_data = data.template_file.aws_ecs_clusters_user_data.rendered
  root_block_device {
    volume_type = "gp2"
    volume_size = "30"
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
    "Name" = "ecs_ec2_instance${var.test_suffix}"
  }
}

data "aws_ami" "aws_ecs_clusters_ami_ubuntu" {
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

data "aws_iam_policy_document" "aws_ecs_clusters_ecs-instance-policy" {
  statement {
    actions = [
      "sts:AssumeRole"]
    principals {
      type = "Service"
      identifiers = [
        "ec2.amazonaws.com"]
    }
  }
}

resource "aws_iam_role" "aws_ecs_clusters_ecs-instance-role" {
  name = "ecs-instance-role_${var.test_prefix}${var.test_suffix}"
  path = "/"
  assume_role_policy = data.aws_iam_policy_document.aws_ecs_clusters_ecs-instance-policy.json
}

resource "aws_iam_role_policy_attachment" "aws_ecs_clusters_ecs-instance-role-attachment" {
  role = aws_iam_role.aws_ecs_clusters_ecs-instance-role.name
  policy_arn = "arn:aws:iam::aws:policy/service-role/AmazonEC2ContainerServiceforEC2Role"
}

resource "aws_iam_instance_profile" "aws_ecs_clusters_instance-profile" {
  name = "ecs-instance-profile_${var.test_prefix}${var.test_suffix}"
  path = "/"
  role = aws_iam_role.aws_ecs_clusters_ecs-instance-role.id
}

data "aws_iam_policy_document" "aws_ecs_clusters_service-policy" {
  statement {
    actions = [
      "sts:AssumeRole"]
    principals {
      type = "Service"
      identifiers = [
        "ecs.amazonaws.com"]
    }
  }
}

resource "aws_iam_role" "aws_ecs_clusters_service-role" {
  name = "ecs_service_role_${var.test_prefix}${var.test_suffix}"
  path = "/"
  assume_role_policy = data.aws_iam_policy_document.aws_ecs_clusters_service-policy.json
}

resource "aws_iam_role_policy_attachment" "aws_ecs_clusters_service-role-attachment" {
  role = aws_iam_role.aws_ecs_clusters_service-role.name
  policy_arn = "arn:aws:iam::aws:policy/service-role/AmazonEC2ContainerServiceRole"
}

resource "aws_ecs_cluster" "aws_ecs_clusters_cluster" {
  name = "ecs_cluster_${var.test_prefix}${var.test_suffix}"
  setting {
    name = "containerInsights"
    value = "enabled"
  }
  tags = {
    name = "ecs_cluster_${var.test_prefix}${var.test_suffix}"
    stage = "test"
  }
}

data "template_file" "aws_ecs_clusters_user_data" {
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

resource "aws_ecs_task_definition" "aws_ecs_clusters_task_definition" {
  container_definitions = jsonencode([
    {
      "name" : "web-server",
      "image" : "nginx",
      "cpu" : 10,
      "memory" : 512,
      "logConfiguration" : {
        "logDriver" : "awslogs",
        "options" : {
          "awslogs-group" : "openapi-devl-cw",
          "awslogs-region" : "eu-west-1",
          "awslogs-stream-prefix" : "ecs"
        }
      },
      "links" : [],
      "portMappings" : [
        {
          "hostPort" : 8080,
          "containerPort" : 8080,
          "protocol" : "tcp"
        }
      ],
      "essential" : true,
      "entryPoint" : [],
      "command" : [],
      "environment" : [],
      "mountPoints" : [],
      "volumesFrom" : []
    }
  ])
  family = "openapi-task-defination"
  network_mode = "awsvpc"
  memory = "2048"
  cpu = "1024"
  requires_compatibilities = [
    "EC2"]
  # TASK running role
}

resource "aws_ecs_service" "aws_ecs_clusters_service" {
  cluster = aws_ecs_cluster.aws_ecs_clusters_cluster.id
  desired_count = 1
  launch_type = "EC2"
  name = "ecs_service_${var.test_prefix}${var.test_suffix}"
  task_definition = aws_ecs_task_definition.aws_ecs_clusters_task_definition.arn

  load_balancer {
    container_name = "web-server"
    container_port = "8080"
    target_group_arn = aws_lb_target_group.aws_elbv2_lb_target_group.arn
  }
  network_configuration {
    security_groups = [
      aws_security_group.aws_ecs_clusters_security_group.id]
    subnets = [
      aws_subnet.aws_vpc_subnet.id,
      aws_subnet.aws_vpc_subnet2.id]
    assign_public_ip = "false"
  }
  depends_on = [
    aws_lb_target_group.aws_elbv2_lb_target_group]
}

// TODO - move to separate file
resource "aws_cloudwatch_log_group" "aws_ecs_clusters_log_group" {
  name = "ecs_clusters_log_group_${var.test_prefix}${var.test_suffix}"
  tags = {
    Environment = "production"
  }
}
