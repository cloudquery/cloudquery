module "cq_provider_aws_s3" {
  source        = "terraform-aws-modules/s3-bucket/aws"
  version       = "~> 2.14"
  bucket        = "cq-provider-aws-bucket"
  force_destroy = true
}