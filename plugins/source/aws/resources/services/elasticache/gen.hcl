//check-for-changes

service = "aws"
output_directory = "."
add_generate = true

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
    rename = "pending_auth_token_status"
    description = "Auth token status that is applied to the cluster in the future or is currently being applied"
  }

  column "pending_modified_values_cache_node_ids_to_remove" {
    rename = "pending_cache_node_ids_to_remove"
  }

  column "pending_modified_values_cache_node_type" {
    rename = "pending_cache_node_type"
  }

  column "pending_modified_values_engine_version" {
    rename = "pending_engine_version"
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
      rename = "cloudwatch_destination_log_group"
      description = "The log group of the CloudWatch Logs destination"
    }

    column "destination_details_kinesis_firehose_details_delivery_stream" {
      rename = "kinesis_firehose_destination_delivery_stream"
      description = "The Kinesis Data Firehose delivery stream of the Kinesis Data Firehose destination"
    }
  }

  # Actually skipping a relation.
  # Skipping because having subrelation is a bit much...
  column "pending_modified_values_log_delivery_configurations" {
    skip = true
  }
}

resource "aws" "elasticache" "engine_versions" {
  path = "github.com/aws/aws-sdk-go-v2/service/elasticache/types.CacheEngineVersion"

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
    primary_keys = [
      "account_id", "region", "engine", "engine_version"
    ]
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

  column "cache_parameter_group_family" {
    description = "The name of the cache parameter group family associated with this cache engine. Valid values are: memcached1.4 \\| memcached1.5 \\| memcached1.6 \\| redis2.6 \\| redis2.8 \\| redis3.2 \\| redis4.0 \\| redis5.0 \\| redis6.x"
  }
}


resource "aws" "elasticache" "parameter_groups" {
  path = "github.com/aws/aws-sdk-go-v2/service/elasticache/types.CacheParameterGroup"

  description = "Provides details about Elasticache parameter groups."

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
    primary_keys = [
      "arn"
    ]
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

  user_relation "aws" "elasticache" "parameters" {
    path = "github.com/aws/aws-sdk-go-v2/service/elasticache/types.Parameter"
  }
}

resource "aws" "elasticache" "subnet_groups" {
  path = "github.com/aws/aws-sdk-go-v2/service/elasticache/types.CacheSubnetGroup"

  description = "Contains information about cache subnet groups"

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
    primary_keys = [
      "arn"
    ]
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
}

resource "aws" "elasticache" "global_replication_groups" {
  path = "github.com/aws/aws-sdk-go-v2/service/elasticache/types.GlobalReplicationGroup"

  ignoreError "IgnoreCommonErrors" {
    path = "github.com/cloudquery/cloudquery/plugins/source/aws/client.IgnoreCommonErrors"
  }

  multiplex "AwsAccountMultiplex" {
    path = "github.com/cloudquery/cloudquery/plugins/source/aws/client.AccountMultiplex"
  }
  deleteFilter "AccountRegionFilter" {
    path = "github.com/cloudquery/cloudquery/plugins/source/aws/client.DeleteAccountFilter"
  }

  options {
    primary_keys = [
      "arn"
    ]
  }

  userDefinedColumn "account_id" {
    description = "The AWS Account ID of the resource."
    type        = "string"
    resolver "resolveAWSAccount" {
      path = "github.com/cloudquery/cloudquery/plugins/source/aws/client.ResolveAWSAccount"
    }
  }
}

resource "aws" "elasticache" "replication_groups" {
  path = "github.com/aws/aws-sdk-go-v2/service/elasticache/types.ReplicationGroup"

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
    primary_keys = [
      "arn"
    ]
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
    description = "Auto minor version upgrade."
  }

  column "global_replication_group_info_global_replication_group_member_role" {
    rename = "global_replication_group_member"
  }

  column "global_replication_group_info_global_replication_group_id" {
    rename = "global_replication_group_id"
  }

  column "pending_modified_values_auth_token_status" {
    rename = "pending_auth_token_status"
    description = "Pending modified auth token status"
  }

  column "pending_modified_values_automatic_failover_status" {
    rename = "pending_automatic_failover_status"
    description = "pending autmatic failover for this redis replication group"
  }

  # Actually skipping a relation.
  column "pending_modified_values_log_delivery_configurations" {
    skip = true
  }

  column "pending_modified_values_primary_cluster_id" {
    rename = "pending_primary_cluster_id"
  }

  column "pending_modified_values_resharding_slot_migration_progress_percentage" {
    rename = "pending_resharding_slot_migration_progress_percentage"
  }

  column "pending_modified_values_user_groups_user_group_ids_to_add" {
    rename = "pending_user_group_ids_to_add"
  }

  column "pending_modified_values_user_groups_user_group_ids_to_remove" {
    rename = "pending_user_group_ids_to_remove"
  }

  relation "aws" "elasticache" "log_delivery_configurations" {
    column "destination_details_cloud_watch_logs_details_log_group" {
      rename = "cloudwatch_destination_log_group"
      description = "The log group of the CloudWatch Logs destination"
    }

    column "destination_details_kinesis_firehose_details_delivery_stream" {
      rename = "kinesis_firehose_destination_delivery_stream"
      description = "The Kinesis Data Firehose delivery stream of the Kinesis Data Firehose destination"
    }
  }

  relation "aws" "elasticache" "node_groups" {
    relation "aws" "elasticache" "node_group_members" {
      rename = "members"
    }
  }
}

resource "aws" "elasticache" "reserved_cache_nodes" {
  path = "github.com/aws/aws-sdk-go-v2/service/elasticache/types.ReservedCacheNode"

  description = "Reserved Elasticache Cache Nodes"

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
    primary_keys = [
      "reservation_arn"
    ]
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
}

resource "aws" "elasticache" "reserved_cache_nodes_offerings" {
  path = "github.com/aws/aws-sdk-go-v2/service/elasticache/types.ReservedCacheNodesOffering"

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
    primary_keys = [
      "account_id",
      "region",
      "reserved_cache_nodes_offering_id"
    ]
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
}

resource "aws" "elasticache" "service_updates" {
  path = "github.com/aws/aws-sdk-go-v2/service/elasticache/types.ServiceUpdate"

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
    primary_keys = [
      "account_id",
      "region",
      "name"
    ]
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

  column "service_update_name" {
    rename = "name"
  }

  column "service_update_description" {
    rename = "description"
  }
  
  column "service_update_end_date" {
    rename = "end_date"
  }

  column "service_update_recommended_apply_by_date" {
    rename = "recommended_apply_by_date"
  }

  column "service_update_release_date" {
    rename = "release_date"
  }

  column "service_update_severity" {
    rename = "severity"
  }

  column "service_update_status" {
    rename = "status"
  }

  column "service_update_type" {
    rename = "type"
  }

}

resource "aws" "elasticache" "snapshots" {
  path = "github.com/aws/aws-sdk-go-v2/service/elasticache/types.Snapshot"

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
    primary_keys = [
      "arn"
    ]
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

  column "data_tiering" {
    description = "Data tiering"
  }

  column "node_group_configuration_node_group_id" {
    rename = "node_group_id"
  }

  column "node_group_configuration_primary_availability_zone" {
    rename = "node_group_primary_availability_zone"
  }

  column "node_group_configuration_primary_outpost_arn" {
    rename = "node_group_primary_outpost_arn"
  }

  column "node_group_configuration_replica_availability_zones" {
    rename =  "node_group_replica_availability_zones"
  }

  column "node_group_configuration_replica_count" {
    rename = "node_group_replica_count"
  }

  column "node_group_configuration_replica_outpost_arns" {
    rename = "node_group_replica_outpost_arns"
  }
}

resource "aws" "elasticache" "user_groups" {
  path = "github.com/aws/aws-sdk-go-v2/service/elasticache/types.UserGroup"

  description = "Describes Elasticache user groups"

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
    primary_keys = [
      "arn"
    ]
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

  column "pending_changes_user_ids_to_add" {
    rename = "pending_user_ids_to_add"
  }

  column "pending_changes_user_ids_to_remove" {
    rename = "pending_user_ids_to_remove"
  }
}


resource "aws" "elasticache" "users" {
  path = "github.com/aws/aws-sdk-go-v2/service/elasticache/types.User"

  description = "Describes Elasticache users"

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
    primary_keys = [
      "arn"
    ]
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
}