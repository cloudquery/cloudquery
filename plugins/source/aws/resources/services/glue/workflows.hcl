//check-for-changes
service          = "aws"
output_directory = "."
add_generate     = true

resource "aws" "glue" "workflows" {
  path = "github.com/aws/aws-sdk-go-v2/service/glue/types.Workflow"
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
    primary_keys = ["arn"]
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
  userDefinedColumn "arn" {
    type              = "string"
    description       = "The Amazon Resource Name (ARN) of the workflow."
    generate_resolver = true
  }

  column "graph" {
    skip = true
  }

  column "last_run_graph" {
    skip = true
  }

  userDefinedColumn "tags" {
    type              = "json"
    description       = "Resource tags."
    generate_resolver = true
  }

  column "last_run_starting_event_batch_condition_batch_size" {
    rename = "last_run_starting_event_batch_condition_size"
  }

  column "last_run_starting_event_batch_condition_batch_window" {
    rename = "last_run_starting_event_batch_condition_window"
  }

  column "blueprint_details_blueprint_name" {
    rename = "blueprint_name"
  }

  column "blueprint_details_run_id" {
    rename = "blueprint_run_id"
  }

  ignore_columns_in_tests = [
    "blueprint_details_blueprint_name",
    "blueprint_details_run_id",
    "default_run_properties",
    "last_run_completed_on",
    "last_run_error_message",
    "last_run_name",
    "last_run_previous_run_id",
    "last_run_started_on",
    "last_run_starting_event_batch_condition_batch_size",
    "last_run_starting_event_batch_condition_batch_window",
    "last_run_workflow_run_id",
    "last_run_workflow_run_properties",
  ]
}
