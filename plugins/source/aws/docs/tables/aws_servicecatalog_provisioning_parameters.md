# Table: aws_servicecatalog_provisioning_parameters

This table shows data for AWS Service Catalog Provisioning Parameters.

https://docs.aws.amazon.com/servicecatalog/latest/dg/API_DescribeProvisioningParameters.html

The primary key for this table is **_cq_id**.
The following fields are used to calculate the value of `_cq_id`: (**account_id**, **region**, **provisioned_product_arn**, **product_id**, **provisioning_artifact_id**, **path_id**).
## Relations

This table depends on [aws_servicecatalog_launch_paths](aws_servicecatalog_launch_paths.md).

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
|path_id|`utf8`|
|constraint_summaries|`json`|
|provisioning_artifact_output_keys|`json`|
|provisioning_artifact_parameters|`json`|
|provisioning_artifact_preferences|`json`|
|tag_options|`json`|
|usage_instructions|`json`|