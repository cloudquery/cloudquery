resource "aws_flow_log" "aws_ec2_flow_logs_fl" {
  log_destination      = aws_s3_bucket.aws_ec2_flow_logs_s3.arn
  log_destination_type = "s3"
  traffic_type         = "ALL"
  vpc_id               = aws_vpc.aws_vpc.id
}

resource "aws_s3_bucket" "aws_ec2_flow_logs_s3" {
  bucket        = "ec2-fl-buck${var.test_prefix}${var.test_suffix}"
  force_destroy = true
}