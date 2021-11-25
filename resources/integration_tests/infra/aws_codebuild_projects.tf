resource "aws_s3_bucket" "codebuild_s3" {
  bucket = "codebuild${var.test_prefix}${var.test_suffix}"
  acl    = "private"
}

resource "aws_iam_role" "codebuild_role" {
  name = "codbuildrole-${var.test_prefix}${var.test_suffix}"

  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Principal": {
        "Service": "codebuild.amazonaws.com"
      },
      "Action": "sts:AssumeRole"
    }
  ]
}
EOF
}

resource "aws_iam_role_policy" "codebuild_role_policy" {
  role = aws_iam_role.codebuild_role.name

  policy = <<POLICY
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Resource": [
        "*"
      ],
      "Action": [
        "logs:CreateLogGroup",
        "logs:CreateLogStream",
        "logs:PutLogEvents"
      ]
    },
    {
      "Effect": "Allow",
      "Action": [
        "ec2:CreateNetworkInterface",
        "ec2:DescribeDhcpOptions",
        "ec2:DescribeNetworkInterfaces",
        "ec2:DeleteNetworkInterface",
        "ec2:DescribeSubnets",
        "ec2:DescribeSecurityGroups",
        "ec2:DescribeVpcs"
      ],
      "Resource": "*"
    },
    {
      "Effect": "Allow",
      "Action": [
        "ec2:CreateNetworkInterfacePermission"
      ],
      "Resource": [
        "arn:aws:ec2:us-east-1:123456789012:network-interface/*"
      ],
      "Condition": {
        "StringEquals": {
          "ec2:Subnet": [
            "${aws_subnet.aws_vpc_subnet.arn}",
            "${aws_subnet.aws_vpc_subnet3.arn}"
          ],
          "ec2:AuthorizedService": "codebuild.amazonaws.com"
        }
      }
    },
    {
      "Effect": "Allow",
      "Action": [
        "s3:*"
      ],
      "Resource": [
        "${aws_s3_bucket.codebuild_s3.arn}",
        "${aws_s3_bucket.codebuild_s3.arn}/*"
      ]
    }
  ]
}
POLICY
}

resource "aws_codebuild_project" "codebuild_project" {
  name          = "project-${var.test_prefix}${var.test_suffix}"
  description   = "test_codebuild_project"
  build_timeout = "5"
  service_role  = aws_iam_role.codebuild_role.arn

  artifacts {
    type = "NO_ARTIFACTS"
  }

  cache {
    type     = "S3"
    location = aws_s3_bucket.codebuild_s3.bucket
  }

  environment {
    compute_type                = "BUILD_GENERAL1_SMALL"
    image                       = "aws/codebuild/standard:1.0"
    type                        = "LINUX_CONTAINER"
    image_pull_credentials_type = "CODEBUILD"

    privileged_mode = true

    environment_variable {
      name  = "SOME_KEY1"
      value = "SOME_VALUE1"
    }

    environment_variable {
      name  = "SOME_KEY2"
      value = "SOME_VALUE2"
      type  = "PARAMETER_STORE"
    }
  }

  logs_config {
    cloudwatch_logs {
      group_name  = "log-group"
      stream_name = "log-stream"
    }

    s3_logs {
      status   = "ENABLED"
      location = "${aws_s3_bucket.codebuild_s3.id}/build-log"
    }
  }

  source {
    type            = "GITHUB"
    location        = "https://github.com/mitchellh/packer.git"
    git_clone_depth = 1

    git_submodules_config {
      fetch_submodules = true
    }


  }
  secondary_sources {
    type              = "S3"
    source_identifier = "package"
    location          = "${aws_s3_bucket.codebuild_s3.bucket}/package.zip"
  }

  secondary_artifacts {
    type                = "S3"
    artifact_identifier = "package"
    location            = aws_s3_bucket.codebuild_s3.bucket
  }

  file_system_locations {
    identifier = "CODEBUILD_MY_EFS"
    location   = "${aws_efs_file_system.codebuild_efs.dns_name}:/path"
    mount_point = "/path"
  }

  source_version = "master"

  vpc_config {
    vpc_id = aws_vpc.aws_vpc.id

    subnets = [
      aws_subnet.aws_vpc_subnet.id,
      aws_subnet.aws_vpc_subnet3.id,
    ]

    security_group_ids = [
      aws_security_group.codebuild_sg.id,
    ]
  }

  tags = {
    Environment = "Test"
  }
}

resource "aws_efs_file_system" "codebuild_efs" {
  creation_token = "efs${var.test_prefix}${var.test_suffix}"

  tags = {
    Name = "MyProduct"
  }
}


resource "aws_security_group" "codebuild_sg" {
  name = "ecs_clusters_sg${var.test_prefix}${var.test_suffix}"

  vpc_id = aws_vpc.aws_vpc.id

  ingress {
    protocol  = "tcp"
    from_port = "80"
    to_port   = "80"
  }

  egress {
    from_port        = 0
    to_port          = 0
    protocol         = "-1"
    cidr_blocks      = [
      "0.0.0.0/0"
    ]
    ipv6_cidr_blocks = [
      "::/0"
    ]
  }
}