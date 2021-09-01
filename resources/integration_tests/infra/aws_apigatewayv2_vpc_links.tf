resource "aws_apigatewayv2_vpc_link" "aws_apigatewayv2_vpc_links_link" {
  name = "apigw-link-${var.test_prefix}-${var.test_suffix}"
  security_group_ids = [
  aws_security_group.aws_apigatewayv2_vpc_links_sg.id]
  subnet_ids = [
    aws_subnet.aws_vpc_subnet2.id,
  aws_subnet.aws_vpc_subnet3.id]
}

resource "aws_security_group" "aws_apigatewayv2_vpc_links_sg" {
  name   = "apigw-sg-${var.test_prefix}-${var.test_suffix}"
  vpc_id = aws_vpc.aws_vpc.id
  ingress {
    from_port = 22
    to_port   = 22
    protocol  = "tcp"
    cidr_blocks = [
    "0.0.0.0/0"]
  }
}
