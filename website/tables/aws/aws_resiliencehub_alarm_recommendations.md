# Table: aws_resiliencehub_alarm_recommendations

This table shows data for AWS Resilience Hub Alarm Recommendations.

https://docs.aws.amazon.com/resilience-hub/latest/APIReference/API_AlarmRecommendation.html

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
|name|`utf8`|
|recommendation_id (PK)|`utf8`|
|reference_id|`utf8`|
|type|`utf8`|
|app_component_name|`utf8`|
|description|`utf8`|
|items|`json`|
|prerequisite|`utf8`|