# Table: aws_batch_jobs

This table shows data for Batch Jobs.

https://docs.aws.amazon.com/batch/latest/APIReference/API_DescribeJobs.html

The primary key for this table is **arn**.

## Relations

This table depends on [aws_batch_job_queues](aws_batch_job_queues).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|tags|`json`|
|arn (PK)|`utf8`|
|job_definition|`utf8`|
|job_id|`utf8`|
|job_name|`utf8`|
|job_queue|`utf8`|
|started_at|`int64`|
|status|`utf8`|
|array_properties|`json`|
|attempts|`json`|
|container|`json`|
|created_at|`int64`|
|depends_on|`json`|
|eks_attempts|`json`|
|eks_properties|`json`|
|is_cancelled|`bool`|
|is_terminated|`bool`|
|job_arn|`utf8`|
|node_details|`json`|
|node_properties|`json`|
|parameters|`json`|
|platform_capabilities|`list<item: utf8, nullable>`|
|propagate_tags|`bool`|
|retry_strategy|`json`|
|scheduling_priority|`int64`|
|share_identifier|`utf8`|
|status_reason|`utf8`|
|stopped_at|`int64`|
|timeout|`json`|