module "appsync" {
  source  = "terraform-aws-modules/appsync/aws"
  version = "1.5.2"
  name    = "${var.prefix}-appsync"
  tags    = var.tags
}
