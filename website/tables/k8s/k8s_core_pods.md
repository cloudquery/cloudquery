# Table: k8s_core_pods

This table shows data for Kubernetes (K8s) Core Pods.

The primary key for this table is **uid**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|utf8|
|_cq_sync_time|timestamp[us, tz=UTC]|
|_cq_id|uuid|
|_cq_parent_id|uuid|
|context|utf8|
|status_host_ip|inet|
|status_pod_ip|inet|
|status_pod_ips|list<item: inet, nullable>|
|kind|utf8|
|api_version|utf8|
|name|utf8|
|namespace|utf8|
|uid (PK)|utf8|
|resource_version|utf8|
|generation|int64|
|deletion_grace_period_seconds|int64|
|labels|json|
|annotations|json|
|owner_references|json|
|finalizers|list<item: utf8, nullable>|
|spec_volumes|json|
|spec_init_containers|json|
|spec_containers|json|
|spec_ephemeral_containers|json|
|spec_restart_policy|utf8|
|spec_termination_grace_period_seconds|int64|
|spec_active_deadline_seconds|int64|
|spec_dns_policy|utf8|
|spec_node_selector|json|
|spec_service_account_name|utf8|
|spec_automount_service_account_token|bool|
|spec_node_name|utf8|
|spec_host_network|bool|
|spec_host_pid|bool|
|spec_host_ipc|bool|
|spec_share_process_namespace|bool|
|spec_security_context|json|
|spec_image_pull_secrets|json|
|spec_hostname|utf8|
|spec_subdomain|utf8|
|spec_affinity|json|
|spec_scheduler_name|utf8|
|spec_tolerations|json|
|spec_host_aliases|json|
|spec_priority_class_name|utf8|
|spec_priority|int64|
|spec_dns_config|json|
|spec_readiness_gates|json|
|spec_runtime_class_name|utf8|
|spec_enable_service_links|bool|
|spec_preemption_policy|utf8|
|spec_overhead|json|
|spec_topology_spread_constraints|json|
|spec_set_hostname_as_fqdn|bool|
|spec_os|json|
|spec_host_users|bool|
|spec_scheduling_gates|json|
|spec_resource_claims|json|
|status_phase|utf8|
|status_conditions|json|
|status_message|utf8|
|status_reason|utf8|
|status_nominated_node_name|utf8|
|status_start_time|timestamp[us, tz=UTC]|
|status_init_container_statuses|json|
|status_container_statuses|json|
|status_qos_class|utf8|
|status_ephemeral_container_statuses|json|