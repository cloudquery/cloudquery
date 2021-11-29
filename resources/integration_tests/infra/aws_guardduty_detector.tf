resource "aws_guardduty_detector" "MyDetector" {
  enable = true

  datasources {
    s3_logs {
      enable = true
    }
  }

  tags = {
    Name = "fguardduty-detector-${var.test_prefix}${var.test_suffix}"
  }
}