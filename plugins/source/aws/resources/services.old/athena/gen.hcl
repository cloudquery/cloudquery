service          = "aws"
output_directory = "."
add_generate     = true

description_modifier "remove_read_only" {
  words = ["  This member is required."]
}

resource "aws" "athena" "data_catalogs" {
  path = "github.com/aws/aws-sdk-go-v2/service/athena/types.DataCatalog"

  ignoreError "IgnoreCommonErrors" {
    path = "github.com/cloudquery/cloudquery/plugins/source/aws/client.IgnoreCommonErrors"
  }
  multiplex "AwsAccountRegion" {
    path   = "github.com/cloudquery/cloudquery/plugins/source/aws/client.ServiceAccountRegionMultiplexer"
    params = ["athena"]
  }
  deleteFilter "AccountRegionFilter" {
    path = "github.com/cloudquery/cloudquery/plugins/source/aws/client.DeleteAccountRegionFilter"
  }

  options {
    primary_keys = [
      "arn",
    ]
  }

  ignore_columns_in_tests = ["description", "parameters"]

  userDefinedColumn "account_id" {
    type        = "string"
    description = "The AWS Account ID of the resource."
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
    description       = "ARN of the resource."
    generate_resolver = true
  }

  user_relation "aws" "athena" "databases" {
    path                    = "github.com/aws/aws-sdk-go-v2/service/athena/types.Database"
    ignore_columns_in_tests = ["description", "parameters"]

    user_relation "aws" "athena" "tables" {
      path                    = "github.com/aws/aws-sdk-go-v2/service/athena/types.TableMetadata"
      ignore_columns_in_tests = ["last_access_time", "table_type"]

      relation "aws" "athena_data_catalog_database_table" "columns" {
        path            = "github.com/aws/aws-sdk-go-v2/service/athena/types.Column"
        ignore_in_tests = true
      }

      relation "aws" "athena_data_catalog_database_table" "partition_keys" {
        path            = "github.com/aws/aws-sdk-go-v2/service/athena/types.Column"
        ignore_in_tests = true
      }
    }
  }

  userDefinedColumn "tags" {
    type              = "json"
    description       = "Tags associated with the Athena data catalog."
    generate_resolver = true
  }
}


resource "aws" "athena" "work_groups" {
  path = "github.com/aws/aws-sdk-go-v2/service/athena/types.WorkGroup"

  ignoreError "IgnoreCommonErrors" {
    path = "github.com/cloudquery/cloudquery/plugins/source/aws/client.IgnoreCommonErrors"
  }
  multiplex "AwsAccountRegion" {
    path   = "github.com/cloudquery/cloudquery/plugins/source/aws/client.ServiceAccountRegionMultiplexer"
    params = ["athena"]
  }
  deleteFilter "AccountRegionFilter" {
    path = "github.com/cloudquery/cloudquery/plugins/source/aws/client.DeleteAccountRegionFilter"
  }

  options {
    primary_keys = [
      "arn"
    ]
  }

  ignore_columns_in_tests = [
    "bytes_scanned_cutoff_per_query",
    "encryption_configuration_kms_key",
    "expected_bucket_owner",
    "output_location"
  ]
  userDefinedColumn "account_id" {
    type        = "string"
    description = "The AWS Account ID of the resource."
    resolver "resolveAWSAccount" {
      path = "github.com/cloudquery/cloudquery/plugins/source/aws/client.ResolveAWSAccount"
    }
  }
  userDefinedColumn "arn" {
    type              = "string"
    description       = "ARN of the resource."
    generate_resolver = true
  }

  userDefinedColumn "region" {
    type        = "string"
    description = "The AWS Region of the resource."
    resolver "resolveAWSRegion" {
      path = "github.com/cloudquery/cloudquery/plugins/source/aws/client.ResolveAWSRegion"
    }
  }

  column "configuration" {
    skip_prefix = true
  }

  column "result_configuration" {
    skip_prefix = true
  }

  column "engine_version" {
    skip_prefix = true
  }


  userDefinedColumn "tags" {
    type              = "json"
    description       = "Tags associated with the Athena work group."
    generate_resolver = true
  }
  column "effective_engine_version" {
    description = "The engine version on which the query runs. If the user requests a valid engine version other than Auto, the effective engine version is the same as the engine version that the user requested."
  }

  user_relation "aws" "athena" "prepared_statements" {
    path            = "github.com/aws/aws-sdk-go-v2/service/athena/types.PreparedStatement"
    ignore_in_tests = true
  }
  user_relation "aws" "athena" "query_executions" {
    path            = "github.com/aws/aws-sdk-go-v2/service/athena/types.QueryExecution"
    ignore_in_tests = true

    column "engine_version" {
      skip_prefix = true
    }
    column "query_execution_context" {
      skip_prefix = true
    }

    column "query_execution_id" {
      rename = "id"
    }

    column "result_configuration" {
      skip_prefix = true
    }

    column "statistics" {
      skip_prefix = true
    }

    column "effective_engine_version" {
      description = "The engine version on which the query runs. If the user requests a valid engine version other than Auto, the effective engine version is the same as the engine version that the user requested."
    }


    column "status" {
      skip_prefix = true
    }
  }
  user_relation "aws" "athena" "named_queries" {
    path = "github.com/aws/aws-sdk-go-v2/service/athena/types.NamedQuery"
  }
}