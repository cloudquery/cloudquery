resource "aws_kms_key" "backup_kms_key" {}

resource "aws_backup_vault" "backup_vault" {
  name        = "${var.prefix}_backup_vault"
  kms_key_arn = aws_kms_key.backup_kms_key.arn
  tags = {
    key = "backup_vault"
  }
}

resource "aws_backup_plan" "plan" {
  name = "${var.prefix}_backup_plan"

  rule {
    rule_name         = "${var.prefix}_backup_rule"
    target_vault_name = aws_backup_vault.backup_vault.name
    schedule          = "cron(0 12 * * ? *)"
  }
  
  tags = {
    key = "backup_plan"
  }
}

resource "aws_iam_role" "role" {
  name               = "${var.prefix}_backup_role"
  assume_role_policy = <<POLICY
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": ["sts:AssumeRole"],
      "Effect": "allow",
      "Principal": {
        "Service": ["backup.amazonaws.com"]
      }
    }
  ]
}
POLICY
  managed_policy_arns = ["arn:aws:iam::aws:policy/AWSBackupFullAccess"]
}

resource "aws_backup_selection" "selection" {
  iam_role_arn = aws_iam_role.role.arn
  name         = "${var.prefix}_backup_selection"
  plan_id      = aws_backup_plan.plan.id

  selection_tag {
    type  = "STRINGEQUALS"
    key   = "foo"
    value = "bar"
  }
}

resource "aws_sns_topic" "backup_sns_topic" {
  name = "${var.prefix}-backup-vault-events"
}

data "aws_iam_policy_document" "document" {
  policy_id = "${var.prefix}-backup-policy-document"

  statement {
    actions = [
      "SNS:Publish",
    ]

    effect = "Allow"

    principals {
      type        = "Service"
      identifiers = ["backup.amazonaws.com"]
    }

    resources = [
      aws_sns_topic.backup_sns_topic.arn,
    ]

    sid = "__default_statement_ID"
  }
}

resource "aws_sns_topic_policy" "test" {
  arn    = aws_sns_topic.backup_sns_topic.arn
  policy = data.aws_iam_policy_document.document.json
}

resource "aws_backup_vault_notifications" "test" {
  backup_vault_name   = aws_backup_vault.backup_vault.name
  sns_topic_arn       = aws_sns_topic.backup_sns_topic.arn
  backup_vault_events = ["BACKUP_JOB_STARTED", "RESTORE_JOB_COMPLETED"]
}
