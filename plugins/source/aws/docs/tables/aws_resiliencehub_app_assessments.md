# Table: aws_resiliencehub_app_assessments

This table shows data for AWS Resilience Hub App Assessments.

https://docs.aws.amazon.com/resilience-hub/latest/APIReference/API_AppAssessment.html

The primary key for this table is **_cq_id**.
The following fields are used to calculate the value of `_cq_id`: (**app_arn**, **arn**).
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
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|app_arn|`utf8`|
|arn|`utf8`|
|assessment_arn|`utf8`|
|assessment_status|`utf8`|
|invoker|`utf8`|
|app_version|`utf8`|
|assessment_name|`utf8`|
|compliance|`json`|
|compliance_status|`utf8`|
|cost|`json`|
|drift_status|`utf8`|
|end_time|`timestamp[us, tz=UTC]`|
|message|`utf8`|
|policy|`json`|
|resiliency_score|`json`|
|resource_errors_details|`json`|
|start_time|`timestamp[us, tz=UTC]`|
|tags|`json`|
|version_name|`utf8`|