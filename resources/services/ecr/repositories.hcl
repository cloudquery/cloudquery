service = "aws"

output_directory = "."

resource "aws" "ecr" "repositories" {
  multiplex "AwsAccountRegion" {
    path   = "github.com/cloudquery/cq-provider-aws/client.ServiceAccountRegionMultiplexer"
    params = ["api.ecr"]
  }


  path = "github.com/aws/aws-sdk-go-v2/service/ecr/types.Repository"
  ignoreError "IgnoreCommonErrors" {
    path = "github.com/cloudquery/cq-provider-aws/client.IgnoreCommonErrors"
  }
  deleteFilter "AccountRegionFilter" {
    path = "github.com/cloudquery/cq-provider-aws/client.DeleteAccountRegionFilter"
  }

  options {
    primary_keys = ["account_id", "arn"]
  }

  column "repository_arn" {
    rename = "arn"
  }
  column "repository_name" {
    rename = "name"
  }

  column "repository_uri" {
    rename = "uri"
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


  user_relation "aws" "ecr" "images" {
    path = "github.com/aws/aws-sdk-go-v2/service/ecr/types.ImageDetail"
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
    // user_relation "aws" "ecr" "image_scan_findings" {
    //   path = "github.com/aws/aws-sdk-go-v2/service/ecr/types.ImageScanFindings"
    // }
  }
}