# Table: aws_wellarchitected_workload_milestones

This table shows data for AWS Well-Architected Workload Milestones.

https://docs.aws.amazon.com/wellarchitected/latest/APIReference/API_MilestoneSummary.html

The composite primary key for this table is (**workload_arn**, **milestone_name**).

## Relations

This table depends on [aws_wellarchitected_workloads](aws_wellarchitected_workloads).

The following tables depend on aws_wellarchitected_workload_milestones:
  - [aws_wellarchitected_lens_reviews](aws_wellarchitected_lens_reviews)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|workload_arn (PK)|`utf8`|
|workload_id|`utf8`|
|milestone_name (PK)|`utf8`|
|milestone_number|`int64`|
|recorded_at|`timestamp[us, tz=UTC]`|