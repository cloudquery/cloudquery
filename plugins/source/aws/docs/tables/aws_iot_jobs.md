# Table: aws_iot_jobs

https://docs.aws.amazon.com/iot/latest/apireference/API_Job.html

The primary key for this table is **arn**.



## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|tags|JSON|
|arn (PK)|String|
|abort_config|JSON|
|comment|String|
|completed_at|Timestamp|
|created_at|Timestamp|
|description|String|
|document_parameters|JSON|
|force_canceled|Bool|
|is_concurrent|Bool|
|job_executions_retry_config|JSON|
|job_executions_rollout_config|JSON|
|job_id|String|
|job_process_details|JSON|
|job_template_arn|String|
|last_updated_at|Timestamp|
|namespace_id|String|
|presigned_url_config|JSON|
|reason_code|String|
|scheduling_config|JSON|
|status|String|
|target_selection|String|
|targets|StringArray|
|timeout_config|JSON|