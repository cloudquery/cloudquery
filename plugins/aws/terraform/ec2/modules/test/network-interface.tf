resource "aws_network_interface" "nif" {
  subnet_id       = element(module.vpc.private_subnets, 1)
  tags = merge(
    {
      Name = "${var.prefix}-nif",
    },
    var.tags
  )
}