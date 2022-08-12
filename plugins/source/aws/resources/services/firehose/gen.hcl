//check-for-changes
service          = "aws"
output_directory = "."
add_generate     = true

description_modifier "remove_read_only" {
  words = ["  This member is required."]
}


resource "aws" "firehose" "delivery_streams" {
  path = "github.com/aws/aws-sdk-go-v2/service/firehose/types.DeliveryStreamDescription"
  ignoreError "IgnoreCommonErrors" {
    path = "github.com/cloudquery/cloudquery/plugins/source/aws/client.IgnoreCommonErrors"
  }
  deleteFilter "AccountRegionFilter" {
    path = "github.com/cloudquery/cloudquery/plugins/source/aws/client.DeleteAccountRegionFilter"
  }
  multiplex "AwsAccountRegion" {
    path   = "github.com/cloudquery/cloudquery/plugins/source/aws/client.ServiceAccountRegionMultiplexer"
    params = ["firehose"]
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
  options {
    primary_keys = ["arn"]
  }
  userDefinedColumn "tags" {
    type              = "json"
    generate_resolver = true
  }
  column "destinations" {
    skip = true
  }
  column "has_more_destinations" {
    skip = true
  }
  userDefinedColumn "arn" {
    type        = "string"
    description = "The Amazon Resource Name (ARN) of the delivery stream"
    resolver "resolveStreamArn" {
      path   = "github.com/cloudquery/cq-provider-sdk/provider/schema.PathResolver"
      params = ["DeliveryStreamARN"]
    }
  }
  column "delivery_stream_encryption_configuration" {
    rename = "encryption_config"
  }
  column "source" {
    rename = "source"
  }
  column "source_kinesis_stream_source_description" {
    rename = "_kinesis_stream"
    // skip_prefix = true
  }
  user_relation "aws" "kinesis" "open_search_destination" {
    path = "github.com/aws/aws-sdk-go-v2/service/firehose/types.AmazonopensearchserviceDestinationDescription"
    resolver "resolveTable" {
      path          = "github.com/cloudquery/cq-provider-sdk/provider/schema.PathTableResolver"
      path_resolver = true
      params        = ["Destinations.AmazonopensearchserviceDestinationDescription"]
    }
    column "s3_destination_description" {
      rename = "s3_destination"
    }
    column "s3_destination_encryption_configuration" {
      skip_prefix = true
    }
    column "s3_destination_cloud_watch_logging_options_" {
      skip_prefix = true
    }
    column "processing_configuration_processors" {
      skip = true
    }
    userDefinedColumn "processing_configuration_processors" {
      type        = "json"
      description = "Describes a data processing configuration"
      resolver "resolveProcessors" {
        path   = "github.com/cloudquery/cq-provider-sdk/provider/schema.PathResolver"
        params = ["ProcessingConfiguration.Processors"]
      }
    }
  }
  user_relation "aws" "kinesis" "elasticsearch_destination" {
    path = "github.com/aws/aws-sdk-go-v2/service/firehose/types.ElasticsearchDestinationDescription"
    resolver "resolveTable" {
      path          = "github.com/cloudquery/cq-provider-sdk/provider/schema.PathTableResolver"
      path_resolver = true
      params        = ["Destinations.ElasticsearchDestinationDescription"]
    }
    column "s3_destination_description" {
      rename = "s3_destination"
    }
    column "s3_destination_encryption_configuration" {
      skip_prefix = true
    }
    column "processing_configuration_processors" {
      skip = true
    }
    userDefinedColumn "processing_configuration_processors" {
      type        = "json"
      description = "Describes a data processing configuration"
      resolver "resolveProcessors" {
        path   = "github.com/cloudquery/cq-provider-sdk/provider/schema.PathResolver"
        params = ["ProcessingConfiguration.Processors"]
      }
    }
  }
  user_relation "aws" "kinesis" "extended_s3_destination" {
    path = "github.com/aws/aws-sdk-go-v2/service/firehose/types.ExtendedS3DestinationDescription"
    resolver "resolveTable" {
      path          = "github.com/cloudquery/cq-provider-sdk/provider/schema.PathTableResolver"
      path_resolver = true
      params        = ["Destinations.ExtendedS3DestinationDescription"]
    }
    column "data_format_conversion_configuration" {
      skip_prefix = true
    }
    column "input_format_configuration" {
      skip_prefix = true
    }
    column "deserializer_open_x_json_ser_de_convert_dots_in_json_keys_to_underscores" {
      rename = "deserializer_open_x_json_ser_de_convert_dots_to_underscores"
    }
    column "dynamic_partitioning_configuration" {
      rename = "dynamic_partitioning"
    }
    column "output_format_configuration" {
      skip_prefix = true
    }
    column "s3_backup_description" {
      rename = "s3_backup"
    }
    column "s3_backup_encryption_configuration" {
      skip_prefix = true
    }
    column "processing_configuration_processors" {
      skip = true
    }
    userDefinedColumn "processing_configuration_processors" {
      type        = "json"
      description = "Describes a data processing configuration"
      resolver "resolveProcessors" {
        path   = "github.com/cloudquery/cq-provider-sdk/provider/schema.PathResolver"
        params = ["ProcessingConfiguration.Processors"]
      }
    }
  }
  user_relation "aws" "kinesis" "http_destination" {
    path = "github.com/aws/aws-sdk-go-v2/service/firehose/types.HttpEndpointDestinationDescription"
    resolver "resolveTable" {
      path          = "github.com/cloudquery/cq-provider-sdk/provider/schema.PathTableResolver"
      path_resolver = true
      params        = ["Destinations.HttpEndpointDestinationDescription"]
    }
    column "s3_destination_description" {
      rename = "s3_destination"
    }
    column "s3_destination_encryption_configuration" {
      skip_prefix = true
    }
    column "processing_configuration_processors" {
      skip = true
    }
    userDefinedColumn "processing_configuration_processors" {
      type        = "json"
      description = "Describes a data processing configuration"
      resolver "resolveProcessors" {
        path   = "github.com/cloudquery/cq-provider-sdk/provider/schema.PathResolver"
        params = ["ProcessingConfiguration.Processors"]
      }
    }
    column "request_configuration_common_attributes" {
      // skip = true
      type = "json"
      resolver "pathResolver" {
        path   = "github.com/cloudquery/cq-provider-sdk/provider/schema.PathResolver"
        params = ["RequestConfiguration.CommonAttributes"]
      }
    }
  }
  user_relation "aws" "kinesis" "redshift_destination" {
    path = "github.com/aws/aws-sdk-go-v2/service/firehose/types.RedshiftDestinationDescription"
    resolver "resolveTable" {
      path          = "github.com/cloudquery/cq-provider-sdk/provider/schema.PathTableResolver"
      path_resolver = true
      params        = ["Destinations.RedshiftDestinationDescription"]
    }
    column "s3_destination_encryption_configuration" {
      skip_prefix = true
    }
    //s3_backup_encryption_configuration_kms_encryption_config_aws_kms_key_arn
    // column "s3_destination_encryption_configuration_kms_encryption" {
    //   skip_prefix = true
    // }
    column "s3_destination_description" {
      rename = "s3_destination"
    }
    column "s3_backup_description" {
      rename = "s3_backup"
    }

    column "s3_backup_encryption_configuration" {
      skip_prefix = true
    }
    column "processing_configuration_processors" {
      skip = true
    }
    userDefinedColumn "processing_configuration_processors" {
      type        = "json"
      description = "Describes a data processing configuration"
      resolver "resolveProcessors" {
        path   = "github.com/cloudquery/cq-provider-sdk/provider/schema.PathResolver"
        params = ["ProcessingConfiguration.Processors"]
      }
    }
  }
  user_relation "aws" "kinesis" "splunk_destination" {
    path = "github.com/aws/aws-sdk-go-v2/service/firehose/types.SplunkDestinationDescription"
    resolver "resolveTable" {
      path          = "github.com/cloudquery/cq-provider-sdk/provider/schema.PathTableResolver"
      path_resolver = true
      params        = ["Destinations.SplunkDestinationDescription"]
    }
    column "s3_destination_description" {
      rename = "s3_destination"
    }
    column "s3_destination_encryption_configuration" {
      skip_prefix = true
    }

    column "processing_configuration_processors" {
      skip = true
    }
    userDefinedColumn "processing_configuration_processors" {
      type        = "json"
      description = "Describes a data processing configuration"
      resolver "resolveProcessors" {
        path   = "github.com/cloudquery/cq-provider-sdk/provider/schema.PathResolver"
        params = ["ProcessingConfiguration.Processors"]
      }
    }
  }
}
