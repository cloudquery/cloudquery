# Table: aws_apigatewayv2_vpc_links

https://docs.aws.amazon.com/apigateway/latest/api/API_VpcLink.html

The primary key for this table is **_cq_id**.



## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id (PK)|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|arn|String|
|name|String|
|security_group_ids|StringArray|
|subnet_ids|StringArray|
|vpc_link_id|String|
|created_date|Timestamp|
|tags|JSON|
|vpc_link_status|String|
|vpc_link_status_message|String|
|vpc_link_version|String|