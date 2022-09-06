//check-for-changes
service          = "aws"
output_directory = "."
add_generate     = true

description_modifier "remove_read_only" {
  words = ["  This member is required."]
}

resource "aws" "apigateway" "rest_apis" {
  path = "github.com/aws/aws-sdk-go-v2/service/apigateway/types.RestApi"
  ignoreError "IgnoreCommonErrors" {
    path = "github.com/cloudquery/cloudquery/plugins/source/aws/client.IgnoreCommonErrors"
  }
  deleteFilter "AccountRegionFilter" {
    path = "github.com/cloudquery/cloudquery/plugins/source/aws/client.DeleteAccountRegionFilter"
  }
  multiplex "AwsAccountRegion" {
    path   = "github.com/cloudquery/cloudquery/plugins/source/aws/client.ServiceAccountRegionMultiplexer"
    params = ["apigateway"]
  }
  options {
    primary_keys = ["arn"]
  }
  userDefinedColumn "account_id" {
    description = "The AWS Account ID of the resource"
    type        = "string"
    resolver "resolveAWSAccount" {
      path = "github.com/cloudquery/cloudquery/plugins/source/aws/client.ResolveAWSAccount"
    }
  }
  userDefinedColumn "region" {
    type        = "string"
    description = "The AWS Region of the resource"
    resolver "resolveAWSRegion" {
      path = "github.com/cloudquery/cloudquery/plugins/source/aws/client.ResolveAWSRegion"
    }
  }

  userDefinedColumn "arn" {
    description = "The Amazon Resource Name (ARN) for the resource"
    type = "string"
    generate_resolver = true
  }

  user_relation "aws" "apigateway" "authorizers" {
    path = "github.com/aws/aws-sdk-go-v2/service/apigateway/types.Authorizer"
    userDefinedColumn "rest_api_id" {
      description = "The API's identifier"
      type = "string"
      resolver "resolveRestApiId" {
        path = "github.com/cloudquery/cq-provider-sdk/provider/schema.ParentPathResolver"
        params = ["Id"]
      }
    }
    userDefinedColumn "arn" {
      description = "The Amazon Resource Name (ARN) for the resource"
      type = "string"
      generate_resolver = true
    }

    column "provider_arn_s" {
      rename = "provider_arns"
    }
  }

  user_relation "aws" "apigateway" "deployments" {
    path = "github.com/aws/aws-sdk-go-v2/service/apigateway/types.Deployment"
    userDefinedColumn "rest_api_id" {
      description = "The API's identifier"
      type = "string"
      resolver "resolveRestApiId" {
        path = "github.com/cloudquery/cq-provider-sdk/provider/schema.ParentPathResolver"
        params = ["Id"]
      }
    }
    userDefinedColumn "arn" {
      description = "The Amazon Resource Name (ARN) for the resource"
      type = "string"
      generate_resolver = true
    }
  }

  user_relation "aws" "apigateway" "documentation_parts" {
    path = "github.com/aws/aws-sdk-go-v2/service/apigateway/types.DocumentationPart"
    userDefinedColumn "rest_api_id" {
      description = "The API's identifier"
      type = "string"
      resolver "resolveRestApiId" {
        path = "github.com/cloudquery/cq-provider-sdk/provider/schema.ParentPathResolver"
        params = ["Id"]
      }
    }
    userDefinedColumn "arn" {
      description = "The Amazon Resource Name (ARN) for the resource"
      type = "string"
      generate_resolver = true
    }
  }

  user_relation "aws" "apigateway" "documentation_versions" {
    path = "github.com/aws/aws-sdk-go-v2/service/apigateway/types.DocumentationVersion"
    userDefinedColumn "rest_api_id" {
      description = "The API's identifier"
      type = "string"
      resolver "resolveRestApiId" {
        path = "github.com/cloudquery/cq-provider-sdk/provider/schema.ParentPathResolver"
        params = ["Id"]
      }
    }
    userDefinedColumn "arn" {
      description = "The Amazon Resource Name (ARN) for the resource"
      type = "string"
      generate_resolver = true
    }
  }

  user_relation "aws" "apigateway" "gateway_responses" {
    path = "github.com/aws/aws-sdk-go-v2/service/apigateway/types.GatewayResponse"
    userDefinedColumn "rest_api_id" {
      description = "The API's identifier"
      type = "string"
      resolver "resolveRestApiId" {
        path = "github.com/cloudquery/cq-provider-sdk/provider/schema.ParentPathResolver"
        params = ["Id"]
      }
    }
    userDefinedColumn "arn" {
      description = "The Amazon Resource Name (ARN) for the resource"
      type = "string"
      generate_resolver = true
    }
  }

  user_relation "aws" "apigateway" "models" {
    path = "github.com/aws/aws-sdk-go-v2/service/apigateway/types.Model"
    userDefinedColumn "rest_api_id" {
      description = "The API's identifier"
      type = "string"
      resolver "resolveRestApiId" {
        path = "github.com/cloudquery/cq-provider-sdk/provider/schema.ParentPathResolver"
        params = ["Id"]
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

  user_relation "aws" "apigateway" "request_validators" {
    path = "github.com/aws/aws-sdk-go-v2/service/apigateway/types.RequestValidator"
    userDefinedColumn "rest_api_id" {
      description = "The API's identifier"
      type = "string"
      resolver "resolveRestApiId" {
        path = "github.com/cloudquery/cq-provider-sdk/provider/schema.ParentPathResolver"
        params = ["Id"]
      }
    }
    userDefinedColumn "arn" {
      description = "The Amazon Resource Name (ARN) for the resource"
      type = "string"
      generate_resolver = true
    }
  }

  user_relation "aws" "apigateway" "resources" {
    path = "github.com/aws/aws-sdk-go-v2/service/apigateway/types.Resource"
    userDefinedColumn "rest_api_id" {
      description = "The API's identifier"
      type = "string"
      resolver "resolveRestApiId" {
        path = "github.com/cloudquery/cq-provider-sdk/provider/schema.ParentPathResolver"
        params = ["Id"]
      }
    }
    userDefinedColumn "arn" {
      description = "The Amazon Resource Name (ARN) for the resource"
      type = "string"
      generate_resolver = true
    }
  }

  user_relation "aws" "apigateway" "stages" {
    path = "github.com/aws/aws-sdk-go-v2/service/apigateway/types.Stage"
    userDefinedColumn "rest_api_id" {
      description = "The API's identifier"
      type = "string"
      resolver "resolveRestApiId" {
        path = "github.com/cloudquery/cq-provider-sdk/provider/schema.ParentPathResolver"
        params = ["Id"]
      }
    }
    userDefinedColumn "arn" {
      description = "The Amazon Resource Name (ARN) for the resource"
      type = "string"
      generate_resolver = true
    }
  }
}
