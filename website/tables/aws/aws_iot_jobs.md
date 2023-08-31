# Table: aws_iot_jobs

This table shows data for AWS IoT Jobs.

https://docs.aws.amazon.com/iot/latest/apireference/API_Job.html

The primary key for this table is **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|tags|`json`|
|arn (PK)|`utf8`|
|abort_config|`json`|
|comment|`utf8`|
|completed_at|`timestamp[us, tz=UTC]`|
|created_at|`timestamp[us, tz=UTC]`|
|description|`utf8`|
|destination_package_versions|`list<item: utf8, nullable>`|
|document_parameters|`json`|
|force_canceled|`bool`|
|is_concurrent|`bool`|
|job_arn|`utf8`|
|job_executions_retry_config|`json`|
|job_executions_rollout_config|`json`|
|job_id|`utf8`|
|job_process_details|`json`|
|job_template_arn|`utf8`|
|last_updated_at|`timestamp[us, tz=UTC]`|
|namespace_id|`utf8`|
|presigned_url_config|`json`|
|reason_code|`utf8`|
|scheduled_job_rollouts|`json`|
|scheduling_config|`json`|
|status|`utf8`|
|target_selection|`utf8`|
|targets|`list<item: utf8, nullable>`|
|timeout_config|`json`|