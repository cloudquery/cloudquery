//check-for-changes
service          = "aws"
output_directory = "."
add_generate     = true

description_modifier "remove_read_only" {
  words = ["  This member is required."]
}

resource "aws" "inspector" "findings" {
  path = "github.com/aws/aws-sdk-go-v2/service/inspector/types.Finding"
  ignoreError "IgnoreCommonErrors" {
    path = "github.com/cloudquery/cloudquery/plugins/source/aws/client.IgnoreCommonErrors"
  }
  multiplex "AwsAccountRegion" {
    path   = "github.com/cloudquery/cloudquery/plugins/source/aws/client.ServiceAccountRegionMultiplexer"
    params = ["inspector"]
  }
  deleteFilter "AccountRegionFilter" {
    path = "github.com/cloudquery/cloudquery/plugins/source/aws/client.DeleteAccountRegionFilter"
  }

  options {
    primary_keys = [
      "arn"
    ]
  }

  userDefinedColumn "account_id" {
    type        = "string"
    description = "The AWS Account ID of the resource"
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

  column "asset_attributes" {
    type = "json"
  }

  column "attributes" {
    type = "json"
    resolver "resolveAttributes" {
      path   = "github.com/cloudquery/cloudquery/plugins/source/aws/client.ResolveTagField"
      params = ["Attributes"]
    }
  }

  column "user_attributes" {
    type = "json"
    resolver "resolveAttributes" {
      path   = "github.com/cloudquery/cloudquery/plugins/source/aws/client.ResolveTagField"
      params = ["UserAttributes"]
    }
  }

  column "service_attributes" {
    type = "json"
  }
}