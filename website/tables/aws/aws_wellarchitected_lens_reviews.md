# Table: aws_wellarchitected_lens_reviews

This table shows data for AWS Well-Architected Lens Reviews.

https://docs.aws.amazon.com/wellarchitected/latest/APIReference/API_LensReview.html

The composite primary key for this table is (**workload_arn**, **milestone_number**, **lens_alias**).

## Relations

This table depends on [aws_wellarchitected_workload_milestones](aws_wellarchitected_workload_milestones).

The following tables depend on aws_wellarchitected_lens_reviews:
  - [aws_wellarchitected_lens_review_improvements](aws_wellarchitected_lens_review_improvements)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|workload_arn (PK)|`utf8`|
|workload_id|`utf8`|
|milestone_number (PK)|`int64`|
|lens_alias (PK)|`utf8`|
|lens_arn|`utf8`|
|lens_name|`utf8`|
|lens_status|`utf8`|
|lens_version|`utf8`|
|next_token|`utf8`|
|notes|`utf8`|
|pillar_review_summaries|`json`|
|prioritized_risk_counts|`json`|
|profiles|`json`|
|risk_counts|`json`|
|updated_at|`timestamp[us, tz=UTC]`|