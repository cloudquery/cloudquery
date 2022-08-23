//check-for-changes
service          = "aws"
output_directory = "."
add_generate     = true

description_modifier "remove_read_only" {
  words = ["  This member is required."]
}

resource "aws" "apigatewayv2" "apis" {
  path = "github.com/aws/aws-sdk-go-v2/service/apigatewayv2/types.Api"
  ignoreError "IgnoreCommonErrors" {
    path = "github.com/cloudquery/cloudquery/plugins/source/aws/client.IgnoreCommonErrors"
  }
  multiplex "AwsAccountRegion" {
    path   = "github.com/cloudquery/cloudquery/plugins/source/aws/client.ServiceAccountRegionMultiplexer"
    params = ["apigateway"]
  }
  options {
    primary_keys = ["arn"]
  }

  deleteFilter "AccountRegionFilter" {
    path = "github.com/cloudquery/cloudquery/plugins/source/aws/client.DeleteAccountRegionFilter"
  }
  userDefinedColumn "account_id" {
    description = "The AWS Account ID of the resource."
    type        = "string"
    resolver "resolveAWSAccount" {
      path = "github.com/cloudquery/cloudquery/plugins/source/aws/client.ResolveAWSAccount"
    }
  }
  userDefinedColumn "region" {
    type        = "string"
    description = "The AWS Region of the resource."
    resolver "resolveAWSRegion" {
      path = "github.com/cloudquery/cloudquery/plugins/source/aws/client.ResolveAWSRegion"
    }
  }
  userDefinedColumn "arn" {
    description = "The Amazon Resource Name (ARN) for the resource"
    type = "string"
    generate_resolver = true
  }

  column "api_id" {
    rename = "id"
  }

  user_relation "aws" "apigatewayv2" "authorizers" {
    path = "github.com/aws/aws-sdk-go-v2/service/apigatewayv2/types.Authorizer"
    userDefinedColumn "api_id" {
      description = "The API id"
      type = "string"
      resolver "resolveApiId" {
        path = "github.com/cloudquery/cq-provider-sdk/provider/schema.ParentPathResolver"
        params = ["ApiId"]
      }
    }
    userDefinedColumn "arn" {
      description = "The Amazon Resource Name (ARN) for the resource"
      type = "string"
      generate_resolver = true
    }
  }

  user_relation "aws" "apigatewayv2" "deployments" {
    path = "github.com/aws/aws-sdk-go-v2/service/apigatewayv2/types.Deployment"
    userDefinedColumn "api_id" {
      description = "The API id"
      type = "string"
      resolver "resolveApiId" {
        path = "github.com/cloudquery/cq-provider-sdk/provider/schema.ParentPathResolver"
        params = ["ApiId"]
      }
    }
    userDefinedColumn "arn" {
      description = "The Amazon Resource Name (ARN) for the resource"
      type = "string"
      generate_resolver = true
    }
  }

  user_relation "aws" "apigatewayv2" "integrations" {
    path = "github.com/aws/aws-sdk-go-v2/service/apigatewayv2/types.Integration"
    userDefinedColumn "api_id" {
      description = "The API id"
      type = "string"
      resolver "resolveApiId" {
        path = "github.com/cloudquery/cq-provider-sdk/provider/schema.ParentPathResolver"
        params = ["ApiId"]
      }
    }
    userDefinedColumn "arn" {
      description = "The Amazon Resource Name (ARN) for the resource"
      type = "string"
      generate_resolver = true
    }

    user_relation "aws" "apigatewayv2" "responses" {
      path = "github.com/aws/aws-sdk-go-v2/service/apigatewayv2/types.IntegrationResponse"
      userDefinedColumn "integration_id" {
        description = "Represents the identifier of an integration"
        type = "string"
        resolver "resolveIntegrationId" {
          path = "github.com/cloudquery/cq-provider-sdk/provider/schema.ParentPathResolver"
          params = ["IntegrationId"]
        }
      }
      userDefinedColumn "arn" {
        description = "The Amazon Resource Name (ARN) for the resource"
        type = "string"
        generate_resolver = true
      }
    }
  }

  user_relation "aws" "apigatewayv2" "models" {
    path = "github.com/aws/aws-sdk-go-v2/service/apigatewayv2/types.Model"
    userDefinedColumn "api_id" {
      description = "The API id"
      type = "string"
      resolver "resolveApiId" {
        path = "github.com/cloudquery/cq-provider-sdk/provider/schema.ParentPathResolver"
        params = ["ApiId"]
      }
    }
    userDefinedColumn "arn" {
      description = "The Amazon Resource Name (ARN) for the resource"
      type = "string"
      generate_resolver = true
    }
    userDefinedColumn "model_template" {
      type              = "string"
      generate_resolver = true
    }
  }

  user_relation "aws" "apigatewayv2" "routes" {
    path = "github.com/aws/aws-sdk-go-v2/service/apigatewayv2/types.Route"
    userDefinedColumn "api_id" {
      description = "The API id"
      type = "string"
      resolver "resolveApiId" {
        path = "github.com/cloudquery/cq-provider-sdk/provider/schema.ParentPathResolver"
        params = ["ApiId"]
      }
    }
    userDefinedColumn "arn" {
      description = "The Amazon Resource Name (ARN) for the resource"
      type = "string"
      generate_resolver = true
    }
    user_relation "aws" "apigatewayv2" "responses" {
      path = "github.com/aws/aws-sdk-go-v2/service/apigatewayv2/types.RouteResponse"
      userDefinedColumn "route_id" {
        description = "The Route id"
        type = "string"
        resolver "resolveRouteId" {
          path = "github.com/cloudquery/cq-provider-sdk/provider/schema.ParentPathResolver"
          params = ["RouteId"]
        }
      }
      userDefinedColumn "arn" {
        description = "The Amazon Resource Name (ARN) for the resource"
        type = "string"
        generate_resolver = true
      }
    }
  }

  user_relation "aws" "apigatewayv2" "stages" {
    path = "github.com/aws/aws-sdk-go-v2/service/apigatewayv2/types.Stage"
    userDefinedColumn "api_id" {
      description = "The API id"
      type = "string"
      resolver "resolveApiId" {
        path = "github.com/cloudquery/cq-provider-sdk/provider/schema.ParentPathResolver"
        params = ["ApiId"]
      }
    }
    userDefinedColumn "arn" {
      description = "The Amazon Resource Name (ARN) for the resource"
      type = "string"
      generate_resolver = true
    }
    column "default_route_settings" {
      rename = "route_settings"
    }
  }
}


resource "aws" "apigatewayv2" "domain_names" {
  path = "github.com/aws/aws-sdk-go-v2/service/apigatewayv2/types.DomainName"
  ignoreError "IgnoreCommonErrors" {
    path = "github.com/cloudquery/cloudquery/plugins/source/aws/client.IgnoreCommonErrors"
  }
  multiplex "AwsAccountRegion" {
    path   = "github.com/cloudquery/cloudquery/plugins/source/aws/client.ServiceAccountRegionMultiplexer"
    params = ["apigateway"]
  }
  deleteFilter "AccountRegionFilter" {
    path = "github.com/cloudquery/cloudquery/plugins/source/aws/client.DeleteAccountRegionFilter"
  }
  options {
    primary_keys = ["arn"]
  }
  userDefinedColumn "account_id" {
    description = "The AWS Account ID of the resource."
    type        = "string"
    resolver "resolveAWSAccount" {
      path = "github.com/cloudquery/cloudquery/plugins/source/aws/client.ResolveAWSAccount"
    }
  }
  userDefinedColumn "region" {
    type        = "string"
    description = "The AWS Region of the resource."
    resolver "resolveAWSRegion" {
      path = "github.com/cloudquery/cloudquery/plugins/source/aws/client.ResolveAWSRegion"
    }
  }
  userDefinedColumn "arn" {
    description = "The Amazon Resource Name (ARN) for the resource"
    type = "string"
    generate_resolver = true
  }

  user_relation "aws" "apigatewayv2" "rest_api_mappings" {
    path = "github.com/aws/aws-sdk-go-v2/service/apigatewayv2/types.ApiMapping"
    userDefinedColumn "arn" {
      description = "The Amazon Resource Name (ARN) for the resource"
      type = "string"
      generate_resolver = true
    }
  }
}

resource "aws" "apigatewayv2" "vpc_links" {
  path = "github.com/aws/aws-sdk-go-v2/service/apigatewayv2/types.VpcLink"
  ignoreError "IgnoreCommonErrors" {
    path = "github.com/cloudquery/cloudquery/plugins/source/aws/client.IgnoreCommonErrors"
  }
  multiplex "AwsAccountRegion" {
    path   = "github.com/cloudquery/cloudquery/plugins/source/aws/client.ServiceAccountRegionMultiplexer"
    params = ["apigateway"]
  }
  deleteFilter "AccountRegionFilter" {
    path = "github.com/cloudquery/cloudquery/plugins/source/aws/client.DeleteAccountRegionFilter"
  }
  options {
    primary_keys = ["arn"]
  }
  userDefinedColumn "account_id" {
    description = "The AWS Account ID of the resource."
    type        = "string"
    resolver "resolveAWSAccount" {
      path = "github.com/cloudquery/cloudquery/plugins/source/aws/client.ResolveAWSAccount"
    }
  }
  userDefinedColumn "region" {
    type        = "string"
    description = "The AWS Region of the resource."
    resolver "resolveAWSRegion" {
      path = "github.com/cloudquery/cloudquery/plugins/source/aws/client.ResolveAWSRegion"
    }
  }
  userDefinedColumn "arn" {
    description = "The Amazon Resource Name (ARN) for the resource"
    type = "string"
    generate_resolver = true
  }

  column "vpc_link_id" {
    rename = "id"
  }
}
