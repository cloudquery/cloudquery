resource "aws_kms_key" "aws_kms_keys_key" {
  description             = "kms-key-${var.test_prefix}${var.test_suffix}"
  deletion_window_in_days = 10
}