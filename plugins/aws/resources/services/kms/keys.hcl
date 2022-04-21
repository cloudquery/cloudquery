service = "aws"

output_directory = "."

resource "aws" "kms" "keys" {
  path = "github.com/aws/aws-sdk-go-v2/service/kms/types.KeyMetadata"
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
    primary_keys = ["arn"]
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

  column "key_manager" {
      rename = "manager"
  }

  column "multi_region_configuration" {
      skip_prefix = true
  }

  column "replica_keys" {
      type = "json"
      generate_resolver = true
  }

  column "customer_master_key_spec" {
      skip = "true"
  }

  column "key_id" {
    rename = "id"
  }

  userDefinedColumn "rotation_enabled" {
      description = "Specifies whether key rotation is enabled."
      type = "bool"
      generate_resolver = true
  }

  userDefinedColumn "tags" {
      description = "Key tags."
      type = "json"
      generate_resolver = true
  }
}
