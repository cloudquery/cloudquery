# Table: aws_resiliencehub_recommendation_templates

This table shows data for AWS Resilience Hub Recommendation Templates.

https://docs.aws.amazon.com/resilience-hub/latest/APIReference/API_RecommendationTemplate.html

The composite primary key for this table is (**arn**, **assessment_arn**, **app_arn**).

## Relations

This table depends on [aws_resiliencehub_app_assessments](aws_resiliencehub_app_assessments).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn (PK)|`utf8`|
|assessment_arn (PK)|`utf8`|
|format|`utf8`|
|name|`utf8`|
|recommendation_template_arn|`utf8`|
|recommendation_types|`list<item: utf8, nullable>`|
|status|`utf8`|
|app_arn (PK)|`utf8`|
|end_time|`timestamp[us, tz=UTC]`|
|message|`utf8`|
|needs_replacements|`bool`|
|recommendation_ids|`list<item: utf8, nullable>`|
|start_time|`timestamp[us, tz=UTC]`|
|tags|`json`|
|templates_location|`json`|