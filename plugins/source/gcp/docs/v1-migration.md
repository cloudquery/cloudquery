# Schema Changes from v0 to v1
This guide summarizes schema changes from CloudQuery v0 to v1. It is automatically generated and
not guaranteed to be complete, but we hope it helps as a starting point and reference when migrating to v1.

Last updated 2022-10-06.

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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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

## gcp_billing_billing_accounts
This table was newly added.

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|project_id|text|added|
|name|text|added|
|open|boolean|added|
|display_name|text|added|
|master_billing_account|text|added|

## gcp_billing_services
This table was newly added.

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|project_id|text|added|
|name|text|added|
|service_id|text|added|
|display_name|text|added|
|business_entity_name|text|added|

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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|id|bigint|updated|Type changed from text to bigint

## gcp_compute_autoscaler_custom_metric_utilizations
Moved to JSON column on [gcp_compute_autoscalers](#gcp_compute_autoscalers)


## gcp_compute_autoscalers

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|deprecated|jsonb|updated|Type changed from text to jsonb
|deprecated_deleted|text|removed|
|deprecated_obsolete|text|removed|
|deprecated_replacement|text|removed|
|deprecated_state|text|removed|
|id|bigint|updated|Type changed from text to bigint

## gcp_compute_disks

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|allowed|jsonb|added|
|denied|jsonb|added|
|id|bigint|updated|Type changed from text to bigint
|log_config|jsonb|added|
|log_config_enable|boolean|removed|
|log_config_metadata|text|removed|

## gcp_compute_forwarding_rules

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|id|bigint|updated|Type changed from text to bigint
|metadata_filters|jsonb|added|
|no_automate_dns_zone|boolean|added|
|psc_connection_status|text|added|
|service_directory_registrations|jsonb|added|

## gcp_compute_images

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|circuit_infos|jsonb|added|
|expected_outages|jsonb|added|
|id|bigint|updated|Type changed from text to bigint
|satisfies_pzs|boolean|added|

## gcp_compute_network_peerings
Moved to JSON column on [gcp_compute_networks](#gcp_compute_networks)


## gcp_compute_networks

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|creation_timestamp|text|updated|Type changed from timestamp without time zone to text
|id|bigint|updated|Type changed from text to bigint

## gcp_compute_target_https_proxies
This table was removed.


## gcp_compute_target_ssl_proxies

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|id|bigint|updated|Type changed from text to bigint
|stack_type|text|added|
|vpn_interfaces|jsonb|added|

## gcp_container_clusters
This table was newly added.

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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

## gcp_dns_managed_zone_dnssec_config_default_key_specs
Moved to JSON column on [gcp_dns_managed_zones](#gcp_dns_managed_zones)


## gcp_dns_managed_zone_forwarding_config_target_name_servers
Moved to JSON column on [gcp_dns_managed_zones](#gcp_dns_managed_zones)


## gcp_dns_managed_zone_private_visibility_config_networks
Moved to JSON column on [gcp_dns_managed_zones](#gcp_dns_managed_zones)


## gcp_dns_managed_zones

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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
|create_time|timestamp without time zone|updated|Type changed from text to timestamp without time zone
|custom_dns_ds_records|jsonb|removed|
|custom_dns_name_servers|text[]|removed|
|dns_settings|jsonb|added|
|expire_time|timestamp without time zone|updated|Type changed from text to timestamp without time zone
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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|project_id|text|added|
|name|text|added|
|description|text|added|
|status|bigint|added|
|entry_point|text|added|
|runtime|text|added|
|timeout|jsonb|added|
|available_memory_mb|bigint|added|
|service_account_email|text|added|
|update_time|timestamp without time zone|added|
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

## gcp_iam_roles

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|

## gcp_iam_service_account_keys

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|disabled|boolean|added|
|private_key_data|text|added|
|private_key_type|text|added|
|project_id|text|added|
|public_key_data|text|added|
|service_account_cq_id|uuid|removed|
|service_account_unique_id|text|added|
|valid_after_time|text|updated|Type changed from timestamp without time zone to text
|valid_before_time|text|updated|Type changed from timestamp without time zone to text

## gcp_iam_service_accounts

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|etag|text|added|
|id|text|removed|
|unique_id|text|added|

## gcp_kms_crypto_keys
This table was newly added.

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|project_id|text|added|
|create_time|text|added|
|crypto_key_backend|text|added|
|destroy_scheduled_duration|text|added|
|import_only|boolean|added|
|labels|jsonb|added|
|name|text|added|
|next_rotation_time|text|added|
|primary|jsonb|added|
|purpose|text|added|
|rotation_period|text|added|
|version_template|jsonb|added|

## gcp_kms_keyring_crypto_keys
Moved to JSON column on [gcp_kms_keyrings](#gcp_kms_keyrings)


## gcp_kms_keyrings

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|bucket_options|jsonb|added|
|create_time|timestamp without time zone|updated|Type changed from text to timestamp without time zone
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
|update_time|timestamp without time zone|updated|Type changed from text to timestamp without time zone
|version|bigint|updated|Type changed from text to bigint

## gcp_logging_sink_exclusions
Moved to JSON column on [gcp_logging_sinks](#gcp_logging_sinks)


## gcp_logging_sinks

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|bigquery_options_use_partitioned_tables|boolean|removed|
|bigquery_options_uses_timestamp_column_partitioning|boolean|removed|
|create_time|timestamp without time zone|updated|Type changed from text to timestamp without time zone
|exclusions|jsonb|added|
|output_version_format|bigint|updated|Type changed from text to bigint
|update_time|timestamp without time zone|updated|Type changed from text to timestamp without time zone

## gcp_memorystore_redis_instance_server_ca_certs
This table was removed.


## gcp_memorystore_redis_instances
This table was removed.


## gcp_monitoring_alert_policies

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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
|create_time|timestamp without time zone|added|
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

## gcp_resource_manager_folders
Renamed to [gcp_resourcemanager_folders](#gcp_resourcemanager_folders)


## gcp_resource_manager_projects
Renamed to [gcp_resourcemanager_projects](#gcp_resourcemanager_projects)


## gcp_resourcemanager_folders
Renamed from [gcp_resource_manager_folders](gcp_resource_manager_folders)


## gcp_resourcemanager_project_policies
This table was newly added.

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|project_id|text|added|
|audit_configs|jsonb|added|
|bindings|jsonb|added|
|etag|text|added|
|version|bigint|added|

## gcp_resourcemanager_projects
Renamed from [gcp_resource_manager_projects](gcp_resource_manager_projects)


## gcp_run_services
This table was newly added.

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|project_id|text|added|
|name|text|added|
|description|text|added|
|uid|text|added|
|generation|bigint|added|
|labels|jsonb|added|
|annotations|jsonb|added|
|create_time|timestamp without time zone|added|
|update_time|timestamp without time zone|added|
|delete_time|timestamp without time zone|added|
|expire_time|timestamp without time zone|added|
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

## gcp_secretmanager_secrets
This table was newly added.

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|project_id|text|added|
|name|text|added|
|replication|jsonb|added|
|create_time|timestamp without time zone|added|
|labels|jsonb|added|
|topics|jsonb|added|
|etag|text|added|
|rotation|jsonb|added|
|version_aliases|jsonb|added|

## gcp_security_secret_user_managed_replicas
This table was removed.


## gcp_security_secrets
This table was removed.


## gcp_serviceusage_services
This table was newly added.

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|project_id|text|added|
|name|text|added|
|parent|text|added|
|config|jsonb|added|
|state|bigint|added|

## gcp_sql_instance_ip_addresses
Moved to JSON column on [gcp_sql_instances](#gcp_sql_instances)


## gcp_sql_instance_settings_deny_maintenance_periods
Moved to JSON column on [gcp_sql_instances](#gcp_sql_instances)


## gcp_sql_instance_settings_ip_config_authorized_networks
Moved to JSON column on [gcp_sql_instances](#gcp_sql_instances)


## gcp_sql_instances

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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


## gcp_storage_bucket_policies
This table was newly added.

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|project_id|text|added|
|bucket_name|text|added|
|bindings|jsonb|added|

## gcp_storage_buckets

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
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

