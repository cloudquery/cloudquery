resource "aws_lightsail_static_ip" "aws_lightsail_static_ip" {
  name = "${var.prefix}_static_ip"
}