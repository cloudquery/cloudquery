################################################################################
# IAM Module - Private
################################################################################

module "iam-roles" {
  source = "terraform-google-modules/iam/google//modules/custom_role_iam"

  target_level         = "project"
  target_id            = local.project
  role_id              = "${replace(local.prefix, "-", "_")}_iam_role"
  title                = "CloudQuery IAM Custom Role"
  description          = "CloudQuery IAM Role"
  permissions          = ["iam.roles.list"]

  members              = [
    "serviceAccount:${module.service_accounts.email}"
  ]
}