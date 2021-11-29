data "aws_iam_policy_document" "dms_assume_role" {
  statement {
    actions = ["sts:AssumeRole"]

    principals {
      identifiers = ["dms.amazonaws.com"]
      type        = "Service"
    }
  }
}

resource "aws_iam_role" "dms_access_for_endpoint" {
  assume_role_policy = data.aws_iam_policy_document.dms_assume_role.json
  name               = "dms-access-for-endpoint-role"
}

resource "aws_iam_role_policy_attachment" "dms_access_for_endpoint_AmazonDMSRedshiftS3Role" {
  policy_arn = "arn:aws:iam::aws:policy/service-role/AmazonDMSRedshiftS3Role"
  role       = aws_iam_role.dms_access_for_endpoint.name
}

resource "aws_iam_role" "dms_cloudwatch_logs_role" {
  assume_role_policy = data.aws_iam_policy_document.dms_assume_role.json
  name               = "dms-cloudwatch-logs-role"
}

resource "aws_iam_role_policy_attachment" "dms_cloudwatch_logs_role_AmazonDMSCloudWatchLogsRole" {
  policy_arn = "arn:aws:iam::aws:policy/service-role/AmazonDMSCloudWatchLogsRole"
  role       = aws_iam_role.dms_cloudwatch_logs_role.name
}

resource "aws_iam_role" "dms_vpc_role" {
  assume_role_policy = data.aws_iam_policy_document.dms_assume_role.json
  name               = "dms-vpc-role"
}

resource "aws_iam_role_policy_attachment" "dms_vpc_role_AmazonDMSVPCManagementRole" {
  policy_arn = "arn:aws:iam::aws:policy/service-role/AmazonDMSVPCManagementRole"
  role       = aws_iam_role.dms_vpc_role.name
}

resource "aws_security_group" "dms_sg" {
  name        = "dms-sg-${var.test_prefix}_${var.test_suffix}"
  description = "Security group for DMS"
}

# Create a new replication instance
resource "aws_dms_replication_instance" "dms_replication_instance" {
  allocated_storage            = 20
  apply_immediately            = true
  auto_minor_version_upgrade   = true
  availability_zone            = "us-east-1a"
  multi_az                     = false
  preferred_maintenance_window = "sun:10:30-sun:14:30"
  publicly_accessible          = false
  replication_instance_class   = "dms.t2.micro"
  replication_instance_id      = "dms-replication-instance-${var.test_prefix}-${var.test_suffix}"

  vpc_security_group_ids = [
    aws_security_group.dms_sg.id
  ]

  tags = {
    Name = "dms-replication-instance-${var.test_prefix}-${var.test_suffix}"
  }

  depends_on = [
    aws_iam_role_policy_attachment.dms_access_for_endpoint_AmazonDMSRedshiftS3Role,
    aws_iam_role_policy_attachment.dms_cloudwatch_logs_role_AmazonDMSCloudWatchLogsRole,
    aws_iam_role_policy_attachment.dms_vpc_role_AmazonDMSVPCManagementRole
  ]
}