# Table: aws_wellarchitected_lens_review_improvements

This table shows data for Wellarchitected Lens Review Improvements.

https://pkg.go.dev/github.com/aws/aws-sdk-go-v2/service/wellarchitected/types#Lens

The composite primary key for this table is (**account_id**, **region**, **pillar_id**, **question_id**).

## Relations

This table depends on [aws_wellarchitected_lens_reviews](aws_wellarchitected_lens_reviews).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|`utf8`|
|_cq_sync_time|`timestamp[us, tz=UTC]`|
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|region (PK)|`utf8`|
|workload_id|`utf8`|
|milestone_number|`int64`|
|lens_alias|`utf8`|
|improvement_plan_url|`utf8`|
|improvement_plans|`json`|
|pillar_id (PK)|`utf8`|
|question_id (PK)|`utf8`|
|question_title|`utf8`|
|risk|`utf8`|