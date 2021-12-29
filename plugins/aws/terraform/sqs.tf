// https://github.com/terraform-aws-modules/terraform-aws-sqs

module "sqs" {
  source  = "terraform-aws-modules/sqs/aws"
  version = "~> 2.0"

  name = "user"
}