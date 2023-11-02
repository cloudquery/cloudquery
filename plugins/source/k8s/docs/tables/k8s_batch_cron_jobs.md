# Table: k8s_batch_cron_jobs

This table shows data for Kubernetes (K8s) Batch Cron Jobs.

The primary key for this table is **uid**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|context|`utf8`|
|kind|`utf8`|
|api_version|`utf8`|
|name|`utf8`|
|namespace|`utf8`|
|uid (PK)|`utf8`|
|resource_version|`utf8`|
|generation|`int64`|
|deletion_grace_period_seconds|`int64`|
|labels|`json`|
|annotations|`json`|
|owner_references|`json`|
|finalizers|`list<item: utf8, nullable>`|
|spec_schedule|`utf8`|
|spec_time_zone|`utf8`|
|spec_starting_deadline_seconds|`int64`|
|spec_concurrency_policy|`utf8`|
|spec_suspend|`bool`|
|spec_job_template|`json`|
|spec_successful_jobs_history_limit|`int64`|
|spec_failed_jobs_history_limit|`int64`|
|status_active|`json`|
|status_last_schedule_time|`timestamp[us, tz=UTC]`|
|status_last_successful_time|`timestamp[us, tz=UTC]`|