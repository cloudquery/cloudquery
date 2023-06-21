# Table: aws_wellarchitected_workloads

This table shows data for Wellarchitected Workloads.

https://docs.aws.amazon.com/wellarchitected/latest/APIReference/API_Workload.html

The composite primary key for this table is (**account_id**, **region**, **arn**).

## Relations

The following tables depend on aws_wellarchitected_workloads:
  - [aws_wellarchitected_workload_milestones](aws_wellarchitected_workload_milestones)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|`utf8`|
|_cq_sync_time|`timestamp[us, tz=UTC]`|
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|region (PK)|`utf8`|
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
|arn (PK)|`utf8`|
|id|`utf8`|
|name|`utf8`|