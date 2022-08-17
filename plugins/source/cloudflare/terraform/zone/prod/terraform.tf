terraform {
  backend "s3" {
    bucket = "cq-plugins-source-cloudflare"
    key    = "account"
    region = "us-east-1"
  }
}
