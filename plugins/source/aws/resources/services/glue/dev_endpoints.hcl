service          = "aws"
output_directory = "."
add_generate     = true


description_modifier "remove_read_only" {
  words = ["  This member is required."]
}

resource "aws" "glue" "dev_endpoints" {
  path = "github.com/aws/aws-sdk-go-v2/service/glue/types.DevEndpoint"
  ignoreError "IgnoreAccessDenied" {
    path = "github.com/cloudquery/cq-provider-aws/client.IgnoreAccessDeniedServiceDisabled"
  }
  deleteFilter "AccountRegionFilter" {
    path = "github.com/cloudquery/cq-provider-aws/client.DeleteAccountRegionFilter"
  }
  multiplex "AwsAccountRegion" {
    path   = "github.com/cloudquery/cq-provider-aws/client.ServiceAccountRegionMultiplexer"
    params = ["glue"]
  }
  options {
    primary_keys = ["arn"]
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
    type              = "string"
    description       = "The Amazon Resource Name (ARN) of the workflow."
    generate_resolver = true
  }

  column "endpoint_name" {
    rename = "name"
  }

  column "worker_type" {
    description = "The type of predefined worker that is allocated to the development endpoint Accepts a value of Standard, G1X, or G2X"
  }

  userDefinedColumn "tags" {
    type = "json"
    description = "Resource tags."
    generate_resolver = true
  }
}
