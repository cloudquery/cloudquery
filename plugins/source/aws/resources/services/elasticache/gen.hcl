//check-for-changes

service          = "aws"
output_directory = "."
add_generate     = true

resource "aws" "elasticache" "clusters" {
  path = "github.com/aws/aws-sdk-go-v2/service/elasticache/types.CacheCluster"

  ignoreError "IgnoreCommonErrors" {
    path = "github.com/cloudquery/cloudquery/plugins/source/aws/client.IgnoreCommonErrors"
  }

  deleteFilter "AccountRegionFilter" {
    path = "github.com/cloudquery/cloudquery/plugins/source/aws/client.DeleteAccountRegionFilter"
  }

  multiplex "AwsAccountRegion" {
    path   = "github.com/cloudquery/cloudquery/plugins/source/aws/client.ServiceAccountRegionMultiplexer"
    params = ["elasticache"]
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

  column "auto_minor_version_upgrade" {
    description = "Auto minor version upgrade"
  }

  column "cache_cluster_create_time" {
    rename = "create_time"
  }

  column "cache_cluster_id" {
    rename = "id"
  }

  column "cache_cluster_status" {
    rename = "status"
  }

  column "notification_configuration_topic_arn" {
    description = "The arn of a notification topic used for publishing ElastiCache events to subscribers using Amazon Simple Notification Service (SNS)"
  }

  column "notification_configuration_topic_status" {
    description = "The current state of a notification topic used for publishing ElastiCache events to subscribers using Amazon Simple Notification Service (SNS)"
  }

  column "pending_modified_values_auth_token_status" {
    rename      = "pending_auth_token_status"
    description = "Auth token status that is applied to the cluster in the future or is currently being applied"
  }

  column "pending_modified_values_cache_node_ids_to_remove" {
    rename = "pending_cache_node_ids_to_remove"
  }

  column "pending_modified_values_cache_node_type" {
    rename = "pending_cache_node_type"
  }

  column "pending_modified_values_engine_version" {
    rename      = "pending_engine_version"
    description = "Cache engine version that is being applied to the cluster (or will be applied)"
  }

  column "pending_modified_values_num_cache_nodes" {
    rename = "pending_num_cache_nodes"
  }

  relation "aws" "elasticache" "cache_nodes" {
    column "cache_node_create_time" {
      rename = "create_time"
    }

    column "cache_node_id" {
      rename = "id"
    }

    column "cache_node_status" {
      rename = "status"
    }
  }

  relation "aws" "elasticache" "cache_security_groups" {
    column "cache_security_group_name" {
      rename = "name"
    }
  }

  relation "aws" "elasticache" "log_delivery_configurations" {
    column "destination_details_cloud_watch_logs_details_log_group" {
      rename      = "cloudwatch_destination_log_group"
      description = "The log group of the CloudWatch Logs destination"
    }

    column "destination_details_kinesis_firehose_details_delivery_stream" {
      rename      = "kinesis_firehose_destination_delivery_stream"
      description = "The Kinesis Data Firehose delivery stream of the Kinesis Data Firehose destination"
    }
  }

  # Actually skipping a relation.
  # Skipping because having subrelation is a bit much...
  column "pending_modified_values_log_delivery_configurations" {
    skip = true
  }
}