service          = "aws"
output_directory = "."
add_generate     = true

description_modifier "remove_read_only" {
  words = ["  This member is required."]
}

resource "aws" "shield" "protections" {
  path = "github.com/aws/aws-sdk-go-v2/service/shield/types.Protection"
  ignoreError "IgnoreAccessDenied" {
    path = "github.com/cloudquery/cq-provider-aws/client.IgnoreAccessDeniedServiceDisabled"
  }
  multiplex "AwsAccount" {
    path = "github.com/cloudquery/cq-provider-aws/client.AccountMultiplex"
  }
  deleteFilter "AccountRegionFilter" {
    path = "github.com/cloudquery/cq-provider-aws/client.DeleteAccountFilter"
  }
  userDefinedColumn "account_id" {
    type        = "string"
    description = "The AWS Account ID of the resource."
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

  column "protection_arn" {
    rename = "arn"
  }

  column "application_layer_automatic_response_configuration_status" {
    rename = "application_automatic_response_configuration_status"
  }

  userDefinedColumn "tags" {
    type              = "json"
    description       = "The AWS tags of the resource."
    generate_resolver = true
  }
}


resource "aws" "shield" "subscriptions" {
  path = "github.com/aws/aws-sdk-go-v2/service/shield/types.Subscription"
  ignoreError "IgnoreAccessDenied" {
    path = "github.com/cloudquery/cq-provider-aws/client.IgnoreAccessDeniedServiceDisabled"
  }
  multiplex "AwsAccount" {
    path = "github.com/cloudquery/cq-provider-aws/client.AccountMultiplex"
  }
  deleteFilter "AccountRegionFilter" {
    path = "github.com/cloudquery/cq-provider-aws/client.DeleteAccountFilter"
  }
  userDefinedColumn "account_id" {
    type        = "string"
    description = "The AWS Account ID of the resource."
    resolver "resolveAWSAccount" {
      path = "github.com/cloudquery/cq-provider-aws/client.ResolveAWSAccount"
    }
  }

  options {
    primary_keys = ["arn"]
  }

  column "subscription_limits" {
    skip_prefix = true
  }

  column "subscription_arn" {
    rename = "arn"
  }
  column "limits" {
    type              = "json"
    generate_resolver = true
  }
  column "protection_limits_protected_resource_type_limits" {
    rename            = "protected_resource_type_limits"
    type              = "json"
    generate_resolver = true
  }
  column "protection_group_limits_pattern_type_limits_arbitrary_pattern_limits" {
    rename = "protection_group_limits_arbitrary_pattern_limits"
  }

  column "protection_group_limits_pattern_type_limitsprotection_group_limits_arbitrary_pattern_limits_max_members" {
    rename            = "protection_group_limits_arbitrary_pattern_limits_max_members"
    type              = "int"
    generate_resolver = true
  }

  column "protection_group_limits_max_protection_groups" {
    type              = "int"
    generate_resolver = true
  }


  column "time_commitment_in_seconds" {
    type              = "int"
    generate_resolver = true
  }
}


resource "aws" "shield" "attacks" {
  path = "github.com/aws/aws-sdk-go-v2/service/shield/types.AttackDetail"
  ignoreError "IgnoreAccessDenied" {
    path = "github.com/cloudquery/cq-provider-aws/client.IgnoreAccessDeniedServiceDisabled"
  }
  multiplex "AwsAccount" {
    path = "github.com/cloudquery/cq-provider-aws/client.AccountMultiplex"
  }
  deleteFilter "AccountRegionFilter" {
    path = "github.com/cloudquery/cq-provider-aws/client.DeleteAccountFilter"
  }
  userDefinedColumn "account_id" {
    type        = "string"
    description = "The AWS Account ID of the resource."
    resolver "resolveAWSAccount" {
      path = "github.com/cloudquery/cq-provider-aws/client.ResolveAWSAccount"
    }
  }
  options {
    primary_keys = ["id"]
  }
  column "attack_id" {
    rename = "id"
  }
  column "attack_property" {
    skip_prefix = true
  }
  column "mitigations" {
    type              = "stringarray"
    generate_resolver = true
  }

  relation "aws" "shield" "attack_properties" {
    column "top_contributors" {
      type              = "json"
      generate_resolver = true
    }
  }

  column "attack_counters" {
    type              = "json"
    generate_resolver = true
  }

  relation "aws" "shield" "sub_resources" {
    column "attack_vectors" {
      type              = "json"
      generate_resolver = true
    }

    column "counters" {
      type              = "json"
      generate_resolver = true
    }
  }
}


resource "aws" "shield" "protection_groups" {
  path = "github.com/aws/aws-sdk-go-v2/service/shield/types.ProtectionGroup"
  ignoreError "IgnoreAccessDenied" {
    path = "github.com/cloudquery/cq-provider-aws/client.IgnoreAccessDeniedServiceDisabled"
  }
  multiplex "AwsAccount" {
    path = "github.com/cloudquery/cq-provider-aws/client.AccountMultiplex"
  }
  deleteFilter "AccountRegionFilter" {
    path = "github.com/cloudquery/cq-provider-aws/client.DeleteAccountFilter"
  }
  userDefinedColumn "account_id" {
    type        = "string"
    description = "The AWS Account ID of the resource."
    resolver "resolveAWSAccount" {
      path = "github.com/cloudquery/cq-provider-aws/client.ResolveAWSAccount"
    }
  }

  options {
    primary_keys = ["arn"]
  }

  column "protection_group_id" {
    rename = "id"
  }

  column "protection_group_arn" {
    rename = "arn"
  }

  userDefinedColumn "tags" {
    type              = "json"
    generate_resolver = true
  }
}