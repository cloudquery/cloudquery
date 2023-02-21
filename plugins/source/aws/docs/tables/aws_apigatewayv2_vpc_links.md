# Table: aws_apigatewayv2_vpc_links

https://docs.aws.amazon.com/apigateway/latest/api/API_VpcLink.html

The composite primary key for this table is (**account_id**, **arn**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id (PK)|String|
|region|String|
|arn (PK)|String|
|name|String|
|security_group_ids|StringArray|
|subnet_ids|StringArray|
|vpc_link_id|String|
|created_date|Timestamp|
|tags|JSON|
|vpc_link_status|String|
|vpc_link_status_message|String|
|vpc_link_version|String|