service          = "aws"
output_directory = "."
add_generate     = true

description_modifier "remove_read_only" {
  words = ["  This member is required."]
}

resource "aws" "glue" "triggers" {
  path = "github.com/aws/aws-sdk-go-v2/service/glue/types.Trigger"
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
    type        = "string"
    description = "The Amazon Resource Name (ARN) of the trigger."
    generate_resolver = true
  }

  userDefinedColumn "tags" {
    type = "json"
    description = "Resource tags."
    generate_resolver = true
  }

  column "event_batching_condition_batch_size" {
    rename = "event_batching_condition_size"
  }

  column "event_batching_condition_batch_window" {
    rename = "event_batching_condition_window"
  }

  relation "aws" "glue" "actions" {
    column "notification_property" {
        skip_prefix = true
    }
  }

  relation "aws" "glue" "predicate_conditions" {
    ignore_in_tests = true
  }
}
