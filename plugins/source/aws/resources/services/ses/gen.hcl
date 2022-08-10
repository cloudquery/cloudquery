service          = "aws"
output_directory = "."
add_generate     = true

resource "aws" "ses" "templates" {
  path        = "github.com/cloudquery/cq-provider-aws/resources/services/ses.Template"
  description = "Amazon Simple Email Service (SES) is a cost-effective, flexible, and scalable email service that enables developers to send mail from within any application."
  ignoreError "IgnoreAccessDenied" {
    path = "github.com/cloudquery/cq-provider-aws/client.IgnoreAccessDeniedServiceDisabled"
  }
  multiplex "AwsAccountRegion" {
    path   = "github.com/cloudquery/cq-provider-aws/client.ServiceAccountRegionMultiplexer"
    params = ["email"]
  }
  deleteFilter "AccountRegionFilter" {
    path = "github.com/cloudquery/cq-provider-aws/client.DeleteAccountRegionFilter"
  }

  options {
    primary_keys = [
      "arn"
    ]
  }

  column "template_name" {
    rename      = "name"
    description = "The name of the template."
  }
  column "email_template_content_html" {
    rename = "html"
  }
  column "email_template_content_subject" {
    rename = "subject"
  }
  column "email_template_content_text" {
    rename = "text"
  }
  column "created_timestamp" {
    type        = "timestamp"
    description = "The time and date the template was created."
  }

  userDefinedColumn "arn" {
    type              = "string"
    description       = "The Amazon Resource Name (ARN) for the resource."
    generate_resolver = true
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
}