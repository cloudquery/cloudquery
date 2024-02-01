# Table: aws_wellarchitected_lens_reviews

This table shows data for AWS Well-Architected Lens Reviews.

https://docs.aws.amazon.com/wellarchitected/latest/APIReference/API_LensReview.html

The primary key for this table is **_cq_id**.
The following fields are used to calculate the value of `_cq_id`: (**workload_arn**, **milestone_number**, **lens_alias**).
## Relations

This table depends on [aws_wellarchitected_workload_milestones](aws_wellarchitected_workload_milestones.md).

The following tables depend on aws_wellarchitected_lens_reviews:
  - [aws_wellarchitected_lens_review_improvements](aws_wellarchitected_lens_review_improvements.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|workload_arn|`utf8`|
|workload_id|`utf8`|
|milestone_number|`int64`|
|lens_alias|`utf8`|
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