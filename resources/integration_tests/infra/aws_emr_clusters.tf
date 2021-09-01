resource "aws_emr_cluster" "aws_emr_clusters_cluster" {
  name          = "emr-cluster-${var.test_prefix}${var.test_suffix}"
  release_label = "emr-5.12.0"
  applications = [
  "Spark"]

  log_uri = "s3://${aws_s3_bucket.aws_emr_cluster_logs.id}/"


  additional_info = <<EOF
{
  "instanceAwsClientConfiguration": {
    "proxyPort": 8099,
    "proxyHost": "myproxy.example.com"
  }
}
EOF

  termination_protection            = false
  keep_job_flow_alive_when_no_steps = true

  ec2_attributes {
    subnet_id                         = aws_subnet.aws_vpc_subnet3.id
    emr_managed_master_security_group = aws_security_group.aws_emr_clusters_security_group.id
    emr_managed_slave_security_group  = aws_security_group.aws_emr_clusters_security_group.id
    instance_profile                  = aws_iam_instance_profile.aws_emr_clusters_instance_profile.arn
  }

  master_instance_group {
    instance_type = "m1.small"
  }

  core_instance_group {
    instance_type  = "m1.small"
    instance_count = 1

    ebs_config {
      size                 = "40"
      type                 = "gp2"
      volumes_per_instance = 1
    }
    bid_price          = "0.60"
    autoscaling_policy = <<EOF
{
"Constraints": {
  "MinCapacity": 1,
  "MaxCapacity": 2
},
"Rules": [
  {
    "Name": "ScaleOutMemoryPercentage",
    "Description": "Scale out if YARNMemoryAvailablePercentage is less than 15",
    "Action": {
      "SimpleScalingPolicyConfiguration": {
        "AdjustmentType": "CHANGE_IN_CAPACITY",
        "ScalingAdjustment": 1,
        "CoolDown": 300
      }
    },
    "Trigger": {
      "CloudWatchAlarmDefinition": {
        "ComparisonOperator": "LESS_THAN",
        "EvaluationPeriods": 1,
        "MetricName": "YARNMemoryAvailablePercentage",
        "Namespace": "AWS/ElasticMapReduce",
        "Period": 300,
        "Statistic": "AVERAGE",
        "Threshold": 15.0,
        "Unit": "PERCENT"
      }
    }
  }
]
}
EOF
  }

  ebs_root_volume_size = 100

  tags = {
    role = "rolename"
    env  = "env"
  }

  bootstrap_action {
    path = "s3://elasticmapreduce/bootstrap-actions/run-if"
    name = "runif"
    args = [
      "instance.isMaster=true",
    "echo running on master node"]
  }

  configurations_json = <<EOF
  [
    {
      "Classification": "hadoop-env",
      "Configurations": [
        {
          "Classification": "export",
          "Properties": {
            "JAVA_HOME": "/usr/lib/jvm/java-1.8.0"
          }
        }
      ],
      "Properties": {}
    },
    {
      "Classification": "spark-env",
      "Configurations": [
        {
          "Classification": "export",
          "Properties": {
            "JAVA_HOME": "/usr/lib/jvm/java-1.8.0"
          }
        }
      ],
      "Properties": {}
    }
  ]
EOF
  autoscaling_role    = aws_iam_role.emr_clusters_autoscaling_role.arn
  service_role = aws_iam_role.emr_clusters_service_role.arn

}

resource "aws_iam_instance_profile" "aws_emr_clusters_instance_profile" {
  name = "emr-cluster-instance_profile_${var.test_prefix}${var.test_suffix}"
  path = "/"
  role = aws_iam_role.emr_clusters_instance_profile_role.id
}

resource "aws_iam_role" "emr_clusters_instance_profile_role" {
  name = "emr-cluster-instance-profile-role_${var.test_prefix}"

  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Sid    = ""
        Principal = {
          Service = [
          "ec2.amazonaws.com"]
        }
      }
    ]
  })

  inline_policy {
    name = "inline-${var.test_prefix}${var.test_suffix}"

    policy = jsonencode({
      Version = "2012-10-17"
      Statement : [
        {
          Action : [
            "cloudwatch:*",
            "dynamodb:*",
            "ec2:Describe*",
            "elasticmapreduce:Describe*",
            "elasticmapreduce:ListBootstrapActions",
            "elasticmapreduce:ListClusters",
            "elasticmapreduce:ListInstanceGroups",
            "elasticmapreduce:ListInstances",
            "elasticmapreduce:ListSteps",
            "kinesis:CreateStream",
            "kinesis:DeleteStream",
            "kinesis:DescribeStream",
            "kinesis:GetRecords",
            "kinesis:GetShardIterator",
            "kinesis:MergeShards",
            "kinesis:PutRecord",
            "kinesis:SplitShard",
            "rds:Describe*",
            "s3:*",
            "sdb:*",
            "sns:*",
            "sqs:*",
            "iam:CreateServiceLinkedRole",
            "glue:CreateDatabase",
            "glue:UpdateDatabase",
            "glue:DeleteDatabase",
            "glue:GetDatabase",
            "glue:GetDatabases",
            "glue:CreateTable",
            "glue:UpdateTable",
            "glue:DeleteTable",
            "glue:GetTable",
            "glue:GetTables",
            "glue:GetTableVersions",
            "glue:CreatePartition",
            "glue:BatchCreatePartition",
            "glue:UpdatePartition",
            "glue:DeletePartition",
            "glue:BatchDeletePartition",
            "glue:GetPartition",
            "glue:GetPartitions",
            "glue:BatchGetPartition",
            "glue:CreateUserDefinedFunction",
            "glue:UpdateUserDefinedFunction",
            "glue:DeleteUserDefinedFunction",
            "glue:GetUserDefinedFunction",
            "glue:GetUserDefinedFunctions"
          ],
          Effect : "Allow",
          Resource : "*"
        }
      ]
    })
  }

}

resource "aws_iam_role" "emr_clusters_autoscaling_role" {
  name = "emr-cluster-autoscaling-role_${var.test_prefix}"

  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Sid    = ""
        Principal = {
          Service = [
            "ec2.amazonaws.com",
          "elasticmapreduce.amazonaws.com"]
        }
      }
    ]
  })

  inline_policy {
    name = "inline-${var.test_prefix}${var.test_suffix}"

    policy = jsonencode({
      Version = "2012-10-17"
      Statement : [
        {
          Action : [
            "cloudwatch:DescribeAlarms",
            "elasticmapreduce:ListInstanceGroups",
            "elasticmapreduce:ModifyInstanceGroups",
            "iam:CreateServiceLinkedRole",
          ],
          Effect : "Allow",
          Resource : "*"
        }
      ]
    })
  }
}

resource "aws_iam_role" "emr_clusters_service_role" {
  name = "emr_clusters_service_role_${var.test_prefix}${var.test_suffix}"

  assume_role_policy = jsonencode({
    Version = "2008-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Sid    = ""
        Principal = {
          Service = [
          "elasticmapreduce.amazonaws.com"]
        }
      }
    ]
  })

  inline_policy {
    name = "inline-${var.test_prefix}${var.test_suffix}"

    policy = jsonencode({
      Version = "2012-10-17"
      Statement : [
        {
          Action : [
            "ec2:AuthorizeSecurityGroupEgress",
            "ec2:AuthorizeSecurityGroupIngress",
            "ec2:CancelSpotInstanceRequests",
            "ec2:CreateFleet",
            "ec2:CreateLaunchTemplate",
            "ec2:CreateNetworkInterface",
            "ec2:CreateSecurityGroup",
            "ec2:CreateTags",
            "ec2:DeleteLaunchTemplate",
            "ec2:DeleteNetworkInterface",
            "ec2:DeleteSecurityGroup",
            "ec2:DeleteTags",
            "ec2:DescribeAvailabilityZones",
            "ec2:DescribeAccountAttributes",
            "ec2:DescribeDhcpOptions",
            "ec2:DescribeImages",
            "ec2:DescribeInstanceStatus",
            "ec2:DescribeInstances",
            "ec2:DescribeKeyPairs",
            "ec2:DescribeLaunchTemplates",
            "ec2:DescribeNetworkAcls",
            "ec2:DescribeNetworkInterfaces",
            "ec2:DescribePrefixLists",
            "ec2:DescribeRouteTables",
            "ec2:DescribeSecurityGroups",
            "ec2:DescribeSpotInstanceRequests",
            "ec2:DescribeSpotPriceHistory",
            "ec2:DescribeSubnets",
            "ec2:DescribeTags",
            "ec2:DescribeVpcAttribute",
            "ec2:DescribeVpcEndpoints",
            "ec2:DescribeVpcEndpointServices",
            "ec2:DescribeVpcs",
            "ec2:DetachNetworkInterface",
            "ec2:ModifyImageAttribute",
            "ec2:ModifyInstanceAttribute",
            "ec2:RequestSpotInstances",
            "ec2:RevokeSecurityGroupEgress",
            "ec2:RunInstances",
            "ec2:TerminateInstances",
            "ec2:DeleteVolume",
            "ec2:DescribeVolumeStatus",
            "ec2:DescribeInstanceAttribute",
            "ec2:DescribeVolumes",
            "ec2:DetachVolume",
            "iam:GetRole",
            "iam:CreateServiceLinkedRole",
            "iam:GetRolePolicy",
            "iam:ListInstanceProfiles",
            "iam:ListRolePolicies",
            "iam:PassRole",
            "s3:CreateBucket",
            "s3:Get*",
            "s3:List*",
            "sdb:BatchPutAttributes",
            "sdb:Select",
            "sqs:CreateQueue",
            "sqs:Delete*",
            "sqs:GetQueue*",
            "sqs:PurgeQueue",
            "sqs:ReceiveMessage",
            "cloudwatch:PutMetricAlarm",
            "cloudwatch:DescribeAlarms",
            "cloudwatch:DeleteAlarms",
            "application-autoscaling:RegisterScalableTarget",
            "application-autoscaling:DeregisterScalableTarget",
            "application-autoscaling:PutScalingPolicy",
            "application-autoscaling:DeleteScalingPolicy",
            "application-autoscaling:Describe*"
          ],
          Effect : "Allow",
          Resource : "*"
        }
      ]
    })
  }

}

resource "aws_s3_bucket" "aws_emr_cluster_logs" {
  bucket        = "emr-cluster-logs${var.test_prefix}${var.test_suffix}"
  acl           = "private"
  force_destroy = true

  tags = {
    Name        = "${var.test_prefix}${var.test_suffix}"
    Environment = "Dev"
  }
}

resource "aws_security_group_rule" "aws_emr_cluster_allow_tcp_from_master_to_service" {
  type                     = "ingress"
  from_port                = 9443
  to_port                  = 9443
  protocol                 = "tcp"
  security_group_id        = aws_security_group.aws_emr_clusters_security_group.id
  source_security_group_id = aws_security_group.aws_emr_clusters_security_group.id
}

resource "aws_security_group" "aws_emr_clusters_security_group" {
  name = "aws_emr_clusters_security_group${var.test_prefix}${var.test_suffix}"

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
