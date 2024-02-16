# Table: aws_servicecatalog_launch_paths

This table shows data for AWS Service Catalog Launch Paths.

https://docs.aws.amazon.com/servicecatalog/latest/dg/API_LaunchPathSummary.html

The primary key for this table is **_cq_id**.
The following fields are used to calculate the value of `_cq_id`: (**account_id**, **region**, **provisioned_product_arn**, **product_id**, **provisioning_artifact_id**).
## Relations

This table depends on [aws_servicecatalog_provisioned_products](aws_servicecatalog_provisioned_products.md).

The following tables depend on aws_servicecatalog_launch_paths:
  - [aws_servicecatalog_provisioning_parameters](aws_servicecatalog_provisioning_parameters.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|provisioned_product_arn|`utf8`|
|product_id|`utf8`|
|provisioning_artifact_id|`utf8`|
|tags|`json`|
|constraint_summaries|`json`|
|id|`utf8`|
|name|`utf8`|