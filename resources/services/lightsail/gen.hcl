service          = "aws"
output_directory = "."
add_generate     = true

resource "aws" "lightsail" "instances" {
  path = "github.com/aws/aws-sdk-go-v2/service/lightsail/types.Instance"
  ignoreError "IgnoreAccessDenied" {
    path = "github.com/cloudquery/cq-provider-aws/client.IgnoreAccessDeniedServiceDisabled"
  }
  deleteFilter "AccountRegionFilter" {
    path = "github.com/cloudquery/cq-provider-aws/client.DeleteAccountRegionFilter"
  }
  multiplex "AwsAccountRegion" {
    path   = "github.com/cloudquery/cq-provider-aws/client.ServiceAccountRegionMultiplexer"
    params = ["lightsail"]
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

  options {
    primary_keys = ["arn"]
  }

  column "tags" {
    type              = "json"
    generate_resolver = true
  }

  relation "aws" "lightsail" "add_ons" {
    ignore_in_tests = true // see https://github.com/hashicorp/terraform-provider-aws/issues/23688
  }

  // disks attached to the instance
  relation "aws" "lightsail" "hardware_disks" {
    path = "github.com/aws/aws-sdk-go-v2/service/lightsail/types.Disk"

    column "tags" {
      type              = "json"
      generate_resolver = true
    }

    // skip columns deprecated by the SDK
    column "gb_in_use" {
      skip = true
    }

    column "attachment_state" {
      skip = true
    }

    relation "aws" "lightsail" "add_ons" {
      ignore_in_tests = true // see https://github.com/hashicorp/terraform-provider-aws/issues/23688
    }
  }
}
