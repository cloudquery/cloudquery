# Table: aws_apigateway_rest_api_models

This table shows data for Amazon API Gateway Rest API Models.

https://docs.aws.amazon.com/apigateway/latest/api/API_Model.html

The composite primary key for this table is (**account_id**, **arn**).

## Relations

This table depends on [aws_apigateway_rest_apis](aws_apigateway_rest_apis).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id (PK)|String|
|region|String|
|rest_api_arn|String|
|arn (PK)|String|
|model_template|String|
|content_type|String|
|description|String|
|id|String|
|name|String|
|schema|String|