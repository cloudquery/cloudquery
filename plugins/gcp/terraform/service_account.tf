################################################################################
# Service Account Module
################################################################################

module "service_accounts" {
  source       = "terraform-google-modules/service-accounts/google"
  version      = "~> 3.0"
  project_id   = local.project
  prefix       = "cq-sa-"
  display_name = "CQ Service Account"
  names        = ["single-account"]
}