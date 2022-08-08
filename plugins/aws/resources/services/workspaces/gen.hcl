service = "aws"
output_directory = "."
add_generate = true

resource "aws" "workspaces" "workspaces" {
  path = "github.com/aws/aws-sdk-go-v2/service/workspaces/types.Workspace"
  ignoreError "IgnoreAccessDenied" {
    path = "github.com/cloudquery/cq-provider-aws/client.IgnoreAccessDeniedServiceDisabled"
  }
  deleteFilter "AccountRegionFilter" {
    path = "github.com/cloudquery/cq-provider-aws/client.DeleteAccountRegionFilter"
  }
  multiplex "AwsAccountRegion" {
    path = "github.com/cloudquery/cq-provider-aws/client.ServiceAccountRegionMultiplexer"
    params = ["workspaces"]
  }


  options {
    primary_keys = ["id"]
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

  column "workspace_id" {
    rename = "id"
  }

  column "workspace_properties" {
    skip_prefix = true
  }

  column "modification_states" {
    type              = "JSON"
    generate_resolver = true
  }

  userDefinedColumn "arn" {
    type        = "string"
    description = "The Amazon Resource Name (ARN) for the workspaces workspace"
    generate_resolver = false
  }
}

resource "aws" "workspaces" "directories" {
  path = "github.com/aws/aws-sdk-go-v2/service/workspaces/types.WorkspaceDirectory"
  ignoreError "IgnoreAccessDenied" {
    path = "github.com/cloudquery/cq-provider-aws/client.IgnoreAccessDeniedServiceDisabled"
  }
  deleteFilter "AccountRegionFilter" {
    path = "github.com/cloudquery/cq-provider-aws/client.DeleteAccountRegionFilter"
  }
  multiplex "AwsAccountRegion" {
    path = "github.com/cloudquery/cq-provider-aws/client.ServiceAccountRegionMultiplexer"
    params = ["workspaces"]
  }


  options {
    primary_keys = ["id"]
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

  column "directory_id" {
    rename = "id"
  }

  column "directory_type" {
    rename = "type"
  }

  column "directory_name" {
    rename = "name"
  }

  column "workspace_creation_properties" {
    skip_prefix = true
  }

  column "workspace_access_properties" {
    skip_prefix = true
  }

  column "selfservice_permissions" {
    skip_prefix = true
  }

  userDefinedColumn "arn" {
    type        = "string"
    description = "The Amazon Resource Name (ARN) for the workspaces directory"
    generate_resolver = false
  }
}