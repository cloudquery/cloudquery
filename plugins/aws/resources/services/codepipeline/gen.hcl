//check-for-changes
service          = "aws"
output_directory = "."
add_generate     = true

description_modifier "remove_read_only" {
  words = ["  This member is required."]
}

description_modifier "remove_extra_info" {
  regex = "([*].+)"
}

resource "aws" "codepipeline" "pipelines" {
  path = "github.com/aws/aws-sdk-go-v2/service/codepipeline.GetPipelineOutput"
  ignoreError "IgnoreCommonErrors" {
    path = "github.com/cloudquery/cq-provider-aws/client.IgnoreCommonErrors"
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
    ignore_columns_in_tests = ["blockers"]

    resolver "fetchCodepipelinePipelineStages" {
      generate = true
    }

    column "blockers" {
      type              = "JSON"
      resolver "pathResolver" {
        path = "github.com/cloudquery/cq-provider-sdk/provider/schema.PathResolver"
        params = ["Blockers"]
      }
    }

    relation "aws" "codepipeline" "actions" {
      path = "github.com/aws/aws-sdk-go-v2/service/codepipeline/types.ActionDeclaration"
      ignore_columns_in_tests = ["namespace", "region", "role_arn"]

      column "action_type_id" {
        skip_prefix = true
      }

      column "input_artifacts" {
        type              = "stringarray"
        resolver "pathResolver" {
          path = "github.com/cloudquery/cq-provider-sdk/provider/schema.PathResolver"
          params = ["InputArtifacts.Name"]
        }
      }

      column "output_artifacts" {
        type              = "stringarray"
        resolver "pathResolver" {
          path = "github.com/cloudquery/cq-provider-sdk/provider/schema.PathResolver"
          params = ["OutputArtifacts.Name"]
        }
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

resource "aws" "codepipeline" "webhooks" {
  path = "github.com/aws/aws-sdk-go-v2/service/codepipeline/types.ListWebhookItem"
  ignoreError "IgnoreCommonErrors" {
    path = "github.com/cloudquery/cq-provider-aws/client.IgnoreCommonErrors"
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

  column "authentication_configuration" {
    rename = "authentication"
  }


  column "definition" {
    skip_prefix = true
  }

  column "tags" {
    type = "JSON"
    resolver "resolveTags" {
      path = "github.com/cloudquery/cq-provider-aws/client.ResolveTags"
    }
    description = "The tags associated with the webhook."
  }
}