//check-for-changes
service          = "aws"
output_directory = "."
add_generate     = true

description_modifier "remove_read_only" {
  words = ["  This member is required."]
}

resource "aws" "fsx" "storage_vms" {
  path = "github.com/aws/aws-sdk-go-v2/service/fsx/types.StorageVirtualMachine"
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

  column "resource_arn" {
    rename = "arn"
  }

  column "tags" {
    type = "json"
    resolver "resolveTags" {
      path = "github.com/cloudquery/cloudquery/plugins/source/aws/client.ResolveTags"
    }
  }

  column "active_directory_configuration" {
    rename = "ad_cfg"
  }

  column "ad_cfg_self_managed_active_directory_configuration" {
    skip_prefix = true
  }

  column "storage_virtual_machine_id" {
    rename = "id"
  }
}
