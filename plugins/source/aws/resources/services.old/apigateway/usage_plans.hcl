//check-for-changes
service          = "aws"
output_directory = "."
add_generate     = true

description_modifier "remove_read_only" {
  words = ["  This member is required."]
}

resource "aws" "apigateway" "usage_plans" {
  path = "github.com/aws/aws-sdk-go-v2/service/apigateway/types.UsagePlan"
  ignoreError "IgnoreCommonErrors" {
    path = "github.com/cloudquery/cloudquery/plugins/source/aws/client.IgnoreCommonErrors"
  }
  deleteFilter "AccountRegionFilter" {
    path = "github.com/cloudquery/cloudquery/plugins/source/aws/client.DeleteAccountRegionFilter"
  }
  multiplex "AwsAccountRegion" {
    path   = "github.com/cloudquery/cloudquery/plugins/source/aws/client.ServiceAccountRegionMultiplexer"
    params = ["apigateway"]
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

  userDefinedColumn "arn" {
    description = "The Amazon Resource Name (ARN) for the resource"
    type = "string"
    generate_resolver = true
  }

  relation "aws" "apigateway" "api_stages" {
    userDefinedColumn "usage_plan_id" {
      description = "The identifier of a UsagePlan resource"
      type = "string"
      resolver "resolveUsagePlanId" {
        path = "github.com/cloudquery/cq-provider-sdk/provider/schema.ParentPathResolver"
        params = ["Id"]
      }
    }
  }

  user_relation "aws" "apigateway" "keys" {
    path = "github.com/aws/aws-sdk-go-v2/service/apigateway/types.UsagePlanKey"
    userDefinedColumn "usage_plan_id" {
      description = "The identifier of a UsagePlan resource"
      type = "string"
      resolver "resolveUsagePlanId" {
        path = "github.com/cloudquery/cq-provider-sdk/provider/schema.ParentPathResolver"
        params = ["Id"]
      }
    }
    userDefinedColumn "arn" {
      description = "The Amazon Resource Name (ARN) for the resource"
      type = "string"
      generate_resolver = true
    }
  }
}
