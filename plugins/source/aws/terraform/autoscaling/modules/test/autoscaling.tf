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

  owners = ["099720109477"] # Canonical
}

resource "aws_launch_configuration" "launch_configuration" {
  name          = "${var.prefix}-launch-configuration"
  image_id      = data.aws_ami.ubuntu.id
  instance_type = "t4g.nano"
}

resource "aws_autoscaling_group" "autoscaling_group" {
  availability_zones        = ["us-east-1a", "us-east-1b"]
  name                      = "${var.prefix}-autoscaling-group"
  max_size                  = 0
  min_size                  = 0
  desired_capacity = 0
  health_check_grace_period = 300
  health_check_type         = "ELB"
  force_delete              = true
  termination_policies      = ["OldestInstance"]
  launch_configuration      = aws_launch_configuration.launch_configuration.name
}

resource "aws_autoscaling_schedule" "autoscaling_schedule" {
  scheduled_action_name  = "${var.prefix}-autoscaling-schedule"
  min_size               = 0
  max_size               = 1
  desired_capacity       = 0
  start_time             = "2999-12-11T18:00:00Z"
  end_time               = "2999-12-12T06:00:00Z"
  autoscaling_group_name = aws_autoscaling_group.autoscaling_group.name
}
