resource "aws_kms_key" "aws_kms_keys_key" {
  description             = "kms-key-${var.test_prefix}${var.test_suffix}"
  deletion_window_in_days = 10
}


resource "time_sleep" "aws_kms_wait_for_key" {
  depends_on = [
  aws_kms_key.aws_kms_keys_key]

  create_duration = "10m"
}