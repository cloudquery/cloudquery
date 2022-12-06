# Table: aws_shield_protections

https://docs.aws.amazon.com/waf/latest/DDOSAPIReference/API_Protection.html

The primary key for this table is **arn**.



## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|arn (PK)|String|
|tags|JSON|
|application_layer_automatic_response_configuration|JSON|
|health_check_ids|StringArray|
|id|String|
|name|String|
|resource_arn|String|