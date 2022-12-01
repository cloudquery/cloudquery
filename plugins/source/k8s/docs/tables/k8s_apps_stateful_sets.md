# Table: k8s_apps_stateful_sets



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
|spec_replicas|Int|
|spec_selector|JSON|
|spec_template|JSON|
|spec_volume_claim_templates|JSON|
|spec_service_name|String|
|spec_pod_management_policy|String|
|spec_update_strategy|JSON|
|spec_revision_history_limit|Int|
|spec_min_ready_seconds|Int|
|spec_persistent_volume_claim_retention_policy|JSON|
|status_observed_generation|Int|
|status_replicas|Int|
|status_ready_replicas|Int|
|status_current_replicas|Int|
|status_updated_replicas|Int|
|status_current_revision|String|
|status_update_revision|String|
|status_collision_count|Int|
|status_conditions|JSON|
|status_available_replicas|Int|