//check-for-changes
service          = "aws"
output_directory = "."
add_generate     = true
resource "aws" "iam" "roles" {
  path        = "github.com/aws/aws-sdk-go-v2/service/iam/types.Role"
  description = "An IAM role is an IAM identity that you can create in your account that has specific permissions."
  ignoreError "IgnoreAccessDenied" {
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
    type              = "json"
    generate_resolver = true
  }
  user_relation "aws" "iam" "role_policies" {
    path = "github.com/aws/aws-sdk-go-v2/service/iam.GetRolePolicyOutput"
    column "policy_document" {
      type              = "json"
      generate_resolver = true
    }
  }
}
resource "aws" "iam" "users" {
  path = "github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/iam.WrappedUser"
  // description = "An IAM user is an IAM identity that you can create in your account that has specific permissions."
  ignoreError "IgnoreAccessDenied" {
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

  column "user_arn" {
    skip = true
  }
  column "user_create_date" {
    rename = "create_date"
  }
  column "user_password_last_used" {
    rename = "password_last_used"
  }
  column "user_path" {
    rename = "path"
  }
  column "user_permissions_boundary_permissions_boundary_arn" {
    rename = "permissions_boundary_arn"
  }
  column "user_permissions_boundary_permissions_boundary_type" {
    rename = "permissions_boundary_type"
  }
  column "user_tags" {
    rename = "tags"
    type   = "json"

  }

  column "user_id" {
    rename = "id"
  }
  column "password_status" {
    type = "bool"
    rename = "password_enabled"
  }
  column "report_user_arn" {
    skip = true
  }
  column "report_user" {
    skip_prefix = true
  }


  column "password_last_changed"{
    type = "timestamp"
  }
  column "password_next_rotation"{
    type = "timestamp"
  }

  column "cert1_last_rotated"{
    type = "timestamp"
  }
  column "access_key1_last_rotated"{
    type = "timestamp"
  }
  column "cert2_last_rotated"{
    type = "timestamp"
  }
  column "access_key2_last_rotated"{
    type = "timestamp"
  }

  userDefinedColumn "account_id" {
    type        = "string"
    description = "The AWS Account ID of the resource."
    resolver "resolveAWSAccount" {
      path = "github.com/cloudquery/cloudquery/plugins/source/aws/client.ResolveAWSAccount"
    }
  }
  postResourceResolver "PostResourceResolver" {
    path = "github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/iam.postIamUserResolver"
  }
  user_relation "aws" "iam" "access_keys" {
    path = "github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/iam.wrappedKey"
    column "access_key_metadata" {
      skip_prefix = true
    }
    column "access_key_last_used" {
      rename = "last_used"
    }
     column "last_used_last_used_date"{
      rename = "last_used_date"
    }
   
  }
  user_relation "aws" "iam" "groups" {
    path = "github.com/aws/aws-sdk-go-v2/service/iam/types.Group"
  }
  user_relation "aws" "iam" "attached_policies" {
    path = "github.com/aws/aws-sdk-go-v2/service/iam/types.AttachedPolicy"
  }
    user_relation "aws" "iam" "user_policies" {
    path = "github.com/aws/aws-sdk-go-v2/service/iam.GetUserPolicyOutput"
    column "policy_document" {
      type              = "json"
      generate_resolver = true
    }
  }

}