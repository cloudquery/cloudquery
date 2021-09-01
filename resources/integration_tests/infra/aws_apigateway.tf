resource "aws_api_gateway_vpc_link" "apigateway_vpc_link" {
  name        = "apigw-vpc-link-${var.test_prefix}-${var.test_suffix}"
  description = "example description"
  target_arns = [
  aws_lb.apigateway_nlb.arn]
}

resource "aws_lb" "apigateway_nlb" {
  name               = "apigateway-nlb-${var.test_suffix}"
  internal           = false
  load_balancer_type = "network"
  subnets = [
    aws_subnet.aws_vpc_subnet2.id,
  aws_subnet.aws_vpc_subnet3.id]

  enable_deletion_protection = false

  tags = {
    Environment = "dev"
  }
}
