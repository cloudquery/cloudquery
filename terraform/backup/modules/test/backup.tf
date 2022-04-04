resource "aws_kms_key" "backup_kms_key" {}

resource "aws_backup_vault" "backup_vault" {
  name        = "${var.prefix}_backup_vault"
  kms_key_arn = aws_kms_key.backup_kms_key.arn
}

resource "aws_backup_plan" "plan" {
  name = "${var.prefix}_backup_plan"

  rule {
    rule_name         = "${var.prefix}_backup_rule"
    target_vault_name = aws_backup_vault.backup_vault.name
    schedule          = "cron(0 12 * * ? *)"
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
