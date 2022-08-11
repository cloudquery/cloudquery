service          = "aws"
output_directory = "."
add_generate     = true

resource "aws" "cloudwatchlogs" "log_groups" {
  path = "github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs/types.LogGroup"
  ignoreError "IgnoreAccessDenied" {
    path = "github.com/cloudquery/cq-provider-aws/client.IgnoreAccessDeniedServiceDisabled"
  }
  deleteFilter "AccountRegionFilter" {
    path = "github.com/cloudquery/cq-provider-aws/client.DeleteAccountRegionFilter"
  }
  multiplex "AwsAccountRegion" {
    path   = "github.com/cloudquery/cq-provider-aws/client.ServiceAccountRegionMultiplexer"
    params = ["logs"]
  }

  userDefinedColumn "account_id" {
    description = "The AWS Account ID of the resource."
    type        = "string"
    resolver "resolveAWSAccount" {
      path = "github.com/cloudquery/cq-provider-aws/client.ResolveAWSAccount"
    }
  }

  userDefinedColumn "region" {
    type        = "string"
    description = "The AWS Region of the resource."
    resolver "resolveAWSRegion" {
      path = "github.com/cloudquery/cq-provider-aws/client.ResolveAWSRegion"
    }
  }

  userDefinedColumn "tags" {
    type              = "json"
    generate_resolver = true
    description       = "The tags for the log group."
  }

  ignore_columns_in_tests = ["kms_key_id","retention_in_days"]

  options {
    primary_keys = ["arn"]
  }

}