resource "awslightsail_bucket" "awslightsail_bucket" {
  name      = "${var.prefix}-lightsail-bucket"
  bundle_id = "small_1_0"
}