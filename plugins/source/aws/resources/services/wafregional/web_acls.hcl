service = "aws"

output_directory = "."

resource "aws" "wafregional" "web_acls" {
  path = "github.com/aws/aws-sdk-go-v2/service/wafregional/types.WebACL"
  multiplex "ServiceAccountRegionMultiplexer" {
    path = "github.com/cloudquery/cq-provider-aws/client.ServiceAccountRegionMultiplexer"
  }
  ignoreError "IgnoreAccessDenied" {
    path = "github.com/cloudquery/cq-provider-aws/client.IgnoreAccessDeniedServiceDisabled"
  }
  deleteFilter "AccountRegionFilter" {
    path = "github.com/cloudquery/cq-provider-aws/client.DeleteAccountRegionFilter"
  }

  options {
    primary_keys = ["account_id", "region", "id"]
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

  column "web_acl_id" {
    rename = "id"
  }

  column "web_acl_arn" {
    rename = "arn"
  }

  column "default_action_type" {
    rename = "default_action"
  }

  userDefinedColumn "tags" {
    type = "json"
    generate_resolver = true
    description = "Web ACL tags."
  }

  relation "aws" "wafregional" "rules" {
    path = "github.com/aws/aws-sdk-go-v2/service/wafregional/types.ActivatedRule"
    description = "The action for each Rule in a WebACL"

    column "action_type" {
      rename = "action"
    }

    column "override_action_type" {
      rename = "override_action"
      description = "Describes an override action for the rule."
    }

    column "excluded_rules" {
      type = "stringarray"
      generate_resolver = true
    }
  }
}
