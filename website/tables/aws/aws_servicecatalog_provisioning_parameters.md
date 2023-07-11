# Table: aws_servicecatalog_provisioning_parameters

This table shows data for AWS Service Catalog Provisioning Parameters.

https://docs.aws.amazon.com/servicecatalog/latest/dg/API_DescribeProvisioningParameters.html

The composite primary key for this table is (**account_id**, **region**, **product_id**, **provisioning_artifact_id**).

## Relations

This table depends on [aws_servicecatalog_provisioned_products](aws_servicecatalog_provisioned_products).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|region (PK)|`utf8`|
|product_id (PK)|`utf8`|
|provisioning_artifact_id (PK)|`utf8`|
|constraint_summaries|`json`|
|provisioning_artifact_output_keys|`json`|
|provisioning_artifact_parameters|`json`|
|provisioning_artifact_preferences|`json`|
|tag_options|`json`|
|usage_instructions|`json`|