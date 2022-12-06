# Table: k8s_core_pods



The primary key for this table is **uid**.



## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|context|String|
|uid (PK)|String|
|status_host_ip|Inet|
|status_pod_ip|Inet|
|status_pod_ips|InetArray|
|kind|String|
|api_version|String|
|name|String|
|namespace|String|
|resource_version|String|
|generation|Int|
|deletion_grace_period_seconds|Int|
|labels|JSON|
|annotations|JSON|
|owner_references|JSON|
|finalizers|StringArray|
|spec_volumes|JSON|
|spec_init_containers|JSON|
|spec_containers|JSON|
|spec_ephemeral_containers|JSON|
|spec_restart_policy|String|
|spec_termination_grace_period_seconds|Int|
|spec_active_deadline_seconds|Int|
|spec_dns_policy|String|
|spec_node_selector|JSON|
|spec_service_account_name|String|
|spec_automount_service_account_token|Bool|
|spec_node_name|String|
|spec_host_network|Bool|
|spec_host_pid|Bool|
|spec_host_ipc|Bool|
|spec_share_process_namespace|Bool|
|spec_security_context|JSON|
|spec_image_pull_secrets|JSON|
|spec_hostname|String|
|spec_subdomain|String|
|spec_affinity|JSON|
|spec_scheduler_name|String|
|spec_tolerations|JSON|
|spec_host_aliases|JSON|
|spec_priority_class_name|String|
|spec_priority|Int|
|spec_dns_config|JSON|
|spec_readiness_gates|JSON|
|spec_runtime_class_name|String|
|spec_enable_service_links|Bool|
|spec_preemption_policy|String|
|spec_overhead|JSON|
|spec_topology_spread_constraints|JSON|
|spec_set_hostname_as_fqdn|Bool|
|spec_os|JSON|
|spec_host_users|Bool|
|status_phase|String|
|status_conditions|JSON|
|status_message|String|
|status_reason|String|
|status_nominated_node_name|String|
|status_start_time|Timestamp|
|status_init_container_statuses|JSON|
|status_container_statuses|JSON|
|status_qos_class|String|
|status_ephemeral_container_statuses|JSON|