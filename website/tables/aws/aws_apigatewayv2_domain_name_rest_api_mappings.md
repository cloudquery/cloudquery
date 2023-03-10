# Table: aws_apigatewayv2_domain_name_rest_api_mappings

This table shows data for AWS API Gateway v2 Domain Name Rest API Mappings.

https://docs.aws.amazon.com/apigateway/latest/api/API_ApiMapping.html

The composite primary key for this table is (**account_id**, **arn**).

## Relations

This table depends on [aws_apigatewayv2_domain_names](aws_apigatewayv2_domain_names).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id (PK)|String|
|region|String|
|domain_name_arn|String|
|arn (PK)|String|
|api_id|String|
|stage|String|
|api_mapping_id|String|
|api_mapping_key|String|