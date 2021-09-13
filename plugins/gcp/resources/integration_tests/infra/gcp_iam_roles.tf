resource "google_project_iam_custom_role" "gcp_iam_roles_role" {
  role_id     = "gcp_iam_roles_${var.test_suffix}"
  title       = "title ${var.test_prefix}${var.test_suffix}"
  description = "A description"
  permissions = [
    "iam.roles.list",
    "iam.roles.create",
  "iam.roles.delete"]
}