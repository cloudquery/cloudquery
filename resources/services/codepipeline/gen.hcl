service          = "aws"
output_directory = "."
add_generate     = true

resource "aws" "codepipeline" "pipelines" {
  path = "github.com/aws/aws-sdk-go-v2/service/codepipeline.GetPipelineOutput"
  ignoreError "IgnoreAccessDenied" {
    path = "github.com/cloudquery/cq-provider-aws/client.IgnoreAccessDeniedServiceDisabled"
  }
  deleteFilter "AccountRegionFilter" {
    path = "github.com/cloudquery/cq-provider-aws/client.DeleteAccountRegionFilter"
  }
  multiplex "AwsAccountRegion" {
    path   = "github.com/cloudquery/cq-provider-aws/client.ServiceAccountRegionMultiplexer"
    params = ["codepipeline"]
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


  column "metadata" {
    skip_prefix = true
  }

  column "pipeline" {
    skip_prefix = true
  }

  column "result_metadata" {
    skip = true
  }

  column "pipeline_arn" {
    rename = "arn"
  }

  relation "aws" "codepipeline" "stages" {

    column "blockers" {
      type              = "JSON"
      generate_resolver = true
    }

    relation "aws" "codepipeline" "actions" {
      path = "github.com/aws/aws-sdk-go-v2/service/codepipeline/types.ActionDeclaration"

      column "action_type_id" {
        skip_prefix = true
      }

      column "input_artifacts" {
        type              = "stringarray"
        generate_resolver = true
      }

      column "output_artifacts" {
        type              = "stringarray"
        generate_resolver = true
      }
    }

    userDefinedColumn "stage_order" {
      type              = "int"
      generate_resolver = false
      description       = "The stage order in the pipeline."
    }
  }

  userDefinedColumn "tags" {
    type              = "json"
    generate_resolver = true
    description       = "The tags associated with the pipeline."
  }

}