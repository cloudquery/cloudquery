# Table: aws_apprunner_services

https://docs.aws.amazon.com/apprunner/latest/api/API_Service.html

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
|auto_scaling_configuration_summary|JSON|
|created_at|Timestamp|
|instance_configuration|JSON|
|network_configuration|JSON|
|service_id|String|
|service_name|String|
|service_url|String|
|source_configuration|JSON|
|status|String|
|updated_at|Timestamp|
|deleted_at|Timestamp|
|encryption_configuration|JSON|
|health_check_configuration|JSON|
|observability_configuration|JSON|