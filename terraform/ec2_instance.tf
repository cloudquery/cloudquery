
// This is based on this module
// https://github.com/terraform-aws-modules/terraform-aws-ec2-instance/blob/master/examples/complete/main.tf



resource "aws_kms_key" "ec2_kms_key" {
}

module "ec2_instance" {
  source  = "terraform-aws-modules/ec2-instance/aws"
  version = "~> 3.4"

  name = "cq-provider-aws-ec2-instance"
  // create_spot_instance = true
  // spot_price           = "0.60"
  // spot_type            = "persistent"

  ami                    = "ami-05d34d340fb1d89e5"
  instance_type          = "t2.micro"
  availability_zone      = element(module.vpc.azs, 0)
  subnet_id              = element(module.vpc.private_subnets, 0)
  vpc_security_group_ids = [module.deny_all_sg.security_group_id]
  //   placement_group             = aws_placement_group.web.id
  associate_public_ip_address = true

  # only one of these can be enabled at a time
  hibernation = true
  # enclave_options_enabled = true

  //   user_data_base64 = base64encode(local.user_data)

  // cpu_core_count       = 2 # default 4
  // cpu_threads_per_core = 1 # default 2

  capacity_reservation_specification = {
    capacity_reservation_preference = "open"
  }

  enable_volume_tags = false
  root_block_device = [
    {
      encrypted   = true
      volume_type = "gp3"
      throughput  = 200
      volume_size = 50
      tags = {
        Name = "my-root-block"
      }
    },
  ]

}