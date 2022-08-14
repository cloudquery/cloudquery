terraform {
  backend "gcs" {
    bucket = "cq-plugins-source-gcp-tf-state"
    prefix = "dns"
  }
}
