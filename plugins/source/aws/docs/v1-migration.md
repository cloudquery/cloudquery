# Schema Changes from v0 to v1
This guide summarizes schema changes from CloudQuery v0 to v1. It is automatically generated and
not guaranteed to be complete, but we hope it helps as a starting point and reference when migrating to v1.

Last updated 2022-10-06.

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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|account_id|text|added|
|region|text|added|
|analyzer_arn|text|added|
|created_at|timestamp without time zone|added|
|filter|jsonb|added|
|rule_name|text|added|
|updated_at|timestamp without time zone|added|

## aws_accessanalyzer_analyzer_findings
This table was newly added.

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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

## aws_accessanalyzer_analyzers
Renamed from [aws_access_analyzer_analyzers](aws_access_analyzer_analyzers)


## aws_accounts
Renamed to [aws_iam_accounts](#aws_iam_accounts)


## aws_acm_certificates

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|

## aws_apigateway_client_certificates

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|client_certificate_id|text|added|
|id|text|removed|

## aws_apigateway_domain_name_base_path_mappings

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|account_id|text|added|
|domain_name|text|removed|
|domain_name_arn|text|added|
|domain_name_cq_id|uuid|removed|
|region|text|added|

## aws_apigateway_domain_names

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|account_id|text|added|
|region|text|added|
|rest_api_arn|text|added|
|rest_api_cq_id|uuid|removed|
|rest_api_id|text|removed|

## aws_apigateway_rest_api_documentation_parts

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|account_id|text|added|
|region|text|added|
|rest_api_arn|text|added|
|rest_api_cq_id|uuid|removed|
|rest_api_id|text|removed|

## aws_apigateway_rest_api_gateway_responses

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|account_id|text|added|
|region|text|added|
|rest_api_arn|text|added|
|rest_api_cq_id|uuid|removed|
|rest_api_id|text|removed|

## aws_apigateway_rest_api_models

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|account_id|text|added|
|region|text|added|
|rest_api_arn|text|added|
|rest_api_cq_id|uuid|removed|
|rest_api_id|text|removed|

## aws_apigateway_rest_api_request_validators

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|account_id|text|added|
|region|text|added|
|rest_api_arn|text|added|
|rest_api_cq_id|uuid|removed|
|rest_api_id|text|removed|

## aws_apigateway_rest_api_resources

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|account_id|text|added|
|region|text|added|
|rest_api_arn|text|added|
|rest_api_cq_id|uuid|removed|
|rest_api_id|text|removed|

## aws_apigateway_rest_api_stages

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|endpoint_configuration|jsonb|added|
|endpoint_configuration_types|text[]|removed|
|endpoint_configuration_vpc_endpoint_ids|text[]|removed|

## aws_apigateway_usage_plan_api_stages
Moved to JSON column on [aws_apigateway_usage_plans](#aws_apigateway_usage_plans)


## aws_apigateway_usage_plan_keys

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|account_id|text|added|
|region|text|added|
|usage_plan_arn|text|added|
|usage_plan_cq_id|uuid|removed|
|usage_plan_id|text|removed|

## aws_apigateway_usage_plans

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|

## aws_apigatewayv2_api_authorizers

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|account_id|text|added|
|api_arn|text|added|
|api_cq_id|uuid|removed|
|region|text|added|

## aws_apigatewayv2_api_integration_responses

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|account_id|text|added|
|api_integration_arn|text|added|
|api_integration_cq_id|uuid|removed|
|region|text|added|

## aws_apigatewayv2_api_integrations

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|account_id|text|added|
|api_arn|text|added|
|api_cq_id|uuid|removed|
|region|text|added|

## aws_apigatewayv2_api_route_responses

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|account_id|text|added|
|api_route_arn|text|added|
|api_route_cq_id|uuid|removed|
|region|text|added|

## aws_apigatewayv2_api_routes

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|account_id|text|added|
|api_arn|text|added|
|api_cq_id|uuid|removed|
|region|text|added|

## aws_apigatewayv2_api_stages

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|account_id|text|added|
|domain_name_arn|text|added|
|domain_name_cq_id|uuid|removed|
|region|text|added|

## aws_apigatewayv2_domain_names

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|domain_name_configurations|jsonb|added|
|mutual_tls_authentication|jsonb|added|
|mutual_tls_authentication_truststore_uri|text|removed|
|mutual_tls_authentication_truststore_version|text|removed|
|mutual_tls_authentication_truststore_warnings|text[]|removed|

## aws_apigatewayv2_vpc_links

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|id|text|removed|
|vpc_link_id|text|added|

## aws_applicationautoscaling_policies

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|account_id|text|added|
|data_catalog_arn|text|added|
|data_catalog_cq_id|uuid|removed|
|region|text|added|

## aws_athena_data_catalogs

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|

## aws_athena_work_group_named_queries

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|account_id|text|added|
|region|text|added|
|work_group_arn|text|added|
|work_group_cq_id|uuid|removed|

## aws_athena_work_group_prepared_statements

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|account_id|text|added|
|region|text|added|
|work_group_arn|text|added|
|work_group_cq_id|uuid|removed|

## aws_athena_work_group_query_executions

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|account_id|text|added|
|global_timeout|bigint|updated|Type changed from integer to bigint
|group_arn|text|added|
|group_cq_id|uuid|removed|
|heartbeat_timeout|bigint|updated|Type changed from integer to bigint
|region|text|added|

## aws_autoscaling_group_scaling_policies

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|account_id|text|added|
|cooldown|bigint|updated|Type changed from integer to bigint
|estimated_instance_warmup|bigint|updated|Type changed from integer to bigint
|group_arn|text|added|
|group_cq_id|uuid|removed|
|min_adjustment_magnitude|bigint|updated|Type changed from integer to bigint
|min_adjustment_step|bigint|updated|Type changed from integer to bigint
|name|text|removed|
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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|desired_capacity|bigint|updated|Type changed from integer to bigint
|max_size|bigint|updated|Type changed from integer to bigint
|min_size|bigint|updated|Type changed from integer to bigint
|name|text|removed|
|scheduled_action_name|text|added|

## aws_backup_global_settings

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|result_metadata|jsonb|added|

## aws_backup_plan_rules
Moved to JSON column on [aws_backup_plans](#aws_backup_plans)


## aws_backup_plan_selections

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|backup_plan|jsonb|added|
|backup_plan_id|text|added|
|deletion_date|timestamp without time zone|added|
|id|text|removed|
|name|text|removed|
|result_metadata|jsonb|added|

## aws_backup_region_settings

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|result_metadata|jsonb|added|

## aws_backup_vault_recovery_points

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|account_id|text|added|
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
|region|text|added|
|vault_arn|text|added|
|vault_cq_id|uuid|removed|

## aws_backup_vaults

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|account_id|text|added|
|data_resources|jsonb|added|
|region|text|added|
|trail_cq_id|uuid|removed|

## aws_cloudtrail_trails

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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

## aws_cloudwatch_alarm_metrics
Moved to JSON column on [aws_cloudwatch_alarms](#aws_cloudwatch_alarms)


## aws_cloudwatch_alarms

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|

## aws_cloudwatchlogs_metric_filters
This table was newly added.

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|account_id|text|added|
|region|text|added|
|arn|text|added|
|creation_time|bigint|added|
|filter_name|text|added|
|filter_pattern|text|added|
|log_group_name|text|added|
|metric_transformations|jsonb|added|

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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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
|updated|timestamp without time zone|removed|
|version|bigint|removed|

## aws_codepipeline_webhook_filters
Moved to JSON column on [aws_codepipeline_webhooks](#aws_codepipeline_webhooks)


## aws_codepipeline_webhooks

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|cognito_identity_providers|jsonb|added|
|open_id_connect_provider_ar_ns|text[]|added|
|open_id_connect_provider_arns|text[]|removed|
|result_metadata|jsonb|added|
|saml_provider_ar_ns|text[]|added|
|saml_provider_arns|text[]|removed|

## aws_cognito_user_pool_identity_providers

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|user_pool_arn|text|added|
|user_pool_cq_id|uuid|removed|

## aws_cognito_user_pool_schema_attributes
Moved to JSON column on [aws_cognito_user_pools](#aws_cognito_user_pools)


## aws_cognito_user_pools

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|recording_group|jsonb|added|
|recording_group_all_supported|boolean|removed|
|recording_group_include_global_resource_types|boolean|removed|
|recording_group_resource_types|text[]|removed|

## aws_config_conformance_pack_rule_compliances

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|conformance_pack_input_parameters|jsonb|removed|
|template_ssm_document_details|jsonb|added|

## aws_dax_cluster_nodes
Moved to JSON column on [aws_dax_clusters](#aws_dax_clusters)


## aws_dax_clusters

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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
|total_nodes|bigint|updated|Type changed from integer to bigint

## aws_directconnect_connection_mac_sec_keys
Moved to JSON column on [aws_directconnect_connections](#aws_directconnect_connections)


## aws_directconnect_connections

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|aws_device|text|added|
|aws_logical_device_id|text|added|
|connection_name|text|added|
|mac_sec_keys|jsonb|added|
|name|text|removed|
|vlan|bigint|updated|Type changed from integer to bigint

## aws_directconnect_gateway_associations

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|account_id|text|added|
|direct_connect_gateway_id|text|added|
|gateway_arn|text|added|
|gateway_cq_id|uuid|removed|
|region|text|added|

## aws_directconnect_gateways

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|state|text|removed|
|virtual_gateway_state|text|added|

## aws_directconnect_virtual_interface_bgp_peers
Moved to JSON column on [aws_directconnect_virtual_interfaces](#aws_directconnect_virtual_interfaces)


## aws_directconnect_virtual_interfaces

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|

## aws_ec2_customer_gateways

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|customer_gateway_id|text|added|
|id|text|removed|

## aws_ec2_ebs_snapshots

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|arn|text|added|
|attribute|jsonb|added|
|create_volume_permissions|jsonb|removed|
|restore_expiry_time|timestamp without time zone|added|
|storage_tier|text|added|
|volume_size|bigint|updated|Type changed from integer to bigint

## aws_ec2_ebs_volume_attachments
Moved to JSON column on [aws_ec2_ebs_volumes](#aws_ec2_ebs_volumes)


## aws_ec2_ebs_volumes

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|attachments|jsonb|added|
|id|text|removed|
|iops|bigint|updated|Type changed from integer to bigint
|size|bigint|updated|Type changed from integer to bigint
|throughput|bigint|updated|Type changed from integer to bigint
|volume_id|text|added|

## aws_ec2_egress_only_internet_gateways

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|egress_only_internet_gateway_id|text|added|
|id|text|removed|

## aws_ec2_eips

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|carrier_ip|text|updated|Type changed from inet to text
|customer_owned_ip|text|updated|Type changed from inet to text
|private_ip_address|text|updated|Type changed from inet to text
|public_ip|text|updated|Type changed from inet to text

## aws_ec2_flow_logs

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|deliver_cross_account_role|text|added|
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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|block_device_mappings|jsonb|added|
|boot_mode|text|added|
|creation_date|text|updated|Type changed from timestamp without time zone to text
|deprecation_time|text|updated|Type changed from timestamp without time zone to text
|id|text|removed|
|image_id|text|added|
|imds_support|text|added|
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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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
|tpm_support|text|added|
|usage_operation|text|added|
|usage_operation_update_time|timestamp without time zone|added|

## aws_ec2_internet_gateway_attachments
Moved to JSON column on [aws_ec2_internet_gateways](#aws_ec2_internet_gateways)


## aws_ec2_internet_gateways

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|attachments|jsonb|added|
|id|text|removed|
|internet_gateway_id|text|added|

## aws_ec2_key_pairs

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|public_key|text|added|

## aws_ec2_nat_gateway_addresses
Moved to JSON column on [aws_ec2_nat_gateways](#aws_ec2_nat_gateways)


## aws_ec2_nat_gateways

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|associations|jsonb|added|
|entries|jsonb|added|
|id|text|removed|
|network_acl_id|text|added|

## aws_ec2_network_interface_private_ip_addresses
Moved to JSON column on [aws_ec2_network_interfaces](#aws_ec2_network_interfaces)


## aws_ec2_network_interfaces

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|

## aws_ec2_route_table_associations
Moved to JSON column on [aws_ec2_route_tables](#aws_ec2_route_tables)


## aws_ec2_route_table_propagating_vgws
Moved to JSON column on [aws_ec2_route_tables](#aws_ec2_route_tables)


## aws_ec2_route_table_routes
Moved to JSON column on [aws_ec2_route_tables](#aws_ec2_route_tables)


## aws_ec2_route_tables

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|group_id|text|added|
|id|text|removed|
|ip_permissions|jsonb|added|
|ip_permissions_egress|jsonb|added|

## aws_ec2_subnet_ipv6_cidr_block_association_sets
Moved to JSON column on [aws_ec2_subnets](#aws_ec2_subnets)


## aws_ec2_subnets

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|accepter_owner_id|text|removed|
|accepter_region|text|removed|
|accepter_tgw_info|jsonb|added|
|accepter_transit_gateway_attachment_id|text|added|
|accepter_transit_gateway_id|text|removed|
|account_id|text|added|
|options|jsonb|added|
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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|account_id|text|added|
|region|text|added|
|transit_gateway_arn|text|added|
|transit_gateway_cq_id|uuid|removed|
|transit_gateway_id|text|added|

## aws_ec2_transit_gateway_vpc_attachments

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|id|text|removed|
|private_dns_names|jsonb|updated|Type changed from text[] to jsonb
|service_id|text|added|
|service_type|jsonb|updated|Type changed from text[] to jsonb
|supported_ip_address_types|text[]|added|

## aws_ec2_vpc_endpoints

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|cidr_block_association_set|jsonb|added|
|id|text|removed|
|ipv6_cidr_block_association_set|jsonb|added|
|vpc_id|text|added|

## aws_ec2_vpn_gateways

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|id|text|removed|
|vpc_attachments|jsonb|added|
|vpn_gateway_id|text|added|

## aws_ecr_repositories

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|encryption_configuration|jsonb|added|
|encryption_configuration_encryption_type|text|removed|
|encryption_configuration_kms_key|text|removed|
|image_scanning_configuration|jsonb|added|
|image_scanning_configuration_scan_on_push|boolean|removed|
|name|text|removed|
|repository_name|text|added|
|repository_uri|text|added|
|uri|text|removed|

## aws_ecr_repository_images

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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
|volumes|jsonb|added|

## aws_efs_filesystems

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|certificate_authority|jsonb|added|
|certificate_authority_data|text|removed|
|connector_config|jsonb|added|
|encryption_config|jsonb|added|
|health|jsonb|added|
|id|text|added|
|identity|jsonb|added|
|identity_oidc_issuer|text|removed|
|kubernetes_network_config|jsonb|added|
|kubernetes_network_config_service_ipv4_cidr|text|removed|
|logging|jsonb|added|
|outpost_config|jsonb|added|
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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|

## aws_elasticache_global_replication_group_global_node_groups
Moved to JSON column on [aws_elasticache_global_replication_groups](#aws_elasticache_global_replication_groups)


## aws_elasticache_global_replication_group_members
Moved to JSON column on [aws_elasticache_global_replication_groups](#aws_elasticache_global_replication_groups)


## aws_elasticache_global_replication_groups

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|global_node_groups|jsonb|added|
|members|jsonb|added|
|region|text|added|

## aws_elasticache_parameter_group_parameters
Moved to JSON column on [aws_elasticache_parameter_groups](#aws_elasticache_parameter_groups)


## aws_elasticache_parameter_groups

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|

## aws_elasticache_replication_group_log_delivery_configurations
Moved to JSON column on [aws_elasticache_replication_groups](#aws_elasticache_replication_groups)


## aws_elasticache_replication_group_node_group_members
Moved to JSON column on [aws_elasticache_replication_groups](#aws_elasticache_replication_groups)


## aws_elasticache_replication_group_node_groups
Moved to JSON column on [aws_elasticache_replication_groups](#aws_elasticache_replication_groups)


## aws_elasticache_replication_groups

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|arn|text|added|
|fixed_price|real|updated|Type changed from float to real
|recurring_charges|jsonb|added|
|usage_price|real|updated|Type changed from float to real

## aws_elasticache_service_updates

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|node_snapshots|jsonb|added|

## aws_elasticache_subnet_group_subnets
Moved to JSON column on [aws_elasticache_subnet_groups](#aws_elasticache_subnet_groups)


## aws_elasticache_subnet_groups

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|subnets|jsonb|added|

## aws_elasticache_user_groups

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|pending_changes|jsonb|added|
|pending_user_ids_to_add|text[]|removed|
|pending_user_ids_to_remove|text[]|removed|

## aws_elasticache_users

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|authentication|jsonb|added|
|authentication_password_count|bigint|removed|
|authentication_type|text|removed|

## aws_elasticbeanstalk_application_versions

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|certificates|jsonb|added|
|default_actions|jsonb|added|
|load_balancer_cq_id|uuid|removed|
|port|bigint|updated|Type changed from integer to bigint

## aws_elbv2_load_balancer_attributes

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|availability_zones|jsonb|added|
|load_balancer_name|text|added|
|name|text|removed|
|state|jsonb|added|
|state_code|text|removed|
|state_reason|text|removed|

## aws_elbv2_target_group_target_health_descriptions

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|block_public_access_configuration|jsonb|added|
|block_public_access_configuration_metadata|jsonb|added|
|block_public_security_group_rules|boolean|removed|
|classification|text|removed|
|configurations|jsonb|removed|
|created_by_arn|text|removed|
|creation_date_time|timestamp without time zone|removed|
|properties|jsonb|removed|
|result_metadata|jsonb|added|

## aws_emr_clusters

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|account_id|text|added|
|event_bus_arn|text|added|
|event_bus_cq_id|uuid|removed|
|region|text|added|

## aws_eventbridge_event_buses

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|

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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|account_id|text|added|
|region|text|added|
|arn|text|added|
|tags|jsonb|added|
|association_id|text|added|
|batch_import_meta_data_on_create|boolean|added|
|creation_time|timestamp without time zone|added|
|data_repository_path|text|added|
|data_repository_subdirectories|text[]|added|
|failure_details|jsonb|added|
|file_cache_id|text|added|
|file_cache_path|text|added|
|file_system_id|text|added|
|file_system_path|text|added|
|imported_file_chunk_size|bigint|added|
|lifecycle|text|added|
|nfs|jsonb|added|
|s3|jsonb|added|

## aws_fsx_data_repository_tasks
This table was newly added.

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|account_id|text|added|
|region|text|added|
|arn|text|added|
|tags|jsonb|added|
|creation_time|timestamp without time zone|added|
|lifecycle|text|added|
|task_id|text|added|
|type|text|added|
|capacity_to_release|bigint|added|
|end_time|timestamp without time zone|added|
|failure_details|jsonb|added|
|file_cache_id|text|added|
|file_system_id|text|added|
|paths|text[]|added|
|report|jsonb|added|
|start_time|timestamp without time zone|added|
|status|jsonb|added|

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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|administrative_actions|jsonb|added|
|lifecycle_transition_reason|jsonb|added|
|lifecycle_transition_reason_message|text|removed|

## aws_fsx_storage_virtual_machines
This table was newly added.

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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

## aws_fsx_storage_vms
This table was removed.


## aws_fsx_volume_ontap_configuration
Moved to JSON column on [aws_fsx_volumes](#aws_fsx_volumes)


## aws_fsx_volume_open_zfs_configuration
Moved to JSON column on [aws_fsx_volumes](#aws_fsx_volumes)


## aws_fsx_volumes

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|target_database|jsonb|added|
|target_database_catalog_id|text|removed|
|target_database_name|text|removed|

## aws_glue_datacatalog_encryption_settings

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|aws_kms_key_id|text|removed|
|connection_password_encryption|jsonb|added|
|encryption_at_rest|jsonb|added|
|encryption_at_rest_catalog_encryption_mode|text|removed|
|encryption_at_rest_sse_aws_kms_key_id|text|removed|
|return_connection_password_encrypted|boolean|removed|

## aws_glue_dev_endpoints

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|endpoint_name|text|added|
|name|text|removed|

## aws_glue_job_runs

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|account_id|text|added|
|dpu_seconds|real|updated|Type changed from float to real
|execution_class|text|added|
|job_arn|text|added|
|job_cq_id|uuid|removed|
|max_capacity|real|updated|Type changed from float to real
|notification_property|jsonb|added|
|notification_property_notify_delay_after|bigint|removed|
|region|text|added|

## aws_glue_jobs

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|command|jsonb|added|
|command_name|text|removed|
|command_python_version|text|removed|
|command_script_location|text|removed|
|connections|jsonb|updated|Type changed from text[] to jsonb
|execution_class|text|added|
|execution_property|jsonb|added|
|execution_property_max_concurrent_runs|bigint|removed|
|max_capacity|real|updated|Type changed from float to real
|notification_property|jsonb|added|
|notification_property_notify_delay_after|bigint|removed|
|source_control_details|jsonb|added|

## aws_glue_ml_transform_input_record_tables
Moved to JSON column on [aws_glue_ml_transforms](#aws_glue_ml_transforms)


## aws_glue_ml_transform_task_runs

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|

## aws_glue_registry_schema_versions

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|account_id|text|added|
|region|text|added|
|registry_cq_id|uuid|removed|
|result_metadata|jsonb|added|

## aws_glue_security_configuration_s3_encryption
Moved to JSON column on [aws_glue_security_configurations](#aws_glue_security_configurations)


## aws_glue_security_configurations

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|actions|jsonb|added|
|event_batching_condition|jsonb|added|
|event_batching_condition_size|bigint|removed|
|event_batching_condition_window|bigint|removed|
|predicate|jsonb|added|
|predicate_logical|text|removed|

## aws_glue_workflows

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|administrator_id|text|added|
|detector_arn|text|added|
|detector_cq_id|uuid|removed|
|invited_at|text|updated|Type changed from timestamp without time zone to text
|region|text|added|
|updated_at|text|updated|Type changed from timestamp without time zone to text

## aws_guardduty_detectors

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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

## aws_iam_group_policies

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|group_arn|text|added|
|group_cq_id|uuid|removed|
|result_metadata|jsonb|added|

## aws_iam_groups

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|group_name|text|added|
|name|text|removed|

## aws_iam_openid_connect_identity_providers

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|result_metadata|jsonb|added|

## aws_iam_password_policies

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|max_password_age|bigint|updated|Type changed from integer to bigint
|minimum_password_length|bigint|updated|Type changed from integer to bigint
|password_reuse_prevention|bigint|updated|Type changed from integer to bigint

## aws_iam_policies

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|account_id|text|added|
|result_metadata|jsonb|added|
|role_arn|text|added|
|role_cq_id|uuid|removed|

## aws_iam_roles

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|permissions_boundary|jsonb|added|
|permissions_boundary_arn|text|removed|
|permissions_boundary_type|text|removed|
|role_last_used|jsonb|added|
|role_last_used_last_used_date|timestamp without time zone|removed|
|role_last_used_region|text|removed|

## aws_iam_saml_identity_providers

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|saml_metadata_document|text|removed|

## aws_iam_server_certificates

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|name|text|removed|
|server_certificate_name|text|added|

## aws_iam_user_access_keys

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|account_id|text|added|
|user_arn|text|added|
|user_cq_id|uuid|removed|
|user_name|text|added|

## aws_iam_user_attached_policies

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|account_id|text|added|
|user_arn|text|added|
|user_cq_id|uuid|removed|

## aws_iam_user_groups

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|account_id|text|added|
|arn|text|added|
|group_arn|text|removed|
|user_arn|text|added|
|user_cq_id|uuid|removed|

## aws_iam_user_policies

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|result_metadata|jsonb|added|
|user_arn|text|added|
|user_cq_id|uuid|removed|

## aws_iam_users

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|aws_account_id|text|added|
|finding_arn|text|removed|
|fix_available|text|added|
|inspector_score|real|updated|Type changed from float to real
|remediation|jsonb|added|
|remediation_recommendation_text|text|removed|
|remediation_recommendation_url|text|removed|
|resources|jsonb|added|

## aws_inspector_findings

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|numeric_severity|real|updated|Type changed from float to real

## aws_iot_billing_groups

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|certificate_id|text|added|
|certificate_mode|text|added|
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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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

## aws_iot_policies

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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

## aws_iot_stream_files
Moved to JSON column on [aws_iot_streams](#aws_iot_streams)


## aws_iot_streams

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|files|jsonb|added|
|id|text|removed|
|stream_id|text|added|
|stream_version|bigint|added|
|version|integer|removed|

## aws_iot_thing_groups

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|name|text|removed|
|thing_name|text|added|
|thing_type_name|text|added|
|type_name|text|removed|

## aws_iot_topic_rule_actions
Moved to JSON column on [aws_iot_topic_rules](#aws_iot_topic_rules)


## aws_iot_topic_rules

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|enhanced_monitoring|jsonb|added|
|stream_arn|text|removed|
|stream_mode_details|jsonb|added|
|stream_mode_details_stream_mode|text|removed|

## aws_kms_keys

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|account_id|text|added|
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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|account_id|text|added|
|amazon_managed_kafka_event_source_config|jsonb|added|
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
|self_managed_kafka_event_source_config|jsonb|added|
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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|account_id|text|added|
|layer_version_arn|text|added|
|layer_version_cq_id|uuid|removed|
|region|text|added|
|result_metadata|jsonb|added|

## aws_lambda_layer_versions

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|account_id|text|added|
|region|text|added|

## aws_lightsail_alarms

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|domain_validation_records|jsonb|added|
|renewal_summary|jsonb|added|
|renewal_summary_reason|text|removed|
|renewal_summary_status|text|removed|
|renewal_summary_updated_at|timestamp without time zone|removed|

## aws_lightsail_container_service_deployments

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|account_id|text|added|
|container_service_arn|text|added|
|container_service_cq_id|uuid|removed|
|region|text|added|

## aws_lightsail_container_services

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|account_id|text|added|
|database_arn|text|added|
|database_cq_id|uuid|removed|
|region|text|added|

## aws_lightsail_database_log_events

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|account_id|text|added|
|database_arn|text|added|
|database_cq_id|uuid|removed|
|region|text|added|

## aws_lightsail_database_parameters

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|availability_zone|text|removed|
|location|jsonb|added|

## aws_lightsail_databases

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|add_ons|jsonb|added|
|location|jsonb|added|
|location_availability_zone|text|removed|
|location_region_name|text|removed|

## aws_lightsail_distributions

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|availability_zone|text|removed|
|from_attached_disks|jsonb|added|
|location|jsonb|added|

## aws_lightsail_instances

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|add_ons|jsonb|added|
|hardware|jsonb|added|
|hardware_cpu_count|bigint|removed|
|hardware_ram_size_in_gb|float|removed|
|location|jsonb|added|
|location_availability_zone|text|removed|
|location_region_name|text|removed|
|metadata_options|jsonb|added|
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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|availability_zone|text|removed|
|instance_health_summary|jsonb|added|
|location|jsonb|added|
|public_ports|bigint[]|updated|Type changed from integer[] to bigint[]
|tls_certificate_summaries|jsonb|added|

## aws_lightsail_static_ips

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|availability_zone|text|removed|
|location|jsonb|added|

## aws_mq_broker_configuration_revisions

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|account_id|text|added|
|broker_configuration_arn|text|added|
|broker_configuration_cq_id|uuid|removed|
|region|text|added|
|result_metadata|jsonb|added|

## aws_mq_broker_configurations

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|broker_arn|text|added|
|broker_cq_id|uuid|removed|
|latest_revision|jsonb|updated|Type changed from integer to jsonb
|latest_revision_created|timestamp without time zone|removed|
|latest_revision_description|text|removed|

## aws_mq_broker_users

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|broker_arn|text|added|
|broker_cq_id|uuid|removed|
|broker_id|text|added|
|result_metadata|jsonb|added|

## aws_mq_brokers

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|

## aws_qldb_ledger_journal_kinesis_streams

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|encryption_description|jsonb|added|
|encryption_status|text|removed|
|inaccessible_kms_key_date_time|timestamp without time zone|removed|
|kms_key_arn|text|removed|
|result_metadata|jsonb|added|

## aws_rds_certificates

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|

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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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

## aws_rds_cluster_parameter_groups

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|db_cluster_parameter_group_name|text|added|
|db_parameter_group_family|text|added|
|family|text|removed|
|name|text|removed|

## aws_rds_cluster_parameters
Moved to JSON column on [aws_rds_clusters](#aws_rds_clusters)


## aws_rds_cluster_snapshots

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|allocated_storage|bigint|updated|Type changed from integer to bigint
|percent_progress|bigint|updated|Type changed from integer to bigint
|port|bigint|updated|Type changed from integer to bigint
|tag_list|jsonb|added|

## aws_rds_cluster_vpc_security_groups
Moved to JSON column on [aws_rds_clusters](#aws_rds_clusters)


## aws_rds_clusters

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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
|network_type|text|added|
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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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

## aws_rds_db_parameter_groups

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|db_parameter_group_family|text|added|
|db_parameter_group_name|text|added|
|family|text|removed|
|name|text|removed|

## aws_rds_db_parameters
This table was removed.


## aws_rds_db_security_groups

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|db_security_group_description|text|added|
|db_security_group_name|text|added|
|description|text|removed|
|name|text|removed|

## aws_rds_db_snapshots

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|allocated_storage|bigint|updated|Type changed from integer to bigint
|iops|bigint|updated|Type changed from integer to bigint
|original_snapshot_create_time|timestamp without time zone|added|
|percent_progress|bigint|updated|Type changed from integer to bigint
|port|bigint|updated|Type changed from integer to bigint
|snapshot_database_time|timestamp without time zone|added|
|snapshot_target|text|added|
|tag_list|jsonb|added|

## aws_rds_event_subscriptions

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|

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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|activity_stream_engine_native_audit_fields_included|boolean|added|
|activity_stream_kinesis_stream_name|text|added|
|activity_stream_kms_key_id|text|added|
|activity_stream_mode|text|added|
|activity_stream_policy_status|text|added|
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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|account_id|text|added|
|cluster_arn|text|added|
|cluster_cq_id|uuid|removed|
|cluster_parameter_status_list|jsonb|added|
|region|text|added|

## aws_redshift_cluster_parameters

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|account_id|text|added|
|cluster_arn|text|added|
|cluster_parameter_group_cq_id|uuid|removed|
|region|text|added|

## aws_redshift_cluster_security_groups
Moved to JSON column on [aws_redshift_clusters](#aws_redshift_clusters)


## aws_redshift_cluster_vpc_security_groups
Moved to JSON column on [aws_redshift_clusters](#aws_redshift_clusters)


## aws_redshift_clusters

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|aqua_configuration|jsonb|added|
|automated_snapshot_retention_period|bigint|updated|Type changed from integer to bigint
|cluster_identifier|text|added|
|cluster_nodes|jsonb|added|
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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|cust_subscription_id|text|added|
|id|text|removed|

## aws_redshift_snapshot_accounts_with_restore_access
Moved to JSON column on [aws_redshift_snapshots](#aws_redshift_snapshots)


## aws_redshift_snapshots

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|subnets|jsonb|added|

## aws_regions

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|

## aws_resourcegroups_resource_groups

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|account_id|text|added|
|arn|text|added|
|name_servers|text[]|added|
|caller_reference|text|added|
|id|text|added|

## aws_route53_domain_nameservers
Moved to JSON column on [aws_route53_domains](#aws_route53_domains)


## aws_route53_domains

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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
|result_metadata|jsonb|added|
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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|account_id|text|added|
|hosted_zone_arn|text|added|
|hosted_zone_cq_id|uuid|removed|
|hosted_zone_id|text|added|

## aws_route53_hosted_zone_resource_record_sets

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|latest_version|bigint|updated|Type changed from integer to bigint
|traffic_policy_count|bigint|updated|Type changed from integer to bigint

## aws_route53_traffic_policy_versions

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|account_id|text|added|
|traffic_policy_arn|text|added|
|traffic_policy_cq_id|uuid|removed|
|version|bigint|updated|Type changed from integer to bigint

## aws_s3_account_config
Moved to JSON column on [aws_s3_accounts](#aws_s3_accounts)


## aws_s3_accounts
This table was newly added.

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|account_id|text|added|
|block_public_acls|boolean|added|
|block_public_policy|boolean|added|
|ignore_public_acls|boolean|added|
|restrict_public_buckets|boolean|added|
|config_exists|boolean|added|

## aws_s3_bucket_cors_rules

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|account_id|text|added|
|bucket_arn|text|added|
|bucket_cq_id|uuid|removed|
|max_age_seconds|bigint|updated|Type changed from integer to bigint

## aws_s3_bucket_encryption_rules

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|account_id|text|added|
|apply_server_side_encryption_by_default|jsonb|added|
|bucket_arn|text|added|
|bucket_cq_id|uuid|removed|
|kms_master_key_id|text|removed|
|sse_algorithm|text|removed|

## aws_s3_bucket_grants

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|replication_rules|jsonb|added|

## aws_sagemaker_endpoint_configuration_production_variants
Moved to JSON column on [aws_sagemaker_endpoint_configurations](#aws_sagemaker_endpoint_configurations)


## aws_sagemaker_endpoint_configurations

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|async_inference_config|jsonb|added|
|endpoint_config_name|text|added|
|explainer_config|jsonb|added|
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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|containers|jsonb|added|
|model_name|text|added|
|name|text|removed|
|result_metadata|jsonb|added|
|vpc_config|jsonb|added|

## aws_sagemaker_notebook_instances

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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
|warm_pool_status|jsonb|added|

## aws_secretsmanager_secrets

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|rotation_rules|jsonb|added|
|rotation_rules_automatically_after_days|bigint|removed|
|secret_versions_to_stages|jsonb|removed|
|version_ids_to_stages|jsonb|added|

## aws_ses_templates

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|name|text|removed|
|template_name|text|added|

## aws_shield_attack_properties
Moved to JSON column on [aws_shield_attacks](#aws_shield_attacks)


## aws_shield_attack_sub_resources
Moved to JSON column on [aws_shield_attacks](#aws_shield_attacks)


## aws_shield_attacks

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|attack_properties|jsonb|added|
|mitigations|jsonb|updated|Type changed from text[] to jsonb
|sub_resources|jsonb|added|

## aws_shield_protection_groups

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|id|text|removed|
|protection_group_id|text|added|

## aws_shield_protections

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|application_automatic_response_configuration_status|text|removed|
|application_layer_automatic_response_configuration|jsonb|added|
|region|text|removed|

## aws_shield_subscriptions

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|protected_resource_type_limits|jsonb|removed|
|protection_group_limits_arbitrary_pattern_limits_max_members|integer|removed|
|protection_group_limits_max_protection_groups|integer|removed|
|subscription_limits|jsonb|added|
|time_commitment_in_seconds|bigint|updated|Type changed from integer to bigint

## aws_sns_subscriptions

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|redrive_policy|jsonb|updated|Type changed from text to jsonb

## aws_sns_topics

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|

## aws_sqs_queues

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|

## aws_ssm_documents

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|account_ids|text[]|removed|
|account_sharing_info_list|jsonb|removed|
|category|text[]|added|
|category_enum|text[]|added|
|permissions|jsonb|added|

## aws_ssm_instance_compliance_items

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|policies|jsonb|added|

## aws_transfer_server_workflow_details_on_upload
Moved to JSON column on [aws_transfer_servers](#aws_transfer_servers)


## aws_transfer_servers

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|id|text|removed|
|metric_name|text|removed|
|rule_group_id|text|added|

## aws_waf_rule_predicates
Moved to JSON column on [aws_waf_rules](#aws_waf_rules)


## aws_waf_rules

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|id|text|removed|
|metric_name|text|removed|
|rule_id|text|added|

## aws_waf_subscribed_rule_groups

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|

## aws_waf_web_acl_logging_configuration
Moved to JSON column on [aws_waf_web_acls](#aws_waf_web_acls)


## aws_waf_web_acl_rules
Moved to JSON column on [aws_waf_web_acls](#aws_waf_web_acls)


## aws_waf_web_acls

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|id|text|removed|
|match_predicates|jsonb|added|
|rule_id|text|added|

## aws_wafregional_rule_groups

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|id|text|removed|
|rule_group_id|text|added|

## aws_wafregional_rule_predicates
Moved to JSON column on [aws_wafregional_rules](#aws_wafregional_rules)


## aws_wafregional_rules

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|id|text|removed|
|predicates|jsonb|added|
|rule_id|text|added|

## aws_wafregional_web_acl_rules
Moved to JSON column on [aws_wafregional_web_acls](#aws_wafregional_web_acls)


## aws_wafregional_web_acls

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|default_action|jsonb|updated|Type changed from text to jsonb
|id|text|removed|
|rules|jsonb|added|
|web_acl_id|text|added|

## aws_wafv2_ipsets

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|addresses|inet[]|updated|Type changed from cidr[] to inet[]
|scope|text|removed|

## aws_wafv2_managed_rule_groups

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|regular_expression_list|jsonb|updated|Type changed from text[] to jsonb
|scope|text|removed|

## aws_wafv2_rule_groups

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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
|saml_properties|jsonb|added|
|selfservice_permissions|jsonb|added|
|switch_running_mode|text|removed|
|type|text|removed|
|user_enabled_as_local_administrator|boolean|removed|
|workspace_access_properties|jsonb|added|
|workspace_creation_properties|jsonb|added|

## aws_workspaces_workspaces

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|

## aws_xray_groups

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|insights_configuration|jsonb|added|
|insights_enabled|boolean|removed|
|notifications_enabled|boolean|removed|

## aws_xray_sampling_rules

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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
