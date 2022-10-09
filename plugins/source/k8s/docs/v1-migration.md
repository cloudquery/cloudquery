# Schema Changes from v0 to v1
This guide summarizes schema changes from CloudQuery v0 to v1. It is automatically generated and
not guaranteed to be complete, but we hope it helps as a starting point and reference when migrating to v1.

Last updated 2022-10-06.

## k8s_apps_daemon_set_selector_match_expressions
Moved to JSON column on [k8s_apps_daemon_sets](#k8s_apps_daemon_sets)


## k8s_apps_daemon_set_status_conditions
Moved to JSON column on [k8s_apps_daemon_sets](#k8s_apps_daemon_sets)


## k8s_apps_daemon_sets

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|api_version|text|added|
|generate_name|text|removed|
|kind|text|added|
|managed_fields|jsonb|removed|
|min_ready_seconds|integer|removed|
|revision_history_limit|integer|removed|
|selector_match_labels|jsonb|removed|
|self_link|text|removed|
|spec_min_ready_seconds|bigint|added|
|spec_revision_history_limit|bigint|added|
|spec_selector|jsonb|added|
|spec_template|jsonb|added|
|spec_update_strategy|jsonb|added|
|status_collision_count|bigint|updated|Type changed from integer to bigint
|status_conditions|jsonb|added|
|status_current_number_scheduled|bigint|updated|Type changed from integer to bigint
|status_desired_number_scheduled|bigint|updated|Type changed from integer to bigint
|status_number_available|bigint|updated|Type changed from integer to bigint
|status_number_misscheduled|bigint|updated|Type changed from integer to bigint
|status_number_ready|bigint|updated|Type changed from integer to bigint
|status_number_unavailable|bigint|updated|Type changed from integer to bigint
|status_updated_number_scheduled|bigint|updated|Type changed from integer to bigint
|template|jsonb|removed|
|update_strategy_rolling_update_max_surge_int_val|integer|removed|
|update_strategy_rolling_update_max_surge_str_val|text|removed|
|update_strategy_rolling_update_max_surge_type|bigint|removed|
|update_strategy_rolling_update_max_unavailable_int_val|integer|removed|
|update_strategy_rolling_update_max_unavailable_str_val|text|removed|
|update_strategy_rolling_update_max_unavailable_type|bigint|removed|
|update_strategy_type|text|removed|
|zzz_cluster_name|text|removed|

## k8s_apps_deployment_selector_match_expressions
Moved to JSON column on [k8s_apps_deployments](#k8s_apps_deployments)


## k8s_apps_deployment_status_conditions
Moved to JSON column on [k8s_apps_deployments](#k8s_apps_deployments)


## k8s_apps_deployments

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|generate_name|text|removed|
|managed_fields|jsonb|removed|
|min_ready_seconds|integer|removed|
|paused|boolean|removed|
|progress_deadline_seconds|integer|removed|
|replicas|integer|removed|
|revision_history_limit|integer|removed|
|selector_match_labels|jsonb|removed|
|self_link|text|removed|
|spec_min_ready_seconds|bigint|added|
|spec_paused|boolean|added|
|spec_progress_deadline_seconds|bigint|added|
|spec_replicas|bigint|added|
|spec_revision_history_limit|bigint|added|
|spec_selector|jsonb|added|
|spec_strategy|jsonb|added|
|spec_template|jsonb|added|
|status_available_replicas|bigint|updated|Type changed from integer to bigint
|status_collision_count|bigint|updated|Type changed from integer to bigint
|status_conditions|jsonb|added|
|status_ready_replicas|bigint|updated|Type changed from integer to bigint
|status_replicas|bigint|updated|Type changed from integer to bigint
|status_unavailable_replicas|bigint|updated|Type changed from integer to bigint
|status_updated_replicas|bigint|updated|Type changed from integer to bigint
|strategy_rolling_update_max_surge_int_val|integer|removed|
|strategy_rolling_update_max_surge_str_val|text|removed|
|strategy_rolling_update_max_surge_type|bigint|removed|
|strategy_rolling_update_max_unavailable_int_val|integer|removed|
|strategy_rolling_update_max_unavailable_str_val|text|removed|
|strategy_rolling_update_max_unavailable_type|bigint|removed|
|strategy_type|text|removed|
|template|jsonb|removed|
|zzz_cluster_name|text|removed|

## k8s_apps_replica_set_selector_match_expressions
Moved to JSON column on [k8s_apps_replica_sets](#k8s_apps_replica_sets)


## k8s_apps_replica_set_status_conditions
Moved to JSON column on [k8s_apps_replica_sets](#k8s_apps_replica_sets)


## k8s_apps_replica_sets

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|api_version|text|added|
|generate_name|text|removed|
|kind|text|added|
|managed_fields|jsonb|removed|
|min_ready_seconds|integer|removed|
|replicas|integer|removed|
|selector_match_labels|jsonb|removed|
|self_link|text|removed|
|spec_min_ready_seconds|bigint|added|
|spec_replicas|bigint|added|
|spec_selector|jsonb|added|
|spec_template|jsonb|added|
|status_available_replicas|bigint|updated|Type changed from integer to bigint
|status_conditions|jsonb|added|
|status_fully_labeled_replicas|bigint|updated|Type changed from integer to bigint
|status_ready_replicas|bigint|updated|Type changed from integer to bigint
|status_replicas|bigint|updated|Type changed from integer to bigint
|template|jsonb|removed|
|zzz_cluster_name|text|removed|

## k8s_apps_stateful_set_selector_match_expressions
Moved to JSON column on [k8s_apps_stateful_sets](#k8s_apps_stateful_sets)


## k8s_apps_stateful_set_status_conditions
Moved to JSON column on [k8s_apps_stateful_sets](#k8s_apps_stateful_sets)


## k8s_apps_stateful_sets

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|api_version|text|added|
|generate_name|text|removed|
|kind|text|added|
|managed_fields|jsonb|removed|
|min_ready_seconds|integer|removed|
|pod_management_policy|text|removed|
|replicas|integer|removed|
|revision_history_limit|integer|removed|
|selector_match_labels|jsonb|removed|
|self_link|text|removed|
|service_name|text|removed|
|spec_min_ready_seconds|bigint|added|
|spec_persistent_volume_claim_retention_policy|jsonb|added|
|spec_pod_management_policy|text|added|
|spec_replicas|bigint|added|
|spec_revision_history_limit|bigint|added|
|spec_selector|jsonb|added|
|spec_service_name|text|added|
|spec_template|jsonb|added|
|spec_update_strategy|jsonb|added|
|spec_volume_claim_templates|jsonb|added|
|status_available_replicas|bigint|updated|Type changed from integer to bigint
|status_collision_count|bigint|updated|Type changed from integer to bigint
|status_conditions|jsonb|added|
|status_current_replicas|bigint|updated|Type changed from integer to bigint
|status_ready_replicas|bigint|updated|Type changed from integer to bigint
|status_replicas|bigint|updated|Type changed from integer to bigint
|status_updated_replicas|bigint|updated|Type changed from integer to bigint
|template|jsonb|removed|
|update_strategy_rolling_update_partition|integer|removed|
|update_strategy_type|text|removed|
|volume_claim_templates|jsonb|removed|
|zzz_cluster_name|text|removed|

## k8s_batch_cron_jobs

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|api_version|text|added|
|concurrency_policy|text|removed|
|failed_jobs_history_limit|integer|removed|
|generate_name|text|removed|
|job_template|jsonb|removed|
|kind|text|added|
|managed_fields|jsonb|removed|
|schedule|text|removed|
|self_link|text|removed|
|spec_concurrency_policy|text|added|
|spec_failed_jobs_history_limit|bigint|added|
|spec_job_template|jsonb|added|
|spec_schedule|text|added|
|spec_starting_deadline_seconds|bigint|added|
|spec_successful_jobs_history_limit|bigint|added|
|spec_suspend|boolean|added|
|spec_time_zone|text|added|
|starting_deadline_seconds|bigint|removed|
|status|jsonb|removed|
|status_active|jsonb|added|
|status_last_schedule_time|timestamp without time zone|added|
|status_last_successful_time|timestamp without time zone|added|
|successful_jobs_history_limit|integer|removed|
|suspend|boolean|removed|
|zzz_cluster_name|text|removed|

## k8s_batch_job_selector_match_expressions
Moved to JSON column on [k8s_batch_jobs](#k8s_batch_jobs)


## k8s_batch_job_status_conditions
Moved to JSON column on [k8s_batch_jobs](#k8s_batch_jobs)


## k8s_batch_jobs

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|active_deadline_seconds|bigint|removed|
|api_version|text|added|
|backoff_limit|integer|removed|
|completion_mode|text|removed|
|completions|integer|removed|
|generate_name|text|removed|
|kind|text|added|
|managed_fields|jsonb|removed|
|manual_selector|boolean|removed|
|parallelism|integer|removed|
|selector_match_labels|jsonb|removed|
|self_link|text|removed|
|spec_active_deadline_seconds|bigint|added|
|spec_backoff_limit|bigint|added|
|spec_completion_mode|text|added|
|spec_completions|bigint|added|
|spec_manual_selector|boolean|added|
|spec_parallelism|bigint|added|
|spec_selector|jsonb|added|
|spec_suspend|boolean|added|
|spec_template|jsonb|added|
|spec_ttl_seconds_after_finished|bigint|added|
|status_active|bigint|updated|Type changed from integer to bigint
|status_completion_time|timestamp without time zone|added|
|status_conditions|jsonb|added|
|status_failed|bigint|updated|Type changed from integer to bigint
|status_ready|bigint|added|
|status_start_time|timestamp without time zone|added|
|status_succeeded|bigint|updated|Type changed from integer to bigint
|status_uncounted_terminated_pods|jsonb|added|
|status_uncounted_terminated_pods_failed|text[]|removed|
|status_uncounted_terminated_pods_succeeded|text[]|removed|
|suspend|boolean|removed|
|template|jsonb|removed|
|ttl_seconds_after_finished|integer|removed|
|zzz_cluster_name|text|removed|

## k8s_core_endpoint_subset_addresses
Moved to JSON column on [k8s_core_endpoints](#k8s_core_endpoints)


## k8s_core_endpoint_subset_not_ready_addresses
Moved to JSON column on [k8s_core_endpoints](#k8s_core_endpoints)


## k8s_core_endpoint_subset_ports
Moved to JSON column on [k8s_core_endpoints](#k8s_core_endpoints)


## k8s_core_endpoint_subsets
Moved to JSON column on [k8s_core_endpoints](#k8s_core_endpoints)


## k8s_core_endpoints

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|generate_name|text|removed|
|managed_fields|jsonb|removed|
|self_link|text|removed|
|subsets|jsonb|added|
|zzz_cluster_name|text|removed|

## k8s_core_limit_range_limits
Moved to JSON column on [k8s_core_limit_ranges](#k8s_core_limit_ranges)


## k8s_core_limit_ranges

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|generate_name|text|removed|
|managed_fields|jsonb|removed|
|self_link|text|removed|
|spec_limits|jsonb|added|
|zzz_cluster_name|text|removed|

## k8s_core_namespaces

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|conditions|jsonb|removed|
|generate_name|text|removed|
|managed_fields|jsonb|removed|
|phase|text|removed|
|self_link|text|removed|
|status_conditions|jsonb|added|
|status_phase|text|added|
|zzz_cluster_name|text|removed|

## k8s_core_node_images
Moved to JSON column on [k8s_core_nodes](#k8s_core_nodes)


## k8s_core_node_volumes_attached
Moved to JSON column on [k8s_core_nodes](#k8s_core_nodes)


## k8s_core_nodes

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|allocatable|jsonb|removed|
|architecture|text|removed|
|boot_id|text|removed|
|capacity|jsonb|removed|
|conditions|jsonb|removed|
|config|jsonb|removed|
|container_runtime_version|text|removed|
|daemon_endpoints_kubelet_endpoint_port|integer|removed|
|external_ip|inet|removed|
|hostname|text|removed|
|internal_ip|inet|removed|
|kernel_version|text|removed|
|kube_proxy_version|text|removed|
|kubelet_version|text|removed|
|machine_id|text|removed|
|operating_system|text|removed|
|os_image|text|removed|
|phase|text|removed|
|pod_cidr|cidr|removed|
|pod_cidrs|cidr[]|removed|
|provider_id|text|removed|
|spec_config_source|jsonb|added|
|spec_pod_cidr|cidr|added|
|spec_pod_cidrs|cidr[]|added|
|spec_provider_id|text|added|
|spec_taints|jsonb|added|
|spec_unschedulable|boolean|added|
|status_addresses|jsonb|added|
|status_allocatable|jsonb|added|
|status_capacity|jsonb|added|
|status_conditions|jsonb|added|
|status_config|jsonb|added|
|status_daemon_endpoints|jsonb|added|
|status_images|jsonb|added|
|status_node_info|jsonb|added|
|status_phase|text|added|
|status_volumes_attached|jsonb|added|
|status_volumes_in_use|text[]|added|
|system_uuid|text|removed|
|taints|jsonb|removed|
|unschedulable|boolean|removed|
|volumes_in_use|text[]|removed|
|zzz_cluster_name|text|removed|

## k8s_core_pod_container_envs
Moved to JSON column on [k8s_core_pods](#k8s_core_pods)


## k8s_core_pod_container_ports
Moved to JSON column on [k8s_core_pods](#k8s_core_pods)


## k8s_core_pod_container_statuses
Moved to JSON column on [k8s_core_pods](#k8s_core_pods)


## k8s_core_pod_container_volume_devices
Moved to JSON column on [k8s_core_pods](#k8s_core_pods)


## k8s_core_pod_container_volume_mounts
Moved to JSON column on [k8s_core_pods](#k8s_core_pods)


## k8s_core_pod_containers
Moved to JSON column on [k8s_core_pods](#k8s_core_pods)


## k8s_core_pod_ephemeral_container_envs
Moved to JSON column on [k8s_core_pods](#k8s_core_pods)


## k8s_core_pod_ephemeral_container_ports
Moved to JSON column on [k8s_core_pods](#k8s_core_pods)


## k8s_core_pod_ephemeral_container_statuses
Moved to JSON column on [k8s_core_pods](#k8s_core_pods)


## k8s_core_pod_ephemeral_container_volume_devices
Moved to JSON column on [k8s_core_pods](#k8s_core_pods)


## k8s_core_pod_ephemeral_container_volume_mounts
Moved to JSON column on [k8s_core_pods](#k8s_core_pods)


## k8s_core_pod_ephemeral_containers
Moved to JSON column on [k8s_core_pods](#k8s_core_pods)


## k8s_core_pod_init_container_envs
Moved to JSON column on [k8s_core_pods](#k8s_core_pods)


## k8s_core_pod_init_container_ports
Moved to JSON column on [k8s_core_pods](#k8s_core_pods)


## k8s_core_pod_init_container_statuses
Moved to JSON column on [k8s_core_pods](#k8s_core_pods)


## k8s_core_pod_init_container_volume_devices
Moved to JSON column on [k8s_core_pods](#k8s_core_pods)


## k8s_core_pod_init_container_volume_mounts
Moved to JSON column on [k8s_core_pods](#k8s_core_pods)


## k8s_core_pod_init_containers
Moved to JSON column on [k8s_core_pods](#k8s_core_pods)


## k8s_core_pod_volumes
Moved to JSON column on [k8s_core_pods](#k8s_core_pods)


## k8s_core_pods

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|active_deadline_seconds|bigint|removed|
|affinity|jsonb|removed|
|automount_service_account_token|boolean|removed|
|conditions|jsonb|removed|
|dns_config|jsonb|removed|
|dns_policy|text|removed|
|enable_service_links|boolean|removed|
|host_aliases|jsonb|removed|
|host_ip|inet|removed|
|host_ipc|boolean|removed|
|host_network|boolean|removed|
|host_pid|boolean|removed|
|hostname|text|removed|
|image_pull_secrets|jsonb|removed|
|message|text|removed|
|node_name|text|removed|
|node_selector|jsonb|removed|
|nominated_node_name|text|removed|
|overhead|jsonb|removed|
|phase|text|removed|
|pod_ip|inet|removed|
|pod_ips|inet[]|removed|
|preemption_policy|text|removed|
|priority|integer|removed|
|priority_class_name|text|removed|
|qos_class|text|removed|
|readiness_gates|jsonb|removed|
|reason|text|removed|
|restart_policy|text|removed|
|runtime_class_name|text|removed|
|scheduler_name|text|removed|
|security_context|jsonb|removed|
|service_account_name|text|removed|
|set_hostname_as_fqdn|boolean|removed|
|share_process_namespace|boolean|removed|
|spec_active_deadline_seconds|bigint|added|
|spec_affinity|jsonb|added|
|spec_automount_service_account_token|boolean|added|
|spec_containers|jsonb|added|
|spec_dns_config|jsonb|added|
|spec_dns_policy|text|added|
|spec_enable_service_links|boolean|added|
|spec_ephemeral_containers|jsonb|added|
|spec_host_aliases|jsonb|added|
|spec_host_ipc|boolean|added|
|spec_host_network|boolean|added|
|spec_host_pid|boolean|added|
|spec_hostname|text|added|
|spec_image_pull_secrets|jsonb|added|
|spec_init_containers|jsonb|added|
|spec_node_name|text|added|
|spec_node_selector|jsonb|added|
|spec_os|jsonb|added|
|spec_overhead|jsonb|added|
|spec_preemption_policy|text|added|
|spec_priority|bigint|added|
|spec_priority_class_name|text|added|
|spec_readiness_gates|jsonb|added|
|spec_restart_policy|text|added|
|spec_runtime_class_name|text|added|
|spec_scheduler_name|text|added|
|spec_security_context|jsonb|added|
|spec_service_account_name|text|added|
|spec_set_hostname_as_fqdn|boolean|added|
|spec_share_process_namespace|boolean|added|
|spec_subdomain|text|added|
|spec_termination_grace_period_seconds|bigint|added|
|spec_tolerations|jsonb|added|
|spec_topology_spread_constraints|jsonb|added|
|spec_volumes|jsonb|added|
|status_conditions|jsonb|added|
|status_container_statuses|jsonb|added|
|status_ephemeral_container_statuses|jsonb|added|
|status_host_ip|inet|added|
|status_init_container_statuses|jsonb|added|
|status_message|text|added|
|status_nominated_node_name|text|added|
|status_phase|text|added|
|status_pod_ip|inet|added|
|status_pod_ips|inet[]|added|
|status_qos_class|text|added|
|status_reason|text|added|
|status_start_time|timestamp without time zone|added|
|subdomain|text|removed|
|termination_grace_period_seconds|bigint|removed|
|tolerations|jsonb|removed|
|topology_spread_constraints|jsonb|removed|
|zzz_cluster_name|text|removed|

## k8s_core_resource_quota_scope_selector_match_expressions
Moved to JSON column on [k8s_core_resource_quotas](#k8s_core_resource_quotas)


## k8s_core_resource_quotas

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|generate_name|text|removed|
|hard|jsonb|removed|
|managed_fields|jsonb|removed|
|scopes|text[]|removed|
|self_link|text|removed|
|spec_hard|jsonb|added|
|spec_scope_selector|jsonb|added|
|spec_scopes|text[]|added|
|zzz_cluster_name|text|removed|

## k8s_core_service_account_secrets
Moved to JSON column on [k8s_core_services](#k8s_core_services)


## k8s_core_service_accounts

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|generate_name|text|removed|
|image_pull_secrets|jsonb|added|
|managed_fields|jsonb|removed|
|pull_secret_names|text[]|removed|
|secrets|jsonb|added|
|self_link|text|removed|
|zzz_cluster_name|text|removed|

## k8s_core_service_conditions
Moved to JSON column on [k8s_core_services](#k8s_core_services)


## k8s_core_service_load_balancer_ingress_ports
Moved to JSON column on [k8s_core_services](#k8s_core_services)


## k8s_core_service_load_balancer_ingresses
Moved to JSON column on [k8s_core_services](#k8s_core_services)


## k8s_core_service_ports
Moved to JSON column on [k8s_core_services](#k8s_core_services)


## k8s_core_services

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|allocate_load_balancer_node_ports|boolean|removed|
|cluster_ip|inet|removed|
|cluster_ips|inet[]|removed|
|external_ips|inet[]|removed|
|external_name|text|removed|
|external_traffic_policy|text|removed|
|health_check_node_port|integer|removed|
|internal_traffic_policy|text|removed|
|ip_families|text[]|removed|
|ip_family_policy|text|removed|
|load_balancer_class|text|removed|
|load_balancer_ip|text|removed|
|load_balancer_source_ranges|text[]|removed|
|publish_not_ready_addresses|boolean|removed|
|selector|jsonb|removed|
|session_affinity|text|removed|
|session_affinity_config_client_ip_timeout_seconds|integer|removed|
|spec_allocate_load_balancer_node_ports|boolean|added|
|spec_cluster_ip|inet|added|
|spec_cluster_ips|inet[]|added|
|spec_external_ips|inet[]|added|
|spec_external_name|text|added|
|spec_external_traffic_policy|text|added|
|spec_health_check_node_port|bigint|added|
|spec_internal_traffic_policy|text|added|
|spec_ip_families|text[]|added|
|spec_ip_family_policy|text|added|
|spec_load_balancer_class|text|added|
|spec_load_balancer_ip|inet|added|
|spec_load_balancer_source_ranges|text[]|added|
|spec_ports|jsonb|added|
|spec_publish_not_ready_addresses|boolean|added|
|spec_selector|jsonb|added|
|spec_session_affinity|text|added|
|spec_session_affinity_config|jsonb|added|
|spec_type|text|added|
|status_conditions|jsonb|added|
|status_load_balancer|jsonb|added|
|type|text|removed|
|zzz_cluster_name|text|removed|

## k8s_meta_owner_references
This table was removed.


## k8s_networking_network_policies

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|generate_name|text|removed|
|managed_fields|jsonb|removed|
|pod_selector_match_labels|jsonb|removed|
|policy_types|text[]|removed|
|self_link|text|removed|
|spec_egress|jsonb|added|
|spec_ingress|jsonb|added|
|spec_pod_selector|jsonb|added|
|spec_policy_types|text[]|added|
|status_conditions|jsonb|added|
|zzz_cluster_name|text|removed|

## k8s_networking_network_policy_egress
Moved to JSON column on [k8s_networking_network_policies](#k8s_networking_network_policies)


## k8s_networking_network_policy_egress_ports
Moved to JSON column on [k8s_networking_network_policies](#k8s_networking_network_policies)


## k8s_networking_network_policy_egress_to
Moved to JSON column on [k8s_networking_network_policies](#k8s_networking_network_policies)


## k8s_networking_network_policy_ingress
Moved to JSON column on [k8s_networking_network_policies](#k8s_networking_network_policies)


## k8s_networking_network_policy_ingress_from
Moved to JSON column on [k8s_networking_network_policies](#k8s_networking_network_policies)


## k8s_networking_network_policy_ingress_ports
Moved to JSON column on [k8s_networking_network_policies](#k8s_networking_network_policies)


## k8s_networking_network_policy_pod_selector_match_expressions
Moved to JSON column on [k8s_networking_network_policies](#k8s_networking_network_policies)


## k8s_rbac_role_binding_subjects
Moved to JSON column on [k8s_rbac_roles](#k8s_rbac_roles)


## k8s_rbac_role_bindings

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|generate_name|text|removed|
|managed_fields|jsonb|removed|
|role_ref|jsonb|added|
|role_ref_api_group|text|removed|
|role_ref_kind|text|removed|
|role_ref_name|text|removed|
|self_link|text|removed|
|subjects|jsonb|added|
|zzz_cluster_name|text|removed|

## k8s_rbac_role_rules
Moved to JSON column on [k8s_rbac_roles](#k8s_rbac_roles)


## k8s_rbac_roles

| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
|_cq_id|uuid|added|
|_cq_parent_id|uuid|added|
|_cq_source_name|text|added|
|_cq_sync_time|timestamp without time zone|added|
|generate_name|text|removed|
|managed_fields|jsonb|removed|
|rules|jsonb|added|
|self_link|text|removed|
|zzz_cluster_name|text|removed|
