resource "aws_route53_health_check" "hch" {
  fqdn = "${var.test_prefix}${var.test_suffix}.com"
  port = 80
  type = "HTTP"
  resource_path = "/"
  failure_threshold = "5"
  request_interval = "10"

  tags = {
    Name = "tf-test-health-check1"
  }
}
