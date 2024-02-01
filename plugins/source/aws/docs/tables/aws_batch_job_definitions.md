# Table: aws_batch_job_definitions

This table shows data for Batch Job Definitions.

https://docs.aws.amazon.com/batch/latest/APIReference/API_DescribeJobDefinitions.html

The primary key for this table is **_cq_id**.
The following field is used to calculate the value of `_cq_id`: **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|tags|`json`|
|arn|`utf8`|
|job_definition_arn|`utf8`|
|job_definition_name|`utf8`|
|revision|`int64`|
|type|`utf8`|
|container_orchestration_type|`utf8`|
|container_properties|`json`|
|eks_properties|`json`|
|node_properties|`json`|
|parameters|`json`|
|platform_capabilities|`list<item: utf8, nullable>`|
|propagate_tags|`bool`|
|retry_strategy|`json`|
|scheduling_priority|`int64`|
|status|`utf8`|
|timeout|`json`|