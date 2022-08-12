service          = "aws"
output_directory = "."
add_generate     = true

resource "aws" "wafregional" "rate_based_rules" {
  path        = "github.com/aws/aws-sdk-go-v2/service/wafregional/types.RateBasedRule"
  description = "A combination of identifiers for web requests that you want to allow, block, or count, including rate limit."
  multiplex "ServiceAccountRegionMultiplexer" {
<<<<<<< HEAD:plugins/source/aws/resources/services/wafregional/rate_based_rules.hcl
    path = "github.com/cloudquery/cloudquery/plugins/source/aws/client.ServiceAccountRegionMultiplexer"
  }
  ignoreError "IgnoreAccessDenied" {
    path = "github.com/cloudquery/cloudquery/plugins/source/aws/client.IgnoreAccessDeniedServiceDisabled"
=======
    path   = "github.com/cloudquery/cq-provider-aws/client.ServiceAccountRegionMultiplexer"
    params = ["waf-regional"]
  }
  ignoreError "IgnoreCommonErrors" {
    path = "github.com/cloudquery/cq-provider-aws/client.IgnoreCommonErrors"
>>>>>>> d1d8e7efe (Add ‘unused resource’ example policy (#1378)):resources/services/wafregional/rate_based_rules.hcl
  }
  deleteFilter "AccountRegionFilter" {
    path = "github.com/cloudquery/cloudquery/plugins/source/aws/client.DeleteAccountRegionFilter"
  }

  options {
    primary_keys = ["account_id", "region", "id"]
  }

  userDefinedColumn "account_id" {
    description = "The AWS Account ID of the resource."
    type        = "string"
    resolver "resolveAWSAccount" {
      path = "github.com/cloudquery/cloudquery/plugins/source/aws/client.ResolveAWSAccount"
    }
  }
  userDefinedColumn "region" {
    description = "The AWS Region of the resource."
    type        = "string"
    resolver "resolveAWSRegion" {
      path = "github.com/cloudquery/cloudquery/plugins/source/aws/client.ResolveAWSRegion"
    }
  }

  column "rule_id" {
    rename = "id"
  }

  userDefinedColumn "arn" {
    type              = "string"
    description       = "ARN of the rate based rule."
    generate_resolver = true
  }

  userDefinedColumn "tags" {
    type              = "json"
    generate_resolver = true
    description       = "Rule tags."
<<<<<<< HEAD:plugins/source/aws/resources/services/wafregional/rate_based_rules.hcl
=======
  }

  relation "aws" "wafregional" "match_predicates" {
    description = "Contains one Predicate element for each ByteMatchSet, IPSet, or SqlInjectionMatchSet object that you want to include in a RateBasedRule."
>>>>>>> d1d8e7efe (Add ‘unused resource’ example policy (#1378)):resources/services/wafregional/rate_based_rules.hcl
  }
}
