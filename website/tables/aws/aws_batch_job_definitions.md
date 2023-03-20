# Table: aws_batch_job_definitions

This table shows data for Batch Job Definitions.

https://docs.aws.amazon.com/batch/latest/APIReference/API_DescribeJobDefinitions.html

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
|job_definition_arn|String|
|job_definition_name|String|
|revision|Int|
|type|String|
|container_orchestration_type|String|
|container_properties|JSON|
|eks_properties|JSON|
|node_properties|JSON|
|parameters|JSON|
|platform_capabilities|StringArray|
|propagate_tags|Bool|
|retry_strategy|JSON|
|scheduling_priority|Int|
|status|String|
|timeout|JSON|