# Table: aws_resiliencehub_apps

This table shows data for AWS Resilience Hub Apps.

https://docs.aws.amazon.com/resilience-hub/latest/APIReference/API_App.html

The primary key for this table is **arn**.

## Relations

The following tables depend on aws_resiliencehub_apps:
  - [aws_resiliencehub_app_assessments](aws_resiliencehub_app_assessments)
  - [aws_resiliencehub_app_versions](aws_resiliencehub_app_versions)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn (PK)|`utf8`|
|app_arn|`utf8`|
|creation_time|`timestamp[us, tz=UTC]`|
|name|`utf8`|
|assessment_schedule|`utf8`|
|compliance_status|`utf8`|
|description|`utf8`|
|last_app_compliance_evaluation_time|`timestamp[us, tz=UTC]`|
|last_resiliency_score_evaluation_time|`timestamp[us, tz=UTC]`|
|policy_arn|`utf8`|
|resiliency_score|`float64`|
|status|`utf8`|
|tags|`json`|