// https://github.com/terraform-aws-modules/terraform-aws-redshift
// https://github.com/hashicorp/terraform-provider-aws/issues/19110
resource "random_password" "redshift_password" {
  length           = 16
  special          = true
  override_special = "_%@"
}

###########
# Redshift
###########
module "redshift" {
  source  = "terraform-aws-modules/redshift/aws"

  cluster_identifier      = "cq-provider-aws"
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