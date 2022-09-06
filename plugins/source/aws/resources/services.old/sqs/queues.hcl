//check-for-changes
service          = "aws"
output_directory = "."
add_generate     = true

resource "aws" "sqs" "queues" {
  path = "github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/sqs.Queue"
  ignoreError "IgnoreAccessDenied" {
    path = "github.com/cloudquery/cloudquery/plugins/source/aws/client.IgnoreCommonErrors"
  }
  deleteFilter "AccountRegionFilter" {
    path = "github.com/cloudquery/cloudquery/plugins/source/aws/client.DeleteAccountRegionFilter"
  }
  multiplex "AwsAccountRegion" {
    path   = "github.com/cloudquery/cloudquery/plugins/source/aws/client.ServiceAccountRegionMultiplexer"
    params = ["sqs"]
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

  userDefinedColumn "tags" {
    type = "json"
    generate_resolver = true
  }

  column "policy" {
    type = "json"
  }

  column "redrive_policy" {
    type = "json"
  }

  column "redrive_allow_policy" {
    type = "json"
  }
}
