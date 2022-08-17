//check-for-changes
service          = "aws"
output_directory = "."
add_generate     = true


description_modifier "remove_read_only" {
  words = ["  This member is required."]
}

resource "aws" "glue" "security_configurations" {
  path = "github.com/aws/aws-sdk-go-v2/service/glue/types.SecurityConfiguration"
  ignoreError "IgnoreAccessDenied" {
    path = "github.com/cloudquery/cloudquery/plugins/source/aws/client.IgnoreAccessDeniedServiceDisabled"
  }
  deleteFilter "AccountRegionFilter" {
    path = "github.com/cloudquery/cloudquery/plugins/source/aws/client.DeleteAccountRegionFilter"
  }
  multiplex "AwsAccountRegion" {
    path   = "github.com/cloudquery/cloudquery/plugins/source/aws/client.ServiceAccountRegionMultiplexer"
    params = ["glue"]
  }
  options {
    primary_keys = ["account_id", "region", "name"]
  }
  userDefinedColumn "account_id" {
    description = "The AWS Account ID of the resource."
    type        = "string"
    resolver "resolveAWSAccount" {
      path = "github.com/cloudquery/cloudquery/plugins/source/aws/client.ResolveAWSAccount"
    }
  }
  userDefinedColumn "region" {
    type        = "string"
    description = "The AWS Region of the resource."
    resolver "resolveAWSRegion" {
      path = "github.com/cloudquery/cloudquery/plugins/source/aws/client.ResolveAWSRegion"
    }
  }
  column "encryption_configuration" {
    skip_prefix = true
  }
}
