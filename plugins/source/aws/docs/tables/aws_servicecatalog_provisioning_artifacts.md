# Table: aws_servicecatalog_provisioning_artifacts

This table shows data for AWS Service Catalog Provisioning Artifacts.

https://docs.aws.amazon.com/servicecatalog/latest/dg/API_DescribeProvisioningArtifact.html

The composite primary key for this table is (**provisioned_product_arn**, **product_id**, **provisioning_artifact_id**).

## Relations

This table depends on [aws_servicecatalog_provisioned_products](aws_servicecatalog_provisioned_products.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|provisioned_product_arn (PK)|`utf8`|
|product_id (PK)|`utf8`|
|provisioning_artifact_id (PK)|`utf8`|
|info|`json`|
|provisioning_artifact_detail|`json`|
|provisioning_artifact_parameters|`json`|
|status|`utf8`|