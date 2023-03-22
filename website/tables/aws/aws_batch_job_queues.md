# Table: aws_batch_job_queues

This table shows data for Batch Job Queues.

https://docs.aws.amazon.com/batch/latest/APIReference/API_DescribeJobQueues.html

The primary key for this table is **arn**.

## Relations

The following tables depend on aws_batch_job_queues:
  - [aws_batch_jobs](aws_batch_jobs)

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
|compute_environment_order|JSON|
|job_queue_arn|String|
|job_queue_name|String|
|priority|Int|
|state|String|
|scheduling_policy_arn|String|
|status|String|
|status_reason|String|