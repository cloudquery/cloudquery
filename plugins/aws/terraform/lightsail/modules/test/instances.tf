# Create a new GitLab Lightsail Instance
resource "aws_lightsail_instance" "aws_lightsail_instance" {
  name              = "${var.prefix}-lightsaleinstance"
  availability_zone = "us-east-1b"
  blueprint_id      = "amazon_linux_2"
  bundle_id         = "nano_2_0"
  key_pair_name     = aws_lightsail_key_pair.aws_lightsail_key_pair.name
}

resource "aws_lightsail_static_ip_attachment" "test" {
  static_ip_name = aws_lightsail_static_ip.aws_lightsail_static_ip.id
  instance_name  = aws_lightsail_instance.aws_lightsail_instance.id
}


resource "awslightsail_lb_attachment" "test" {
  load_balancer_name = awslightsail_lb.awslightsail_lb.name
  instance_name      = aws_lightsail_instance.aws_lightsail_instance.name
}
