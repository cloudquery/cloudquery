terraform {
  backend "gcs" {
    bucket = "cq-provider-gcp-tf-state"
    prefix = "kms"
  }
}