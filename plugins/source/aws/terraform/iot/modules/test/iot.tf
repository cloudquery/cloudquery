
resource "aws_iot_thing_type" "example" {
  name = "${var.prefix}-iot"
  tags = var.tags
}

resource "aws_iot_thing" "example" {
  name = "${var.prefix}-iot"

  attributes = {
    First = "examplevalue"
  }
  thing_type_name = aws_iot_thing_type.example.name
}

resource "aws_iot_thing_group" "parent" {
  name = "${var.prefix}-iot-parent"
  
}

resource "aws_iot_thing_group" "child" {
  name = "${var.prefix}-iot-child"

  parent_group_name = aws_iot_thing_group.parent.name

  properties {
    attribute_payload {
      attributes = {
        One = "11111"
        Two = "TwoTwo"
      }
    }
    description = "This is my thing group"
  }

  tags = var.tags
}

resource "aws_iot_thing_group_membership" "example" {
  thing_name       = aws_iot_thing.example.name
  thing_group_name = aws_iot_thing_group.parent.name

  override_dynamic_group = true
}

resource "aws_iot_policy" "example" {
  name = "${var.prefix}-iot"

  # Terraform's "jsonencode" function converts a
  # Terraform expression result to valid JSON syntax.
  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = [
          "iot:*",
        ]
        Effect   = "Allow"
        Resource = "*"
      },
    ]
  })
}

resource "aws_iot_certificate" "example" {
  active = true
}

resource "aws_iot_policy_attachment" "integration_test_iot_policy_attachment" {
  policy = aws_iot_policy.example.name
  target = aws_iot_certificate.example.arn
}

resource "aws_iot_topic_rule" "example" {
  name        = "${var.prefix}_iot"
  description = "Example rule"
  enabled     = true
  sql         = "SELECT * FROM 'topic/test'"
  sql_version = "2016-03-23"

  sns {
    message_format = "RAW"
    role_arn       = aws_iam_role.integration_test_iot_role.arn
    target_arn     = aws_sns_topic.example.arn
  }

  error_action {
    sns {
      message_format = "RAW"
      role_arn       = aws_iam_role.integration_test_iot_role.arn
      target_arn     = aws_sns_topic.example.arn
    }
  }


}

resource "aws_sns_topic" "example" {
  name = "${var.prefix}-iot"
}

resource "aws_iam_role" "integration_test_iot_role" {
  name = "${var.prefix}-iot"

  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Principal": {
        "Service": "iot.amazonaws.com"
      },
      "Action": "sts:AssumeRole"
    }
  ]
}
EOF
}

resource "aws_iam_role_policy" "integration_test_tot_poicy" {
  name = "${var.prefix}-iot"
  role = aws_iam_role.integration_test_iot_role.id

  policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
        "Effect": "Allow",
        "Action": [
            "sns:Publish"
        ],
        "Resource": "${aws_sns_topic.example.arn}"
    }
  ]
}
EOF
}
