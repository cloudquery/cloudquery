module "cq_provider_aws_s3" {
  source        = "terraform-aws-modules/s3-bucket/aws"
  version       = "~> 2.14"
  bucket        = "${var.prefix}-s3-cq-provider-aws"
  force_destroy = true
}