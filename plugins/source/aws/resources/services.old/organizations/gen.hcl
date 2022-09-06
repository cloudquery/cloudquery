service          = "aws"
output_directory = "."
add_generate     = true

description_modifier "remove_read_only" {
  words = ["  This member is required."]
}

resource "aws" "organizations" "accounts" {
  path = "github.com/aws/aws-sdk-go-v2/service/organizations/types.Account"
  ignoreError "IgnoreAccessDenied" {
    path = "github.com/cloudquery/cloudquery/plugins/source/aws/client.IgnoreCommonErrors"
  }

  multiplex "AccountMultiplexer" {
    path = "github.com/cloudquery/cloudquery/plugins/source/aws/client.AccountMultiplex"
  }

  deleteFilter "DeleteAccountFilter" {
    path = "github.com/cloudquery/cloudquery/plugins/source/aws/client.DeleteAccountFilter"
  }

  options {
    primary_keys = ["account_id", "id"]
  }

  userDefinedColumn "account_id" {
    type        = "string"
    description = "The AWS Account ID of the resource."
    resolver "resolveAWSAccount" {
      path = "github.com/cloudquery/cloudquery/plugins/source/aws/client.ResolveAWSAccount"
    }
  }

  userDefinedColumn "tags" {
    type              = "json"
    description       = "The AWS tags of the resource."
    generate_resolver = true
  }
}