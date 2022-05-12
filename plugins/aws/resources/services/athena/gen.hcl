service          = "aws"
output_directory = "."
add_generate     = true

description_modifier "remove_read_only" {
  words = ["  This member is required."]
}

resource "aws" "athena" "data_catalogs" {
  path = "github.com/aws/aws-sdk-go-v2/service/athena/types.DataCatalog"

  ignoreError "IgnoreAccessDenied" {
    path = "github.com/cloudquery/cq-provider-aws/client.IgnoreAccessDeniedServiceDisabled"
  }
  multiplex "AwsAccountRegion" {
    path   = "github.com/cloudquery/cq-provider-aws/client.ServiceAccountRegionMultiplexer"
    params = ["athena"]
  }
  deleteFilter "AccountRegionFilter" {
    path = "github.com/cloudquery/cq-provider-aws/client.DeleteAccountRegionFilter"
  }

  options {
    primary_keys = [
      "arn",
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

  userDefinedColumn "arn" {
    type              = "string"
    description       = "ARN of the resource."
    generate_resolver = true
  }

  user_relation "aws" "athena" "databases" {
    path = "github.com/aws/aws-sdk-go-v2/service/athena/types.Database"
    user_relation "aws" "athena" "tables" {
      path = "github.com/aws/aws-sdk-go-v2/service/athena/types.TableMetadata"
    }
  }

  userDefinedColumn "tags" {
    type              = "json"
    generate_resolver = true
  }
}


resource "aws" "athena" "work_groups" {
  path = "github.com/aws/aws-sdk-go-v2/service/athena/types.WorkGroup"

  ignoreError "IgnoreAccessDenied" {
    path = "github.com/cloudquery/cq-provider-aws/client.IgnoreAccessDeniedServiceDisabled"
  }
  multiplex "AwsAccountRegion" {
    path   = "github.com/cloudquery/cq-provider-aws/client.ServiceAccountRegionMultiplexer"
    params = ["athena"]
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
  userDefinedColumn "arn" {
    type              = "string"
    description       = "ARN of the resource."
    generate_resolver = true
  }

  userDefinedColumn "region" {
    type        = "string"
    description = "The AWS Region of the resource."
    resolver "resolveAWSRegion" {
      path = "github.com/cloudquery/cq-provider-aws/client.ResolveAWSRegion"
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
    generate_resolver = true
  }
  column "effective_engine_version" {
    description = "The engine version on which the query runs. If the user requests a valid engine version other than Auto, the effective engine version is the same as the engine version that the user requested."
  }

  user_relation "aws" "athena" "prepared_statements" {
    path = "github.com/aws/aws-sdk-go-v2/service/athena/types.PreparedStatement"
  }
  user_relation "aws" "athena" "query_executions" {
    path = "github.com/aws/aws-sdk-go-v2/service/athena/types.QueryExecution"
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