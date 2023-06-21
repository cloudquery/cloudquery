# Table: aws_wellarchitected_workload_milestones

This table shows data for Wellarchitected Workload Milestones.

https://docs.aws.amazon.com/wellarchitected/latest/APIReference/API_Workload.html

The composite primary key for this table is (**account_id**, **region**, **workload_id**, **name**).

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
|name (PK)|`utf8`|
|number|`int64`|
|recorded_at|`timestamp[us, tz=UTC]`|