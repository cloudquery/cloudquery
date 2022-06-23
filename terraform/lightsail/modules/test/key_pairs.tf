# Create a new Lightsail Key Pair
resource "aws_lightsail_key_pair" "aws_lightsail_key_pair" {
  name = "${var.prefix}_key_pair"
}