resource "aws_xray_sampling_rule" "xray_sampling_rule" {
  rule_name      = "${var.prefix}-xray-sampling-rule"
  priority       = 100
  version        = 1
  reservoir_size = 1
  fixed_rate     = 0.05
  url_path       = "*"
  host           = "*"
  http_method    = "*"
  service_type   = "*"
  service_name   = "*"
  resource_arn   = "*"

  attributes = {
    Hello = "Tris"
  }

  tags = merge(
    { Name = "${var.prefix}-xray-sampling-rule" },
    var.tags
  )
}