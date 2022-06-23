resource "awslightsail_disk" "awslightsail_disk" {
  name              = "${var.prefix}_awslightsail_disk"
  size_in_gb        = 8
  availability_zone = "us-east-1a"
}