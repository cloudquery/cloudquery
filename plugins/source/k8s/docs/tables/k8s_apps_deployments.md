# Table: k8s_apps_deployments



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
|spec_strategy|JSON|
|spec_min_ready_seconds|Int|
|spec_revision_history_limit|Int|
|spec_paused|Bool|
|spec_progress_deadline_seconds|Int|
|status_observed_generation|Int|
|status_replicas|Int|
|status_updated_replicas|Int|
|status_ready_replicas|Int|
|status_available_replicas|Int|
|status_unavailable_replicas|Int|
|status_conditions|JSON|
|status_collision_count|Int|