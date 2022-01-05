################################################################################
# CloudQuery GCP Provider Local Variables
################################################################################

locals {
  project = "cq-provider-gcp"
  region  = "us-central1"
  prefix  = "cq-integration-test"

  domain = "cq.example.com"

  labels = {
    "integration-test" = "true"
    "environment"      = "dev"
  }

  subnet_ids = [
    "10.10.10.0",
    "10.10.20.0",
    "10.10.30.0"
  ]

  tags = [
    "integrationtest",
    "environmentdev"
  ]
}