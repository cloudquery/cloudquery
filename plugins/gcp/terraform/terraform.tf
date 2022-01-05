terraform {
  backend "gcs" {
    bucket = "cq-provider-gcp-tf-state"
    prefix = "tf-state"
  }
}


provider "google" {
  project = local.project
  region  = local.region
}