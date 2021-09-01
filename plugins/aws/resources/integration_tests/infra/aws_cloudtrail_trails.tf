resource "aws_s3_bucket" "aws_cloudtrail_trails_bucket" {
  bucket        = "cloudtrail-buc-${var.test_prefix}-${var.test_suffix}"
  force_destroy = true
}

resource "aws_s3_bucket" "aws_cloudtrail_trails_target_bucket" {
  bucket        = "cloudtrail-target-buc-${var.test_prefix}-${var.test_suffix}"
  force_destroy = true

  versioning {
    enabled = true
  }

  policy = <<POLICY
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Sid": "AWSCloudTrailAclCheck",
            "Effect": "Allow",
            "Principal": {
              "Service": "cloudtrail.amazonaws.com"
            },
            "Action": "s3:GetBucketAcl",
            "Resource": "arn:aws:s3:::cloudtrail-target-buc-${var.test_prefix}-${var.test_suffix}"
        },
        {
            "Sid": "AWSCloudTrailWrite",
            "Effect": "Allow",
            "Principal": {
              "Service": "cloudtrail.amazonaws.com"
            },
            "Action": "s3:PutObject",
            "Resource": "arn:aws:s3:::cloudtrail-target-buc-${var.test_prefix}-${var.test_suffix}/*",
            "Condition": {
                "StringEquals": {
                    "s3:x-amz-acl": "bucket-owner-full-control"
                }
            }
        },
        {
            "Sid": "AWSGlueReadLogs",
            "Effect": "Allow",
            "Principal": {
              "Service": "glue.amazonaws.com"
            },
            "Action": "s3:GetObject",
            "Resource": "arn:aws:s3:::cloudtrail-target-buc-${var.test_prefix}-${var.test_suffix}/*"
        }
    ]
}
POLICY

  tags = {
    "test" : "test"
  }
}

resource "aws_cloudtrail" "aws_cloudtrail_trails_trail" {
  name                          = "cloudtrail-${var.test_prefix}-${var.test_suffix}"
  s3_bucket_name                = aws_s3_bucket.aws_cloudtrail_trails_target_bucket.id
  s3_key_prefix                 = "cloudtrail"
  include_global_service_events = true
  enable_log_file_validation    = true
  enable_logging                = true
  is_multi_region_trail         = true

  event_selector {
    read_write_type           = "All"
    include_management_events = true

    data_resource {

      type = "AWS::S3::Object"

      values = [
      "${aws_s3_bucket.aws_cloudtrail_trails_bucket.arn}/"]
    }
  }
}