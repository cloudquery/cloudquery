# Table: aws_wellarchitected_workload_milestones

This table shows data for AWS Well-Architected Workload Milestones.

https://docs.aws.amazon.com/wellarchitected/latest/APIReference/API_MilestoneSummary.html

The primary key for this table is **_cq_id**.
The following fields are used to calculate the value of `_cq_id`: (**workload_arn**, **milestone_name**).
## Relations

This table depends on [aws_wellarchitected_workloads](aws_wellarchitected_workloads.md).

The following tables depend on aws_wellarchitected_workload_milestones:
  - [aws_wellarchitected_lens_reviews](aws_wellarchitected_lens_reviews.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|workload_arn|`utf8`|
|workload_id|`utf8`|
|milestone_name|`utf8`|
|milestone_number|`int64`|
|recorded_at|`timestamp[us, tz=UTC]`|