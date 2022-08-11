// https://github.com/terraform-aws-modules/terraform-aws-redshift
// https://github.com/hashicorp/terraform-provider-aws/issues/19110
resource "random_password" "redshift_password" {
  length           = 16
  // special          = true
  // override_special = "_%@"
}

module "vpc" {
  source  = "terraform-aws-modules/vpc/aws"
  version = "~> 3.0"

  name = "${var.prefix}-redshift"
  cidr = "10.0.0.0/16"

  azs             = ["us-east-1a", "us-east-1b"]
  private_subnets = ["10.0.1.0/24", "10.0.2.0/24"]
  //   public_subnets  = ["10.0.101.0/24", "10.0.102.0/24"]
  enable_ipv6            = true
  enable_nat_gateway     = false
  create_egress_only_igw = false
}

module "deny_all_sg" {
  source  = "terraform-aws-modules/security-group/aws"
  version = "~> 4.8"

  name        = "${var.prefix}-redshift-deny-all-sg"
  description = "${var.prefix}-redshift deny all sg"
  vpc_id      = module.vpc.vpc_id

}

###########
# Redshift
###########
module "redshift" {
  source  = "terraform-aws-modules/redshift/aws"
  version = "~> 3.4"

  cluster_identifier      = "${var.prefix}-redshift"
  cluster_node_type       = "dc2.large"
  cluster_number_of_nodes = 1

  cluster_database_name   = "mydb"
  cluster_master_username = "mydbuser"
  cluster_master_password = random_password.redshift_password.result

  subnets                = module.vpc.private_subnets
  vpc_security_group_ids = [module.deny_all_sg.security_group_id]
  # Group parameters
  wlm_json_configuration = "[]"
  #  redshift_subnet_group_name = module.vpc.redshift_subnet_group
}

resource "aws_sns_topic" "sns_redshift" {
  name = "${var.prefix}-redshift-events"
}

resource "aws_redshift_event_subscription" "default" {
  name          = "${var.prefix}-redshift-event-sub"
  sns_topic_arn = aws_sns_topic.sns_redshift.arn

  source_type = "cluster"
  source_ids  = [module.redshift.redshift_cluster_id]

  severity = "INFO"

  event_categories = [
    "configuration",
    "management",
    "monitoring",
    "security",
  ]
}
