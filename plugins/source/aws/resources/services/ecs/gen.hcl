
//check-for-changes
service          = "aws"
output_directory = "."
add_generate     = true

description_modifier "remove_read_only" {
  words = ["  This member is required."]
}


resource "aws" "ecs" "clusters" {
  path = "github.com/aws/aws-sdk-go-v2/service/ecs/types.Cluster"
  ignoreError "IgnoreAccessDenied" {
    path = "github.com/cloudquery/cq-provider-aws/client.IgnoreAccessDeniedServiceDisabled"
  }
  multiplex "AwsAccountRegion" {
    path = "github.com/cloudquery/cq-provider-aws/client.AccountRegionMultiplex"
  }
  deleteFilter "AccountRegionFilter" {
    path = "github.com/cloudquery/cq-provider-aws/client.DeleteAccountRegionFilter"
  }
  userDefinedColumn "account_id" {
    type        = "string"
    description = "The AWS Account ID of the resource."
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

  options {
    primary_keys = [
      "arn"
    ]
  }

  column "configuration" {
    skip_prefix = true
  }
  column "cluster_arn" {
    rename = "arn"
  }
  column "cluster_name" {
    rename = "name"
  }
  column "execute_command_configuration_kms_key_id" {
    rename = "execute_config_kms_key_id"
  }
  column "execute_command_configuration_log_configuration_cloud_watch_encryption_enabled" {
    rename = "execute_config_logs_cloud_watch_encryption_enabled"
  }
  column "execute_command_configuration_log_configuration_cloud_watch_log_group_name" {
    rename = "execute_config_log_cloud_watch_log_group_name"
  }
  column "execute_command_configuration_log_configuration_s3_bucket_name" {
    rename = "execute_config_log_s3_bucket_name"
  }
  column "execute_command_configuration_log_configuration_s3_encryption_enabled" {
    rename = "execute_config_log_s3_encryption_enabled"
  }
  column "execute_command_configuration_log_configuration_s3_key_prefix" {
    rename = "execute_config_log_s3_key_prefix"
  }
  column "execute_command_configuration_logging" {
    rename = "execute_config_logging"
  }
  column "settings" {
    type              = "json"
    generate_resolver = true
  }
  column "statistics" {
    type              = "json"
    generate_resolver = true
  }
  column "tags" {
    type              = "json"
    generate_resolver = true
  }
  column "default_capacity_provider_strategy" {
    type              = "json"
    generate_resolver = true
  }

  relation "aws" "ecs" "attachments" {
    path = "github.com/aws/aws-sdk-go-v2/service/ecs/types.Attachment"

    options {
      primary_keys = [
        "cluster_cq_id",
        "id"
      ]
    }
    column "details" {
      type              = "json"
      generate_resolver = true
    }

  }

  relation "aws" "ecs" "services" {
    path = "github.com/aws/aws-sdk-go-v2/service/ecs/types.Service"

    column "capacity_provider_strategy" {
      type              = "json"
      generate_resolver = true
    }

    column "enable_e_c_s_managed_tags" {
      rename = "enable_ecs_managed_tags"
    }

    column "service_arn" {
      rename = "arn"
    }

    column "service_name" {
      rename = "name"
    }
    column "placement_constraints" {
      type              = "json"
      generate_resolver = true
    }

    column "placement_strategy" {
      type              = "json"
      generate_resolver = true
    }

    column "tags" {
      type              = "json"
      generate_resolver = true
    }

    relation "aws" "ecs" "deployments" {
      path = "github.com/aws/aws-sdk-go-v2/service/ecs/types.Deployment"

      column "capacity_provider_strategy" {
        type              = "json"
        generate_resolver = true
      }
    }

    relation "aws" "ecs" "task_sets" {
      path = "github.com/aws/aws-sdk-go-v2/service/ecs/types.TaskSet"
      column "task_set_arn" {
        rename = "arn"
      }
      column "capacity_provider_strategy" {
        type              = "json"
        generate_resolver = true
      }

      column "tags" {
        type              = "json"
        generate_resolver = true
      }
      relation "aws" "ecs" "service_registries" {
        path = "github.com/aws/aws-sdk-go-v2/service/ecs/types.ServiceRegistry"

        column "registry_arn" {
          rename = "arn"
        }
      }
    }

    relation "aws" "ecs" "registries" {
      path = "github.com/aws/aws-sdk-go-v2/service/ecs/types.Registry"

      column "registry_arn" {
        rename = "arn"
      }
    }

  }

  relation "aws" "ecs" "container_instances" {
    path = "github.com/aws/aws-sdk-go-v2/service/ecs/types.ContainerInstance"

    relation "aws" "ecs" "attachments" {
      path = "github.com/aws/aws-sdk-go-v2/service/ecs/types.Attachment"

      column "details" {
        type              = "json"
        generate_resolver = true
      }

    }
    column "tags" {
      type              = "json"
      generate_resolver = true
    }

  }


}