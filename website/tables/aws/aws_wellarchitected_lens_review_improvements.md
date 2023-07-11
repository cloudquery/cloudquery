# Table: aws_wellarchitected_lens_review_improvements

This table shows data for AWS Well-Architected Lens Review Improvements.

https://docs.aws.amazon.com/wellarchitected/latest/APIReference/API_ImprovementSummary.html

The composite primary key for this table is (**workload_arn**, **milestone_number**, **lens_alias**, **pillar_id**, **question_id**).

## Relations

This table depends on [aws_wellarchitected_lens_reviews](aws_wellarchitected_lens_reviews).

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
|improvement_plan_url|`utf8`|
|improvement_plans|`json`|
|pillar_id (PK)|`utf8`|
|question_id (PK)|`utf8`|
|question_title|`utf8`|
|risk|`utf8`|