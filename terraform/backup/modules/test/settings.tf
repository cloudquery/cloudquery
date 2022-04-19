
// TODO - required org admin permission
#resource "aws_backup_global_settings" "backup_global_settings" {
#  global_settings = {
#    "isCrossAccountBackupEnabled" = "true"
#  }
#}

resource "aws_backup_region_settings" "backup_region_settings" {
  resource_type_opt_in_preference = {
    "Aurora"          = true
    "DocumentDB"      = true
    "DynamoDB"        = true
    "EBS"             = true
    "EC2"             = true
    "EFS"             = true
    "FSx"             = true
    "Neptune"         = true
    "RDS"             = true
    "Storage Gateway" = true
    "VirtualMachine"  = true
  }

  resource_type_management_preference = {
    "DynamoDB" = true
    "EFS"      = true
  }
}