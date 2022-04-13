service          = "aws"
output_directory = "."
add_generate     = true

resource "aws" "ec2" "egress_only_internet_gateways" {
  path = "github.com/aws/aws-sdk-go-v2/service/ec2/types.EgressOnlyInternetGateway"
  ignoreError "IgnoreAccessDenied" {
    path = "github.com/cloudquery/cq-provider-aws/client.IgnoreAccessDeniedServiceDisabled"
  }
  deleteFilter "AccountRegionFilter" {
    path = "github.com/cloudquery/cq-provider-aws/client.DeleteAccountRegionFilter"
  }
  multiplex "AwsAccountRegion" {
    path   = "github.com/cloudquery/cq-provider-aws/client.ServiceAccountRegionMultiplexer"
    params = ["ec2"]
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

  column "egress_only_internet_gateway_id" {
    rename = "id"
  }

  column "attachments" {
    type              = "json"
    generate_resolver = true
    description       = "Information about the attachment of the egress-only internet gateway."
  }

  column "tags" {
    type              = "json"
    generate_resolver = false
    description       = "The tags assigned to the egress-only internet gateway."
  }

  userDefinedColumn "arn" {
    type        = "string"
    description = "The Amazon Resource Name (ARN) for the egress-only internet gateway."
    generate_resolver = false
  }

}