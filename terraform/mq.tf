module "mq_broker" {
  source  = "cloudposse/mq-broker/aws"
  version = "0.15.0"

  namespace                  = "eg"
  stage                      = "test"
  name                       = "mq-broker"
  apply_immediately          = true
  auto_minor_version_upgrade = true
  deployment_mode            = "ACTIVE_STANDBY_MULTI_AZ"
  engine_type                = "ActiveMQ"
  engine_version             = "5.15.15"
  host_instance_type         = "mq.t3.micro"
  publicly_accessible        = false
  general_log_enabled        = true
  audit_log_enabled          = true
  encryption_enabled         = true
  use_aws_owned_key          = true
  vpc_id                     = module.vpc.vpc_id
  subnet_ids                 = module.vpc.private_subnets
  // security_groups            = module.deny_all_sg.security_group_id
}
