service = "aws"
output_directory = "."
add_generate = true



resource "aws" "cloudformation" "stacks" {
  path = "github.com/aws/aws-sdk-go-v2/service/cloudformation/types.Stack"
  ignoreError "IgnoreAccessDenied" {
    path = "github.com/cloudquery/cq-provider-aws/client.IgnoreAccessDeniedServiceDisabled"
  }
  multiplex "AwsAccountRegion" {
    path = "github.com/cloudquery/cq-provider-aws/client.ServiceAccountRegionMultiplexer"
    params = ["cloudformation"]
  }

  deleteFilter "AccountRegionFilter" {
    path = "github.com/cloudquery/cq-provider-aws/client.DeleteAccountRegionFilter"
  }

  options {
    primary_keys = ["id"]
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
    type = "string"
    resolver "resolveStackArn" {
      path = "github.com/cloudquery/cq-provider-sdk/provider/schema.PathResolver"
      path_resolver = true
      // TODO: require manual changing from ARN -> StackId for the path resolver as its not supported by cq-gen yet
      params = ["StackId"]
    }
  }
  column "stack_id" {
    rename = "id"
  }

  column "stack_name" {
    rename = "stack"
  }

  column "notification_arn_s" {
    rename = "notification_arns"
  }

  column "stack_status" {
    rename = "status"
  }

  column "parameters" {
    type = "json"
  }

  column "drift_information_stack_drift_status" {
    rename = "stack_drift_status"
  }

  column "drift_information_last_check_timestamp" {
    rename = "drift_last_check_timestamp"
  }

  column "tags" {
    type = "json"
    generate_resolver = true
  }

  column "rollback_configuration_rollback_triggers" {
    type = "json"
  }


  user_relation "aws" "cloudformation" "resources" {
    path = "github.com/aws/aws-sdk-go-v2/service/cloudformation/types.StackResourceSummary"

    column "drift_information_stack_resource_drift_status" {
      rename = "stack_resource_drift_status"
    }

    column "drift_information_last_check_timestamp" {
      rename = "drift_last_check_timestamp"
    }
  }

}


