service          = "aws"
output_directory = "."
add_generate     = true

resource "aws" "qldb" "ledgers" {
  path = "github.com/aws/aws-sdk-go-v2/service/qldb.DescribeLedgerOutput"
  ignoreError "IgnoreAccessDenied" {
    path = "github.com/cloudquery/cq-provider-aws/client.IgnoreAccessDeniedServiceDisabled"
  }
  deleteFilter "AccountRegionFilter" {
    path = "github.com/cloudquery/cq-provider-aws/client.DeleteAccountRegionFilter"
  }
  multiplex "AwsAccountRegion" {
    path   = "github.com/cloudquery/cq-provider-aws/client.ServiceAccountRegionMultiplexer"
    params = ["qldb"]
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

  user_relation "aws" "qldb" "journal_kinesis_streams" {
    path = "github.com/aws/aws-sdk-go-v2/service/qldb/types.JournalKinesisStreamDescription"

    column "kinesis_configuration" {
      skip_prefix = true
    }
  }

  user_relation "aws" "qldb" "journal_s3_exports" {
    path = "github.com/aws/aws-sdk-go-v2/service/qldb/types.JournalS3ExportDescription"

    column "s3_export_configuration" {
      skip_prefix = true
    }

    column "encryption_configuration" {
      skip_prefix = true
    }
  }

  column "encryption_description" {
    skip_prefix = true
  }

  column "result_metadata" {
    skip = true
  }


  userDefinedColumn "tags" {
    type              = "json"
    generate_resolver = true
    description       = "The tags associated with the pipeline."
  }

}