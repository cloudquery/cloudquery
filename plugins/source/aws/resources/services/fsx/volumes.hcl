//check-for-changes
service          = "aws"
output_directory = "."
add_generate     = true

description_modifier "remove_read_only" {
  words = ["  This member is required."]
}

resource "aws" "fsx" "volumes" {
  path = "github.com/aws/aws-sdk-go-v2/service/fsx/types.Volume"
  ignoreError "IgnoreAccessDenied" {
    path = "github.com/cloudquery/cloudquery/plugins/source/aws/client.IgnoreAccessDeniedServiceDisabled"
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

  column "administrative_actions" {
    skip = true
  }

  column "volume_id" {
    rename = "id"
  }

  column "ontap_configuration" {
    skip = true
  }

  user_relation "aws" "fsx" "ontap_configuration" {
    path = "github.com/aws/aws-sdk-go-v2/service/fsx/types.OntapVolumeConfiguration"
    resolver "resolveOntapConfiguration" {
      path = "github.com/cloudquery/cq-provider-sdk/provider/schema.PathTableResolver"
      params = ["OntapConfiguration"]
    }

    column "ontap_volume_type" {
      rename = "volume_type"
    }
  }

  column "open_z_f_s_configuration" {
    skip = true
  }

  user_relation "aws" "fsx" "open_zfs_configuration" {
    path = "github.com/aws/aws-sdk-go-v2/service/fsx/types.OpenZFSVolumeConfiguration"
    resolver "resolveOpenZFSConfiguration" {
      path = "github.com/cloudquery/cq-provider-sdk/provider/schema.PathTableResolver"
      params = ["OpenZFSConfiguration"]
    }

    column "origin_snapshot_snapshot_arn" {
      rename = "origin_snapshot_arn"
    }

    column "record_size_ki_b" {
      rename = "record_size"
    }

    column "storage_capacity_quota_gi_b" {
      rename = "storage_capacity_quota"
    }

    column "storage_capacity_reservation_gi_b" {
      rename = "storage_capacity_reservation"
    }

    column "nfs_exports" {
      type = "json"
    }

    column "user_and_group_quotas" {
      type = "json"
    }
  }
}
