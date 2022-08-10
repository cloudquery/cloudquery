provider "aws" {

  resource "*" {
    ignore_identifiers = [ ]
    ignore_attributes = [ "unknown_fields" ]

    iac {
      terraform {
        # map of attributes from cloud provider to iac provider
        attribute_map = [
          "tags=tags_all"
        ]
      }
    }
  }

  resource "accessanalyzer.analyzers" {
    identifiers = [ "name" ]
    ignore_attributes = [ "last_resource_analyzed", "last_resource_analyzed_at", "status" ]

    iac {
      terraform {
        type = "aws_accessanalyzer_analyzer"

        attribute_map = [
          "name=id"
        ]
      }
    }
  }

  resource "acm.certificates" {
    ignore_attributes = [
      "created_at",
      "extended_key_usages",
      "failure_reason",
      "imported_at",
      "in_use_by",
      "issued_at",
      "issuer",
      "key_algorithm",
      "key_usages",
      "not_after",
      "not_before",
      "renewal_eligibility",
      "renewal_summary_domain_validation_options",
      "renewal_summary_status",
      "renewal_summary_updated_at",
      "renewal_summary_failure_reason",
      "revocation_reason",
      "revoked_at",
      "serial",
      "signature_algorithm",
      "subject",
      "type"
    ]

    iac {
      terraform {
        type = "aws_acm_certificate"
        identifiers = [ "arn" ]

        attribute_map = [
          "certificate_transparency_logging_preference=options.0.certificate_transparency_logging_preference"
        ]
      }
    }
  }

  resource "apigateway.api_keys" {
    iac {
      terraform {
        type = "aws_api_gateway_api_key"
      }
    }
  }

  resource "apigateway.client_certificates" {
    iac {
      terraform {
        type = "aws_api_gateway_client_certificate"
      }
    }
  }

  resource "apigateway.domain_names" {
    iac {
      terraform {
        type = "aws_api_gateway_domain_name"
        identifiers = [ "domain_name" ]
      }
    }
  }

  resource "apigateway.rest_apis" {
    iac {
      terraform {
        type = "aws_api_gateway_rest_api"
      }
    }
  }

  resource "aws_apigateway_rest_api_authorizers" {
    identifiers = [ "rest_api_id", "id" ]
    ignore_attributes = [ "arn", "auth_type" ]

    iac {
      terraform {
        type = "aws_api_gateway_authorizer"
        identifiers = [ "rest_api_id", "id" ]
      }
    }
  }

  resource "aws_apigateway_rest_api_deployments" {
    identifiers = [ "rest_api_id", "id" ]
    ignore_attributes = [ "arn", "created_date" ]

    iac {
      terraform {
        type = "aws_api_gateway_deployment"
        identifiers = [ "rest_api_id", "id" ]
      }
    }
  }

  resource "aws_apigateway_rest_api_documentation_parts" {
    identifiers = [ sql("CONCAT(c.rest_api_id, '/', c.id)") ]

    iac {
      terraform {
        type = "aws_api_gateway_documentation_part"
      }
    }
  }

  resource "aws_apigateway_rest_api_documentation_versions" {
    identifiers = [ sql("CONCAT(c.rest_api_id, '/', c.version)") ]

    iac {
      terraform {
        type = "aws_api_gateway_documentation_version"
      }
    }
  }

  resource "aws_apigateway_rest_api_models" {
    identifiers = [ "rest_api_id", "id" ]
    ignore_attributes = [ "arn", "model_template" ]

    iac {
      terraform {
        type = "aws_api_gateway_model"
        identifiers = [ "rest_api_id", "id" ]
      }
    }
  }

  resource "aws_apigateway_rest_api_request_validators" {
    identifiers = [ "rest_api_id", "id" ]
    ignore_attributes = [ "arn" ]

    iac {
      terraform {
        type = "aws_api_gateway_request_validator"
        identifiers = [ "rest_api_id", "id" ]
      }
    }
  }

  resource "aws_apigateway_rest_api_resources" {
    identifiers = [ "rest_api_id", "parent_id", "id" ]
    ignore_attributes = [ "arn", "resource_methods" ]

    iac {
      terraform {
        type = "aws_api_gateway_resource"
        identifiers = [ "rest_api_id", "parent_id", "id" ]
      }
    }
  }

  resource "aws_apigateway_rest_api_stages" {
    identifiers = [ "arn" ]
    ignore_attributes = [ "cache_cluster_status", "canary_settings_percent_traffic", "canary_settings_use_stage_cache", "canary_settings_deployment_id", "canary_settings_stage_variable_overrides", "created_date", "last_updated_date", "method_settings" ]

    iac {
      terraform {
        type = "aws_api_gateway_stage"
        identifiers = [ "arn" ]
        attribute_map = [
          "tracing_enabled=xray_tracing_enabled"
        ]
      }
    }
  }

  resource "apigateway.usage_plans" {
    iac {
      terraform {
        type = "aws_api_gateway_usage_plan"
      }
    }
  }

  resource "aws_apigateway_usage_plan_api_stages" {
    identifiers = [ "usage_plan_id", "api_id", "stage" ]
    iac {
      terraform {
        type = "aws_api_gateway_usage_plan"
        path = "api_stages"
        identifiers = [ "root.id", "api_id", "stage" ]
        attribute_map = [
          "usage_plan_id=root.id"
        ]
      }
    }
  }

  resource "aws_apigateway_usage_plan_keys" {
    identifiers = [ "usage_plan_id", "id" ]
    ignore_attributes = [ "arn" ]
    iac {
      terraform {
        type = "aws_api_gateway_usage_plan_key"
        identifiers = [ "usage_plan_id", "id" ]
        attribute_map = [
          "type=key_type"
        ]
      }
    }
  }

  resource "apigateway.vpc_links" {
    iac {
      terraform {
        type = "aws_api_gateway_vpc_link"
      }
    }
  }

  resource "apigatewayv2.apis" {
    identifiers = [ "arn" ]
    ignore_attributes = [ "created_date", "api_gateway_managed", "disable_schema_validation" ]
    sets = [ "cors_configuration_allow_headers", "cors_configuration_allow_methods", "cors_configuration_allow_origins", "cors_configuration_expose_headers" ]
    iac {
      terraform {
        type = "aws_apigatewayv2_api"
        identifiers = [ "arn" ]
        attribute_map = [
          "cors_configuration_allow_credentials=cors_configuration.0.allow_credentials",
          "cors_configuration_allow_headers=cors_configuration.0.allow_headers",
          "cors_configuration_allow_methods=cors_configuration.0.allow_methods",
          "cors_configuration_allow_origins=cors_configuration.0.allow_origins",
          "cors_configuration_allow_credentials=cors_configuration.0.allow_credentials",
          "cors_configuration_expose_headers=cors_configuration.0.expose_headers",
          "cors_configuration_max_age=cors_configuration.0.max_age"
        ]
      }
    }
  }

  resource "aws_apigatewayv2_api_authorizers" {
    identifiers = [ "api_id", "authorizer_id" ]
    ignore_attributes = [ "arn" ]

    iac {
      terraform {
        type = "aws_apigatewayv2_authorizer"
        identifiers = [ "api_id", "id" ]
        attribute_map = [
          "authorizer_id=id",
          "identity_source=identity_sources"
        ]
      }
    }
  }

  resource "aws_apigatewayv2_api_deployments" {
    identifiers = [ "api_id", "deployment_id" ]
    ignore_attributes = [ "arn", "created_date", "deployment_status" ]

    iac {
      terraform {
        type = "aws_apigatewayv2_deployment"
        identifiers = [ "api_id", "id" ]
        attribute_map = [
          "deployment_id=id"
        ]
      }
    }
  }

  resource "aws_apigatewayv2_api_integrations" {
    identifiers = [ "api_id", "integration_id" ]
    ignore_attributes = [ "arn", "api_gateway_managed", "deployment_status" ]

    iac {
      terraform {
        type = "aws_apigatewayv2_integration"
        identifiers = [ "api_id", "id" ]
        attribute_map = [
          "integration_id=id",
          "timeout_in_millis=timeout_milliseconds"
        ]
      }
    }
  }

  resource "aws_apigatewayv2_api_integration_responses" {
    identifiers = [ "parent.api_id", "c.integration_id", "c.integration_response_id" ]
    ignore_attributes = [ "arn" ]

    iac {
      terraform {
        type = "aws_apigatewayv2_integration_response"
        identifiers = [ "api_id", "integration_id", "id" ]
        attribute_map = [
          "integration_response_id=id"
        ]
      }
    }
  }

  resource "aws_apigatewayv2_api_models" {
    identifiers = [ "api_id", "model_id" ]
    ignore_attributes = [ "arn", "model_template" ]

    iac {
      terraform {
        type = "aws_apigatewayv2_model"
        identifiers = [ "api_id", "id" ]
        attribute_map = [
          "model_id=id"
        ]
      }
    }
  }

  resource "aws_apigatewayv2_api_routes" {
    identifiers = [ "api_id", "route_id" ]
    ignore_attributes = [ "arn", "api_gateway_managed" ]

    iac {
      terraform {
        type = "aws_apigatewayv2_route"
        identifiers = [ "api_id", "id" ]
        attribute_map = [
          "route_id=id"
        ]
      }
    }
  }

  resource "aws_apigatewayv2_api_route_responses" {
    identifiers = [ "parent.api_id", "c.route_id", "c.route_response_id" ]
    ignore_attributes = [ "arn" ]

    iac {
      terraform {
        type = "aws_apigatewayv2_route_response"
        identifiers = [ "api_id", "route_id", "id" ]
        attribute_map = [
          "route_response_id=id"
        ]
      }
    }
  }

  resource "aws_apigatewayv2_api_stages" {
    identifiers = [ "arn" ]
    ignore_attributes = [ "api_gateway_managed", "created_date", "last_updated_date", "last_deployment_status_message" ]

    iac {
      terraform {
        type = "aws_apigatewayv2_stage"
        identifiers = [ "arn" ]
        attribute_map = [
          "stage_name=id",
          "route_settings_data_trace_enabled=route_settings.0|@getbool:data_trace_enabled",
          "route_settings_detailed_metrics_enabled=route_settings.0|@getbool:detailed_metrics_enabled",
          "route_settings_logging_level=route_settings.0.logging_level",
          "route_settings_throttling_burst_limit=route_settings.0.throttling_burst_limit",
          "route_settings_throttling_rate_limit=route_settings.0.throttling_rate_limit"
        ]
      }
    }
  }

  resource "apigatewayv2.domain_names" {
    identifiers = [ "arn" ]
    ignore_attributes = [ "mutual_tls_authentication_truststore_warnings" ]

    iac {
      terraform {
        type = "aws_apigatewayv2_domain_name"
        identifiers = [ "arn" ]

        attribute_map = [
          "mutual_tls_authentication_truststore_uri=mutual_tls_authentication.0.truststore_uri",
          "mutual_tls_authentication_truststore_version=mutual_tls_authentication.0.truststore_version"
        ]
      }
    }
  }

  resource "apigatewayv2.vpc_links" {
    ignore_attributes = [ "created_date", "vpc_link_status", "vpc_link_status_message", "vpc_link_version", "arn" ]
    sets = [ "security_group_ids", "subnet_ids" ]
    iac {
      terraform {
        type = "aws_apigatewayv2_vpc_link"
      }
    }
  }

  resource "applicationautoscaling.policies" {
    iac {
      terraform {
        type = "aws_appautoscaling_policy"
        identifiers = [ "arn" ]
      }
    }
  }

  resource "autoscaling.groups" {
    identifiers = [ "arn" ]
    ignore_attributes = [ "created_time", "load_balancers", "notifications_configurations", "metrics", "status" ]
    sets = [ "availability_zones", "load_balancer_target_groups", "suspended_processes", "vpc_zone_identifier" ]

    iac {
      terraform {
        type = "aws_autoscaling_group"
        identifiers = [ "arn" ]

        attribute_map = [
          "load_balancer_target_groups|@keys=target_group_arns", # object keys to array
          "suspended_processes|@keys=suspended_processes", # object keys to array
          "vpc_zone_identifier|@split:,=vpc_zone_identifier", # comma-separated string to array
          "load_balancer_names=load_balancers",
          "launch_configuration_name=launch_configuration",
          "launch_template_id=launch_template.0.id",
          "launch_template_name=launch_template.0.name",
          "launch_template_version=launch_template.0.version",
          "new_instances_protected_from_scale_in=protect_from_scale_in",
          "mixed_instances_policy=mixed_instances_policy.0"
        ]
      }
    }
  }

  resource "autoscaling.launch_configurations" {
    identifiers = [ "launch_configuration_name" ]
    ignore_attributes = [ "created_time" ]

    iac {
      terraform {
        type = "aws_launch_configuration"
        attribute_map = [
          "launch_configuration_name=name"
        ]
      }
    }
  }

  resource "aws_autoscaling_launch_configuration_block_device_mappings" {
    identifiers = [ "parent.launch_configuration_name", "device_name" ]

    iac {
      terraform {
        type = "aws_launch_configuration"
        path = "ebs_block_device"
        identifiers = [ "root.id", "device_name" ]
      }
    }
  }

  resource "autoscaling.scheduled_actions" {
    ignore_attributes = [ "time" ]
    iac {
      terraform {
        type = "aws_autoscaling_schedule"
        identifiers = [ "arn" ]
        attribute_map = [
          "auto_scaling_group_name=autoscaling_group_name",
          "name=scheduled_action_name"
        ]
      }
    }
  }

  resource "backup.plans" {
    ignore_attributes = [ "creation_date", "creator_request_id", "last_execution_date" ]
    iac {
      terraform {
        type = "aws_backup_plan"
        identifiers = [ "arn" ]
        attribute_map = [
          "advanced_backup_settings=advanced_backup_setting|0",
          "version_id=version"
        ]
      }
    }
  }

  resource "aws_backup_plan_rules" {
    identifiers = [ "parent.arn", "c.name" ]
    ignore_attributes = [ "id", "name" ]
    iac {
      terraform {
        type = "aws_backup_plan"
        path = "rule"
        identifiers = [ "root.arn", "rule_name" ]
        attribute_map = [
          "completion_window_minutes=completion_window",
          "schedule_expression=schedule",
          "start_window_minutes=start_window",
          "target_backup_vault_name=target_vault_name"
        ]
      }
    }
  }

  resource "aws_backup_plan_selections" {
    identifiers = [ "parent.id", "c.selection_id" ]
    ignore_attributes = [ "creation_date", "creator_request_id", "conditions", "list_of_tags" ]
#    sets = [ "conditions", "list_of_tags" ] # TODO CamelCase vs snake_case in keys of map
    iac {
      terraform {
        type = "aws_backup_selection"
        identifiers = [ "plan_id", "id" ]
        attribute_map = [
          "selection_id=id",
          "selection_name=name"
#          "conditions=condition",
#          "list_of_tags=selection_tag"
        ]
      }
    }
  }

  resource "backup.vaults" {
    ignore_attributes = [ "creation_date", "creator_request_id", "locked", "notification_events", "notification_sns_topic_arn", "lock_date", "max_retention_days", "min_retention_days", "access_policy" ]
    iac {
      terraform {
        type = "aws_backup_vault"
        identifiers = [ "arn" ]
        attribute_map = [
          "encryption_key_arn=kms_key_arn",
          "number_of_recovery_points=recovery_points"
        ]
      }
    }
  }

  resource "backup.vaults#notif" {
    identifiers = [ "arn" ]
    ignore_attributes = [ "creation_date", "creator_request_id", "locked", "encryption_key_arn", "number_of_recovery_points", "lock_date", "max_retention_days", "min_retention_days", "tags", "access_policy" ]
    sets = [ "notification_events" ]
    iac {
      terraform {
        type = "aws_backup_vault_notifications"
        identifiers = [ "backup_vault_arn" ]
        attribute_map = [
          "arn=backup_vault_arn",
          "name=backup_vault_name",
          "notification_events=backup_vault_events",
          "notification_sns_topic_arn=sns_topic_arn"
        ]
      }
    }
  }

  resource "cloudformation.stacks" {
    identifiers = [ "id" ]
    ignore_attributes = [ "status", "stack_drift_status" ]
    iac {
      terraform {
        type = "aws_cloudformation_stack"
        attribute_map = [
          "arn=id",
          "stack=name"
        ]
      }
    }
  }

  resource "cloudfront.cache_policies" {
    iac {
      terraform {
        type = "aws_cloudfront_cache_policy"
      }
    }
  }

  resource "cloudfront.distributions" {
    identifiers = [ "id" ]

    iac {
      terraform {
        type = "aws_cloudfront_distribution"
      }
    }
  }

  resource "aws_cloudfront_distribution_cache_behaviors" {
    identifiers = [ "parent.id", "path_pattern", "target_origin_id", "viewer_protocol_policy" ]
    sets = [ "allowed_methods", "cached_methods" ]

    iac {
      terraform {
        type = "aws_cloudfront_distribution"
        path = "ordered_cache_behavior"
        identifiers = [ "root.id", "path_pattern", "target_origin_id", "viewer_protocol_policy" ]
      }
    }
  }

  resource "aws_cloudfront_distribution_custom_error_responses" {
    identifiers = [ "parent.id", "error_code", "response_code", "response_page_path" ]

    iac {
      terraform {
        type = "aws_cloudfront_distribution"
        path = "custom_error_response"
        identifiers = [ "root.id", "error_code", "response_code", "response_page_path" ]
      }
    }
  }

  resource "aws_cloudfront_distribution_origins" {
    identifiers = [ "domain_name", "id", "s3_origin_config_origin_access_identity" ]
    filters = [
      "c.s3_origin_config_origin_access_identity IS NOT NULL AND c.s3_origin_config_origin_access_identity!=''"
    ]

    iac {
      terraform {
        type = "aws_cloudfront_distribution"
        path = "origin"
        identifiers = [ "domain_name", "origin_id", "s3_origin_config.0.origin_access_identity" ]
        attribute_map = [
          "id=origin_id",
          "s3_origin_config_origin_access_identity=s3_origin_config.0.origin_access_identity",
          "custom_origin_config_http_port=s3_origin_config.0.http_port",
          "custom_origin_config_https_port=s3_origin_config.0.https_port",
          "custom_origin_config_protocol_policy=s3_origin_config.0.origin_protocol_policy",
          "custom_origin_config_ssl_protocols=s3_origin_config.0.origin_ssl_protocols",
          "custom_origin_config_keepalive_timeout=s3_origin_config.0.origin_keepalive_timeout",
          "custom_origin_config_read_timeout=s3_origin_config.0.origin_read_timeout",
          "origin_shield_enabled=origin_shield.0|@getbool:enabled",
          "origin_shield_region=origin_shield.0.origin_shield_region",
          "custom_headers=custom_header"
        ]
      }
    }
  }

  # TODO: aws_cloudfront_distribution_origin_groups (tf res type: "aws_cloudfront_distribution".attributes->"origin_group")

  resource "cloudtrail.trails" {
    identifiers = [ "name" ]

    iac {
      terraform {
        type = "aws_cloudtrail"
      }
    }
  }

  resource "aws_cloudtrail_trail_event_selectors" {
    identifiers = [ "parent.name", sql("include_management_events::varchar"), "read_write_type" ]

    iac {
      terraform {
        type = "aws_cloudtrail"
        path = "event_selector"
        identifiers = [ "root.id", "include_management_events", "read_write_type" ]
      }
    }
  }

  resource "cloudwatch.alarms" {
    identifiers = [ "name" ]

    iac {
      terraform {
        type = "aws_cloudwatch_metric_alarm"
      }
    }
  }

  resource "aws_cloudwatch_alarm_metrics" {
    identifiers = [ "parent.name", "id" ]

    iac {
      terraform {
        type = "aws_cloudwatch_metric_alarm"
        path = "metric_query"
        identifiers = [ "root.id", "id" ]
      }
    }
  }

  resource "cloudwatchlogs.filters" {
    identifiers = [ "name", "log_group_name" ]

    iac {
      terraform {
        identifiers = [ "id", "log_group_name" ]
        type = "aws_cloudwatch_log_metric_filter"
      }
    }
  }

  resource "aws_cloudwatchlogs_filter_metric_transformations" {
    identifiers = [ "parent.name", "metric_namespace", "metric_name" ]
    ignore_attributes = [ "default_value" ]

    iac {
      terraform {
        type = "aws_cloudwatch_log_metric_filter"
        path = "metric_transformation"
        identifiers = [ "root.id", "namespace", "name" ]
        attribute_map = [
          "metric_namespace=namespace",
          "metric_name=name",
          "metric_value=value",
        ]
      }
    }
  }

  resource "codebuild.projects" {
    identifiers = [ "arn" ]
    ignore_attributes = [ "created", "last_modified" ]

    iac {
      terraform {
        type = "aws_codebuild_project"
        identifiers = [ "arn" ]
        attribute_map = [
          "artifacts_type=artifacts.0.type",
          "artifacts_encryption_disabled=artifacts.0.encryption_disabled",
          "artifacts_override_artifact_name=artifacts.0.override_artifact_name",
          "artifacts_packaging=artifacts.0.packaging",
          "artifacts_name=artifacts.0.name",
          "cache_type=cache.0.type",
          "environment_compute_type=environment.0.compute_type",
          "environment_image=environment.0.image",
          "environment_type=environment.0.type",
          "environment_privileged_mode=environment.0.privileged_mode",
          "environment_image_pull_credentials_type=environment.0.image_pull_credentials_type",
          "logs_config_cloud_watch_logs_status=logs_config.0.cloudwatch_logs.0.status",
          "logs_config_s3_logs_status=logs_config.0.s3_logs.0.status",
          "logs_config_s3_logs_encryption_disabled=logs_config.0.s3_logs.0.encryption_disabled",
          "queued_timeout_in_minutes=queued_timeout",
          "source_type=source.0.type",
          "source_insecure_ssl=source.0.insecure_ssl",
          "timeout_in_minutes=build_timeout"
        ]
      }
    }
  }

  resource "codepipeline.pipelines" {
    identifiers = [ "arn" ]
    ignore_attributes = [ "created", "updated", "artifact_stores", "version" ]

    iac {
      terraform {
        type = "aws_codepipeline"
        identifiers = [ "arn" ]
        attribute_map = [
          "artifact_store_location=artifact_store.0.location",
          "artifact_store_type=artifact_store.0.type",
          "artifact_store_encryption_key_id=artifact_store.0.encryption_key.0.id",
          "artifact_store_encryption_key_type=artifact_store.0.encryption_key.0.type"
        ]
      }
    }
  }

  resource "aws_codepipeline_pipeline_stages" {
    identifiers = [ "parent.arn", "c.name" ]
    ignore_attributes = [ "blockers", "stage_order" ]

    iac {
      terraform {
        type = "aws_codepipeline"
        identifiers = [ "root.arn", "name" ]
        path = "stage"
      }
    }
  }

  resource "aws_codepipeline_pipeline_stage_actions" {
    identifiers = [ "parent1.arn", "parent.name", "c.name" ]
    ignore_attributes = [ "run_order" ]

    iac {
      terraform {
        type = "aws_codepipeline"
        path = "stage"
        identifiers = [ "root.arn", "name", "action.#.name" ]
        attribute_map = [
          "category=action.#.category|0",
          "configuration=action.#.configuration|0",
          "provider=action.#.provider|0",
          "version=action.#.version|0",
          "owner=action.#.owner|0",
          "region=action.#.region|0",
          "role_arn=action.#.role_arn|0",
          "namespace=action.#.namespace|0",
          "input_artifacts=action.#.input_artifacts|0",
          "output_artifacts=action.#.output_artifacts|0"
        ]
      }
    }
  }

  resource "codepipeline.webhooks" {
    identifiers = [ "arn" ]
    ignore_attributes = [ "last_triggered" ]

    iac {
      terraform {
        type = "aws_codepipeline_webhook"
        identifiers = [ "arn" ]
      }
    }
  }

  resource "aws_codepipeline_webhook_filters" {
    identifiers = [ "parent.arn", "c.json_path", "c.match_equals" ]

    iac {
      terraform {
        type = "aws_codepipeline_webhook"
        identifiers = [ "root.arn", "json_path", "match_equals" ]
        path = "filter"
      }
    }
  }

  resource "cognito.identity_pools" {
    iac {
      terraform {
        type = "aws_cognito_identity_pool"
      }
    }
  }

  resource "cognito.user_pools" {
    iac {
      terraform {
        type = "aws_cognito_user_pool"
      }
    }
  }

  resource "config.configuration_recorders" {
    identifiers = [ "name" ]

    iac {
      terraform {
        type = "aws_config_configuration_recorder"
      }
    }
  }

  resource "config.conformance_packs" {
    identifiers = [ "conformance_pack_name" ]

    iac {
      terraform {
        type = "aws_config_conformance_pack"
      }
    }
  }

  resource "dax.clusters" {
    iac {
      terraform {
        type = "aws_dax_cluster"
        identifiers = [ "arn" ]
      }
    }
  }

  resource "directconnect.connections" {
    iac {
      terraform {
        type = "aws_dx_connection"
      }
    }
  }

  resource "directconnect.gateways" {
    iac {
      terraform {
        type = "aws_dx_gateway"
      }
    }
  }

  resource "aws_directconnect_gateway_associations" {
    identifiers = [ sql("CONCAT('ga-', c.gateway_id, c.associated_gateway_id)") ]
    iac {
      terraform {
        type = "aws_dx_gateway_association"
      }
    }
  }

  resource "directconnect.lags" {
    iac {
      terraform {
        type = "aws_dx_lag"
      }
    }
  }

  resource "directconnect.virtual_interfaces" {
    iac {
      terraform {
        type = "aws_dx_public_virtual_interface"
      }
    }
  }

  resource "dms.replication_instances" {
    identifiers = [ "arn" ]
    ignore_attributes = [ "instance_create_time" ]

    iac {
      terraform {
        type = "aws_dms_replication_instance"
        identifiers = [ "replication_instance_arn" ]
      }
    }
  }

  resource "dynamodb.tables" {
    identifiers = [ "name", "region" ]
    ignore_attributes = [ "creation_date_time" ]
    filters = [
      "c.global_table_version IS NULL OR c.global_table_version != '2017.11.29'"
    ]

    iac {
      terraform {
        type = "aws_dynamodb_table"
        identifiers = [ "name", "region" ]
      }
    }
  }

  resource "dynamodb.tables#globalv1" {
    identifiers = [ "name" ]
    ignore_attributes = [ "creation_date_time" ]
    filters = [
      "c.global_table_version='2017.11.29'"
    ]

    iac {
      terraform {
        type = "aws_dynamodb_global_table"
        identifiers = [ "replication_instance_arn" ]
      }
    }
  }

  resource "ec2.customer_gateways" {
    iac {
      terraform {
        type = "aws_customer_gateway"
      }
    }
  }

  resource "ec2.ebs_snapshots" {
    ignore_attributes = [ "start_time" ]

    iac {
      terraform {
        type = "aws_ebs_snapshot"
      }
    }
  }

  resource "ec2.ebs_volumes" {
    iac {
      terraform {
        type = "aws_ebs_volume"
      }
    }
  }

  resource "aws_ec2_ebs_volume_attachments" {
    identifiers = [ "instance_id", "volume_id", "device" ]
    ignore_attributes = [ "attach_time", "state" ]
    iac {
      terraform {
        type = "aws_instance"
        path = "root_block_device"
        identifiers =  [ "root.id", "volume_id", "device_name" ]
        attribute_map = [
          "instance_id=root.id",
          "device=device_name"
        ]
      }
    }
  }

  resource "ec2.egress_only_internet_gateways" {
    identifiers = [ "id" ]
    ignore_attributes = [ "arn", "attachments" ]
    iac {
      terraform {
        type = "aws_egress_only_internet_gateway"
      }
    }
  }

  resource "ec2.egress_only_internet_gateways#attachments" {
    identifiers = [ "id", sql("(attachments->>0)::jsonb->>'VpcId'") ] // attachments.0.VpcId
    ignore_attributes = [ "arn", "attachments" ]
    iac {
      terraform {
        type = "aws_egress_only_internet_gateway"
        identifiers = [ "id", "vpc_id" ]
      }
    }
  }

  resource "ec2.eips" {
    ignore_attributes = [ "network_interface_owner_id" ]
    iac {
      terraform {
        type = "aws_eip"
        identifiers = [ "allocation_id" ]
        attribute_map = [
          "network_interface_id=network_interface",
          "private_ip_address=private_ip"
        ]
      }
    }
  }

  resource "ec2.flow_logs" {
    iac {
      terraform {
        type = "aws_flow_log"
      }
    }
  }

  resource "ec2.hosts" {
    iac {
      terraform {
        type = "aws_ec2_host"
        identifiers = [ "arn" ]
      }
    }
  }

  resource "ec2.images" {
    identifiers = [ "id", "region" ]
    iac {
      terraform {
        type = "aws_imagebuilder_image"
        identifiers = [
          "output_resources.#.amis.#.image|@flatten|0",
          "output_resources.#.amis.#.region|@flatten|0",
        ]
      }
    }
  }

  resource "ec2.instances" {
    ignore_attributes = [ "launch_time", "ami_launch_index", "architecture", "client_token", "ena_support", "hypervisor",
      "metadata_options_http_protocol_ipv6", "metadata_options_state", "placement_availability_zone", "root_device_type",
      "state_code", "state_name", "virtualization_type", "vpc_id" ]

    iac {
      terraform {
        type = "aws_instance"
        attribute_map = [
          "cap_reservation_preference=capacity_reservation_specification.0.capacity_reservation_preference",
          "cpu_options_core_count=cpu_core_count",
          "cpu_options_threads_per_core=cpu_threads_per_core",
          "enclave_options_enabled=enclave_options.0.enabled",
          "hibernation_options_configured=hibernation",
          "image_id=ami",
          "metadata_options_http_endpoint=metadata_options.0.http_endpoint",
          "metadata_options_http_protocol_ipv6=metadata_options.0.",
          "metadata_options_http_put_response_hop_limit=metadata_options.0.http_put_response_hop_limit",
          "metadata_options_http_tokens=metadata_options.0.http_tokens",
          "monitoring_state=monitoring",
          "placement_tenancy=tenancy",
          "private_dns_name=private_dns",
          "private_ip_address=private_ip",
          "public_dns_name=public_dns",
          "public_ip_address=public_ip",
          "root_device_name=root_block_device.0.device_name"
        ]
      }
    }
  }

  resource "ec2.internet_gateways" {
    filters = [
      "NOT EXISTS (SELECT 1 FROM aws_ec2_internet_gateway_attachments a JOIN aws_ec2_vpcs v ON v.id=a.vpc_id WHERE a.internet_gateway_cq_id=c.cq_id AND v.is_default)",
      "NOT EXISTS (SELECT 1 FROM aws_ec2_route_table_routes WHERE gateway_id=c.id AND destination_cidr_block='0.0.0.0/0')"
    ]

    iac {
      terraform {
        type = "aws_internet_gateway"
      }
    }
  }

  resource "ec2.nat_gateways" {
    ignore_attributes = [ "arn", "create_time", "state", "vpc_id" ]
    iac {
      terraform {
        type = "aws_nat_gateway"
      }
    }
  }

  resource "aws_ec2_nat_gateway_addresses" {
    identifiers = [ "parent.id", "c.network_interface_id" ]
    iac {
      terraform {
        type = "aws_nat_gateway"
        identifiers = [ "id", "network_interface_id" ]
      }
    }
  }

  resource "ec2.network_acls" {
    filters = [ "c.is_default!=true" ]

    iac {
      terraform {
        type = "aws_network_acl"
      }
    }
  }

  #        resource "aws_ec2_network_acl_entries" {
  #            # TODO: no CRC32 function, no data in tests to verify
  #            identifiers = [ sql("CONCAT('nacl-',CRC32(CONCAT(parent.id,'-',c.rule_number,'-',CASE WHEN c.egress THEN 'true' ELSE 'false' END,'-',c.protocol,'-')))") ]
  #            filters = [ "((c.cidr_block='0.0.0.0/0' AND c.rule_number=32767) OR (c.ipv6_cidr_block=':/0' AND c.rule_number=32768)) AND c.rule_action='deny' AND c.protocol='-1'" ]
  #
  #            iac {
  #                terraform {
  #                    type = "aws_network_acl_rule"
  #                }
  #            }
  #        }

  resource "ec2.network_interfaces" {
    ignore_attributes = [ "availability_zone", "requester_id", "requester_managed", "status", "vpc_id" ]
    iac {
      terraform {
        type = "aws_network_interface"
        identifiers = [ "arn" ]
        attribute_map = [
          "groups|#.GroupId=security_groups",
          "private_ip_address=private_ip"
        ]
      }
    }
  }

  resource "ec2.regional_config#key" { # TODO: add account/region support
    identifiers = [ "ebs_default_kms_key_id" ]
    ignore_attributes = [ "ebs_encryption_enabled_by_default" ]
    iac {
      terraform {
        type = "aws_ebs_default_kms_key"
        identifiers = [ "key_arn" ]
      }
    }
  }

  resource "ec2.regional_config#ebs" { # TODO: add account/region support
    identifiers = [ sql("ebs_encryption_enabled_by_default::varchar") ]
    ignore_attributes = [ "ebs_default_kms_key_id" ]
    iac {
      terraform {
        type = "aws_ebs_default_kms_key"
        identifiers = [ "enabled" ]
      }
    }
  }

  resource "ec2.route_tables" {
    filters = [
      "NOT EXISTS (SELECT 1 FROM aws_ec2_route_table_associations WHERE route_table_cq_id=c.cq_id AND main)",
      "NOT EXISTS (SELECT 1 FROM aws_ec2_route_table_routes WHERE route_table_cq_id=c.cq_id AND origin='CreateRouteTable')"
    ]

    iac {
      terraform {
        type = "aws_route_table"
      }
    }
  }

  resource "ec2.security_groups" {
    filters = [ "c.group_name!='default'" ]

    iac {
      terraform {
        type = "aws_security_group"
        attribute_map = [
          "group_name=name"
        ]
      }
    }
  }

  resource "ec2.subnets" {
    ignore_attributes = [ "state", "default_for_az", "available_ip_address_count" ]
    filters = [ "c.default_for_az!=true" ]

    iac {
      terraform {
        type = "aws_subnet"
      }
    }
  }

  resource "ec2.transit_gateways" {
    iac {
      terraform {
        type = "aws_ec2_transit_gateway"
      }
    }
  }

  resource "ec2.vpc_endpoints" {
    iac {
      terraform {
        type = "aws_vpc_endpoint"
      }
    }
  }

  resource "ec2.vpc_peering_connections" {
    iac {
      terraform {
        type = "aws_vpc_peering_connection"
      }
    }
  }

  resource "ec2.vpcs" {
    ignore_attributes = [ "is_default", "state" ]
    filters = [ "c.is_default!=true" ]

    iac {
      terraform {
        type = "aws_vpc"
      }
    }
  }

  resource "ec2.vpn_gateways" {
    iac {
      terraform {
        type = "aws_vpn_gateway"
      }
    }
  }

  resource "ecr.repositories" {
    identifiers = [ "name" ]

    iac {
      terraform {
        type = "aws_ecr_repository"
      }
    }
  }

  resource "ecs.clusters" {
    identifiers = [ "arn" ]
    ignore_attributes = [ "active_services_count", "pending_tasks_count", "registered_container_instances_count", "running_tasks_count", "status" ]

    iac {
      terraform {
        type = "aws_ecs_cluster"
        attribute_map = [
#          "execute_config_logs_cloud_watch_encryption_enabled=configuration.0.execute_command_configuration.0.log_configuration.0|@getbool:cloud_watch_encryption_enabled",
#          "execute_config_log_s3_encryption_enabled=configuration.0.execute_command_configuration.0.log_configuration.0|@getbool:s3_bucket_encryption_enabled",
          "default_capacity_provider_strategy|0.Base=default_capacity_provider_strategy.0.base",
          "default_capacity_provider_strategy|0.CapacityProvider=default_capacity_provider_strategy.0.capacity_provider",
          "default_capacity_provider_strategy|0.Weight=default_capacity_provider_strategy.0.weight"
        ]
      }
    }
  }

  resource "ecs.task_definitions" {
    identifiers = [ "arn" ]
    ignore_attributes = [ "registered_at", "registered_by", "deregistered_at", "requires_attributes" ]

    iac {
      terraform {
        type = "aws_ecs_task_definition"
        attribute_map = [
          "compatibilities=requires_compatibilities"
        ]
      }
    }
  }

  resource "efs.filesystems" {
    iac {
      terraform {
        type = "aws_efs_file_system"
      }
    }
  }

  resource "eks.clusters" {
    identifiers = [ "name" ]

    iac {
      terraform {
        type = "aws_eks_cluster"
      }
    }
  }

  resource "elasticbeanstalk.applications" {
    identifiers = [ "arn" ]
    ignore_attributes = [ "date_created", "date_updated", "versions" ]

    iac {
      terraform {
        type = "aws_elastic_beanstalk_application"
        identifiers = [ "arn" ]
      }
    }
  }


  resource "elasticbeanstalk.environments" {
    iac {
      terraform {
        type = "aws_elastic_beanstalk_environment"
      }
    }
  }

  resource "elasticsearch.domains" {
    identifiers = [ "arn" ]

    iac {
      terraform {
        type = "aws_elasticsearch_domain"
      }
    }
  }

  resource "elbv1.load_balancers" {
    iac {
      terraform {
        type = "aws_elb"
      }
    }
  }

  resource "elbv2.load_balancers" {
    ignore_attributes = [ "created_time", "scheme", "state_code", "state_reason" ]
    sets = [ "subnets", "security_groups" ]
    iac {
      terraform {
        type = "aws_lb"
        attribute_map = [
          "canonical_hosted_zone_id=zone_id",
          "type=load_balancer_type"
        ]
      }
    }
  }

  resource "elbv2.target_groups" {
    ignore_attributes = [ "load_balancer_arns" ]
    iac {
      terraform {
        type = "aws_lb_target_group"
        attribute_map = [
          "health_check_enabled=health_check.0.enabled",
          "health_check_interval_seconds=health_check.0.interval",
          "health_check_path=health_check.0.path",
          "health_check_timeout_seconds=health_check.0.timeout",
          "healthy_threshold_count=health_check.0.healthy_threshold",
          "matcher_http_code=health_check.0.matcher",
          "unhealthy_threshold_count=health_check.0.unhealthy_threshold"
        ]
      }
    }
  }

  resource "emr.clusters" {
    identifiers = [ "id" ]
    iac {
      terraform {
        type = "aws_emr_cluster"
      }
    }
  }

  resource "fsx.backups" {
    iac {
      terraform {
        type = "aws_fsx_backup"
      }
    }
  }

  resource "guardduty.detectors" {
    identifiers = [ "arn" ]
    ignore_attributes = [ "created_at", "updated_at" ]

    iac {
      terraform {
        type = "aws_guardduty_detector"
        identifiers = [ "arn" ]
      }
    }
  }

  resource "iam.accounts#aliases" {
    identifiers = [ sql("array_to_string(aliases,',')") ]
    ignore_attributes = [
      "users", "users_quota",
      "groups", "groups_quota",
      "server_certificates", "server_certificates_quota",
      "user_policy_size_quota", "group_policy_size_quota", "groups_per_user_quota",
      "signing_certificates_per_user_quota", "access_keys_per_user_quota",
      "mfa_devices", "mfa_devices_in_use", "account_mfa_enabled",
      "account_access_keys_present", "account_signing_certificates_present",
      "attached_policies_per_group_quota", "policies", "policies_quota", "policy_size_quota", "policy_versions_in_use", "policy_versions_in_use_quota",
      "versions_per_policy_quota", "global_endpoint_token_version"
    ]
    filters = [
      "aliases IS NOT NULL"
    ]
    iac {
      terraform {
        type = "aws_iam_account_alias"
        identifiers = [ "account_alias" ]
        attribute_map = [
          "aliases|0=account_alias"
        ]
      }
    }
  }

  resource "iam.groups" {
    identifiers = [ "name" ]
    iac {
      terraform {
        type = "aws_iam_group"
      }
    }
  }

  resource "iam.openid_connect_identity_providers" {
    iac {
      terraform {
        type = "aws_iam_openid_connect_provider"
      }
    }
  }

  resource "iam.policies" {
    identifiers = [ "arn" ]
    ignore_attributes = [ "attachment_count", "create_date", "permissions_boundary_usage_count", "update_date", "is_attachable", "default_version_id" ]

    iac {
      terraform {
        type = "aws_iam_policy"
        attribute_map = [
          "id=policy_id"
        ]
      }
    }
  }

  resource "iam.roles" {
    identifiers = [ "name" ]
    ignore_attributes = [ "role_last_used_region", "role_last_used_last_used_date", "permissions_boundary_type", "policies" ]
    iac {
      terraform {
        type = "aws_iam_role"

        attribute_map = [
          "name=id",
          "id=unique_id",
          "permissions_boundary_arn=permissions_boundary",
          "assume_role_policy_document=assume_role_policy"
        ]
      }
    }
  }

  resource "iam.saml_identity_providers" {
    iac {
      terraform {
        type = "aws_iam_saml_provider"
      }
    }
  }

  resource "iam.server_certificates" {
    iac {
      terraform {
        type = "aws_iam_server_certificate"
      }
    }
  }

  resource "iam.users" {
    identifiers       = ["user_name"]
    attributes = [ "arn", "path", "permissions_boundary_arn", "permissions_boundary_type", "tags" ]

    iac {
      terraform {
        type = "aws_iam_user"
      }
    }
  }

  resource "aws_iam_user_groups" {
    identifiers = [ "group_name" ]

    iac {
      terraform {
        type = "aws_iam_group"
      }
    }
  }

  resource "aws_iam_user_access_keys" {
    iac {
      terraform {
        type = "aws_iam_access_key"
      }
    }
  }

  resource "aws_iam_user_attached_policies" {
    identifiers = [ sql("CONCAT(parent.user_name, ':user_', c.policy_name)") ]

    iac {
      terraform {
        type = "aws_iam_user_policy"
      }
    }
  }

  resource "aws_iam_user_policies" {
    identifiers = [ sql("CONCAT(parent.user_name, ':', c.policy_name)") ]

    iac {
      terraform {
        type = "aws_iam_user_policy"
      }
    }
  }

  resource "iam.virtual_mfa_devices" {
    identifiers = [ "serial_number" ]
    ignore_attributes = [ "enable_date" ]

    iac {
      terraform {
        type = "aws_iam_virtual_mfa_device"
        identifiers = [ "arn" ]
      }
    }
  }

  resource "iot.certificates" {
    identifiers = [ "id" ]

    ignore_attributes = [ "customer_version", "ca_certificate_id", "mode", "policies", "last_modified_date", "generation_id", "owned_by", "previous_owned_by", "transfer_data_accept_date", "transfer_data_reject_date", "transfer_data_reject_reason", "transfer_data_transfer_date", "transfer_data_transfer_message", "validity_not_after", "validity_not_before" ]

    iac {
      terraform {
        type = "aws_iot_certificate"

        attribute_map = [
          "pem=certificate_pem",
          "status=active|@iftrue:ACTIVE"
        ]
      }
    }
  }

  resource "iot.policies" {
    identifiers = [ "name" ]
    ignore_attributes = [ "generation_id", "last_modified_date" ]

    iac {
      terraform {
        type = "aws_iot_policy"

        attribute_map = [
          "name=id",
          "document=policy",
        ]
      }
    }
  }

  resource "iot.thing_groups" {
    identifiers = [ "name" ]
    ignore_attributes = [ "id", "things_in_group", "index_name", "query_string", "query_version", "status", "root_to_parent_thing_groups", "attribute_payload_merge" ]

    iac {
      terraform {
        type = "aws_iot_thing_group"

        attribute_map = [
          "name=id",
          "thing_group_description=properties.#.description|@flatten|0",
          "attribute_payload_attributes=properties.#.attribute_payload.#.attributes|@flatten|0"
        ]
      }
    }
  }

  resource "iot.thing_types" {
    identifiers = [ "name" ]

    iac {
      terraform {
        type = "aws_iot_thing_type"

        attribute_map = [
          "description=properties.#.description|@flatten|0"
        ]
      }
    }
  }

  resource "iot.things" {
    identifiers = [ "name" ]

    iac {
      terraform {
        type = "aws_iot_thing"

        attribute_map = [
          "type_name=thing_type_name"
        ]
      }
    }
  }

  resource "iot.topic_rules" {
    identifiers = [ "rule_name" ]
    ignore_attributes = [
      "error_action_cloudwatch_logs_log_group_name", "error_action_cloudwatch_logs_role_arn", "error_action_firehose_batch_mode",
      "error_action_http_url", "error_action_http_auth_sigv4_role_arn", "error_action_http_auth_sigv4_service_name", "error_action_http_auth_sigv4_signing_region", "error_action_http_confirmation_url", "error_action_http_headers",
      "error_action_iot_analytics_batch_mode", "error_action_iot_analytics_channel_arn",
      "error_action_iot_events_batch_mode",
      "error_action_iot_site_wise",
      "error_action_kafka_client_properties", "error_action_kafka_destination_arn", "error_action_kafka_topic", "error_action_kafka_key", "error_action_kafka_partition",
      "error_action_open_search_endpoint", "error_action_open_search_id", "error_action_open_search_index", "error_action_open_search_role_arn", "error_action_open_search_type",
      "error_action_salesforce_token", "error_action_salesforce_url",
      "error_action_s3_canned_acl",
      "error_action_timestream_database_name", "error_action_timestream_dimensions", "error_action_timestream_role_arn", "error_action_timestream_table_name", "error_action_timestream_timestamp_unit", "error_action_timestream_timestamp_value"
    ]

    iac {
      terraform {
        type = "aws_iot_topic_rule"

        attribute_map = [
          "rule_name=name",
          "rule_disabled=enabled|@inverse",
          "aws_iot_sql_version=sql_version",
          "error_action_cloudwatch_alarm_name=error_action.#.cloudwatch_alarm.#.name|@flatten|0",
          "error_action_cloudwatch_alarm_role_arn=error_action.#.cloudwatch_alarm.#.role_arn|@flatten|0",
          "error_action_cloudwatch_alarm_state_reason=error_action.#.cloudwatch_alarm.#.state_reason|@flatten|0",
          "error_action_cloudwatch_alarm_state_value=error_action.#.cloudwatch_alarm.#.state_value|@flatten|0",
          "error_action_cloudwatch_metric_metric_name=error_action.#.cloudwatch_metric.#.metric_name|@flatten|0",
          "error_action_cloudwatch_metric_metric_namespace=error_action.#.cloudwatch_metric.#.metric_namespace|@flatten|0",
          "error_action_cloudwatch_metric_unit=error_action.#.cloudwatch_metric.#.metric_unit|@flatten|0",
          "error_action_cloudwatch_metric_value=error_action.#.cloudwatch_metric.#.metric_value|@flatten|0",
          "error_action_cloudwatch_metric_role_arn=error_action.#.cloudwatch_metric.#.role_arn|@flatten|0",
          "error_action_cloudwatch_metric_timestamp=error_action.#.cloudwatch_metric.#.metric_timestamp|@flatten|0",
          "error_action_dynamo_db_hash_key_field=error_action.#.dynamodb.#.hash_key_field|@flatten|0",
          "error_action_dynamo_db_hash_key_value=error_action.#.dynamodb.#.hash_key_value|@flatten|0",
          "error_action_dynamo_db_role_arn=error_action.#.dynamodb.#.role_arn|@flatten|0",
          "error_action_dynamo_db_table_name=error_action.#.dynamodb.#.table_name|@flatten|0",
          "error_action_dynamo_db_hash_key_type=error_action.#.dynamodb.#.hash_key_type|@flatten|0",
          "error_action_dynamo_db_operation=error_action.#.dynamodb.#.operation|@flatten|0",
          "error_action_dynamo_db_payload_field=error_action.#.dynamodb.#.payload_field|@flatten|0",
          "error_action_dynamo_db_range_key_field=error_action.#.dynamodb.#.range_key_field|@flatten|0",
          "error_action_dynamo_db_range_key_type=error_action.#.dynamodb.#.range_key_type|@flatten|0",
          "error_action_dynamo_db_range_key_value=error_action.#.dynamodb.#.range_key_value|@flatten|0",
          "error_action_dynamo_db_v2_put_item_table_name=error_action.#.dynamodbv2.#.put_item.table_name|@flatten|0",
          "error_action_dynamo_db_v2_role_arn=error_action.#.dynamodbv2.#.role_arn|@flatten|0",
          "error_action_elasticsearch_endpoint=error_action.#.elasticsearch.#.endpoint|@flatten|0",
          "error_action_elasticsearch_id=error_action.#.elasticsearch.#.id|@flatten|0",
          "error_action_elasticsearch_index=error_action.#.elasticsearch.#.index|@flatten|0",
          "error_action_elasticsearch_role_arn=error_action.#.elasticsearch.#.role_arn|@flatten|0",
          "error_action_elasticsearch_type=error_action.#.elasticsearch.#.type|@flatten|0",
          "error_action_firehose_delivery_stream_name=error_action.#.firehose.#.delivery_stream_name|@flatten|0",
          "error_action_firehose_role_arn=error_action.#.firehose.#.role_arn|@flatten|0",
          "error_action_firehose_separator=error_action.#.firehose.#.separator|@flatten|0",
          "error_action_iot_analytics_channel_name=error_action.#.iot_analytics.#.channel_name|@flatten|0",
          "error_action_iot_analytics_role_arn=error_action.#.iot_analytics.#.role_arn|@flatten|0",
          "error_action_iot_events_input_name=error_action.#.iot_events.#.input_name|@flatten|0",
          "error_action_iot_events_role_arn=error_action.#.iot_events.#.role_arn|@flatten|0",
          "error_action_iot_events_message_id=error_action.#.iot_events.#.message_id|@flatten|0",
          "error_action_kinesis_stream_name=error_action.#.kinesis.#.stream_name|@flatten|0",
          "error_action_kinesis_partition_key=error_action.#.kinesis.#.partition_key|@flatten|0",
          "error_action_kinesis_role_arn=error_action.#.kinesis.#.role_arn|@flatten|0",
          "error_action_lambda_function_arn=error_action.#.lambda.#.function_arn|@flatten|0",
          "error_action_republish_topic=error_action.#.republish.#.topic|@flatten|0",
          "error_action_republish_qos=error_action.#.republish.#.qos|@flatten|0",
          "error_action_republish_role_arn=error_action.#.republish.#.role_arn|@flatten|0",
          "error_action_s3_bucket_name=error_action.#.s3.#.bucket_name|@flatten|0",
          "error_action_s3_key=error_action.#.s3.#.key|@flatten|0",
          "error_action_s3_role_arn=error_action.#.s3.#.role_arn|@flatten|0",
          "error_action_sns_role_arn=error_action.#.sns.#.role_arn|@flatten|0",
          "error_action_sns_target_arn=error_action.#.sns.#.target_arn|@flatten|0",
          "error_action_sns_message_format=error_action.#.sns.#.message_format|@flatten|0",
          "error_action_sqs_queue_url=error_action.#.sqs.#.queue_url|@flatten|0",
          "error_action_sqs_role_arn=error_action.#.sqs.#.role_arn|@flatten|0",
          "error_action_sqs_use_base64=error_action.#.sqs.#.use_base64|@flatten|0",
          "error_action_step_functions_role_arn=error_action.#.step_functions.#.role_arn|@flatten|0",
          "error_action_step_functions_state_machine_name=error_action.#.step_functions.#.state_machine_name|@flatten|0",
          "error_action_step_functions_execution_name_prefix=error_action.#.step_functions.#.execution_name_prefix|@flatten|0",
        ]
      }
    }
  }

  resource "kms.keys" {
    identifiers = [ "id" ]
    ignore_attributes = [ "encryption_algorithms", "manager", "key_state", "origin", "deletion_date", "valid_to" ]
    iac {
      terraform {
        type = "aws_kms_key"
        attribute_map = [
          "rotation_enabled=enable_key_rotation",
          "enabled=is_enabled"
        ]
      }
    }
  }

  resource "lambda.functions" {
    identifiers = [ "arn" ]
    ignore_attributes =  [ "code_location", "code_repository_type", "last_update_status", "revision_id", "state", "policy_document", "policy_revision_id" ]
    sets = [ "vpc_config_subnet_ids", "vpc_config_security_group_ids" ]
    iac {
      terraform {
        type = "aws_lambda_function"
        identifiers = [ "arn" ]
        attribute_map = [
          "name=id",
          "code_size=source_code_size",
          "code_sha256=source_code_hash",
          "dead_letter_config_target_arn=dead_letter_config.#.target_arn|@flatten|0",
          "environment_variables=environment.#.variables|0",
          "tracing_config_mode=tracing_config.#.mode|@flatten|0",
          "vpc_config_security_group_ids=vpc_config.0.security_group_ids",
          "vpc_config_subnet_ids=vpc_config.0.subnet_ids",
          "vpc_config_vpc_id=vpc_config.0.vpc_id"
        ]
      }
    }
  }

  resource "aws_lambda_layer_versions" {
    identifiers = [ "layer_version_arn" ]

    iac {
      terraform {
        type = "aws_lambda_layer_version"
        identifiers = [ "arn" ]
        attribute_map = [
          "layer_version_arn=arn"
        ]
      }
    }
  }

  resource "mq.brokers" {
    ignore_attributes = [ "created", "broker_state" ]
    sets = [ "instances", "subnet_ids" ]
    iac {
      terraform {
        type = "aws_mq_broker"
        attribute_map = [
          "broker_instances=instances", # TODO CamelCase vs snake_case in keys of map
          "encryption_options_use_aws_owned_key=encryption_options.0.use_aws_owned_key",
          "logs|@getbool:Audit=logs.0|@getbool:audit",
          "logs|@getbool:General=logs.0|@getbool:general",
          "maintenance_window_start_time=maintenance_window_start_time.0" # TODO CamelCase vs snake_case in keys of map
        ]
      }
    }
  }

  resource "organizations.accounts" {
    #    identifiers = [ "account_id", "id" ]
    identifiers = [ "arn" ]
    ignore_attributes = [ "joined_timestamp", "joined_method" ]
    iac {
      terraform {
        type = "aws_organizations_account"
        #        identifiers = [ "parent_id", "id" ]
        identifiers = [ "arn" ]
      }
    }
  }

  resource "qldb.ledgers" {
    ignore_attributes = [ "creation_date_time", "state", "inaccessible_kms_key_date_time" ]
    iac {
      terraform {
        type = "aws_qldb_ledger"
        identifiers = [ "arn" ]
        attribute_map = [
          "kms_key_arn=kms_key"
        ]
      }
    }
  }

  resource "rds.clusters" {
    identifiers = [ "db_cluster_identifier" ]
    iac {
      terraform {
        type = "aws_rds_cluster"
      }
    }
  }

  resource "rds.cluster_parameter_groups" {
    iac {
      terraform {
        type = "aws_rds_cluster_parameter_group"
        identifiers = [ "arn" ]
      }
    }
  }

  resource "aws_rds_cluster_parameters" {
    identifiers = [ "parent.arn", "c.parameter_name" ]
    ignore_attributes = [ "allowed_values", "apply_type", "data_type", "description", "is_modifiable", "minimum_engine_version", "source", "supported_engine_modes" ]
    filters = [ "c.source NOT IN ('engine-default', 'system')" ]

    iac {
      terraform {
        type = "aws_rds_cluster_parameter_group"
        path = "parameter"
        identifiers = [ "root.arn", "name" ]
        attribute_map = [
          "parameter_name=name",
          "parameter_value=value"
        ]
      }
    }
  }

  resource "rds.cluster_snapshots" {
    identifiers = [ "db_cluster_identifier", "arn" ]
    ignore_attributes = [ "snapshot_create_time" ]
    iac {
      terraform {
        type = "aws_db_cluster_snapshot"
        identifiers = [ "db_cluster_identifier", "db_cluster_snapshot_arn" ]
      }
    }
  }

  resource "rds.db_parameter_groups" {
    identifiers = [ "arn" ]
    iac {
      terraform {
        type = "aws_db_parameter_group"
        identifiers = [ "arn" ]
      }
    }
  }

  resource "aws_rds_db_parameters" {
    identifiers = [ "parent.arn", "c.parameter_name" ]
    ignore_attributes = [ "allowed_values", "apply_type", "data_type", "description", "is_modifiable", "minimum_engine_version", "source", "supported_engine_modes" ]
    filters = [ "c.source NOT IN ('engine-default', 'system')" ]

    iac {
      terraform {
        type = "aws_db_parameter_group"
        path = "parameter"
        identifiers = [ "root.arn", "name" ]
        attribute_map = [
          "parameter_name=name",
          "parameter_value=value"
        ]
      }
    }
  }


  resource "rds.db_security_groups" {
    identifiers = [ "arn" ]
    iac {
      terraform {
        type = "aws_db_security_group"
        identifiers = [ "arn" ]
      }
    }
  }

  resource "rds.db_snapshots" {
    identifiers = [ "db_instance_identifier", "arn" ]
    ignore_attributes = [ "snapshot_create_time" ]
    iac {
      terraform {
        type = "aws_db_snapshot"
        identifiers = [ "db_instance_identifier", "db_snapshot_arn" ]
      }
    }
  }

  resource "rds.db_subnet_groups" {
    identifiers = [ "name" ]
    ignore_attributes = [ "status", "vpc_id" ]
    sets = [ "subnet_ids" ]
    filters = [
      "NOT EXISTS (SELECT 1 FROM aws_ec2_vpcs WHERE id=c.vpc_id AND is_default)",
    ]
    iac {
      terraform {
        type = "aws_db_subnet_group"
      }
    }
  }

  resource "rds.event_subscriptions" {
    identifiers = [ "arn" ]
    ignore_attributes = [ "cust_subscription_id", "subscription_creation_time" ]
    sets = [ "event_categories_list", "source_id_list" ]
    iac {
      terraform {
        type = "aws_db_event_subscription"
        identifiers = [ "arn" ]
        attribute_map = [
          "event_categories_list=event_categories",
          "sns_topic_arn=sns_topic",
          "source_id_list=source_ids",
          "status=enabled|@iftrue:active"
        ]
      }
    }
  }

  resource "rds.instances" {
    identifiers = [ "arn" ]
    ignore_attributes = [
      "allocated_storage", "backup_retention_period", "customer_owned_ip_enabled", "user_instance_id", "db_instance_status", "db_name",
      "subnet_group_description", "subnet_group_subnet_group_status", "subnet_group_vpc_id", "instance_port", "deletion_protection",
      "endpoint_hosted_zone_id", "iam_database_authentication_enabled", "instance_create_time", "license_model", "listener_endpoint_port",
      "master_username", "multi_az", "performance_insights_retention_period", "storage_type"
    ]

    iac {
      terraform {
        identifiers = [ "arn" ]
        type = "aws_rds_cluster_instance"
        attribute_map = [
          "id=dbi_resource_id",
          "ca_certificate_identifier=ca_cert_identifier",
          "endpoint_address=endpoint",
          "endpoint_port=port",
          "db_instance_class=instance_class",
          "subnet_group_name=db_subnet_group_name"
        ]
      }
    }
  }

  resource "redshift.clusters" {
    ignore_attributes = [ "cluster_create_time", "cluster_availability_status", "availability_zone_relocation_status", "availability_zone_relocation_status",
      "cluster_namespace_arn", "cluster_snapshot_copy_status_manual_snapshot_retention_period", "cluster_snapshot_copy_status_retention_period",
      "cluster_status", "data_transfer_progress_data_transferred_in_mega_bytes", "data_transfer_progress_total_data_in_mega_bytes", "maintenance_track_name",
      "manual_snapshot_retention_period", "next_maintenance_window_start_time", "resize_info_allow_cancel_resize", "restore_status_current_restore_rate_in_mega_bytes_per_second",
      "restore_status_elapsed_time_in_seconds", "restore_status_estimated_time_to_completion_in_seconds", "restore_status_progress_in_mega_bytes", "restore_status_snapshot_size_in_mega_bytes",
      "total_storage_capacity_in_mega_bytes", "vpc_id" ]
    iac {
      terraform {
        type = "aws_redshift_cluster"
        attribute_map = [
          "db_name=database_name",
          "endpoint_address=endpoint|@split::|0",
          "endpoint_port=port",
          "logging_status|BucketName=logging.0.bucket_name",
          "logging_status|@getbool:LoggingEnabled=logging.0|@getbool:enable",
          "logging_status|S3KeyPrefix=logging.0.s3_key_prefix",
        ]
      }
    }
  }

  resource "redshift.event_subscriptions" {
    identifiers = [ "id" ]
    ignore_attributes = [ "subscription_creation_time", "status" ]
    sets = [ "source_ids_list", "event_categories_list" ]
    iac {
      terraform {
        type = "aws_redshift_event_subscription"
        attribute_map = [
          "source_ids_list=source_ids",
          "event_categories_list=event_categories"
        ]
      }
    }
  }

  resource "redshift.subnet_groups" {
    ignore_attributes = [ "subnet_group_status", "vpc_id" ]
    filters = [
      "NOT EXISTS (SELECT 1 FROM aws_ec2_vpcs WHERE id=c.vpc_id AND is_default)",
    ]
    iac {
      terraform {
        type = "aws_redshift_subnet_group"
        attribute_map = [
          "cluster_subnet_group_name=name"
        ]
      }
    }
  }

  resource "route53.domains" {
    iac {
      terraform {
        type = "aws_route53domains_registered_domain"
        identifiers = [ "domain_name" ]
      }
    }
  }

  resource "route53.health_checks" {
    iac {
      terraform {
        type = "aws_route53_health_check"
      }
    }
  }

  resource "route53.hosted_zones" {
    iac {
      terraform {
        type = "aws_route53_zone"
      }
    }
  }

  resource "route53.reusable_delegation_sets" {
    identifiers = [ sql("SPLIT_PART(c.id, '/', 3)") ]
    iac {
      terraform {
        type = "aws_route53_delegation_set"
      }
    }
  }

  resource "aws_route53_traffic_policy_versions" {
    identifiers = [ "id", "version" ]
    ignore_attributes = [ "comment" ]
    iac {
      terraform {
        type = "aws_route53_traffic_policy"
        identifiers = [ "id", "version" ]
      }
    }
  }

#  resource "s3.accounts" {
#    ignore_attributes = [ "config_exists" ]
#
#    iac {
#      terraform {
#        type = "aws_s3_account_public_access_block"
#      }
#    }
#  }

  resource "s3.buckets" {
    ignore_attributes = [ "name" ]

    iac {
      terraform {
        type = "aws_s3_bucket"
      }
    }
  }

  resource "sagemaker.endpoint_configurations" {
    ignore_attributes = [ "creation_time" ]
    iac {
      terraform {
        type = "aws_sagemaker_endpoint_configuration"
        identifiers = [ "arn" ]
      }
    }
  }

  resource "sagemaker.models" {
    ignore_attributes = [ "creation_time" ]
    iac {
      terraform {
        type = "aws_sagemaker_model"
        identifiers = [ "arn" ]
      }
    }
  }

  resource "sagemaker.notebook_instances" {
    ignore_attributes = ["creation_time", "last_modified_time", "notebook_instance_status", "accelerator_types" ]
    iac {
      terraform {
        type        = "aws_sagemaker_notebook_instance"
        identifiers = ["arn"]
        attribute_map = [
          "volume_size_in_gb=volume_size",
          "direct_internet_access=direct_internet_access|@if:Enabled,true",
          "notebook_instance_lifecycle_config_name=lifecycle_config_name"
        ]
      }
    }
  }

  resource "secretsmanager.secrets" {
    ignore_attributes = [ "created_date", "deleted_date", "last_accessed_date", "last_changed_date", "last_rotated_date" ]
    iac {
      terraform {
        type = "aws_secretsmanager_secret"
        identifiers = ["arn"]
      }
    }
  }

  resource "sns.subscriptions" {
    iac {
      terraform {
        type = "aws_sns_topic_subscription"
      }
    }
  }

  resource "sns.topics" {
    ignore_attributes = [ "subscriptions_confirmed", "subscriptions_deleted", "subscriptions_pending", "effective_delivery_policy" ]
    iac {
      terraform {
        type = "aws_sns_topic"
      }
    }
  }

  resource "sqs.queues" {
    identifiers = [ "url" ]
    ignore_attributes = [
      "created_timestamp", "last_modified_timestamp",
      "approximate_number_of_messages", "approximate_number_of_messages_not_visible", "approximate_number_of_messages_delayed",
      "policy", "redrive_policy" # string type in TF, json type in CQ
    ]
    iac {
      terraform {
        type = "aws_sqs_queue"
        attribute_map = [
          "maximum_message_size=max_message_size",
          "message_retention_period=message_retention_seconds",
          "receive_message_wait_time_seconds=receive_wait_time_seconds",
          "kms_data_key_reuse_period_seconds=kms_data_key_reuse_period_seconds",
          "visibility_timeout=visibility_timeout_seconds"
        ]
      }
    }
  }

  resource "ssm.documents" {
    ignore_attributes = [ "created_date" ]
    sets = [ "platform_types" ]
    iac {
      terraform {
        type = "aws_ssm_document"
        identifiers = ["arn"]
      }
    }
  }

  resource "waf.rule_groups" {
    sets = [ "rule_ids" ]
    iac {
      terraform {
        type = "aws_waf_rule_group"
        attribute_map = [
          "rule_ids=activated_rule.#.rule_id"
        ]
      }
    }
  }

  resource "waf.rules" {
    iac {
      terraform {
        type = "aws_waf_rule"
      }
    }
  }

  resource "waf.web_acls" {
    iac {
      terraform {
        type = "aws_waf_web_acl"
      }
    }
  }


  resource "wafregional.rule_groups" {
    sets = [ "rule_ids" ]
    iac {
      terraform {
        type = "aws_wafregional_rule_group"
        attribute_map = [
          "rule_ids=activated_rule.#.rule_id"
        ]
      }
    }
  }

  resource "wafregional.rules" {
    iac {
      terraform {
        type = "aws_wafregional_rule"
      }
    }
  }

  resource "wafregional.rate_based_rules" {
    iac {
      terraform {
        type = "aws_wafregional_rate_based_rule"
      }
    }
  }

  resource "wafregional.web_acls" {
    iac {
      terraform {
        type = "aws_wafregional_web_acl"
        attribute_map = [
          "default_action=default_action.0.type"
        ]
      }
    }
  }

  resource "wafv2.ipsets" {
    iac {
      terraform {
        type = "aws_wafv2_ip_set"
        identifiers = [ "arn" ]
      }
    }
  }

  resource "wafv2.regex_pattern_sets" {
    sets = [ "regular_expression_list" ]
    iac {
      terraform {
        type = "aws_wafv2_regex_pattern_set"
        identifiers = [ "arn" ]
        attribute_map = [
          "regular_expression_list=regular_expression.#.regex_string"
        ]
      }
    }
  }

  resource "wafv2.rule_groups" {
    iac {
      terraform {
        type = "aws_wafv2_rule_group"
      }
    }
  }

  resource "wafv2.web_acls" {
    iac {
      terraform {
        type = "aws_wafv2_web_acl"
      }
    }
  }

  resource "workspaces.directories" {
    identifiers = [ "id" ]
    ignore_attributes = [ "arn", "state", "tenancy" ]
    sets = [ "subnet_ids" ]
    iac {
      terraform {
        type = "aws_workspaces_directory"
        identifiers = [ "id" ]
        attribute_map = [
          "name=directory_name",
          "type=directory_type",
          "change_compute_type=self_service_permissions.0.change_compute_type|@iftrue:ENABLED",
          "increase_volume_size=self_service_permissions.0.increase_volume_size|@iftrue:ENABLED",
          "rebuild_workspace=self_service_permissions.0.rebuild_workspace|@iftrue:ENABLED",
          "restart_workspace=self_service_permissions.0.restart_workspace|@iftrue:ENABLED",
          "switch_running_mode=self_service_permissions.0.switch_running_mode|@iftrue:ENABLED",
          "device_type_android=workspace_access_properties.0.device_type_android",
          "device_type_chrome_os=workspace_access_properties.0.device_type_chromeos",
          "device_type_ios=workspace_access_properties.0.device_type_ios",
          "device_type_linux=workspace_access_properties.0.device_type_linux",
          "device_type_osx=workspace_access_properties.0.device_type_osx",
          "device_type_web=workspace_access_properties.0.device_type_web",
          "device_type_windows=workspace_access_properties.0.device_type_windows",
          "device_type_zero_client=workspace_access_properties.0.device_type_zeroclient",
          "custom_security_group_id=workspace_creation_properties.0.custom_security_group_id",
          "default_ou=workspace_creation_properties.0.default_ou",
          "enable_internet_access=workspace_creation_properties.0|@getbool:enable_internet_access",
          "enable_maintenance_mode=workspace_creation_properties.0|@getbool:enable_maintenance_mode",
          "enable_work_docs=workspace_creation_properties.0|@getbool:enable_work_docs",
          "user_enabled_as_local_administrator=workspace_creation_properties.0|@getbool:user_enabled_as_local_administrator"
        ]
      }
    }
  }

  resource "workspaces.workspaces" {
    identifiers = [ "id" ]
    ignore_attributes = [ "arn", "subnet_id" ]
    iac {
      terraform {
        type = "aws_workspaces_workspace"
        identifiers = [ "id" ]
        attribute_map = [
          "compute_type_name=workspace_properties.0.compute_type_name",
          "root_volume_size_gib=workspace_properties.0.root_volume_size_gib",
          "running_mode=workspace_properties.0.running_mode",
          "running_mode_auto_stop_timeout_in_minutes=workspace_properties.0.running_mode_auto_stop_timeout_in_minutes",
          "user_volume_size_gib=workspace_properties.0.user_volume_size_gib",
        ]
      }
    }
  }

  resource "xray.encryption_config" {
    identifiers = [ "region" ]
    ignore_attributes = [ "status" ]
    iac {
      terraform {
        type = "aws_xray_encryption_config"
      }
    }
  }

}
