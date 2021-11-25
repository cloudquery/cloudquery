resource "aws_placement_group" "autoscaling_groups" {
  name     = "ag-${var.test_prefix}${var.test_suffix}"
  strategy = "cluster"
}

data "aws_ami" "autoscaling_groups_ami" {
  most_recent = true

  filter {
    name   = "name"
    values = [
      "ubuntu/images/hvm-ssd/ubuntu-focal-20.04-amd64-server-*"
    ]
  }

  filter {
    name   = "virtualization-type"
    values = [
      "hvm"
    ]
  }

  owners = [
    "099720109477"
  ]
}

resource "aws_launch_template" "autoscaling_groups_lt" {
  name_prefix   = "example"
  image_id      = data.aws_ami.autoscaling_groups_ami.id
  instance_type = "t2.nano"
}

resource "aws_autoscaling_group" "autoscaling_group" {
  name                  = "ag-${var.test_prefix}${var.test_suffix}"
  capacity_rebalance    = true
  desired_capacity      = 1
  max_size              = 1
  min_size              = 1
  wait_for_elb_capacity = 0
  health_check_type         = "ELB"
  health_check_grace_period = 300
  vpc_zone_identifier   = [aws_subnet.aws_vpc_subnet.id, aws_subnet.aws_vpc_subnet3.id]
  load_balancers        = [aws_elb.elbv1-loadbalancer.name]
  tag {
    key                 = "foo"
    value               = "bar"
    propagate_at_launch = true
  }

  timeouts {
    delete = "15m"
  }

  tag {
    key                 = "lorem"
    value               = "ipsum"
    propagate_at_launch = false
  }

  mixed_instances_policy {
    instances_distribution {
      on_demand_base_capacity                  = 1
      on_demand_percentage_above_base_capacity = 0
      #      spot_allocation_strategy                 = "capacity-optimized"
    }
    launch_template {
      launch_template_specification {
        launch_template_id = aws_launch_template.autoscaling_groups_lt.id
        version            = aws_launch_template.autoscaling_groups_lt.latest_version
      }

      override {
        instance_type = "t2.nano"
      }
      override {
        instance_type = "t2.small"
      }
    }
  }
}


resource "aws_autoscaling_policy" "bat" {
  name                   = "policy-${var.test_prefix}-${var.test_suffix}"
  scaling_adjustment     = 1
  adjustment_type        = "ChangeInCapacity"
  cooldown               = 300
  autoscaling_group_name = aws_autoscaling_group.autoscaling_group.name
}

resource "aws_autoscaling_notification" "example_notifications" {
  group_names = [
    aws_autoscaling_group.autoscaling_group.name,
    aws_autoscaling_group.autoscaling_group.name,
  ]


  notifications = [
    "autoscaling:EC2_INSTANCE_LAUNCH",
    "autoscaling:EC2_INSTANCE_TERMINATE",
    "autoscaling:EC2_INSTANCE_LAUNCH_ERROR",
    "autoscaling:EC2_INSTANCE_TERMINATE_ERROR",
  ]

  topic_arn = aws_sns_topic.autoscaling_group_hook_sns.arn
}

resource "aws_sns_topic" "autoscaling_group_hook_sns" {
  name         = "ag-topic-${var.test_suffix}"
  display_name = "${var.test_prefix}-${var.test_suffix}"

  delivery_policy = <<EOF
      {
        "http": {
          "defaultHealthyRetryPolicy": {
            "minDelayTarget": 20,
            "maxDelayTarget": 20,
            "numRetries": 3,
            "numMaxDelayRetries": 0,
            "numNoDelayRetries": 0,
            "numMinDelayRetries": 0,
            "backoffFunction": "linear"
          },
          "disableSubscriptionOverrides": false,
          "defaultThrottlePolicy": {
            "maxReceivesPerSecond": 1
          }
        }
      }
  EOF
}

resource "aws_autoscaling_lifecycle_hook" "autoscaling_group_hook" {
  name                   = "foobar${var.test_prefix}${var.test_suffix}"
  default_result         = "CONTINUE"
  heartbeat_timeout      = 2000
  autoscaling_group_name = aws_autoscaling_group.autoscaling_group.name

  lifecycle_transition    = "autoscaling:EC2_INSTANCE_LAUNCHING"
  notification_metadata   = jsonencode({ foo : "bar" })
  notification_target_arn = aws_sqs_queue.autoscaling_group_sqs.arn
  role_arn                = aws_iam_role.autoscaling_group_lifecycle_role.arn
}

resource "aws_sqs_queue" "autoscaling_group_sqs" {
  name = "ag-queue-${var.test_suffix}"
  #  fifo_queue                  = true
  #  content_based_deduplication = true
}

resource "aws_iam_role" "autoscaling_group_lifecycle_role" {
  name = "ag_role_${var.test_prefix}${var.test_suffix}"

  assume_role_policy = jsonencode({
    Version   = "2008-10-17"
    Statement = [
      {
        Action    = "sts:AssumeRole"
        Effect    = "Allow"
        Sid       = ""
        Principal = {
          Service = [
            "autoscaling.amazonaws.com"
          ]
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
            "s3:*",
            "sqs:SendMessage",
            "sqs:GetQueueUrl",
            "sns:Publish",
          ],
          Effect : "Allow",
          Resource : "*"
        }
      ]
    })
  }

}