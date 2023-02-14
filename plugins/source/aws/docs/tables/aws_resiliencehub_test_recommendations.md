# Table: aws_resiliencehub_test_recommendations

https://docs.aws.amazon.com/resilience-hub/latest/APIReference/API_TestRecommendation.html

The composite primary key for this table is (**app_arn**, **assessment_arn**, **recommendation_id**).

## Relations

This table depends on [aws_resiliencehub_app_assessments](aws_resiliencehub_app_assessments.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|app_arn (PK)|String|
|assessment_arn (PK)|String|
|reference_id|String|
|app_component_name|String|
|depends_on_alarms|StringArray|
|description|String|
|intent|String|
|items|JSON|
|name|String|
|prerequisite|String|
|recommendation_id (PK)|String|
|risk|String|
|type|String|