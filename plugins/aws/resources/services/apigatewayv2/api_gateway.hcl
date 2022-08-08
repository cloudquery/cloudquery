service = "aws"
output_directory = "."
add_generate = true



resource "aws" "apigatewayv2" "apis" {
  path = "github.com/aws/aws-sdk-go-v2/service/apigatewayv2/types.Api"
  ignoreError "IgnoreAccessDenied" {
    path = "github.com/cloudquery/cq-provider-aws/client.IgnoreAccessDeniedServiceDisabled"
  }
  multiplex "AwsAccountRegion" {
    path = "github.com/cloudquery/cq-provider-aws/client.ServiceAccountRegionMultiplexer"
    params = ["apigatewayv2", 2]
  }

  deleteFilter "AccountRegionFilter" {
    path = "github.com/cloudquery/cq-provider-aws/client.DeleteAccountRegionFilter"
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

  user_relation "aws" "apigatewayv2" "authorizers" {
    path = "github.com/aws/aws-sdk-go-v2/service/apigatewayv2/types.Authorizer"
  }

  user_relation "aws" "apigatewayv2" "deployments" {
    path = "github.com/aws/aws-sdk-go-v2/service/apigatewayv2/types.Deployment"
  }

  user_relation "aws" "apigatewayv2" "integrations" {
    path = "github.com/aws/aws-sdk-go-v2/service/apigatewayv2/types.Integration"

    user_relation "aws" "apigatewayv2" "responses" {
      path = "github.com/aws/aws-sdk-go-v2/service/apigatewayv2/types.IntegrationResponse"
    }
  }

  user_relation "aws" "apigatewayv2" "models" {
    path = "github.com/aws/aws-sdk-go-v2/service/apigatewayv2/types.Model"
    userDefinedColumn "model_template" {
      type              = "string"
      generate_resolver = true
    }
  }

  user_relation "aws" "apigatewayv2" "routes" {
    path = "github.com/aws/aws-sdk-go-v2/service/apigatewayv2/types.Route"
    user_relation "aws" "apigatewayv2" "responses" {
      path = "github.com/aws/aws-sdk-go-v2/service/apigatewayv2/types.RouteResponse"
    }
  }

  user_relation "aws" "apigatewayv2" "stages" {
    path = "github.com/aws/aws-sdk-go-v2/service/apigatewayv2/types.Stage"
    column "default_route_settings_data_trace_enabled" {
      rename = "route_settings_data_trace_enabled"
    }
    column "default_route_settings_detailed_metrics_enabled" {
      rename = "route_settings_detailed_metrics_enabled"
    }
    column "default_route_settings_logging_level" {
      rename = "route_settings_logging_level"
    }
    column "default_route_settings_throttling_burst_limit" {
      rename = "route_settings_throttling_burst_limit"
    }
    column "default_route_settings_throttling_rate_limit" {
      rename = "route_settings_throttling_rate_limit"
    }
  }
}


resource "aws" "apigatewayv2" "domain_names" {
  path = "github.com/aws/aws-sdk-go-v2/service/apigatewayv2/types.DomainName"
  ignoreError "IgnoreAccessDenied" {
    path = "github.com/cloudquery/cq-provider-aws/client.IgnoreAccessDeniedServiceDisabled"
  }
  multiplex "AwsAccountRegion" {
    path = "github.com/cloudquery/cq-provider-aws/client.AccountRegionMultiplex"
  }
  deleteFilter "AccountRegionFilter" {
    path = "github.com/cloudquery/cq-provider-aws/client.DeleteAccountRegionFilter"
  }
  userDefinedColumn "account_id" {
    description = "The AWS Account ID of the resource."
    type        = "string"
    resolver "resolveAWSAccount" {
      path = "github.com/cloudquery/cq-provider-aws/client.ResolveAWSAccount"
    }
  }

  relation "aws" "apigatewayv2" "rest_api_mappings" {
    path = "github.com/aws/aws-sdk-go-v2/service/apigatewayv2/types.ApiMapping"

    //    column "rest_api_id" {
    //      skip = true
    //    }
  }
  userDefinedColumn "region" {
    type        = "string"
    description = "The AWS Region of the resource."
    resolver "resolveAWSRegion" {
      path = "github.com/cloudquery/cq-provider-aws/client.ResolveAWSRegion"
    }
  }
}

resource "aws" "apigatewayv2" "vpc_links" {
  path = "github.com/aws/aws-sdk-go-v2/service/apigatewayv2/types.VpcLink"
  ignoreError "IgnoreAccessDenied" {
    path = "github.com/cloudquery/cq-provider-aws/client.IgnoreAccessDeniedServiceDisabled"
  }
  multiplex "AwsAccountRegion" {
    path = "github.com/cloudquery/cq-provider-aws/client.AccountRegionMultiplex"
  }
  deleteFilter "AccountRegionFilter" {
    path = "github.com/cloudquery/cq-provider-aws/client.DeleteAccountRegionFilter"
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

  options {
    primary_keys = [
      "account_id",
      "id"
    ]
  }

  column "vpc_link_id" {
    rename = "id"
  }

  column "tags" {
    generate_resolver = true
  }
}


resource "aws" "apigateway" "rest_apis" {
  path = "github.com/aws/aws-sdk-go-v2/service/apigateway/types.RestApi"
  ignoreError "IgnoreAccessDenied" {
    path = "github.com/cloudquery/cq-provider-aws/client.IgnoreAccessDeniedServiceDisabled"
  }
  multiplex "AwsAccountRegion" {
    path = "github.com/cloudquery/cq-provider-aws/client.AccountRegionMultiplex"
  }
  deleteFilter "AccountRegionFilter" {
    path = "github.com/cloudquery/cq-provider-aws/client.DeleteAccountRegionFilter"
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

  column "id" {
    rename = "resource_id"
  }

  relation "aws" "apigateway" "rest_api_authorizers" {
    path = "github.com/aws/aws-sdk-go-v2/service/apigateway/types.Authorizer"

    column "id" {
      rename = "resource_id"
    }

    column "provider_arn_s" {
      rename = "provider_arns"
    }
  }

  relation "aws" "apigateway" "rest_api_deployments" {
    path = "github.com/aws/aws-sdk-go-v2/service/apigateway/types.Deployment"
    column "id" {
      rename = "resource_id"
    }
  }

  relation "aws" "apigateway" "rest_api_documentation_parts" {
    path = "github.com/aws/aws-sdk-go-v2/service/apigateway/types.DocumentationPart"
    column "id" {
      rename = "documentation_part_id"
    }
  }

  relation "aws" "apigateway" "rest_api_documentation_versions" {
    path = "github.com/aws/aws-sdk-go-v2/service/apigateway/types.DocumentationVersion"
  }

  relation "aws" "apigateway" "rest_api_gateway_responses" {
    path = "github.com/aws/aws-sdk-go-v2/service/apigateway/types.GatewayResponse"
  }

  relation "aws" "apigateway" "rest_api_models" {
    path = "github.com/aws/aws-sdk-go-v2/service/apigateway/types.Model"

    column "id" {
      rename = "resource_id"
    }

    userDefinedColumn "model_template" {
      type              = "string"
      generate_resolver = true
    }
  }
  relation "aws" "apigateway" "rest_api_request_validators" {
    path = "github.com/aws/aws-sdk-go-v2/service/apigateway/types.RequestValidator"
    column "id" {
      rename = "resource_id"
    }
  }

  relation "aws" "apigateway" "rest_api_resources" {
    path = "github.com/aws/aws-sdk-go-v2/service/apigateway/types.Resource"
    column "id" {
      rename = "resource_id"
    }
  }

  relation "aws" "apigateway" "rest_api_stages" {
    path = "github.com/aws/aws-sdk-go-v2/service/apigateway/types.Stage"
  }
}

resource "aws" "apigateway" "usage_plans" {
  path = "github.com/aws/aws-sdk-go-v2/service/apigateway/types.UsagePlan"
  ignoreError "IgnoreAccessDenied" {
    path = "github.com/cloudquery/cq-provider-aws/client.IgnoreAccessDeniedServiceDisabled"
  }
  multiplex "AwsAccountRegion" {
    path = "github.com/cloudquery/cq-provider-aws/client.AccountRegionMultiplex"
  }
  deleteFilter "AccountRegionFilter" {
    path = "github.com/cloudquery/cq-provider-aws/client.DeleteAccountRegionFilter"
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

  column "id" {
    rename = "resource_id"
  }

  relation "aws" "apigateway" "usage_plan_keys" {
    path = "github.com/aws/aws-sdk-go-v2/service/apigateway/types.UsagePlanKey"

    column "id" {
      rename = "resource_id"
    }
  }
}

resource "aws" "apigateway" "domain_names" {
  path = "github.com/aws/aws-sdk-go-v2/service/apigateway/types.DomainName"
  ignoreError "IgnoreAccessDenied" {
    path = "github.com/cloudquery/cq-provider-aws/client.IgnoreAccessDeniedServiceDisabled"
  }
  multiplex "AwsAccountRegion" {
    path = "github.com/cloudquery/cq-provider-aws/client.AccountRegionMultiplex"
  }
  deleteFilter "AccountRegionFilter" {
    path = "github.com/cloudquery/cq-provider-aws/client.DeleteAccountRegionFilter"
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

  relation "aws" "apigateway" "domain_name_base_path_mappings" {
    path = "github.com/aws/aws-sdk-go-v2/service/apigateway/types.BasePathMapping"
  }
}


resource "aws" "apigateway" "client_certificates" {
  path = "github.com/aws/aws-sdk-go-v2/service/apigateway/types.ClientCertificate"
  ignoreError "IgnoreAccessDenied" {
    path = "github.com/cloudquery/cq-provider-aws/client.IgnoreAccessDeniedServiceDisabled"
  }
  multiplex "AwsAccountRegion" {
    path = "github.com/cloudquery/cq-provider-aws/client.AccountRegionMultiplex"
  }
  deleteFilter "AccountRegionFilter" {
    path = "github.com/cloudquery/cq-provider-aws/client.DeleteAccountRegionFilter"
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
}


resource "aws" "apigateway" "api_keys" {
  path = "github.com/aws/aws-sdk-go-v2/service/apigateway/types.ApiKey"
  ignoreError "IgnoreAccessDenied" {
    path = "github.com/cloudquery/cq-provider-aws/client.IgnoreAccessDeniedServiceDisabled"
  }
  multiplex "AwsAccountRegion" {
    path = "github.com/cloudquery/cq-provider-aws/client.AccountRegionMultiplex"
  }
  deleteFilter "AccountRegionFilter" {
    path = "github.com/cloudquery/cq-provider-aws/client.DeleteAccountRegionFilter"
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
  column "id" {
    rename = "resource_id"
  }
}

resource "aws" "apigateway" "vpc_links" {
  path = "github.com/aws/aws-sdk-go-v2/service/apigateway/types.VpcLink"
  ignoreError "IgnoreAccessDenied" {
    path = "github.com/cloudquery/cq-provider-aws/client.IgnoreAccessDeniedServiceDisabled"
  }
  multiplex "AwsAccountRegion" {
    path = "github.com/cloudquery/cq-provider-aws/client.AccountRegionMultiplex"
  }
  deleteFilter "AccountRegionFilter" {
    path = "github.com/cloudquery/cq-provider-aws/client.DeleteAccountRegionFilter"
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

  column "id" {
    rename = "resource_id"
  }
}
