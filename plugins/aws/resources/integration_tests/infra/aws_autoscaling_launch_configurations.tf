data "aws_ami" "ubuntu" {
  most_recent = true

  filter {
    name   = "name"
    values = ["ubuntu/images/hvm-ssd/ubuntu-trusty-14.04-amd64-server-*"]
  }

  filter {
    name   = "virtualization-type"
    values = ["hvm"]
  }

  owners = ["099720109477"]
}

resource "aws_launch_configuration" "aws_lc" {
  name          = "lc-${var.test_prefix}-${var.test_suffix}"
  image_id      = data.aws_ami.ubuntu.id
  instance_type = "t2.nano"

  ebs_block_device {
    device_name = "ebs_block-${var.test_prefix}${var.test_suffix}"
    volume_type = "standard"
    volume_size = 5
  }
}
