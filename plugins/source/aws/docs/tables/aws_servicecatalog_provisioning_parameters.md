# Table: aws_servicecatalog_provisioning_parameters

This table shows data for AWS Service Catalog Provisioning Parameters.

https://docs.aws.amazon.com/servicecatalog/latest/dg/API_DescribeProvisioningParameters.html

The composite primary key for this table is (**account_id**, **region**, **provisioned_product_arn**, **product_id**, **provisioning_artifact_id**, **path_id**).

## Relations

This table depends on [aws_servicecatalog_launch_paths](aws_servicecatalog_launch_paths.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|region (PK)|`utf8`|
|provisioned_product_arn (PK)|`utf8`|
|product_id (PK)|`utf8`|
|provisioning_artifact_id (PK)|`utf8`|
|path_id (PK)|`utf8`|
|constraint_summaries|`json`|
|provisioning_artifact_output_keys|`json`|
|provisioning_artifact_parameters|`json`|
|provisioning_artifact_preferences|`json`|
|tag_options|`json`|
|usage_instructions|`json`|