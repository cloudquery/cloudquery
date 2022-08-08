service          = "aws"
output_directory = "."
add_generate     = true


description_modifier "remove_read_only" {
  words = ["  This member is required."]
}

resource "aws" "glue" "ml_transforms" {
  path = "github.com/aws/aws-sdk-go-v2/service/glue/types.MLTransform"
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
  userDefinedColumn "arn" {
    type              = "string"
    description       = "The Amazon Resource Name (ARN) of the workflow."
    generate_resolver = true
  }

  column "evaluation_metrics_find_matches_metrics_confusion_matrix" {
    type = "json"
  }
  column "transform_encryption_ml_user_data_encryption_ml_user_data_encryption_mode" {
    rename = "transform_encryption_user_data_encryption_mode"
  }

  column "schema" {
    type              = "json"
    generate_resolver = true
  }

  column "evaluation_metrics_find_matches_metrics_column_importances" {
    type              = "json"
    generate_resolver = true
  }

  column "evaluation_metrics_find_matches_metrics_area_under_p_r_curve" {
    rename = "evaluation_metrics_find_matches_metrics_area_under_pr_curve"
  }

  column "transform_id" {
    rename = "id"
  }

  userDefinedColumn "tags" {
    type              = "json"
    description       = "Resource tags"
    generate_resolver = true
  }

  user_relation "aws" "glue" "task_runs" {
    path = "github.com/aws/aws-sdk-go-v2/service/glue/types.TaskRun"
    column "properties" {
      skip_prefix = true
    }
    column "task_run_id" {
      rename = "id"
    }
  }
}
