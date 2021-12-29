module "cq_provider_aws_s3" {
  source = "terraform-aws-modules/s3-bucket/aws"
  bucket        = "cq-provider-aws-bucket"
  force_destroy = true
}