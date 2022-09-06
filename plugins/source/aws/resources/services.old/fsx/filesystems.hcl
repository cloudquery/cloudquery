//check-for-changes
service          = "aws"
output_directory = "."
add_generate     = true

#description_modifier "remove_read_only" {
#  words = ["  This member is required."]
#}

resource "aws" "fsx" "filesystems" {
  path = "github.com/aws/aws-sdk-go-v2/service/fsx/types.FileSystem"
  ignoreError "IgnoreCommonErrors" {
    path = "github.com/cloudquery/cloudquery/plugins/source/aws/client.IgnoreCommonErrors"
  }
  deleteFilter "AccountRegionFilter" {
    path = "github.com/cloudquery/cloudquery/plugins/source/aws/client.DeleteAccountRegionFilter"
  }
  multiplex "AwsAccountRegion" {
    path   = "github.com/cloudquery/cloudquery/plugins/source/aws/client.ServiceAccountRegionMultiplexer"
    params = ["fsx"]
  }
  options {
    primary_keys = ["arn"]
  }
  userDefinedColumn "account_id" {
    description = "The AWS Account ID of the resource"
    type        = "string"
    resolver "resolveAWSAccount" {
      path = "github.com/cloudquery/cloudquery/plugins/source/aws/client.ResolveAWSAccount"
    }
  }
  userDefinedColumn "region" {
    type        = "string"
    description = "The AWS Region of the resource"
    resolver "resolveAWSRegion" {
      path = "github.com/cloudquery/cloudquery/plugins/source/aws/client.ResolveAWSRegion"
    }
  }

  column "file_system_id" {
    rename = "id"
  }

  column "file_system_type" {
    rename = "type"
  }

  column "file_system_type_version" {
    rename = "version"
  }

  column "resource_arn" {
    rename = "arn"
  }

  column "administrative_actions" {
    skip = true
  }

  column "tags" {
    type = "json"
    resolver "resolveTags" {
      path = "github.com/cloudquery/cloudquery/plugins/source/aws/client.ResolveTags"
    }
  }

  column "lustre_configuration" {
    skip = true
  }

  user_relation "aws" "fsx" "lustre_configuration" {
    path = "github.com/aws/aws-sdk-go-v2/service/fsx/types.LustreFileSystemConfiguration"
    resolver "resolveLustreConfiguration" {
      path = "github.com/cloudquery/cq-provider-sdk/provider/schema.PathTableResolver"
      params = ["LustreConfiguration"]
    }

    column "data_repository_configuration" {
      rename = "data_repo_cfg"
    }
  }

  column "ontap_configuration" {
    skip = true
  }

  user_relation "aws" "fsx" "ontap_configuration" {
    path = "github.com/aws/aws-sdk-go-v2/service/fsx/types.OntapFileSystemConfiguration"
    resolver "resolveOntapConfiguration" {
      path = "github.com/cloudquery/cq-provider-sdk/provider/schema.PathTableResolver"
      params = ["OntapConfiguration"]
    }
  }

  column "open_z_f_s_configuration" {
    skip = true
  }

  user_relation "aws" "fsx" "open_zfs_configuration" {
    path = "github.com/aws/aws-sdk-go-v2/service/fsx/types.OpenZFSFileSystemConfiguration"
    resolver "resolveOpenZfsConfiguration" {
      path = "github.com/cloudquery/cq-provider-sdk/provider/schema.PathTableResolver"
      params = ["OpenZFSConfiguration"]
    }
  }

  column "windows_configuration" {
    skip = true
  }

  user_relation "aws" "fsx" "windows_configuration" {
    path = "github.com/aws/aws-sdk-go-v2/service/fsx/types.WindowsFileSystemConfiguration"
    resolver "resolveWindowsConfiguration" {
      path = "github.com/cloudquery/cq-provider-sdk/provider/schema.PathTableResolver"
      params = ["WindowsConfiguration"]
    }

    column "self_managed_active_directory_configuration" {
      rename = "self_managed_ad_config"
    }

    column "aliases" {
      type = "json"
    }

    column "audit_log_configuration" {
      skip_prefix = true
    }
  }
}
