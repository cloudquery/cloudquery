# Table: aws_batch_compute_environments

This table shows data for Batch Compute Environments.

https://docs.aws.amazon.com/batch/latest/APIReference/API_DescribeComputeEnvironments.html

The primary key for this table is **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|`utf8`|
|_cq_sync_time|`timestamp[us, tz=UTC]`|
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|tags|`json`|
|arn (PK)|`utf8`|
|compute_environment_arn|`utf8`|
|compute_environment_name|`utf8`|
|compute_resources|`json`|
|container_orchestration_type|`utf8`|
|ecs_cluster_arn|`utf8`|
|eks_configuration|`json`|
|service_role|`utf8`|
|state|`utf8`|
|status|`utf8`|
|status_reason|`utf8`|
|type|`utf8`|
|unmanagedv_cpus|`int64`|
|update_policy|`json`|
|uuid|`utf8`|