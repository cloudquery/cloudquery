# Table: aws_batch_jobs

This table shows data for Batch Jobs.

https://docs.aws.amazon.com/batch/latest/APIReference/API_DescribeJobs.html

The primary key for this table is **arn**.

## Relations

This table depends on [aws_batch_job_queues](aws_batch_job_queues).

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
|job_definition|String|
|job_id|String|
|job_name|String|
|job_queue|String|
|started_at|Int|
|status|String|
|array_properties|JSON|
|attempts|JSON|
|container|JSON|
|created_at|Int|
|depends_on|JSON|
|eks_attempts|JSON|
|eks_properties|JSON|
|is_cancelled|Bool|
|is_terminated|Bool|
|job_arn|String|
|node_details|JSON|
|node_properties|JSON|
|parameters|JSON|
|platform_capabilities|StringArray|
|propagate_tags|Bool|
|retry_strategy|JSON|
|scheduling_priority|Int|
|share_identifier|String|
|status_reason|String|
|stopped_at|Int|
|timeout|JSON|