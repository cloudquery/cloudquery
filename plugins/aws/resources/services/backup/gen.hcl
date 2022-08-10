service          = "aws"
output_directory = "."
add_generate     = true

resource "aws" "backup" "global_settings" {
  path = "github.com/aws/aws-sdk-go-v2/service/backup.DescribeGlobalSettingsOutput"
  ignoreError "IgnoreAccessDenied" {
    path = "github.com/cloudquery/cloudquery/plugins/aws/client.IgnoreAccessDeniedServiceDisabled"
  }
  deleteFilter "AccountFilter" {
    path = "github.com/cloudquery/cloudquery/plugins/aws/client.DeleteAccountFilter"
  }
  multiplex "AwsAccount" {
    path = "github.com/cloudquery/cloudquery/plugins/aws/client.AccountMultiplex"
  }


  options {
    primary_keys = ["account_id"]
  }

  userDefinedColumn "account_id" {
    description = "The AWS Account ID of the resource."
    type        = "string"
    resolver "resolveAWSAccount" {
      path = "github.com/cloudquery/cloudquery/plugins/aws/client.ResolveAWSAccount"
    }
  }

  column "result_metadata" {
    skip = true
  }
}

resource "aws" "backup" "region_settings" {
  path = "github.com/aws/aws-sdk-go-v2/service/backup.DescribeRegionSettingsOutput"
  ignoreError "IgnoreAccessDenied" {
    path = "github.com/cloudquery/cloudquery/plugins/aws/client.IgnoreAccessDeniedServiceDisabled"
  }
  deleteFilter "AccountRegionFilter" {
    path = "github.com/cloudquery/cloudquery/plugins/aws/client.DeleteAccountRegionFilter"
  }
  multiplex "AwsAccountRegion" {
    path   = "github.com/cloudquery/cloudquery/plugins/aws/client.ServiceAccountRegionMultiplexer"
    params = ["backup"]
  }


  options {
    primary_keys = ["account_id", "region"]
  }

  userDefinedColumn "account_id" {
    description = "The AWS Account ID of the resource."
    type        = "string"
    resolver "resolveAWSAccount" {
      path = "github.com/cloudquery/cloudquery/plugins/aws/client.ResolveAWSAccount"
    }
  }

  userDefinedColumn "region" {
    type        = "string"
    description = "The AWS Region of the resource."
    resolver "resolveAWSRegion" {
      path = "github.com/cloudquery/cloudquery/plugins/aws/client.ResolveAWSRegion"
    }
  }

  column "result_metadata" {
    skip = true
  }
}
