service          = "aws"
output_directory = "."
add_generate     = true

description_modifier "remove_read_only" {
  words = ["  This member is required."]
}

resource "aws" "lightsail" "instances" {
  path = "github.com/aws/aws-sdk-go-v2/service/lightsail/types.Instance"
  ignoreError "IgnoreAccessDenied" {
    path = "github.com/cloudquery/cq-provider-aws/client.IgnoreAccessDeniedServiceDisabled"
  }
  deleteFilter "AccountRegionFilter" {
    path = "github.com/cloudquery/cq-provider-aws/client.DeleteAccountRegionFilter"
  }
  multiplex "AwsAccountRegion" {
    path   = "github.com/cloudquery/cq-provider-aws/client.ServiceAccountRegionMultiplexer"
    params = ["lightsail"]
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
    primary_keys = ["arn"]
  }

  column "tags" {
    type = "json"
    resolver "resolveTags" {
      path = "github.com/cloudquery/cq-provider-aws/client.ResolveTags"
    }
  }

  relation "aws" "lightsail" "add_ons" {
    ignore_in_tests = true // see https://github.com/hashicorp/terraform-provider-aws/issues/23688
  }

  // disks attached to the instance
  relation "aws" "lightsail" "hardware_disks" {
    path = "github.com/aws/aws-sdk-go-v2/service/lightsail/types.Disk"

    ignore_columns_in_tests = ["gb_in_use"]

    column "tags" {
      type = "json"
      resolver "resolveTags" {
        path = "github.com/cloudquery/cq-provider-aws/client.ResolveTags"
      }
    }
    column "attachment_state" {
      skip = true
    }

    relation "aws" "lightsail" "add_ons" {
      ignore_in_tests = true // see https://github.com/hashicorp/terraform-provider-aws/issues/23688
    }
  }

  // todo maybe skip original ports column
  user_relation "aws" "lightsail" "port_states" {
    path = "github.com/aws/aws-sdk-go-v2/service/lightsail/types.InstancePortState"
  }


  userDefinedColumn "access_details" {
    type              = "json"
    generate_resolver = true
  }
}

resource "aws" "lightsail" "buckets" {
  path = "github.com/aws/aws-sdk-go-v2/service/lightsail/types.Bucket"
  ignoreError "IgnoreAccessDenied" {
    path = "github.com/cloudquery/cq-provider-aws/client.IgnoreAccessDeniedServiceDisabled"
  }
  multiplex "AwsAccountRegion" {
    path   = "github.com/cloudquery/cq-provider-aws/client.ServiceAccountRegionMultiplexer"
    params = ["lightsail"]
  }
  deleteFilter "AccountRegionFilter" {
    path = "github.com/cloudquery/cq-provider-aws/client.DeleteAccountRegionFilter"
  }

  options {
    primary_keys = [
      "arn"
    ]
  }

  ignore_columns_in_tests = [
    "access_log_config_enabled",
    "access_log_config_destination",
    "access_log_config_prefix",
    "resources_receiving_access",
    "state_message",
  ]

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


  column "tags" {
    type = "json"
    resolver "resolveTags" {
      path = "github.com/cloudquery/cq-provider-aws/client.ResolveTags"
    }
  }

  column "resources_receiving_access" {
    type = "json"
  }

  user_relation "aws" "lightsail" "access_keys" {
    path            = "github.com/aws/aws-sdk-go-v2/service/lightsail/types.AccessKey"
    ignore_in_tests = true
  }
}


resource "aws" "lightsail" "disks" {
  path = "github.com/aws/aws-sdk-go-v2/service/lightsail/types.Disk"
  ignoreError "IgnoreAccessDenied" {
    path = "github.com/cloudquery/cq-provider-aws/client.IgnoreAccessDeniedServiceDisabled"
  }
  multiplex "AwsAccountRegion" {
    path   = "github.com/cloudquery/cq-provider-aws/client.ServiceAccountRegionMultiplexer"
    params = ["lightsail"]
  }
  deleteFilter "AccountRegionFilter" {
    path = "github.com/cloudquery/cq-provider-aws/client.DeleteAccountRegionFilter"
  }

  options {
    primary_keys = [
      "arn"
    ]
  }

  ignore_columns_in_tests = ["gb_in_use"]

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

  column "tags" {
    type = "json"
    resolver "resolveTags" {
      path = "github.com/cloudquery/cq-provider-aws/client.ResolveTags"
    }
  }

  relation "aws" "lightsail" "add_ons" {
    ignore_in_tests = true
  }
  user_relation "aws" "lightsail" "disk_snapshot" {
    path            = "github.com/aws/aws-sdk-go-v2/service/lightsail/types.DiskSnapshot"
    ignore_in_tests = true

    column "tags" {
      type = "json"
      resolver "resolveTags" {
        path = "github.com/cloudquery/cq-provider-aws/client.ResolveTags"
      }
    }
  }
}


resource "aws" "lightsail" "alarms" {
  path            = "github.com/aws/aws-sdk-go-v2/service/lightsail/types.Alarm"
  ignore_in_tests = true

  ignoreError "IgnoreAccessDenied" {
    path = "github.com/cloudquery/cq-provider-aws/client.IgnoreAccessDeniedServiceDisabled"
  }
  multiplex "AwsAccountRegion" {
    path   = "github.com/cloudquery/cq-provider-aws/client.ServiceAccountRegionMultiplexer"
    params = ["lightsail"]
  }
  deleteFilter "AccountRegionFilter" {
    path = "github.com/cloudquery/cq-provider-aws/client.DeleteAccountRegionFilter"
  }

  options {
    primary_keys = [
      "arn"
    ]
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

  column "location" {
    skip_prefix = true
  }
  column "region_name" {
    skip = true
  }
  column "monitored_resource_info_name" {
    rename = "monitored_resource_name"
  }
  column "monitored_resource_info_resource_type" {
    rename = "monitored_resource_resource_type"
  }
}


resource "aws" "lightsail" "certificates" {
  path = "github.com/aws/aws-sdk-go-v2/service/lightsail/types.Certificate"
  ignoreError "IgnoreAccessDenied" {
    path = "github.com/cloudquery/cq-provider-aws/client.IgnoreAccessDeniedServiceDisabled"
  }
  multiplex "AwsAccountRegion" {
    path   = "github.com/cloudquery/cq-provider-aws/client.ServiceAccountRegionMultiplexer"
    params = ["lightsail"]
  }
  deleteFilter "AccountRegionFilter" {
    path = "github.com/cloudquery/cq-provider-aws/client.DeleteAccountRegionFilter"
  }

  options {
    primary_keys = [
      "arn"
    ]
  }

  ignore_columns_in_tests = [
    "issued_at",
    "not_after",
    "not_before",
    "renewal_summary_reason",
    "renewal_summary_updated_at",
    "revocation_reason",
    "revoked_at",
    "serial_number",
    "renewal_summary_renewal_status_reason",
  ]

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


  column "tags" {
    type = "json"
    resolver "resolveTags" {
      path = "github.com/cloudquery/cq-provider-aws/client.ResolveTags"
    }
  }

  column "renewal_summary_renewal_status" {
    rename = "renewal_summary_status"
  }

  column "renewal_summary_renewal_status_reason" {
    rename = "renewal_summary_reason"
  }

  relation "aws" "lightsail" "domain_validation_records" {
    ignore_in_tests = true

    column "resource_record" {
      skip_prefix = true
    }
  }

  relation "aws" "lightsail" "renewal_summary_domain_validation_records" {
    ignore_in_tests = true

    column "resource_record" {
      skip_prefix = true
    }
  }
}

resource "aws" "lightsail" "static_ips" {
  path = "github.com/aws/aws-sdk-go-v2/service/lightsail/types.StaticIp"
  ignoreError "IgnoreAccessDenied" {
    path = "github.com/cloudquery/cq-provider-aws/client.IgnoreAccessDeniedServiceDisabled"
  }
  multiplex "AwsAccountRegion" {
    path   = "github.com/cloudquery/cq-provider-aws/client.ServiceAccountRegionMultiplexer"
    params = ["lightsail"]
  }
  deleteFilter "AccountRegionFilter" {
    path = "github.com/cloudquery/cq-provider-aws/client.DeleteAccountRegionFilter"
  }
  options {
    primary_keys = [
      "arn"
    ]
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

  column "location" {
    skip_prefix = true
  }

  column "region_name" {
    skip = true
  }
}


resource "aws" "lightsail" "database_snapshots" {
  path = "github.com/aws/aws-sdk-go-v2/service/lightsail/types.RelationalDatabaseSnapshot"
  ignoreError "IgnoreAccessDenied" {
    path = "github.com/cloudquery/cq-provider-aws/client.IgnoreAccessDeniedServiceDisabled"
  }
  multiplex "AwsAccountRegion" {
    path   = "github.com/cloudquery/cq-provider-aws/client.ServiceAccountRegionMultiplexer"
    params = ["lightsail"]
  }
  deleteFilter "AccountRegionFilter" {
    path = "github.com/cloudquery/cq-provider-aws/client.DeleteAccountRegionFilter"
  }

  options {
    primary_keys = [
      "arn"
    ]
  }

  ignore_in_tests = true

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
  column "tags" {
    type = "json"
    resolver "resolveTags" {
      path = "github.com/cloudquery/cq-provider-aws/client.ResolveTags"
    }
  }
  column "location" {
    skip_prefix = true
  }
  column "region_name" {
    skip = true
  }
}


resource "aws" "lightsail" "load_balancers" {
  path = "github.com/aws/aws-sdk-go-v2/service/lightsail/types.LoadBalancer"
  ignoreError "IgnoreAccessDenied" {
    path = "github.com/cloudquery/cq-provider-aws/client.IgnoreAccessDeniedServiceDisabled"
  }
  multiplex "AwsAccountRegion" {
    path   = "github.com/cloudquery/cq-provider-aws/client.ServiceAccountRegionMultiplexer"
    params = ["lightsail"]
  }
  deleteFilter "AccountRegionFilter" {
    path = "github.com/cloudquery/cq-provider-aws/client.DeleteAccountRegionFilter"
  }

  options {
    primary_keys = [
      "arn"
    ]
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
  column "tags" {
    type = "json"
    resolver "resolveTags" {
      path = "github.com/cloudquery/cq-provider-aws/client.ResolveTags"
    }
  }

  column "public_ports" {
    generate_resolver = true
  }

  column "location" {
    skip_prefix = true
  }

  column "region_name" {
    skip = true
  }

  column "resource_type" {
    description = "Type of the lightsail resource"
  }


  relation "aws" "lightsail" "tls_certificate_summaries" {
    ignore_in_tests = true
  }

  user_relation "aws" "lightsail" "tls_certificates" {
    path            = "github.com/aws/aws-sdk-go-v2/service/lightsail/types.LoadBalancerTlsCertificate"
    ignore_in_tests = true

    column "tags" {
      type = "json"
      resolver "resolveTags" {
        path = "github.com/cloudquery/cq-provider-aws/client.ResolveTags"
      }
    }

    column "renewal_summary_domain_validation_options" {
      type = "json"
    }

    column "location" {
      skip_prefix = true
    }


    column "domain_validation_records" {
      type = "json"
    }
  }
}


resource "aws" "lightsail" "databases" {
  path = "github.com/aws/aws-sdk-go-v2/service/lightsail/types.RelationalDatabase"
  ignoreError "IgnoreAccessDenied" {
    path = "github.com/cloudquery/cq-provider-aws/client.IgnoreAccessDeniedServiceDisabled"
  }
  multiplex "AwsAccountRegion" {
    path   = "github.com/cloudquery/cq-provider-aws/client.ServiceAccountRegionMultiplexer"
    params = ["lightsail"]
  }
  deleteFilter "AccountRegionFilter" {
    path = "github.com/cloudquery/cq-provider-aws/client.DeleteAccountRegionFilter"
  }

  options {
    primary_keys = [
      "arn"
    ]
  }

  ignore_columns_in_tests = [
    "pending_modified_values",
    "secondary_availability_zone",
  ]

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


  column "tags" {
    type = "json"
    resolver "resolveTags" {
      path = "github.com/cloudquery/cq-provider-aws/client.ResolveTags"
    }
  }
  column "pending_modified_values" {
    type = "json"
  }

  column "location" {
    skip_prefix = true
  }

  column "region_name" {
    skip = true
  }

  relation "aws" "lightsail" "pending_maintenance_actions" {
    ignore_columns_in_tests = ["current_apply_date"]
  }

  user_relation "aws" "lightsail" "parameters" {
    path = "github.com/aws/aws-sdk-go-v2/service/lightsail/types.RelationalDatabaseParameter"

    column "parameter_name" {
      rename = "name"
    }

    column "parameter_value" {
      rename = "value"
    }
  }

  user_relation "aws" "lightsail" "events" {
    path = "github.com/aws/aws-sdk-go-v2/service/lightsail/types.RelationalDatabaseEvent"
  }

  user_relation "aws" "lightsail" "log_events" {
    path            = "github.com/cloudquery/cq-provider-aws/resources/services/lightsail.LogEventWrapper"
    ignore_in_tests = true

    column "log_event" {
      skip_prefix = true
    }
  }
}


resource "aws" "lightsail" "instance_snapshots" {
  path = "github.com/aws/aws-sdk-go-v2/service/lightsail/types.InstanceSnapshot"
  ignoreError "IgnoreAccessDenied" {
    path = "github.com/cloudquery/cq-provider-aws/client.IgnoreAccessDeniedServiceDisabled"
  }
  multiplex "AwsAccountRegion" {
    path   = "github.com/cloudquery/cq-provider-aws/client.ServiceAccountRegionMultiplexer"
    params = ["lightsail"]
  }
  deleteFilter "AccountRegionFilter" {
    path = "github.com/cloudquery/cq-provider-aws/client.DeleteAccountRegionFilter"
  }

  options {
    primary_keys = [
      "arn"
    ]
  }

  ignore_in_tests = true

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

  column "location" {
    skip_prefix = true
  }
  column "region_name" {
    skip = true
  }

  column "tags" {
    type              = "json"
    generate_resolver = true
  }

  relation "aws" "lightsail" "from_attached_disks" {
    column "tags" {
      type = "json"
      resolver "resolveTags" {
        path = "github.com/cloudquery/cq-provider-aws/client.ResolveTags"
      }
    }
  }
}


resource "aws" "lightsail" "distributions" {
  path = "github.com/cloudquery/cq-provider-aws/resources/services/lightsail.DistributionWrapper"

  ignoreError "IgnoreAccessDenied" {
    path = "github.com/cloudquery/cq-provider-aws/client.IgnoreAccessDeniedServiceDisabled"
  }
  multiplex "AwsAccountRegion" {
    path = "github.com/cloudquery/cq-provider-aws/client.AccountMultiplex"
  }
  deleteFilter "AccountRegionFilter" {
    path = "github.com/cloudquery/cq-provider-aws/client.DeleteAccountFilter"
  }

  ignore_in_tests = true

  options {
    primary_keys = [
      "arn"
    ]
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

  column "get_distribution_latest_cache_reset_output_create_time" {
    rename = "cache_reset_create_time"
  }

  column "get_distribution_latest_cache_reset_output_status" {
    rename = "cache_reset_status"
  }

  column "lightsail_distribution" {
    skip_prefix = true
  }

  column "cache_behavior_settings" {
    type = "json"
  }

  column "cache_behaviors" {
    type = "json"
  }

  column "location" {
    skip_prefix = true
  }
  column "region_name" {
    skip = true
  }

  column "tags" {
    type = "json"
    resolver "resolveTags" {
      path = "github.com/cloudquery/cq-provider-aws/client.ResolveTags"
    }
  }
}