# Table: aws_wellarchitected_workload_shares

This table shows data for Wellarchitected Workload Shares.

https://docs.aws.amazon.com/wellarchitected/latest/APIReference/API_WorkloadShareSummary.html

The composite primary key for this table is (**account_id**, **region**, **workload_id**, **id**).

## Relations

This table depends on [aws_wellarchitected_workloads](aws_wellarchitected_workloads).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|`utf8`|
|_cq_sync_time|`timestamp[us, tz=UTC]`|
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|region (PK)|`utf8`|
|workload_id (PK)|`utf8`|
|permission_type|`utf8`|
|id (PK)|`utf8`|
|shared_with|`utf8`|
|status|`utf8`|
|status_message|`utf8`|