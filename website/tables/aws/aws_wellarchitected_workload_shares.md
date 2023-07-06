# Table: aws_wellarchitected_workload_shares

This table shows data for AWS Well-Architected Workload Shares.

https://docs.aws.amazon.com/wellarchitected/latest/APIReference/API_WorkloadShareSummary.html

The composite primary key for this table is (**workload_arn**, **share_id**).

## Relations

This table depends on [aws_wellarchitected_workloads](aws_wellarchitected_workloads).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|workload_arn (PK)|`utf8`|
|permission_type|`utf8`|
|share_id (PK)|`utf8`|
|shared_with|`utf8`|
|status|`utf8`|
|status_message|`utf8`|