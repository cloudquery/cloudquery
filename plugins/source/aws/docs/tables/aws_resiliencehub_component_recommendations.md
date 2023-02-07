# Table: aws_resiliencehub_component_recommendations

https://docs.aws.amazon.com/resilience-hub/latest/APIReference/API_ComponentRecommendation.html

The composite primary key for this table is (**account_id**, **region**, **app_component_name**).

## Relations

This table depends on [aws_resiliencehub_app_assessments](aws_resiliencehub_app_assessments.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id (PK)|String|
|region (PK)|String|
|app_component_name (PK)|String|
|config_recommendations|JSON|
|recommendation_status|String|