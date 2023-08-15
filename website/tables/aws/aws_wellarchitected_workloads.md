# Table: aws_wellarchitected_workloads

This table shows data for AWS Well-Architected Workloads.

https://docs.aws.amazon.com/wellarchitected/latest/APIReference/API_Workload.html

The primary key for this table is **arn**.

## Relations

The following tables depend on aws_wellarchitected_workloads:
  - [aws_wellarchitected_workload_milestones](aws_wellarchitected_workload_milestones)
  - [aws_wellarchitected_workload_shares](aws_wellarchitected_workload_shares)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn (PK)|`utf8`|
|account_ids|`list<item: utf8, nullable>`|
|applications|`list<item: utf8, nullable>`|
|architectural_design|`utf8`|
|aws_regions|`list<item: utf8, nullable>`|
|description|`utf8`|
|discovery_config|`json`|
|environment|`utf8`|
|improvement_status|`utf8`|
|industry|`utf8`|
|industry_type|`utf8`|
|is_review_owner_update_acknowledged|`bool`|
|lenses|`list<item: utf8, nullable>`|
|non_aws_regions|`list<item: utf8, nullable>`|
|notes|`utf8`|
|owner|`utf8`|
|pillar_priorities|`list<item: utf8, nullable>`|
|prioritized_risk_counts|`json`|
|profiles|`json`|
|review_owner|`utf8`|
|review_restriction_date|`timestamp[us, tz=UTC]`|
|risk_counts|`json`|
|share_invitation_id|`utf8`|
|tags|`json`|
|updated_at|`timestamp[us, tz=UTC]`|
|workload_arn|`utf8`|
|workload_id|`utf8`|
|workload_name|`utf8`|