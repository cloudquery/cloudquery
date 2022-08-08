################################################################################
# KMS Module
################################################################################

module "kms" {
  source  = "terraform-google-modules/kms/google"
  version = "~> 2.1.0"

  project_id         = var.project_id
  location           = "us"
  keyring            = "${var.prefix}-kms"
  keys               = ["key1"]
  # set_owners_for     = ["key_for_db"]
  # owners = [
    # "serviceAccount:${module.service_accounts.email}",
  # ]

  labels = var.labels
}