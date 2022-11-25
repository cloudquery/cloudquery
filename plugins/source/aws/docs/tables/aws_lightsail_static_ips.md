# Table: aws_lightsail_static_ips

https://docs.aws.amazon.com/lightsail/2016-11-28/api-reference/API_StaticIp.html

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
|attached_to|String|
|created_at|Timestamp|
|ip_address|String|
|is_attached|Bool|
|location|JSON|
|name|String|
|resource_type|String|
|support_code|String|