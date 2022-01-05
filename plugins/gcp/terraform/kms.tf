################################################################################
# KMS Module
################################################################################

module "kms" {
  source  = "terraform-google-modules/kms/google"
  version = "~> 1.2"

  project_id         = local.project
  location           = "us"
  keyring            = local.prefix
  keys               = ["key_for_db"]
  set_owners_for     = ["key_for_db"]
  owners = [
    "serviceAccount:${module.service_accounts.email}",
  ]

  labels = local.labels
}