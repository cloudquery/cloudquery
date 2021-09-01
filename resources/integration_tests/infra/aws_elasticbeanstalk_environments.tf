resource "aws_elastic_beanstalk_application" "aws_elasticbeanstalk_environments_app" {
  name        = "beanstalk-ea-${var.test_suffix}"
  description = "tf-test-desc"
}

resource "aws_elastic_beanstalk_environment" "aws_elasticbeanstalk_environments_env" {
  name                = substr("beanstalk-ee-${var.test_suffix}", 0, 40)
  application         = aws_elastic_beanstalk_application.aws_elasticbeanstalk_environments_app.name
  solution_stack_name = "64bit Amazon Linux 2 v3.3.4 running Go 1"


  setting {
    namespace = "aws:ec2:vpc"
    name      = "VPCId"
    value     = aws_vpc.aws_vpc.id
  }

  setting {
    namespace = "aws:ec2:vpc"
    name      = "AssociatePublicIpAddress"
    value     = "True"
  }

  setting {
    namespace = "aws:ec2:vpc"
    name      = "Subnets"
    value = join(",", [
      aws_subnet.aws_vpc_subnet2.id,
    aws_subnet.aws_vpc_subnet3.id])
  }

  setting {
    namespace = "aws:autoscaling:launchconfiguration"
    name      = "IamInstanceProfile"
    value     = aws_iam_instance_profile.aws_elasticbeanstalk_environments_ip.id
  }


  setting {
    namespace = "aws:elasticbeanstalk:environment:process:default"
    name      = "MatcherHTTPCode"
    value     = "200"
  }
  setting {
    namespace = "aws:elasticbeanstalk:environment"
    name      = "LoadBalancerType"
    value     = "application"
  }
  setting {
    namespace = "aws:autoscaling:launchconfiguration"
    name      = "InstanceType"
    value     = "t2.nano"
  }
  setting {
    namespace = "aws:ec2:vpc"
    name      = "ELBScheme"
    value     = "internet facing"
  }
  setting {
    namespace = "aws:autoscaling:asg"
    name      = "MinSize"
    value     = 1
  }
  setting {
    namespace = "aws:autoscaling:asg"
    name      = "MaxSize"
    value     = 2
  }
  setting {
    namespace = "aws:elasticbeanstalk:healthreporting:system"
    name      = "SystemType"
    value     = "enhanced"
  }

  tags = {
    name = substr("e-${var.test_suffix}", 0, 40)
  }

}

resource "aws_iam_role" "aws_elasticbeanstalk_environments_ir" {
  name = "beanstalk-role_${var.test_prefix}${var.test_suffix}"

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
    }]
  })

  inline_policy {
    name = "beanstalk-inline-${var.test_prefix}${var.test_suffix}"

    policy = jsonencode({
      Version = "2012-10-17"
      Statement : [
        {
          Sid : "BucketAccess",
          Action : [
            "s3:Get*",
            "s3:List*",
            "s3:PutObject"
          ],
          Effect : "Allow",
          Resource : [
            "arn:aws:s3:::elasticbeanstalk-*",
            "arn:aws:s3:::elasticbeanstalk-*/*"
          ]
        },
        {
          Sid : "XRayAccess",
          Action : [
            "xray:PutTraceSegments",
            "xray:PutTelemetryRecords",
            "xray:GetSamplingRules",
            "xray:GetSamplingTargets",
            "xray:GetSamplingStatisticSummaries"
          ],
          Effect : "Allow",
          Resource : "*"
        },
        {
          Sid : "CloudWatchLogsAccess",
          Action : [
            "logs:PutLogEvents",
            "logs:CreateLogStream",
            "logs:DescribeLogStreams",
            "logs:DescribeLogGroups"
          ],
          Effect : "Allow",
          Resource : [
            "arn:aws:logs:*:*:log-group:/aws/elasticbeanstalk*"
          ]
        },
        {
          Sid : "ElasticBeanstalkHealthAccess",
          Action : [
            "elasticbeanstalk:PutInstanceStatistics"
          ],
          Effect : "Allow",
          Resource : [
            "arn:aws:elasticbeanstalk:*:*:application/*",
            "arn:aws:elasticbeanstalk:*:*:environment/*"
          ]
        }
      ]
    })
  }
}

resource "aws_iam_instance_profile" "aws_elasticbeanstalk_environments_ip" {
  name = "beanstalk-ip_${var.test_prefix}${var.test_suffix}"
  path = "/"
  role = aws_iam_role.aws_elasticbeanstalk_environments_ir.id
}