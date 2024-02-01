# Table: aws_servicecatalog_provisioning_artifacts

This table shows data for AWS Service Catalog Provisioning Artifacts.

https://docs.aws.amazon.com/servicecatalog/latest/dg/API_DescribeProvisioningArtifact.html

The primary key for this table is **_cq_id**.
The following fields are used to calculate the value of `_cq_id`: (**provisioned_product_arn**, **product_id**, **provisioning_artifact_id**).
## Relations

This table depends on [aws_servicecatalog_provisioned_products](aws_servicecatalog_provisioned_products.md).

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
|info|`json`|
|provisioning_artifact_detail|`json`|
|provisioning_artifact_parameters|`json`|
|status|`utf8`|