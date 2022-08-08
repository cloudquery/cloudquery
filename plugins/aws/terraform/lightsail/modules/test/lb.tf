resource "awslightsail_lb" "awslightsail_lb" {
  name              = "${var.prefix}_load_balancer"
  health_check_path = "/"
  instance_port     = "80"
}