resource "aws_xray_group" "xray-group" {
  group_name        = "${var.prefix}-xray-group"
  filter_expression = "responsetime > 5"

  tags = merge(
    { Name = "${var.prefix}-xray-group" },
    var.tags
  )
}