data "aws_workspaces_bundle" "value_windows_10" {
  bundle_id = "wsb-gk1wpk43z" # wsb-gk1wpk43z,STANDARD,Windows 10 Experience powered by Windows Server 2019 with PCoIP 2 vCPU 4GiB Memory 50GB Storage (English)
}

resource "aws_workspaces_workspace" "workspace" {
  directory_id = aws_workspaces_directory.directory.id
  bundle_id    = data.aws_workspaces_bundle.value_windows_10.id
  user_name    = "Admin"

  root_volume_encryption_enabled = false
  user_volume_encryption_enabled = false

  workspace_properties {
    compute_type_name                         = "VALUE"
    user_volume_size_gib                      = 10
    root_volume_size_gib                      = 80
    running_mode                              = "AUTO_STOP"
    running_mode_auto_stop_timeout_in_minutes = 60
  }

  tags = {
    Department = "IT"
  }

  timeouts {
    create = "60m"
    update = "60m"
    delete = "30m"
  }
}