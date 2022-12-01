# Table: aws_servicecatalog_provisioned_products

https://docs.aws.amazon.com/servicecatalog/latest/dg/API_ProvisionedProductAttribute.html

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
|created_time|Timestamp|
|id|String|
|idempotency_token|String|
|last_provisioning_record_id|String|
|last_record_id|String|
|last_successful_provisioning_record_id|String|
|name|String|
|physical_id|String|
|product_id|String|
|product_name|String|
|provisioning_artifact_id|String|
|provisioning_artifact_name|String|
|status|String|
|status_message|String|
|type|String|
|user_arn|String|
|user_arn_session|String|