resource "aws_route53_health_check" "route53_health_check" {
  fqdn = "${var.test_prefix}${var.test_suffix}.com"
  port = 80
  type = "HTTP"
  resource_path = "/"
  failure_threshold = "5"
  request_interval = "10"

  tags = {
    Name = "health-check${var.test_prefix}${var.test_suffix}"
  }
}
