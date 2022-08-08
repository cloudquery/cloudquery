terraform {
  backend "s3" {
    bucket         = "cq-provider-aws-tf"
    key            = "apigatewayv2"
    region         = "us-east-1"
  }
}
