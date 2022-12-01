# Table: k8s_batch_cron_jobs



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
|spec_schedule|String|
|spec_time_zone|String|
|spec_starting_deadline_seconds|Int|
|spec_concurrency_policy|String|
|spec_suspend|Bool|
|spec_job_template|JSON|
|spec_successful_jobs_history_limit|Int|
|spec_failed_jobs_history_limit|Int|
|status_active|JSON|
|status_last_schedule_time|Timestamp|
|status_last_successful_time|Timestamp|