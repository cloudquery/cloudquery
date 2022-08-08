description_modifier "remove_read_only" {
  words = ["  This member is required."]
}

resource "aws" "kinesis" "streams" {
  path = "github.com/aws/aws-sdk-go-v2/service/kinesis/types.StreamDescriptionSummary"
  ignoreError "IgnoreAccessDenied" {
    path = "github.com/cloudquery/cq-provider-aws/client.IgnoreAccessDeniedServiceDisabled"
  }
  deleteFilter "AccountRegionFilter" {
    path = "github.com/cloudquery/cq-provider-aws/client.DeleteAccountRegionFilter"
  }
  multiplex "AwsAccountRegion" {
    path   = "github.com/cloudquery/cq-provider-aws/client.ServiceAccountRegionMultiplexer"
    params = ["kinesis"]
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
  ignore_columns_in_tests = ["kms_key_id", "retention_in_days"]

  options {
    primary_keys = ["arn"]
  }

  userDefinedColumn "arn" {
    type = "string"
    resolver "resolveStreamArn" {
      path          = "github.com/cloudquery/cq-provider-sdk/provider/schema.PathResolver"
      path_resolver = true
      // TODO: require manual changing from ARN -> StreamARN for the path resolver as its not supported by cq-gen yet
      params = ["stream_arn"]
    }
  }

  userDefinedColumn "tags" {
    type              = "json"
    generate_resolver = true
  }
}
