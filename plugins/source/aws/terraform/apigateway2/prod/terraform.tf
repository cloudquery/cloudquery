terraform {
  backend "s3" {
    bucket         = "cq-plugins-source-aws-tf"
    key            = "apigatewayv2"
    region         = "us-east-1"
  }
}
