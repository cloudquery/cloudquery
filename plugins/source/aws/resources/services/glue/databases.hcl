service          = "aws"
output_directory = "."
add_generate     = true

description_modifier "remove_read_only" {
  words = ["  This member is required."]
}

resource "aws" "glue" "databases" {
  path = "github.com/aws/aws-sdk-go-v2/service/glue/types.Database"
  ignoreError "IgnoreAccessDenied" {
    path = "github.com/cloudquery/cq-provider-aws/client.IgnoreAccessDeniedServiceDisabled"
  }
  deleteFilter "AccountRegionFilter" {
    path = "github.com/cloudquery/cq-provider-aws/client.DeleteAccountRegionFilter"
  }
  multiplex "AwsAccountRegion" {
    path   = "github.com/cloudquery/cq-provider-aws/client.ServiceAccountRegionMultiplexer"
    params = ["glue"]
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
    type        = "string"
    description = "The AWS Region of the resource."
    resolver "resolveAWSRegion" {
      path = "github.com/cloudquery/cq-provider-aws/client.ResolveAWSRegion"
    }
  }
  userDefinedColumn "arn" {
    type              = "string"
    description       = "The Amazon Resource Name (ARN) of the workflow."
    generate_resolver = true
  }

  user_relation "aws" "glue" "tables" {
    path = "github.com/aws/aws-sdk-go-v2/service/glue/types.Table"

    column "storage_descriptor" {
      type = "json"
    }
  }

  column "target_database_database_name" {
    rename = "target_database_name"
  }

  column "create_table_default_permissions" {
    type = "json"
  }
}
