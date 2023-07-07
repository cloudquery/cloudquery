# Table: aws_servicecatalog_provisioned_products

This table shows data for AWS Service Catalog Provisioned Products.

https://docs.aws.amazon.com/servicecatalog/latest/dg/API_ProvisionedProductAttribute.html

The primary key for this table is **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|arn (PK)|`utf8`|
|tags|`json`|
|created_time|`timestamp[us, tz=UTC]`|
|id|`utf8`|
|idempotency_token|`utf8`|
|last_provisioning_record_id|`utf8`|
|last_record_id|`utf8`|
|last_successful_provisioning_record_id|`utf8`|
|name|`utf8`|
|physical_id|`utf8`|
|product_id|`utf8`|
|product_name|`utf8`|
|provisioning_artifact_id|`utf8`|
|provisioning_artifact_name|`utf8`|
|status|`utf8`|
|status_message|`utf8`|
|type|`utf8`|
|user_arn|`utf8`|
|user_arn_session|`utf8`|