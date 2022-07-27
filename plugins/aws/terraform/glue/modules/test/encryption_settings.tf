resource "aws_glue_data_catalog_encryption_settings" "example" {
  data_catalog_encryption_settings {
    connection_password_encryption {
      aws_kms_key_id                       = aws_kms_key.aws_kms_key.arn
      return_connection_password_encrypted = true
    }

    encryption_at_rest {
      catalog_encryption_mode = "SSE-KMS"
      sse_aws_kms_key_id      = aws_kms_key.aws_kms_key.arn
    }
  }
}

resource "aws_glue_security_configuration" "example" {
  name = "${var.prefix}-glue-security-config"

  encryption_configuration {
    cloudwatch_encryption {
      cloudwatch_encryption_mode = "DISABLED"
    }

    job_bookmarks_encryption {
      job_bookmarks_encryption_mode = "DISABLED"
    }

    s3_encryption {
      kms_key_arn        = aws_kms_key.aws_kms_key.arn
      s3_encryption_mode = "SSE-KMS"
    }
  }
}

resource "aws_kms_key" "aws_kms_key" {
  description             = "KMS key 1"
  deletion_window_in_days = 10
}