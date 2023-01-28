# Table: aws_resiliencehub_alarm_recommendations

https://docs.aws.amazon.com/resilience-hub/latest/APIReference/API_AlarmRecommendation.html

The primary key for this table is **name**.

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
|name (PK)|String|
|recommendation_id|String|
|reference_id|String|
|type|String|
|app_component_name|String|
|description|String|
|items|JSON|
|prerequisite|String|