# Table: aws_sagemaker_apps

This table shows data for Amazon SageMaker Apps.

https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_DescribeApp.html

The primary key for this table is **arn**.

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
|tags|JSON|
|app_arn|String|
|app_name|String|
|app_type|String|
|creation_time|Timestamp|
|domain_id|String|
|failure_reason|String|
|last_health_check_timestamp|Timestamp|
|last_user_activity_timestamp|Timestamp|
|resource_spec|JSON|
|space_name|String|
|status|String|
|user_profile_name|String|