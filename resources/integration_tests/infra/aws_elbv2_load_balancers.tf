resource "aws_lb" "aws_elbv2_load_balancer" {
  name = "elbv2-${var.test_suffix}"

  subnets = [aws_subnet.aws_vpc_subnet.id, aws_subnet.aws_vpc_subnet2.id]

  tags = {
    test = "test"
  }
}

resource "aws_lb_target_group" "aws_elbv2_lb_target_group" {
  name = "elbv2-tg-${var.test_suffix}"
  port = "80"
  protocol = "HTTP"
  vpc_id = aws_vpc.aws_vpc.id
  target_type = "ip"


  #STEP 1 - ECS task Running
  health_check {
    healthy_threshold = "3"
    interval = "10"
    port = "8080"
    path = "/index.html"
    protocol = "HTTP"
    unhealthy_threshold = "3"
  }

  depends_on = [aws_lb.aws_elbv2_load_balancer]
}

resource "aws_lb_listener" "aws_elbv2_lb_listener" {

  default_action {
    target_group_arn = aws_lb_target_group.aws_elbv2_lb_target_group.id
    type = "forward"
  }

  load_balancer_arn = aws_lb.aws_elbv2_load_balancer.arn
  port = "80"
  protocol = "HTTP"
}
