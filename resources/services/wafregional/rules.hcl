service = "aws"

output_directory = "."

resource "aws" "wafregional" "rules" {
  path = "github.com/aws/aws-sdk-go-v2/service/wafregional/types.Rule"
  multiplex "ServiceAccountRegionMultiplexer" {
    path = "github.com/cloudquery/cq-provider-aws/client.ServiceAccountRegionMultiplexer"
  }
  ignoreError "IgnoreAccessDenied" {
    path = "github.com/cloudquery/cq-provider-aws/client.IgnoreAccessDeniedServiceDisabled"
  }
  deleteFilter "AccountRegionFilter" {
    path = "github.com/cloudquery/cq-provider-aws/client.DeleteAccountRegionFilter"
  }

  options {
    primary_keys = ["account_id", "region", "id"]
  }

  userDefinedColumn "account_id" {
    description = "The AWS Account ID of the resource."
    type        = "string"
    resolver "resolveAWSAccount" {
      path = "github.com/cloudquery/cq-provider-aws/client.ResolveAWSAccount"
    }
  }
  userDefinedColumn "region" {
    description = "The AWS Region of the resource."
    type        = "string"
    resolver "resolveAWSRegion" {
      path = "github.com/cloudquery/cq-provider-aws/client.ResolveAWSRegion"
    }
  }

  column "rule_id" {
    rename = "id"
  }

  userDefinedColumn "arn" {
    type = "string"
    description = "ARN of the rule."
    generate_resolver = true
  }

  userDefinedColumn "tags" {
    type = "json"
    generate_resolver = true
    description = "Rule tags."
  }
}
