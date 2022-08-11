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
    iac {
      terraform {
        type = "aws_apigatewayv2_api"
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
    ignore_attributes = [ "api_gateway_managed", "created_date", "last_updated_date" ]

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
    ignore_attributes = [
      "created_time", "load_balancers", "notifications_configurations", "metrics", "status",
      "load_balancer_target_groups", "suspended_processes", "vpc_zone_identifier" # These are supported in v2
    ]
    sets = [ "availability_zones" ]

    iac {
      terraform {
        type = "aws_autoscaling_group"
        identifiers = [ "arn" ]

        attribute_map = [
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

    iac {
      terraform {
        type = "aws_launch_configuration"
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
    ignore_attributes = [ "created" ]

    iac {
      terraform {
        type = "aws_codebuild_project"
        identifiers = [ "arn" ]
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
    iac {
      terraform {
        type = "aws_instance"
        path = "root_block_device"
        identifiers =  [ "root.id", "volume_id", "device_name" ]
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
    ignore_attributes = ["launch_time"]

    iac {
      terraform {
        type = "aws_instance"
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
      }
    }
  }

  resource "ec2.subnets" {
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

    iac {
      terraform {
        type = "aws_ecs_cluster"
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
    iac {
      terraform {
        type = "aws_lb"
      }
    }
  }

  resource "elbv2.target_groups" {
    iac {
      terraform {
        type = "aws_lb_target_group"
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

#  resource "iam.accounts" {
#    identifiers = [ "aliases" ]
#    iac {
#      terraform {
#        type = "aws_iam_account_alias"
#        identifiers = [ "aliases" ]
#      }
#    }
#  }

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

    iac {
      terraform {
        type = "aws_iam_policy"
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
    iac {
      terraform {
        type = "aws_kms_key"
      }
    }
  }

  resource "lambda.functions" {
    identifiers = [ "arn" ]
    ignore_attributes =  [ "code_location", "code_repository_type", "last_update_status", "revision_id", "state" ]
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
          "tracing_config_mode=tracing_config.#.mode|@flatten|0"
        ]
      }
    }
  }

  resource "aws_lambda_layer_versions" {
    identifiers = [ sql("CONCAT(parent.arn, ':', c.version)") ]

    iac {
      terraform {
        type = "aws_lambda_layer_version"
      }
    }
  }

  resource "mq.brokers" {
    iac {
      terraform {
        type = "aws_mq_broker"
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
    iac {
      terraform {
        type = "aws_redshift_cluster"
      }
    }
  }

  resource "redshift.event_subscriptions" {
    identifiers = [ "id" ]
    ignore_attributes = [ "subscription_creation_time", "status" ]
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
    filters = [
      "NOT EXISTS (SELECT 1 FROM aws_ec2_vpcs WHERE id=c.vpc_id AND is_default)",
    ]
    iac {
      terraform {
        type = "aws_redshift_subnet_group"
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
    ignore_attributes = [ "policy", "redrive_policy" ] # string type in TF, json type in CQ
    iac {
      terraform {
        type = "aws_sqs_queue"
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

}
