// TODO - this resource is currently unavailable due to high costs

#resource "aws_ec2_host" "ec2-dedicated-host" {
#  instance_type     = "a1.medium"
#  availability_zone = element(module.vpc.azs, 1)
#  host_recovery     = "on"
#  auto_placement    = "on"
#
#  tags = merge(
#    {
#      Name = "${var.prefix}-ec2-dedicated-host"
#    },
#    var.tags
#  )
#}