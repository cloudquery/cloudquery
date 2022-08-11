service          = "aws"
output_directory = "."
add_generate     = true

description_modifier "remove_read_only" {
  words = ["  This member is required."]
}

resource "aws" "lambda" "functions" {
  path        = "github.com/aws/aws-sdk-go-v2/service/lambda.GetFunctionOutput"
  description = "AWS Lambda is a serverless compute service that lets you run code without provisioning or managing servers, creating workload-aware cluster scaling logic, maintaining event integrations, or managing runtimes"
  ignoreError "IgnoreAccessDenied" {
    path = "github.com/cloudquery/cq-provider-aws/client.IgnoreAccessDeniedServiceDisabled"
  }
  multiplex "AwsAccountRegion" {
    path   = "github.com/cloudquery/cq-provider-aws/client.ServiceAccountRegionMultiplexer"
    params = ["lambda"]
  }
  deleteFilter "AccountRegionFilter" {
    path = "github.com/cloudquery/cq-provider-aws/client.DeleteAccountRegionFilter"
  }

  options {
    primary_keys = [
      "arn"
    ]
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

  postResourceResolver "resolvePolicyCodeSigningConfig" {
    path     = "github.com/cloudquery/cq-provider-sdk/provider/schema.RowResolver"
    generate = true
  }

  userDefinedColumn "policy_document" {
    type        = "json"
    description = "The resource-based policy."
  }
  userDefinedColumn "policy_revision_id" {
    type        = "string"
    description = "A unique identifier for the current revision of the policy."
  }

  userDefinedColumn "code_signing_allowed_publishers_version_arns" {
    type        = "stringarray"
    description = "The Amazon Resource Name (ARN) for each of the signing profiles. A signing profile defines a trusted user who can sign a code package."
  }
  userDefinedColumn "code_signing_config_arn" {
    description = "The Amazon Resource Name (ARN) of the Code signing configuration."
    type        = "string"
  }
  userDefinedColumn "code_signing_config_id" {
    description = "Unique identifier for the Code signing configuration."
    type        = "string"
  }
  userDefinedColumn "code_signing_policies_untrusted_artifact_on_deployment" {
    description = "Code signing configuration policy for deployment validation failure."
    type        = "string"
  }
  userDefinedColumn "code_signing_description" {
    description = "Code signing configuration description."
    type        = "string"
  }
  userDefinedColumn "code_signing_last_modified" {
    description = "The date and time that the Code signing configuration was last modified, in ISO-8601 format (YYYY-MM-DDThh:mm:ss.sTZD)."

    type = "timestamp"
  }

  column "function_arn" {
    rename = "arn"
  }

  column "function_name" {
    rename = "name"
  }
  column "configuration" {
    skip_prefix = true
  }


  column "result_metadata_values" {
    skip = true
  }

  column "get_function_output" {
    skip_prefix = true
  }

  column "tags" {
    generate_resolver = true
  }

  column "image_config_response" {
    skip_prefix = true
  }

  column "environment_error_error_code" {
    rename = "environment_error_code"
  }

  user_relation "aws" "lambda" "event_invoke_configs" {
    path        = "github.com/aws/aws-sdk-go-v2/service/lambda/types.FunctionEventInvokeConfig"
    description = "A configuration object that specifies the destination of an event after Lambda processes it. "
    column "destination_config" {
      skip_prefix = true
    }
  }


  user_relation "aws" "lambda" "aliases" {
    path = "github.com/cloudquery/cq-provider-aws/resources/services/lambda.AliasWrapper"

    options {
      primary_keys = [
        "function_cq_id",
        "arn"
      ]
    }
    column "url_config_result_metadata_values" {
      skip = true
    }

    column "url_config_last_modified_time" {
      type = "timestamp"
      resolver "dateResolver" {
        path   = "github.com/cloudquery/cq-provider-sdk/provider/schema.DateResolver"
        params = ["UrlConfig.LastModifiedTime"]
      }
    }

    column "url_config_creation_time" {
      type = "timestamp"
      resolver "dateResolver" {
        path   = "github.com/cloudquery/cq-provider-sdk/provider/schema.DateResolver"
        params = ["UrlConfig.CreationTime"]
      }
    }

    column "url_config_cors" {
      type              = "json"
      generate_resolver = true
    }
    column "alias_configuration" {
      skip_prefix = true
    }
    column "alias_arn" {
      rename = "arn"
    }
    userDefinedColumn "function_arn" {
      type        = "string"
      description = "The Amazon Resource Name (ARN) of the lambda function"
      resolver "resolveArn" {
        //argument arn
        path   = "github.com/cloudquery/cq-provider-sdk/provider/schema.ParentResourceFieldResolver"
        params = ["arn"]
      }
    }
  }

  user_relation "aws" "lambda" "versions" {
    path = "github.com/aws/aws-sdk-go-v2/service/lambda/types.FunctionConfiguration"

    options {
      primary_keys = [
        "function_cq_id",
        "version"
      ]
    }

    column "image_config_response" {
      skip_prefix = true
    }
    relation "aws" "lambda" "layers" {
      options {
        primary_keys = [
          "function_version_cq_id", "arn"
        ]
      }
    }
  }

  relation "aws" "lambda" "layers" {
    options {
      primary_keys = [
        "function_cq_id",
        "arn"
      ]
    }
    userDefinedColumn "function_arn" {
      type        = "string"
      description = "The Amazon Resource Name (ARN) of the lambda function"
      resolver "resolveArn" {
        //argument arn
        path   = "github.com/cloudquery/cq-provider-sdk/provider/schema.ParentResourceFieldResolver"
        params = ["arn"]
      }
    }
    relation "aws" "lambda" "file_system_configs" {
      options {
        primary_keys = [
          "function_version_cq_id",
          "arn"
        ]
      }
    }
  }

  relation "aws" "lambda" "file_system_configs" {
    description = "Details about the connection between a Lambda function and an Amazon EFS file system. "
    options {
      primary_keys = [
        "function_cq_id",
        "arn"
      ]
    }

    userDefinedColumn "function_arn" {
      type        = "string"
      description = "The Amazon Resource Name (ARN) of the lambda function"
      resolver "resolveArn" {
        //argument arn
        path   = "github.com/cloudquery/cq-provider-sdk/provider/schema.ParentResourceFieldResolver"
        params = ["arn"]
      }
    }
  }

  user_relation "aws" "lambda" "function_concurrency_configs" {
    path = "github.com/aws/aws-sdk-go-v2/service/lambda/types.ProvisionedConcurrencyConfigListItem"
    column "image_config_response" {
      skip_prefix = true
    }
  }

  user_relation "aws" "lambda" "function_event_source_mappings" {
    path = "github.com/aws/aws-sdk-go-v2/service/lambda/types.EventSourceMappingConfiguration"
    column "image_config_response" {
      skip_prefix = true
    }

    options {
      primary_keys = [
        "function_cq_id", "uuid"
      ]
    }
    column "source_access_configurations" {
      type              = "json"
      generate_resolver = true
    }

    column "filter_criteria_filters" {
      rename            = "criteria_filters"
      type              = "stringarray"
      generate_resolver = true
    }

    column "destination_config" {
      skip_prefix = true
    }
  }
}