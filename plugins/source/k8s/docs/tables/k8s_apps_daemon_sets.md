# Table: k8s_apps_daemon_sets



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
|spec_selector|JSON|
|spec_template|JSON|
|spec_update_strategy|JSON|
|spec_min_ready_seconds|Int|
|spec_revision_history_limit|Int|
|status_current_number_scheduled|Int|
|status_number_misscheduled|Int|
|status_desired_number_scheduled|Int|
|status_number_ready|Int|
|status_observed_generation|Int|
|status_updated_number_scheduled|Int|
|status_number_available|Int|
|status_number_unavailable|Int|
|status_collision_count|Int|
|status_conditions|JSON|