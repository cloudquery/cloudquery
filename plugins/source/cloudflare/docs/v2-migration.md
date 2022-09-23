# Schema Changes from v1 to v2
This guide summarizes schema changes from CloudQuery v1 to v2. It is automatically generated and
not guaranteed to be complete, but we hope it helps as a starting point and reference when migrating to v2.

Last updated Fri Sep 23 15:00:02 BST 2022.

## aws_access_analyzer_analyzer_archive_rules
Moved to JSON column on [aws_accessanalyzer_analyzers](#aws_accessanalyzer_analyzers)


## aws_access_analyzer_analyzer_finding_sources
Moved to JSON column on [aws_accessanalyzer_analyzers](#aws_accessanalyzer_analyzers)


## aws_access_analyzer_analyzer_findings
Moved to JSON column on [aws_accessanalyzer_analyzers](#aws_accessanalyzer_analyzers)


## aws_access_analyzer_analyzers
Renamed to [aws_accessanalyzer_analyzers](#aws_accessanalyzer_analyzers)


## aws_accessanalyzer_analyzer_archive_rules
This table was newly added.

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|account_id|text|added|
|region|text|added|
|analyzer_arn|text|added|
|created_at|timestamp without time zone|added|
|filter|jsonb|added|
|rule_name|text|added|
|updated_at|timestamp without time zone|added|
|_cq_id|uuid|added|
|_cq_fetch_time|timestamp without time zone|added|

## aws_accessanalyzer_analyzer_findings
This table was newly added.

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|account_id|text|added|
|region|text|added|
|arn|text|added|
|analyzer_arn|text|added|
|analyzed_at|timestamp without time zone|added|
|condition|jsonb|added|
|created_at|timestamp without time zone|added|
|id|text|added|
|resource_owner_account|text|added|
|resource_type|text|added|
|status|text|added|
|updated_at|timestamp without time zone|added|
|action|text[]|added|
|error|text|added|
|is_public|boolean|added|
|principal|jsonb|added|
|resource|text|added|
|sources|jsonb|added|
|_cq_id|uuid|added|
|_cq_fetch_time|timestamp without time zone|added|

## aws_accessanalyzer_analyzers
Renamed from [aws_access_analyzer_analyzers](aws_access_analyzer_analyzers)


## aws_accounts
Renamed to [aws_iam_accounts](#aws_iam_accounts)


## aws_acm_certificates

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|certificate_transparency_logging_preference|text|removed|
|key_usages|jsonb|updated|Type changed from text[] to jsonb
|options|jsonb|added|
|renewal_summary|jsonb|added|
|renewal_summary_domain_validation_options|jsonb|removed|
|renewal_summary_failure_reason|text|removed|
|renewal_summary_status|text|removed|
|renewal_summary_updated_at|timestamp without time zone|removed|

## aws_apigateway_api_keys

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|

## aws_apigateway_client_certificates

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|client_certificate_id|text|added|
|id|text|removed|

## aws_apigateway_domain_name_base_path_mappings

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|account_id|text|added|
|domain_name|text|removed|
|domain_name_arn|text|added|
|domain_name_cq_id|uuid|removed|
|region|text|added|

## aws_apigateway_domain_names

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|endpoint_configuration|jsonb|added|
|endpoint_configuration_types|text[]|removed|
|endpoint_configuration_vpc_endpoint_ids|text[]|removed|
|mutual_tls_authentication|jsonb|added|
|mutual_tls_authentication_truststore_uri|text|removed|
|mutual_tls_authentication_truststore_version|text|removed|
|mutual_tls_authentication_truststore_warnings|text[]|removed|

## aws_apigateway_rest_api_authorizers

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|account_id|text|added|
|provider_ar_ns|text[]|added|
|provider_arns|text[]|removed|
|region|text|added|
|rest_api_arn|text|added|
|rest_api_cq_id|uuid|removed|
|rest_api_id|text|removed|

## aws_apigateway_rest_api_deployments

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|account_id|text|added|
|region|text|added|
|rest_api_arn|text|added|
|rest_api_cq_id|uuid|removed|
|rest_api_id|text|removed|

## aws_apigateway_rest_api_documentation_parts

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|account_id|text|added|
|location|jsonb|added|
|location_method|text|removed|
|location_name|text|removed|
|location_path|text|removed|
|location_status_code|text|removed|
|location_type|text|removed|
|region|text|added|
|rest_api_arn|text|added|
|rest_api_cq_id|uuid|removed|
|rest_api_id|text|removed|

## aws_apigateway_rest_api_documentation_versions

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|account_id|text|added|
|region|text|added|
|rest_api_arn|text|added|
|rest_api_cq_id|uuid|removed|
|rest_api_id|text|removed|

## aws_apigateway_rest_api_gateway_responses

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|account_id|text|added|
|region|text|added|
|rest_api_arn|text|added|
|rest_api_cq_id|uuid|removed|
|rest_api_id|text|removed|

## aws_apigateway_rest_api_models

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|account_id|text|added|
|region|text|added|
|rest_api_arn|text|added|
|rest_api_cq_id|uuid|removed|
|rest_api_id|text|removed|

## aws_apigateway_rest_api_request_validators

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|account_id|text|added|
|region|text|added|
|rest_api_arn|text|added|
|rest_api_cq_id|uuid|removed|
|rest_api_id|text|removed|

## aws_apigateway_rest_api_resources

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|account_id|text|added|
|region|text|added|
|rest_api_arn|text|added|
|rest_api_cq_id|uuid|removed|
|rest_api_id|text|removed|

## aws_apigateway_rest_api_stages

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|access_log_settings|jsonb|added|
|access_log_settings_destination_arn|text|removed|
|access_log_settings_format|text|removed|
|account_id|text|added|
|canary_settings|jsonb|added|
|canary_settings_deployment_id|text|removed|
|canary_settings_percent_traffic|float|removed|
|canary_settings_stage_variable_overrides|jsonb|removed|
|canary_settings_use_stage_cache|boolean|removed|
|region|text|added|
|rest_api_arn|text|added|
|rest_api_cq_id|uuid|removed|
|rest_api_id|text|removed|

## aws_apigateway_rest_apis

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|endpoint_configuration|jsonb|added|
|endpoint_configuration_types|text[]|removed|
|endpoint_configuration_vpc_endpoint_ids|text[]|removed|

## aws_apigateway_usage_plan_api_stages
Moved to JSON column on [aws_apigateway_usage_plans](#aws_apigateway_usage_plans)


## aws_apigateway_usage_plan_keys

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|account_id|text|added|
|region|text|added|
|usage_plan_arn|text|added|
|usage_plan_cq_id|uuid|removed|
|usage_plan_id|text|removed|

## aws_apigateway_usage_plans

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|api_stages|jsonb|added|
|quota|jsonb|added|
|quota_limit|bigint|removed|
|quota_offset|bigint|removed|
|quota_period|text|removed|
|throttle|jsonb|added|
|throttle_burst_limit|bigint|removed|
|throttle_rate_limit|float|removed|

## aws_apigateway_vpc_links

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|

## aws_apigatewayv2_api_authorizers

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|account_id|text|added|
|api_arn|text|added|
|api_cq_id|uuid|removed|
|authorizer_result_ttl_in_seconds|bigint|updated|Type changed from integer to bigint
|jwt_configuration|jsonb|added|
|jwt_configuration_audience|text[]|removed|
|jwt_configuration_issuer|text|removed|
|region|text|added|

## aws_apigatewayv2_api_deployments

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|account_id|text|added|
|api_arn|text|added|
|api_cq_id|uuid|removed|
|region|text|added|

## aws_apigatewayv2_api_integration_responses

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|account_id|text|added|
|api_integration_arn|text|added|
|api_integration_cq_id|uuid|removed|
|region|text|added|

## aws_apigatewayv2_api_integrations

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|account_id|text|added|
|api_arn|text|added|
|api_cq_id|uuid|removed|
|region|text|added|
|timeout_in_millis|bigint|updated|Type changed from integer to bigint
|tls_config|jsonb|added|
|tls_config_server_name_to_verify|text|removed|

## aws_apigatewayv2_api_models

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|account_id|text|added|
|api_arn|text|added|
|api_cq_id|uuid|removed|
|region|text|added|

## aws_apigatewayv2_api_route_responses

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|account_id|text|added|
|api_route_arn|text|added|
|api_route_cq_id|uuid|removed|
|region|text|added|

## aws_apigatewayv2_api_routes

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|account_id|text|added|
|api_arn|text|added|
|api_cq_id|uuid|removed|
|region|text|added|

## aws_apigatewayv2_api_stages

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|access_log_settings|jsonb|added|
|access_log_settings_destination_arn|text|removed|
|access_log_settings_format|text|removed|
|account_id|text|added|
|api_arn|text|added|
|api_cq_id|uuid|removed|
|default_route_settings|jsonb|added|
|region|text|added|
|route_settings_data_trace_enabled|boolean|removed|
|route_settings_detailed_metrics_enabled|boolean|removed|
|route_settings_logging_level|text|removed|
|route_settings_throttling_burst_limit|integer|removed|
|route_settings_throttling_rate_limit|float|removed|

## aws_apigatewayv2_apis

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|cors_configuration|jsonb|added|
|cors_configuration_allow_credentials|boolean|removed|
|cors_configuration_allow_headers|text[]|removed|
|cors_configuration_allow_methods|text[]|removed|
|cors_configuration_allow_origins|text[]|removed|
|cors_configuration_expose_headers|text[]|removed|
|cors_configuration_max_age|integer|removed|

## aws_apigatewayv2_domain_name_configurations
Moved to JSON column on [aws_apigatewayv2_domain_names](#aws_apigatewayv2_domain_names)


## aws_apigatewayv2_domain_name_rest_api_mappings

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|account_id|text|added|
|domain_name_arn|text|added|
|domain_name_cq_id|uuid|removed|
|region|text|added|

## aws_apigatewayv2_domain_names

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|domain_name_configurations|jsonb|added|
|mutual_tls_authentication|jsonb|added|
|mutual_tls_authentication_truststore_uri|text|removed|
|mutual_tls_authentication_truststore_version|text|removed|
|mutual_tls_authentication_truststore_warnings|text[]|removed|

## aws_apigatewayv2_vpc_links

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|id|text|removed|
|vpc_link_id|text|added|

## aws_applicationautoscaling_policies

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|name|text|removed|
|namespace|text|removed|
|policy_name|text|added|
|policy_type|text|added|
|type|text|removed|

## aws_appsync_graphql_api_additional_authentication_providers
Moved to JSON column on [aws_appsync_graphql_apis](#aws_appsync_graphql_apis)


## aws_appsync_graphql_apis

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|additional_authentication_providers|jsonb|added|
|api_id|text|added|
|id|text|removed|
|lambda_authorizer_config|jsonb|added|
|lambda_authorizer_config_authorizer_result_ttl_in_seconds|bigint|removed|
|lambda_authorizer_config_authorizer_uri|text|removed|
|lambda_authorizer_config_identity_validation_expression|text|removed|
|log_config|jsonb|added|
|log_config_cloud_watch_logs_role_arn|text|removed|
|log_config_exclude_verbose_content|boolean|removed|
|log_config_field_log_level|text|removed|
|open_id_connect_config|jsonb|added|
|open_id_connect_config_auth_ttl|bigint|removed|
|open_id_connect_config_client_id|text|removed|
|open_id_connect_config_iat_ttl|bigint|removed|
|open_id_connect_config_issuer|text|removed|
|user_pool_config|jsonb|added|
|user_pool_config_app_id_client_regex|text|removed|
|user_pool_config_aws_region|text|removed|
|user_pool_config_default_action|text|removed|
|user_pool_config_user_pool_id|text|removed|

## aws_athena_data_catalog_database_table_columns
Moved to JSON column on [aws_athena_data_catalogs](#aws_athena_data_catalogs)


## aws_athena_data_catalog_database_table_partition_keys
Moved to JSON column on [aws_athena_data_catalogs](#aws_athena_data_catalogs)


## aws_athena_data_catalog_database_tables

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|account_id|text|added|
|columns|jsonb|added|
|data_catalog_arn|text|added|
|data_catalog_database_cq_id|uuid|removed|
|data_catalog_database_name|text|added|
|partition_keys|jsonb|added|
|region|text|added|

## aws_athena_data_catalog_databases

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|account_id|text|added|
|data_catalog_arn|text|added|
|data_catalog_cq_id|uuid|removed|
|region|text|added|

## aws_athena_data_catalogs

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|

## aws_athena_work_group_named_queries

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|account_id|text|added|
|region|text|added|
|work_group_arn|text|added|
|work_group_cq_id|uuid|removed|

## aws_athena_work_group_prepared_statements

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|account_id|text|added|
|region|text|added|
|work_group_arn|text|added|
|work_group_cq_id|uuid|removed|

## aws_athena_work_group_query_executions

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|account_id|text|added|
|acl_configuration_s3_acl_option|text|removed|
|athena_error_error_category|bigint|removed|
|athena_error_error_message|text|removed|
|athena_error_error_type|bigint|removed|
|athena_error_retryable|boolean|removed|
|catalog|text|removed|
|completion_date_time|timestamp without time zone|removed|
|data_manifest_location|text|removed|
|data_scanned_in_bytes|bigint|removed|
|database|text|removed|
|effective_engine_version|text|removed|
|encryption_configuration_encryption_option|text|removed|
|encryption_configuration_kms_key|text|removed|
|engine_execution_time_in_millis|bigint|removed|
|engine_version|jsonb|added|
|expected_bucket_owner|text|removed|
|id|text|removed|
|output_location|text|removed|
|query_execution_context|jsonb|added|
|query_execution_id|text|added|
|query_planning_time_in_millis|bigint|removed|
|query_queue_time_in_millis|bigint|removed|
|region|text|added|
|result_configuration|jsonb|added|
|selected_engine_version|text|removed|
|service_processing_time_in_millis|bigint|removed|
|state|text|removed|
|state_change_reason|text|removed|
|statistics|jsonb|added|
|status|jsonb|added|
|submission_date_time|timestamp without time zone|removed|
|total_execution_time_in_millis|bigint|removed|
|work_group_arn|text|added|
|work_group_cq_id|uuid|removed|

## aws_athena_work_groups

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|acl_configuration_s3_acl_option|text|removed|
|bytes_scanned_cutoff_per_query|bigint|removed|
|configuration|jsonb|added|
|effective_engine_version|text|removed|
|encryption_configuration_encryption_option|text|removed|
|encryption_configuration_kms_key|text|removed|
|enforce_work_group_configuration|boolean|removed|
|expected_bucket_owner|text|removed|
|output_location|text|removed|
|publish_cloud_watch_metrics_enabled|boolean|removed|
|requester_pays_enabled|boolean|removed|
|selected_engine_version|text|removed|

## aws_autoscaling_group_instances
Moved to JSON column on [aws_autoscaling_groups](#aws_autoscaling_groups)


## aws_autoscaling_group_lifecycle_hooks

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|account_id|text|added|
|global_timeout|bigint|updated|Type changed from integer to bigint
|group_arn|text|added|
|group_cq_id|uuid|removed|
|heartbeat_timeout|bigint|updated|Type changed from integer to bigint
|region|text|added|

## aws_autoscaling_group_scaling_policies

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|account_id|text|added|
|arn|text|removed|
|cooldown|bigint|updated|Type changed from integer to bigint
|estimated_instance_warmup|bigint|updated|Type changed from integer to bigint
|group_arn|text|added|
|group_cq_id|uuid|removed|
|min_adjustment_magnitude|bigint|updated|Type changed from integer to bigint
|min_adjustment_step|bigint|updated|Type changed from integer to bigint
|name|text|removed|
|policy_arn|text|added|
|policy_name|text|added|
|policy_type|text|added|
|predictive_scaling_configuration|jsonb|added|
|region|text|added|
|scaling_adjustment|bigint|updated|Type changed from integer to bigint
|target_tracking_configuration|jsonb|added|
|target_tracking_configuration_customized_metric_dimensions|jsonb|removed|
|target_tracking_configuration_customized_metric_name|text|removed|
|target_tracking_configuration_customized_metric_namespace|text|removed|
|target_tracking_configuration_customized_metric_statistic|text|removed|
|target_tracking_configuration_customized_metric_unit|text|removed|
|target_tracking_configuration_disable_scale_in|boolean|removed|
|target_tracking_configuration_predefined_metric_resource_label|text|removed|
|target_tracking_configuration_predefined_metric_type|text|removed|
|target_tracking_configuration_target_value|float|removed|
|type|text|removed|

## aws_autoscaling_group_tags
Moved to JSON column on [aws_autoscaling_groups](#aws_autoscaling_groups)


## aws_autoscaling_groups

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|auto_scaling_group_name|text|added|
|context|text|added|
|default_cooldown|bigint|updated|Type changed from integer to bigint
|default_instance_warmup|bigint|added|
|desired_capacity|bigint|updated|Type changed from integer to bigint
|desired_capacity_type|text|added|
|health_check_grace_period|bigint|updated|Type changed from integer to bigint
|instances|jsonb|added|
|launch_template|jsonb|added|
|launch_template_id|text|removed|
|launch_template_name|text|removed|
|launch_template_version|text|removed|
|max_instance_lifetime|bigint|updated|Type changed from integer to bigint
|max_size|bigint|updated|Type changed from integer to bigint
|min_size|bigint|updated|Type changed from integer to bigint
|name|text|removed|
|notification_configurations|jsonb|added|
|notifications_configurations|jsonb|removed|
|predicted_capacity|bigint|added|
|tags|jsonb|added|
|target_group_ar_ns|text[]|added|
|target_group_arns|text[]|removed|
|warm_pool_configuration|jsonb|added|
|warm_pool_size|bigint|added|

## aws_autoscaling_launch_configuration_block_device_mappings
Moved to JSON column on [aws_autoscaling_launch_configurations](#aws_autoscaling_launch_configurations)


## aws_autoscaling_launch_configurations

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|block_device_mappings|jsonb|added|
|instance_monitoring|jsonb|added|
|instance_monitoring_enabled|boolean|removed|
|metadata_options|jsonb|added|
|metadata_options_http_endpoint|text|removed|
|metadata_options_http_put_response_hop_limit|integer|removed|
|metadata_options_http_tokens|text|removed|

## aws_autoscaling_scheduled_actions

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|desired_capacity|bigint|updated|Type changed from integer to bigint
|max_size|bigint|updated|Type changed from integer to bigint
|min_size|bigint|updated|Type changed from integer to bigint
|name|text|removed|
|scheduled_action_name|text|added|

## aws_backup_global_settings

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|result_metadata|jsonb|added|

## aws_backup_plan_rules
Moved to JSON column on [aws_backup_plans](#aws_backup_plans)


## aws_backup_plan_selections

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|account_id|text|added|
|backup_plan_id|text|added|
|backup_selection|jsonb|added|
|conditions|jsonb|removed|
|iam_role_arn|text|removed|
|list_of_tags|jsonb|removed|
|not_resources|text[]|removed|
|plan_arn|text|added|
|plan_cq_id|uuid|removed|
|region|text|added|
|resources|text[]|removed|
|result_metadata|jsonb|added|
|selection_name|text|removed|

## aws_backup_plans

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|backup_plan|jsonb|added|
|backup_plan_id|text|added|
|deletion_date|timestamp without time zone|added|
|id|text|removed|
|name|text|removed|
|result_metadata|jsonb|added|

## aws_backup_region_settings

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|result_metadata|jsonb|added|

## aws_backup_vault_recovery_points

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|account_id|text|added|
|arn|text|removed|
|backup_size|bigint|removed|
|backup_size_in_bytes|bigint|added|
|backup_vault_arn|text|added|
|backup_vault_name|text|added|
|calculated_delete_at|timestamp without time zone|removed|
|calculated_lifecycle|jsonb|added|
|calculated_move_to_cold_storage_at|timestamp without time zone|removed|
|delete_after|bigint|removed|
|lifecycle|jsonb|added|
|move_to_cold_storage_after|bigint|removed|
|recovery_point_arn|text|added|
|region|text|added|
|vault_arn|text|added|
|vault_cq_id|uuid|removed|

## aws_backup_vaults

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|backup_vault_name|text|added|
|name|text|removed|
|notification_events|text[]|removed|
|notification_sns_topic_arn|text|removed|
|notifications|jsonb|added|

## aws_cloudformation_stack_outputs
Moved to JSON column on [aws_cloudformation_stacks](#aws_cloudformation_stacks)


## aws_cloudformation_stack_resources

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|account_id|text|added|
|drift_information|jsonb|added|
|drift_last_check_timestamp|timestamp without time zone|removed|
|module_info|jsonb|added|
|module_info_logical_id_hierarchy|text|removed|
|module_info_type_hierarchy|text|removed|
|region|text|added|
|stack_cq_id|uuid|removed|
|stack_resource_drift_status|text|removed|

## aws_cloudformation_stacks

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|drift_information|jsonb|added|
|drift_last_check_timestamp|timestamp without time zone|removed|
|notification_ar_ns|text[]|added|
|notification_arns|text[]|removed|
|outputs|jsonb|added|
|rollback_configuration|jsonb|added|
|rollback_configuration_monitoring_time_in_minutes|integer|removed|
|rollback_configuration_rollback_triggers|jsonb|removed|
|stack|text|removed|
|stack_drift_status|text|removed|
|stack_name|text|added|
|stack_status|text|added|
|status|text|removed|
|timeout_in_minutes|bigint|updated|Type changed from integer to bigint

## aws_cloudfront_cache_policies

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|cache_policy|jsonb|added|
|comment|text|removed|
|cookies|text[]|removed|
|cookies_behavior|text|removed|
|cookies_quantity|integer|removed|
|default_ttl|bigint|removed|
|enable_accept_encoding_brotli|boolean|removed|
|enable_accept_encoding_gzip|boolean|removed|
|headers|text[]|removed|
|headers_behavior|text|removed|
|headers_quantity|integer|removed|
|id|text|removed|
|last_modified_time|timestamp without time zone|removed|
|max_ttl|bigint|removed|
|min_ttl|bigint|removed|
|name|text|removed|
|query_strings|text[]|removed|
|query_strings_behavior|text|removed|
|query_strings_quantity|integer|removed|

## aws_cloudfront_distribution_cache_behavior_lambda_functions
Moved to JSON column on [aws_cloudfront_distributions](#aws_cloudfront_distributions)


## aws_cloudfront_distribution_cache_behaviors
Moved to JSON column on [aws_cloudfront_distributions](#aws_cloudfront_distributions)


## aws_cloudfront_distribution_custom_error_responses
Moved to JSON column on [aws_cloudfront_distributions](#aws_cloudfront_distributions)


## aws_cloudfront_distribution_default_cache_behavior_functions
Moved to JSON column on [aws_cloudfront_distributions](#aws_cloudfront_distributions)


## aws_cloudfront_distribution_origin_groups
Moved to JSON column on [aws_cloudfront_distributions](#aws_cloudfront_distributions)


## aws_cloudfront_distribution_origins
Moved to JSON column on [aws_cloudfront_distributions](#aws_cloudfront_distributions)


## aws_cloudfront_distributions

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|active_trusted_key_groups_enabled|boolean|removed|
|active_trusted_signers_enabled|boolean|removed|
|aliases|text[]|removed|
|cache_behavior_allowed_methods|text[]|removed|
|cache_behavior_allowed_methods_cached_methods|text[]|removed|
|cache_behavior_cache_policy_id|text|removed|
|cache_behavior_compress|boolean|removed|
|cache_behavior_default_ttl|bigint|removed|
|cache_behavior_field_level_encryption_id|text|removed|
|cache_behavior_forwarded_values_cookies_forward|text|removed|
|cache_behavior_forwarded_values_cookies_whitelisted_names|text[]|removed|
|cache_behavior_forwarded_values_headers|text[]|removed|
|cache_behavior_forwarded_values_query_string|boolean|removed|
|cache_behavior_forwarded_values_query_string_cache_keys|text[]|removed|
|cache_behavior_max_ttl|bigint|removed|
|cache_behavior_min_ttl|bigint|removed|
|cache_behavior_origin_request_policy_id|text|removed|
|cache_behavior_realtime_log_config_arn|text|removed|
|cache_behavior_smooth_streaming|boolean|removed|
|cache_behavior_target_origin_id|text|removed|
|cache_behavior_trusted_key_groups|text[]|removed|
|cache_behavior_trusted_key_groups_enabled|boolean|removed|
|cache_behavior_trusted_signers|text[]|removed|
|cache_behavior_trusted_signers_enabled|boolean|removed|
|cache_behavior_viewer_protocol_policy|text|removed|
|caller_reference|text|removed|
|comment|text|removed|
|default_root_object|text|removed|
|distribution_config|jsonb|added|
|enabled|boolean|removed|
|geo_restriction_type|text|removed|
|geo_restrictions|text[]|removed|
|http_version|text|removed|
|in_progress_invalidation_batches|bigint|updated|Type changed from integer to bigint
|ipv6_enabled|boolean|removed|
|logging_bucket|text|removed|
|logging_enabled|boolean|removed|
|logging_include_cookies|boolean|removed|
|logging_prefix|text|removed|
|price_class|text|removed|
|viewer_certificate|text|removed|
|viewer_certificate_acm_certificate_arn|text|removed|
|viewer_certificate_cloudfront_default_certificate|boolean|removed|
|viewer_certificate_iam_certificate_id|text|removed|
|viewer_certificate_minimum_protocol_version|text|removed|
|viewer_certificate_source|text|removed|
|viewer_certificate_ssl_support_method|text|removed|
|web_acl_id|text|removed|

## aws_cloudtrail_trail_event_selectors

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|account_id|text|added|
|data_resources|jsonb|added|
|region|text|added|
|trail_cq_id|uuid|removed|

## aws_cloudtrail_trails

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|home_region|text|added|
|is_logging|boolean|removed|
|latest_cloud_watch_logs_delivery_error|text|removed|
|latest_cloud_watch_logs_delivery_time|timestamp without time zone|removed|
|latest_delivery_error|text|removed|
|latest_delivery_time|timestamp without time zone|removed|
|latest_digest_delivery_error|text|removed|
|latest_digest_delivery_time|timestamp without time zone|removed|
|latest_notification_error|text|removed|
|latest_notification_time|timestamp without time zone|removed|
|start_logging_time|timestamp without time zone|removed|
|status|jsonb|added|
|stop_logging_time|timestamp without time zone|removed|
|tags|jsonb|removed|

## aws_cloudwatch_alarm_metrics
Moved to JSON column on [aws_cloudwatch_alarms](#aws_cloudwatch_alarms)


## aws_cloudwatch_alarms

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|actions|text[]|removed|
|alarm_actions|text[]|added|
|alarm_configuration_updated_timestamp|timestamp without time zone|added|
|alarm_description|text|added|
|alarm_name|text|added|
|configuration_updated_timestamp|timestamp without time zone|removed|
|datapoints_to_alarm|bigint|updated|Type changed from integer to bigint
|description|text|removed|
|evaluation_periods|bigint|updated|Type changed from integer to bigint
|metrics|jsonb|added|
|name|text|removed|
|period|bigint|updated|Type changed from integer to bigint
|threshold|real|updated|Type changed from float to real

## aws_cloudwatchlogs_filter_metric_transformations
This table was removed.


## aws_cloudwatchlogs_filters
This table was removed.


## aws_cloudwatchlogs_log_groups

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|

## aws_cloudwatchlogs_metric_filters
This table was newly added.

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|account_id|text|added|
|region|text|added|
|arn|text|added|
|creation_time|bigint|added|
|filter_name|text|added|
|filter_pattern|text|added|
|log_group_name|text|added|
|metric_transformations|jsonb|added|
|_cq_id|uuid|added|
|_cq_fetch_time|timestamp without time zone|added|

## aws_codebuild_project_environment_variables
Moved to JSON column on [aws_codebuild_projects](#aws_codebuild_projects)


## aws_codebuild_project_file_system_locations
Moved to JSON column on [aws_codebuild_projects](#aws_codebuild_projects)


## aws_codebuild_project_secondary_artifacts
Moved to JSON column on [aws_codebuild_projects](#aws_codebuild_projects)


## aws_codebuild_project_secondary_sources
Moved to JSON column on [aws_codebuild_projects](#aws_codebuild_projects)


## aws_codebuild_projects

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|artifacts|jsonb|added|
|artifacts_artifact_identifier|text|removed|
|artifacts_bucket_owner_access|text|removed|
|artifacts_encryption_disabled|boolean|removed|
|artifacts_location|text|removed|
|artifacts_name|text|removed|
|artifacts_namespace_type|text|removed|
|artifacts_override_artifact_name|boolean|removed|
|artifacts_packaging|text|removed|
|artifacts_path|text|removed|
|artifacts_type|text|removed|
|badge|jsonb|added|
|badge_enabled|boolean|removed|
|badge_request_url|text|removed|
|build_batch_config|jsonb|added|
|build_batch_config_batch_report_mode|text|removed|
|build_batch_config_combine_artifacts|boolean|removed|
|build_batch_config_restrictions_compute_types_allowed|text[]|removed|
|build_batch_config_restrictions_maximum_builds_allowed|integer|removed|
|build_batch_config_service_role|text|removed|
|build_batch_config_timeout_in_mins|integer|removed|
|cache|jsonb|added|
|cache_location|text|removed|
|cache_modes|text[]|removed|
|cache_type|text|removed|
|concurrent_build_limit|bigint|updated|Type changed from integer to bigint
|environment|jsonb|added|
|environment_certificate|text|removed|
|environment_compute_type|text|removed|
|environment_image|text|removed|
|environment_image_pull_credentials_type|text|removed|
|environment_privileged_mode|boolean|removed|
|environment_registry_credential|text|removed|
|environment_registry_credential_credential_provider|text|removed|
|environment_type|text|removed|
|file_system_locations|jsonb|added|
|logs_config|jsonb|added|
|logs_config_cloud_watch_logs_group_name|text|removed|
|logs_config_cloud_watch_logs_status|text|removed|
|logs_config_cloud_watch_logs_stream_name|text|removed|
|logs_config_s3_logs_bucket_owner_access|text|removed|
|logs_config_s3_logs_encryption_disabled|boolean|removed|
|logs_config_s3_logs_location|text|removed|
|logs_config_s3_logs_status|text|removed|
|queued_timeout_in_minutes|bigint|updated|Type changed from integer to bigint
|secondary_artifacts|jsonb|added|
|secondary_sources|jsonb|added|
|source|jsonb|added|
|source_auth_resource|text|removed|
|source_auth_type|text|removed|
|source_build_status_config_context|text|removed|
|source_build_status_config_target_url|text|removed|
|source_buildspec|text|removed|
|source_git_clone_depth|integer|removed|
|source_git_submodules_config_fetch_submodules|boolean|removed|
|source_identifier|text|removed|
|source_insecure_ssl|boolean|removed|
|source_location|text|removed|
|source_report_build_status|boolean|removed|
|source_type|text|removed|
|timeout_in_minutes|bigint|updated|Type changed from integer to bigint
|vpc_config|jsonb|added|
|vpc_config_security_group_ids|text[]|removed|
|vpc_config_subnets|text[]|removed|
|vpc_config_vpc_id|text|removed|
|webhook|jsonb|added|
|webhook_branch_filter|text|removed|
|webhook_build_type|text|removed|
|webhook_filter_groups|jsonb|removed|
|webhook_last_modified_secret|timestamp without time zone|removed|
|webhook_payload_url|text|removed|
|webhook_secret|text|removed|
|webhook_url|text|removed|

## aws_codepipeline_pipeline_stage_actions
Moved to JSON column on [aws_codepipeline_pipelines](#aws_codepipeline_pipelines)


## aws_codepipeline_pipeline_stages
Moved to JSON column on [aws_codepipeline_pipelines](#aws_codepipeline_pipelines)


## aws_codepipeline_pipelines

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|artifact_store_encryption_key_id|text|removed|
|artifact_store_encryption_key_type|text|removed|
|artifact_store_location|text|removed|
|artifact_store_type|text|removed|
|artifact_stores|jsonb|removed|
|created|timestamp without time zone|removed|
|metadata|jsonb|added|
|name|text|removed|
|pipeline|jsonb|added|
|role_arn|text|removed|
|tags|text|updated|Type changed from jsonb to text
|updated|timestamp without time zone|removed|
|version|bigint|removed|

## aws_codepipeline_webhook_filters
Moved to JSON column on [aws_codepipeline_webhooks](#aws_codepipeline_webhooks)


## aws_codepipeline_webhooks

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|authentication|text|removed|
|authentication_allowed_ip_range|text|removed|
|authentication_secret_token|text|removed|
|definition|jsonb|added|
|name|text|removed|
|target_action|text|removed|
|target_pipeline|text|removed|

## aws_cognito_identity_pool_cognito_identity_providers
Moved to JSON column on [aws_cognito_identity_pools](#aws_cognito_identity_pools)


## aws_cognito_identity_pools

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|cognito_identity_providers|jsonb|added|
|open_id_connect_provider_ar_ns|text[]|added|
|open_id_connect_provider_arns|text[]|removed|
|result_metadata|jsonb|added|
|saml_provider_ar_ns|text[]|added|
|saml_provider_arns|text[]|removed|

## aws_cognito_user_pool_identity_providers

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|user_pool_arn|text|added|
|user_pool_cq_id|uuid|removed|

## aws_cognito_user_pool_schema_attributes
Moved to JSON column on [aws_cognito_user_pools](#aws_cognito_user_pools)


## aws_cognito_user_pools

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|admin_create_user_admin_only|boolean|removed|
|admin_create_user_config|jsonb|added|
|admin_create_user_config_unused_account_validity_days|integer|removed|
|admin_create_user_invite_email_message|text|removed|
|admin_create_user_invite_email_subject|text|removed|
|admin_create_user_invite_sms|text|removed|
|challenge_required_on_new_device|boolean|removed|
|device_configuration|jsonb|added|
|device_only_remembered_on_user_prompt|boolean|removed|
|email_configuration|jsonb|added|
|email_configuration_from|text|removed|
|email_configuration_reply_to_address|text|removed|
|email_configuration_sending_account|text|removed|
|email_configuration_set|text|removed|
|email_configuration_source_arn|text|removed|
|estimated_number_of_users|bigint|updated|Type changed from integer to bigint
|lambda_config|jsonb|added|
|lambda_config_create_auth_challenge|text|removed|
|lambda_config_custom_email_sender_lambda_arn|text|removed|
|lambda_config_custom_email_sender_lambda_version|text|removed|
|lambda_config_custom_message|text|removed|
|lambda_config_custom_sms_sender_lambda_arn|text|removed|
|lambda_config_custom_sms_sender_lambda_version|text|removed|
|lambda_config_define_auth_challenge|text|removed|
|lambda_config_kms_key_id|text|removed|
|lambda_config_post_authentication|text|removed|
|lambda_config_post_confirmation|text|removed|
|lambda_config_pre_authentication|text|removed|
|lambda_config_pre_sign_up|text|removed|
|lambda_config_pre_token_generation|text|removed|
|lambda_config_user_migration|text|removed|
|lambda_config_verify_auth_challenge_response|text|removed|
|policies|jsonb|added|
|policies_password_policy_minimum_length|integer|removed|
|policies_password_policy_require_lowercase|boolean|removed|
|policies_password_policy_require_numbers|boolean|removed|
|policies_password_policy_require_symbols|boolean|removed|
|policies_password_policy_require_uppercase|boolean|removed|
|policies_password_policy_temporary_password_validity_days|integer|removed|
|schema_attributes|jsonb|added|
|sms_configuration|jsonb|added|
|sms_configuration_external_id|text|removed|
|sms_configuration_sns_caller_arn|text|removed|
|user_attribute_update_settings|jsonb|added|
|user_pool_add_ons|jsonb|added|
|user_pool_add_ons_advanced_security_mode|text|removed|
|username_configuration|jsonb|added|
|username_configuration_case_sensitive|boolean|removed|
|verification_message_template|jsonb|added|
|verification_message_template_default_email_option|text|removed|
|verification_message_template_email_message|text|removed|
|verification_message_template_email_message_by_link|text|removed|
|verification_message_template_email_subject|text|removed|
|verification_message_template_email_subject_by_link|text|removed|
|verification_message_template_sms_message|text|removed|

## aws_config_configuration_recorders

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|recording_group|jsonb|added|
|recording_group_all_supported|boolean|removed|
|recording_group_include_global_resource_types|boolean|removed|
|recording_group_resource_types|text[]|removed|

## aws_config_conformance_pack_rule_compliances

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|account_id|text|added|
|conformance_pack_arn|text|added|
|conformance_pack_cq_id|uuid|removed|
|evaluation_result_identifier|jsonb|added|
|ordering_timestamp|timestamp without time zone|removed|
|region|text|added|
|resource_id|text|removed|
|resource_type|text|removed|

## aws_config_conformance_packs

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|conformance_pack_input_parameters|jsonb|removed|

## aws_dax_cluster_nodes
Moved to JSON column on [aws_dax_clusters](#aws_dax_clusters)


## aws_dax_clusters

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|active_nodes|bigint|updated|Type changed from integer to bigint
|cluster_discovery_endpoint|jsonb|added|
|cluster_discovery_endpoint_address|text|removed|
|cluster_discovery_endpoint_port|integer|removed|
|cluster_discovery_endpoint_url|text|removed|
|cluster_name|text|added|
|name|text|removed|
|node_ids_to_reboot|text[]|removed|
|nodes|jsonb|added|
|notification_configuration|jsonb|added|
|notification_configuration_topic_arn|text|removed|
|notification_configuration_topic_status|text|removed|
|parameter_apply_status|text|removed|
|parameter_group|jsonb|added|
|parameter_group_name|text|removed|
|sse_description|jsonb|added|
|sse_description_status|text|removed|
|tags|text|updated|Type changed from jsonb to text
|total_nodes|bigint|updated|Type changed from integer to bigint

## aws_directconnect_connection_mac_sec_keys
Moved to JSON column on [aws_directconnect_connections](#aws_directconnect_connections)


## aws_directconnect_connections

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|aws_device|text|added|
|aws_logical_device_id|text|added|
|connection_name|text|added|
|mac_sec_keys|jsonb|added|
|name|text|removed|
|vlan|bigint|updated|Type changed from integer to bigint

## aws_directconnect_gateway_associations

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|account_id|text|added|
|allowed_prefixes_to_direct_connect_gateway|jsonb|updated|Type changed from text[] to jsonb
|associated_gateway|jsonb|added|
|associated_gateway_id|text|removed|
|associated_gateway_owner_account|text|removed|
|associated_gateway_region|text|removed|
|associated_gateway_type|text|removed|
|direct_connect_gateway_id|text|added|
|gateway_arn|text|added|
|gateway_cq_id|uuid|removed|
|region|text|added|
|resource_id|text|removed|
|virtual_gateway_region|text|added|

## aws_directconnect_gateway_attachments

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|account_id|text|added|
|direct_connect_gateway_id|text|added|
|gateway_arn|text|added|
|gateway_cq_id|uuid|removed|
|region|text|added|

## aws_directconnect_gateways

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|direct_connect_gateway_name|text|added|
|direct_connect_gateway_state|text|added|
|name|text|removed|
|region|text|added|
|state|text|removed|

## aws_directconnect_lag_mac_sec_keys
Moved to JSON column on [aws_directconnect_lags](#aws_directconnect_lags)


## aws_directconnect_lags

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|aws_device|text|added|
|aws_logical_device_id|text|added|
|connection_ids|text[]|removed|
|connections|jsonb|added|
|lag_name|text|added|
|lag_state|text|added|
|mac_sec_keys|jsonb|added|
|minimum_links|bigint|updated|Type changed from integer to bigint
|name|text|removed|
|number_of_connections|bigint|updated|Type changed from integer to bigint
|state|text|removed|

## aws_directconnect_virtual_gateways

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|state|text|removed|
|virtual_gateway_state|text|added|

## aws_directconnect_virtual_interface_bgp_peers
Moved to JSON column on [aws_directconnect_virtual_interfaces](#aws_directconnect_virtual_interfaces)


## aws_directconnect_virtual_interfaces

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|asn|bigint|updated|Type changed from integer to bigint
|aws_logical_device_id|text|added|
|bgp_peers|jsonb|added|
|mtu|bigint|updated|Type changed from integer to bigint
|route_filter_prefixes|jsonb|updated|Type changed from text[] to jsonb
|site_link_enabled|boolean|added|
|vlan|bigint|updated|Type changed from integer to bigint

## aws_dms_replication_instance_replication_subnet_group_subnets
Moved to JSON column on [aws_dms_replication_instances](#aws_dms_replication_instances)


## aws_dms_replication_instance_vpc_security_groups
Moved to JSON column on [aws_dms_replication_instances](#aws_dms_replication_instances)


## aws_dms_replication_instances

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|allocated_storage|bigint|updated|Type changed from integer to bigint
|class|text|removed|
|identifier|text|removed|
|pending_modified_values|jsonb|added|
|pending_modified_values_allocated_storage|integer|removed|
|pending_modified_values_class|text|removed|
|pending_modified_values_engine_version|text|removed|
|pending_modified_values_multi_az|boolean|removed|
|private_ip_address|inet|removed|
|private_ip_addresses|inet[]|removed|
|public_ip_address|inet|removed|
|public_ip_addresses|inet[]|removed|
|replication_instance_class|text|added|
|replication_instance_identifier|text|added|
|replication_instance_private_ip_address|text|added|
|replication_instance_private_ip_addresses|text[]|added|
|replication_instance_public_ip_address|text|added|
|replication_instance_public_ip_addresses|text[]|added|
|replication_instance_status|text|added|
|replication_subnet_group|jsonb|added|
|replication_subnet_group_description|text|removed|
|replication_subnet_group_identifier|text|removed|
|replication_subnet_group_subnet_group_status|text|removed|
|replication_subnet_group_vpc_id|text|removed|
|status|text|removed|
|vpc_security_groups|jsonb|added|

## aws_dynamodb_table_continuous_backups

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|account_id|text|added|
|earliest_restorable_date_time|timestamp without time zone|removed|
|latest_restorable_date_time|timestamp without time zone|removed|
|point_in_time_recovery_description|jsonb|added|
|point_in_time_recovery_status|text|removed|
|region|text|added|
|table_arn|text|added|
|table_cq_id|uuid|removed|

## aws_dynamodb_table_global_secondary_indexes
Moved to JSON column on [aws_dynamodb_tables](#aws_dynamodb_tables)


## aws_dynamodb_table_local_secondary_indexes
Moved to JSON column on [aws_dynamodb_tables](#aws_dynamodb_tables)


## aws_dynamodb_table_replica_auto_scalings

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|account_id|text|added|
|read_capacity|jsonb|removed|
|region|text|added|
|replica_provisioned_read_capacity_auto_scaling_settings|jsonb|added|
|replica_provisioned_write_capacity_auto_scaling_settings|jsonb|added|
|table_arn|text|added|
|table_cq_id|uuid|removed|
|write_capacity|jsonb|removed|

## aws_dynamodb_table_replicas
Moved to JSON column on [aws_dynamodb_tables](#aws_dynamodb_tables)


## aws_dynamodb_tables

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|global_secondary_indexes|jsonb|added|
|id|text|removed|
|inaccessible_encryption_date_time|timestamp without time zone|removed|
|kms_master_key_arn|text|removed|
|local_secondary_indexes|jsonb|added|
|name|text|removed|
|provisioned_throughput|jsonb|added|
|provisioned_throughput_last_decrease_date_time|timestamp without time zone|removed|
|provisioned_throughput_last_increase_date_time|timestamp without time zone|removed|
|provisioned_throughput_number_of_decreases_today|bigint|removed|
|provisioned_throughput_read_capacity_units|bigint|removed|
|provisioned_throughput_write_capacity_units|bigint|removed|
|replicas|jsonb|added|
|size_bytes|bigint|removed|
|sse_description|jsonb|added|
|sse_status|text|removed|
|sse_type|text|removed|
|status|text|removed|
|table_class|text|removed|
|table_class_last_update|timestamp without time zone|removed|
|table_class_summary|jsonb|added|
|table_id|text|added|
|table_name|text|added|
|table_size_bytes|bigint|added|
|table_status|text|added|

## aws_ec2_byoip_cidrs

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|

## aws_ec2_customer_gateways

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|customer_gateway_id|text|added|
|id|text|removed|

## aws_ec2_ebs_snapshots

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|arn|text|added|
|attribute|text|added|
|create_volume_permissions|jsonb|removed|
|restore_expiry_time|timestamp without time zone|added|
|storage_tier|text|added|
|volume_size|bigint|updated|Type changed from integer to bigint

## aws_ec2_ebs_volume_attachments
Moved to JSON column on [aws_ec2_ebs_volumes](#aws_ec2_ebs_volumes)


## aws_ec2_ebs_volumes

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|attachments|jsonb|added|
|id|text|removed|
|iops|bigint|updated|Type changed from integer to bigint
|size|bigint|updated|Type changed from integer to bigint
|throughput|bigint|updated|Type changed from integer to bigint
|volume_id|text|added|

## aws_ec2_egress_only_internet_gateways

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|egress_only_internet_gateway_id|text|added|
|id|text|removed|

## aws_ec2_eips

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|carrier_ip|text|updated|Type changed from inet to text
|customer_owned_ip|text|updated|Type changed from inet to text
|private_ip_address|text|updated|Type changed from inet to text
|public_ip|text|updated|Type changed from inet to text

## aws_ec2_flow_logs

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|destination_options|jsonb|added|
|id|text|removed|
|max_aggregation_interval|bigint|updated|Type changed from integer to bigint

## aws_ec2_host_available_instance_capacity
Moved to JSON column on [aws_ec2_hosts](#aws_ec2_hosts)


## aws_ec2_host_instances
Moved to JSON column on [aws_ec2_hosts](#aws_ec2_hosts)


## aws_ec2_hosts

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|available_capacity|jsonb|added|
|available_vcpus|integer|removed|
|cores|integer|removed|
|host_id|text|added|
|host_properties|jsonb|added|
|host_reservation_id|text|added|
|id|text|removed|
|instance_family|text|removed|
|instance_type|text|removed|
|instances|jsonb|added|
|outpost_arn|text|added|
|reservation_id|text|removed|
|sockets|integer|removed|
|total_vcpus|integer|removed|

## aws_ec2_image_block_device_mappings
Moved to JSON column on [aws_ec2_images](#aws_ec2_images)


## aws_ec2_images

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|block_device_mappings|jsonb|added|
|boot_mode|text|added|
|creation_date|text|updated|Type changed from timestamp without time zone to text
|deprecation_time|text|updated|Type changed from timestamp without time zone to text
|id|text|removed|
|image_id|text|added|
|last_launched_time|timestamp without time zone|removed|
|state_reason|jsonb|added|
|state_reason_code|text|removed|
|state_reason_message|text|removed|
|tpm_support|text|added|

## aws_ec2_instance_block_device_mappings
Moved to JSON column on [aws_ec2_instances](#aws_ec2_instances)


## aws_ec2_instance_elastic_gpu_associations
Moved to JSON column on [aws_ec2_instances](#aws_ec2_instances)


## aws_ec2_instance_elastic_inference_accelerator_associations
Moved to JSON column on [aws_ec2_instances](#aws_ec2_instances)


## aws_ec2_instance_network_interface_groups
Moved to JSON column on [aws_ec2_instances](#aws_ec2_instances)


## aws_ec2_instance_network_interface_ipv6_addresses
Moved to JSON column on [aws_ec2_instances](#aws_ec2_instances)


## aws_ec2_instance_network_interface_private_ip_addresses
Moved to JSON column on [aws_ec2_instances](#aws_ec2_instances)


## aws_ec2_instance_network_interfaces
Moved to JSON column on [aws_ec2_instances](#aws_ec2_instances)


## aws_ec2_instance_product_codes
Moved to JSON column on [aws_ec2_instances](#aws_ec2_instances)


## aws_ec2_instance_security_groups
Moved to JSON column on [aws_ec2_instances](#aws_ec2_instances)


## aws_ec2_instance_status_events
Moved to JSON column on [aws_ec2_instances](#aws_ec2_instances)


## aws_ec2_instance_statuses

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|details|jsonb|removed|
|events|jsonb|added|
|instance_state|jsonb|added|
|instance_state_code|integer|removed|
|instance_state_name|text|removed|
|instance_status|jsonb|added|
|status|text|removed|
|system_status|jsonb|updated|Type changed from text to jsonb
|system_status_details|jsonb|removed|

## aws_ec2_instance_type_fpga_info_fpgas
Moved to JSON column on [aws_ec2_instances](#aws_ec2_instances)


## aws_ec2_instance_type_gpu_info_gpus
Moved to JSON column on [aws_ec2_instances](#aws_ec2_instances)


## aws_ec2_instance_type_inference_accelerator_info_accelerators
Moved to JSON column on [aws_ec2_instances](#aws_ec2_instances)


## aws_ec2_instance_type_instance_storage_info_disks
Moved to JSON column on [aws_ec2_instances](#aws_ec2_instances)


## aws_ec2_instance_type_network_info_network_cards
Moved to JSON column on [aws_ec2_instances](#aws_ec2_instances)


## aws_ec2_instance_types

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|arn|text|added|
|ebs_info|jsonb|added|
|ebs_info_ebs_optimized_info_baseline_bandwidth_in_mbps|bigint|removed|
|ebs_info_ebs_optimized_info_baseline_iops|bigint|removed|
|ebs_info_ebs_optimized_info_baseline_throughput_in_mb_ps|float|removed|
|ebs_info_ebs_optimized_info_maximum_bandwidth_in_mbps|bigint|removed|
|ebs_info_ebs_optimized_info_maximum_iops|bigint|removed|
|ebs_info_ebs_optimized_info_maximum_throughput_in_mb_ps|float|removed|
|ebs_info_ebs_optimized_support|text|removed|
|ebs_info_encryption_support|text|removed|
|ebs_info_nvme_support|text|removed|
|fpga_info|jsonb|added|
|fpga_info_total_fpga_memory_in_mi_b|bigint|removed|
|gpu_info|jsonb|added|
|gpu_info_total_gpu_memory_in_mi_b|bigint|removed|
|inference_accelerator_info|jsonb|added|
|instance_storage_info|jsonb|added|
|instance_storage_info_encryption_support|text|removed|
|instance_storage_info_nvme_support|text|removed|
|instance_storage_info_total_size_in_gb|bigint|removed|
|memory_info|jsonb|added|
|memory_info_size_in_mi_b|bigint|removed|
|network_info|jsonb|added|
|network_info_default_network_card_index|bigint|removed|
|network_info_efa_info_maximum_efa_interfaces|bigint|removed|
|network_info_efa_supported|boolean|removed|
|network_info_ena_support|text|removed|
|network_info_encryption_in_transit_supported|boolean|removed|
|network_info_ipv4_addresses_per_interface|bigint|removed|
|network_info_ipv6_addresses_per_interface|bigint|removed|
|network_info_ipv6_supported|boolean|removed|
|network_info_maximum_network_cards|bigint|removed|
|network_info_maximum_network_interfaces|bigint|removed|
|network_info_network_performance|text|removed|
|placement_group_info|jsonb|added|
|placement_group_info_supported_strategies|text[]|removed|
|processor_info|jsonb|added|
|processor_info_supported_architectures|text[]|removed|
|processor_info_sustained_clock_speed_in_ghz|float|removed|
|v_cpu_info|jsonb|added|
|v_cpu_info_default_cores|bigint|removed|
|v_cpu_info_default_threads_per_core|bigint|removed|
|v_cpu_info_default_v_cpus|bigint|removed|
|v_cpu_info_valid_cores|integer[]|removed|
|v_cpu_info_valid_threads_per_core|integer[]|removed|

## aws_ec2_instances

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|ami_launch_index|bigint|updated|Type changed from integer to bigint
|block_device_mappings|jsonb|added|
|cap_reservation_preference|text|removed|
|cap_reservation_target_capacity_reservation_id|text|removed|
|cap_reservation_target_capacity_reservation_rg_arn|text|removed|
|capacity_reservation_specification|jsonb|added|
|cpu_options|jsonb|added|
|cpu_options_core_count|integer|removed|
|cpu_options_threads_per_core|integer|removed|
|elastic_gpu_associations|jsonb|added|
|elastic_inference_accelerator_associations|jsonb|added|
|enclave_options|jsonb|added|
|enclave_options_enabled|boolean|removed|
|hibernation_options|jsonb|added|
|hibernation_options_configured|boolean|removed|
|iam_instance_profile|jsonb|added|
|iam_instance_profile_arn|text|removed|
|iam_instance_profile_id|text|removed|
|id|text|removed|
|instance_id|text|added|
|ipv6_address|text|added|
|licenses|jsonb|updated|Type changed from text[] to jsonb
|maintenance_options|jsonb|added|
|metadata_options|jsonb|added|
|metadata_options_http_endpoint|text|removed|
|metadata_options_http_protocol_ipv6|text|removed|
|metadata_options_http_put_response_hop_limit|integer|removed|
|metadata_options_http_tokens|text|removed|
|metadata_options_state|text|removed|
|monitoring|jsonb|added|
|monitoring_state|text|removed|
|network_interfaces|jsonb|added|
|placement|jsonb|added|
|placement_affinity|text|removed|
|placement_availability_zone|text|removed|
|placement_group_name|text|removed|
|placement_host_id|text|removed|
|placement_host_resource_group_arn|text|removed|
|placement_partition_number|integer|removed|
|placement_spread_domain|text|removed|
|placement_tenancy|text|removed|
|platform_details|text|added|
|private_dns_name_options|jsonb|added|
|product_codes|jsonb|added|
|security_groups|jsonb|added|
|state|jsonb|added|
|state_code|integer|removed|
|state_name|text|removed|
|state_reason|jsonb|added|
|state_reason_code|text|removed|
|state_reason_message|text|removed|
|state_transition_reason_time|timestamp without time zone|removed|
|tpm_support|text|added|
|usage_operation|text|added|
|usage_operation_update_time|timestamp without time zone|added|

## aws_ec2_internet_gateway_attachments
Moved to JSON column on [aws_ec2_internet_gateways](#aws_ec2_internet_gateways)


## aws_ec2_internet_gateways

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|attachments|jsonb|added|
|id|text|removed|
|internet_gateway_id|text|added|

## aws_ec2_key_pairs

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|public_key|text|added|

## aws_ec2_nat_gateway_addresses
Moved to JSON column on [aws_ec2_nat_gateways](#aws_ec2_nat_gateways)


## aws_ec2_nat_gateways

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|connectivity_type|text|added|
|id|text|removed|
|nat_gateway_addresses|jsonb|added|
|nat_gateway_id|text|added|
|provisioned_bandwidth|jsonb|added|
|provisioned_bandwidth_provision_time|timestamp without time zone|removed|
|provisioned_bandwidth_provisioned|text|removed|
|provisioned_bandwidth_request_time|timestamp without time zone|removed|
|provisioned_bandwidth_requested|text|removed|
|provisioned_bandwidth_status|text|removed|

## aws_ec2_network_acl_associations
Moved to JSON column on [aws_ec2_network_acls](#aws_ec2_network_acls)


## aws_ec2_network_acl_entries
Moved to JSON column on [aws_ec2_network_acls](#aws_ec2_network_acls)


## aws_ec2_network_acls

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|associations|jsonb|added|
|entries|jsonb|added|
|id|text|removed|
|network_acl_id|text|added|

## aws_ec2_network_interface_private_ip_addresses
Moved to JSON column on [aws_ec2_network_interfaces](#aws_ec2_network_interfaces)


## aws_ec2_network_interfaces

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|association|jsonb|added|
|association_allocation_id|text|removed|
|association_carrier_ip|text|removed|
|association_customer_owned_ip|text|removed|
|association_id|text|removed|
|association_ip_owner_id|text|removed|
|association_public_dns_name|text|removed|
|association_public_ip|text|removed|
|attachment|jsonb|added|
|attachment_attach_time|timestamp without time zone|removed|
|attachment_delete_on_termination|boolean|removed|
|attachment_device_index|integer|removed|
|attachment_id|text|removed|
|attachment_instance_id|text|removed|
|attachment_instance_owner_id|text|removed|
|attachment_network_card_index|integer|removed|
|attachment_status|text|removed|
|id|text|removed|
|ipv4_prefixes|jsonb|updated|Type changed from text[] to jsonb
|ipv6_addresses|jsonb|updated|Type changed from text[] to jsonb
|ipv6_prefixes|jsonb|updated|Type changed from text[] to jsonb
|network_interface_id|text|added|
|private_ip_addresses|jsonb|added|
|tag_set|jsonb|added|
|tags|jsonb|removed|

## aws_ec2_regional_config

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|

## aws_ec2_route_table_associations
Moved to JSON column on [aws_ec2_route_tables](#aws_ec2_route_tables)


## aws_ec2_route_table_propagating_vgws
Moved to JSON column on [aws_ec2_route_tables](#aws_ec2_route_tables)


## aws_ec2_route_table_routes
Moved to JSON column on [aws_ec2_route_tables](#aws_ec2_route_tables)


## aws_ec2_route_tables

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|associations|jsonb|added|
|id|text|removed|
|propagating_vgws|jsonb|added|
|route_table_id|text|added|
|routes|jsonb|added|

## aws_ec2_security_group_ip_permission_ip_ranges
Moved to JSON column on [aws_ec2_security_groups](#aws_ec2_security_groups)


## aws_ec2_security_group_ip_permission_prefix_list_ids
Moved to JSON column on [aws_ec2_security_groups](#aws_ec2_security_groups)


## aws_ec2_security_group_ip_permission_user_id_group_pairs
Moved to JSON column on [aws_ec2_security_groups](#aws_ec2_security_groups)


## aws_ec2_security_group_ip_permissions
Moved to JSON column on [aws_ec2_security_groups](#aws_ec2_security_groups)


## aws_ec2_security_groups

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|group_id|text|added|
|id|text|removed|
|ip_permissions|jsonb|added|
|ip_permissions_egress|jsonb|added|

## aws_ec2_subnet_ipv6_cidr_block_association_sets
Moved to JSON column on [aws_ec2_subnets](#aws_ec2_subnets)


## aws_ec2_subnets

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|available_ip_address_count|bigint|updated|Type changed from integer to bigint
|enable_dns64|boolean|added|
|enable_lni_at_device_index|bigint|added|
|id|text|removed|
|ipv6_cidr_block_association_set|jsonb|added|
|ipv6_native|boolean|added|
|private_dns_name_options_on_launch|jsonb|added|
|subnet_arn|text|added|
|subnet_id|text|added|

## aws_ec2_transit_gateway_attachments

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|account_id|text|added|
|association|jsonb|added|
|association_route_table_id|text|removed|
|association_state|text|removed|
|region|text|added|
|transit_gateway_arn|text|added|
|transit_gateway_attachment_id|text|added|
|transit_gateway_cq_id|uuid|removed|
|transit_gateway_id|text|added|

## aws_ec2_transit_gateway_multicast_domains

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|account_id|text|added|
|auto_accept_shared_associations|text|removed|
|igmpv2_support|text|removed|
|options|jsonb|added|
|region|text|added|
|static_sources_support|text|removed|
|transit_gateway_arn|text|added|
|transit_gateway_cq_id|uuid|removed|
|transit_gateway_id|text|added|

## aws_ec2_transit_gateway_peering_attachments

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|accepter_owner_id|text|removed|
|accepter_region|text|removed|
|accepter_tgw_info|jsonb|added|
|accepter_transit_gateway_id|text|removed|
|account_id|text|added|
|region|text|added|
|requester_owner_id|text|removed|
|requester_region|text|removed|
|requester_tgw_info|jsonb|added|
|requester_transit_gateway_id|text|removed|
|status|jsonb|added|
|status_code|text|removed|
|status_message|text|removed|
|transit_gateway_arn|text|added|
|transit_gateway_cq_id|uuid|removed|

## aws_ec2_transit_gateway_route_tables

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|account_id|text|added|
|region|text|added|
|transit_gateway_arn|text|added|
|transit_gateway_cq_id|uuid|removed|
|transit_gateway_id|text|added|

## aws_ec2_transit_gateway_vpc_attachments

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|account_id|text|added|
|appliance_mode_support|text|removed|
|dns_support|text|removed|
|ipv6_support|text|removed|
|options|jsonb|added|
|region|text|added|
|subnet_ids|text[]|added|
|transit_gateway_arn|text|added|
|transit_gateway_cq_id|uuid|removed|
|transit_gateway_id|text|added|

## aws_ec2_transit_gateways

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|amazon_side_asn|bigint|removed|
|association_default_route_table_id|text|removed|
|auto_accept_shared_attachments|text|removed|
|default_route_table_association|text|removed|
|default_route_table_propagation|text|removed|
|dns_support|text|removed|
|multicast_support|text|removed|
|options|jsonb|added|
|propagation_default_route_table_id|text|removed|
|region|text|removed|
|transit_gateway_cidr_blocks|text[]|removed|
|vpn_ecmp_support|text|removed|

## aws_ec2_vpc_attachment
Moved to JSON column on [aws_ec2_vpcs](#aws_ec2_vpcs)


## aws_ec2_vpc_cidr_block_association_sets
Moved to JSON column on [aws_ec2_vpcs](#aws_ec2_vpcs)


## aws_ec2_vpc_endpoint_dns_entries
Moved to JSON column on [aws_ec2_vpcs](#aws_ec2_vpcs)


## aws_ec2_vpc_endpoint_groups
Moved to JSON column on [aws_ec2_vpcs](#aws_ec2_vpcs)


## aws_ec2_vpc_endpoint_service_configurations

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|private_dns_name_configuration|jsonb|added|
|private_dns_name_configuration_name|text|removed|
|private_dns_name_configuration_state|text|removed|
|private_dns_name_configuration_type|text|removed|
|private_dns_name_configuration_value|text|removed|
|service_type|jsonb|updated|Type changed from text[] to jsonb
|supported_ip_address_types|text[]|added|

## aws_ec2_vpc_endpoint_services

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|id|text|removed|
|private_dns_names|jsonb|updated|Type changed from text[] to jsonb
|service_id|text|added|
|service_type|jsonb|updated|Type changed from text[] to jsonb
|supported_ip_address_types|text[]|added|

## aws_ec2_vpc_endpoints

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|dns_entries|jsonb|added|
|dns_options|jsonb|added|
|groups|jsonb|added|
|id|text|removed|
|ip_address_type|text|added|
|last_error|jsonb|added|
|last_error_code|text|removed|
|last_error_message|text|removed|
|vpc_endpoint_id|text|added|

## aws_ec2_vpc_ipv6_cidr_block_association_sets
Moved to JSON column on [aws_ec2_vpcs](#aws_ec2_vpcs)


## aws_ec2_vpc_peering_connections

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|accepter_allow_dns_resolution_from_remote_vpc|boolean|removed|
|accepter_allow_egress_local_classic_link_to_remote_vpc|boolean|removed|
|accepter_allow_egress_local_vpc_to_remote_classic_link|boolean|removed|
|accepter_cidr_block|text|removed|
|accepter_cidr_block_set|text[]|removed|
|accepter_ipv6_cidr_block_set|text[]|removed|
|accepter_owner_id|text|removed|
|accepter_vpc_id|text|removed|
|accepter_vpc_info|jsonb|added|
|accepter_vpc_region|text|removed|
|id|text|removed|
|requester_allow_dns_resolution_from_remote_vpc|boolean|removed|
|requester_allow_egress_local_classic_link_to_remote_vpc|boolean|removed|
|requester_allow_egress_local_vpc_to_remote_classic_link|boolean|removed|
|requester_cidr_block|text|removed|
|requester_cidr_block_set|text[]|removed|
|requester_ipv6_cidr_block_set|text[]|removed|
|requester_owner_id|text|removed|
|requester_vpc_id|text|removed|
|requester_vpc_info|jsonb|added|
|requester_vpc_region|text|removed|
|status|jsonb|added|
|status_code|text|removed|
|status_message|text|removed|
|vpc_peering_connection_id|text|added|

## aws_ec2_vpcs

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|cidr_block_association_set|jsonb|added|
|id|text|removed|
|ipv6_cidr_block_association_set|jsonb|added|
|vpc_id|text|added|

## aws_ec2_vpn_gateways

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|id|text|removed|
|vpc_attachments|jsonb|added|
|vpn_gateway_id|text|added|

## aws_ecr_repositories

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|encryption_configuration|jsonb|added|
|encryption_configuration_encryption_type|text|removed|
|encryption_configuration_kms_key|text|removed|
|image_scanning_configuration|jsonb|added|
|image_scanning_configuration_scan_on_push|boolean|removed|
|name|text|removed|
|repository_name|text|added|
|repository_uri|text|added|
|tags|text|updated|Type changed from jsonb to text
|uri|text|removed|

## aws_ecr_repository_images

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|arn|text|added|
|image_scan_findings_summary|jsonb|added|
|image_scan_findings_summary_finding_severity_counts|jsonb|removed|
|image_scan_findings_summary_image_scan_completed_at|timestamp without time zone|removed|
|image_scan_findings_summary_vulnerability_source_updated_at|timestamp without time zone|removed|
|image_scan_status|jsonb|updated|Type changed from text to jsonb
|image_scan_status_description|text|removed|
|repository_cq_id|uuid|removed|

## aws_ecs_cluster_attachments
Moved to JSON column on [aws_ecs_clusters](#aws_ecs_clusters)


## aws_ecs_cluster_container_instance_attachments
Moved to JSON column on [aws_ecs_clusters](#aws_ecs_clusters)


## aws_ecs_cluster_container_instance_attributes
Moved to JSON column on [aws_ecs_clusters](#aws_ecs_clusters)


## aws_ecs_cluster_container_instance_health_status_details
Moved to JSON column on [aws_ecs_clusters](#aws_ecs_clusters)


## aws_ecs_cluster_container_instance_registered_resources
Moved to JSON column on [aws_ecs_clusters](#aws_ecs_clusters)


## aws_ecs_cluster_container_instance_remaining_resources
Moved to JSON column on [aws_ecs_clusters](#aws_ecs_clusters)


## aws_ecs_cluster_container_instances

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|account_id|text|added|
|attachments|jsonb|added|
|attributes|jsonb|added|
|cluster_arn|text|added|
|cluster_cq_id|uuid|removed|
|health_status|jsonb|added|
|health_status_overall_status|text|removed|
|pending_tasks_count|bigint|updated|Type changed from integer to bigint
|region|text|added|
|registered_resources|jsonb|added|
|remaining_resources|jsonb|added|
|running_tasks_count|bigint|updated|Type changed from integer to bigint
|version_info|jsonb|added|
|version_info_agent_hash|text|removed|
|version_info_agent_version|text|removed|
|version_info_docker_version|text|removed|

## aws_ecs_cluster_service_deployments
Moved to JSON column on [aws_ecs_clusters](#aws_ecs_clusters)


## aws_ecs_cluster_service_events
Moved to JSON column on [aws_ecs_clusters](#aws_ecs_clusters)


## aws_ecs_cluster_service_load_balancers
Moved to JSON column on [aws_ecs_clusters](#aws_ecs_clusters)


## aws_ecs_cluster_service_service_registries
Moved to JSON column on [aws_ecs_clusters](#aws_ecs_clusters)


## aws_ecs_cluster_service_task_set_load_balancers
Moved to JSON column on [aws_ecs_clusters](#aws_ecs_clusters)


## aws_ecs_cluster_service_task_set_service_registries
Moved to JSON column on [aws_ecs_clusters](#aws_ecs_clusters)


## aws_ecs_cluster_service_task_sets
Moved to JSON column on [aws_ecs_clusters](#aws_ecs_clusters)


## aws_ecs_cluster_services

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|account_id|text|added|
|cluster_cq_id|uuid|removed|
|deployment_configuration|jsonb|added|
|deployment_configuration_deployment_circuit_breaker_enable|boolean|removed|
|deployment_configuration_deployment_circuit_breaker_rollback|boolean|removed|
|deployment_configuration_maximum_percent|integer|removed|
|deployment_configuration_minimum_healthy_percent|integer|removed|
|deployment_controller|jsonb|added|
|deployment_controller_type|text|removed|
|deployments|jsonb|added|
|desired_count|bigint|updated|Type changed from integer to bigint
|events|jsonb|added|
|health_check_grace_period_seconds|bigint|updated|Type changed from integer to bigint
|load_balancers|jsonb|added|
|name|text|removed|
|network_configuration|jsonb|added|
|network_configuration_awsvpc_configuration_assign_public_ip|text|removed|
|network_configuration_awsvpc_configuration_security_groups|text[]|removed|
|network_configuration_awsvpc_configuration_subnets|text[]|removed|
|pending_count|bigint|updated|Type changed from integer to bigint
|region|text|added|
|running_count|bigint|updated|Type changed from integer to bigint
|service_name|text|added|
|service_registries|jsonb|added|
|task_sets|jsonb|added|

## aws_ecs_cluster_task_attachments
Moved to JSON column on [aws_ecs_clusters](#aws_ecs_clusters)


## aws_ecs_cluster_task_containers
Moved to JSON column on [aws_ecs_clusters](#aws_ecs_clusters)


## aws_ecs_cluster_tasks

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|account_id|text|added|
|attachments|jsonb|added|
|cluster_cq_id|uuid|removed|
|containers|jsonb|added|
|ephemeral_storage|jsonb|added|
|ephemeral_storage_size_in_gib|integer|removed|
|region|text|added|

## aws_ecs_clusters

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|active_services_count|bigint|updated|Type changed from integer to bigint
|attachments|jsonb|added|
|cluster_name|text|added|
|configuration|jsonb|added|
|execute_config_kms_key_id|text|removed|
|execute_config_log_cloud_watch_log_group_name|text|removed|
|execute_config_log_s3_bucket_name|text|removed|
|execute_config_log_s3_encryption_enabled|boolean|removed|
|execute_config_log_s3_key_prefix|text|removed|
|execute_config_logging|text|removed|
|execute_config_logs_cloud_watch_encryption_enabled|boolean|removed|
|name|text|removed|
|pending_tasks_count|bigint|updated|Type changed from integer to bigint
|registered_container_instances_count|bigint|updated|Type changed from integer to bigint
|running_tasks_count|bigint|updated|Type changed from integer to bigint

## aws_ecs_task_definition_container_definitions
Moved to JSON column on [aws_ecs_task_definitions](#aws_ecs_task_definitions)


## aws_ecs_task_definition_volumes
Moved to JSON column on [aws_ecs_task_definitions](#aws_ecs_task_definitions)


## aws_ecs_task_definitions

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|arn|text|removed|
|container_definitions|jsonb|added|
|ephemeral_storage|jsonb|added|
|ephemeral_storage_size|integer|removed|
|proxy_configuration|jsonb|added|
|proxy_configuration_container_name|text|removed|
|proxy_configuration_properties|jsonb|removed|
|proxy_configuration_type|text|removed|
|revision|bigint|updated|Type changed from integer to bigint
|runtime_platform|jsonb|added|
|runtime_platform_cpu_architecture|text|removed|
|runtime_platform_os_family|text|removed|
|task_definition_arn|text|added|
|volumes|jsonb|added|

## aws_efs_filesystems

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|file_system_id|text|added|
|id|text|removed|
|number_of_mount_targets|bigint|updated|Type changed from integer to bigint
|provisioned_throughput_in_mibps|real|updated|Type changed from float to real
|size_in_bytes|jsonb|added|
|size_in_bytes_timestamp|timestamp without time zone|removed|
|size_in_bytes_value|bigint|removed|
|size_in_bytes_value_in_ia|bigint|removed|
|size_in_bytes_value_in_standard|bigint|removed|

## aws_eks_cluster_encryption_configs
Moved to JSON column on [aws_eks_clusters](#aws_eks_clusters)


## aws_eks_cluster_loggings
Moved to JSON column on [aws_eks_clusters](#aws_eks_clusters)


## aws_eks_clusters

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|certificate_authority|jsonb|added|
|certificate_authority_data|text|removed|
|connector_config|jsonb|added|
|encryption_config|jsonb|added|
|identity|jsonb|added|
|identity_oidc_issuer|text|removed|
|kubernetes_network_config|jsonb|added|
|kubernetes_network_config_service_ipv4_cidr|text|removed|
|logging|jsonb|added|
|resources_vpc_config|jsonb|added|
|resources_vpc_config_cluster_security_group_id|text|removed|
|resources_vpc_config_endpoint_private_access|boolean|removed|
|resources_vpc_config_endpoint_public_access|boolean|removed|
|resources_vpc_config_public_access_cidrs|text[]|removed|
|resources_vpc_config_security_group_ids|text[]|removed|
|resources_vpc_config_subnet_ids|text[]|removed|
|resources_vpc_config_vpc_id|text|removed|

## aws_elasticache_cluster_cache_nodes
Moved to JSON column on [aws_elasticache_clusters](#aws_elasticache_clusters)


## aws_elasticache_cluster_cache_security_groups
Moved to JSON column on [aws_elasticache_clusters](#aws_elasticache_clusters)


## aws_elasticache_cluster_log_delivery_configurations
Moved to JSON column on [aws_elasticache_clusters](#aws_elasticache_clusters)


## aws_elasticache_cluster_security_groups
Moved to JSON column on [aws_elasticache_clusters](#aws_elasticache_clusters)


## aws_elasticache_clusters

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|cache_cluster_create_time|timestamp without time zone|added|
|cache_cluster_id|text|added|
|cache_cluster_status|text|added|
|cache_nodes|jsonb|added|
|cache_parameter_group|jsonb|added|
|cache_parameter_group_cache_node_ids_to_reboot|text[]|removed|
|cache_parameter_group_name|text|removed|
|cache_parameter_group_parameter_apply_status|text|removed|
|cache_security_groups|jsonb|added|
|configuration_endpoint|jsonb|added|
|configuration_endpoint_address|text|removed|
|configuration_endpoint_port|bigint|removed|
|create_time|timestamp without time zone|removed|
|id|text|removed|
|log_delivery_configurations|jsonb|added|
|notification_configuration|jsonb|added|
|notification_configuration_topic_arn|text|removed|
|notification_configuration_topic_status|text|removed|
|pending_auth_token_status|text|removed|
|pending_cache_node_ids_to_remove|text[]|removed|
|pending_cache_node_type|text|removed|
|pending_engine_version|text|removed|
|pending_modified_values|jsonb|added|
|pending_num_cache_nodes|bigint|removed|
|security_groups|jsonb|added|
|status|text|removed|

## aws_elasticache_engine_versions

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|

## aws_elasticache_global_replication_group_global_node_groups
Moved to JSON column on [aws_elasticache_global_replication_groups](#aws_elasticache_global_replication_groups)


## aws_elasticache_global_replication_group_members
Moved to JSON column on [aws_elasticache_global_replication_groups](#aws_elasticache_global_replication_groups)


## aws_elasticache_global_replication_groups

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|global_node_groups|jsonb|added|
|members|jsonb|added|
|region|text|added|

## aws_elasticache_parameter_group_parameters
Moved to JSON column on [aws_elasticache_parameter_groups](#aws_elasticache_parameter_groups)


## aws_elasticache_parameter_groups

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|

## aws_elasticache_replication_group_log_delivery_configurations
Moved to JSON column on [aws_elasticache_replication_groups](#aws_elasticache_replication_groups)


## aws_elasticache_replication_group_node_group_members
Moved to JSON column on [aws_elasticache_replication_groups](#aws_elasticache_replication_groups)


## aws_elasticache_replication_group_node_groups
Moved to JSON column on [aws_elasticache_replication_groups](#aws_elasticache_replication_groups)


## aws_elasticache_replication_groups

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|configuration_endpoint|jsonb|added|
|configuration_endpoint_address|text|removed|
|configuration_endpoint_port|bigint|removed|
|global_replication_group_id|text|removed|
|global_replication_group_info|jsonb|added|
|global_replication_group_member|text|removed|
|log_delivery_configurations|jsonb|added|
|node_groups|jsonb|added|
|pending_auth_token_status|text|removed|
|pending_automatic_failover_status|text|removed|
|pending_modified_values|jsonb|added|
|pending_primary_cluster_id|text|removed|
|pending_resharding_slot_migration_progress_percentage|float|removed|
|pending_user_group_ids_to_add|text[]|removed|
|pending_user_group_ids_to_remove|text[]|removed|

## aws_elasticache_reserved_cache_node_recurring_charges
Moved to JSON column on [aws_elasticache_reserved_cache_nodes](#aws_elasticache_reserved_cache_nodes)


## aws_elasticache_reserved_cache_nodes

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|arn|text|added|
|fixed_price|real|updated|Type changed from float to real
|recurring_charges|jsonb|added|
|reservation_arn|text|removed|
|usage_price|real|updated|Type changed from float to real

## aws_elasticache_reserved_cache_nodes_offering_recurring_charges
Moved to JSON column on [aws_elasticache_reserved_cache_nodes](#aws_elasticache_reserved_cache_nodes)


## aws_elasticache_reserved_cache_nodes_offerings

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|arn|text|added|
|fixed_price|real|updated|Type changed from float to real
|recurring_charges|jsonb|added|
|usage_price|real|updated|Type changed from float to real

## aws_elasticache_service_updates

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|arn|text|added|
|description|text|removed|
|end_date|timestamp without time zone|removed|
|name|text|removed|
|recommended_apply_by_date|timestamp without time zone|removed|
|release_date|timestamp without time zone|removed|
|service_update_description|text|added|
|service_update_end_date|timestamp without time zone|added|
|service_update_name|text|added|
|service_update_recommended_apply_by_date|timestamp without time zone|added|
|service_update_release_date|timestamp without time zone|added|
|service_update_severity|text|added|
|service_update_status|text|added|
|service_update_type|text|added|
|severity|text|removed|
|status|text|removed|
|type|text|removed|

## aws_elasticache_snapshot_node_snapshots
Moved to JSON column on [aws_elasticache_snapshots](#aws_elasticache_snapshots)


## aws_elasticache_snapshots

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|node_snapshots|jsonb|added|

## aws_elasticache_subnet_group_subnets
Moved to JSON column on [aws_elasticache_subnet_groups](#aws_elasticache_subnet_groups)


## aws_elasticache_subnet_groups

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|subnets|jsonb|added|

## aws_elasticache_user_groups

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|pending_changes|jsonb|added|
|pending_user_ids_to_add|text[]|removed|
|pending_user_ids_to_remove|text[]|removed|

## aws_elasticache_users

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|authentication|jsonb|added|
|authentication_password_count|bigint|removed|
|authentication_type|text|removed|

## aws_elasticbeanstalk_application_versions

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|source_build_information|jsonb|added|
|source_bundle|jsonb|added|
|source_bundle_s3_bucket|text|removed|
|source_bundle_s3_key|text|removed|
|source_location|text|removed|
|source_repository|text|removed|
|source_type|text|removed|

## aws_elasticbeanstalk_applications

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|application_name|text|added|
|max_age_rule_delete_source_from_s3|boolean|removed|
|max_age_rule_enabled|boolean|removed|
|max_age_rule_max_age_in_days|integer|removed|
|max_count_rule_delete_source_from_s3|boolean|removed|
|max_count_rule_enabled|boolean|removed|
|max_count_rule_max_count|integer|removed|
|name|text|removed|
|resource_lifecycle_config|jsonb|added|
|resource_lifecycle_config_service_role|text|removed|

## aws_elasticbeanstalk_configuration_options

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|account_id|text|added|
|environment_cq_id|uuid|removed|
|environment_id|text|added|
|max_length|bigint|updated|Type changed from integer to bigint
|max_value|bigint|updated|Type changed from integer to bigint
|min_value|bigint|updated|Type changed from integer to bigint
|regex|jsonb|added|
|regex_label|text|removed|
|regex_pattern|text|removed|
|region|text|added|

## aws_elasticbeanstalk_configuration_setting_options
Moved to JSON column on [aws_elasticbeanstalk_configuration_settings](#aws_elasticbeanstalk_configuration_settings)


## aws_elasticbeanstalk_configuration_settings

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|account_id|text|added|
|environment_cq_id|uuid|removed|
|environment_id|text|added|
|option_settings|jsonb|added|
|region|text|added|

## aws_elasticbeanstalk_environment_links
Moved to JSON column on [aws_elasticbeanstalk_environments](#aws_elasticbeanstalk_environments)


## aws_elasticbeanstalk_environments

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|arn|text|removed|
|environment_arn|text|added|
|environment_links|jsonb|added|
|environment_name|text|added|
|load_balancer_domain|text|removed|
|load_balancer_name|text|removed|
|name|text|removed|
|resources|jsonb|added|
|tier|jsonb|added|
|tier_name|text|removed|
|tier_type|text|removed|
|tier_version|text|removed|

## aws_elasticsearch_domains

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|advanced_security_enabled|boolean|removed|
|advanced_security_internal_user_database_enabled|boolean|removed|
|advanced_security_options|jsonb|added|
|advanced_security_options_saml_options_roles_key|text|removed|
|advanced_security_saml_enabled|boolean|removed|
|advanced_security_saml_idp_entity_id|text|removed|
|advanced_security_saml_roles_key|text|removed|
|advanced_security_saml_session_timeout_minutes|integer|removed|
|advanced_security_saml_subject_key|text|removed|
|auto_tune_error_message|text|removed|
|auto_tune_options|jsonb|added|
|auto_tune_options_state|text|removed|
|change_progress_details|jsonb|added|
|cluster_cold_storage_options_enabled|boolean|removed|
|cluster_dedicated_master_count|integer|removed|
|cluster_dedicated_master_enabled|boolean|removed|
|cluster_dedicated_master_type|text|removed|
|cluster_instance_count|integer|removed|
|cluster_instance_type|text|removed|
|cluster_warm_count|integer|removed|
|cluster_warm_enabled|boolean|removed|
|cluster_warm_type|text|removed|
|cluster_zone_awareness_config_availability_zone_count|integer|removed|
|cluster_zone_awareness_enabled|boolean|removed|
|cognito_enabled|boolean|removed|
|cognito_identity_pool_id|text|removed|
|cognito_options|jsonb|added|
|cognito_role_arn|text|removed|
|cognito_user_pool_id|text|removed|
|domain_endpoint_custom|text|removed|
|domain_endpoint_custom_certificate_arn|text|removed|
|domain_endpoint_custom_enabled|boolean|removed|
|domain_endpoint_enforce_https|boolean|removed|
|domain_endpoint_options|jsonb|added|
|domain_endpoint_tls_security_policy|text|removed|
|domain_name|text|added|
|ebs_enabled|boolean|removed|
|ebs_iops|integer|removed|
|ebs_options|jsonb|added|
|ebs_volume_size|integer|removed|
|ebs_volume_type|text|removed|
|elasticsearch_cluster_config|jsonb|added|
|encryption_at_rest_enabled|boolean|removed|
|encryption_at_rest_kms_key_id|text|removed|
|encryption_at_rest_options|jsonb|added|
|name|text|removed|
|node_to_node_encryption_enabled|boolean|removed|
|node_to_node_encryption_options|jsonb|added|
|service_software_automated_update_date|timestamp without time zone|removed|
|service_software_cancellable|boolean|removed|
|service_software_current_version|text|removed|
|service_software_description|text|removed|
|service_software_new_version|text|removed|
|service_software_optional_deployment|boolean|removed|
|service_software_options|jsonb|added|
|service_software_update_available|boolean|removed|
|service_software_update_status|text|removed|
|snapshot_options|jsonb|added|
|snapshot_options_automated_snapshot_start_hour|integer|removed|
|vpc_availability_zones|text[]|removed|
|vpc_options|jsonb|added|
|vpc_security_group_ids|text[]|removed|
|vpc_subnet_ids|text[]|removed|
|vpc_vpc_id|text|removed|

## aws_elbv1_load_balancer_backend_server_descriptions
Moved to JSON column on [aws_elbv1_load_balancers](#aws_elbv1_load_balancers)


## aws_elbv1_load_balancer_listeners
Moved to JSON column on [aws_elbv1_load_balancers](#aws_elbv1_load_balancers)


## aws_elbv1_load_balancer_policies

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|account_id|text|added|
|load_balance_name|text|removed|
|load_balancer_arn|text|added|
|load_balancer_cq_id|uuid|removed|
|load_balancer_name|text|added|
|region|text|added|

## aws_elbv1_load_balancer_policies_app_cookie_stickiness
Moved to JSON column on [aws_elbv1_load_balancers](#aws_elbv1_load_balancers)


## aws_elbv1_load_balancer_policies_lb_cookie_stickiness
Moved to JSON column on [aws_elbv1_load_balancers](#aws_elbv1_load_balancers)


## aws_elbv1_load_balancers

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|attributes|jsonb|added|
|attributes_access_log_emit_interval|integer|removed|
|attributes_access_log_enabled|boolean|removed|
|attributes_access_log_s3_bucket_name|text|removed|
|attributes_access_log_s3_bucket_prefix|text|removed|
|attributes_additional_attributes|jsonb|removed|
|attributes_connection_draining_enabled|boolean|removed|
|attributes_connection_draining_timeout|integer|removed|
|attributes_connection_settings_idle_timeout|integer|removed|
|attributes_cross_zone_load_balancing_enabled|boolean|removed|
|backend_server_descriptions|jsonb|added|
|health_check|jsonb|added|
|health_check_healthy_threshold|integer|removed|
|health_check_interval|integer|removed|
|health_check_target|text|removed|
|health_check_timeout|integer|removed|
|health_check_unhealthy_threshold|integer|removed|
|instances|jsonb|updated|Type changed from text[] to jsonb
|listener_descriptions|jsonb|added|
|load_balancer_name|text|added|
|name|text|removed|
|other_policies|text[]|removed|
|policies|jsonb|added|
|source_security_group|jsonb|added|
|source_security_group_name|text|removed|
|source_security_group_owner_alias|text|removed|

## aws_elbv2_listener_certificates

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|account_id|text|added|
|listener_arn|text|added|
|listener_cq_id|uuid|removed|
|region|text|added|

## aws_elbv2_listener_default_action_forward_config_target_groups
Moved to JSON column on [aws_elbv2_listeners](#aws_elbv2_listeners)


## aws_elbv2_listener_default_actions
Moved to JSON column on [aws_elbv2_listeners](#aws_elbv2_listeners)


## aws_elbv2_listeners

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|certificates|jsonb|added|
|default_actions|jsonb|added|
|load_balancer_cq_id|uuid|removed|
|port|bigint|updated|Type changed from integer to bigint

## aws_elbv2_load_balancer_attributes

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|access_logs_s3_bucket|text|removed|
|access_logs_s3_enabled|boolean|removed|
|access_logs_s3_prefix|text|removed|
|account_id|text|added|
|deletion_protection|boolean|removed|
|idle_timeout|integer|removed|
|key|text|added|
|load_balancer_arn|text|added|
|load_balancer_cq_id|uuid|removed|
|load_balancing_cross_zone|boolean|removed|
|region|text|added|
|routing_http2|boolean|removed|
|routing_http_desync_mitigation_mode|text|removed|
|routing_http_drop_invalid_header_fields|boolean|removed|
|routing_http_xamzntls_enabled|boolean|removed|
|routing_http_xff_client_port|boolean|removed|
|value|text|added|
|waf_fail_open|boolean|removed|

## aws_elbv2_load_balancer_availability_zone_addresses
Moved to JSON column on [aws_elbv2_load_balancers](#aws_elbv2_load_balancers)


## aws_elbv2_load_balancer_availability_zones
Moved to JSON column on [aws_elbv2_load_balancers](#aws_elbv2_load_balancers)


## aws_elbv2_load_balancers

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|availability_zones|jsonb|added|
|load_balancer_name|text|added|
|name|text|removed|
|state|jsonb|added|
|state_code|text|removed|
|state_reason|text|removed|

## aws_elbv2_target_group_target_health_descriptions

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|account_id|text|added|
|region|text|added|
|target|jsonb|added|
|target_availability_zone|text|removed|
|target_group_arn|text|added|
|target_group_cq_id|uuid|removed|
|target_health|jsonb|added|
|target_health_description|text|removed|
|target_health_reason|text|removed|
|target_health_state|text|removed|
|target_id|text|removed|
|target_port|integer|removed|

## aws_elbv2_target_groups

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|health_check_interval_seconds|bigint|updated|Type changed from integer to bigint
|health_check_timeout_seconds|bigint|updated|Type changed from integer to bigint
|healthy_threshold_count|bigint|updated|Type changed from integer to bigint
|ip_address_type|text|added|
|matcher|jsonb|added|
|matcher_grpc_code|text|removed|
|matcher_http_code|text|removed|
|name|text|removed|
|port|bigint|updated|Type changed from integer to bigint
|target_group_name|text|added|
|unhealthy_threshold_count|bigint|updated|Type changed from integer to bigint

## aws_emr_block_public_access_config_port_ranges
Moved to JSON column on [aws_emr_block_public_access_configs](#aws_emr_block_public_access_configs)


## aws_emr_block_public_access_configs

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|block_public_access_configuration|jsonb|added|
|block_public_access_configuration_metadata|jsonb|added|
|block_public_security_group_rules|boolean|removed|
|classification|text|removed|
|configurations|jsonb|removed|
|created_by_arn|text|removed|
|creation_date_time|timestamp without time zone|removed|
|properties|jsonb|removed|

## aws_emr_clusters

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|creation_date_time|timestamp without time zone|removed|
|ebs_root_volume_size|bigint|updated|Type changed from integer to bigint
|ec2_instance_attribute_additional_master_security_groups|text[]|removed|
|ec2_instance_attribute_additional_slave_security_groups|text[]|removed|
|ec2_instance_attribute_availability_zone|text|removed|
|ec2_instance_attribute_emr_managed_master_security_group|text|removed|
|ec2_instance_attribute_emr_managed_slave_security_group|text|removed|
|ec2_instance_attribute_iam_instance_profile|text|removed|
|ec2_instance_attribute_key_name|text|removed|
|ec2_instance_attribute_requested_availability_zones|text[]|removed|
|ec2_instance_attribute_requested_subnet_ids|text[]|removed|
|ec2_instance_attribute_service_access_security_group|text|removed|
|ec2_instance_attribute_subnet_id|text|removed|
|ec2_instance_attributes|jsonb|added|
|end_date_time|timestamp without time zone|removed|
|kerberos_ad_domain_join_password|text|removed|
|kerberos_ad_domain_join_user|text|removed|
|kerberos_attributes|jsonb|added|
|kerberos_cross_realm_trust_principal_password|text|removed|
|kerberos_kdc_admin_password|text|removed|
|kerberos_realm|text|removed|
|normalized_instance_hours|bigint|updated|Type changed from integer to bigint
|os_release_label|text|added|
|ready_date_time|timestamp without time zone|removed|
|state|text|removed|
|state_change_reason_code|text|removed|
|state_change_reason_message|text|removed|
|status|jsonb|added|
|step_concurrency_level|bigint|updated|Type changed from integer to bigint

## aws_eventbridge_event_bus_rules

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|account_id|text|added|
|event_bus_arn|text|added|
|event_bus_cq_id|uuid|removed|
|region|text|added|

## aws_eventbridge_event_buses

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|

## aws_firehose_delivery_stream_elasticsearch_destination
Moved to JSON column on [aws_firehose_delivery_streams](#aws_firehose_delivery_streams)


## aws_firehose_delivery_stream_extended_s3_destination
Moved to JSON column on [aws_firehose_delivery_streams](#aws_firehose_delivery_streams)


## aws_firehose_delivery_stream_http_destination
Moved to JSON column on [aws_firehose_delivery_streams](#aws_firehose_delivery_streams)


## aws_firehose_delivery_stream_open_search_destination
Moved to JSON column on [aws_firehose_delivery_streams](#aws_firehose_delivery_streams)


## aws_firehose_delivery_stream_redshift_destination
Moved to JSON column on [aws_firehose_delivery_streams](#aws_firehose_delivery_streams)


## aws_firehose_delivery_stream_splunk_destination
Moved to JSON column on [aws_firehose_delivery_streams](#aws_firehose_delivery_streams)


## aws_firehose_delivery_streams

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|delivery_stream_arn|text|removed|
|delivery_stream_encryption_configuration|jsonb|added|
|destinations|jsonb|added|
|encryption_config_failure_description_details|text|removed|
|encryption_config_failure_description_type|text|removed|
|encryption_config_key_arn|text|removed|
|encryption_config_key_type|text|removed|
|encryption_config_status|text|removed|
|failure_description|jsonb|added|
|failure_description_details|text|removed|
|failure_description_type|text|removed|
|has_more_destinations|boolean|added|
|source|jsonb|added|
|source_kinesis_stream_delivery_start_timestamp|timestamp without time zone|removed|
|source_kinesis_stream_kinesis_stream_arn|text|removed|
|source_kinesis_stream_role_arn|text|removed|

## aws_fsx_backups

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|arn|text|removed|
|directory_information|jsonb|added|
|directory_information_active_directory_id|text|removed|
|directory_information_domain_name|text|removed|
|failure_details|jsonb|added|
|failure_details_message|text|removed|
|file_system|jsonb|added|
|owner_id|text|added|
|progress_percent|bigint|updated|Type changed from integer to bigint
|resource_arn|text|added|
|resource_type|text|added|
|source_backup_id|text|added|
|source_backup_region|text|added|
|volume|jsonb|added|

## aws_fsx_data_repo_associations
This table was removed.


## aws_fsx_data_repo_tasks
This table was removed.


## aws_fsx_data_repository_associations
This table was newly added.

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|account_id|text|added|
|region|text|added|
|arn|text|added|
|tags|jsonb|added|
|association_id|text|added|
|batch_import_meta_data_on_create|boolean|added|
|creation_time|timestamp without time zone|added|
|data_repository_path|text|added|
|failure_details|jsonb|added|
|file_system_id|text|added|
|file_system_path|text|added|
|imported_file_chunk_size|bigint|added|
|lifecycle|text|added|
|s3|jsonb|added|
|_cq_id|uuid|added|
|_cq_fetch_time|timestamp without time zone|added|

## aws_fsx_data_repository_tasks
This table was newly added.

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|account_id|text|added|
|region|text|added|
|arn|text|added|
|tags|jsonb|added|
|creation_time|timestamp without time zone|added|
|file_system_id|text|added|
|lifecycle|text|added|
|task_id|text|added|
|type|text|added|
|end_time|timestamp without time zone|added|
|failure_details|jsonb|added|
|paths|text[]|added|
|report|jsonb|added|
|start_time|timestamp without time zone|added|
|status|jsonb|added|
|_cq_id|uuid|added|
|_cq_fetch_time|timestamp without time zone|added|

## aws_fsx_file_systems
Renamed from [aws_fsx_filesystems](aws_fsx_filesystems)


## aws_fsx_filesystem_lustre_configuration
Moved to JSON column on [aws_fsx_file_systems](#aws_fsx_file_systems)


## aws_fsx_filesystem_ontap_configuration
Moved to JSON column on [aws_fsx_file_systems](#aws_fsx_file_systems)


## aws_fsx_filesystem_open_zfs_configuration
Moved to JSON column on [aws_fsx_file_systems](#aws_fsx_file_systems)


## aws_fsx_filesystem_windows_configuration
Moved to JSON column on [aws_fsx_file_systems](#aws_fsx_file_systems)


## aws_fsx_filesystems
Renamed to [aws_fsx_file_systems](#aws_fsx_file_systems)


## aws_fsx_snapshots

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|administrative_actions|jsonb|added|
|lifecycle_transition_reason|jsonb|added|
|lifecycle_transition_reason_message|text|removed|

## aws_fsx_storage_virtual_machines
This table was newly added.

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|account_id|text|added|
|region|text|added|
|arn|text|added|
|tags|jsonb|added|
|active_directory_configuration|jsonb|added|
|creation_time|timestamp without time zone|added|
|endpoints|jsonb|added|
|file_system_id|text|added|
|lifecycle|text|added|
|lifecycle_transition_reason|jsonb|added|
|name|text|added|
|root_volume_security_style|text|added|
|storage_virtual_machine_id|text|added|
|subtype|text|added|
|uuid|text|added|
|_cq_id|uuid|added|
|_cq_fetch_time|timestamp without time zone|added|

## aws_fsx_storage_vms
This table was removed.


## aws_fsx_volume_ontap_configuration
Moved to JSON column on [aws_fsx_volumes](#aws_fsx_volumes)


## aws_fsx_volume_open_zfs_configuration
Moved to JSON column on [aws_fsx_volumes](#aws_fsx_volumes)


## aws_fsx_volumes

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|administrative_actions|jsonb|added|
|id|text|removed|
|lifecycle_transition_reason|jsonb|added|
|lifecycle_transition_reason_message|text|removed|
|ontap_configuration|jsonb|added|
|open_zfs_configuration|jsonb|added|
|volume_id|text|added|

## aws_glue_classifiers

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|csv_classifier|jsonb|added|
|csv_classifier_allow_single_column|boolean|removed|
|csv_classifier_contains_header|text|removed|
|csv_classifier_creation_time|timestamp without time zone|removed|
|csv_classifier_delimiter|text|removed|
|csv_classifier_disable_value_trimming|boolean|removed|
|csv_classifier_header|text[]|removed|
|csv_classifier_last_updated|timestamp without time zone|removed|
|csv_classifier_name|text|removed|
|csv_classifier_quote_symbol|text|removed|
|csv_classifier_version|bigint|removed|
|grok_classifier|jsonb|added|
|grok_classifier_classification|text|removed|
|grok_classifier_creation_time|timestamp without time zone|removed|
|grok_classifier_custom_patterns|text|removed|
|grok_classifier_grok_pattern|text|removed|
|grok_classifier_last_updated|timestamp without time zone|removed|
|grok_classifier_name|text|removed|
|grok_classifier_version|bigint|removed|
|json_classifier|jsonb|added|
|json_classifier_creation_time|timestamp without time zone|removed|
|json_classifier_json_path|text|removed|
|json_classifier_last_updated|timestamp without time zone|removed|
|json_classifier_name|text|removed|
|json_classifier_version|bigint|removed|
|xml_classifier|jsonb|added|
|xml_classifier_classification|text|removed|
|xml_classifier_creation_time|timestamp without time zone|removed|
|xml_classifier_last_updated|timestamp without time zone|removed|
|xml_classifier_name|text|removed|
|xml_classifier_row_tag|text|removed|
|xml_classifier_version|bigint|removed|

## aws_glue_connections

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|availability_zone|text|removed|
|physical_connection_requirements|jsonb|added|
|security_group_id_list|text[]|removed|
|subnet_id|text|removed|

## aws_glue_crawler_targets_catalog_targets
Moved to JSON column on [aws_glue_crawlers](#aws_glue_crawlers)


## aws_glue_crawler_targets_delta_targets
Moved to JSON column on [aws_glue_crawlers](#aws_glue_crawlers)


## aws_glue_crawler_targets_dynamo_db_targets
Moved to JSON column on [aws_glue_crawlers](#aws_glue_crawlers)


## aws_glue_crawler_targets_jdbc_targets
Moved to JSON column on [aws_glue_crawlers](#aws_glue_crawlers)


## aws_glue_crawler_targets_mongo_db_targets
Moved to JSON column on [aws_glue_crawlers](#aws_glue_crawlers)


## aws_glue_crawler_targets_s3_targets
Moved to JSON column on [aws_glue_crawlers](#aws_glue_crawlers)


## aws_glue_crawlers

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|lake_formation_configuration|jsonb|added|
|lake_formation_configuration_account_id|text|removed|
|lake_formation_configuration_use_lake_formation_credentials|boolean|removed|
|last_crawl|jsonb|added|
|last_crawl_error_message|text|removed|
|last_crawl_log_group|text|removed|
|last_crawl_log_stream|text|removed|
|last_crawl_message_prefix|text|removed|
|last_crawl_start_time|timestamp without time zone|removed|
|last_crawl_status|text|removed|
|lineage_configuration|jsonb|added|
|lineage_configuration_crawler_lineage_settings|text|removed|
|recrawl_behavior|text|removed|
|recrawl_policy|jsonb|added|
|schedule|jsonb|added|
|schedule_expression|text|removed|
|schedule_state|text|removed|
|schema_change_policy|jsonb|added|
|schema_change_policy_delete_behavior|text|removed|
|schema_change_policy_update_behavior|text|removed|
|targets|jsonb|added|

## aws_glue_database_table_columns
Moved to JSON column on [aws_glue_databases](#aws_glue_databases)


## aws_glue_database_table_indexes

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|account_id|text|added|
|database_arn|text|added|
|database_table_cq_id|uuid|removed|
|database_table_name|text|added|
|region|text|added|

## aws_glue_database_table_partition_keys
Moved to JSON column on [aws_glue_databases](#aws_glue_databases)


## aws_glue_database_tables

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|account_id|text|added|
|additional_locations|text[]|removed|
|bucket_columns|text[]|removed|
|compressed|boolean|removed|
|database_arn|text|added|
|database_cq_id|uuid|removed|
|input_format|text|removed|
|location|text|removed|
|number_of_buckets|bigint|removed|
|output_format|text|removed|
|partition_keys|jsonb|added|
|region|text|added|
|schema_reference_schema_id|jsonb|removed|
|schema_reference_schema_version_id|text|removed|
|schema_reference_schema_version_number|bigint|removed|
|serde_info|jsonb|removed|
|skewed_info|jsonb|removed|
|sort_columns|jsonb|removed|
|storage_descriptor|jsonb|added|
|storage_parameters|jsonb|removed|
|stored_as_sub_directories|boolean|removed|
|target_table|jsonb|added|
|target_table_catalog_id|text|removed|
|target_table_database_name|text|removed|
|target_table_name|text|removed|

## aws_glue_databases

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|target_database|jsonb|added|
|target_database_catalog_id|text|removed|
|target_database_name|text|removed|

## aws_glue_datacatalog_encryption_settings

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|aws_kms_key_id|text|removed|
|connection_password_encryption|jsonb|added|
|encryption_at_rest|jsonb|added|
|encryption_at_rest_catalog_encryption_mode|text|removed|
|encryption_at_rest_sse_aws_kms_key_id|text|removed|
|return_connection_password_encrypted|boolean|removed|

## aws_glue_dev_endpoints

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|endpoint_name|text|added|
|name|text|removed|

## aws_glue_job_runs

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|account_id|text|added|
|dpu_seconds|real|updated|Type changed from float to real
|job_arn|text|added|
|job_cq_id|uuid|removed|
|max_capacity|real|updated|Type changed from float to real
|notification_property|jsonb|added|
|notification_property_notify_delay_after|bigint|removed|
|region|text|added|

## aws_glue_jobs

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|command|jsonb|added|
|command_name|text|removed|
|command_python_version|text|removed|
|command_script_location|text|removed|
|connections|jsonb|updated|Type changed from text[] to jsonb
|execution_property|jsonb|added|
|execution_property_max_concurrent_runs|bigint|removed|
|max_capacity|real|updated|Type changed from float to real
|notification_property|jsonb|added|
|notification_property_notify_delay_after|bigint|removed|

## aws_glue_ml_transform_input_record_tables
Moved to JSON column on [aws_glue_ml_transforms](#aws_glue_ml_transforms)


## aws_glue_ml_transform_task_runs

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|account_id|text|added|
|export_labels_task_run_properties_output_s3_path|text|removed|
|find_matches_task_run_properties_job_id|text|removed|
|find_matches_task_run_properties_job_name|text|removed|
|find_matches_task_run_properties_job_run_id|text|removed|
|id|text|removed|
|import_labels_task_run_properties_input_s3_path|text|removed|
|import_labels_task_run_properties_replace|boolean|removed|
|labeling_set_generation_task_run_properties_output_s3_path|text|removed|
|ml_transform_arn|text|added|
|ml_transform_cq_id|uuid|removed|
|properties|jsonb|added|
|region|text|added|
|task_run_id|text|added|
|task_type|text|removed|

## aws_glue_ml_transforms

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|evaluation_metrics|jsonb|added|
|evaluation_metrics_find_matches_metrics_area_under_pr_curve|float|removed|
|evaluation_metrics_find_matches_metrics_column_importances|jsonb|removed|
|evaluation_metrics_find_matches_metrics_confusion_matrix|jsonb|removed|
|evaluation_metrics_find_matches_metrics_f1|float|removed|
|evaluation_metrics_find_matches_metrics_precision|float|removed|
|evaluation_metrics_find_matches_metrics_recall|float|removed|
|evaluation_metrics_transform_type|text|removed|
|id|text|removed|
|input_record_tables|jsonb|added|
|max_capacity|real|updated|Type changed from float to real
|parameters|jsonb|added|
|parameters_find_matches_parameters_accuracy_cost_tradeoff|float|removed|
|parameters_find_matches_parameters_enforce_provided_labels|boolean|removed|
|parameters_find_matches_parameters_precision_recall_tradeoff|float|removed|
|parameters_find_matches_parameters_primary_key_column_name|text|removed|
|parameters_transform_type|text|removed|
|transform_encryption|jsonb|added|
|transform_encryption_ml_user_data_encryption_kms_key_id|text|removed|
|transform_encryption_task_run_security_configuration_name|text|removed|
|transform_encryption_user_data_encryption_mode|text|removed|
|transform_id|text|added|

## aws_glue_registries

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|

## aws_glue_registry_schema_versions

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|account_id|text|added|
|id|text|removed|
|region|text|added|
|registry_schema_arn|text|added|
|registry_schema_cq_id|uuid|removed|
|result_metadata|jsonb|added|
|schema_version_id|text|added|

## aws_glue_registry_schemas

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|account_id|text|added|
|region|text|added|
|registry_cq_id|uuid|removed|
|result_metadata|jsonb|added|

## aws_glue_security_configuration_s3_encryption
Moved to JSON column on [aws_glue_security_configurations](#aws_glue_security_configurations)


## aws_glue_security_configurations

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|cloud_watch_encryption_kms_key_arn|text|removed|
|cloud_watch_encryption_mode|text|removed|
|encryption_configuration|jsonb|added|
|job_bookmarks_encryption_kms_key_arn|text|removed|
|job_bookmarks_encryption_mode|text|removed|

## aws_glue_trigger_actions
Moved to JSON column on [aws_glue_triggers](#aws_glue_triggers)


## aws_glue_trigger_predicate_conditions
Moved to JSON column on [aws_glue_triggers](#aws_glue_triggers)


## aws_glue_triggers

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|actions|jsonb|added|
|event_batching_condition|jsonb|added|
|event_batching_condition_size|bigint|removed|
|event_batching_condition_window|bigint|removed|
|predicate|jsonb|added|
|predicate_logical|text|removed|

## aws_glue_workflows

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|blueprint_details|jsonb|added|
|blueprint_name|text|removed|
|blueprint_run_id|text|removed|
|graph|jsonb|added|
|last_run|jsonb|added|
|last_run_completed_on|timestamp without time zone|removed|
|last_run_error_message|text|removed|
|last_run_name|text|removed|
|last_run_previous_run_id|text|removed|
|last_run_started_on|timestamp without time zone|removed|
|last_run_starting_event_batch_condition_size|bigint|removed|
|last_run_starting_event_batch_condition_window|bigint|removed|
|last_run_statistics_failed_actions|bigint|removed|
|last_run_statistics_running_actions|bigint|removed|
|last_run_statistics_stopped_actions|bigint|removed|
|last_run_statistics_succeeded_actions|bigint|removed|
|last_run_statistics_timeout_actions|bigint|removed|
|last_run_statistics_total_actions|bigint|removed|
|last_run_status|text|removed|
|last_run_workflow_run_id|text|removed|
|last_run_workflow_run_properties|jsonb|removed|

## aws_guardduty_detector_members

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|administrator_id|text|added|
|detector_arn|text|added|
|detector_cq_id|uuid|removed|
|invited_at|text|updated|Type changed from timestamp without time zone to text
|region|text|added|
|updated_at|text|updated|Type changed from timestamp without time zone to text

## aws_guardduty_detectors

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|created_at|text|updated|Type changed from timestamp without time zone to text
|data_sources|jsonb|added|
|data_sources_cloud_trail_status|text|removed|
|data_sources_dns_logs_status|text|removed|
|data_sources_flow_logs_status|text|removed|
|data_sources_s3_logs_status|text|removed|
|result_metadata|jsonb|added|
|updated_at|text|updated|Type changed from timestamp without time zone to text

## aws_iam_accounts
Renamed from [aws_accounts](aws_accounts)


## aws_iam_credential_reports
This table was newly added.

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|arn|text|added|
|user_creation_time|timestamp without time zone|added|
|password_last_changed|timestamp without time zone|added|
|password_next_rotation|timestamp without time zone|added|
|access_key_1_last_rotated|timestamp without time zone|added|
|access_key_2_last_rotated|timestamp without time zone|added|
|cert_1_last_rotated|timestamp without time zone|added|
|cert_2_last_rotated|timestamp without time zone|added|
|access_key_1_last_used_date|timestamp without time zone|added|
|access_key_2_last_used_date|timestamp without time zone|added|
|password_last_used|timestamp without time zone|added|
|user|text|added|
|password_status|text|added|
|mfa_active|boolean|added|
|access_key1_active|boolean|added|
|access_key2_active|boolean|added|
|cert1_active|boolean|added|
|cert2_active|boolean|added|
|access_key1_last_used_region|text|added|
|access_key1_last_used_service|text|added|
|access_key2_last_used_region|text|added|
|access_key2_last_used_service|text|added|
|_cq_id|uuid|added|
|_cq_fetch_time|timestamp without time zone|added|

## aws_iam_group_policies

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|group_arn|text|added|
|group_cq_id|uuid|removed|
|result_metadata|jsonb|added|

## aws_iam_groups

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|group_name|text|added|
|name|text|removed|

## aws_iam_openid_connect_identity_providers

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|result_metadata|jsonb|added|

## aws_iam_password_policies

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|max_password_age|bigint|updated|Type changed from integer to bigint
|minimum_password_length|bigint|updated|Type changed from integer to bigint
|password_reuse_prevention|bigint|updated|Type changed from integer to bigint

## aws_iam_policies

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|attachment_count|bigint|updated|Type changed from integer to bigint
|name|text|removed|
|permissions_boundary_usage_count|bigint|updated|Type changed from integer to bigint
|policy_name|text|added|
|policy_version_list|jsonb|added|

## aws_iam_policy_versions
Moved to JSON column on [aws_iam_policies](#aws_iam_policies)


## aws_iam_role_policies

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|account_id|text|added|
|result_metadata|jsonb|added|
|role_arn|text|added|
|role_cq_id|uuid|removed|

## aws_iam_roles

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|permissions_boundary|jsonb|added|
|permissions_boundary_arn|text|removed|
|permissions_boundary_type|text|removed|
|role_last_used|jsonb|added|
|role_last_used_last_used_date|timestamp without time zone|removed|
|role_last_used_region|text|removed|

## aws_iam_saml_identity_providers

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|saml_metadata_document|text|removed|

## aws_iam_server_certificates

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|name|text|removed|
|server_certificate_name|text|added|

## aws_iam_user_access_keys

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|account_id|text|added|
|user_arn|text|added|
|user_cq_id|uuid|removed|
|user_name|text|added|

## aws_iam_user_attached_policies

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|account_id|text|added|
|user_arn|text|added|
|user_cq_id|uuid|removed|

## aws_iam_user_groups

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|account_id|text|added|
|arn|text|added|
|group_arn|text|removed|
|user_arn|text|added|
|user_cq_id|uuid|removed|

## aws_iam_user_policies

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|result_metadata|jsonb|added|
|user_arn|text|added|
|user_cq_id|uuid|removed|

## aws_iam_users

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|access_key_1_active|boolean|removed|
|access_key_1_last_rotated|timestamp without time zone|removed|
|access_key_2_active|boolean|removed|
|access_key_2_last_rotated|timestamp without time zone|removed|
|cert_1_active|boolean|removed|
|cert_1_last_rotated|timestamp without time zone|removed|
|cert_2_active|boolean|removed|
|cert_2_last_rotated|timestamp without time zone|removed|
|mfa_active|boolean|removed|
|password_enabled|boolean|removed|
|password_last_changed|timestamp without time zone|removed|
|password_next_rotation|timestamp without time zone|removed|
|password_status|text|removed|
|permissions_boundary|jsonb|added|
|permissions_boundary_arn|text|removed|
|permissions_boundary_type|text|removed|
|user_id|text|removed|

## aws_iam_virtual_mfa_devices

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|base32_string_seed|bigint[]|updated|Type changed from bytea to bigint[]
|qr_code_png|bigint[]|updated|Type changed from bytea to bigint[]
|user|jsonb|added|
|user_arn|text|removed|
|user_create_date|timestamp without time zone|removed|
|user_id|text|removed|
|user_name|text|removed|
|user_password_last_used|timestamp without time zone|removed|
|user_path|text|removed|
|user_permissions_boundary_permissions_boundary_arn|text|removed|
|user_permissions_boundary_permissions_boundary_type|text|removed|

## aws_inspector2_finding_resources
Moved to JSON column on [aws_inspector2_findings](#aws_inspector2_findings)


## aws_inspector2_findings

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|aws_account_id|text|added|
|finding_arn|text|removed|
|inspector_score|real|updated|Type changed from float to real
|remediation|jsonb|added|
|remediation_recommendation_text|text|removed|
|remediation_recommendation_url|text|removed|
|resources|jsonb|added|

## aws_inspector_findings

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|numeric_severity|real|updated|Type changed from float to real

## aws_iot_billing_groups

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|billing_group_id|text|added|
|billing_group_metadata|jsonb|added|
|billing_group_name|text|added|
|billing_group_properties|jsonb|added|
|creation_date|timestamp without time zone|removed|
|description|text|removed|
|id|text|removed|
|name|text|removed|
|result_metadata|jsonb|added|

## aws_iot_ca_certificates

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|certificate_id|text|added|
|certificate_pem|text|added|
|customer_version|bigint|updated|Type changed from integer to bigint
|id|text|removed|
|pem|text|removed|
|validity|jsonb|added|
|validity_not_after|timestamp without time zone|removed|
|validity_not_before|timestamp without time zone|removed|

## aws_iot_certificates

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|certificate_id|text|added|
|certificate_mode|text|added|
|certificate_pem|text|added|
|customer_version|bigint|updated|Type changed from integer to bigint
|id|text|removed|
|mode|text|removed|
|pem|text|removed|
|transfer_data|jsonb|added|
|transfer_data_accept_date|timestamp without time zone|removed|
|transfer_data_reject_date|timestamp without time zone|removed|
|transfer_data_reject_reason|text|removed|
|transfer_data_transfer_date|timestamp without time zone|removed|
|transfer_data_transfer_message|text|removed|
|validity|jsonb|added|
|validity_not_after|timestamp without time zone|removed|
|validity_not_before|timestamp without time zone|removed|

## aws_iot_jobs
This table was newly added.

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|account_id|text|added|
|region|text|added|
|tags|jsonb|added|
|arn|text|added|
|abort_config|jsonb|added|
|comment|text|added|
|completed_at|timestamp without time zone|added|
|created_at|timestamp without time zone|added|
|description|text|added|
|document_parameters|jsonb|added|
|force_canceled|boolean|added|
|is_concurrent|boolean|added|
|job_executions_retry_config|jsonb|added|
|job_executions_rollout_config|jsonb|added|
|job_id|text|added|
|job_process_details|jsonb|added|
|job_template_arn|text|added|
|last_updated_at|timestamp without time zone|added|
|namespace_id|text|added|
|presigned_url_config|jsonb|added|
|reason_code|text|added|
|status|text|added|
|target_selection|text|added|
|targets|text[]|added|
|timeout_config|jsonb|added|
|_cq_id|uuid|added|
|_cq_fetch_time|timestamp without time zone|added|

## aws_iot_policies

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|creation_date|timestamp without time zone|removed|
|default_version_id|text|removed|
|document|text|removed|
|generation_id|text|removed|
|last_modified_date|timestamp without time zone|removed|
|name|text|removed|
|policy_name|text|added|

## aws_iot_security_profiles
This table was newly added.

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|account_id|text|added|
|region|text|added|
|targets|text[]|added|
|tags|jsonb|added|
|arn|text|added|
|additional_metrics_to_retain|text[]|added|
|additional_metrics_to_retain_v2|jsonb|added|
|alert_targets|jsonb|added|
|behaviors|jsonb|added|
|creation_date|timestamp without time zone|added|
|last_modified_date|timestamp without time zone|added|
|security_profile_description|text|added|
|security_profile_name|text|added|
|version|bigint|added|
|result_metadata|jsonb|added|
|_cq_id|uuid|added|
|_cq_fetch_time|timestamp without time zone|added|

## aws_iot_stream_files
Moved to JSON column on [aws_iot_streams](#aws_iot_streams)


## aws_iot_streams

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|files|jsonb|added|
|id|text|removed|
|stream_id|text|added|
|stream_version|bigint|added|
|version|integer|removed|

## aws_iot_thing_groups

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|attribute_payload_attributes|jsonb|removed|
|attribute_payload_merge|boolean|removed|
|creation_date|timestamp without time zone|removed|
|id|text|removed|
|name|text|removed|
|parent_group_name|text|removed|
|result_metadata|jsonb|added|
|root_to_parent_thing_groups|jsonb|removed|
|thing_group_description|text|removed|
|thing_group_id|text|added|
|thing_group_metadata|jsonb|added|
|thing_group_name|text|added|
|thing_group_properties|jsonb|added|

## aws_iot_thing_types

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|creation_date|timestamp without time zone|removed|
|deprecated|boolean|removed|
|deprecation_date|timestamp without time zone|removed|
|description|text|removed|
|name|text|removed|
|searchable_attributes|text[]|removed|
|thing_type_metadata|jsonb|added|
|thing_type_name|text|added|
|thing_type_properties|jsonb|added|

## aws_iot_things

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|name|text|removed|
|thing_name|text|added|
|thing_type_name|text|added|
|type_name|text|removed|

## aws_iot_topic_rule_actions
Moved to JSON column on [aws_iot_topic_rules](#aws_iot_topic_rules)


## aws_iot_topic_rules

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|aws_iot_sql_version|text|removed|
|created_at|timestamp without time zone|removed|
|description|text|removed|
|error_action_cloudwatch_alarm_name|text|removed|
|error_action_cloudwatch_alarm_role_arn|text|removed|
|error_action_cloudwatch_alarm_state_reason|text|removed|
|error_action_cloudwatch_alarm_state_value|text|removed|
|error_action_cloudwatch_logs_log_group_name|text|removed|
|error_action_cloudwatch_logs_role_arn|text|removed|
|error_action_cloudwatch_metric_metric_name|text|removed|
|error_action_cloudwatch_metric_metric_namespace|text|removed|
|error_action_cloudwatch_metric_role_arn|text|removed|
|error_action_cloudwatch_metric_timestamp|text|removed|
|error_action_cloudwatch_metric_unit|text|removed|
|error_action_cloudwatch_metric_value|text|removed|
|error_action_dynamo_db_hash_key_field|text|removed|
|error_action_dynamo_db_hash_key_type|text|removed|
|error_action_dynamo_db_hash_key_value|text|removed|
|error_action_dynamo_db_operation|text|removed|
|error_action_dynamo_db_payload_field|text|removed|
|error_action_dynamo_db_range_key_field|text|removed|
|error_action_dynamo_db_range_key_type|text|removed|
|error_action_dynamo_db_range_key_value|text|removed|
|error_action_dynamo_db_role_arn|text|removed|
|error_action_dynamo_db_table_name|text|removed|
|error_action_dynamo_db_v2_put_item_table_name|text|removed|
|error_action_dynamo_db_v2_role_arn|text|removed|
|error_action_elasticsearch_endpoint|text|removed|
|error_action_elasticsearch_id|text|removed|
|error_action_elasticsearch_index|text|removed|
|error_action_elasticsearch_role_arn|text|removed|
|error_action_elasticsearch_type|text|removed|
|error_action_firehose_batch_mode|boolean|removed|
|error_action_firehose_delivery_stream_name|text|removed|
|error_action_firehose_role_arn|text|removed|
|error_action_firehose_separator|text|removed|
|error_action_http_auth_sigv4_role_arn|text|removed|
|error_action_http_auth_sigv4_service_name|text|removed|
|error_action_http_auth_sigv4_signing_region|text|removed|
|error_action_http_confirmation_url|text|removed|
|error_action_http_headers|jsonb|removed|
|error_action_http_url|text|removed|
|error_action_iot_analytics_batch_mode|boolean|removed|
|error_action_iot_analytics_channel_arn|text|removed|
|error_action_iot_analytics_channel_name|text|removed|
|error_action_iot_analytics_role_arn|text|removed|
|error_action_iot_events_batch_mode|boolean|removed|
|error_action_iot_events_input_name|text|removed|
|error_action_iot_events_message_id|text|removed|
|error_action_iot_events_role_arn|text|removed|
|error_action_iot_site_wise|jsonb|removed|
|error_action_kafka_client_properties|jsonb|removed|
|error_action_kafka_destination_arn|text|removed|
|error_action_kafka_key|text|removed|
|error_action_kafka_partition|text|removed|
|error_action_kafka_topic|text|removed|
|error_action_kinesis_partition_key|text|removed|
|error_action_kinesis_role_arn|text|removed|
|error_action_kinesis_stream_name|text|removed|
|error_action_lambda_function_arn|text|removed|
|error_action_open_search_endpoint|text|removed|
|error_action_open_search_id|text|removed|
|error_action_open_search_index|text|removed|
|error_action_open_search_role_arn|text|removed|
|error_action_open_search_type|text|removed|
|error_action_republish_qos|integer|removed|
|error_action_republish_role_arn|text|removed|
|error_action_republish_topic|text|removed|
|error_action_s3_bucket_name|text|removed|
|error_action_s3_canned_acl|text|removed|
|error_action_s3_key|text|removed|
|error_action_s3_role_arn|text|removed|
|error_action_salesforce_token|text|removed|
|error_action_salesforce_url|text|removed|
|error_action_sns_message_format|text|removed|
|error_action_sns_role_arn|text|removed|
|error_action_sns_target_arn|text|removed|
|error_action_sqs_queue_url|text|removed|
|error_action_sqs_role_arn|text|removed|
|error_action_sqs_use_base64|boolean|removed|
|error_action_step_functions_execution_name_prefix|text|removed|
|error_action_step_functions_role_arn|text|removed|
|error_action_step_functions_state_machine_name|text|removed|
|error_action_timestream_database_name|text|removed|
|error_action_timestream_dimensions|jsonb|removed|
|error_action_timestream_role_arn|text|removed|
|error_action_timestream_table_name|text|removed|
|error_action_timestream_timestamp_unit|text|removed|
|error_action_timestream_timestamp_value|text|removed|
|result_metadata|jsonb|added|
|rule|jsonb|added|
|rule_disabled|boolean|removed|
|rule_name|text|removed|
|sql|text|removed|

## aws_kinesis_stream_enhanced_monitoring
Moved to JSON column on [aws_kinesis_streams](#aws_kinesis_streams)


## aws_kinesis_streams

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|enhanced_monitoring|jsonb|added|
|stream_arn|text|removed|
|stream_mode_details|jsonb|added|
|stream_mode_details_stream_mode|text|removed|

## aws_kms_keys

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|customer_master_key_spec|text|added|
|id|text|removed|
|key_id|text|added|
|key_manager|text|added|
|manager|text|removed|
|multi_region_configuration|jsonb|added|
|multi_region_key_type|text|removed|
|pending_deletion_window_in_days|bigint|updated|Type changed from integer to bigint
|primary_key_arn|text|removed|
|primary_key_region|text|removed|

## aws_lambda_function_aliases

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|account_id|text|added|
|alias_arn|text|added|
|arn|text|removed|
|function_cq_id|uuid|removed|
|region|text|added|
|routing_config|jsonb|added|
|routing_config_additional_version_weights|jsonb|removed|
|url_config|jsonb|added|
|url_config_auth_type|text|removed|
|url_config_cors|jsonb|removed|
|url_config_creation_time|timestamp without time zone|removed|
|url_config_function_arn|text|removed|
|url_config_function_url|text|removed|
|url_config_last_modified_time|timestamp without time zone|removed|

## aws_lambda_function_concurrency_configs

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|account_id|text|added|
|allocated_provisioned_concurrent_executions|bigint|updated|Type changed from integer to bigint
|available_provisioned_concurrent_executions|bigint|updated|Type changed from integer to bigint
|function_cq_id|uuid|removed|
|last_modified|text|updated|Type changed from timestamp without time zone to text
|region|text|added|
|requested_provisioned_concurrent_executions|bigint|updated|Type changed from integer to bigint

## aws_lambda_function_event_invoke_configs

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|account_id|text|added|
|destination_config|jsonb|added|
|function_cq_id|uuid|removed|
|maximum_event_age_in_seconds|bigint|updated|Type changed from integer to bigint
|maximum_retry_attempts|bigint|updated|Type changed from integer to bigint
|on_failure_destination|text|removed|
|on_success_destination|text|removed|
|region|text|added|

## aws_lambda_function_event_source_mappings

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|account_id|text|added|
|batch_size|bigint|updated|Type changed from integer to bigint
|criteria_filters|text[]|removed|
|destination_config|jsonb|added|
|filter_criteria|jsonb|added|
|function_cq_id|uuid|removed|
|maximum_batching_window_in_seconds|bigint|updated|Type changed from integer to bigint
|maximum_record_age_in_seconds|bigint|updated|Type changed from integer to bigint
|maximum_retry_attempts|bigint|updated|Type changed from integer to bigint
|on_failure_destination|text|removed|
|on_success_destination|text|removed|
|parallelization_factor|bigint|updated|Type changed from integer to bigint
|region|text|added|
|self_managed_event_source|jsonb|added|
|self_managed_event_source_endpoints|jsonb|removed|
|tumbling_window_in_seconds|bigint|updated|Type changed from integer to bigint

## aws_lambda_function_file_system_configs
Moved to JSON column on [aws_lambda_functions](#aws_lambda_functions)


## aws_lambda_function_layers
Moved to JSON column on [aws_lambda_functions](#aws_lambda_functions)


## aws_lambda_function_version_file_system_configs
Moved to JSON column on [aws_lambda_functions](#aws_lambda_functions)


## aws_lambda_function_version_layers
Moved to JSON column on [aws_lambda_functions](#aws_lambda_functions)


## aws_lambda_function_versions

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|account_id|text|added|
|dead_letter_config|jsonb|added|
|dead_letter_config_target_arn|text|removed|
|environment|jsonb|added|
|environment_error_error_code|text|removed|
|environment_error_message|text|removed|
|environment_variables|jsonb|removed|
|ephemeral_storage|jsonb|added|
|ephemeral_storage_size|integer|removed|
|error_code|text|removed|
|error_message|text|removed|
|file_system_configs|jsonb|added|
|function_cq_id|uuid|removed|
|image_config_command|text[]|removed|
|image_config_entry_point|text[]|removed|
|image_config_response|jsonb|added|
|image_config_working_directory|text|removed|
|last_modified|text|updated|Type changed from timestamp without time zone to text
|layers|jsonb|added|
|memory_size|bigint|updated|Type changed from integer to bigint
|region|text|added|
|timeout|bigint|updated|Type changed from integer to bigint
|tracing_config|jsonb|added|
|tracing_config_mode|text|removed|
|vpc_config|jsonb|added|
|vpc_config_security_group_ids|text[]|removed|
|vpc_config_subnet_ids|text[]|removed|
|vpc_config_vpc_id|text|removed|

## aws_lambda_functions

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|architectures|text[]|removed|
|code|jsonb|added|
|code_image_uri|text|removed|
|code_location|text|removed|
|code_resolved_image_uri|text|removed|
|code_sha256|text|removed|
|code_signing_allowed_publishers_version_arns|text[]|removed|
|code_signing_config|jsonb|added|
|code_signing_config_arn|text|removed|
|code_signing_config_id|text|removed|
|code_signing_description|text|removed|
|code_signing_last_modified|timestamp without time zone|removed|
|code_signing_policies_untrusted_artifact_on_deployment|text|removed|
|code_size|bigint|removed|
|concurrency|jsonb|added|
|concurrency_reserved_concurrent_executions|integer|removed|
|configuration|jsonb|added|
|dead_letter_config_target_arn|text|removed|
|description|text|removed|
|environment_error_code|text|removed|
|environment_error_message|text|removed|
|environment_variables|jsonb|removed|
|ephemeral_storage_size|integer|removed|
|error_code|text|removed|
|error_message|text|removed|
|handler|text|removed|
|image_config_command|text[]|removed|
|image_config_entry_point|text[]|removed|
|image_config_working_directory|text|removed|
|kms_key_arn|text|removed|
|last_modified|timestamp without time zone|removed|
|last_update_status|text|removed|
|last_update_status_reason|text|removed|
|last_update_status_reason_code|text|removed|
|master_arn|text|removed|
|memory_size|integer|removed|
|name|text|removed|
|package_type|text|removed|
|result_metadata|jsonb|added|
|revision_id|text|removed|
|role|text|removed|
|runtime|text|removed|
|signing_job_arn|text|removed|
|signing_profile_version_arn|text|removed|
|state|text|removed|
|state_reason|text|removed|
|state_reason_code|text|removed|
|timeout|integer|removed|
|tracing_config_mode|text|removed|
|version|text|removed|
|vpc_config_security_group_ids|text[]|removed|
|vpc_config_subnet_ids|text[]|removed|
|vpc_config_vpc_id|text|removed|

## aws_lambda_layer_version_policies

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|account_id|text|added|
|layer_version_arn|text|added|
|layer_version_cq_id|uuid|removed|
|region|text|added|
|result_metadata|jsonb|added|

## aws_lambda_layer_versions

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|account_id|text|added|
|arn|text|added|
|compatible_architectures|text[]|added|
|created_date|text|updated|Type changed from timestamp without time zone to text
|layer_arn|text|added|
|layer_cq_id|uuid|removed|
|layer_version_arn|text|removed|
|region|text|added|

## aws_lambda_layers

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|latest_matching_version|jsonb|updated|Type changed from bigint to jsonb
|latest_matching_version_compatible_runtimes|text[]|removed|
|latest_matching_version_created_date|timestamp without time zone|removed|
|latest_matching_version_description|text|removed|
|latest_matching_version_layer_version_arn|text|removed|
|latest_matching_version_license_info|text|removed|
|layer_arn|text|added|
|layer_name|text|added|
|name|text|removed|

## aws_lambda_runtimes

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|account_id|text|added|
|region|text|added|

## aws_lightsail_alarms

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|availability_zone|text|removed|
|location|jsonb|added|
|monitored_resource_info|jsonb|added|
|monitored_resource_info_arn|text|removed|
|monitored_resource_name|text|removed|
|monitored_resource_resource_type|text|removed|
|threshold|real|updated|Type changed from float to real

## aws_lightsail_bucket_access_keys

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|account_id|text|added|
|bucket_arn|text|added|
|bucket_cq_id|uuid|removed|
|last_used|jsonb|added|
|last_used_date|timestamp without time zone|removed|
|last_used_region|text|removed|
|last_used_service_name|text|removed|
|region|text|added|

## aws_lightsail_buckets

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|access_log_config|jsonb|added|
|access_log_config_destination|text|removed|
|access_log_config_enabled|boolean|removed|
|access_log_config_prefix|text|removed|
|access_rules|jsonb|added|
|access_rules_allow_public_overrides|boolean|removed|
|access_rules_get_object|text|removed|
|location|jsonb|added|
|location_availability_zone|text|removed|
|location_region_name|text|removed|
|state|jsonb|added|
|state_code|text|removed|
|state_message|text|removed|

## aws_lightsail_certificate_domain_validation_records
Moved to JSON column on [aws_lightsail_certificates](#aws_lightsail_certificates)


## aws_lightsail_certificate_renewal_summary_domain_validation_records
Moved to JSON column on [aws_lightsail_certificates](#aws_lightsail_certificates)


## aws_lightsail_certificates

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|domain_validation_records|jsonb|added|
|renewal_summary|jsonb|added|
|renewal_summary_reason|text|removed|
|renewal_summary_status|text|removed|
|renewal_summary_updated_at|timestamp without time zone|removed|

## aws_lightsail_container_service_deployments

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|account_id|text|added|
|container_service_arn|text|added|
|container_service_cq_id|uuid|removed|
|public_endpoint|jsonb|added|
|public_endpoint_container_name|text|removed|
|public_endpoint_container_port|bigint|removed|
|public_endpoint_health_check|jsonb|removed|
|region|text|added|

## aws_lightsail_container_service_images

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|account_id|text|added|
|container_service_arn|text|added|
|container_service_cq_id|uuid|removed|
|region|text|added|

## aws_lightsail_container_services

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|availability_zone|text|removed|
|current_deployment|jsonb|added|
|current_deployment_containers|jsonb|removed|
|current_deployment_created_at|timestamp without time zone|removed|
|current_deployment_public_endpoint_container_name|text|removed|
|current_deployment_public_endpoint_container_port|bigint|removed|
|current_deployment_public_endpoint_health_check|jsonb|removed|
|current_deployment_state|text|removed|
|current_deployment_version|bigint|removed|
|location|jsonb|added|
|next_deployment|jsonb|added|
|next_deployment_containers|jsonb|removed|
|next_deployment_created_at|timestamp without time zone|removed|
|next_deployment_public_endpoint_container_name|text|removed|
|next_deployment_public_endpoint_container_port|bigint|removed|
|next_deployment_public_endpoint_health_check|jsonb|removed|
|next_deployment_state|text|removed|
|next_deployment_version|bigint|removed|
|private_registry_access|jsonb|added|
|private_registry_access_ecr_image_puller_role_is_active|boolean|removed|
|private_registry_access_ecr_image_puller_role_principal_arn|text|removed|
|state_detail|jsonb|added|
|state_detail_code|text|removed|
|state_detail_message|text|removed|

## aws_lightsail_database_events

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|account_id|text|added|
|database_arn|text|added|
|database_cq_id|uuid|removed|
|region|text|added|

## aws_lightsail_database_log_events

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|account_id|text|added|
|database_arn|text|added|
|database_cq_id|uuid|removed|
|region|text|added|

## aws_lightsail_database_parameters

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|account_id|text|added|
|database_arn|text|added|
|database_cq_id|uuid|removed|
|name|text|removed|
|parameter_name|text|added|
|parameter_value|text|added|
|region|text|added|
|value|text|removed|

## aws_lightsail_database_pending_maintenance_actions
Moved to JSON column on [aws_lightsail_databases](#aws_lightsail_databases)


## aws_lightsail_database_snapshots

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|availability_zone|text|removed|
|location|jsonb|added|

## aws_lightsail_databases

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|availability_zone|text|removed|
|hardware|jsonb|added|
|hardware_cpu_count|bigint|removed|
|hardware_disk_size_in_gb|bigint|removed|
|hardware_ram_size_in_gb|float|removed|
|location|jsonb|added|
|master_endpoint|jsonb|added|
|master_endpoint_address|text|removed|
|master_endpoint_port|bigint|removed|
|pending_maintenance_actions|jsonb|added|

## aws_lightsail_disk_add_ons
Moved to JSON column on [aws_lightsail_disks](#aws_lightsail_disks)


## aws_lightsail_disk_snapshot

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|account_id|text|added|
|disk_arn|text|added|
|disk_cq_id|uuid|removed|
|location|jsonb|added|
|location_availability_zone|text|removed|
|location_region_name|text|removed|
|region|text|added|

## aws_lightsail_disks

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|add_ons|jsonb|added|
|location|jsonb|added|
|location_availability_zone|text|removed|
|location_region_name|text|removed|

## aws_lightsail_distributions

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|availability_zone|text|removed|
|cache_reset_create_time|timestamp without time zone|removed|
|cache_reset_status|text|removed|
|default_cache_behavior|jsonb|updated|Type changed from text to jsonb
|latest_cache_reset|jsonb|added|
|location|jsonb|added|
|origin|jsonb|added|
|origin_name|text|removed|
|origin_protocol_policy|text|removed|
|origin_region_name|text|removed|
|origin_resource_type|text|removed|

## aws_lightsail_instance_add_ons
Moved to JSON column on [aws_lightsail_instances](#aws_lightsail_instances)


## aws_lightsail_instance_hardware_disk_add_ons
Moved to JSON column on [aws_lightsail_instances](#aws_lightsail_instances)


## aws_lightsail_instance_hardware_disks
Moved to JSON column on [aws_lightsail_instances](#aws_lightsail_instances)


## aws_lightsail_instance_networking_ports
Moved to JSON column on [aws_lightsail_instances](#aws_lightsail_instances)


## aws_lightsail_instance_port_states

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|account_id|text|added|
|instance_arn|text|added|
|instance_cq_id|uuid|removed|
|region|text|added|

## aws_lightsail_instance_snapshot_from_attached_disk_add_ons
Moved to JSON column on [aws_lightsail_instances](#aws_lightsail_instances)


## aws_lightsail_instance_snapshot_from_attached_disks
Moved to JSON column on [aws_lightsail_instances](#aws_lightsail_instances)


## aws_lightsail_instance_snapshots

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|availability_zone|text|removed|
|from_attached_disks|jsonb|added|
|location|jsonb|added|

## aws_lightsail_instances

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|add_ons|jsonb|added|
|hardware|jsonb|added|
|hardware_cpu_count|bigint|removed|
|hardware_ram_size_in_gb|float|removed|
|location|jsonb|added|
|location_availability_zone|text|removed|
|location_region_name|text|removed|
|networking|jsonb|added|
|networking_monthly_transfer_gb_per_month_allocated|bigint|removed|
|state|jsonb|added|
|state_code|bigint|removed|
|state_name|text|removed|

## aws_lightsail_load_balancer_instance_health_summary
Moved to JSON column on [aws_lightsail_load_balancers](#aws_lightsail_load_balancers)


## aws_lightsail_load_balancer_tls_certificate_summaries
Moved to JSON column on [aws_lightsail_load_balancers](#aws_lightsail_load_balancers)


## aws_lightsail_load_balancer_tls_certificates

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|account_id|text|added|
|availability_zone|text|removed|
|load_balancer_arn|text|added|
|load_balancer_cq_id|uuid|removed|
|location|jsonb|added|
|region|text|added|
|region_name|text|removed|
|renewal_summary|jsonb|added|
|renewal_summary_domain_validation_options|jsonb|removed|
|renewal_summary_renewal_status|text|removed|

## aws_lightsail_load_balancers

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|availability_zone|text|removed|
|instance_health_summary|jsonb|added|
|location|jsonb|added|
|public_ports|bigint[]|updated|Type changed from integer[] to bigint[]
|tls_certificate_summaries|jsonb|added|

## aws_lightsail_static_ips

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|availability_zone|text|removed|
|location|jsonb|added|

## aws_mq_broker_configuration_revisions

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|account_id|text|added|
|broker_configuration_arn|text|added|
|broker_configuration_cq_id|uuid|removed|
|region|text|added|
|result_metadata|jsonb|added|

## aws_mq_broker_configurations

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|broker_arn|text|added|
|broker_cq_id|uuid|removed|
|latest_revision|jsonb|updated|Type changed from integer to jsonb
|latest_revision_created|timestamp without time zone|removed|
|latest_revision_description|text|removed|

## aws_mq_broker_users

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|broker_arn|text|added|
|broker_cq_id|uuid|removed|
|broker_id|text|added|
|result_metadata|jsonb|added|

## aws_mq_brokers

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|actions_required|jsonb|added|
|broker_id|text|added|
|configurations|jsonb|added|
|encryption_options|jsonb|added|
|encryption_options_kms_key_id|text|removed|
|encryption_options_use_aws_owned_key|boolean|removed|
|id|text|removed|
|result_metadata|jsonb|added|
|users|jsonb|added|

## aws_organizations_accounts

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|

## aws_qldb_ledger_journal_kinesis_streams

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|account_id|text|added|
|aggregation_enabled|boolean|removed|
|kinesis_configuration|jsonb|added|
|ledger_arn|text|added|
|ledger_cq_id|uuid|removed|
|region|text|added|
|stream_arn|text|removed|

## aws_qldb_ledger_journal_s3_exports

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|account_id|text|added|
|bucket|text|removed|
|kms_key_arn|text|removed|
|ledger_arn|text|added|
|ledger_cq_id|uuid|removed|
|object_encryption_type|text|removed|
|prefix|text|removed|
|region|text|added|
|s3_export_configuration|jsonb|added|

## aws_qldb_ledgers

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|encryption_description|jsonb|added|
|encryption_status|text|removed|
|inaccessible_kms_key_date_time|timestamp without time zone|removed|
|kms_key_arn|text|removed|
|result_metadata|jsonb|added|

## aws_rds_certificates

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|

## aws_rds_cluster_associated_roles
Moved to JSON column on [aws_rds_clusters](#aws_rds_clusters)


## aws_rds_cluster_db_cluster_members
Moved to JSON column on [aws_rds_clusters](#aws_rds_clusters)


## aws_rds_cluster_domain_memberships
Moved to JSON column on [aws_rds_clusters](#aws_rds_clusters)


## aws_rds_cluster_parameter_group_parameters
This table was newly added.

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|account_id|text|added|
|region|text|added|
|cluster_parameter_group_arn|text|added|
|allowed_values|text|added|
|apply_method|text|added|
|apply_type|text|added|
|data_type|text|added|
|description|text|added|
|is_modifiable|boolean|added|
|minimum_engine_version|text|added|
|parameter_name|text|added|
|parameter_value|text|added|
|source|text|added|
|supported_engine_modes|text[]|added|
|_cq_id|uuid|added|
|_cq_fetch_time|timestamp without time zone|added|

## aws_rds_cluster_parameter_groups

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|db_cluster_parameter_group_name|text|added|
|db_parameter_group_family|text|added|
|family|text|removed|
|name|text|removed|

## aws_rds_cluster_parameters
Moved to JSON column on [aws_rds_clusters](#aws_rds_clusters)


## aws_rds_cluster_snapshots

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|allocated_storage|bigint|updated|Type changed from integer to bigint
|percent_progress|bigint|updated|Type changed from integer to bigint
|port|bigint|updated|Type changed from integer to bigint
|tag_list|jsonb|added|

## aws_rds_cluster_vpc_security_groups
Moved to JSON column on [aws_rds_clusters](#aws_rds_clusters)


## aws_rds_clusters

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|allocated_storage|bigint|updated|Type changed from integer to bigint
|associated_roles|jsonb|added|
|auto_minor_version_upgrade|boolean|added|
|automatic_restart_time|timestamp without time zone|added|
|backup_retention_period|bigint|updated|Type changed from integer to bigint
|capacity|bigint|updated|Type changed from integer to bigint
|db_cluster_instance_class|text|added|
|db_cluster_members|jsonb|added|
|db_cluster_resource_id|text|added|
|domain_memberships|jsonb|added|
|id|text|removed|
|iops|bigint|added|
|monitoring_interval|bigint|added|
|monitoring_role_arn|text|added|
|pending_cloudwatch_logs_types_to_disable|text[]|removed|
|pending_cloudwatch_logs_types_to_enable|text[]|removed|
|pending_modified_values|jsonb|added|
|pending_modified_values_db_cluster_identifier|text|removed|
|pending_modified_values_engine_version|text|removed|
|pending_modified_values_iam_database_authentication_enabled|boolean|removed|
|pending_modified_values_master_user_password|text|removed|
|performance_insights_enabled|boolean|added|
|performance_insights_kms_key_id|text|added|
|performance_insights_retention_period|bigint|added|
|port|bigint|updated|Type changed from integer to bigint
|publicly_accessible|boolean|added|
|scaling_configuration_info|jsonb|added|
|scaling_configuration_info_auto_pause|boolean|removed|
|scaling_configuration_info_max_capacity|integer|removed|
|scaling_configuration_info_min_capacity|integer|removed|
|scaling_configuration_info_seconds_until_auto_pause|integer|removed|
|scaling_configuration_info_timeout_action|text|removed|
|serverless_v2_scaling_configuration|jsonb|added|
|storage_type|text|added|
|tag_list|jsonb|added|
|vpc_security_groups|jsonb|added|

## aws_rds_db_parameter_group_db_parameters
This table was newly added.

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|account_id|text|added|
|region|text|added|
|db_parameter_group_arn|text|added|
|allowed_values|text|added|
|apply_method|text|added|
|apply_type|text|added|
|data_type|text|added|
|description|text|added|
|is_modifiable|boolean|added|
|minimum_engine_version|text|added|
|parameter_name|text|added|
|parameter_value|text|added|
|source|text|added|
|supported_engine_modes|text[]|added|
|_cq_id|uuid|added|
|_cq_fetch_time|timestamp without time zone|added|

## aws_rds_db_parameter_groups

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|db_parameter_group_family|text|added|
|db_parameter_group_name|text|added|
|family|text|removed|
|name|text|removed|

## aws_rds_db_parameters
This table was removed.


## aws_rds_db_security_groups

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|db_security_group_description|text|added|
|db_security_group_name|text|added|
|description|text|removed|
|name|text|removed|

## aws_rds_db_snapshots

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|allocated_storage|bigint|updated|Type changed from integer to bigint
|iops|bigint|updated|Type changed from integer to bigint
|original_snapshot_create_time|timestamp without time zone|added|
|percent_progress|bigint|updated|Type changed from integer to bigint
|port|bigint|updated|Type changed from integer to bigint
|snapshot_target|text|added|
|tag_list|jsonb|added|

## aws_rds_event_subscriptions

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|

## aws_rds_instance_associated_roles
Moved to JSON column on [aws_rds_instances](#aws_rds_instances)


## aws_rds_instance_db_instance_automated_backups_replications
Moved to JSON column on [aws_rds_instances](#aws_rds_instances)


## aws_rds_instance_db_parameter_groups
Moved to JSON column on [aws_rds_instances](#aws_rds_instances)


## aws_rds_instance_db_security_groups
Moved to JSON column on [aws_rds_instances](#aws_rds_instances)


## aws_rds_instance_db_subnet_group_subnets
Moved to JSON column on [aws_rds_instances](#aws_rds_instances)


## aws_rds_instance_domain_memberships
Moved to JSON column on [aws_rds_instances](#aws_rds_instances)


## aws_rds_instance_option_group_memberships
Moved to JSON column on [aws_rds_instances](#aws_rds_instances)


## aws_rds_instance_vpc_security_groups
Moved to JSON column on [aws_rds_instances](#aws_rds_instances)


## aws_rds_instances

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|activity_stream_engine_native_audit_fields_included|boolean|added|
|activity_stream_kinesis_stream_name|text|added|
|activity_stream_kms_key_id|text|added|
|activity_stream_mode|text|added|
|activity_stream_status|text|added|
|allocated_storage|bigint|updated|Type changed from integer to bigint
|associated_roles|jsonb|added|
|automatic_restart_time|timestamp without time zone|added|
|automation_mode|text|added|
|backup_retention_period|bigint|updated|Type changed from integer to bigint
|backup_target|text|added|
|cluster_identifier|text|removed|
|custom_iam_instance_profile|text|added|
|db_cluster_identifier|text|added|
|db_instance_automated_backups_replications|jsonb|added|
|db_instance_identifier|text|added|
|db_instance_port|bigint|added|
|db_parameter_groups|jsonb|added|
|db_security_groups|jsonb|added|
|db_subnet_group|jsonb|added|
|dbi_resource_id|text|added|
|domain_memberships|jsonb|added|
|endpoint|jsonb|added|
|endpoint_address|text|removed|
|endpoint_hosted_zone_id|text|removed|
|endpoint_port|integer|removed|
|id|text|removed|
|instance_port|integer|removed|
|iops|bigint|updated|Type changed from integer to bigint
|listener_endpoint|jsonb|added|
|listener_endpoint_address|text|removed|
|listener_endpoint_hosted_zone_id|text|removed|
|listener_endpoint_port|integer|removed|
|max_allocated_storage|bigint|updated|Type changed from integer to bigint
|monitoring_interval|bigint|updated|Type changed from integer to bigint
|network_type|text|added|
|option_group_memberships|jsonb|added|
|pending_cloudwatch_logs_types_to_disable|text[]|removed|
|pending_cloudwatch_logs_types_to_enable|text[]|removed|
|pending_modified_values|jsonb|added|
|pending_modified_values_allocated_storage|integer|removed|
|pending_modified_values_backup_retention_period|integer|removed|
|pending_modified_values_ca_certificate_identifier|text|removed|
|pending_modified_values_db_instance_class|text|removed|
|pending_modified_values_db_instance_identifier|text|removed|
|pending_modified_values_db_subnet_group_name|text|removed|
|pending_modified_values_engine_version|text|removed|
|pending_modified_values_iam_database_authentication_enabled|boolean|removed|
|pending_modified_values_iops|integer|removed|
|pending_modified_values_license_model|text|removed|
|pending_modified_values_master_user_password|text|removed|
|pending_modified_values_multi_az|boolean|removed|
|pending_modified_values_port|integer|removed|
|pending_modified_values_processor_features|jsonb|removed|
|pending_modified_values_storage_type|text|removed|
|performance_insights_retention_period|bigint|updated|Type changed from integer to bigint
|promotion_tier|bigint|updated|Type changed from integer to bigint
|resume_full_automation_mode_time|timestamp without time zone|added|
|subnet_group_arn|text|removed|
|subnet_group_description|text|removed|
|subnet_group_name|text|removed|
|subnet_group_subnet_group_status|text|removed|
|subnet_group_vpc_id|text|removed|
|tag_list|jsonb|added|
|user_instance_id|text|removed|
|vpc_security_groups|jsonb|added|

## aws_rds_subnet_group_subnets
Moved to JSON column on [aws_rds_subnet_groups](#aws_rds_subnet_groups)


## aws_rds_subnet_groups

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|db_subnet_group_description|text|added|
|db_subnet_group_name|text|added|
|description|text|removed|
|name|text|removed|
|status|text|removed|
|subnet_group_status|text|added|
|subnets|jsonb|added|
|supported_network_types|text[]|added|

## aws_redshift_cluster_deferred_maintenance_windows
Moved to JSON column on [aws_redshift_clusters](#aws_redshift_clusters)


## aws_redshift_cluster_endpoint_vpc_endpoint_network_interfaces
Moved to JSON column on [aws_redshift_clusters](#aws_redshift_clusters)


## aws_redshift_cluster_endpoint_vpc_endpoints
Moved to JSON column on [aws_redshift_clusters](#aws_redshift_clusters)


## aws_redshift_cluster_iam_roles
Moved to JSON column on [aws_redshift_clusters](#aws_redshift_clusters)


## aws_redshift_cluster_nodes
Moved to JSON column on [aws_redshift_clusters](#aws_redshift_clusters)


## aws_redshift_cluster_parameter_group_status_lists
Moved to JSON column on [aws_redshift_clusters](#aws_redshift_clusters)


## aws_redshift_cluster_parameter_groups
Moved to JSON column on [aws_redshift_clusters](#aws_redshift_clusters)


## aws_redshift_cluster_parameters
Moved to JSON column on [aws_redshift_clusters](#aws_redshift_clusters)


## aws_redshift_cluster_security_groups
Moved to JSON column on [aws_redshift_clusters](#aws_redshift_clusters)


## aws_redshift_cluster_vpc_security_groups
Moved to JSON column on [aws_redshift_clusters](#aws_redshift_clusters)


## aws_redshift_clusters

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|aqua_configuration|jsonb|added|
|automated_snapshot_retention_period|bigint|updated|Type changed from integer to bigint
|cluster_identifier|text|added|
|cluster_nodes|jsonb|added|
|cluster_parameter_groups|jsonb|added|
|cluster_security_groups|jsonb|added|
|cluster_snapshot_copy_status|jsonb|added|
|cluster_snapshot_copy_status_destination_region|text|removed|
|cluster_snapshot_copy_status_manual_snapshot_retention_period|integer|removed|
|cluster_snapshot_copy_status_retention_period|bigint|removed|
|cluster_snapshot_copy_status_snapshot_copy_grant_name|text|removed|
|data_transfer_progress|jsonb|added|
|data_transfer_progress_current_rate_in_mega_bytes_per_second|float|removed|
|data_transfer_progress_data_transferred_in_mega_bytes|bigint|removed|
|data_transfer_progress_elapsed_time_in_seconds|bigint|removed|
|data_transfer_progress_estimated_time_to_completion_in_seconds|bigint|removed|
|data_transfer_progress_status|text|removed|
|data_transfer_progress_total_data_in_mega_bytes|bigint|removed|
|default_iam_role_arn|text|added|
|deferred_maintenance_windows|jsonb|added|
|elastic_ip_status|jsonb|updated|Type changed from text to jsonb
|elastic_ip_status_elastic_ip|text|removed|
|endpoint|jsonb|added|
|endpoint_address|text|removed|
|endpoint_port|integer|removed|
|hsm_status|jsonb|updated|Type changed from text to jsonb
|hsm_status_hsm_client_certificate_identifier|text|removed|
|hsm_status_hsm_configuration_identifier|text|removed|
|iam_roles|jsonb|added|
|id|text|removed|
|manual_snapshot_retention_period|bigint|updated|Type changed from integer to bigint
|number_of_nodes|bigint|updated|Type changed from integer to bigint
|pending_modified_values|jsonb|added|
|pending_modified_values_automated_snapshot_retention_period|integer|removed|
|pending_modified_values_cluster_identifier|text|removed|
|pending_modified_values_cluster_type|text|removed|
|pending_modified_values_cluster_version|text|removed|
|pending_modified_values_encryption_type|text|removed|
|pending_modified_values_enhanced_vpc_routing|boolean|removed|
|pending_modified_values_maintenance_track_name|text|removed|
|pending_modified_values_master_user_password|text|removed|
|pending_modified_values_node_type|text|removed|
|pending_modified_values_number_of_nodes|integer|removed|
|pending_modified_values_publicly_accessible|boolean|removed|
|reserved_node_exchange_status|jsonb|added|
|resize_info|jsonb|added|
|resize_info_allow_cancel_resize|boolean|removed|
|resize_info_resize_type|text|removed|
|restore_status|jsonb|updated|Type changed from text to jsonb
|restore_status_current_restore_rate_in_mega_bytes_per_second|float|removed|
|restore_status_elapsed_time_in_seconds|bigint|removed|
|restore_status_estimated_time_to_completion_in_seconds|bigint|removed|
|restore_status_progress_in_mega_bytes|bigint|removed|
|restore_status_snapshot_size_in_mega_bytes|bigint|removed|
|vpc_security_groups|jsonb|added|

## aws_redshift_event_subscriptions

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|cust_subscription_id|text|added|
|id|text|removed|

## aws_redshift_snapshot_accounts_with_restore_access
Moved to JSON column on [aws_redshift_snapshots](#aws_redshift_snapshots)


## aws_redshift_snapshots

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|account_id|text|added|
|accounts_with_restore_access|jsonb|added|
|actual_incremental_backup_size|float|removed|
|actual_incremental_backup_size_in_mega_bytes|real|added|
|backup_progress|float|removed|
|backup_progress_in_mega_bytes|real|added|
|cluster_cq_id|uuid|removed|
|current_backup_rate|float|removed|
|current_backup_rate_in_mega_bytes_per_second|real|added|
|elapsed_time|bigint|removed|
|elapsed_time_in_seconds|bigint|added|
|manual_snapshot_remaining_days|bigint|updated|Type changed from integer to bigint
|manual_snapshot_retention_period|bigint|updated|Type changed from integer to bigint
|number_of_nodes|bigint|updated|Type changed from integer to bigint
|port|bigint|updated|Type changed from integer to bigint
|region|text|added|
|total_backup_size_in_mega_bytes|real|updated|Type changed from float to real

## aws_redshift_subnet_group_subnets
Moved to JSON column on [aws_redshift_subnet_groups](#aws_redshift_subnet_groups)


## aws_redshift_subnet_groups

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|subnets|jsonb|added|

## aws_regions

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|

## aws_resourcegroups_resource_groups

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|arn|text|removed|
|description|text|added|
|group|text|removed|
|group_arn|text|added|
|group_description|text|removed|
|name|text|added|
|query|text|added|
|resource_query|text|removed|
|resource_query_type|text|removed|
|type|text|added|

## aws_route53_delegation_sets
This table was newly added.

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|account_id|text|added|
|arn|text|added|
|name_servers|text[]|added|
|caller_reference|text|added|
|id|text|added|
|_cq_id|uuid|added|
|_cq_fetch_time|timestamp without time zone|added|

## aws_route53_domain_nameservers
Moved to JSON column on [aws_route53_domains](#aws_route53_domains)


## aws_route53_domains

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|admin_contact|jsonb|added|
|admin_contact_address_line1|text|removed|
|admin_contact_address_line2|text|removed|
|admin_contact_city|text|removed|
|admin_contact_country_code|text|removed|
|admin_contact_email|text|removed|
|admin_contact_extra_params|jsonb|removed|
|admin_contact_fax|text|removed|
|admin_contact_first_name|text|removed|
|admin_contact_last_name|text|removed|
|admin_contact_organization_name|text|removed|
|admin_contact_phone_number|text|removed|
|admin_contact_state|text|removed|
|admin_contact_type|text|removed|
|admin_contact_zip_code|text|removed|
|nameservers|jsonb|added|
|registrant_contact|jsonb|added|
|registrant_contact_address_line1|text|removed|
|registrant_contact_address_line2|text|removed|
|registrant_contact_city|text|removed|
|registrant_contact_country_code|text|removed|
|registrant_contact_email|text|removed|
|registrant_contact_extra_params|jsonb|removed|
|registrant_contact_fax|text|removed|
|registrant_contact_first_name|text|removed|
|registrant_contact_last_name|text|removed|
|registrant_contact_organization_name|text|removed|
|registrant_contact_phone_number|text|removed|
|registrant_contact_state|text|removed|
|registrant_contact_type|text|removed|
|registrant_contact_zip_code|text|removed|
|status_list|jsonb|updated|Type changed from text[] to jsonb
|tech_contact|jsonb|added|
|tech_contact_address_line1|text|removed|
|tech_contact_address_line2|text|removed|
|tech_contact_city|text|removed|
|tech_contact_country_code|text|removed|
|tech_contact_email|text|removed|
|tech_contact_extra_params|jsonb|removed|
|tech_contact_fax|text|removed|
|tech_contact_first_name|text|removed|
|tech_contact_last_name|text|removed|
|tech_contact_organization_name|text|removed|
|tech_contact_phone_number|text|removed|
|tech_contact_state|text|removed|
|tech_contact_type|text|removed|
|tech_contact_zip_code|text|removed|

## aws_route53_health_checks

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|alarm_identifier_name|text|removed|
|alarm_identifier_region|text|removed|
|child_health_checks|text[]|removed|
|cloud_watch_alarm_config_comparison_operator|text|removed|
|cloud_watch_alarm_config_evaluation_periods|integer|removed|
|cloud_watch_alarm_config_metric_name|text|removed|
|cloud_watch_alarm_config_namespace|text|removed|
|cloud_watch_alarm_config_period|integer|removed|
|cloud_watch_alarm_config_statistic|text|removed|
|cloud_watch_alarm_config_threshold|float|removed|
|cloud_watch_alarm_configuration|jsonb|added|
|disabled|boolean|removed|
|enable_sni|boolean|removed|
|failure_threshold|integer|removed|
|fully_qualified_domain_name|text|removed|
|health_check_config|jsonb|added|
|health_threshold|integer|removed|
|insufficient_data_health_status|text|removed|
|inverted|boolean|removed|
|ip_address|text|removed|
|linked_service|jsonb|added|
|linked_service_description|text|removed|
|linked_service_service_principal|text|removed|
|measure_latency|boolean|removed|
|port|integer|removed|
|regions|text[]|removed|
|request_interval|integer|removed|
|resource_path|text|removed|
|search_string|text|removed|
|type|text|removed|

## aws_route53_hosted_zone_query_logging_configs

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|account_id|text|added|
|hosted_zone_arn|text|added|
|hosted_zone_cq_id|uuid|removed|
|hosted_zone_id|text|added|

## aws_route53_hosted_zone_resource_record_sets

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|account_id|text|added|
|alias_target|jsonb|added|
|cidr_routing_config|jsonb|added|
|dns_name|text|removed|
|evaluate_target_health|boolean|removed|
|geo_location|jsonb|added|
|geo_location_continent_code|text|removed|
|geo_location_country_code|text|removed|
|geo_location_subdivision_code|text|removed|
|hosted_zone_arn|text|added|
|hosted_zone_cq_id|uuid|removed|
|resource_records|jsonb|updated|Type changed from text[] to jsonb

## aws_route53_hosted_zone_traffic_policy_instances

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|account_id|text|added|
|hosted_zone_arn|text|added|
|hosted_zone_cq_id|uuid|removed|
|hosted_zone_id|text|added|
|traffic_policy_version|bigint|updated|Type changed from integer to bigint

## aws_route53_hosted_zone_vpc_association_authorizations
Moved to JSON column on [aws_route53_hosted_zones](#aws_route53_hosted_zones)


## aws_route53_hosted_zones

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|config|jsonb|added|
|config_comment|text|removed|
|config_private_zone|boolean|removed|
|delegation_set_id|text|removed|
|linked_service|jsonb|added|
|linked_service_description|text|removed|
|linked_service_principal|text|removed|
|tags|jsonb|removed|

## aws_route53_reusable_delegation_sets
This table was removed.


## aws_route53_traffic_policies

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|latest_version|bigint|updated|Type changed from integer to bigint
|traffic_policy_count|bigint|updated|Type changed from integer to bigint

## aws_route53_traffic_policy_versions

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|account_id|text|added|
|document|text|updated|Type changed from jsonb to text
|traffic_policy_arn|text|added|
|traffic_policy_cq_id|uuid|removed|
|version|bigint|updated|Type changed from integer to bigint

## aws_s3_account_config
Moved to JSON column on [aws_s3_accounts](#aws_s3_accounts)


## aws_s3_accounts
This table was newly added.

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|account_id|text|added|
|block_public_acls|boolean|added|
|block_public_policy|boolean|added|
|ignore_public_acls|boolean|added|
|restrict_public_buckets|boolean|added|
|config_exists|boolean|added|
|_cq_id|uuid|added|
|_cq_fetch_time|timestamp without time zone|added|

## aws_s3_bucket_cors_rules

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|account_id|text|added|
|bucket_arn|text|added|
|bucket_cq_id|uuid|removed|
|max_age_seconds|bigint|updated|Type changed from integer to bigint

## aws_s3_bucket_encryption_rules

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|account_id|text|added|
|apply_server_side_encryption_by_default|jsonb|added|
|bucket_arn|text|added|
|bucket_cq_id|uuid|removed|
|kms_master_key_id|text|removed|
|sse_algorithm|text|removed|

## aws_s3_bucket_grants

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|account_id|text|added|
|bucket_arn|text|added|
|bucket_cq_id|uuid|removed|
|display_name|text|removed|
|email_address|text|removed|
|grantee|jsonb|added|
|grantee_id|text|removed|
|type|text|removed|
|uri|text|removed|

## aws_s3_bucket_lifecycles

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|abort_incomplete_multipart_upload|jsonb|added|
|abort_incomplete_multipart_upload_days_after_initiation|integer|removed|
|account_id|text|added|
|bucket_arn|text|added|
|bucket_cq_id|uuid|removed|
|expiration|jsonb|added|
|expiration_date|timestamp without time zone|removed|
|expiration_days|integer|removed|
|expiration_expired_object_delete_marker|boolean|removed|
|filter|jsonb|removed|
|noncurrent_version_expiration|jsonb|added|
|noncurrent_version_expiration_days|integer|removed|

## aws_s3_bucket_replication_rules
Moved to JSON column on [aws_s3_buckets](#aws_s3_buckets)


## aws_s3_buckets

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|replication_rules|jsonb|added|

## aws_sagemaker_endpoint_configuration_production_variants
Moved to JSON column on [aws_sagemaker_endpoint_configurations](#aws_sagemaker_endpoint_configurations)


## aws_sagemaker_endpoint_configurations

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|async_inference_config|jsonb|added|
|endpoint_config_name|text|added|
|name|text|removed|
|production_variants|jsonb|added|
|result_metadata|jsonb|added|

## aws_sagemaker_model_containers
Moved to JSON column on [aws_sagemaker_models](#aws_sagemaker_models)


## aws_sagemaker_model_vpc_config
Moved to JSON column on [aws_sagemaker_models](#aws_sagemaker_models)


## aws_sagemaker_models

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|containers|jsonb|added|
|model_name|text|added|
|name|text|removed|
|result_metadata|jsonb|added|
|vpc_config|jsonb|added|

## aws_sagemaker_notebook_instances

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|direct_internet_access|text|updated|Type changed from boolean to text
|failure_reason|text|added|
|instance_metadata_service_configuration|jsonb|added|
|name|text|removed|
|notebook_instance_name|text|added|
|platform_identifier|text|added|
|result_metadata|jsonb|added|
|role_arn|text|added|
|root_access|text|added|
|security_groups|text[]|updated|Type changed from jsonb to text[]
|volume_size_in_gb|bigint|updated|Type changed from integer to bigint

## aws_sagemaker_training_job_algorithm_specification
Moved to JSON column on [aws_sagemaker_training_jobs](#aws_sagemaker_training_jobs)


## aws_sagemaker_training_job_debug_hook_config
Moved to JSON column on [aws_sagemaker_training_jobs](#aws_sagemaker_training_jobs)


## aws_sagemaker_training_job_debug_rule_configurations
Moved to JSON column on [aws_sagemaker_training_jobs](#aws_sagemaker_training_jobs)


## aws_sagemaker_training_job_debug_rule_evaluation_statuses
Moved to JSON column on [aws_sagemaker_training_jobs](#aws_sagemaker_training_jobs)


## aws_sagemaker_training_job_input_data_config
Moved to JSON column on [aws_sagemaker_training_jobs](#aws_sagemaker_training_jobs)


## aws_sagemaker_training_job_profiler_rule_configurations
Moved to JSON column on [aws_sagemaker_training_jobs](#aws_sagemaker_training_jobs)


## aws_sagemaker_training_job_profiler_rule_evaluation_statuses
Moved to JSON column on [aws_sagemaker_training_jobs](#aws_sagemaker_training_jobs)


## aws_sagemaker_training_jobs

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|algorithm_specification|jsonb|added|
|billable_time_in_seconds|bigint|updated|Type changed from integer to bigint
|debug_hook_config|jsonb|added|
|debug_rule_configurations|jsonb|added|
|debug_rule_evaluation_statuses|jsonb|added|
|input_data_config|jsonb|added|
|name|text|removed|
|profiler_rule_configurations|jsonb|added|
|profiler_rule_evaluation_statuses|jsonb|added|
|result_metadata|jsonb|added|
|retry_strategy|jsonb|added|
|training_job_name|text|added|
|training_time_in_seconds|bigint|updated|Type changed from integer to bigint

## aws_secretsmanager_secrets

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|rotation_rules|jsonb|added|
|rotation_rules_automatically_after_days|bigint|removed|
|secret_versions_to_stages|jsonb|removed|
|version_ids_to_stages|jsonb|added|

## aws_ses_templates

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|name|text|removed|
|template_name|text|added|

## aws_shield_attack_properties
Moved to JSON column on [aws_shield_attacks](#aws_shield_attacks)


## aws_shield_attack_sub_resources
Moved to JSON column on [aws_shield_attacks](#aws_shield_attacks)


## aws_shield_attacks

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|attack_properties|jsonb|added|
|mitigations|jsonb|updated|Type changed from text[] to jsonb
|sub_resources|jsonb|added|

## aws_shield_protection_groups

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|id|text|removed|
|protection_group_id|text|added|

## aws_shield_protections

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|application_automatic_response_configuration_status|text|removed|
|application_layer_automatic_response_configuration|jsonb|added|
|region|text|removed|

## aws_shield_subscriptions

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|protected_resource_type_limits|jsonb|removed|
|protection_group_limits_arbitrary_pattern_limits_max_members|integer|removed|
|protection_group_limits_max_protection_groups|integer|removed|
|subscription_limits|jsonb|added|
|time_commitment_in_seconds|bigint|updated|Type changed from integer to bigint

## aws_sns_subscriptions

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|delivery_policy|text|updated|Type changed from jsonb to text
|effective_delivery_policy|text|updated|Type changed from jsonb to text
|filter_policy|text|updated|Type changed from jsonb to text

## aws_sns_topics

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|delivery_policy|text|updated|Type changed from jsonb to text
|effective_delivery_policy|text|updated|Type changed from jsonb to text
|policy|text|updated|Type changed from jsonb to text

## aws_sqs_queues

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|policy|text|updated|Type changed from jsonb to text
|redrive_allow_policy|text|updated|Type changed from jsonb to text
|redrive_policy|text|updated|Type changed from jsonb to text

## aws_ssm_documents

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|account_ids|text[]|removed|
|account_sharing_info_list|jsonb|removed|
|category|text[]|added|
|category_enum|text[]|added|
|permissions|jsonb|added|

## aws_ssm_instance_compliance_items

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|account_id|text|added|
|execution_summary|jsonb|added|
|execution_summary_execution_id|text|removed|
|execution_summary_execution_time|timestamp without time zone|removed|
|execution_summary_execution_type|text|removed|
|instance_arn|text|added|
|instance_cq_id|uuid|removed|
|region|text|added|

## aws_ssm_instances

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|association_instance_status_aggregated_count|jsonb|removed|
|association_overview|jsonb|added|
|association_overview_detailed_status|text|removed|
|ip_address|text|updated|Type changed from inet to text
|source_id|text|added|
|source_type|text|added|

## aws_ssm_parameter_policies
Moved to JSON column on [aws_ssm_parameters](#aws_ssm_parameters)


## aws_ssm_parameters

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|policies|jsonb|added|

## aws_transfer_server_workflow_details_on_upload
Moved to JSON column on [aws_transfer_servers](#aws_transfer_servers)


## aws_transfer_servers

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|endpoint_details|jsonb|added|
|endpoint_details_address_allocation_ids|text[]|removed|
|endpoint_details_security_group_ids|text[]|removed|
|endpoint_details_subnet_ids|text[]|removed|
|endpoint_details_vpc_endpoint_id|text|removed|
|endpoint_details_vpc_id|text|removed|
|identity_provider_details|jsonb|added|
|identity_provider_details_directory_id|text|removed|
|identity_provider_details_function|text|removed|
|identity_provider_details_invocation_role|text|removed|
|identity_provider_details_url|text|removed|
|protocol_details|jsonb|added|
|protocol_details_as2_transports|text[]|removed|
|protocol_details_passive_ip|text|removed|
|protocol_details_set_stat_option|text|removed|
|protocol_details_tls_session_resumption_mode|text|removed|
|workflow_details|jsonb|added|

## aws_waf_rule_groups

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|id|text|removed|
|metric_name|text|removed|
|rule_group_id|text|added|

## aws_waf_rule_predicates
Moved to JSON column on [aws_waf_rules](#aws_waf_rules)


## aws_waf_rules

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|id|text|removed|
|metric_name|text|removed|
|rule_id|text|added|

## aws_waf_subscribed_rule_groups

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|

## aws_waf_web_acl_logging_configuration
Moved to JSON column on [aws_waf_web_acls](#aws_waf_web_acls)


## aws_waf_web_acl_rules
Moved to JSON column on [aws_waf_web_acls](#aws_waf_web_acls)


## aws_waf_web_acls

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|default_action_type|text|removed|
|id|text|removed|
|logging_configuration|jsonb|updated|Type changed from text[] to jsonb
|metric_name|text|removed|
|web_acl_id|text|added|

## aws_wafregional_rate_based_rule_match_predicates
Moved to JSON column on [aws_wafregional_rate_based_rules](#aws_wafregional_rate_based_rules)


## aws_wafregional_rate_based_rules

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|id|text|removed|
|match_predicates|jsonb|added|
|rule_id|text|added|

## aws_wafregional_rule_groups

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|id|text|removed|
|rule_group_id|text|added|

## aws_wafregional_rule_predicates
Moved to JSON column on [aws_wafregional_rules](#aws_wafregional_rules)


## aws_wafregional_rules

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|id|text|removed|
|predicates|jsonb|added|
|rule_id|text|added|

## aws_wafregional_web_acl_rules
Moved to JSON column on [aws_wafregional_web_acls](#aws_wafregional_web_acls)


## aws_wafregional_web_acls

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|default_action|jsonb|updated|Type changed from text to jsonb
|id|text|removed|
|rules|jsonb|added|
|web_acl_id|text|added|

## aws_wafv2_ipsets

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|addresses|inet[]|updated|Type changed from cidr[] to inet[]
|scope|text|removed|

## aws_wafv2_managed_rule_groups

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|available_labels|text[]|removed|
|capacity|bigint|removed|
|consumed_labels|text[]|removed|
|label_namespace|text|removed|
|properties|jsonb|added|
|rules|jsonb|removed|
|versioning_supported|boolean|added|

## aws_wafv2_regex_pattern_sets

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|regular_expression_list|jsonb|updated|Type changed from text[] to jsonb
|scope|text|removed|

## aws_wafv2_rule_groups

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|available_labels|jsonb|updated|Type changed from text[] to jsonb
|consumed_labels|jsonb|updated|Type changed from text[] to jsonb
|scope|text|removed|
|visibility_config|jsonb|added|
|visibility_config_cloud_watch_metrics_enabled|boolean|removed|
|visibility_config_metric_name|text|removed|
|visibility_config_sampled_requests_enabled|boolean|removed|

## aws_wafv2_web_acl_logging_configuration
Moved to JSON column on [aws_wafv2_web_acls](#aws_wafv2_web_acls)


## aws_wafv2_web_acl_post_process_firewall_manager_rule_groups
Moved to JSON column on [aws_wafv2_web_acls](#aws_wafv2_web_acls)


## aws_wafv2_web_acl_pre_process_firewall_manager_rule_groups
Moved to JSON column on [aws_wafv2_web_acls](#aws_wafv2_web_acls)


## aws_wafv2_web_acl_rules
Moved to JSON column on [aws_wafv2_web_acls](#aws_wafv2_web_acls)


## aws_wafv2_web_acls

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|captcha_config|jsonb|added|
|logging_configuration|jsonb|updated|Type changed from text[] to jsonb
|post_process_firewall_manager_rule_groups|jsonb|added|
|pre_process_firewall_manager_rule_groups|jsonb|added|
|rules|jsonb|added|
|scope|text|removed|
|visibility_config|jsonb|added|
|visibility_config_cloud_watch_metrics_enabled|boolean|removed|
|visibility_config_metric_name|text|removed|
|visibility_config_sampled_requests_enabled|boolean|removed|

## aws_workspaces_directories

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|change_compute_type|text|removed|
|custom_security_group_id|text|removed|
|default_ou|text|removed|
|device_type_android|text|removed|
|device_type_chrome_os|text|removed|
|device_type_ios|text|removed|
|device_type_linux|text|removed|
|device_type_osx|text|removed|
|device_type_web|text|removed|
|device_type_windows|text|removed|
|device_type_zero_client|text|removed|
|directory_id|text|added|
|directory_name|text|added|
|directory_type|text|added|
|enable_internet_access|boolean|removed|
|enable_maintenance_mode|boolean|removed|
|enable_work_docs|boolean|removed|
|id|text|removed|
|increase_volume_size|text|removed|
|name|text|removed|
|rebuild_workspace|text|removed|
|region|text|removed|
|restart_workspace|text|removed|
|selfservice_permissions|jsonb|added|
|switch_running_mode|text|removed|
|type|text|removed|
|user_enabled_as_local_administrator|boolean|removed|
|workspace_access_properties|jsonb|added|
|workspace_creation_properties|jsonb|added|

## aws_workspaces_workspaces

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|compute_type_name|text|removed|
|id|text|removed|
|region|text|removed|
|root_volume_size_gib|integer|removed|
|running_mode|text|removed|
|running_mode_auto_stop_timeout_in_minutes|integer|removed|
|user_volume_size_gib|integer|removed|
|workspace_id|text|added|
|workspace_properties|jsonb|added|

## aws_xray_encryption_config

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|

## aws_xray_groups

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|insights_configuration|jsonb|added|
|insights_enabled|boolean|removed|
|notifications_enabled|boolean|removed|

## aws_xray_sampling_rules

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|arn|text|removed|
|attributes|jsonb|removed|
|fixed_rate|float|removed|
|host|text|removed|
|http_method|text|removed|
|priority|bigint|removed|
|reservoir_size|bigint|removed|
|resource_arn|text|removed|
|rule_name|text|removed|
|sampling_rule|jsonb|added|
|service_name|text|removed|
|service_type|text|removed|
|url_path|text|removed|
|version|bigint|removed|

## azure_account_location_paired_region
This table was removed.


## azure_account_locations
This table was removed.


## azure_authorization_role_assignments

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|principal_id|text|removed|
|properties_principal_id|text|added|
|properties_role_definition_id|text|added|
|properties_scope|text|added|
|role_definition_id|text|removed|
|scope|text|removed|

## azure_authorization_role_definition_permissions
Moved to JSON column on [azure_authorization_role_definitions](#azure_authorization_role_definitions)


## azure_authorization_role_definitions

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|permissions|jsonb|added|

## azure_batch_account_private_endpoint_connections
Moved to JSON column on [azure_batch_accounts](#azure_batch_accounts)


## azure_batch_accounts

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|active_job_and_job_schedule_quota|bigint|updated|Type changed from integer to bigint
|auto_storage|jsonb|added|
|auto_storage_authentication_mode|text|removed|
|auto_storage_last_key_sync_time|timestamp without time zone|removed|
|auto_storage_node_identity_reference_resource_id|text|removed|
|auto_storage_storage_account_id|text|removed|
|dedicated_core_quota|bigint|updated|Type changed from integer to bigint
|encryption|jsonb|added|
|encryption_key_source|text|removed|
|encryption_key_vault_properties_key_identifier|text|removed|
|identity|jsonb|added|
|identity_principal_id|text|removed|
|identity_tenant_id|text|removed|
|identity_type|text|removed|
|identity_user_assigned_identities|jsonb|removed|
|key_vault_reference|jsonb|added|
|key_vault_reference_id|text|removed|
|key_vault_reference_url|text|removed|
|low_priority_core_quota|bigint|updated|Type changed from integer to bigint
|pool_quota|bigint|updated|Type changed from integer to bigint
|private_endpoint_connections|jsonb|added|

## azure_cdn_custom_domains
This table was newly added.

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|subscription_id|text|added|
|cdn_endpoint_id|uuid|added|
|host_name|text|added|
|resource_state|text|added|
|custom_https_provisioning_state|text|added|
|custom_https_provisioning_substate|text|added|
|validation_data|text|added|
|provisioning_state|text|added|
|id|text|added|
|name|text|added|
|type|text|added|
|system_data|jsonb|added|
|_cq_id|uuid|added|
|_cq_fetch_time|timestamp without time zone|added|

## azure_cdn_endpoints
This table was newly added.

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|subscription_id|text|added|
|cdn_profile_id|uuid|added|
|host_name|text|added|
|origins|jsonb|added|
|origin_groups|jsonb|added|
|resource_state|text|added|
|provisioning_state|text|added|
|origin_path|text|added|
|content_types_to_compress|text[]|added|
|origin_host_header|text|added|
|is_compression_enabled|boolean|added|
|is_http_allowed|boolean|added|
|is_https_allowed|boolean|added|
|query_string_caching_behavior|text|added|
|optimization_type|text|added|
|probe_path|text|added|
|geo_filters|jsonb|added|
|default_origin_group|jsonb|added|
|url_signing_keys|jsonb|added|
|delivery_policy|jsonb|added|
|web_application_firewall_policy_link|jsonb|added|
|location|text|added|
|tags|jsonb|added|
|id|text|added|
|name|text|added|
|type|text|added|
|system_data|jsonb|added|
|_cq_id|uuid|added|
|_cq_fetch_time|timestamp without time zone|added|

## azure_cdn_profile_endpoint_custom_domains
Moved to JSON column on [azure_cdn_profiles](#azure_cdn_profiles)


## azure_cdn_profile_endpoint_delivery_policy_rules
Moved to JSON column on [azure_cdn_profiles](#azure_cdn_profiles)


## azure_cdn_profile_endpoint_geo_filters
Moved to JSON column on [azure_cdn_profiles](#azure_cdn_profiles)


## azure_cdn_profile_endpoint_origin_groups
Moved to JSON column on [azure_cdn_profiles](#azure_cdn_profiles)


## azure_cdn_profile_endpoint_origins
Moved to JSON column on [azure_cdn_profiles](#azure_cdn_profiles)


## azure_cdn_profile_endpoint_routes
Moved to JSON column on [azure_cdn_profiles](#azure_cdn_profiles)


## azure_cdn_profile_endpoint_url_signing_keys
Moved to JSON column on [azure_cdn_profiles](#azure_cdn_profiles)


## azure_cdn_profile_endpoints
Moved to JSON column on [azure_cdn_profiles](#azure_cdn_profiles)


## azure_cdn_profile_rule_set_rules
Moved to JSON column on [azure_cdn_profiles](#azure_cdn_profiles)


## azure_cdn_profile_rule_sets
Moved to JSON column on [azure_cdn_profiles](#azure_cdn_profiles)


## azure_cdn_profile_security_policies
Moved to JSON column on [azure_cdn_profiles](#azure_cdn_profiles)


## azure_cdn_profiles

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|created_at_time|timestamp without time zone|removed|
|created_by|text|removed|
|created_by_type|text|removed|
|last_modified_at_time|timestamp without time zone|removed|
|last_modified_by|text|removed|
|last_modified_by_type|text|removed|
|resource_state|text|added|
|sku|jsonb|added|
|sku_name|text|removed|
|state|text|removed|
|system_data|jsonb|added|

## azure_cdn_routes
This table was newly added.

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|subscription_id|text|added|
|cdn_endpoint_id|uuid|added|
|custom_domains|jsonb|added|
|origin_group|jsonb|added|
|origin_path|text|added|
|rule_sets|jsonb|added|
|supported_protocols|text[]|added|
|patterns_to_match|text[]|added|
|query_string_caching_behavior|text|added|
|forwarding_protocol|text|added|
|link_to_default_domain|text|added|
|https_redirect|text|added|
|enabled_state|text|added|
|provisioning_state|text|added|
|deployment_status|text|added|
|id|text|added|
|name|text|added|
|type|text|added|
|system_data|jsonb|added|
|_cq_id|uuid|added|
|_cq_fetch_time|timestamp without time zone|added|

## azure_cdn_rule_sets
This table was newly added.

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|subscription_id|text|added|
|cdn_profile_id|uuid|added|
|provisioning_state|text|added|
|deployment_status|text|added|
|id|text|added|
|name|text|added|
|type|text|added|
|system_data|jsonb|added|
|_cq_id|uuid|added|
|_cq_fetch_time|timestamp without time zone|added|

## azure_cdn_rules
This table was newly added.

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|subscription_id|text|added|
|cdn_rule_set_id|uuid|added|
|order|bigint|added|
|conditions|jsonb|added|
|actions|jsonb|added|
|match_processing_behavior|text|added|
|provisioning_state|text|added|
|deployment_status|text|added|
|id|text|added|
|name|text|added|
|type|text|added|
|system_data|jsonb|added|
|_cq_id|uuid|added|
|_cq_fetch_time|timestamp without time zone|added|

## azure_cdn_security_policies
This table was newly added.

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|subscription_id|text|added|
|cdn_profile_id|uuid|added|
|provisioning_state|text|added|
|deployment_status|text|added|
|id|text|added|
|name|text|added|
|type|text|added|
|system_data|jsonb|added|
|_cq_id|uuid|added|
|_cq_fetch_time|timestamp without time zone|added|

## azure_compute_disk_encryption_settings
Moved to JSON column on [azure_compute_disks](#azure_compute_disks)


## azure_compute_disks

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|bursting_enabled|boolean|added|
|creation_data|jsonb|added|
|creation_data_create_option|text|removed|
|creation_data_gallery_image_reference_id|text|removed|
|creation_data_gallery_image_reference_lun|integer|removed|
|creation_data_image_reference_id|text|removed|
|creation_data_image_reference_lun|integer|removed|
|creation_data_source_resource_id|text|removed|
|creation_data_source_unique_id|text|removed|
|creation_data_source_uri|text|removed|
|creation_data_storage_account_id|text|removed|
|creation_data_upload_size_bytes|bigint|removed|
|disk_m_bps_read_only|bigint|added|
|disk_m_bps_read_write|bigint|added|
|disk_mbps_read_only|bigint|removed|
|disk_mbps_read_write|bigint|removed|
|disk_size_gb|bigint|updated|Type changed from integer to bigint
|encryption|jsonb|added|
|encryption_disk_encryption_set_id|text|removed|
|encryption_settings_collection|jsonb|added|
|encryption_settings_collection_enabled|boolean|removed|
|encryption_settings_collection_encryption_settings_version|text|removed|
|encryption_type|text|removed|
|extended_location|jsonb|added|
|hyper_v_generation|text|added|
|hyperv_generation|text|removed|
|max_shares|bigint|updated|Type changed from integer to bigint
|property_updates_in_progress|jsonb|added|
|purchase_plan|jsonb|added|
|security_profile|jsonb|added|
|share_info|jsonb|updated|Type changed from text[] to jsonb
|sku|jsonb|added|
|sku_name|text|removed|
|sku_tier|text|removed|
|supports_hibernation|boolean|added|
|tier|text|added|

## azure_compute_instance_views
This table was newly added.

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|subscription_id|text|added|
|compute_virtual_machine_id|uuid|added|
|platform_update_domain|bigint|added|
|platform_fault_domain|bigint|added|
|computer_name|text|added|
|os_name|text|added|
|os_version|text|added|
|hyper_v_generation|text|added|
|rdp_thumb_print|text|added|
|vm_agent|jsonb|added|
|maintenance_redeploy_status|jsonb|added|
|disks|jsonb|added|
|extensions|jsonb|added|
|vm_health|jsonb|added|
|boot_diagnostics|jsonb|added|
|assigned_host|text|added|
|statuses|jsonb|added|
|patch_status|jsonb|added|
|_cq_id|uuid|added|
|_cq_fetch_time|timestamp without time zone|added|

## azure_compute_virtual_machine_extensions
This table was newly added.

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|subscription_id|text|added|
|compute_virtual_machine_id|uuid|added|
|force_update_tag|text|added|
|publisher|text|added|
|type_handler_version|text|added|
|auto_upgrade_minor_version|boolean|added|
|enable_automatic_upgrade|boolean|added|
|provisioning_state|text|added|
|instance_view|jsonb|added|
|id|text|added|
|name|text|added|
|location|text|added|
|tags|jsonb|added|
|type|text|added|
|_cq_id|uuid|added|
|_cq_fetch_time|timestamp without time zone|added|

## azure_compute_virtual_machine_resources
Moved to JSON column on [azure_compute_virtual_machines](#azure_compute_virtual_machines)


## azure_compute_virtual_machine_scale_set_extensions
Moved to JSON column on [azure_compute_virtual_machines](#azure_compute_virtual_machines)


## azure_compute_virtual_machine_scale_set_os_profile_secrets
Moved to JSON column on [azure_compute_virtual_machines](#azure_compute_virtual_machines)


## azure_compute_virtual_machine_scale_sets

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|additional_capabilities|jsonb|added|
|additional_capabilities_ultra_ssd_enabled|boolean|removed|
|automatic_repairs_policy|jsonb|added|
|automatic_repairs_policy_enabled|boolean|removed|
|automatic_repairs_policy_grace_period|text|removed|
|billing_profile_max_price|float|removed|
|diagnostics_profile|jsonb|removed|
|eviction_policy|text|removed|
|extended_location|jsonb|added|
|extended_location_name|text|removed|
|extended_location_type|text|removed|
|extension_profile_extensions_time_budget|text|removed|
|host_group|jsonb|added|
|host_group_id|text|removed|
|identity|jsonb|added|
|identity_principal_id|text|removed|
|identity_tenant_id|text|removed|
|identity_type|text|removed|
|identity_user_assigned_identities|jsonb|removed|
|license_type|text|removed|
|network_profile|jsonb|removed|
|os_profile_admin_password|text|removed|
|os_profile_admin_username|text|removed|
|os_profile_computer_name_prefix|text|removed|
|os_profile_custom_data|text|removed|
|os_profile_linux_configuration|jsonb|removed|
|os_profile_windows_configuration|jsonb|removed|
|plan|jsonb|added|
|plan_name|text|removed|
|plan_product|text|removed|
|plan_promotion_code|text|removed|
|plan_publisher|text|removed|
|platform_fault_domain_count|bigint|updated|Type changed from integer to bigint
|priority|text|removed|
|proximity_placement_group|jsonb|added|
|proximity_placement_group_id|text|removed|
|scale_in_policy|jsonb|added|
|scale_in_policy_rules|text[]|removed|
|scheduled_events_profile|jsonb|removed|
|security_profile|jsonb|removed|
|sku|jsonb|added|
|sku_capacity|bigint|removed|
|sku_name|text|removed|
|sku_tier|text|removed|
|storage_profile|jsonb|removed|
|user_data|text|removed|
|virtual_machine_profile|jsonb|added|

## azure_compute_virtual_machine_secret_vault_certificates
Moved to JSON column on [azure_compute_virtual_machines](#azure_compute_virtual_machines)


## azure_compute_virtual_machine_secrets
Moved to JSON column on [azure_compute_virtual_machines](#azure_compute_virtual_machines)


## azure_compute_virtual_machine_win_config_rm_listeners
Moved to JSON column on [azure_compute_virtual_machines](#azure_compute_virtual_machines)


## azure_compute_virtual_machines

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|additional_capabilities|jsonb|added|
|additional_capabilities_ultra_ssd_enabled|boolean|removed|
|admin_password|text|removed|
|admin_username|text|removed|
|allow_extension_operations|boolean|removed|
|availability_set|jsonb|added|
|availability_set_id|text|removed|
|billing_profile|jsonb|added|
|billing_profile_max_price|float|removed|
|computer_name|text|removed|
|custom_data|text|removed|
|diagnostics_profile|jsonb|added|
|diagnostics_profile_boot_diagnostics_enabled|boolean|removed|
|diagnostics_profile_boot_diagnostics_storage_uri|text|removed|
|extended_location|jsonb|added|
|extended_location_name|text|removed|
|extended_location_type|text|removed|
|hardware_profile|jsonb|added|
|hardware_profile_vm_size|text|removed|
|host|jsonb|added|
|host_group|jsonb|added|
|host_group_id|text|removed|
|host_id|text|removed|
|identity|jsonb|added|
|identity_principal_id|text|removed|
|identity_tenant_id|text|removed|
|identity_type|text|removed|
|identity_user_assigned_identities|jsonb|removed|
|linux_configuration_disable_password_authentication|boolean|removed|
|linux_configuration_patch_settings_assessment_mode|text|removed|
|linux_configuration_patch_settings_patch_mode|text|removed|
|linux_configuration_provision_vm_agent|boolean|removed|
|linux_configuration_ssh_public_keys|jsonb|removed|
|network_profile|jsonb|added|
|network_profile_network_api_version|text|removed|
|network_profile_network_interface_configurations|jsonb|removed|
|network_profile_network_interfaces|jsonb|removed|
|os_profile|jsonb|added|
|plan|jsonb|added|
|plan_name|text|removed|
|plan_product|text|removed|
|plan_promotion_code|text|removed|
|plan_publisher|text|removed|
|platform_fault_domain|bigint|updated|Type changed from integer to bigint
|proximity_placement_group|jsonb|added|
|proximity_placement_group_id|text|removed|
|require_guest_provision_signal|boolean|removed|
|resources|jsonb|added|
|security_profile|jsonb|added|
|security_profile_encryption_at_host|boolean|removed|
|security_profile_security_type|text|removed|
|security_profile_uefi_settings_secure_boot_enabled|boolean|removed|
|security_profile_uefi_settings_v_tpm_enabled|boolean|removed|
|virtual_machine_scale_set|jsonb|added|
|virtual_machine_scale_set_id|text|removed|
|windows_configuration_additional_unattend_content|jsonb|removed|
|windows_configuration_enable_automatic_updates|boolean|removed|
|windows_configuration_patch_settings_assessment_mode|text|removed|
|windows_configuration_patch_settings_enable_hotpatching|boolean|removed|
|windows_configuration_patch_settings_patch_mode|text|removed|
|windows_configuration_provision_vm_agent|boolean|removed|
|windows_configuration_time_zone|text|removed|

## azure_container_managed_cluster_agent_pool_profiles
Moved to JSON column on [azure_container_managed_clusters](#azure_container_managed_clusters)


## azure_container_managed_cluster_pip_user_assigned_id_exceptions
Moved to JSON column on [azure_container_managed_clusters](#azure_container_managed_clusters)


## azure_container_managed_cluster_pip_user_assigned_identities
Moved to JSON column on [azure_container_managed_clusters](#azure_container_managed_clusters)


## azure_container_managed_cluster_private_link_resources
Moved to JSON column on [azure_container_managed_clusters](#azure_container_managed_clusters)


## azure_container_managed_clusters

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|aad_profile|jsonb|added|
|aad_profile_admin_group_object_ids|text[]|removed|
|aad_profile_client_app_id|text|removed|
|aad_profile_enable_azure_rbac|boolean|removed|
|aad_profile_managed|boolean|removed|
|aad_profile_server_app_id|text|removed|
|aad_profile_server_app_secret|text|removed|
|aad_profile_tenant_id|text|removed|
|agent_pool_profiles|jsonb|added|
|api_server_access_profile|jsonb|added|
|api_server_access_profile_authorized_ip_ranges|text[]|removed|
|api_server_access_profile_enable_private_cluster|boolean|removed|
|api_server_access_profile_private_dns_zone|text|removed|
|auto_scaler_profile|jsonb|added|
|auto_scaler_profile_expander|text|removed|
|auto_upgrade_profile|jsonb|added|
|auto_upgrade_profile_upgrade_channel|text|removed|
|enable_pod_security_policy|boolean|added|
|extended_location|jsonb|added|
|extended_location_name|text|removed|
|extended_location_type|text|removed|
|http_proxy_config|jsonb|added|
|http_proxy_config_http_proxy|text|removed|
|http_proxy_config_https_proxy|text|removed|
|http_proxy_config_no_proxy|text[]|removed|
|http_proxy_config_trusted_ca|text|removed|
|identity|jsonb|added|
|identity_principal_id|text|removed|
|identity_tenant_id|text|removed|
|identity_type|text|removed|
|identity_user_assigned_identities|jsonb|removed|
|linux_profile|jsonb|added|
|linux_profile_admin_username|text|removed|
|max_agent_pools|bigint|updated|Type changed from integer to bigint
|network_profile|jsonb|added|
|network_profile_dns_service_ip|text|removed|
|network_profile_docker_bridge_cidr|text|removed|
|network_profile_load_balancer_allocated_outbound_ports|integer|removed|
|network_profile_load_balancer_effective_outbound_ips|text[]|removed|
|network_profile_load_balancer_idle_timeout|integer|removed|
|network_profile_load_balancer_managed_outbound_ips_count|integer|removed|
|network_profile_load_balancer_outbound_ip_prefixes|text[]|removed|
|network_profile_load_balancer_outbound_ips|text[]|removed|
|network_profile_load_balancer_sku|text|removed|
|network_profile_network_mode|text|removed|
|network_profile_network_plugin|text|removed|
|network_profile_network_policy|text|removed|
|network_profile_outbound_type|text|removed|
|network_profile_pod_cidr|text|removed|
|network_profile_service_cidr|text|removed|
|pod_identity_profile|jsonb|added|
|pod_identity_profile_allow_network_plugin_kubenet|boolean|removed|
|pod_identity_profile_enabled|boolean|removed|
|power_state|jsonb|added|
|power_state_code|text|removed|
|private_link_resources|jsonb|added|
|service_principal_profile|jsonb|added|
|service_principal_profile_client_id|text|removed|
|service_principal_profile_secret|text|removed|
|sku|jsonb|added|
|sku_name|text|removed|
|sku_tier|text|removed|
|windows_profile|jsonb|added|
|windows_profile_admin_password|text|removed|
|windows_profile_admin_username|text|removed|
|windows_profile_enable_csi_proxy|boolean|removed|
|windows_profile_license_type|text|removed|

## azure_container_registries

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|network_rule_set|jsonb|added|
|network_rule_set_default_action|text|removed|
|policies|jsonb|added|
|quarantine_policy_status|text|removed|
|retention_policy_days|integer|removed|
|retention_policy_last_updated_time|timestamp without time zone|removed|
|retention_policy_status|text|removed|
|sku|jsonb|added|
|sku_name|text|removed|
|sku_tier|text|removed|
|status|jsonb|updated|Type changed from text to jsonb
|status_message|text|removed|
|status_timestamp|timestamp without time zone|removed|
|storage_account|jsonb|added|
|storage_account_id|text|removed|
|trust_policy_status|text|removed|
|trust_policy_type|text|removed|

## azure_container_registry_network_rule_set_ip_rules
Moved to JSON column on [azure_container_registries](#azure_container_registries)


## azure_container_registry_network_rule_set_virtual_network_rules
Moved to JSON column on [azure_container_registries](#azure_container_registries)


## azure_container_registry_replications
Moved to JSON column on [azure_container_registries](#azure_container_registries)


## azure_container_replications
This table was newly added.

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|subscription_id|text|added|
|container_registry_id|uuid|added|
|provisioning_state|text|added|
|status|jsonb|added|
|id|text|added|
|name|text|added|
|type|text|added|
|location|text|added|
|tags|jsonb|added|
|_cq_id|uuid|added|
|_cq_fetch_time|timestamp without time zone|added|

## azure_cosmosdb_account_cors
Moved to JSON column on [azure_cosmosdb_accounts](#azure_cosmosdb_accounts)


## azure_cosmosdb_account_failover_policies
Moved to JSON column on [azure_cosmosdb_accounts](#azure_cosmosdb_accounts)


## azure_cosmosdb_account_locations
Moved to JSON column on [azure_cosmosdb_accounts](#azure_cosmosdb_accounts)


## azure_cosmosdb_account_private_endpoint_connections
Moved to JSON column on [azure_cosmosdb_accounts](#azure_cosmosdb_accounts)


## azure_cosmosdb_account_read_locations
Moved to JSON column on [azure_cosmosdb_accounts](#azure_cosmosdb_accounts)


## azure_cosmosdb_account_write_locations
Moved to JSON column on [azure_cosmosdb_accounts](#azure_cosmosdb_accounts)


## azure_cosmosdb_accounts

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|api_properties|jsonb|added|
|api_properties_server_version|text|removed|
|capabilities|jsonb|updated|Type changed from text[] to jsonb
|consistency_policy|jsonb|added|
|consistency_policy_default_consistency_level|text|removed|
|consistency_policy_max_interval_in_seconds|integer|removed|
|consistency_policy_max_staleness_prefix|bigint|removed|
|cors|jsonb|added|
|failover_policies|jsonb|added|
|ip_rules|jsonb|updated|Type changed from text[] to jsonb
|kind|text|added|
|locations|jsonb|added|
|private_endpoint_connections|jsonb|added|
|read_locations|jsonb|added|
|write_locations|jsonb|added|

## azure_cosmosdb_mongo_db_databases
Renamed from [azure_cosmosdb_mongodb_databases](azure_cosmosdb_mongodb_databases)


## azure_cosmosdb_mongodb_databases
Renamed to [azure_cosmosdb_mongo_db_databases](#azure_cosmosdb_mongo_db_databases)


## azure_cosmosdb_sql_databases

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|autoscale_settings_max_throughput|integer|removed|
|cosmosdb_account_id|uuid|added|
|database_colls|text|removed|
|database_etag|text|removed|
|database_id|text|removed|
|database_rid|text|removed|
|database_ts|float|removed|
|database_users|text|removed|
|options|jsonb|added|
|resource|jsonb|added|
|sql_database_get_properties_throughput|integer|removed|

## azure_datalake_analytics_account_compute_policies
Moved to JSON column on [azure_datalake_analytics_accounts](#azure_datalake_analytics_accounts)


## azure_datalake_analytics_account_data_lake_store_accounts
Moved to JSON column on [azure_datalake_analytics_accounts](#azure_datalake_analytics_accounts)


## azure_datalake_analytics_account_firewall_rules
Moved to JSON column on [azure_datalake_analytics_accounts](#azure_datalake_analytics_accounts)


## azure_datalake_analytics_account_storage_accounts
Moved to JSON column on [azure_datalake_analytics_accounts](#azure_datalake_analytics_accounts)


## azure_datalake_analytics_accounts

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|compute_policies|jsonb|added|
|data_lake_store_accounts|jsonb|added|
|firewall_rules|jsonb|added|
|max_degree_of_parallelism|bigint|updated|Type changed from integer to bigint
|max_degree_of_parallelism_per_job|bigint|updated|Type changed from integer to bigint
|max_job_count|bigint|updated|Type changed from integer to bigint
|min_priority_per_job|bigint|updated|Type changed from integer to bigint
|query_store_retention|bigint|updated|Type changed from integer to bigint
|storage_accounts|jsonb|added|
|system_max_degree_of_parallelism|bigint|updated|Type changed from integer to bigint
|system_max_job_count|bigint|updated|Type changed from integer to bigint

## azure_datalake_storage_account_firewall_rules
This table was removed.


## azure_datalake_storage_account_trusted_id_providers
This table was removed.


## azure_datalake_storage_account_virtual_network_rules
This table was removed.


## azure_datalake_storage_accounts
This table was removed.


## azure_datalake_store_accounts
This table was newly added.

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|subscription_id|text|added|
|identity|jsonb|added|
|default_group|text|added|
|encryption_config|jsonb|added|
|encryption_state|text|added|
|encryption_provisioning_state|text|added|
|firewall_rules|jsonb|added|
|virtual_network_rules|jsonb|added|
|firewall_state|text|added|
|firewall_allow_azure_ips|text|added|
|trusted_id_providers|jsonb|added|
|trusted_id_provider_state|text|added|
|new_tier|text|added|
|current_tier|text|added|
|account_id|uuid|added|
|provisioning_state|text|added|
|state|text|added|
|creation_time|timestamp without time zone|added|
|last_modified_time|timestamp without time zone|added|
|endpoint|text|added|
|id|text|added|
|name|text|added|
|type|text|added|
|location|text|added|
|tags|jsonb|added|
|_cq_id|uuid|added|
|_cq_fetch_time|timestamp without time zone|added|

## azure_eventhub_namespace_encryption_key_vault_properties
Moved to JSON column on [azure_eventhub_namespaces](#azure_eventhub_namespaces)


## azure_eventhub_namespaces

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|created_at|timestamp without time zone|added|
|created_at_time|timestamp without time zone|removed|
|encryption|jsonb|added|
|encryption_key_source|text|removed|
|identity|jsonb|added|
|identity_principal_id|text|removed|
|identity_tenant_id|text|removed|
|identity_type|text|removed|
|maximum_throughput_units|bigint|updated|Type changed from integer to bigint
|network_rule_set|jsonb|removed|
|sku|jsonb|added|
|sku_capacity|integer|removed|
|sku_name|text|removed|
|sku_tier|text|removed|
|updated_at|timestamp without time zone|added|
|updated_at_time|timestamp without time zone|removed|

## azure_eventhub_network_rule_sets
This table was newly added.

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|subscription_id|text|added|
|eventhub_namespace_id|uuid|added|
|trusted_service_access_enabled|boolean|added|
|default_action|text|added|
|virtual_network_rules|jsonb|added|
|ip_rules|jsonb|added|
|id|text|added|
|name|text|added|
|type|text|added|
|_cq_id|uuid|added|
|_cq_fetch_time|timestamp without time zone|added|

## azure_front_door_backend_pool_backends
This table was removed.


## azure_front_door_backend_pools
This table was removed.


## azure_front_door_frontend_endpoints
This table was removed.


## azure_front_door_health_probe_settings
This table was removed.


## azure_front_door_load_balancing_settings
This table was removed.


## azure_front_door_routing_rules
This table was removed.


## azure_front_door_rules_engine_rule_match_conditions
This table was removed.


## azure_front_door_rules_engine_rules
This table was removed.


## azure_front_door_rules_engines
This table was removed.


## azure_front_doors
This table was removed.


## azure_frontdoor_doors
This table was newly added.

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|subscription_id|text|added|
|resource_state|text|added|
|provisioning_state|text|added|
|cname|text|added|
|frontdoor_id|text|added|
|rules_engines|jsonb|added|
|friendly_name|text|added|
|routing_rules|jsonb|added|
|load_balancing_settings|jsonb|added|
|health_probe_settings|jsonb|added|
|backend_pools|jsonb|added|
|frontend_endpoints|jsonb|added|
|backend_pools_settings|jsonb|added|
|enabled_state|text|added|
|id|text|added|
|name|text|added|
|type|text|added|
|location|text|added|
|tags|jsonb|added|
|_cq_id|uuid|added|
|_cq_fetch_time|timestamp without time zone|added|

## azure_iothub_devices
This table was newly added.

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|subscription_id|text|added|
|etag|text|added|
|properties_authorization_policies|jsonb|added|
|properties_disable_local_auth|boolean|added|
|properties_disable_device_sas|boolean|added|
|properties_disable_module_sas|boolean|added|
|properties_restrict_outbound_network_access|boolean|added|
|properties_allowed_fqdn_list|text[]|added|
|properties_public_network_access|text|added|
|properties_ip_filter_rules|jsonb|added|
|properties_network_rule_sets|jsonb|added|
|properties_min_tls_version|text|added|
|properties_private_endpoint_connections|jsonb|added|
|properties_provisioning_state|text|added|
|properties_state|text|added|
|properties_host_name|text|added|
|properties_event_hub_endpoints|jsonb|added|
|properties_routing|jsonb|added|
|properties_storage_endpoints|jsonb|added|
|properties_messaging_endpoints|jsonb|added|
|properties_enable_file_upload_notifications|boolean|added|
|properties_cloud_to_device|jsonb|added|
|properties_comments|text|added|
|properties_features|text|added|
|properties_locations|jsonb|added|
|properties_enable_data_residency|boolean|added|
|sku|jsonb|added|
|identity|jsonb|added|
|system_data|jsonb|added|
|id|text|added|
|name|text|added|
|type|text|added|
|location|text|added|
|tags|jsonb|added|
|_cq_id|uuid|added|
|_cq_fetch_time|timestamp without time zone|added|

## azure_iothub_hub_authorization_policies
This table was removed.


## azure_iothub_hub_ip_filter_rules
This table was removed.


## azure_iothub_hub_network_rule_sets_ip_rules
This table was removed.


## azure_iothub_hub_private_endpoint_connections
This table was removed.


## azure_iothub_hub_routing_endpoints_event_hubs
This table was removed.


## azure_iothub_hub_routing_endpoints_service_bus_queues
This table was removed.


## azure_iothub_hub_routing_endpoints_service_bus_topics
This table was removed.


## azure_iothub_hub_routing_endpoints_storage_containers
This table was removed.


## azure_iothub_hub_routing_routes
This table was removed.


## azure_iothub_hubs
This table was removed.


## azure_keyvault_keys
This table was newly added.

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|subscription_id|text|added|
|keyvault_vault_id|uuid|added|
|kid|text|added|
|attributes|jsonb|added|
|tags|jsonb|added|
|managed|boolean|added|
|_cq_id|uuid|added|
|_cq_fetch_time|timestamp without time zone|added|

## azure_keyvault_managed_hsm
Moved to JSON column on [azure_keyvault_managed_hsms](#azure_keyvault_managed_hsms)


## azure_keyvault_managed_hsms
This table was newly added.

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|subscription_id|text|added|
|properties_tenant_id|uuid|added|
|properties_initial_admin_object_ids|text[]|added|
|properties_hsm_uri|text|added|
|properties_enable_soft_delete|boolean|added|
|properties_soft_delete_retention_in_days|bigint|added|
|properties_enable_purge_protection|boolean|added|
|properties_create_mode|text|added|
|properties_status_message|text|added|
|properties_provisioning_state|text|added|
|id|text|added|
|name|text|added|
|type|text|added|
|location|text|added|
|sku|jsonb|added|
|tags|jsonb|added|
|_cq_id|uuid|added|
|_cq_fetch_time|timestamp without time zone|added|

## azure_keyvault_secrets
This table was newly added.

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|subscription_id|text|added|
|keyvault_vault_id|uuid|added|
|id|text|added|
|attributes|jsonb|added|
|tags|jsonb|added|
|content_type|text|added|
|managed|boolean|added|
|_cq_id|uuid|added|
|_cq_fetch_time|timestamp without time zone|added|

## azure_keyvault_vault_access_policies
Moved to JSON column on [azure_keyvault_vaults](#azure_keyvault_vaults)


## azure_keyvault_vault_keys
Moved to JSON column on [azure_keyvault_vaults](#azure_keyvault_vaults)


## azure_keyvault_vault_private_endpoint_connections
Moved to JSON column on [azure_keyvault_vaults](#azure_keyvault_vaults)


## azure_keyvault_vault_secrets
Moved to JSON column on [azure_keyvault_vaults](#azure_keyvault_vaults)


## azure_keyvault_vaults

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|create_mode|text|removed|
|enable_purge_protection|boolean|removed|
|enable_rbac_authorization|boolean|removed|
|enable_soft_delete|boolean|removed|
|enabled_for_deployment|boolean|removed|
|enabled_for_disk_encryption|boolean|removed|
|enabled_for_template_deployment|boolean|removed|
|network_acls_bypass|text|removed|
|network_acls_default_action|text|removed|
|network_acls_ip_rules|text[]|removed|
|network_acls_virtual_network_rules|text[]|removed|
|properties_access_policies|jsonb|added|
|properties_create_mode|text|added|
|properties_enable_purge_protection|boolean|added|
|properties_enable_rbac_authorization|boolean|added|
|properties_enable_soft_delete|boolean|added|
|properties_enabled_for_deployment|boolean|added|
|properties_enabled_for_disk_encryption|boolean|added|
|properties_enabled_for_template_deployment|boolean|added|
|properties_network_acls|jsonb|added|
|properties_private_endpoint_connections|jsonb|added|
|properties_sku|jsonb|added|
|properties_soft_delete_retention_in_days|bigint|added|
|properties_tenant_id|uuid|added|
|properties_vault_uri|text|added|
|sku_family|text|removed|
|sku_name|text|removed|
|soft_delete_retention_in_days|integer|removed|
|tenant_id|uuid|removed|
|vault_uri|text|removed|

## azure_logic_app_workflows
This table was removed.


## azure_logic_diagnostic_settings
This table was newly added.

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|subscription_id|text|added|
|logic_workflow_id|uuid|added|
|storage_account_id|text|added|
|service_bus_rule_id|text|added|
|event_hub_authorization_rule_id|text|added|
|event_hub_name|text|added|
|metrics|jsonb|added|
|logs|jsonb|added|
|workspace_id|text|added|
|log_analytics_destination_type|text|added|
|id|text|added|
|name|text|added|
|type|text|added|
|_cq_id|uuid|added|
|_cq_fetch_time|timestamp without time zone|added|

## azure_logic_workflows
This table was newly added.

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|subscription_id|text|added|
|provisioning_state|text|added|
|created_time|timestamp without time zone|added|
|changed_time|timestamp without time zone|added|
|state|text|added|
|version|text|added|
|access_endpoint|text|added|
|endpoints_configuration|jsonb|added|
|access_control|jsonb|added|
|sku|jsonb|added|
|integration_account|jsonb|added|
|integration_service_environment|jsonb|added|
|parameters|jsonb|added|
|identity|jsonb|added|
|id|text|added|
|name|text|added|
|type|text|added|
|location|text|added|
|tags|jsonb|added|
|_cq_id|uuid|added|
|_cq_fetch_time|timestamp without time zone|added|

## azure_mariadb_configurations
This table was newly added.

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|subscription_id|text|added|
|mariadb_server_id|uuid|added|
|value|text|added|
|description|text|added|
|default_value|text|added|
|data_type|text|added|
|allowed_values|text|added|
|source|text|added|
|id|text|added|
|name|text|added|
|type|text|added|
|_cq_id|uuid|added|
|_cq_fetch_time|timestamp without time zone|added|

## azure_mariadb_server_configurations
Moved to JSON column on [azure_mariadb_servers](#azure_mariadb_servers)


## azure_mariadb_server_private_endpoint_connections
Moved to JSON column on [azure_mariadb_servers](#azure_mariadb_servers)


## azure_mariadb_servers

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|backup_retention_days|integer|removed|
|earliest_restore_date|timestamp without time zone|added|
|earliest_restore_date_time|timestamp without time zone|removed|
|geo_redundant_backup|text|removed|
|private_endpoint_connections|jsonb|added|
|replica_capacity|bigint|updated|Type changed from integer to bigint
|sku|jsonb|added|
|sku_capacity|integer|removed|
|sku_family|text|removed|
|sku_name|text|removed|
|sku_size|text|removed|
|sku_tier|text|removed|
|storage_autogrow|text|removed|
|storage_mb|integer|removed|
|storage_profile|jsonb|added|

## azure_monitor_activity_log_alert_action_groups
Moved to JSON column on [azure_monitor_activity_logs](#azure_monitor_activity_logs)


## azure_monitor_activity_log_alert_conditions
Moved to JSON column on [azure_monitor_activity_logs](#azure_monitor_activity_logs)


## azure_monitor_activity_log_alerts

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|actions|jsonb|added|
|condition|jsonb|added|

## azure_monitor_activity_logs

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|authorization|jsonb|added|
|authorization_action|text|removed|
|authorization_role|text|removed|
|authorization_scope|text|removed|
|category|jsonb|added|
|category_localized_value|text|removed|
|category_value|text|removed|
|event_name|jsonb|added|
|event_name_localized_value|text|removed|
|event_name_value|text|removed|
|event_timestamp|timestamp without time zone|added|
|event_timestamp_time|timestamp without time zone|removed|
|http_request|jsonb|added|
|http_request_client_ip_address|text|removed|
|http_request_client_request_id|text|removed|
|http_request_method|text|removed|
|http_request_uri|text|removed|
|operation_name|jsonb|added|
|operation_name_localized_value|text|removed|
|operation_name_value|text|removed|
|resource_provider_name|jsonb|added|
|resource_provider_name_localized_value|text|removed|
|resource_provider_name_value|text|removed|
|resource_type|jsonb|added|
|resource_type_localized_value|text|removed|
|resource_type_value|text|removed|
|status|jsonb|added|
|status_localized_value|text|removed|
|status_value|text|removed|
|sub_status|jsonb|added|
|sub_status_localized_value|text|removed|
|sub_status_value|text|removed|
|submission_timestamp|timestamp without time zone|added|
|submission_timestamp_time|timestamp without time zone|removed|

## azure_monitor_diagnostic_setting_logs
Moved to JSON column on [azure_monitor_diagnostic_settings](#azure_monitor_diagnostic_settings)


## azure_monitor_diagnostic_setting_metrics
Moved to JSON column on [azure_monitor_diagnostic_settings](#azure_monitor_diagnostic_settings)


## azure_monitor_diagnostic_settings

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|logs|jsonb|added|
|metrics|jsonb|added|
|monitor_resource_id|uuid|added|

## azure_monitor_log_profiles

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|etag|text|added|
|kind|text|added|
|retention_policy|jsonb|added|
|retention_policy_days|integer|removed|
|retention_policy_enabled|boolean|removed|

## azure_monitor_resources
This table was newly added.

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|id|text|added|
|_cq_id|uuid|added|
|_cq_fetch_time|timestamp without time zone|added|

## azure_mysql_configurations
This table was newly added.

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|subscription_id|text|added|
|mysql_server_id|uuid|added|
|value|text|added|
|description|text|added|
|default_value|text|added|
|data_type|text|added|
|allowed_values|text|added|
|source|text|added|
|id|text|added|
|name|text|added|
|type|text|added|
|_cq_id|uuid|added|
|_cq_fetch_time|timestamp without time zone|added|

## azure_mysql_server_configurations
Moved to JSON column on [azure_mysql_servers](#azure_mysql_servers)


## azure_mysql_server_private_endpoint_connections
Moved to JSON column on [azure_mysql_servers](#azure_mysql_servers)


## azure_mysql_servers

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|earliest_restore_date|timestamp without time zone|added|
|earliest_restore_date_time|timestamp without time zone|removed|
|identity|jsonb|added|
|identity_principal_id|uuid|removed|
|identity_tenant_id|uuid|removed|
|identity_type|text|removed|
|private_endpoint_connections|jsonb|added|
|replica_capacity|bigint|updated|Type changed from integer to bigint
|sku|jsonb|added|
|sku_capacity|integer|removed|
|sku_family|text|removed|
|sku_name|text|removed|
|sku_size|text|removed|
|sku_tier|text|removed|
|storage_profile|jsonb|added|
|storage_profile_backup_retention_days|integer|removed|
|storage_profile_geo_redundant_backup|text|removed|
|storage_profile_storage_autogrow|text|removed|
|storage_profile_storage_mb|integer|removed|

## azure_network_express_route_circuit_authorizations
Moved to JSON column on [azure_network_express_route_circuits](#azure_network_express_route_circuits)


## azure_network_express_route_circuit_connections
Moved to JSON column on [azure_network_express_route_circuits](#azure_network_express_route_circuits)


## azure_network_express_route_circuit_peerings
Moved to JSON column on [azure_network_express_route_circuits](#azure_network_express_route_circuits)


## azure_network_express_route_circuits

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|authorizations|jsonb|added|
|bandwidth_in_gbps|real|updated|Type changed from float to real
|express_route_port|jsonb|added|
|express_route_port_id|text|removed|
|peerings|jsonb|added|
|service_provider_properties|jsonb|added|
|service_provider_properties_bandwidth_in_mbps|integer|removed|
|service_provider_properties_peering_location|text|removed|
|service_provider_properties_service_provider_name|text|removed|
|sku|jsonb|added|
|sku_family|text|removed|
|sku_name|text|removed|
|sku_tier|text|removed|
|stag|bigint|updated|Type changed from integer to bigint

## azure_network_express_route_connections
This table was removed.


## azure_network_express_route_gateways

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|auto_scale_configuration|jsonb|added|
|auto_scale_configuration_bound_max|integer|removed|
|auto_scale_configuration_bound_min|integer|removed|
|express_route_connections|jsonb|added|
|virtual_hub|jsonb|added|
|virtual_hub_id|text|removed|

## azure_network_express_route_links
This table was removed.


## azure_network_express_route_ports

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|bandwidth_in_gbps|bigint|updated|Type changed from integer to bigint
|circuits|jsonb|updated|Type changed from text[] to jsonb
|identity|jsonb|added|
|identity_principal_id|text|removed|
|identity_tenant_id|text|removed|
|identity_type|text|removed|
|identity_user_assigned_identities|jsonb|removed|
|links|jsonb|added|
|provisioned_bandwidth_in_gbps|real|updated|Type changed from float to real

## azure_network_flow_logs
This table was newly added.

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|subscription_id|text|added|
|network_watcher_id|uuid|added|
|target_resource_id|text|added|
|target_resource_guid|text|added|
|storage_id|text|added|
|enabled|boolean|added|
|retention_policy|jsonb|added|
|format|jsonb|added|
|flow_analytics_configuration|jsonb|added|
|provisioning_state|text|added|
|etag|text|added|
|id|text|added|
|name|text|added|
|type|text|added|
|location|text|added|
|tags|jsonb|added|
|_cq_id|uuid|added|
|_cq_fetch_time|timestamp without time zone|added|

## azure_network_interface_ip_configurations
Moved to JSON column on [azure_network_interfaces](#azure_network_interfaces)


## azure_network_interfaces

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|dns_settings|jsonb|added|
|dns_settings_applied_dns_servers|text[]|removed|
|dns_settings_dns_servers|text[]|removed|
|dns_settings_internal_dns_name_label|text|removed|
|dns_settings_internal_domain_name_suffix|text|removed|
|dns_settings_internal_fqdn|text|removed|
|dscp_configuration|jsonb|added|
|dscp_configuration_id|text|removed|
|extended_location|jsonb|added|
|extended_location_name|text|removed|
|extended_location_type|text|removed|
|ip_configurations|jsonb|added|
|network_security_group|jsonb|updated|Type changed from text to jsonb
|private_endpoint|jsonb|updated|Type changed from text to jsonb
|virtual_machine|jsonb|added|
|virtual_machine_id|text|removed|

## azure_network_peer_express_route_circuit_connections
This table was removed.


## azure_network_public_ip_addresses

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|ddos_settings|jsonb|added|
|ddos_settings_ddos_custom_policy_id|text|removed|
|ddos_settings_protected_ip|boolean|removed|
|ddos_settings_protection_coverage|text|removed|
|dns_settings|jsonb|added|
|dns_settings_domain_name_label|text|removed|
|dns_settings_fqdn|text|removed|
|dns_settings_reverse_fqdn|text|removed|
|extended_location|jsonb|added|
|extended_location_name|text|removed|
|extended_location_type|text|removed|
|idle_timeout_in_minutes|bigint|updated|Type changed from integer to bigint
|ip_address|text|updated|Type changed from inet to text
|public_ip_prefix|jsonb|added|
|public_ip_prefix_id|text|removed|
|sku|jsonb|added|
|sku_name|text|removed|
|sku_tier|text|removed|

## azure_network_route_filter_rules
Moved to JSON column on [azure_network_route_filters](#azure_network_route_filters)


## azure_network_route_filters

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|ipv6_peerings|jsonb|removed|
|ipv_6_peerings|jsonb|added|
|rules|jsonb|added|

## azure_network_route_table_routes
Moved to JSON column on [azure_network_route_tables](#azure_network_route_tables)


## azure_network_route_tables

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|route_table_subnets|text[]|removed|
|routes|jsonb|added|
|subnets|jsonb|added|

## azure_network_security_group_default_security_rules
Moved to JSON column on [azure_network_security_groups](#azure_network_security_groups)


## azure_network_security_group_flow_logs
Moved to JSON column on [azure_network_security_groups](#azure_network_security_groups)


## azure_network_security_group_security_rules
Moved to JSON column on [azure_network_security_groups](#azure_network_security_groups)


## azure_network_security_groups

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|default_security_rules|jsonb|added|
|flow_logs|jsonb|added|
|network_interfaces|jsonb|added|
|security_rules|jsonb|added|
|subnets|jsonb|added|

## azure_network_virtual_network_gateway_connections

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|network_virtual_network_gateway_id|uuid|added|
|peer|jsonb|added|
|peer_id|text|removed|
|routing_weight|bigint|updated|Type changed from integer to bigint
|subscription_id|text|added|
|virtual_network_gateway_cq_id|uuid|removed|

## azure_network_virtual_network_gateways

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|bgp_settings|jsonb|added|
|bgp_settings_asn|bigint|removed|
|bgp_settings_bgp_peering_address|text|removed|
|bgp_settings_bgp_peering_addresses|jsonb|removed|
|bgp_settings_peer_weight|integer|removed|
|custom_routes|jsonb|added|
|custom_routes_address_prefixes|text[]|removed|
|extended_location|jsonb|added|
|extended_location_name|text|removed|
|extended_location_type|text|removed|
|gateway_default_site|jsonb|added|
|gateway_default_site_id|text|removed|
|network_virtual_network_id|uuid|added|
|sku|jsonb|added|
|sku_capacity|integer|removed|
|sku_name|text|removed|
|sku_tier|text|removed|
|subscription_id|text|added|
|v_net_extended_location_resource_id|text|added|
|virtual_network_cq_id|uuid|removed|
|vnet_extended_location_resource_id|text|removed|
|vpn_client_configuration|jsonb|added|
|vpn_client_configuration_aad_audience|text|removed|
|vpn_client_configuration_aad_issuer|text|removed|
|vpn_client_configuration_aad_tenant|text|removed|
|vpn_client_configuration_address_pool|text[]|removed|
|vpn_client_configuration_authentication_types|text[]|removed|
|vpn_client_configuration_ipsec_policies|jsonb|removed|
|vpn_client_configuration_protocols|text[]|removed|
|vpn_client_configuration_radius_server_address|text|removed|
|vpn_client_configuration_radius_server_secret|text|removed|
|vpn_client_configuration_radius_servers|jsonb|removed|
|vpn_client_configuration_revoked_certificates|jsonb|removed|
|vpn_client_configuration_root_certificates|jsonb|removed|

## azure_network_virtual_network_peerings
Moved to JSON column on [azure_network_virtual_networks](#azure_network_virtual_networks)


## azure_network_virtual_network_subnets
Moved to JSON column on [azure_network_virtual_networks](#azure_network_virtual_networks)


## azure_network_virtual_networks

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|address_space|jsonb|added|
|address_space_address_prefixes|text[]|removed|
|bgp_communities|jsonb|added|
|bgp_communities_regional_community|text|removed|
|bgp_communities_virtual_network_community|text|removed|
|ddos_protection_plan|jsonb|added|
|ddos_protection_plan_id|text|removed|
|dhcp_options|jsonb|added|
|dhcp_options_dns_servers|inet[]|removed|
|extended_location|jsonb|added|
|extended_location_name|text|removed|
|extended_location_type|text|removed|
|ip_allocations|jsonb|updated|Type changed from text[] to jsonb
|subnets|jsonb|added|
|virtual_network_peerings|jsonb|added|

## azure_network_watchers

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|

## azure_postgresql_configurations
This table was newly added.

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|subscription_id|text|added|
|postgresql_server_id|uuid|added|
|value|text|added|
|description|text|added|
|default_value|text|added|
|data_type|text|added|
|allowed_values|text|added|
|source|text|added|
|id|text|added|
|name|text|added|
|type|text|added|
|_cq_id|uuid|added|
|_cq_fetch_time|timestamp without time zone|added|

## azure_postgresql_firewall_rules
This table was newly added.

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|subscription_id|text|added|
|postgresql_server_id|uuid|added|
|start_ip_address|text|added|
|end_ip_address|text|added|
|id|text|added|
|name|text|added|
|type|text|added|
|_cq_id|uuid|added|
|_cq_fetch_time|timestamp without time zone|added|

## azure_postgresql_server_configurations
Moved to JSON column on [azure_postgresql_servers](#azure_postgresql_servers)


## azure_postgresql_server_firewall_rules
Moved to JSON column on [azure_postgresql_servers](#azure_postgresql_servers)


## azure_postgresql_server_private_endpoint_connections
Moved to JSON column on [azure_postgresql_servers](#azure_postgresql_servers)


## azure_postgresql_servers

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|earliest_restore_date|timestamp without time zone|added|
|earliest_restore_date_time|timestamp without time zone|removed|
|identity|jsonb|added|
|identity_principal_id|uuid|removed|
|identity_tenant_id|uuid|removed|
|identity_type|text|removed|
|private_endpoint_connections|jsonb|added|
|replica_capacity|bigint|updated|Type changed from integer to bigint
|sku|jsonb|added|
|sku_capacity|integer|removed|
|sku_family|text|removed|
|sku_name|text|removed|
|sku_size|text|removed|
|sku_tier|text|removed|
|storage_profile|jsonb|added|
|storage_profile_backup_retention_days|integer|removed|
|storage_profile_geo_redundant_backup|text|removed|
|storage_profile_storage_autogrow|text|removed|
|storage_profile_storage_mb|integer|removed|

## azure_redis_caches
This table was newly added.

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|subscription_id|text|added|
|provisioning_state|text|added|
|host_name|text|added|
|port|bigint|added|
|ssl_port|bigint|added|
|access_keys|jsonb|added|
|linked_servers|jsonb|added|
|instances|jsonb|added|
|private_endpoint_connections|jsonb|added|
|sku|jsonb|added|
|subnet_id|text|added|
|static_ip|text|added|
|redis_configuration|jsonb|added|
|redis_version|text|added|
|enable_non_ssl_port|boolean|added|
|replicas_per_master|bigint|added|
|replicas_per_primary|bigint|added|
|tenant_settings|jsonb|added|
|shard_count|bigint|added|
|minimum_tls_version|text|added|
|public_network_access|text|added|
|zones|text[]|added|
|tags|jsonb|added|
|location|text|added|
|id|text|added|
|name|text|added|
|type|text|added|
|_cq_id|uuid|added|
|_cq_fetch_time|timestamp without time zone|added|

## azure_redis_services
This table was removed.


## azure_resources_groups

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|

## azure_resources_links

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|notes|text|removed|
|properties_notes|text|added|
|properties_source_id|text|added|
|properties_target_id|text|added|
|source_id|text|removed|
|target_id|text|removed|
|type|text|removed|

## azure_resources_policy_assignments

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|identity|jsonb|added|
|identity_principal_id|text|removed|
|identity_tenant_id|text|removed|
|identity_type|text|removed|
|metadata|jsonb|removed|
|sku|jsonb|added|
|sku_name|text|removed|
|sku_tier|text|removed|

## azure_search_service_private_endpoint_connections
Moved to JSON column on [azure_search_services](#azure_search_services)


## azure_search_service_shared_private_link_resources
Moved to JSON column on [azure_search_services](#azure_search_services)


## azure_search_services

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|identity|jsonb|added|
|identity_principal_id|text|removed|
|identity_tenant_id|text|removed|
|identity_type|text|removed|
|network_rule_set|jsonb|added|
|network_rule_set_ip_rules|inet[]|removed|
|partition_count|bigint|updated|Type changed from integer to bigint
|private_endpoint_connections|jsonb|added|
|replica_count|bigint|updated|Type changed from integer to bigint
|shared_private_link_resources|jsonb|added|
|sku|jsonb|added|
|sku_name|text|removed|

## azure_security_assessments

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|azure_portal_uri|text|removed|
|cause|text|removed|
|code|text|removed|
|description|text|removed|
|links|jsonb|added|
|metadata|jsonb|added|
|metadata_assessment_type|text|removed|
|metadata_categories|text[]|removed|
|metadata_description|text|removed|
|metadata_display_name|text|removed|
|metadata_implementation_effort|text|removed|
|metadata_partner_data_partner_name|text|removed|
|metadata_partner_data_product_name|text|removed|
|metadata_policy_definition_id|text|removed|
|metadata_preview|boolean|removed|
|metadata_remediation_description|text|removed|
|metadata_severity|text|removed|
|metadata_threats|text[]|removed|
|metadata_user_impact|text|removed|
|partner_name|text|removed|
|partners_data|jsonb|added|
|resource_details|jsonb|removed|
|status|jsonb|added|

## azure_security_auto_provisioning_settings

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|resource_type|text|removed|
|type|text|added|

## azure_security_contacts

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|resource_type|text|removed|
|type|text|added|

## azure_security_jit_network_access_policies

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|requests|jsonb|added|
|virtual_machines|jsonb|added|

## azure_security_jit_network_access_policy_requests
Moved to JSON column on [azure_security_jit_network_access_policies](#azure_security_jit_network_access_policies)


## azure_security_jit_network_access_policy_virtual_machines
Moved to JSON column on [azure_security_jit_network_access_policies](#azure_security_jit_network_access_policies)


## azure_security_pricings

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|free_trial_remaining_time|text|added|
|pricing_properties_free_trial_remaining_time|text|removed|
|pricing_properties_tier|text|removed|
|pricing_tier|text|added|
|resource_type|text|removed|
|type|text|added|

## azure_security_settings

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|resource_type|text|removed|
|type|text|added|

## azure_servicebus_access_keys
This table was newly added.

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|subscription_id|text|added|
|servicebus_authorization_rule_id|uuid|added|
|primary_connection_string|text|added|
|secondary_connection_string|text|added|
|alias_primary_connection_string|text|added|
|alias_secondary_connection_string|text|added|
|primary_key|text|added|
|secondary_key|text|added|
|key_name|text|added|
|_cq_id|uuid|added|
|_cq_fetch_time|timestamp without time zone|added|

## azure_servicebus_authorization_rules
This table was newly added.

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|subscription_id|text|added|
|servicebus_topic_id|uuid|added|
|rights|text[]|added|
|system_data|jsonb|added|
|id|text|added|
|name|text|added|
|type|text|added|
|_cq_id|uuid|added|
|_cq_fetch_time|timestamp without time zone|added|

## azure_servicebus_namespace_private_endpoint_connections
Moved to JSON column on [azure_servicebus_namespaces](#azure_servicebus_namespaces)


## azure_servicebus_namespace_topic_authorization_rules
Moved to JSON column on [azure_servicebus_namespaces](#azure_servicebus_namespaces)


## azure_servicebus_namespace_topics
Moved to JSON column on [azure_servicebus_namespaces](#azure_servicebus_namespaces)


## azure_servicebus_namespaces

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|created_at|timestamp without time zone|added|
|created_at_time|timestamp without time zone|removed|
|encryption|jsonb|added|
|identity|jsonb|added|
|identity_principal_id|text|removed|
|identity_tenant_id|text|removed|
|identity_type|text|removed|
|key_source|text|removed|
|key_vault_properties|jsonb|removed|
|private_endpoint_connections|jsonb|added|
|require_infrastructure_encryption|boolean|removed|
|sku|jsonb|added|
|sku_capacity|integer|removed|
|sku_name|text|removed|
|sku_tier|text|removed|
|updated_at|timestamp without time zone|added|
|updated_at_time|timestamp without time zone|removed|
|user_assigned_identities|jsonb|removed|

## azure_servicebus_topics
This table was newly added.

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|subscription_id|text|added|
|servicebus_namespace_id|uuid|added|
|size_in_bytes|bigint|added|
|created_at|timestamp without time zone|added|
|updated_at|timestamp without time zone|added|
|accessed_at|timestamp without time zone|added|
|subscription_count|bigint|added|
|count_details|jsonb|added|
|default_message_time_to_live|text|added|
|max_size_in_megabytes|bigint|added|
|max_message_size_in_kilobytes|bigint|added|
|requires_duplicate_detection|boolean|added|
|duplicate_detection_history_time_window|text|added|
|enable_batched_operations|boolean|added|
|status|text|added|
|support_ordering|boolean|added|
|auto_delete_on_idle|text|added|
|enable_partitioning|boolean|added|
|enable_express|boolean|added|
|system_data|jsonb|added|
|id|text|added|
|name|text|added|
|type|text|added|
|_cq_id|uuid|added|
|_cq_fetch_time|timestamp without time zone|added|

## azure_sql_backup_long_term_retention_policies
This table was newly added.

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|subscription_id|text|added|
|sql_database_id|uuid|added|
|weekly_retention|text|added|
|monthly_retention|text|added|
|yearly_retention|text|added|
|week_of_year|bigint|added|
|id|text|added|
|name|text|added|
|type|text|added|
|_cq_id|uuid|added|
|_cq_fetch_time|timestamp without time zone|added|

## azure_sql_database_blob_auditing_policies
This table was newly added.

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|subscription_id|text|added|
|sql_database_id|uuid|added|
|kind|text|added|
|state|text|added|
|storage_endpoint|text|added|
|storage_account_access_key|text|added|
|retention_days|bigint|added|
|audit_actions_and_groups|text[]|added|
|storage_account_subscription_id|uuid|added|
|is_storage_secondary_key_in_use|boolean|added|
|is_azure_monitor_target_enabled|boolean|added|
|queue_delay_ms|bigint|added|
|id|text|added|
|name|text|added|
|type|text|added|
|_cq_id|uuid|added|
|_cq_fetch_time|timestamp without time zone|added|

## azure_sql_database_db_blob_auditing_policies
Moved to JSON column on [azure_sql_databases](#azure_sql_databases)


## azure_sql_database_db_threat_detection_policies
Moved to JSON column on [azure_sql_databases](#azure_sql_databases)


## azure_sql_database_db_vulnerability_assessment_scans
Moved to JSON column on [azure_sql_databases](#azure_sql_databases)


## azure_sql_database_db_vulnerability_assessments
Moved to JSON column on [azure_sql_databases](#azure_sql_databases)


## azure_sql_database_threat_detection_policies
This table was newly added.

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|subscription_id|text|added|
|sql_database_id|uuid|added|
|location|text|added|
|kind|text|added|
|state|text|added|
|disabled_alerts|text|added|
|email_addresses|text|added|
|email_account_admins|text|added|
|storage_endpoint|text|added|
|storage_account_access_key|text|added|
|retention_days|bigint|added|
|use_server_default|text|added|
|id|text|added|
|name|text|added|
|type|text|added|
|_cq_id|uuid|added|
|_cq_fetch_time|timestamp without time zone|added|

## azure_sql_database_vulnerability_assessment_scans
This table was newly added.

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|subscription_id|text|added|
|sql_database_id|uuid|added|
|scan_id|text|added|
|trigger_type|text|added|
|state|text|added|
|start_time|timestamp without time zone|added|
|end_time|timestamp without time zone|added|
|errors|jsonb|added|
|storage_container_path|text|added|
|number_of_failed_security_checks|bigint|added|
|id|text|added|
|name|text|added|
|type|text|added|
|_cq_id|uuid|added|
|_cq_fetch_time|timestamp without time zone|added|

## azure_sql_database_vulnerability_assessments
This table was newly added.

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|subscription_id|text|added|
|sql_database_id|uuid|added|
|storage_container_path|text|added|
|storage_container_sas_key|text|added|
|storage_account_access_key|text|added|
|recurring_scans|jsonb|added|
|id|text|added|
|name|text|added|
|type|text|added|
|_cq_id|uuid|added|
|_cq_fetch_time|timestamp without time zone|added|

## azure_sql_databases

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|auto_pause_delay|bigint|updated|Type changed from integer to bigint
|backup_long_term_retention_policy|jsonb|removed|
|creation_date|timestamp without time zone|added|
|creation_date_time|timestamp without time zone|removed|
|current_sku|jsonb|added|
|current_sku_capacity|integer|removed|
|current_sku_family|text|removed|
|current_sku_name|text|removed|
|current_sku_size|text|removed|
|current_sku_tier|text|removed|
|earliest_restore_date|timestamp without time zone|added|
|earliest_restore_date_time|timestamp without time zone|removed|
|high_availability_replica_count|bigint|updated|Type changed from integer to bigint
|min_capacity|real|updated|Type changed from float to real
|paused_date|timestamp without time zone|added|
|paused_date_time|timestamp without time zone|removed|
|resumed_date|timestamp without time zone|added|
|resumed_date_time|timestamp without time zone|removed|
|server_cq_id|uuid|removed|
|sku|jsonb|added|
|sku_capacity|integer|removed|
|sku_family|text|removed|
|sku_name|text|removed|
|sku_size|text|removed|
|sku_tier|text|removed|
|source_database_deletion_date|timestamp without time zone|added|
|source_database_deletion_date_time|timestamp without time zone|removed|
|sql_server_id|uuid|added|
|subscription_id|text|added|
|transparent_data_encryption|jsonb|removed|

## azure_sql_encryption_protectors
This table was newly added.

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|subscription_id|text|added|
|sql_server_id|uuid|added|
|kind|text|added|
|location|text|added|
|subregion|text|added|
|server_key_name|text|added|
|server_key_type|text|added|
|uri|text|added|
|thumbprint|text|added|
|id|text|added|
|name|text|added|
|type|text|added|
|_cq_id|uuid|added|
|_cq_fetch_time|timestamp without time zone|added|

## azure_sql_firewall_rules
This table was newly added.

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|subscription_id|text|added|
|sql_server_id|uuid|added|
|kind|text|added|
|location|text|added|
|start_ip_address|text|added|
|end_ip_address|text|added|
|id|text|added|
|name|text|added|
|type|text|added|
|_cq_id|uuid|added|
|_cq_fetch_time|timestamp without time zone|added|

## azure_sql_managed_database_vulnerability_assessment_scans

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|managed_database_cq_id|uuid|removed|
|number_of_failed_security_checks|bigint|updated|Type changed from integer to bigint
|sql_managed_database_id|uuid|added|
|subscription_id|text|added|

## azure_sql_managed_database_vulnerability_assessments

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|managed_database_cq_id|uuid|removed|
|recurring_scans|jsonb|added|
|recurring_scans_email_subscription_admins|boolean|removed|
|recurring_scans_emails|text[]|removed|
|recurring_scans_is_enabled|boolean|removed|
|sql_managed_database_id|uuid|added|
|subscription_id|text|added|

## azure_sql_managed_databases

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|creation_date|timestamp without time zone|added|
|creation_date_time|timestamp without time zone|removed|
|earliest_restore_point|timestamp without time zone|added|
|earliest_restore_point_time|timestamp without time zone|removed|
|sql_managed_instance_id|uuid|added|
|subscription_id|text|added|

## azure_sql_managed_instance_encryption_protectors

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|managed_instance_cq_id|uuid|removed|
|sql_managed_instance_id|uuid|added|
|subscription_id|text|added|

## azure_sql_managed_instance_private_endpoint_connections
Moved to JSON column on [azure_sql_managed_instances](#azure_sql_managed_instances)


## azure_sql_managed_instance_vulnerability_assessments

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|managed_instance_cq_id|uuid|removed|
|recurring_scans|jsonb|added|
|recurring_scans_email_subscription_admins|boolean|removed|
|recurring_scans_emails|text[]|removed|
|recurring_scans_is_enabled|boolean|removed|
|sql_managed_instance_id|uuid|added|
|subscription_id|text|added|

## azure_sql_managed_instances

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|administrator_login_password|text|added|
|identity|jsonb|added|
|identity_principal_id|uuid|removed|
|identity_tenant_id|uuid|removed|
|identity_type|text|removed|
|private_endpoint_connections|jsonb|added|
|sku|jsonb|added|
|sku_capacity|integer|removed|
|sku_family|text|removed|
|sku_name|text|removed|
|sku_size|text|removed|
|sku_tier|text|removed|
|storage_size_in_gb|bigint|updated|Type changed from integer to bigint
|v_cores|bigint|updated|Type changed from integer to bigint

## azure_sql_server_admins

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|server_cq_id|uuid|removed|
|sql_server_id|uuid|added|
|subscription_id|text|added|

## azure_sql_server_blob_auditing_policies
This table was newly added.

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|subscription_id|text|added|
|sql_server_id|uuid|added|
|state|text|added|
|storage_endpoint|text|added|
|storage_account_access_key|text|added|
|retention_days|bigint|added|
|audit_actions_and_groups|text[]|added|
|storage_account_subscription_id|uuid|added|
|is_storage_secondary_key_in_use|boolean|added|
|is_azure_monitor_target_enabled|boolean|added|
|queue_delay_ms|bigint|added|
|id|text|added|
|name|text|added|
|type|text|added|
|_cq_id|uuid|added|
|_cq_fetch_time|timestamp without time zone|added|

## azure_sql_server_db_blob_auditing_policies
Moved to JSON column on [azure_sql_servers](#azure_sql_servers)


## azure_sql_server_dev_ops_auditing_settings
This table was newly added.

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|subscription_id|text|added|
|sql_server_id|uuid|added|
|system_data|jsonb|added|
|is_azure_monitor_target_enabled|boolean|added|
|state|text|added|
|storage_endpoint|text|added|
|storage_account_access_key|text|added|
|storage_account_subscription_id|uuid|added|
|id|text|added|
|name|text|added|
|type|text|added|
|_cq_id|uuid|added|
|_cq_fetch_time|timestamp without time zone|added|

## azure_sql_server_devops_audit_settings
Moved to JSON column on [azure_sql_servers](#azure_sql_servers)


## azure_sql_server_encryption_protectors
Moved to JSON column on [azure_sql_servers](#azure_sql_servers)


## azure_sql_server_firewall_rules
Moved to JSON column on [azure_sql_servers](#azure_sql_servers)


## azure_sql_server_private_endpoint_connections
Moved to JSON column on [azure_sql_servers](#azure_sql_servers)


## azure_sql_server_security_alert_policies
This table was newly added.

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|subscription_id|text|added|
|sql_server_id|uuid|added|
|state|text|added|
|disabled_alerts|text[]|added|
|email_addresses|text[]|added|
|email_account_admins|boolean|added|
|storage_endpoint|text|added|
|storage_account_access_key|text|added|
|retention_days|bigint|added|
|creation_time|timestamp without time zone|added|
|id|text|added|
|name|text|added|
|type|text|added|
|_cq_id|uuid|added|
|_cq_fetch_time|timestamp without time zone|added|

## azure_sql_server_security_alert_policy
Moved to JSON column on [azure_sql_servers](#azure_sql_servers)


## azure_sql_server_virtual_network_rules
Moved to JSON column on [azure_sql_servers](#azure_sql_servers)


## azure_sql_server_vulnerability_assessments

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|recurring_scans|jsonb|added|
|recurring_scans_email_subscription_admins|boolean|removed|
|recurring_scans_emails|text[]|removed|
|recurring_scans_is_enabled|boolean|removed|
|server_cq_id|uuid|removed|
|sql_server_id|uuid|added|
|subscription_id|text|added|

## azure_sql_servers

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|identity|jsonb|added|
|identity_principal_id|uuid|removed|
|identity_tenant_id|uuid|removed|
|identity_type|text|removed|
|private_endpoint_connections|jsonb|added|

## azure_sql_transparent_data_encryptions
This table was newly added.

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|subscription_id|text|added|
|sql_database_id|uuid|added|
|location|text|added|
|status|text|added|
|id|text|added|
|name|text|added|
|type|text|added|
|_cq_id|uuid|added|
|_cq_fetch_time|timestamp without time zone|added|

## azure_sql_virtual_network_rules
This table was newly added.

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|subscription_id|text|added|
|sql_server_id|uuid|added|
|virtual_network_subnet_id|text|added|
|ignore_missing_vnet_service_endpoint|boolean|added|
|state|text|added|
|id|text|added|
|name|text|added|
|type|text|added|
|_cq_id|uuid|added|
|_cq_fetch_time|timestamp without time zone|added|

## azure_storage_account_network_rule_set_ip_rules
Moved to JSON column on [azure_storage_accounts](#azure_storage_accounts)


## azure_storage_account_network_rule_set_virtual_network_rules
Moved to JSON column on [azure_storage_accounts](#azure_storage_accounts)


## azure_storage_account_private_endpoint_connections
Moved to JSON column on [azure_storage_accounts](#azure_storage_accounts)


## azure_storage_accounts

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|allow_shared_key_access|boolean|added|
|azure_files_identity_based_authentication|jsonb|added|
|blob_restore_status|jsonb|updated|Type changed from text to jsonb
|blob_restore_status_failure_reason|text|removed|
|blob_restore_status_parameters_blob_ranges|jsonb|removed|
|blob_restore_status_parameters_time_to_restore_time|timestamp without time zone|removed|
|blob_restore_status_restore_id|text|removed|
|custom_domain|jsonb|added|
|custom_domain_name|text|removed|
|custom_domain_use_sub_domain_name|boolean|removed|
|enable_https_traffic_only|boolean|removed|
|encryption|jsonb|added|
|encryption_key_current_versioned_key_identifier|text|removed|
|encryption_key_last_key_rotation_timestamp_time|timestamp without time zone|removed|
|encryption_key_source|text|removed|
|encryption_key_vault_properties_key_name|text|removed|
|encryption_key_vault_properties_key_vault_uri|text|removed|
|encryption_key_vault_properties_key_version|text|removed|
|encryption_require_infrastructure_encryption|boolean|removed|
|encryption_services_blob_enabled|boolean|removed|
|encryption_services_blob_key_type|text|removed|
|encryption_services_blob_last_enabled_time|timestamp without time zone|removed|
|encryption_services_file_enabled|boolean|removed|
|encryption_services_file_key_type|text|removed|
|encryption_services_file_last_enabled_time|timestamp without time zone|removed|
|encryption_services_queue_enabled|boolean|removed|
|encryption_services_queue_key_type|text|removed|
|encryption_services_queue_last_enabled_time|timestamp without time zone|removed|
|encryption_services_table_enabled|boolean|removed|
|encryption_services_table_key_type|text|removed|
|encryption_services_table_last_enabled_time|timestamp without time zone|removed|
|extended_location|jsonb|added|
|files_identity_auth_ad_properties_azure_storage_sid|text|removed|
|files_identity_auth_ad_properties_domain_guid|text|removed|
|files_identity_auth_ad_properties_domain_name|text|removed|
|files_identity_auth_ad_properties_forest_name|text|removed|
|files_identity_auth_ad_properties_net_bios_domain_name|text|removed|
|files_identity_auth_ad_properties_net_bios_domain_sid|text|removed|
|files_identity_auth_directory_service_options|text|removed|
|geo_replication_stats|jsonb|added|
|geo_replication_stats_can_failover|boolean|removed|
|geo_replication_stats_last_sync_time|timestamp without time zone|removed|
|geo_replication_stats_status|text|removed|
|identity|jsonb|added|
|identity_principal_id|text|removed|
|identity_tenant_id|text|removed|
|identity_type|text|removed|
|is_nfs_v3_enabled|boolean|added|
|network_acls|jsonb|added|
|network_rule_set_bypass|text|removed|
|network_rule_set_default_action|text|removed|
|primary_endpoints|jsonb|added|
|primary_endpoints_blob|text|removed|
|primary_endpoints_dfs|text|removed|
|primary_endpoints_file|text|removed|
|primary_endpoints_internet_endpoints_blob|text|removed|
|primary_endpoints_internet_endpoints_dfs|text|removed|
|primary_endpoints_internet_endpoints_file|text|removed|
|primary_endpoints_internet_endpoints_web|text|removed|
|primary_endpoints_microsoft_endpoints_blob|text|removed|
|primary_endpoints_microsoft_endpoints_dfs|text|removed|
|primary_endpoints_microsoft_endpoints_file|text|removed|
|primary_endpoints_microsoft_endpoints_queue|text|removed|
|primary_endpoints_microsoft_endpoints_table|text|removed|
|primary_endpoints_microsoft_endpoints_web|text|removed|
|primary_endpoints_queue|text|removed|
|primary_endpoints_table|text|removed|
|primary_endpoints_web|text|removed|
|private_endpoint_connections|jsonb|added|
|routing_preference|jsonb|added|
|routing_preference_publish_internet_endpoints|boolean|removed|
|routing_preference_publish_microsoft_endpoints|boolean|removed|
|routing_preference_routing_choice|text|removed|
|secondary_endpoints|jsonb|added|
|secondary_endpoints_blob|text|removed|
|secondary_endpoints_dfs|text|removed|
|secondary_endpoints_file|text|removed|
|secondary_endpoints_internet_endpoints_blob|text|removed|
|secondary_endpoints_internet_endpoints_dfs|text|removed|
|secondary_endpoints_internet_endpoints_file|text|removed|
|secondary_endpoints_internet_endpoints_web|text|removed|
|secondary_endpoints_microsoft_endpoints_blob|text|removed|
|secondary_endpoints_microsoft_endpoints_dfs|text|removed|
|secondary_endpoints_microsoft_endpoints_file|text|removed|
|secondary_endpoints_microsoft_endpoints_queue|text|removed|
|secondary_endpoints_microsoft_endpoints_table|text|removed|
|secondary_endpoints_microsoft_endpoints_web|text|removed|
|secondary_endpoints_queue|text|removed|
|secondary_endpoints_table|text|removed|
|secondary_endpoints_web|text|removed|
|sku|jsonb|added|
|sku_name|text|removed|
|sku_tier|text|removed|
|supports_https_traffic_only|boolean|added|

## azure_storage_blob_service_cors_rules
Moved to JSON column on [azure_storage_blob_services](#azure_storage_blob_services)


## azure_storage_blob_services

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|account_cq_id|uuid|removed|
|change_feed|jsonb|added|
|change_feed_enabled|boolean|removed|
|change_feed_retention_in_days|integer|removed|
|container_delete_retention_policy|jsonb|added|
|container_delete_retention_policy_days|integer|removed|
|container_delete_retention_policy_enabled|boolean|removed|
|cors|jsonb|added|
|delete_retention_policy|jsonb|added|
|delete_retention_policy_days|integer|removed|
|delete_retention_policy_enabled|boolean|removed|
|last_access_time_tracking_policy|jsonb|added|
|last_access_time_tracking_policy_blob_type|text[]|removed|
|last_access_time_tracking_policy_enable|boolean|removed|
|last_access_time_tracking_policy_name|text|removed|
|last_access_time_tracking_policy_tracking_granularity_in_days|integer|removed|
|restore_policy|jsonb|added|
|restore_policy_days|integer|removed|
|restore_policy_enabled|boolean|removed|
|restore_policy_last_enabled_time|timestamp without time zone|removed|
|restore_policy_min_restore_time|timestamp without time zone|removed|
|sku|jsonb|added|
|sku_name|text|removed|
|sku_tier|text|removed|
|storage_account_id|uuid|added|
|subscription_id|text|added|

## azure_storage_containers

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|account_cq_id|uuid|removed|
|account_id|text|removed|
|remaining_retention_days|bigint|updated|Type changed from integer to bigint
|storage_account_id|uuid|added|

## azure_streamanalytics_jobs
This table was removed.


## azure_streamanalytics_streaming_jobs
This table was newly added.

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|subscription_id|text|added|
|sku|jsonb|added|
|job_id|text|added|
|provisioning_state|text|added|
|job_state|text|added|
|job_type|text|added|
|output_start_mode|text|added|
|output_start_time|timestamp without time zone|added|
|last_output_event_time|timestamp without time zone|added|
|events_out_of_order_policy|text|added|
|output_error_policy|text|added|
|events_out_of_order_max_delay_in_seconds|bigint|added|
|events_late_arrival_max_delay_in_seconds|bigint|added|
|data_locale|text|added|
|compatibility_level|text|added|
|created_date|timestamp without time zone|added|
|inputs|jsonb|added|
|transformation|jsonb|added|
|outputs|jsonb|added|
|functions|jsonb|added|
|etag|text|added|
|job_storage_account|jsonb|added|
|content_storage_policy|text|added|
|cluster|jsonb|added|
|identity|jsonb|added|
|tags|jsonb|added|
|location|text|added|
|id|text|added|
|name|text|added|
|type|text|added|
|_cq_id|uuid|added|
|_cq_fetch_time|timestamp without time zone|added|

## azure_subscription_subscriptions
Moved to JSON column on [azure_subscriptions](#azure_subscriptions)


## azure_subscription_tenants
Moved to JSON column on [azure_subscriptions](#azure_subscriptions)


## azure_subscriptions
This table was newly added.

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|subscription_id|text|added|
|authorization_source|text|added|
|managed_by_tenants|jsonb|added|
|subscription_policies|jsonb|added|
|tags|jsonb|added|
|display_name|text|added|
|id|text|added|
|state|text|added|
|tenant_id|text|added|
|_cq_id|uuid|added|
|_cq_fetch_time|timestamp without time zone|added|

## azure_subscriptions_locations
This table was newly added.

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|subscription_id|text|added|
|metadata|jsonb|added|
|display_name|text|added|
|id|text|added|
|name|text|added|
|regional_display_name|text|added|
|type|text|added|
|_cq_id|uuid|added|
|_cq_fetch_time|timestamp without time zone|added|

## azure_subscriptions_tenants
This table was newly added.

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|subscription_id|text|added|
|country|text|added|
|country_code|text|added|
|default_domain|text|added|
|display_name|text|added|
|domains|jsonb|added|
|id|text|added|
|tenant_branding_logo_url|text|added|
|tenant_category|text|added|
|tenant_id|text|added|
|tenant_type|text|added|
|_cq_id|uuid|added|
|_cq_fetch_time|timestamp without time zone|added|

## azure_web_app_auth_settings
Moved to JSON column on [azure_web_apps](#azure_web_apps)


## azure_web_app_host_name_ssl_states
Moved to JSON column on [azure_web_apps](#azure_web_apps)


## azure_web_app_publishing_profiles
Moved to JSON column on [azure_web_apps](#azure_web_apps)


## azure_web_apps

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|cloning_info|jsonb|added|
|cloning_info_app_settings_overrides|jsonb|removed|
|cloning_info_clone_custom_host_names|boolean|removed|
|cloning_info_clone_source_control|boolean|removed|
|cloning_info_configure_load_balancing|boolean|removed|
|cloning_info_correlation_id|uuid|removed|
|cloning_info_hosting_environment|text|removed|
|cloning_info_overwrite|boolean|removed|
|cloning_info_source_web_app_id|text|removed|
|cloning_info_source_web_app_location|text|removed|
|cloning_info_traffic_manager_profile_id|text|removed|
|cloning_info_traffic_manager_profile_name|text|removed|
|container_size|bigint|updated|Type changed from integer to bigint
|daily_memory_time_quota|bigint|updated|Type changed from integer to bigint
|host_name_ssl_states|jsonb|added|
|hosting_environment_profile|jsonb|added|
|hosting_environment_profile_id|text|removed|
|hosting_environment_profile_name|text|removed|
|hosting_environment_profile_type|text|removed|
|identity|jsonb|added|
|identity_principal_id|text|removed|
|identity_tenant_id|text|removed|
|identity_type|text|removed|
|identity_user_assigned_identities|jsonb|removed|
|last_modified_time_utc|timestamp without time zone|added|
|last_modified_time_utc_time|timestamp without time zone|removed|
|max_number_of_workers|bigint|updated|Type changed from integer to bigint
|slot_swap_status|jsonb|added|
|slot_swap_status_destination_slot_name|text|removed|
|slot_swap_status_source_slot_name|text|removed|
|slot_swap_status_timestamp_utc_time|timestamp without time zone|removed|
|suspended_till|timestamp without time zone|added|
|suspended_till_time|timestamp without time zone|removed|
|virtual_network_subnet_id|text|added|
|vnet_connection|jsonb|removed|

## azure_web_publishing_profiles
This table was newly added.

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|subscription_id|text|added|
|web_app_id|uuid|added|
|publish_url|text|added|
|user_name|text|added|
|user_pwd|text|added|
|_cq_id|uuid|added|
|_cq_fetch_time|timestamp without time zone|added|

## azure_web_site_auth_settings
This table was newly added.

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|subscription_id|text|added|
|web_app_id|uuid|added|
|enabled|boolean|added|
|runtime_version|text|added|
|unauthenticated_client_action|text|added|
|token_store_enabled|boolean|added|
|allowed_external_redirect_urls|text[]|added|
|default_provider|text|added|
|token_refresh_extension_hours|real|added|
|client_id|text|added|
|client_secret|text|added|
|client_secret_setting_name|text|added|
|client_secret_certificate_thumbprint|text|added|
|issuer|text|added|
|validate_issuer|boolean|added|
|allowed_audiences|text[]|added|
|additional_login_params|text[]|added|
|aad_claims_authorization|text|added|
|google_client_id|text|added|
|google_client_secret|text|added|
|google_client_secret_setting_name|text|added|
|google_o_auth_scopes|text[]|added|
|facebook_app_id|text|added|
|facebook_app_secret|text|added|
|facebook_app_secret_setting_name|text|added|
|facebook_o_auth_scopes|text[]|added|
|git_hub_client_id|text|added|
|git_hub_client_secret|text|added|
|git_hub_client_secret_setting_name|text|added|
|git_hub_o_auth_scopes|text[]|added|
|twitter_consumer_key|text|added|
|twitter_consumer_secret|text|added|
|twitter_consumer_secret_setting_name|text|added|
|microsoft_account_client_id|text|added|
|microsoft_account_client_secret|text|added|
|microsoft_account_client_secret_setting_name|text|added|
|microsoft_account_o_auth_scopes|text[]|added|
|is_auth_from_file|text|added|
|auth_file_path|text|added|
|config_version|text|added|
|id|text|added|
|name|text|added|
|kind|text|added|
|type|text|added|
|_cq_id|uuid|added|
|_cq_fetch_time|timestamp without time zone|added|

## azure_web_vnet_connections
This table was newly added.

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|subscription_id|text|added|
|web_app_id|uuid|added|
|vnet_resource_id|text|added|
|cert_thumbprint|text|added|
|cert_blob|text|added|
|routes|jsonb|added|
|resync_required|boolean|added|
|dns_servers|text|added|
|is_swift|boolean|added|
|id|text|added|
|name|text|added|
|kind|text|added|
|type|text|added|
|_cq_id|uuid|added|
|_cq_fetch_time|timestamp without time zone|added|

## cloudflare_access_groups

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|

## cloudflare_account_member_roles
Moved to JSON column on [cloudflare_accounts](#cloudflare_accounts)


## cloudflare_account_members

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|account_id|text|removed|
|roles|jsonb|added|
|user|jsonb|added|
|user_email|text|removed|
|user_first_name|text|removed|
|user_id|text|removed|
|user_last_name|text|removed|
|user_two_factor_authentication_enabled|boolean|removed|

## cloudflare_accounts

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|enforce_two_factor|boolean|removed|
|settings|jsonb|added|

## cloudflare_certificate_pack_certificates
Moved to JSON column on [cloudflare_certificate_packs](#cloudflare_certificate_packs)


## cloudflare_certificate_pack_validation_errors
Moved to JSON column on [cloudflare_certificate_packs](#cloudflare_certificate_packs)


## cloudflare_certificate_pack_validation_records
Moved to JSON column on [cloudflare_certificate_packs](#cloudflare_certificate_packs)


## cloudflare_certificate_packs

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|certificates|jsonb|added|
|validation_errors|jsonb|added|
|validation_records|jsonb|added|

## cloudflare_dns_records

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|priority|bigint|updated|Type changed from integer to bigint

## cloudflare_images

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|require_signed_ur_ls|boolean|added|
|require_signed_url_s|boolean|removed|
|variants|text[]|updated|Type changed from jsonb to text[]

## cloudflare_ips

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|

## cloudflare_waf
This table was removed.


## cloudflare_waf_groups
This table was newly added.

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|waf_package_cq_id|uuid|added|
|id|text|added|
|name|text|added|
|description|text|added|
|rules_count|bigint|added|
|modified_rules_count|bigint|added|
|package_id|text|added|
|mode|text|added|
|allowed_modes|text[]|added|
|_cq_id|uuid|added|
|_cq_fetch_time|timestamp without time zone|added|

## cloudflare_waf_overrides

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|url_s|text[]|removed|
|urls|text[]|added|

## cloudflare_waf_packages
This table was newly added.

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|account_id|text|added|
|id|text|added|
|name|text|added|
|description|text|added|
|zone_id|text|added|
|detection_mode|text|added|
|sensitivity|text|added|
|action_mode|text|added|
|_cq_id|uuid|added|
|_cq_fetch_time|timestamp without time zone|added|

## cloudflare_waf_rule_groups
Moved to JSON column on [cloudflare_waf_rules](#cloudflare_waf_rules)


## cloudflare_waf_rules

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|account_id|text|removed|
|waf_cq_id|uuid|removed|
|waf_package_cq_id|uuid|added|
|zone_id|text|removed|

## cloudflare_worker_cron_triggers
This table was newly added.

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|worker_meta_data_cq_id|uuid|added|
|cron|text|added|
|created_on|timestamp without time zone|added|
|modified_on|timestamp without time zone|added|
|_cq_id|uuid|added|
|_cq_fetch_time|timestamp without time zone|added|

## cloudflare_worker_meta_data
This table was newly added.

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|account_id|text|added|
|id|text|added|
|etag|text|added|
|size|bigint|added|
|created_on|timestamp without time zone|added|
|modified_on|timestamp without time zone|added|
|_cq_id|uuid|added|
|_cq_fetch_time|timestamp without time zone|added|

## cloudflare_worker_routes
This table was newly added.

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|account_id|text|added|
|zone_id|text|added|
|id|text|added|
|pattern|text|added|
|enabled|boolean|added|
|script|text|added|
|_cq_id|uuid|added|
|_cq_fetch_time|timestamp without time zone|added|

## cloudflare_workers_routes
This table was removed.


## cloudflare_workers_script_cron_triggers
This table was removed.


## cloudflare_workers_script_secrets
This table was removed.


## cloudflare_workers_scripts
This table was removed.


## cloudflare_workers_secrets
This table was newly added.

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|worker_meta_data_cq_id|uuid|added|
|name|text|added|
|type|text|added|
|_cq_id|uuid|added|
|_cq_fetch_time|timestamp without time zone|added|

## cloudflare_zones

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|account|jsonb|added|
|deact_reason|text|added|
|deactivation_reason|text|removed|
|host|jsonb|added|
|host_name|text|removed|
|host_website|text|removed|
|meta|jsonb|added|
|owner|jsonb|added|
|owner_email|text|removed|
|owner_id|text|removed|
|owner_name|text|removed|
|owner_type|text|removed|
|page_rule_quota|bigint|removed|
|phishing_detected|boolean|removed|
|plan|jsonb|added|
|plan_can_subscribe|boolean|removed|
|plan_currency|text|removed|
|plan_externally_managed|boolean|removed|
|plan_frequency|text|removed|
|plan_id|text|removed|
|plan_is_subscribed|boolean|removed|
|plan_legacy_discount|boolean|removed|
|plan_legacy_id|text|removed|
|plan_name|text|removed|
|plan_pending|jsonb|added|
|plan_pending_can_subscribe|boolean|removed|
|plan_pending_currency|text|removed|
|plan_pending_externally_managed|boolean|removed|
|plan_pending_frequency|text|removed|
|plan_pending_id|text|removed|
|plan_pending_is_subscribed|boolean|removed|
|plan_pending_legacy_discount|boolean|removed|
|plan_pending_legacy_id|text|removed|
|plan_pending_name|text|removed|
|plan_pending_price|bigint|removed|
|plan_price|bigint|removed|
|wildcard_proxiable|boolean|removed|

## digitalocean_accounts

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|reserved_ip_limit|bigint|added|
|team|jsonb|added|

## digitalocean_balance
Moved to JSON column on [digitalocean_balances](#digitalocean_balances)


## digitalocean_balances
This table was newly added.

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|month_to_date_balance|text|added|
|account_balance|text|added|
|month_to_date_usage|text|added|
|generated_at|timestamp without time zone|added|
|_cq_id|uuid|added|
|_cq_fetch_time|timestamp without time zone|added|

## digitalocean_billing_history

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|

## digitalocean_cdn
Moved to JSON column on [digitalocean_cdns](#digitalocean_cdns)


## digitalocean_cdns
This table was newly added.

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|id|text|added|
|origin|text|added|
|endpoint|text|added|
|created_at|timestamp without time zone|added|
|ttl|bigint|added|
|certificate_id|text|added|
|custom_domain|text|added|
|_cq_id|uuid|added|
|_cq_fetch_time|timestamp without time zone|added|

## digitalocean_certificates

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|created|text|removed|
|created_at|text|added|
|s_h_a1_fingerprint|text|removed|
|sha_1_fingerprint|text|added|

## digitalocean_database_backups

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|database_cq_id|uuid|removed|
|size_gigabytes|real|updated|Type changed from float to real

## digitalocean_database_firewall_rules

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|database_cq_id|uuid|removed|

## digitalocean_database_replicas

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|connection|jsonb|added|
|connection_database|text|removed|
|connection_host|text|removed|
|connection_password|text|removed|
|connection_port|bigint|removed|
|connection_ssl|boolean|removed|
|connection_uri|text|removed|
|connection_user|text|removed|
|database_cq_id|uuid|removed|
|private_connection|jsonb|added|
|private_connection_database|text|removed|
|private_connection_host|text|removed|
|private_connection_password|text|removed|
|private_connection_port|bigint|removed|
|private_connection_ssl|boolean|removed|
|private_connection_uri|text|removed|
|private_connection_user|text|removed|

## digitalocean_database_users
Moved to JSON column on [digitalocean_databases](#digitalocean_databases)


## digitalocean_databases

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|connection|jsonb|added|
|connection_database|text|removed|
|connection_host|text|removed|
|connection_password|text|removed|
|connection_port|bigint|removed|
|connection_ssl|boolean|removed|
|connection_uri|text|removed|
|connection_user|text|removed|
|maintenance_window|jsonb|added|
|maintenance_window_day|text|removed|
|maintenance_window_description|text[]|removed|
|maintenance_window_hour|text|removed|
|maintenance_window_pending|boolean|removed|
|private_connection|jsonb|added|
|private_connection_database|text|removed|
|private_connection_host|text|removed|
|private_connection_password|text|removed|
|private_connection_port|bigint|removed|
|private_connection_ssl|boolean|removed|
|private_connection_uri|text|removed|
|private_connection_user|text|removed|
|project_id|text|added|
|region|text|added|
|region_slug|text|removed|
|size|text|added|
|size_slug|text|removed|
|users|jsonb|added|

## digitalocean_domain_records

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|domain_cq_id|uuid|removed|
|id|text|updated|Type changed from bigint to text

## digitalocean_domains

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|

## digitalocean_droplet_neighbors

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|droplet_cq_id|uuid|removed|

## digitalocean_droplet_networks_v4
Moved to JSON column on [digitalocean_droplets](#digitalocean_droplets)


## digitalocean_droplet_networks_v6
Moved to JSON column on [digitalocean_droplets](#digitalocean_droplets)


## digitalocean_droplets

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|backup_ids|bigint[]|updated|Type changed from integer[] to bigint[]
|created|text|removed|
|created_at|text|added|
|image|jsonb|added|
|image_created|text|removed|
|image_description|text|removed|
|image_distribution|text|removed|
|image_error_message|text|removed|
|image_id|bigint|removed|
|image_min_disk_size|bigint|removed|
|image_name|text|removed|
|image_public|boolean|removed|
|image_regions|text[]|removed|
|image_size_giga_bytes|float|removed|
|image_slug|text|removed|
|image_status|text|removed|
|image_tags|text[]|removed|
|image_type|text|removed|
|kernel|jsonb|added|
|kernel_id|bigint|removed|
|kernel_name|text|removed|
|kernel_version|text|removed|
|networks|jsonb|added|
|next_backup_window|jsonb|added|
|next_backup_window_end_time|timestamp without time zone|removed|
|next_backup_window_start_time|timestamp without time zone|removed|
|region|jsonb|added|
|region_available|boolean|removed|
|region_features|text[]|removed|
|region_name|text|removed|
|region_sizes|text[]|removed|
|region_slug|text|removed|
|size|jsonb|added|
|size_available|boolean|removed|
|size_description|text|removed|
|size_disk|bigint|removed|
|size_memory|bigint|removed|
|size_price_hourly|float|removed|
|size_price_monthly|float|removed|
|size_regions|text[]|removed|
|size_transfer|float|removed|
|size_vcpus|bigint|removed|
|snapshot_ids|bigint[]|updated|Type changed from integer[] to bigint[]
|volume_ids|bigint[]|updated|Type changed from text[] to bigint[]

## digitalocean_firewall_droplets
Moved to JSON column on [digitalocean_firewalls](#digitalocean_firewalls)


## digitalocean_firewall_inbound_rules
Moved to JSON column on [digitalocean_firewalls](#digitalocean_firewalls)


## digitalocean_firewall_outbound_rules
Moved to JSON column on [digitalocean_firewalls](#digitalocean_firewalls)


## digitalocean_firewall_pending_changes
Moved to JSON column on [digitalocean_firewalls](#digitalocean_firewalls)


## digitalocean_firewalls

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|created|text|removed|
|created_at|text|added|
|droplet_ids|bigint[]|updated|Type changed from integer[] to bigint[]
|id|text|updated|Type changed from uuid to text
|inbound_rules|jsonb|added|
|outbound_rules|jsonb|added|
|pending_changes|jsonb|added|

## digitalocean_floating_ips

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|droplet|jsonb|added|
|droplet_id|bigint|removed|
|ip|text|updated|Type changed from cidr to text
|region|jsonb|added|
|region_available|boolean|removed|
|region_features|text[]|removed|
|region_name|text|removed|
|region_sizes|text[]|removed|
|region_slug|text|removed|

## digitalocean_images

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|created|text|removed|
|created_at|text|added|
|size_giga_bytes|float|removed|
|size_gigabytes|real|added|

## digitalocean_keys

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|

## digitalocean_load_balancer_droplets
This table was removed.


## digitalocean_load_balancer_forwarding_rules
This table was removed.


## digitalocean_load_balancers
This table was removed.


## digitalocean_monitoring_alert_policies
This table was newly added.

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|uuid|text|added|
|type|text|added|
|description|text|added|
|compare|text|added|
|value|real|added|
|window|text|added|
|entities|text[]|added|
|tags|text[]|added|
|alerts|jsonb|added|
|enabled|boolean|added|
|_cq_id|uuid|added|
|_cq_fetch_time|timestamp without time zone|added|

## digitalocean_project_resources

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|assigned_at|text|updated|Type changed from timestamp without time zone to text
|links|jsonb|added|
|links_self|text|removed|
|project_cq_id|uuid|removed|

## digitalocean_projects

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|

## digitalocean_regions
This table was removed.


## digitalocean_registries
This table was newly added.

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|name|text|added|
|storage_usage_bytes|bigint|added|
|storage_usage_bytes_updated_at|timestamp without time zone|added|
|created_at|timestamp without time zone|added|
|region|text|added|
|_cq_id|uuid|added|
|_cq_fetch_time|timestamp without time zone|added|

## digitalocean_registry
Moved to JSON column on [digitalocean_registries](#digitalocean_registries)


## digitalocean_registry_repositories

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|latest_tag|jsonb|updated|Type changed from text to jsonb
|latest_tag_compressed_size_bytes|bigint|removed|
|latest_tag_manifest_digest|text|removed|
|latest_tag_registry_name|text|removed|
|latest_tag_repository|text|removed|
|latest_tag_size_bytes|bigint|removed|
|latest_tag_updated_at|timestamp without time zone|removed|
|registry_cq_id|uuid|removed|

## digitalocean_sizes

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|price_hourly|real|updated|Type changed from float to real
|price_monthly|real|updated|Type changed from float to real
|transfer|real|updated|Type changed from float to real

## digitalocean_snapshots

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|created|text|removed|
|created_at|text|added|
|size_giga_bytes|float|removed|
|size_gigabytes|real|added|

## digitalocean_space_acls
Moved to JSON column on [digitalocean_spaces](#digitalocean_spaces)


## digitalocean_space_cors

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|expose_headers|text[]|added|
|id|text|added|
|max_age_seconds|bigint|updated|Type changed from integer to bigint
|space_cq_id|uuid|removed|
|space_name|text|removed|

## digitalocean_spaces

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|acls|jsonb|added|
|bucket|jsonb|added|
|creation_date|timestamp without time zone|removed|
|name|text|removed|

## digitalocean_storage_volumes
This table was newly added.

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|id|text|added|
|droplet_ids|bigint[]|added|
|region|jsonb|added|
|name|text|added|
|size_gigabytes|bigint|added|
|description|text|added|
|created_at|timestamp without time zone|added|
|filesystem_type|text|added|
|filesystem_label|text|added|
|tags|text[]|added|
|_cq_id|uuid|added|
|_cq_fetch_time|timestamp without time zone|added|

## digitalocean_volume_droplets
This table was removed.


## digitalocean_volumes
This table was removed.


## digitalocean_vpc_members

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|id|text|removed|
|type|text|removed|
|vpc_cq_id|uuid|removed|

## digitalocean_vpcs

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|ip_range|text|updated|Type changed from cidr to text
|region|text|added|
|region_slug|text|removed|

## gcp_bigquery_dataset_accesses
Moved to JSON column on [gcp_bigquery_datasets](#gcp_bigquery_datasets)


## gcp_bigquery_dataset_table_dataset_model_training_runs
Moved to JSON column on [gcp_bigquery_datasets](#gcp_bigquery_datasets)


## gcp_bigquery_dataset_table_user_defined_functions
Moved to JSON column on [gcp_bigquery_datasets](#gcp_bigquery_datasets)


## gcp_bigquery_dataset_tables
Moved to JSON column on [gcp_bigquery_datasets](#gcp_bigquery_datasets)


## gcp_bigquery_datasets

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|access|jsonb|added|
|dataset_reference|jsonb|added|
|default_collation|text|added|
|default_encryption_configuration|jsonb|added|
|default_encryption_configuration_kms_key_name|text|removed|
|is_case_insensitive|boolean|added|
|max_time_travel_hours|bigint|added|
|tags|jsonb|added|

## gcp_bigquery_tables
This table was newly added.

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|project_id|text|added|
|clone_definition|jsonb|added|
|clustering|jsonb|added|
|creation_time|bigint|added|
|default_collation|text|added|
|description|text|added|
|encryption_configuration|jsonb|added|
|etag|text|added|
|expiration_time|bigint|added|
|external_data_configuration|jsonb|added|
|friendly_name|text|added|
|id|text|added|
|kind|text|added|
|labels|jsonb|added|
|last_modified_time|bigint|added|
|location|text|added|
|materialized_view|jsonb|added|
|max_staleness|text|added|
|model|jsonb|added|
|num_bytes|bigint|added|
|num_long_term_bytes|bigint|added|
|num_physical_bytes|bigint|added|
|num_rows|bigint|added|
|num_active_logical_bytes|bigint|added|
|num_active_physical_bytes|bigint|added|
|num_long_term_logical_bytes|bigint|added|
|num_long_term_physical_bytes|bigint|added|
|num_partitions|bigint|added|
|num_time_travel_physical_bytes|bigint|added|
|num_total_logical_bytes|bigint|added|
|num_total_physical_bytes|bigint|added|
|range_partitioning|jsonb|added|
|require_partition_filter|boolean|added|
|schema|jsonb|added|
|self_link|text|added|
|snapshot_definition|jsonb|added|
|streaming_buffer|jsonb|added|
|table_reference|jsonb|added|
|time_partitioning|jsonb|added|
|type|text|added|
|view|jsonb|added|
|_cq_id|uuid|added|
|_cq_fetch_time|timestamp without time zone|added|

## gcp_billing_billing_accounts
This table was newly added.

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|project_id|text|added|
|name|text|added|
|open|boolean|added|
|display_name|text|added|
|master_billing_account|text|added|
|_cq_id|uuid|added|
|_cq_fetch_time|timestamp without time zone|added|

## gcp_billing_services
This table was newly added.

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|project_id|text|added|
|name|text|added|
|service_id|text|added|
|display_name|text|added|
|business_entity_name|text|added|
|_cq_id|uuid|added|
|_cq_fetch_time|timestamp without time zone|added|

## gcp_cloudbilling_accounts
This table was removed.


## gcp_cloudbilling_service_sku_pricing_info
This table was removed.


## gcp_cloudbilling_service_sku_pricing_info_tiered_rates
This table was removed.


## gcp_cloudbilling_service_skus
This table was removed.


## gcp_cloudbilling_services
This table was removed.


## gcp_cloudfunctions_functions
This table was removed.


## gcp_cloudrun_service_metadata_owner_references
This table was removed.


## gcp_cloudrun_service_spec_template_container_env
This table was removed.


## gcp_cloudrun_service_spec_template_container_volume_mounts
This table was removed.


## gcp_cloudrun_service_spec_template_containers
This table was removed.


## gcp_cloudrun_service_spec_template_metadata_owner_references
This table was removed.


## gcp_cloudrun_service_spec_template_volume_config_map_items
This table was removed.


## gcp_cloudrun_service_spec_template_volume_secret_items
This table was removed.


## gcp_cloudrun_service_spec_template_volumes
This table was removed.


## gcp_cloudrun_service_spec_traffic
This table was removed.


## gcp_cloudrun_service_status_conditions
This table was removed.


## gcp_cloudrun_service_status_traffic
This table was removed.


## gcp_cloudrun_services
This table was removed.


## gcp_compute_addresses

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|id|bigint|updated|Type changed from text to bigint

## gcp_compute_autoscaler_custom_metric_utilizations
Moved to JSON column on [gcp_compute_autoscalers](#gcp_compute_autoscalers)


## gcp_compute_autoscalers

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|autoscaling_policy|jsonb|added|
|cool_down_period_sec|bigint|removed|
|cpu_utilization_predictive_method|text|removed|
|cpu_utilization_utilization_target|float|removed|
|id|bigint|updated|Type changed from text to bigint
|load_balancing_utilization_utilization_target|float|removed|
|max_num_replicas|bigint|removed|
|min_num_replicas|bigint|removed|
|mode|text|removed|
|scale_in_control_max_scaled_in_replicas_calculated|bigint|removed|
|scale_in_control_max_scaled_in_replicas_fixed|bigint|removed|
|scale_in_control_max_scaled_in_replicas_percent|bigint|removed|
|scale_in_control_time_window_sec|bigint|removed|
|scaling_schedules|jsonb|removed|

## gcp_compute_backend_service_backends
Moved to JSON column on [gcp_compute_backend_services](#gcp_compute_backend_services)


## gcp_compute_backend_services

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|backends|jsonb|added|
|cdn_policy|jsonb|added|
|cdn_policy_bypass_cache_on_request_headers|text[]|removed|
|cdn_policy_cache_key_policy_include_host|boolean|removed|
|cdn_policy_cache_key_policy_include_protocol|boolean|removed|
|cdn_policy_cache_key_policy_include_query_string|boolean|removed|
|cdn_policy_cache_key_policy_query_string_blacklist|text[]|removed|
|cdn_policy_cache_key_policy_query_string_whitelist|text[]|removed|
|cdn_policy_cache_mode|text|removed|
|cdn_policy_client_ttl|bigint|removed|
|cdn_policy_default_ttl|bigint|removed|
|cdn_policy_max_ttl|bigint|removed|
|cdn_policy_negative_caching|boolean|removed|
|cdn_policy_negative_caching_policy|jsonb|removed|
|cdn_policy_request_coalescing|boolean|removed|
|cdn_policy_serve_while_stale|bigint|removed|
|cdn_policy_signed_url_cache_max_age_sec|bigint|removed|
|cdn_policy_signed_url_key_names|text[]|removed|
|circuit_breakers|jsonb|added|
|circuit_breakers_max_connections|bigint|removed|
|circuit_breakers_max_pending_requests|bigint|removed|
|circuit_breakers_max_requests|bigint|removed|
|circuit_breakers_max_requests_per_connection|bigint|removed|
|circuit_breakers_max_retries|bigint|removed|
|compression_mode|text|added|
|connection_draining|jsonb|added|
|connection_draining_draining_timeout_sec|bigint|removed|
|connection_tracking_policy|jsonb|added|
|consistent_hash|jsonb|added|
|consistent_hash_http_cookie_name|text|removed|
|consistent_hash_http_cookie_path|text|removed|
|consistent_hash_http_cookie_ttl_nanos|bigint|removed|
|consistent_hash_http_cookie_ttl_seconds|bigint|removed|
|consistent_hash_http_header_name|text|removed|
|consistent_hash_minimum_ring_size|bigint|removed|
|edge_security_policy|text|added|
|failover_policy|jsonb|added|
|failover_policy_disable_connection_drain_on_failover|boolean|removed|
|failover_policy_drop_traffic_if_unhealthy|boolean|removed|
|failover_policy_failover_ratio|float|removed|
|iap|jsonb|added|
|iap_enabled|boolean|removed|
|iap_oauth2_client_id|text|removed|
|iap_oauth2_client_secret|text|removed|
|iap_oauth2_client_secret_sha256|text|removed|
|id|bigint|updated|Type changed from text to bigint
|locality_lb_policies|jsonb|added|
|log_config|jsonb|added|
|log_config_enable|boolean|removed|
|log_config_sample_rate|float|removed|
|max_stream_duration|jsonb|added|
|max_stream_duration_nanos|bigint|removed|
|max_stream_duration_seconds|bigint|removed|
|outlier_detection|jsonb|added|
|outlier_detection_base_ejection_time_nanos|bigint|removed|
|outlier_detection_base_ejection_time_seconds|bigint|removed|
|outlier_detection_consecutive_errors|bigint|removed|
|outlier_detection_consecutive_gateway_failure|bigint|removed|
|outlier_detection_enforcing_consecutive_errors|bigint|removed|
|outlier_detection_enforcing_consecutive_gateway_failure|bigint|removed|
|outlier_detection_enforcing_success_rate|bigint|removed|
|outlier_detection_interval_nanos|bigint|removed|
|outlier_detection_interval_seconds|bigint|removed|
|outlier_detection_max_ejection_percent|bigint|removed|
|outlier_detection_success_rate_minimum_hosts|bigint|removed|
|outlier_detection_success_rate_request_volume|bigint|removed|
|outlier_detection_success_rate_stdev_factor|bigint|removed|
|security_settings|jsonb|added|
|security_settings_client_tls_policy|text|removed|
|security_settings_subject_alt_names|text[]|removed|
|service_bindings|text[]|added|
|subsetting|jsonb|added|

## gcp_compute_disk_types

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|deprecated|jsonb|updated|Type changed from text to jsonb
|deprecated_deleted|text|removed|
|deprecated_obsolete|text|removed|
|deprecated_replacement|text|removed|
|deprecated_state|text|removed|
|id|bigint|updated|Type changed from text to bigint

## gcp_compute_disks

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|architecture|text|added|
|disk_encryption_key|jsonb|added|
|disk_encryption_key_kms_key_name|text|removed|
|disk_encryption_key_kms_key_service_account|text|removed|
|disk_encryption_key_raw_key|text|removed|
|disk_encryption_key_sha256|text|removed|
|guest_os_features|jsonb|updated|Type changed from text[] to jsonb
|id|bigint|updated|Type changed from text to bigint
|license_codes|bigint[]|added|
|params|jsonb|added|
|source_image_encryption_key|jsonb|added|
|source_image_encryption_key_kms_key_name|text|removed|
|source_image_encryption_key_kms_key_service_account|text|removed|
|source_image_encryption_key_raw_key|text|removed|
|source_image_encryption_key_sha256|text|removed|
|source_snapshot_encryption_key|jsonb|added|
|source_snapshot_encryption_key_kms_key_name|text|removed|
|source_snapshot_encryption_key_kms_key_service_account|text|removed|
|source_snapshot_encryption_key_raw_key|text|removed|
|source_snapshot_encryption_key_sha256|text|removed|

## gcp_compute_firewall_allowed
Moved to JSON column on [gcp_compute_firewalls](#gcp_compute_firewalls)


## gcp_compute_firewall_denied
Moved to JSON column on [gcp_compute_firewalls](#gcp_compute_firewalls)


## gcp_compute_firewalls

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|allowed|jsonb|added|
|denied|jsonb|added|
|id|bigint|updated|Type changed from text to bigint
|log_config|jsonb|added|
|log_config_enable|boolean|removed|
|log_config_metadata|text|removed|

## gcp_compute_forwarding_rules

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|id|bigint|updated|Type changed from text to bigint
|metadata_filters|jsonb|added|
|no_automate_dns_zone|boolean|added|
|psc_connection_status|text|added|
|service_directory_registrations|jsonb|added|

## gcp_compute_images

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|architecture|text|added|
|deprecated|jsonb|updated|Type changed from text to jsonb
|deprecated_deleted|text|removed|
|deprecated_obsolete|text|removed|
|deprecated_replacement|text|removed|
|deprecated_state|text|removed|
|guest_os_features|jsonb|updated|Type changed from text[] to jsonb
|id|bigint|updated|Type changed from text to bigint
|image_encryption_key|jsonb|added|
|image_encryption_key_kms_key_name|text|removed|
|image_encryption_key_kms_key_service_account|text|removed|
|image_encryption_key_raw_key|text|removed|
|image_encryption_key_sha256|text|removed|
|license_codes|bigint[]|added|
|raw_disk|jsonb|added|
|raw_disk_container_type|text|removed|
|raw_disk_source|text|removed|
|shielded_instance_initial_state|jsonb|added|
|shielded_instance_initial_state_pk_content|text|removed|
|shielded_instance_initial_state_pk_file_type|text|removed|
|source_disk_encryption_key|jsonb|added|
|source_disk_encryption_key_kms_key_name|text|removed|
|source_disk_encryption_key_kms_key_service_account|text|removed|
|source_disk_encryption_key_raw_key|text|removed|
|source_disk_encryption_key_sha256|text|removed|
|source_image_encryption_key|jsonb|added|
|source_image_encryption_key_kms_key_name|text|removed|
|source_image_encryption_key_kms_key_service_account|text|removed|
|source_image_encryption_key_raw_key|text|removed|
|source_image_encryption_key_sha256|text|removed|
|source_snapshot_encryption_key|jsonb|added|
|source_snapshot_encryption_key_kms_key_name|text|removed|
|source_snapshot_encryption_key_kms_key_service_account|text|removed|
|source_snapshot_encryption_key_raw_key|text|removed|
|source_snapshot_encryption_key_sha256|text|removed|

## gcp_compute_instance_disks
Moved to JSON column on [gcp_compute_instances](#gcp_compute_instances)


## gcp_compute_instance_group_instances
Moved to JSON column on [gcp_compute_instances](#gcp_compute_instances)


## gcp_compute_instance_groups

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|creation_timestamp|text|updated|Type changed from timestamp without time zone to text
|id|bigint|updated|Type changed from text to bigint

## gcp_compute_instance_network_interface_access_configs
Moved to JSON column on [gcp_compute_instances](#gcp_compute_instances)


## gcp_compute_instance_network_interface_alias_ip_ranges
Moved to JSON column on [gcp_compute_instances](#gcp_compute_instances)


## gcp_compute_instance_network_interfaces
Moved to JSON column on [gcp_compute_instances](#gcp_compute_instances)


## gcp_compute_instance_scheduling_node_affinities
Moved to JSON column on [gcp_compute_instances](#gcp_compute_instances)


## gcp_compute_instance_service_accounts
Moved to JSON column on [gcp_compute_instances](#gcp_compute_instances)


## gcp_compute_instances

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|advanced_machine_features|jsonb|added|
|advanced_machine_features_enable_nested_virtualization|boolean|removed|
|confidential_instance_config|jsonb|added|
|confidential_instance_config_enable_confidential_compute|boolean|removed|
|disks|jsonb|added|
|display_device|jsonb|added|
|display_device_enable_display|boolean|removed|
|id|bigint|updated|Type changed from text to bigint
|key_revocation_action_type|text|added|
|metadata|jsonb|added|
|metadata_fingerprint|text|removed|
|metadata_items|jsonb|removed|
|metadata_kind|text|removed|
|network_interfaces|jsonb|added|
|network_performance_config|jsonb|added|
|params|jsonb|added|
|reservation_affinity|jsonb|added|
|reservation_affinity_consume_reservation_type|text|removed|
|reservation_affinity_key|text|removed|
|reservation_affinity_values|text[]|removed|
|scheduling|jsonb|added|
|scheduling_automatic_restart|boolean|removed|
|scheduling_location_hint|text|removed|
|scheduling_min_node_cpus|bigint|removed|
|scheduling_on_host_maintenance|text|removed|
|scheduling_preemptible|boolean|removed|
|service_accounts|jsonb|added|
|shielded_instance_config|jsonb|added|
|shielded_instance_config_enable_integrity_monitoring|boolean|removed|
|shielded_instance_config_enable_secure_boot|boolean|removed|
|shielded_instance_config_enable_vtpm|boolean|removed|
|shielded_instance_integrity_policy|jsonb|added|
|shielded_instance_integrity_policy_update_auto_learn_policy|boolean|removed|
|source_machine_image|text|added|
|source_machine_image_encryption_key|jsonb|added|
|tags|jsonb|added|
|tags_fingerprint|text|removed|
|tags_items|text[]|removed|

## gcp_compute_interconnect_circuit_infos
Moved to JSON column on [gcp_compute_interconnects](#gcp_compute_interconnects)


## gcp_compute_interconnect_expected_outages
Moved to JSON column on [gcp_compute_interconnects](#gcp_compute_interconnects)


## gcp_compute_interconnects

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|circuit_infos|jsonb|added|
|expected_outages|jsonb|added|
|id|bigint|updated|Type changed from text to bigint
|satisfies_pzs|boolean|added|

## gcp_compute_network_peerings
Moved to JSON column on [gcp_compute_networks](#gcp_compute_networks)


## gcp_compute_networks

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|enable_ula_internal_ipv6|boolean|added|
|firewall_policy|text|added|
|gateway_ip_v4|text|removed|
|gateway_ipv4|text|added|
|id|bigint|updated|Type changed from text to bigint
|internal_ipv6_range|text|added|
|ip_v4_range|text|removed|
|ipv4_range|text|added|
|network_firewall_policy_enforcement_order|text|added|
|peerings|jsonb|added|
|routing_config|jsonb|added|
|routing_config_routing_mode|text|removed|
|self_link_with_id|text|added|

## gcp_compute_project_quotas
Moved to JSON column on [gcp_compute_projects](#gcp_compute_projects)


## gcp_compute_projects

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|common_instance_metadata|jsonb|added|
|common_instance_metadata_fingerprint|text|removed|
|common_instance_metadata_items|jsonb|removed|
|common_instance_metadata_kind|text|removed|
|compute_project_id|text|removed|
|creation_timestamp|text|updated|Type changed from timestamp without time zone to text
|id|bigint|added|
|quotas|jsonb|added|
|usage_export_location|jsonb|added|
|usage_export_location_bucket_name|text|removed|
|usage_export_location_report_name_prefix|text|removed|

## gcp_compute_ssl_certificates

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|id|bigint|updated|Type changed from text to bigint
|managed|jsonb|added|
|managed_domain_status|jsonb|removed|
|managed_domains|text[]|removed|
|managed_status|text|removed|
|self_managed|jsonb|added|
|self_managed_certificate|text|removed|
|self_managed_private_key|text|removed|

## gcp_compute_ssl_policies

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|creation_timestamp|text|updated|Type changed from timestamp without time zone to text
|id|bigint|updated|Type changed from text to bigint
|region|text|added|
|warnings|jsonb|added|

## gcp_compute_ssl_policy_warnings
Moved to JSON column on [gcp_compute_ssl_policies](#gcp_compute_ssl_policies)


## gcp_compute_subnetwork_secondary_ip_ranges
Moved to JSON column on [gcp_compute_subnetworks](#gcp_compute_subnetworks)


## gcp_compute_subnetworks

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|external_ipv6_prefix|text|added|
|id|bigint|updated|Type changed from text to bigint
|internal_ipv6_prefix|text|added|
|ipv6_access_type|text|added|
|log_config|jsonb|added|
|log_config_aggregation_interval|text|removed|
|log_config_enable|boolean|removed|
|log_config_filter_expr|text|removed|
|log_config_flow_sampling|float|removed|
|log_config_metadata|text|removed|
|log_config_metadata_fields|text[]|removed|
|secondary_ip_ranges|jsonb|added|
|stack_type|text|added|

## gcp_compute_target_http_proxies

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|creation_timestamp|text|updated|Type changed from timestamp without time zone to text
|id|bigint|updated|Type changed from text to bigint

## gcp_compute_target_https_proxies
This table was removed.


## gcp_compute_target_ssl_proxies

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|certificate_map|text|added|
|creation_timestamp|text|updated|Type changed from timestamp without time zone to text
|id|bigint|updated|Type changed from text to bigint

## gcp_compute_url_map_host_rules
Moved to JSON column on [gcp_compute_url_maps](#gcp_compute_url_maps)


## gcp_compute_url_map_path_matchers
Moved to JSON column on [gcp_compute_url_maps](#gcp_compute_url_maps)


## gcp_compute_url_map_tests
Moved to JSON column on [gcp_compute_url_maps](#gcp_compute_url_maps)


## gcp_compute_url_map_weighted_backend_services
Moved to JSON column on [gcp_compute_url_maps](#gcp_compute_url_maps)


## gcp_compute_url_maps

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|cors_policy_allow_credentials|boolean|removed|
|cors_policy_allow_headers|text[]|removed|
|cors_policy_allow_methods|text[]|removed|
|cors_policy_allow_origin_regexes|text[]|removed|
|cors_policy_allow_origins|text[]|removed|
|cors_policy_disabled|boolean|removed|
|cors_policy_expose_headers|text[]|removed|
|cors_policy_max_age|bigint|removed|
|default_route_action|jsonb|added|
|default_url_redirect|jsonb|added|
|default_url_redirect_host_redirect|text|removed|
|default_url_redirect_https_redirect|boolean|removed|
|default_url_redirect_path_redirect|text|removed|
|default_url_redirect_prefix_redirect|text|removed|
|default_url_redirect_redirect_response_code|text|removed|
|default_url_redirect_strip_query|boolean|removed|
|fault_injection_policy_abort_http_status|bigint|removed|
|fault_injection_policy_abort_percentage|float|removed|
|fault_injection_policy_delay_fixed_delay_nanos|bigint|removed|
|fault_injection_policy_delay_fixed_delay_seconds|bigint|removed|
|fault_injection_policy_delay_percentage|float|removed|
|header_action|jsonb|added|
|header_action_request_headers_to_add|jsonb|removed|
|header_action_request_headers_to_remove|text[]|removed|
|header_action_response_headers_to_add|jsonb|removed|
|header_action_response_headers_to_remove|text[]|removed|
|host_rules|jsonb|added|
|id|bigint|updated|Type changed from text to bigint
|max_stream_duration_nanos|bigint|removed|
|max_stream_duration_seconds|bigint|removed|
|path_matchers|jsonb|added|
|request_mirror_policy_backend_service|text|removed|
|retry_policy_num_retries|bigint|removed|
|retry_policy_per_try_timeout_nanos|bigint|removed|
|retry_policy_per_try_timeout_seconds|bigint|removed|
|retry_policy_retry_conditions|text[]|removed|
|tests|jsonb|added|
|timeout_nanos|bigint|removed|
|timeout_seconds|bigint|removed|
|url_rewrite_host_rewrite|text|removed|
|url_rewrite_path_prefix_rewrite|text|removed|

## gcp_compute_vpn_gateway_vpn_interfaces
Moved to JSON column on [gcp_compute_vpn_gateways](#gcp_compute_vpn_gateways)


## gcp_compute_vpn_gateways

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|id|bigint|updated|Type changed from text to bigint
|stack_type|text|added|
|vpn_interfaces|jsonb|added|

## gcp_container_clusters
This table was newly added.

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|project_id|text|added|
|self_link|text|added|
|name|text|added|
|description|text|added|
|initial_node_count|bigint|added|
|node_config|jsonb|added|
|master_auth|jsonb|added|
|logging_service|text|added|
|monitoring_service|text|added|
|network|text|added|
|cluster_ipv4_cidr|text|added|
|addons_config|jsonb|added|
|subnetwork|text|added|
|node_pools|jsonb|added|
|locations|text[]|added|
|enable_kubernetes_alpha|boolean|added|
|resource_labels|jsonb|added|
|label_fingerprint|text|added|
|legacy_abac|jsonb|added|
|network_policy|jsonb|added|
|ip_allocation_policy|jsonb|added|
|master_authorized_networks_config|jsonb|added|
|maintenance_policy|jsonb|added|
|binary_authorization|jsonb|added|
|autoscaling|jsonb|added|
|network_config|jsonb|added|
|default_max_pods_constraint|jsonb|added|
|resource_usage_export_config|jsonb|added|
|authenticator_groups_config|jsonb|added|
|private_cluster_config|jsonb|added|
|database_encryption|jsonb|added|
|vertical_pod_autoscaling|jsonb|added|
|shielded_nodes|jsonb|added|
|release_channel|jsonb|added|
|workload_identity_config|jsonb|added|
|mesh_certificates|jsonb|added|
|notification_config|jsonb|added|
|confidential_nodes|jsonb|added|
|identity_service_config|jsonb|added|
|zone|text|added|
|endpoint|text|added|
|initial_cluster_version|text|added|
|current_master_version|text|added|
|current_node_version|text|added|
|create_time|text|added|
|status|bigint|added|
|status_message|text|added|
|node_ipv4_cidr_size|bigint|added|
|services_ipv4_cidr|text|added|
|instance_group_urls|text[]|added|
|current_node_count|bigint|added|
|expire_time|text|added|
|location|text|added|
|enable_tpu|boolean|added|
|tpu_ipv4_cidr_block|text|added|
|conditions|jsonb|added|
|autopilot|jsonb|added|
|id|text|added|
|node_pool_defaults|jsonb|added|
|logging_config|jsonb|added|
|monitoring_config|jsonb|added|
|node_pool_auto_config|jsonb|added|
|_cq_id|uuid|added|
|_cq_fetch_time|timestamp without time zone|added|

## gcp_dns_managed_zone_dnssec_config_default_key_specs
Moved to JSON column on [gcp_dns_managed_zones](#gcp_dns_managed_zones)


## gcp_dns_managed_zone_forwarding_config_target_name_servers
Moved to JSON column on [gcp_dns_managed_zones](#gcp_dns_managed_zones)


## gcp_dns_managed_zone_private_visibility_config_networks
Moved to JSON column on [gcp_dns_managed_zones](#gcp_dns_managed_zones)


## gcp_dns_managed_zones

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|cloud_logging_config|jsonb|added|
|dnssec_config|jsonb|added|
|dnssec_config_kind|text|removed|
|dnssec_config_non_existence|text|removed|
|dnssec_config_state|text|removed|
|forwarding_config|jsonb|added|
|forwarding_config_kind|text|removed|
|id|bigint|updated|Type changed from text to bigint
|peering_config|jsonb|added|
|peering_config_kind|text|removed|
|peering_config_target_network_deactivate_time|text|removed|
|peering_config_target_network_kind|text|removed|
|peering_config_target_network_network_url|text|removed|
|private_visibility_config|jsonb|added|
|private_visibility_config_kind|text|removed|
|reverse_lookup_config|jsonb|added|
|reverse_lookup_config_kind|text|removed|
|service_directory_config|jsonb|added|
|service_directory_config_kind|text|removed|
|service_directory_config_namespace_deletion_time|text|removed|
|service_directory_config_namespace_kind|text|removed|
|service_directory_config_namespace_namespace_url|text|removed|

## gcp_dns_policies

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|alternative_name_server_config|jsonb|added|
|alternative_name_server_config_kind|text|removed|
|id|bigint|updated|Type changed from text to bigint
|networks|jsonb|added|

## gcp_dns_policy_alternative_name_servers
Moved to JSON column on [gcp_dns_policies](#gcp_dns_policies)


## gcp_dns_policy_networks
Moved to JSON column on [gcp_dns_policies](#gcp_dns_policies)


## gcp_domains_registration_glue_records
Moved to JSON column on [gcp_domains_registrations](#gcp_domains_registrations)


## gcp_domains_registrations

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|admin_contact_email|text|removed|
|admin_contact_fax_number|text|removed|
|admin_contact_phone_number|text|removed|
|admin_contact_postal_address_address_lines|text[]|removed|
|admin_contact_postal_address_administrative_area|text|removed|
|admin_contact_postal_address_language_code|text|removed|
|admin_contact_postal_address_locality|text|removed|
|admin_contact_postal_address_organization|text|removed|
|admin_contact_postal_address_postal_code|text|removed|
|admin_contact_postal_address_recipients|text[]|removed|
|admin_contact_postal_address_region_code|text|removed|
|admin_contact_postal_address_revision|bigint|removed|
|admin_contact_postal_address_sorting_code|text|removed|
|admin_contact_postal_address_sublocality|text|removed|
|contact_settings|jsonb|added|
|create_time|jsonb|updated|Type changed from text to jsonb
|custom_dns_ds_records|jsonb|removed|
|custom_dns_name_servers|text[]|removed|
|dns_settings|jsonb|added|
|expire_time|jsonb|updated|Type changed from text to jsonb
|google_domains_dns_ds_records|jsonb|removed|
|google_domains_dns_ds_state|text|removed|
|google_domains_dns_name_servers|text[]|removed|
|issues|bigint[]|updated|Type changed from text[] to bigint[]
|management_settings|jsonb|added|
|management_settings_renewal_method|text|removed|
|management_settings_transfer_lock_state|text|removed|
|pending_contact_settings|jsonb|added|
|privacy|text|removed|
|registrant_contact_email|text|removed|
|registrant_contact_fax_number|text|removed|
|registrant_contact_phone_number|text|removed|
|registrant_contact_postal_address_address_lines|text[]|removed|
|registrant_contact_postal_address_administrative_area|text|removed|
|registrant_contact_postal_address_language_code|text|removed|
|registrant_contact_postal_address_locality|text|removed|
|registrant_contact_postal_address_organization|text|removed|
|registrant_contact_postal_address_postal_code|text|removed|
|registrant_contact_postal_address_recipients|text[]|removed|
|registrant_contact_postal_address_region_code|text|removed|
|registrant_contact_postal_address_revision|bigint|removed|
|registrant_contact_postal_address_sorting_code|text|removed|
|registrant_contact_postal_address_sublocality|text|removed|
|state|bigint|updated|Type changed from text to bigint
|supported_privacy|bigint[]|updated|Type changed from text[] to bigint[]
|technical_contact_email|text|removed|
|technical_contact_fax_number|text|removed|
|technical_contact_phone_number|text|removed|
|technical_contact_postal_address_address_lines|text[]|removed|
|technical_contact_postal_address_administrative_area|text|removed|
|technical_contact_postal_address_language_code|text|removed|
|technical_contact_postal_address_locality|text|removed|
|technical_contact_postal_address_organization|text|removed|
|technical_contact_postal_address_postal_code|text|removed|
|technical_contact_postal_address_recipients|text[]|removed|
|technical_contact_postal_address_region_code|text|removed|
|technical_contact_postal_address_revision|bigint|removed|
|technical_contact_postal_address_sorting_code|text|removed|
|technical_contact_postal_address_sublocality|text|removed|

## gcp_functions_functions
This table was newly added.

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|project_id|text|added|
|name|text|added|
|description|text|added|
|status|bigint|added|
|entry_point|text|added|
|runtime|text|added|
|timeout|jsonb|added|
|available_memory_mb|bigint|added|
|service_account_email|text|added|
|update_time|jsonb|added|
|version_id|bigint|added|
|labels|jsonb|added|
|environment_variables|jsonb|added|
|build_environment_variables|jsonb|added|
|network|text|added|
|max_instances|bigint|added|
|min_instances|bigint|added|
|vpc_connector|text|added|
|vpc_connector_egress_settings|bigint|added|
|ingress_settings|bigint|added|
|kms_key_name|text|added|
|build_worker_pool|text|added|
|build_id|text|added|
|build_name|text|added|
|secret_environment_variables|jsonb|added|
|secret_volumes|jsonb|added|
|source_token|text|added|
|docker_repository|text|added|
|docker_registry|bigint|added|
|_cq_id|uuid|added|
|_cq_fetch_time|timestamp without time zone|added|

## gcp_iam_roles

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|

## gcp_iam_service_account_keys
Moved to JSON column on [gcp_iam_service_accounts](#gcp_iam_service_accounts)


## gcp_iam_service_accounts

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|etag|text|added|
|id|text|removed|
|unique_id|text|added|

## gcp_kms_crypto_keys
This table was newly added.

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|project_id|text|added|
|name|text|added|
|primary|jsonb|added|
|purpose|bigint|added|
|create_time|jsonb|added|
|next_rotation_time|jsonb|added|
|version_template|jsonb|added|
|labels|jsonb|added|
|import_only|boolean|added|
|destroy_scheduled_duration|jsonb|added|
|crypto_key_backend|text|added|
|_cq_id|uuid|added|
|_cq_fetch_time|timestamp without time zone|added|

## gcp_kms_keyring_crypto_keys
Moved to JSON column on [gcp_kms_keyrings](#gcp_kms_keyrings)


## gcp_kms_keyrings

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|create_time|jsonb|updated|Type changed from timestamp without time zone to jsonb
|location|text|removed|

## gcp_kubernetes_cluster_node_pools
This table was removed.


## gcp_kubernetes_clusters
This table was removed.


## gcp_logging_metric_descriptor_labels
Moved to JSON column on [gcp_logging_metrics](#gcp_logging_metrics)


## gcp_logging_metrics

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|bucket_options|jsonb|added|
|create_time|jsonb|updated|Type changed from text to jsonb
|disabled|boolean|added|
|exponential_buckets_options_growth_factor|float|removed|
|exponential_buckets_options_num_finite_buckets|bigint|removed|
|exponential_buckets_options_scale|float|removed|
|linear_buckets_options_num_finite_buckets|bigint|removed|
|linear_buckets_options_offset|float|removed|
|linear_buckets_options_width|float|removed|
|metric_descriptor|jsonb|added|
|metric_descriptor_description|text|removed|
|metric_descriptor_display_name|text|removed|
|metric_descriptor_launch_stage|text|removed|
|metric_descriptor_metadata_ingest_delay|text|removed|
|metric_descriptor_metadata_sample_period|text|removed|
|metric_descriptor_metric_kind|text|removed|
|metric_descriptor_monitored_resource_types|text[]|removed|
|metric_descriptor_name|text|removed|
|metric_descriptor_type|text|removed|
|metric_descriptor_unit|text|removed|
|metric_descriptor_value_type|text|removed|
|update_time|jsonb|updated|Type changed from text to jsonb
|version|bigint|updated|Type changed from text to bigint

## gcp_logging_sink_exclusions
Moved to JSON column on [gcp_logging_sinks](#gcp_logging_sinks)


## gcp_logging_sinks

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|bigquery_options_use_partitioned_tables|boolean|removed|
|bigquery_options_uses_timestamp_column_partitioning|boolean|removed|
|create_time|jsonb|updated|Type changed from text to jsonb
|exclusions|jsonb|added|
|output_version_format|bigint|updated|Type changed from text to bigint
|update_time|jsonb|updated|Type changed from text to jsonb

## gcp_memorystore_redis_instance_server_ca_certs
This table was removed.


## gcp_memorystore_redis_instances
This table was removed.


## gcp_monitoring_alert_policies

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|alert_strategy|jsonb|added|
|combiner|bigint|updated|Type changed from text to bigint
|conditions|jsonb|added|
|creation_record|jsonb|added|
|creation_record_mutate_time|text|removed|
|creation_record_mutated_by|text|removed|
|documentation|jsonb|added|
|documentation_content|text|removed|
|documentation_mime_type|text|removed|
|enabled|jsonb|updated|Type changed from boolean to jsonb
|labels|jsonb|removed|
|mutate_time|text|removed|
|mutated_by|text|removed|
|mutation_record|jsonb|added|
|user_labels|jsonb|added|
|validity|jsonb|added|
|validity_code|bigint|removed|
|validity_message|text|removed|

## gcp_monitoring_alert_policy_condition_absent_aggregations
Moved to JSON column on [gcp_monitoring_alert_policies](#gcp_monitoring_alert_policies)


## gcp_monitoring_alert_policy_condition_denominator_aggs
Moved to JSON column on [gcp_monitoring_alert_policies](#gcp_monitoring_alert_policies)


## gcp_monitoring_alert_policy_condition_threshold_aggregations
Moved to JSON column on [gcp_monitoring_alert_policies](#gcp_monitoring_alert_policies)


## gcp_monitoring_alert_policy_conditions
Moved to JSON column on [gcp_monitoring_alert_policies](#gcp_monitoring_alert_policies)


## gcp_redis_instances
This table was newly added.

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|project_id|text|added|
|name|text|added|
|display_name|text|added|
|labels|jsonb|added|
|location_id|text|added|
|alternative_location_id|text|added|
|redis_version|text|added|
|reserved_ip_range|text|added|
|secondary_ip_range|text|added|
|host|text|added|
|port|bigint|added|
|current_location_id|text|added|
|create_time|jsonb|added|
|state|bigint|added|
|status_message|text|added|
|redis_configs|jsonb|added|
|tier|bigint|added|
|memory_size_gb|bigint|added|
|authorized_network|text|added|
|persistence_iam_identity|text|added|
|connect_mode|bigint|added|
|auth_enabled|boolean|added|
|server_ca_certs|jsonb|added|
|transit_encryption_mode|bigint|added|
|maintenance_policy|jsonb|added|
|maintenance_schedule|jsonb|added|
|replica_count|bigint|added|
|nodes|jsonb|added|
|read_endpoint|text|added|
|read_endpoint_port|bigint|added|
|read_replicas_mode|bigint|added|
|_cq_id|uuid|added|
|_cq_fetch_time|timestamp without time zone|added|

## gcp_resource_manager_folders
Renamed to [gcp_resourcemanager_folders](#gcp_resourcemanager_folders)


## gcp_resource_manager_projects
Renamed to [gcp_resourcemanager_projects](#gcp_resourcemanager_projects)


## gcp_resourcemanager_folders
Renamed from [gcp_resource_manager_folders](gcp_resource_manager_folders)


## gcp_resourcemanager_projects
Renamed from [gcp_resource_manager_projects](gcp_resource_manager_projects)


## gcp_run_services
This table was newly added.

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|project_id|text|added|
|name|text|added|
|description|text|added|
|uid|text|added|
|generation|bigint|added|
|labels|jsonb|added|
|annotations|jsonb|added|
|create_time|jsonb|added|
|update_time|jsonb|added|
|delete_time|jsonb|added|
|expire_time|jsonb|added|
|creator|text|added|
|last_modifier|text|added|
|client|text|added|
|client_version|text|added|
|ingress|bigint|added|
|launch_stage|bigint|added|
|binary_authorization|jsonb|added|
|template|jsonb|added|
|traffic|jsonb|added|
|observed_generation|bigint|added|
|terminal_condition|jsonb|added|
|conditions|jsonb|added|
|latest_ready_revision|text|added|
|latest_created_revision|text|added|
|traffic_statuses|jsonb|added|
|uri|text|added|
|reconciling|boolean|added|
|etag|text|added|
|_cq_id|uuid|added|
|_cq_fetch_time|timestamp without time zone|added|

## gcp_secretmanager_secrets
This table was newly added.

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|project_id|text|added|
|name|text|added|
|replication|jsonb|added|
|create_time|jsonb|added|
|labels|jsonb|added|
|topics|jsonb|added|
|etag|text|added|
|rotation|jsonb|added|
|version_aliases|jsonb|added|
|_cq_id|uuid|added|
|_cq_fetch_time|timestamp without time zone|added|

## gcp_security_secret_user_managed_replicas
This table was removed.


## gcp_security_secrets
This table was removed.


## gcp_serviceusage_services
This table was newly added.

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|project_id|text|added|
|name|text|added|
|parent|text|added|
|config|jsonb|added|
|state|bigint|added|
|_cq_id|uuid|added|
|_cq_fetch_time|timestamp without time zone|added|

## gcp_sql_instance_ip_addresses
Moved to JSON column on [gcp_sql_instances](#gcp_sql_instances)


## gcp_sql_instance_settings_deny_maintenance_periods
Moved to JSON column on [gcp_sql_instances](#gcp_sql_instances)


## gcp_sql_instance_settings_ip_config_authorized_networks
Moved to JSON column on [gcp_sql_instances](#gcp_sql_instances)


## gcp_sql_instances

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|available_maintenance_versions|text[]|added|
|configuration_kind|text|removed|
|create_time|text|added|
|database_installed_version|text|added|
|disk_encryption_configuration|jsonb|added|
|disk_encryption_configuration_kind|text|removed|
|disk_encryption_configuration_kms_key_name|text|removed|
|disk_encryption_status|jsonb|added|
|disk_encryption_status_kind|text|removed|
|disk_encryption_status_kms_key_version_name|text|removed|
|failover_replica|jsonb|added|
|failover_replica_available|boolean|removed|
|failover_replica_name|text|removed|
|failover_target|boolean|removed|
|id|text|removed|
|ip_addresses|jsonb|added|
|maintenance_version|text|added|
|mysql_replica_configuration_ca_certificate|text|removed|
|mysql_replica_configuration_client_certificate|text|removed|
|mysql_replica_configuration_client_key|text|removed|
|mysql_replica_configuration_connect_retry_interval|bigint|removed|
|mysql_replica_configuration_dump_file_path|text|removed|
|mysql_replica_configuration_kind|text|removed|
|mysql_replica_configuration_master_heartbeat_period|bigint|removed|
|mysql_replica_configuration_password|text|removed|
|mysql_replica_configuration_ssl_cipher|text|removed|
|mysql_replica_configuration_username|text|removed|
|mysql_replica_configuration_verify_server_certificate|boolean|removed|
|on_premises_configuration|jsonb|added|
|on_premises_configuration_ca_certificate|text|removed|
|on_premises_configuration_client_certificate|text|removed|
|on_premises_configuration_client_key|text|removed|
|on_premises_configuration_dump_file_path|text|removed|
|on_premises_configuration_host_port|text|removed|
|on_premises_configuration_kind|text|removed|
|on_premises_configuration_password|text|removed|
|on_premises_configuration_username|text|removed|
|out_of_disk_report|jsonb|added|
|replica_configuration|jsonb|added|
|scheduled_maintenance|jsonb|added|
|scheduled_maintenance_can_defer|boolean|removed|
|scheduled_maintenance_can_reschedule|boolean|removed|
|scheduled_maintenance_start_time|text|removed|
|server_ca_cert|jsonb|updated|Type changed from text to jsonb
|server_ca_cert_cert_serial_number|text|removed|
|server_ca_cert_common_name|text|removed|
|server_ca_cert_create_time|text|removed|
|server_ca_cert_expiration_time|text|removed|
|server_ca_cert_instance|text|removed|
|server_ca_cert_kind|text|removed|
|server_ca_cert_self_link|text|removed|
|server_ca_cert_sha1_fingerprint|text|removed|
|settings|jsonb|added|
|settings_activation_policy|text|removed|
|settings_active_directory_config_domain|text|removed|
|settings_active_directory_config_kind|text|removed|
|settings_authorized_gae_applications|text[]|removed|
|settings_availability_type|text|removed|
|settings_backup_binary_log_enabled|boolean|removed|
|settings_backup_enabled|boolean|removed|
|settings_backup_kind|text|removed|
|settings_backup_location|text|removed|
|settings_backup_point_in_time_recovery_enabled|boolean|removed|
|settings_backup_replication_log_archiving_enabled|boolean|removed|
|settings_backup_retention_settings_retained_backups|bigint|removed|
|settings_backup_retention_settings_retention_unit|text|removed|
|settings_backup_start_time|text|removed|
|settings_backup_transaction_log_retention_days|bigint|removed|
|settings_collation|text|removed|
|settings_crash_safe_replication_enabled|boolean|removed|
|settings_data_disk_size_gb|bigint|removed|
|settings_data_disk_type|text|removed|
|settings_database_flags|jsonb|removed|
|settings_database_replication_enabled|boolean|removed|
|settings_insights_config_query_insights_enabled|boolean|removed|
|settings_insights_config_query_string_length|bigint|removed|
|settings_insights_config_record_application_tags|boolean|removed|
|settings_insights_config_record_client_address|boolean|removed|
|settings_ip_configuration_ipv4_enabled|boolean|removed|
|settings_ip_configuration_private_network|text|removed|
|settings_ip_configuration_require_ssl|boolean|removed|
|settings_kind|text|removed|
|settings_location_preference_follow_gae_application|text|removed|
|settings_location_preference_kind|text|removed|
|settings_location_preference_secondary_zone|text|removed|
|settings_location_preference_zone|text|removed|
|settings_maintenance_window_day|bigint|removed|
|settings_maintenance_window_hour|bigint|removed|
|settings_maintenance_window_kind|text|removed|
|settings_maintenance_window_update_track|text|removed|
|settings_pricing_plan|text|removed|
|settings_replication_type|text|removed|
|settings_storage_auto_resize|boolean|removed|
|settings_storage_auto_resize_limit|bigint|removed|
|settings_tier|text|removed|
|settings_user_labels|jsonb|removed|
|settings_version|bigint|removed|

## gcp_storage_bucket_acls
Moved to JSON column on [gcp_storage_buckets](#gcp_storage_buckets)


## gcp_storage_bucket_cors
Moved to JSON column on [gcp_storage_buckets](#gcp_storage_buckets)


## gcp_storage_bucket_default_object_acls
Moved to JSON column on [gcp_storage_buckets](#gcp_storage_buckets)


## gcp_storage_bucket_lifecycle_rules
Moved to JSON column on [gcp_storage_buckets](#gcp_storage_buckets)


## gcp_storage_buckets

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|acl|jsonb|added|
|billing_requester_pays|boolean|removed|
|bucket_policy_only|jsonb|added|
|cors|jsonb|added|
|created|timestamp without time zone|added|
|custom_placement_config|jsonb|added|
|default_object_acl|jsonb|added|
|encryption|jsonb|added|
|encryption_default_kms_key_name|text|removed|
|encryption_type|text|removed|
|iam_configuration_bucket_policy_only_enabled|boolean|removed|
|iam_configuration_bucket_policy_only_locked_time|text|removed|
|iam_configuration_public_access_prevention|text|removed|
|iam_configuration_uniform_bucket_level_access_enabled|boolean|removed|
|iam_configuration_uniform_bucket_level_access_locked_time|text|removed|
|id|text|removed|
|kind|text|removed|
|lifecycle|jsonb|added|
|logging|jsonb|added|
|logging_log_bucket|text|removed|
|logging_log_object_prefix|text|removed|
|meta_generation|bigint|added|
|metageneration|bigint|removed|
|owner_entity|text|removed|
|owner_entity_id|text|removed|
|policy|jsonb|removed|
|predefined_acl|text|added|
|predefined_default_object_acl|text|added|
|public_access_prevention|bigint|added|
|requester_pays|boolean|added|
|retention_policy|jsonb|added|
|retention_policy_effective_time|text|removed|
|retention_policy_is_locked|boolean|removed|
|retention_policy_retention_period|bigint|removed|
|rpo|bigint|added|
|satisfies_pzs|boolean|removed|
|self_link|text|removed|
|time_created|text|removed|
|uniform_bucket_level_access|jsonb|added|
|updated|text|removed|
|website|jsonb|added|
|website_main_page_suffix|text|removed|
|website_not_found_page|text|removed|
|zone_affinity|text[]|removed|

## gcp_storage_metrics
This table was removed.


## github_action_billing
This table was removed.


## github_billing_action
This table was newly added.

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|org|text|added|
|total_minutes_used|bigint|added|
|total_paid_minutes_used|real|added|
|included_minutes|bigint|added|
|minutes_used_breakdown|jsonb|added|
|_cq_id|uuid|added|
|_cq_fetch_time|timestamp without time zone|added|

## github_billing_package
This table was newly added.

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|org|text|added|
|total_gigabytes_bandwidth_used|bigint|added|
|total_paid_gigabytes_bandwidth_used|bigint|added|
|included_gigabytes_bandwidth|bigint|added|
|_cq_id|uuid|added|
|_cq_fetch_time|timestamp without time zone|added|

## github_billing_storage
This table was newly added.

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|org|text|added|
|days_left_in_billing_cycle|bigint|added|
|estimated_paid_storage_for_month|real|added|
|estimated_storage_for_month|bigint|added|
|_cq_id|uuid|added|
|_cq_fetch_time|timestamp without time zone|added|

## github_external_group_members
Moved to JSON column on [github_external_groups](#github_external_groups)


## github_external_group_teams
Moved to JSON column on [github_external_groups](#github_external_groups)


## github_external_groups

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|members|jsonb|added|
|teams|jsonb|added|
|updated_at|timestamp without time zone|added|
|updated_at_time|timestamp without time zone|removed|

## github_hook_deliveries

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|delivered_at|timestamp without time zone|added|
|delivered_at_time|timestamp without time zone|removed|
|duration|real|updated|Type changed from float to real
|hook_cq_id|uuid|removed|
|hook_id|bigint|added|
|org|text|added|
|request|text|added|
|request_headers|jsonb|removed|
|request_raw_payload|bytea|removed|
|response|text|added|
|response_headers|jsonb|removed|
|response_raw_payload|bytea|removed|

## github_hooks

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|

## github_installations

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|account|jsonb|added|
|account_avatar_url|text|removed|
|account_bio|text|removed|
|account_blog|text|removed|
|account_collaborators|bigint|removed|
|account_company|text|removed|
|account_created_at_time|timestamp without time zone|removed|
|account_disk_usage|bigint|removed|
|account_email|text|removed|
|account_events_url|text|removed|
|account_followers|bigint|removed|
|account_followers_url|text|removed|
|account_following|bigint|removed|
|account_following_url|text|removed|
|account_gists_url|text|removed|
|account_gravatar_id|text|removed|
|account_hireable|boolean|removed|
|account_html_url|text|removed|
|account_id|bigint|removed|
|account_ldap_dn|text|removed|
|account_location|text|removed|
|account_login|text|removed|
|account_name|text|removed|
|account_node_id|text|removed|
|account_organizations_url|text|removed|
|account_owned_private_repos|bigint|removed|
|account_permissions|jsonb|removed|
|account_plan_collaborators|bigint|removed|
|account_plan_filled_seats|bigint|removed|
|account_plan_name|text|removed|
|account_plan_private_repos|bigint|removed|
|account_plan_seats|bigint|removed|
|account_plan_space|bigint|removed|
|account_private_gists|bigint|removed|
|account_public_gists|bigint|removed|
|account_public_repos|bigint|removed|
|account_received_events_url|text|removed|
|account_repos_url|text|removed|
|account_role_name|text|removed|
|account_site_admin|boolean|removed|
|account_starred_url|text|removed|
|account_subscriptions_url|text|removed|
|account_suspended_at_time|timestamp without time zone|removed|
|account_text_matches|jsonb|removed|
|account_total_private_repos|bigint|removed|
|account_twitter_username|text|removed|
|account_two_factor_authentication|boolean|removed|
|account_type|text|removed|
|account_updated_at_time|timestamp without time zone|removed|
|account_url|text|removed|
|created_at|timestamp without time zone|added|
|created_at_time|timestamp without time zone|removed|
|permissions|jsonb|added|
|permissions_actions|text|removed|
|permissions_administration|text|removed|
|permissions_blocking|text|removed|
|permissions_checks|text|removed|
|permissions_content_references|text|removed|
|permissions_contents|text|removed|
|permissions_deployments|text|removed|
|permissions_emails|text|removed|
|permissions_environments|text|removed|
|permissions_followers|text|removed|
|permissions_issues|text|removed|
|permissions_members|text|removed|
|permissions_metadata|text|removed|
|permissions_organization_administration|text|removed|
|permissions_organization_hooks|text|removed|
|permissions_organization_plan|text|removed|
|permissions_organization_pre_receive_hooks|text|removed|
|permissions_organization_projects|text|removed|
|permissions_organization_secrets|text|removed|
|permissions_organization_self_hosted_runners|text|removed|
|permissions_organization_user_blocking|text|removed|
|permissions_packages|text|removed|
|permissions_pages|text|removed|
|permissions_pull_requests|text|removed|
|permissions_repository_hooks|text|removed|
|permissions_repository_pre_receive_hooks|text|removed|
|permissions_repository_projects|text|removed|
|permissions_secret_scanning_alerts|text|removed|
|permissions_secrets|text|removed|
|permissions_security_events|text|removed|
|permissions_single_file|text|removed|
|permissions_statuses|text|removed|
|permissions_team_discussions|text|removed|
|permissions_vulnerability_alerts|text|removed|
|permissions_workflows|text|removed|
|suspended_at|timestamp without time zone|added|
|suspended_at_time|timestamp without time zone|removed|
|suspended_by|jsonb|added|
|suspended_by_avatar_url|text|removed|
|suspended_by_bio|text|removed|
|suspended_by_blog|text|removed|
|suspended_by_collaborators|bigint|removed|
|suspended_by_company|text|removed|
|suspended_by_created_at_time|timestamp without time zone|removed|
|suspended_by_disk_usage|bigint|removed|
|suspended_by_email|text|removed|
|suspended_by_events_url|text|removed|
|suspended_by_followers|bigint|removed|
|suspended_by_followers_url|text|removed|
|suspended_by_following|bigint|removed|
|suspended_by_following_url|text|removed|
|suspended_by_gists_url|text|removed|
|suspended_by_gravatar_id|text|removed|
|suspended_by_hireable|boolean|removed|
|suspended_by_html_url|text|removed|
|suspended_by_id|bigint|removed|
|suspended_by_ldap_dn|text|removed|
|suspended_by_location|text|removed|
|suspended_by_login|text|removed|
|suspended_by_name|text|removed|
|suspended_by_node_id|text|removed|
|suspended_by_organizations_url|text|removed|
|suspended_by_owned_private_repos|bigint|removed|
|suspended_by_permissions|jsonb|removed|
|suspended_by_plan_collaborators|bigint|removed|
|suspended_by_plan_filled_seats|bigint|removed|
|suspended_by_plan_name|text|removed|
|suspended_by_plan_private_repos|bigint|removed|
|suspended_by_plan_seats|bigint|removed|
|suspended_by_plan_space|bigint|removed|
|suspended_by_private_gists|bigint|removed|
|suspended_by_public_gists|bigint|removed|
|suspended_by_public_repos|bigint|removed|
|suspended_by_received_events_url|text|removed|
|suspended_by_repos_url|text|removed|
|suspended_by_role_name|text|removed|
|suspended_by_site_admin|boolean|removed|
|suspended_by_starred_url|text|removed|
|suspended_by_subscriptions_url|text|removed|
|suspended_by_suspended_at_time|timestamp without time zone|removed|
|suspended_by_text_matches|jsonb|removed|
|suspended_by_total_private_repos|bigint|removed|
|suspended_by_twitter_username|text|removed|
|suspended_by_two_factor_authentication|boolean|removed|
|suspended_by_type|text|removed|
|suspended_by_updated_at_time|timestamp without time zone|removed|
|suspended_by_url|text|removed|
|updated_at|timestamp without time zone|added|
|updated_at_time|timestamp without time zone|removed|

## github_issue_assignees
Moved to JSON column on [github_issues](#github_issues)


## github_issue_labels
Moved to JSON column on [github_issues](#github_issues)


## github_issues

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|assignee|jsonb|added|
|assignee_avatar_url|text|removed|
|assignee_bio|text|removed|
|assignee_blog|text|removed|
|assignee_collaborators|bigint|removed|
|assignee_company|text|removed|
|assignee_created_at_time|timestamp without time zone|removed|
|assignee_disk_usage|bigint|removed|
|assignee_email|text|removed|
|assignee_events_url|text|removed|
|assignee_followers|bigint|removed|
|assignee_followers_url|text|removed|
|assignee_following|bigint|removed|
|assignee_following_url|text|removed|
|assignee_gists_url|text|removed|
|assignee_gravatar_id|text|removed|
|assignee_hireable|boolean|removed|
|assignee_html_url|text|removed|
|assignee_id|bigint|removed|
|assignee_ldap_dn|text|removed|
|assignee_location|text|removed|
|assignee_login|text|removed|
|assignee_name|text|removed|
|assignee_node_id|text|removed|
|assignee_organizations_url|text|removed|
|assignee_owned_private_repos|bigint|removed|
|assignee_permissions|jsonb|removed|
|assignee_plan_collaborators|bigint|removed|
|assignee_plan_filled_seats|bigint|removed|
|assignee_plan_name|text|removed|
|assignee_plan_private_repos|bigint|removed|
|assignee_plan_seats|bigint|removed|
|assignee_plan_space|bigint|removed|
|assignee_private_gists|bigint|removed|
|assignee_public_gists|bigint|removed|
|assignee_public_repos|bigint|removed|
|assignee_received_events_url|text|removed|
|assignee_repos_url|text|removed|
|assignee_role_name|text|removed|
|assignee_site_admin|boolean|removed|
|assignee_starred_url|text|removed|
|assignee_subscriptions_url|text|removed|
|assignee_suspended_at_time|timestamp without time zone|removed|
|assignee_text_matches|jsonb|removed|
|assignee_total_private_repos|bigint|removed|
|assignee_twitter_username|text|removed|
|assignee_two_factor_authentication|boolean|removed|
|assignee_type|text|removed|
|assignee_updated_at_time|timestamp without time zone|removed|
|assignee_url|text|removed|
|assignees|jsonb|added|
|closed_by|jsonb|added|
|closed_by_avatar_url|text|removed|
|closed_by_bio|text|removed|
|closed_by_blog|text|removed|
|closed_by_collaborators|bigint|removed|
|closed_by_company|text|removed|
|closed_by_created_at_time|timestamp without time zone|removed|
|closed_by_disk_usage|bigint|removed|
|closed_by_email|text|removed|
|closed_by_events_url|text|removed|
|closed_by_followers|bigint|removed|
|closed_by_followers_url|text|removed|
|closed_by_following|bigint|removed|
|closed_by_following_url|text|removed|
|closed_by_gists_url|text|removed|
|closed_by_gravatar_id|text|removed|
|closed_by_hireable|boolean|removed|
|closed_by_html_url|text|removed|
|closed_by_id|bigint|removed|
|closed_by_ldap_dn|text|removed|
|closed_by_location|text|removed|
|closed_by_login|text|removed|
|closed_by_name|text|removed|
|closed_by_node_id|text|removed|
|closed_by_organizations_url|text|removed|
|closed_by_owned_private_repos|bigint|removed|
|closed_by_permissions|jsonb|removed|
|closed_by_plan_collaborators|bigint|removed|
|closed_by_plan_filled_seats|bigint|removed|
|closed_by_plan_name|text|removed|
|closed_by_plan_private_repos|bigint|removed|
|closed_by_plan_seats|bigint|removed|
|closed_by_plan_space|bigint|removed|
|closed_by_private_gists|bigint|removed|
|closed_by_public_gists|bigint|removed|
|closed_by_public_repos|bigint|removed|
|closed_by_received_events_url|text|removed|
|closed_by_repos_url|text|removed|
|closed_by_role_name|text|removed|
|closed_by_site_admin|boolean|removed|
|closed_by_starred_url|text|removed|
|closed_by_subscriptions_url|text|removed|
|closed_by_suspended_at_time|timestamp without time zone|removed|
|closed_by_text_matches|jsonb|removed|
|closed_by_total_private_repos|bigint|removed|
|closed_by_twitter_username|text|removed|
|closed_by_two_factor_authentication|boolean|removed|
|closed_by_type|text|removed|
|closed_by_updated_at_time|timestamp without time zone|removed|
|closed_by_url|text|removed|
|labels|jsonb|added|
|milestone|jsonb|added|
|milestone_closed_at|timestamp without time zone|removed|
|milestone_closed_issues|bigint|removed|
|milestone_created_at|timestamp without time zone|removed|
|milestone_creator_avatar_url|text|removed|
|milestone_creator_bio|text|removed|
|milestone_creator_blog|text|removed|
|milestone_creator_collaborators|bigint|removed|
|milestone_creator_company|text|removed|
|milestone_creator_created_at_time|timestamp without time zone|removed|
|milestone_creator_disk_usage|bigint|removed|
|milestone_creator_email|text|removed|
|milestone_creator_events_url|text|removed|
|milestone_creator_followers|bigint|removed|
|milestone_creator_followers_url|text|removed|
|milestone_creator_following|bigint|removed|
|milestone_creator_following_url|text|removed|
|milestone_creator_gists_url|text|removed|
|milestone_creator_gravatar_id|text|removed|
|milestone_creator_hireable|boolean|removed|
|milestone_creator_html_url|text|removed|
|milestone_creator_id|bigint|removed|
|milestone_creator_ldap_dn|text|removed|
|milestone_creator_location|text|removed|
|milestone_creator_login|text|removed|
|milestone_creator_name|text|removed|
|milestone_creator_node_id|text|removed|
|milestone_creator_organizations_url|text|removed|
|milestone_creator_owned_private_repos|bigint|removed|
|milestone_creator_permissions|jsonb|removed|
|milestone_creator_plan_collaborators|bigint|removed|
|milestone_creator_plan_filled_seats|bigint|removed|
|milestone_creator_plan_name|text|removed|
|milestone_creator_plan_private_repos|bigint|removed|
|milestone_creator_plan_seats|bigint|removed|
|milestone_creator_plan_space|bigint|removed|
|milestone_creator_private_gists|bigint|removed|
|milestone_creator_public_gists|bigint|removed|
|milestone_creator_public_repos|bigint|removed|
|milestone_creator_received_events_url|text|removed|
|milestone_creator_repos_url|text|removed|
|milestone_creator_role_name|text|removed|
|milestone_creator_site_admin|boolean|removed|
|milestone_creator_starred_url|text|removed|
|milestone_creator_subscriptions_url|text|removed|
|milestone_creator_suspended_at_time|timestamp without time zone|removed|
|milestone_creator_text_matches|jsonb|removed|
|milestone_creator_total_private_repos|bigint|removed|
|milestone_creator_twitter_username|text|removed|
|milestone_creator_two_factor_authentication|boolean|removed|
|milestone_creator_type|text|removed|
|milestone_creator_updated_at_time|timestamp without time zone|removed|
|milestone_creator_url|text|removed|
|milestone_description|text|removed|
|milestone_due_on|timestamp without time zone|removed|
|milestone_html_url|text|removed|
|milestone_id|bigint|removed|
|milestone_labels_url|text|removed|
|milestone_node_id|text|removed|
|milestone_number|bigint|removed|
|milestone_open_issues|bigint|removed|
|milestone_state|text|removed|
|milestone_title|text|removed|
|milestone_updated_at|timestamp without time zone|removed|
|milestone_url|text|removed|
|pull_request|jsonb|added|
|pull_request_links_diff_url|text|removed|
|pull_request_links_html_url|text|removed|
|pull_request_links_patch_url|text|removed|
|pull_request_links_url|text|removed|
|reactions|jsonb|added|
|reactions_confused|bigint|removed|
|reactions_eyes|bigint|removed|
|reactions_heart|bigint|removed|
|reactions_hooray|bigint|removed|
|reactions_laugh|bigint|removed|
|reactions_plus_one|bigint|removed|
|reactions_rocket|bigint|removed|
|reactions_total_count|bigint|removed|
|reactions_url|text|removed|
|repository|jsonb|added|
|repository_id|bigint|removed|
|user|jsonb|added|
|user_avatar_url|text|removed|
|user_bio|text|removed|
|user_blog|text|removed|
|user_collaborators|bigint|removed|
|user_company|text|removed|
|user_created_at_time|timestamp without time zone|removed|
|user_disk_usage|bigint|removed|
|user_email|text|removed|
|user_events_url|text|removed|
|user_followers|bigint|removed|
|user_followers_url|text|removed|
|user_following|bigint|removed|
|user_following_url|text|removed|
|user_gists_url|text|removed|
|user_gravatar_id|text|removed|
|user_hireable|boolean|removed|
|user_html_url|text|removed|
|user_id|bigint|removed|
|user_ldap_dn|text|removed|
|user_location|text|removed|
|user_login|text|removed|
|user_name|text|removed|
|user_node_id|text|removed|
|user_organizations_url|text|removed|
|user_owned_private_repos|bigint|removed|
|user_permissions|jsonb|removed|
|user_plan_collaborators|bigint|removed|
|user_plan_filled_seats|bigint|removed|
|user_plan_name|text|removed|
|user_plan_private_repos|bigint|removed|
|user_plan_seats|bigint|removed|
|user_plan_space|bigint|removed|
|user_private_gists|bigint|removed|
|user_public_gists|bigint|removed|
|user_public_repos|bigint|removed|
|user_received_events_url|text|removed|
|user_repos_url|text|removed|
|user_role_name|text|removed|
|user_site_admin|boolean|removed|
|user_starred_url|text|removed|
|user_subscriptions_url|text|removed|
|user_suspended_at_time|timestamp without time zone|removed|
|user_text_matches|jsonb|removed|
|user_total_private_repos|bigint|removed|
|user_twitter_username|text|removed|
|user_two_factor_authentication|boolean|removed|
|user_type|text|removed|
|user_updated_at_time|timestamp without time zone|removed|
|user_url|text|removed|

## github_organization_member_membership
Moved to JSON column on [github_organizations](#github_organizations)


## github_organization_members

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|created_at|timestamp without time zone|added|
|created_at_time|timestamp without time zone|removed|
|membership|jsonb|added|
|organization_cq_id|uuid|removed|
|plan|jsonb|added|
|plan_collaborators|bigint|removed|
|plan_filled_seats|bigint|removed|
|plan_name|text|removed|
|plan_private_repos|bigint|removed|
|plan_seats|bigint|removed|
|plan_space|bigint|removed|
|suspended_at|timestamp without time zone|added|
|suspended_at_time|timestamp without time zone|removed|
|updated_at|timestamp without time zone|added|
|updated_at_time|timestamp without time zone|removed|

## github_organizations

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|default_repo_permission|text|removed|
|default_repo_settings|text|removed|
|default_repository_permission|text|added|
|default_repository_settings|text|added|
|members_can_create_internal_repos|boolean|removed|
|members_can_create_internal_repositories|boolean|added|
|members_can_create_private_repos|boolean|removed|
|members_can_create_private_repositories|boolean|added|
|members_can_create_public_repos|boolean|removed|
|members_can_create_public_repositories|boolean|added|
|members_can_create_repos|boolean|removed|
|members_can_create_repositories|boolean|added|
|members_can_fork_private_repos|boolean|removed|
|members_can_fork_private_repositories|boolean|added|
|org|text|added|
|plan|jsonb|added|
|plan_collaborators|bigint|removed|
|plan_filled_seats|bigint|removed|
|plan_name|text|removed|
|plan_private_repos|bigint|removed|
|plan_seats|bigint|removed|
|plan_space|bigint|removed|

## github_package_billing
This table was removed.


## github_repositories

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|code_of_conduct|jsonb|added|
|code_of_conduct_body|text|removed|
|code_of_conduct_key|text|removed|
|code_of_conduct_name|text|removed|
|code_of_conduct_url|text|removed|
|created_at|timestamp without time zone|added|
|created_at_time|timestamp without time zone|removed|
|license|jsonb|added|
|license_body|text|removed|
|license_conditions|text[]|removed|
|license_description|text|removed|
|license_featured|boolean|removed|
|license_html_url|text|removed|
|license_implementation|text|removed|
|license_key|text|removed|
|license_limitations|text[]|removed|
|license_name|text|removed|
|license_permissions|text[]|removed|
|license_spdx_id|text|removed|
|license_url|text|removed|
|organization|jsonb|added|
|organization_avatar_url|text|removed|
|organization_billing_email|text|removed|
|organization_blog|text|removed|
|organization_collaborators|bigint|removed|
|organization_company|text|removed|
|organization_created_at|timestamp without time zone|removed|
|organization_default_repo_permission|text|removed|
|organization_default_repo_settings|text|removed|
|organization_description|text|removed|
|organization_disk_usage|bigint|removed|
|organization_email|text|removed|
|organization_events_url|text|removed|
|organization_followers|bigint|removed|
|organization_following|bigint|removed|
|organization_has_organization_projects|boolean|removed|
|organization_has_repository_projects|boolean|removed|
|organization_hooks_url|text|removed|
|organization_html_url|text|removed|
|organization_id|bigint|removed|
|organization_is_verified|boolean|removed|
|organization_issues_url|text|removed|
|organization_location|text|removed|
|organization_login|text|removed|
|organization_members_allowed_repository_creation_type|text|removed|
|organization_members_can_create_internal_repos|boolean|removed|
|organization_members_can_create_pages|boolean|removed|
|organization_members_can_create_private_pages|boolean|removed|
|organization_members_can_create_private_repos|boolean|removed|
|organization_members_can_create_public_pages|boolean|removed|
|organization_members_can_create_public_repos|boolean|removed|
|organization_members_can_create_repos|boolean|removed|
|organization_members_can_fork_private_repos|boolean|removed|
|organization_members_url|text|removed|
|organization_name|text|removed|
|organization_node_id|text|removed|
|organization_owned_private_repos|bigint|removed|
|organization_plan_collaborators|bigint|removed|
|organization_plan_filled_seats|bigint|removed|
|organization_plan_name|text|removed|
|organization_plan_private_repos|bigint|removed|
|organization_plan_seats|bigint|removed|
|organization_plan_space|bigint|removed|
|organization_private_gists|bigint|removed|
|organization_public_gists|bigint|removed|
|organization_public_members_url|text|removed|
|organization_public_repos|bigint|removed|
|organization_repos_url|text|removed|
|organization_total_private_repos|bigint|removed|
|organization_twitter_username|text|removed|
|organization_two_factor_requirement_enabled|boolean|removed|
|organization_type|text|removed|
|organization_updated_at|timestamp without time zone|removed|
|organization_url|text|removed|
|owner|jsonb|added|
|owner_avatar_url|text|removed|
|owner_bio|text|removed|
|owner_blog|text|removed|
|owner_collaborators|bigint|removed|
|owner_company|text|removed|
|owner_created_at_time|timestamp without time zone|removed|
|owner_disk_usage|bigint|removed|
|owner_email|text|removed|
|owner_events_url|text|removed|
|owner_followers|bigint|removed|
|owner_followers_url|text|removed|
|owner_following|bigint|removed|
|owner_following_url|text|removed|
|owner_gists_url|text|removed|
|owner_gravatar_id|text|removed|
|owner_hireable|boolean|removed|
|owner_html_url|text|removed|
|owner_id|bigint|removed|
|owner_ldap_dn|text|removed|
|owner_location|text|removed|
|owner_login|text|removed|
|owner_name|text|removed|
|owner_node_id|text|removed|
|owner_organizations_url|text|removed|
|owner_owned_private_repos|bigint|removed|
|owner_permissions|jsonb|removed|
|owner_plan_collaborators|bigint|removed|
|owner_plan_filled_seats|bigint|removed|
|owner_plan_name|text|removed|
|owner_plan_private_repos|bigint|removed|
|owner_plan_seats|bigint|removed|
|owner_plan_space|bigint|removed|
|owner_private_gists|bigint|removed|
|owner_public_gists|bigint|removed|
|owner_public_repos|bigint|removed|
|owner_received_events_url|text|removed|
|owner_repos_url|text|removed|
|owner_role_name|text|removed|
|owner_site_admin|boolean|removed|
|owner_starred_url|text|removed|
|owner_subscriptions_url|text|removed|
|owner_suspended_at_time|timestamp without time zone|removed|
|owner_text_matches|jsonb|removed|
|owner_total_private_repos|bigint|removed|
|owner_twitter_username|text|removed|
|owner_two_factor_authentication|boolean|removed|
|owner_type|text|removed|
|owner_updated_at_time|timestamp without time zone|removed|
|owner_url|text|removed|
|parent|jsonb|updated|Type changed from bigint to jsonb
|pushed_at|timestamp without time zone|added|
|pushed_at_time|timestamp without time zone|removed|
|security_and_analysis|jsonb|added|
|security_and_analysis_advanced_security_status|text|removed|
|security_and_analysis_secret_scanning_status|text|removed|
|source|jsonb|updated|Type changed from bigint to jsonb
|template_repository|jsonb|updated|Type changed from bigint to jsonb
|updated_at|timestamp without time zone|added|
|updated_at_time|timestamp without time zone|removed|

## github_storage_billing
This table was removed.


## github_team_members

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|created_at|timestamp without time zone|added|
|created_at_time|timestamp without time zone|removed|
|membership|jsonb|added|
|plan|jsonb|added|
|plan_collaborators|bigint|removed|
|plan_filled_seats|bigint|removed|
|plan_name|text|removed|
|plan_private_repos|bigint|removed|
|plan_seats|bigint|removed|
|plan_space|bigint|removed|
|suspended_at|timestamp without time zone|added|
|suspended_at_time|timestamp without time zone|removed|
|team_cq_id|uuid|removed|
|updated_at|timestamp without time zone|added|
|updated_at_time|timestamp without time zone|removed|

## github_team_repositories

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|code_of_conduct|jsonb|added|
|code_of_conduct_body|text|removed|
|code_of_conduct_key|text|removed|
|code_of_conduct_name|text|removed|
|code_of_conduct_url|text|removed|
|created_at|timestamp without time zone|added|
|created_at_time|timestamp without time zone|removed|
|license|jsonb|added|
|license_body|text|removed|
|license_conditions|text[]|removed|
|license_description|text|removed|
|license_featured|boolean|removed|
|license_html_url|text|removed|
|license_implementation|text|removed|
|license_key|text|removed|
|license_limitations|text[]|removed|
|license_name|text|removed|
|license_permissions|text[]|removed|
|license_s_p_d_x_id|text|removed|
|license_url|text|removed|
|org|text|added|
|organization|jsonb|added|
|organization_avatar_url|text|removed|
|organization_billing_email|text|removed|
|organization_blog|text|removed|
|organization_collaborators|bigint|removed|
|organization_company|text|removed|
|organization_created_at|timestamp without time zone|removed|
|organization_default_repo_permission|text|removed|
|organization_default_repo_settings|text|removed|
|organization_description|text|removed|
|organization_disk_usage|bigint|removed|
|organization_email|text|removed|
|organization_events_url|text|removed|
|organization_followers|bigint|removed|
|organization_following|bigint|removed|
|organization_has_organization_projects|boolean|removed|
|organization_has_repository_projects|boolean|removed|
|organization_hooks_url|text|removed|
|organization_html_url|text|removed|
|organization_id|bigint|removed|
|organization_is_verified|boolean|removed|
|organization_issues_url|text|removed|
|organization_location|text|removed|
|organization_login|text|removed|
|organization_members_allowed_repository_creation_type|text|removed|
|organization_members_can_create_internal_repos|boolean|removed|
|organization_members_can_create_pages|boolean|removed|
|organization_members_can_create_private_pages|boolean|removed|
|organization_members_can_create_private_repos|boolean|removed|
|organization_members_can_create_public_pages|boolean|removed|
|organization_members_can_create_public_repos|boolean|removed|
|organization_members_can_create_repos|boolean|removed|
|organization_members_can_fork_private_repos|boolean|removed|
|organization_members_url|text|removed|
|organization_name|text|removed|
|organization_node_id|text|removed|
|organization_owned_private_repos|bigint|removed|
|organization_plan_collaborators|bigint|removed|
|organization_plan_filled_seats|bigint|removed|
|organization_plan_name|text|removed|
|organization_plan_private_repos|bigint|removed|
|organization_plan_seats|bigint|removed|
|organization_plan_space|bigint|removed|
|organization_private_gists|bigint|removed|
|organization_public_gists|bigint|removed|
|organization_public_members_url|text|removed|
|organization_public_repos|bigint|removed|
|organization_repos_url|text|removed|
|organization_total_private_repos|bigint|removed|
|organization_twitter_username|text|removed|
|organization_two_factor_requirement_enabled|boolean|removed|
|organization_type|text|removed|
|organization_updated_at|timestamp without time zone|removed|
|organization_url|text|removed|
|owner|jsonb|added|
|owner_avatar_url|text|removed|
|owner_bio|text|removed|
|owner_blog|text|removed|
|owner_collaborators|bigint|removed|
|owner_company|text|removed|
|owner_created_at_time|timestamp without time zone|removed|
|owner_disk_usage|bigint|removed|
|owner_email|text|removed|
|owner_events_url|text|removed|
|owner_followers|bigint|removed|
|owner_followers_url|text|removed|
|owner_following|bigint|removed|
|owner_following_url|text|removed|
|owner_gists_url|text|removed|
|owner_gravatar_id|text|removed|
|owner_hireable|boolean|removed|
|owner_html_url|text|removed|
|owner_id|bigint|removed|
|owner_ldap_dn|text|removed|
|owner_location|text|removed|
|owner_login|text|removed|
|owner_name|text|removed|
|owner_node_id|text|removed|
|owner_organizations_url|text|removed|
|owner_owned_private_repos|bigint|removed|
|owner_permissions|jsonb|removed|
|owner_plan_collaborators|bigint|removed|
|owner_plan_filled_seats|bigint|removed|
|owner_plan_name|text|removed|
|owner_plan_private_repos|bigint|removed|
|owner_plan_seats|bigint|removed|
|owner_plan_space|bigint|removed|
|owner_private_gists|bigint|removed|
|owner_public_gists|bigint|removed|
|owner_public_repos|bigint|removed|
|owner_received_events_url|text|removed|
|owner_repos_url|text|removed|
|owner_role_name|text|removed|
|owner_site_admin|boolean|removed|
|owner_starred_url|text|removed|
|owner_subscriptions_url|text|removed|
|owner_suspended_at_time|timestamp without time zone|removed|
|owner_text_matches|jsonb|removed|
|owner_total_private_repos|bigint|removed|
|owner_twitter_username|text|removed|
|owner_two_factor_authentication|boolean|removed|
|owner_type|text|removed|
|owner_updated_at_time|timestamp without time zone|removed|
|owner_url|text|removed|
|parent|jsonb|updated|Type changed from bigint to jsonb
|pushed_at|timestamp without time zone|added|
|pushed_at_time|timestamp without time zone|removed|
|s_v_n_url|text|removed|
|security_and_analysis|jsonb|added|
|security_and_analysis_advanced_security_status|text|removed|
|security_and_analysis_secret_scanning_status|text|removed|
|source|jsonb|updated|Type changed from bigint to jsonb
|svn_url|text|added|
|team_cq_id|uuid|removed|
|template_repository|jsonb|updated|Type changed from bigint to jsonb
|updated_at|timestamp without time zone|added|
|updated_at_time|timestamp without time zone|removed|
|use_squash_p_r_title_as_default|boolean|removed|
|use_squash_pr_title_as_default|boolean|added|

## github_teams

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|ldap_dn|text|added|
|ldapdn|text|removed|
|organization|jsonb|added|
|parent|jsonb|updated|Type changed from bigint to jsonb

## heroku_account_features

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|

## heroku_add_on_attachments

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|

## heroku_add_on_configs

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|

## heroku_add_on_region_capabilities

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|

## heroku_add_on_services

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|

## heroku_add_on_webhook_deliveries

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|created_at|timestamp without time zone|removed|
|event|jsonb|removed|
|id|text|removed|
|last_attempt|jsonb|removed|
|next_attempt_at|timestamp without time zone|removed|
|num_attempts|integer|removed|
|status|text|removed|
|updated_at|timestamp without time zone|removed|
|webhook|jsonb|removed|

## heroku_add_on_webhook_events

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|created_at|timestamp without time zone|removed|
|id|text|removed|
|include|text|removed|
|payload|jsonb|removed|
|updated_at|timestamp without time zone|removed|

## heroku_add_on_webhooks

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|addon|jsonb|removed|

## heroku_add_ons

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|

## heroku_app_features

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|

## heroku_app_transfers

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|

## heroku_app_webhook_deliveries

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|num_attempts|bigint|updated|Type changed from integer to bigint

## heroku_app_webhook_events

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|

## heroku_app_webhooks

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|app|jsonb|removed|
|created_at|timestamp without time zone|removed|
|id|text|removed|
|include|text[]|removed|
|level|text|removed|
|updated_at|timestamp without time zone|removed|
|url|text|removed|

## heroku_apps

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|repo_size|bigint|updated|Type changed from integer to bigint
|slug_size|bigint|updated|Type changed from integer to bigint

## heroku_buildpack_installations

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|ordinal|bigint|updated|Type changed from integer to bigint

## heroku_builds

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|

## heroku_collaborators

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|

## heroku_credits

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|amount|real|updated|Type changed from float to real
|balance|real|updated|Type changed from float to real

## heroku_domains

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|

## heroku_dyno_sizes

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|compute|bigint|updated|Type changed from integer to bigint
|dyno_units|bigint|updated|Type changed from integer to bigint
|memory|real|updated|Type changed from float to real

## heroku_dynos

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|

## heroku_enterprise_account_members

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|

## heroku_enterprise_accounts

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|

## heroku_formations

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|quantity|bigint|updated|Type changed from integer to bigint

## heroku_inbound_rulesets

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|

## heroku_invoices

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|charges_total|real|updated|Type changed from float to real
|credits_total|real|updated|Type changed from float to real
|number|bigint|updated|Type changed from integer to bigint
|state|bigint|updated|Type changed from integer to bigint
|total|real|updated|Type changed from float to real

## heroku_keys

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|

## heroku_log_drains

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|

## heroku_oauth_authorizations

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|

## heroku_oauth_clients

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|

## heroku_outbound_rulesets

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|

## heroku_peerings

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|

## heroku_permission_entities

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|

## heroku_pipeline_builds

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|app|jsonb|removed|
|buildpacks|jsonb|removed|
|created_at|timestamp without time zone|removed|
|id|text|removed|
|output_stream_url|text|removed|
|release|jsonb|removed|
|slug|jsonb|removed|
|source_blob|jsonb|removed|
|stack|text|removed|
|status|text|removed|
|updated_at|timestamp without time zone|removed|
|user|jsonb|removed|

## heroku_pipeline_couplings

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|

## heroku_pipeline_deployments

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|addon_plan_names|text[]|removed|
|app|jsonb|removed|
|created_at|timestamp without time zone|removed|
|current|boolean|removed|
|description|text|removed|
|id|text|removed|
|output_stream_url|text|removed|
|slug|jsonb|removed|
|status|text|removed|
|updated_at|timestamp without time zone|removed|
|user|jsonb|removed|
|version|integer|removed|

## heroku_pipeline_releases

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|addon_plan_names|text[]|removed|
|app|jsonb|removed|
|created_at|timestamp without time zone|removed|
|current|boolean|removed|
|description|text|removed|
|id|text|removed|
|output_stream_url|text|removed|
|slug|jsonb|removed|
|status|text|removed|
|updated_at|timestamp without time zone|removed|
|user|jsonb|removed|
|version|integer|removed|

## heroku_pipelines

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|

## heroku_regions

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|

## heroku_releases

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|version|bigint|updated|Type changed from integer to bigint

## heroku_review_apps

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|pr_number|bigint|updated|Type changed from integer to bigint

## heroku_space_app_accesses

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|

## heroku_spaces

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|

## heroku_stacks

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|

## heroku_team_app_permissions

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|

## heroku_team_features

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|

## heroku_team_invitations

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|

## heroku_team_invoices

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|addons_total|bigint|updated|Type changed from integer to bigint
|charges_total|bigint|updated|Type changed from integer to bigint
|credits_total|bigint|updated|Type changed from integer to bigint
|database_total|bigint|updated|Type changed from integer to bigint
|dyno_units|real|updated|Type changed from float to real
|number|bigint|updated|Type changed from integer to bigint
|platform_total|bigint|updated|Type changed from integer to bigint
|state|bigint|updated|Type changed from integer to bigint
|total|bigint|updated|Type changed from integer to bigint
|weighted_dyno_hours|real|updated|Type changed from float to real

## heroku_team_members

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|

## heroku_team_spaces

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|cidr|text|removed|
|created_at|timestamp without time zone|removed|
|data_cidr|text|removed|
|id|text|removed|
|name|text|removed|
|organization|jsonb|removed|
|region|jsonb|removed|
|shield|boolean|removed|
|state|text|removed|
|team|jsonb|removed|
|updated_at|timestamp without time zone|removed|

## heroku_teams

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|membership_limit|real|updated|Type changed from float to real

## heroku_vpn_connections

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
|ike_version|bigint|updated|Type changed from integer to bigint

## k8s_apps_daemon_set_selector_match_expressions
This table was removed.


## k8s_apps_daemon_set_status_conditions
This table was removed.


## k8s_apps_daemon_sets
This table was removed.


## k8s_apps_deployment_selector_match_expressions
This table was removed.


## k8s_apps_deployment_status_conditions
This table was removed.


## k8s_apps_deployments
This table was removed.


## k8s_apps_replica_set_selector_match_expressions
This table was removed.


## k8s_apps_replica_set_status_conditions
This table was removed.


## k8s_apps_replica_sets
This table was removed.


## k8s_apps_stateful_set_selector_match_expressions
This table was removed.


## k8s_apps_stateful_set_status_conditions
This table was removed.


## k8s_apps_stateful_sets
This table was removed.


## k8s_batch_cron_jobs
This table was removed.


## k8s_batch_job_selector_match_expressions
This table was removed.


## k8s_batch_job_status_conditions
This table was removed.


## k8s_batch_jobs
This table was removed.


## k8s_core_endpoint_subset_addresses
This table was removed.


## k8s_core_endpoint_subset_not_ready_addresses
This table was removed.


## k8s_core_endpoint_subset_ports
This table was removed.


## k8s_core_endpoint_subsets
This table was removed.


## k8s_core_endpoints
This table was removed.


## k8s_core_limit_range_limits
This table was removed.


## k8s_core_limit_ranges
This table was removed.


## k8s_core_namespaces
This table was removed.


## k8s_core_node_images
This table was removed.


## k8s_core_node_volumes_attached
This table was removed.


## k8s_core_nodes
This table was removed.


## k8s_core_pod_container_envs
This table was removed.


## k8s_core_pod_container_ports
This table was removed.


## k8s_core_pod_container_statuses
This table was removed.


## k8s_core_pod_container_volume_devices
This table was removed.


## k8s_core_pod_container_volume_mounts
This table was removed.


## k8s_core_pod_containers
This table was removed.


## k8s_core_pod_ephemeral_container_envs
This table was removed.


## k8s_core_pod_ephemeral_container_ports
This table was removed.


## k8s_core_pod_ephemeral_container_statuses
This table was removed.


## k8s_core_pod_ephemeral_container_volume_devices
This table was removed.


## k8s_core_pod_ephemeral_container_volume_mounts
This table was removed.


## k8s_core_pod_ephemeral_containers
This table was removed.


## k8s_core_pod_init_container_envs
This table was removed.


## k8s_core_pod_init_container_ports
This table was removed.


## k8s_core_pod_init_container_statuses
This table was removed.


## k8s_core_pod_init_container_volume_devices
This table was removed.


## k8s_core_pod_init_container_volume_mounts
This table was removed.


## k8s_core_pod_init_containers
This table was removed.


## k8s_core_pod_volumes
This table was removed.


## k8s_core_pods
This table was removed.


## k8s_core_resource_quota_scope_selector_match_expressions
This table was removed.


## k8s_core_resource_quotas
This table was removed.


## k8s_core_service_account_secrets
This table was removed.


## k8s_core_service_accounts
This table was removed.


## k8s_core_service_conditions
This table was removed.


## k8s_core_service_load_balancer_ingress_ports
This table was removed.


## k8s_core_service_load_balancer_ingresses
This table was removed.


## k8s_core_service_ports
This table was removed.


## k8s_core_services
This table was removed.


## k8s_meta_owner_references
This table was removed.


## k8s_networking_network_policies
This table was removed.


## k8s_networking_network_policy_egress
This table was removed.


## k8s_networking_network_policy_egress_ports
This table was removed.


## k8s_networking_network_policy_egress_to
This table was removed.


## k8s_networking_network_policy_ingress
This table was removed.


## k8s_networking_network_policy_ingress_from
This table was removed.


## k8s_networking_network_policy_ingress_ports
This table was removed.


## k8s_networking_network_policy_pod_selector_match_expressions
This table was removed.


## k8s_rbac_role_binding_subjects
This table was removed.


## k8s_rbac_role_bindings
This table was removed.


## k8s_rbac_role_rules
This table was removed.


## k8s_rbac_roles
This table was removed.


## okta_users

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|

## tf_data

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|

## tf_resource_instances

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|

## tf_resources

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_fetch_time|timestamp without time zone|added|
|_cq_id|uuid|added|
