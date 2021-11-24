
resource "aws_eip" "lb" {
  instance = aws_instance.aws_ec2_instances_ec2_instance.id
  vpc      = true
}