# Table: aws_resiliencehub_app_assessments

https://docs.aws.amazon.com/resilience-hub/latest/APIReference/API_AppAssessment.html

The composite primary key for this table is (**app_arn**, **arn**).

## Relations

This table depends on [aws_resiliencehub_apps](aws_resiliencehub_apps.md).

The following tables depend on aws_resiliencehub_app_assessments:
  - [aws_resiliencehub_alarm_recommendations](aws_resiliencehub_alarm_recommendations.md)
  - [aws_resiliencehub_app_component_compliances](aws_resiliencehub_app_component_compliances.md)
  - [aws_resiliencehub_component_recommendations](aws_resiliencehub_component_recommendations.md)
  - [aws_resiliencehub_recommendation_templates](aws_resiliencehub_recommendation_templates.md)
  - [aws_resiliencehub_sop_recommendations](aws_resiliencehub_sop_recommendations.md)
  - [aws_resiliencehub_test_recommendations](aws_resiliencehub_test_recommendations.md)

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
|arn (PK)|String|
|assessment_arn|String|
|assessment_status|String|
|invoker|String|
|app_version|String|
|assessment_name|String|
|compliance|JSON|
|compliance_status|String|
|cost|JSON|
|end_time|Timestamp|
|message|String|
|policy|JSON|
|resiliency_score|JSON|
|resource_errors_details|JSON|
|start_time|Timestamp|
|tags|JSON|