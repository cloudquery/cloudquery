resource "aws_iot_thing_type" "integration_test_thing_types" {
  name = "integration-test-thing-type"
}

resource "aws_iot_thing" "integration_test_thing" {
  name = "integration_test_thing"

  attributes = {
    First = "examplevalue"
  }
  thing_type_name = aws_iot_thing_type.integration_test_thing_types.name
}

resource "aws_iot_thing_group" "iot_thing_group_parent" {
  name = "iot_group_parent"
}

resource "aws_iot_thing_group" "iot_thing_group_group" {
  name = "iot_group_child"

  parent_group_name = aws_iot_thing_group.iot_thing_group_parent.name

  properties {
    attribute_payload {
      attributes = {
        One = "11111"
        Two = "TwoTwo"
      }
    }
    description = "This is my thing group"
  }

  tags = {
    terraform = "true"
  }
}

resource "aws_iot_thing_group_membership" "example" {
  thing_name       = aws_iot_thing.integration_test_thing.name
  thing_group_name = aws_iot_thing_group.iot_thing_group_group.name

  override_dynamic_group = true
}

resource "aws_iot_policy" "iot_policy" {
  name = "integration_test_policy"

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

resource "aws_iot_certificate" "iot_certificate" {
  csr    = file("iot_csr.csr")
  active = true
}


resource "aws_iot_policy_attachment" "integration_test_iot_policy_attachment" {
  policy = aws_iot_policy.iot_policy.name
  target = aws_iot_certificate.iot_certificate.arn
}

resource "aws_iot_topic_rule" "integration_test_iot_topic_rule" {
  name        = "integration_test_iot_topic_rule"
  description = "Example rule"
  enabled     = true
  sql         = "SELECT * FROM 'topic/test'"
  sql_version = "2016-03-23"

  sns {
    message_format = "RAW"
    role_arn       = aws_iam_role.integration_test_iot_role.arn
    target_arn     = aws_sns_topic.iot_test_topic.arn
  }

  error_action {
    sns {
      message_format = "RAW"
      role_arn       = aws_iam_role.integration_test_iot_role.arn
      target_arn     = aws_sns_topic.iot_test_topic.arn
    }
  }


}

resource "aws_sns_topic" "iot_test_topic" {
  name = "iot_test_topic"
}


resource "aws_iam_role" "integration_test_iot_role" {
  name = "integration_test_iot_role"

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
  name = "integration_test_tot_poicy"
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
        "Resource": "${aws_sns_topic.iot_test_topic.arn}"
    }
  ]
}
EOF
}