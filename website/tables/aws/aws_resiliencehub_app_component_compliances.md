# Table: aws_resiliencehub_app_component_compliances

This table shows data for AWS Resilience Hub App Component Compliances.

https://docs.aws.amazon.com/resilience-hub/latest/APIReference/API_AppComponentCompliance.html

The composite primary key for this table is (**app_arn**, **assessment_arn**, **app_component_name**).

## Relations

This table depends on [aws_resiliencehub_app_assessments](aws_resiliencehub_app_assessments).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|app_arn (PK)|`utf8`|
|assessment_arn (PK)|`utf8`|
|app_component_name (PK)|`utf8`|
|compliance|`json`|
|cost|`json`|
|message|`utf8`|
|resiliency_score|`json`|
|status|`utf8`|