# Table: aws_resiliencehub_test_recommendations

This table shows data for AWS Resilience Hub Test Recommendations.

https://docs.aws.amazon.com/resilience-hub/latest/APIReference/API_TestRecommendation.html

The composite primary key for this table is (**app_arn**, **assessment_arn**, **recommendation_id**).

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
|reference_id|`utf8`|
|app_component_name|`utf8`|
|depends_on_alarms|`list<item: utf8, nullable>`|
|description|`utf8`|
|intent|`utf8`|
|items|`json`|
|name|`utf8`|
|prerequisite|`utf8`|
|recommendation_id (PK)|`utf8`|
|risk|`utf8`|
|type|`utf8`|