resource "aws_eip" "elastic_ip" {
  instance = aws_instance.aws_ec2_instances_ec2_instance.id
  vpc      = true

  tags = {
    Name = "elastic-ip-${var.test_prefix}${var.test_suffix}"
  }
}