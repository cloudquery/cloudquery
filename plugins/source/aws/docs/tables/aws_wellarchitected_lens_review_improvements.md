# Table: aws_wellarchitected_lens_review_improvements

This table shows data for AWS Well-Architected Lens Review Improvements.

https://docs.aws.amazon.com/wellarchitected/latest/APIReference/API_ImprovementSummary.html

The primary key for this table is **_cq_id**.
The following fields are used to calculate the value of `_cq_id`: (**workload_arn**, **milestone_number**, **lens_alias**, **pillar_id**, **question_id**).
## Relations

This table depends on [aws_wellarchitected_lens_reviews](aws_wellarchitected_lens_reviews.md).

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
|improvement_plan_url|`utf8`|
|improvement_plans|`json`|
|pillar_id|`utf8`|
|question_id|`utf8`|
|question_title|`utf8`|
|risk|`utf8`|