service          = "aws"
output_directory = "."
add_generate     = true

resource "aws" "glue" "registries" {
  path = "github.com/aws/aws-sdk-go-v2/service/glue/types.RegistryListItem"
  ignoreError "IgnoreAccessDenied" {
    path = "github.com/cloudquery/cq-provider-aws/client.IgnoreAccessDeniedServiceDisabled"
  }
  deleteFilter "AccountRegionFilter" {
    path = "github.com/cloudquery/cq-provider-aws/client.DeleteAccountRegionFilter"
  }
  multiplex "AwsAccountRegion" {
    path   = "github.com/cloudquery/cq-provider-aws/client.ServiceAccountRegionMultiplexer"
    params = ["glue"]
  }
  options {
    primary_keys = ["arn"]
  }
  column "registry_arn" {
    rename = "arn"
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

  userDefinedColumn "tags" {
    type              = "json"
    description       = "Resource tags."
    generate_resolver = true
  }
  user_relation "aws" "glue" "schemas" {
    path        = "github.com/aws/aws-sdk-go-v2/service/glue.GetSchemaOutput"
    description = "An object that contains minimal details for a schema"
    userDefinedColumn "tags" {
      type              = "json"
      description       = "Resource tags."
      generate_resolver = true
    }
    column "schema_arn" {
      rename = "arn"
    }

    user_relation "aws" "glue" "versions" {
      path        = "github.com/aws/aws-sdk-go-v2/service/glue.GetSchemaVersionOutput"
      description = "An object containing the details about a schema version"
      userDefinedColumn "metadata" {
        type              = "json"
        generate_resolver = true
      }

      column "schema_version_id" {
        rename = "id"
      }
    }
  }
}
