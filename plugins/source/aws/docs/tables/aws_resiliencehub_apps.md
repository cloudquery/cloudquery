# Table: aws_resiliencehub_apps

https://docs.aws.amazon.com/resilience-hub/latest/APIReference/API_App.html

The primary key for this table is **arn**.

## Relations

The following tables depend on aws_resiliencehub_apps:
  - [aws_resiliencehub_app_assessments](aws_resiliencehub_app_assessments.md)
  - [aws_resiliencehub_app_versions](aws_resiliencehub_app_versions.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|arn (PK)|String|
|app_arn|String|
|creation_time|Timestamp|
|name|String|
|assessment_schedule|String|
|compliance_status|String|
|description|String|
|last_app_compliance_evaluation_time|Timestamp|
|last_resiliency_score_evaluation_time|Timestamp|
|policy_arn|String|
|resiliency_score|Float|
|status|String|
|tags|JSON|