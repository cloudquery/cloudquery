//check-for-changes
service          = "aws"
output_directory = "."
add_generate     = true

description_modifier "remove_read_only" {
  words = ["  This member is required."]
}

resource "aws" "applicationautoscaling" "policies" {
  path = "github.com/aws/aws-sdk-go-v2/service/applicationautoscaling/types.ScalingPolicy"
  ignoreError "IgnoreCommonErrors" {
    path = "github.com/cloudquery/cloudquery/plugins/source/aws/client.IgnoreCommonErrors"
  }
  deleteFilter "AccountRegionFilter" {
    path = "github.com/cloudquery/cloudquery/plugins/source/aws/client.DeleteAccountRegionFilter"
  }
  multiplex "ServiceAccountRegionNamespaceMultiplexer" {
    path   = "github.com/cloudquery/cloudquery/plugins/source/aws/client.ServiceAccountRegionNamespaceMultiplexer"
    params = ["application-autoscaling"]
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
  userDefinedColumn "namespace" {
    type        = "string"
    description = "The AWS Service Namespace of the resource"
    resolver "resolveAWSNamespace" {
      path = "github.com/cloudquery/cloudquery/plugins/source/aws/client.ResolveAWSNamespace"
    }
  }

  column "policy_arn" {
    rename = "arn"
  }
  column "policy_name" {
    rename = "name"
  }
  column "policy_type" {
    rename = "type"
  }

  column "step_scaling_policy_configuration" {
    type = "json"
  }

  column "target_tracking_scaling_policy_configuration" {
    type = "json"
  }

  column "alarms" {
    type = "json"
  }
}
