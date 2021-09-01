resource "aws_lb_target_group" "elbv2_target_groups_tg" {
  name     = "lbv2target${var.test_prefix}"
  port     = 80
  protocol = "HTTP"
  vpc_id   = aws_vpc.aws_vpc.id
  tags = {
    test = "test"
  }
}