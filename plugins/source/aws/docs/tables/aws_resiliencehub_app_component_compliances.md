# Table: aws_resiliencehub_app_component_compliances

https://docs.aws.amazon.com/resilience-hub/latest/APIReference/API_AppComponentCompliance.html

The primary key for this table is **app_component_name**.

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
|app_component_name (PK)|String|
|compliance|JSON|
|cost|JSON|
|message|String|
|resiliency_score|JSON|
|status|String|