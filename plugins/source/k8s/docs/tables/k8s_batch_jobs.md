# Table: k8s_batch_jobs



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
|spec_parallelism|Int|
|spec_completions|Int|
|spec_active_deadline_seconds|Int|
|spec_pod_failure_policy|JSON|
|spec_backoff_limit|Int|
|spec_selector|JSON|
|spec_manual_selector|Bool|
|spec_template|JSON|
|spec_ttl_seconds_after_finished|Int|
|spec_completion_mode|String|
|spec_suspend|Bool|
|status_conditions|JSON|
|status_start_time|Timestamp|
|status_completion_time|Timestamp|
|status_active|Int|
|status_succeeded|Int|
|status_failed|Int|
|status_completed_indexes|String|
|status_uncounted_terminated_pods|JSON|
|status_ready|Int|