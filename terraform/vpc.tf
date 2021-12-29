module "deny_all_sg" {
  source  = "terraform-aws-modules/security-group/aws"
  version = "~> 4"

  name        = "cq-provider-aws-deny-all-sg"
  description = "cq-provider-aws deny all sg"
  vpc_id      = module.vpc.vpc_id

}

module "vpc" {
  source  = "terraform-aws-modules/vpc/aws"
  version = "~> 3.0"

  name = "cq-provider-aws-vpc"
  cidr = "10.0.0.0/16"

  azs              = ["eu-central-1a", "eu-central-1b"]
  private_subnets = ["10.0.1.0/24", "10.0.2.0/24"]
//   public_subnets  = ["10.0.101.0/24", "10.0.102.0/24"]
  enable_ipv6 = true
  enable_nat_gateway = false
  create_egress_only_igw = false
}