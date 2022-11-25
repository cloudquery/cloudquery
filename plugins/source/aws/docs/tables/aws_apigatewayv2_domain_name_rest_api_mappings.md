# Table: aws_apigatewayv2_domain_name_rest_api_mappings

https://docs.aws.amazon.com/apigateway/latest/api/API_ApiMapping.html

The primary key for this table is **_cq_id**.

## Relations
This table depends on [aws_apigatewayv2_domain_names](aws_apigatewayv2_domain_names.md).


## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id (PK)|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|domain_name_arn|String|
|arn|String|
|api_id|String|
|stage|String|
|api_mapping_id|String|
|api_mapping_key|String|