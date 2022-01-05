################################################################################
# Bucket Module
################################################################################

module "gcs_buckets" {
  source     = "terraform-google-modules/cloud-storage/google"
  version    = "~> 3.1.0"
  project_id = local.project
  names      = ["storage-bucket"]
  prefix     = local.prefix
  location   = local.region

  versioning = {
    first = true
  }

  bucket_policy_only = {
    "storage-bucket" = false
  }

  lifecycle_rules = [
    {
      action    = {
        type = "Delete"
      }
      condition = {
        age = "10"
      }
    }
  ]

  labels = local.labels
}