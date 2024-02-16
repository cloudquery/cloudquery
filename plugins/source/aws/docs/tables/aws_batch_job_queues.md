# Table: aws_batch_job_queues

This table shows data for Batch Job Queues.

https://docs.aws.amazon.com/batch/latest/APIReference/API_DescribeJobQueues.html

The primary key for this table is **_cq_id**.
The following field is used to calculate the value of `_cq_id`: **arn**.
## Relations

The following tables depend on aws_batch_job_queues:
  - [aws_batch_jobs](aws_batch_jobs.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|tags|`json`|
|arn|`utf8`|
|compute_environment_order|`json`|
|job_queue_arn|`utf8`|
|job_queue_name|`utf8`|
|priority|`int64`|
|state|`utf8`|
|scheduling_policy_arn|`utf8`|
|status|`utf8`|
|status_reason|`utf8`|