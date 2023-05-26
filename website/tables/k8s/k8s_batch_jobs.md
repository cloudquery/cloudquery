# Table: k8s_batch_jobs

This table shows data for Kubernetes (K8s) Batch Jobs.

The primary key for this table is **uid**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|utf8|
|_cq_sync_time|timestamp[us, tz=UTC]|
|_cq_id|uuid|
|_cq_parent_id|uuid|
|context|utf8|
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
|spec_parallelism|int64|
|spec_completions|int64|
|spec_active_deadline_seconds|int64|
|spec_pod_failure_policy|json|
|spec_backoff_limit|int64|
|spec_selector|json|
|spec_manual_selector|bool|
|spec_template|json|
|spec_ttl_seconds_after_finished|int64|
|spec_completion_mode|utf8|
|spec_suspend|bool|
|status_conditions|json|
|status_start_time|timestamp[us, tz=UTC]|
|status_completion_time|timestamp[us, tz=UTC]|
|status_active|int64|
|status_succeeded|int64|
|status_failed|int64|
|status_completed_indexes|utf8|
|status_uncounted_terminated_pods|json|
|status_ready|int64|