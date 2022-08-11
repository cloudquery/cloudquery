resource "aws_kms_key" "example" {
  description             = "${var.prefix}-xray-kms-key"
  deletion_window_in_days = 7

}

resource "aws_xray_encryption_config" "xray-enc-conf" {
  type   = "KMS"
  key_id = aws_kms_key.example.arn
}
