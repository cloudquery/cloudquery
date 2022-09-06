//check-for-changes
service          = "aws"
output_directory = "."
add_generate     = true

description_modifier "remove_read_only" {
  words = ["  This member is required."]
}

resource "aws" "acm" "certificates" {
  path = "github.com/aws/aws-sdk-go-v2/service/acm/types.CertificateDetail"
  ignoreError "IgnoreCommonErrors" {
    path = "github.com/cloudquery/cloudquery/plugins/source/aws/client.IgnoreCommonErrors"
  }
  deleteFilter "AccountRegionFilter" {
    path = "github.com/cloudquery/cloudquery/plugins/source/aws/client.DeleteAccountRegionFilter"
  }
  multiplex "AwsAccountRegion" {
    path   = "github.com/cloudquery/cloudquery/plugins/source/aws/client.ServiceAccountRegionMultiplexer"
    params = ["acm"]
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

  column "certificate_arn" {
    rename = "arn"
  }

  column "domain_validation_options" {
    type = "json"
  }

  column "extended_key_usages" {
    type = "json"
  }

  column "key_usages" {
    type = "stringArray"
    resolver "resolveKeyUsages" {
      path = "github.com/cloudquery/cq-provider-sdk/provider/schema.PathResolver"
      params = ["KeyUsages.Name"]
    }
  }

  column "renewal_summary_domain_validation_options" {
    type = "json"
  }

  column "renewal_summary_renewal_status" {
    rename = "renewal_summary_status"
  }

  column "renewal_summary_renewal_status_reason" {
    rename = "renewal_summary_failure_reason"
  }

  column "options" {
    skip_prefix = true
  }

  userDefinedColumn "tags" {
    description = "The tags that have been applied to the ACM certificate"
    type = "json"
    generate_resolver = true
  }
}
