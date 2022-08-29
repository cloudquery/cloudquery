//check-for-changes
service          = "aws"
output_directory = "."
add_generate     = true

description_modifier "remove_read_only" {
  words = ["  This member is required."]
}

resource "aws" "iam" "roles" {
  path        = "github.com/aws/aws-sdk-go-v2/service/iam/types.Role"
  description = "An IAM role is an IAM identity that you can create in your account that has specific permissions."
  ignoreError "IgnoreCommonErrors" {
    path = "github.com/cloudquery/cloudquery/plugins/source/aws/client.IgnoreCommonErrors"
  }
  multiplex "AwsAccount" {
    path = "github.com/cloudquery/cloudquery/plugins/source/aws/client.AccountMultiplex"
  }
  deleteFilter "AccountDeleteFilter" {
    path = "github.com/cloudquery/cloudquery/plugins/source/aws/client.DeleteAccountFilter"
  }
  options {
    primary_keys = ["account_id", "id"]
  }
  column "role_id" {
    rename = "id"
  }
  userDefinedColumn "account_id" {
    type        = "string"
    description = "The AWS Account ID of the resource."
    resolver "resolveAWSAccount" {
      path = "github.com/cloudquery/cloudquery/plugins/source/aws/client.ResolveAWSAccount"
    }
  }
  userDefinedColumn "policies" {
    type              = "json"
    generate_resolver = true
    description       = "List of policies attached to group."
  }
  column "assume_role_policy_document" {
    type              = "json"
    generate_resolver = true
  }
  column "tags" {
    // TypeJson
    type = "json"
    resolver "resolveTags" {
      path = "github.com/cloudquery/cloudquery/plugins/source/aws/client.ResolveTags"
    }

  }
  user_relation "aws" "iam" "role_policies" {
    path = "github.com/aws/aws-sdk-go-v2/service/iam.GetRolePolicyOutput"
    description = "Inline policies that are embedded in the specified IAM role"
    column "policy_document" {
      type              = "json"
      generate_resolver = true
    }
  }
}
