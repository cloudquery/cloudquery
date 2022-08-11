service          = "aws"
output_directory = "."
add_generate     = true

#description_modifier "remove_read_only" {
#  words = ["  This member is required."]
#}


resource "aws" "mq" "brokers" {
  path = "github.com/aws/aws-sdk-go-v2/service/mq.DescribeBrokerOutput"
  ignoreError "IgnoreAccessDenied" {
    path = "github.com/cloudquery/cq-provider-aws/client.IgnoreCommonErrors"
  }
  multiplex "AwsAccountRegion" {
    path   = "github.com/cloudquery/cq-provider-aws/client.ServiceAccountRegionMultiplexer"
    params = ["mq"]
  }
  deleteFilter "AccountRegionFilter" {
    path = "github.com/cloudquery/cq-provider-aws/client.DeleteAccountRegionFilter"
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

  options {
    primary_keys = [
      "account_id", "id"
    ]
  }

  column "broker_arn" {
    rename = "arn"
  }


  column "auto_minor_version_upgrade" {
    description = "Enables automatic upgrades to new minor versions for brokers, as Apache releases the versions."
  }

  column "broker_instances" {
    type              = "json"
    generate_resolver = true
  }

  column "configurations" {
    skip = true
  }

  column "broker_id" {
    rename = "id"
  }

  column "deployment_mode" {
    description = "The deployment mode of the broker."
  }

  column "encryption_options_use_aws_owned_key" {
    description = "Enables the use of an AWS owned CMK using AWS Key Management Service (KMS)."
  }

  column "encryption_options_kms_key_id" {
    description = "The symmetric customer master key (CMK) to use for the AWS Key Management Service (KMS)."
  }

  column "engine_type" {
    description = "The type of broker engine."
  }

  column "ldap_server_metadata" {
    type              = "json"
    generate_resolver = true
  }

  column "logs" {
    type              = "json"
    generate_resolver = true
  }

  column "maintenance_window_start_time" {
    type              = "json"
    generate_resolver = true
  }

  column "pending_ldap_server_metadata" {
    type              = "json"
    generate_resolver = true
  }

  column "publicly_accessible" {
    description = "Enables connections from applications outside of the VPC that hosts the broker's subnets."
  }

  column "users" {
    skip = true
  }

  user_relation "aws" "mq" "configurations" {
    path = "github.com/aws/aws-sdk-go-v2/service/mq.DescribeConfigurationOutput"

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

    column "result_metadata_values" {
      skip = true
    }
    column "arn" {
      description = "The ARN of the configuration."
    }

    column "created" {
      description = "The date and time of the configuration revision."
    }

    column "description" {
      description = "The description of the configuration."
    }

    column "engine_type" {
      description = "The type of broker engine."
    }

    column "engine_version" {
      description = "The version of the broker engine."
    }

    column "authentication_strategy" {
      description = "The authentication strategy associated with the configuration."
    }

    column "id" {
      description = "The unique ID that Amazon MQ generates for the configuration."
      type        = "string"
    }

    column "latest_revision_created" {
      description = "The date and time of the configuration revision."
    }

    column "latest_revision_description" {
      description = "The description of the configuration revision."
    }

    column "latest_revision" {
      description = "The revision number of the configuration."
    }

    column "name" {
      description = "The name of the configuration."
    }

    user_relation "aws" "mq" "revisions" {
      path = "github.com/aws/aws-sdk-go-v2/service/mq.DescribeConfigurationRevisionOutput"

      userDefinedColumn "id" {

      }
      column "data" {
        type              = "json"
        generate_resolver = true
      }

      column "result_metadata_values" {
        skip = true
      }
    }
  }

  user_relation "aws" "mq" "users" {
    path = "github.com/aws/aws-sdk-go-v2/service/mq.DescribeUserOutput"
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

    column "broker_id" {
      skip = true
    }

    column "pending" {
      type              = "json"
      generate_resolver = true
    }
    column "result_metadata_values" {
      skip = true
    }
    column "username" {
      description = "The username of the ActiveMQ user."
    }
  }
}
