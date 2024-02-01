# Table: aws_servicecatalog_products

This table shows data for AWS Service Catalog Products.

https://docs.aws.amazon.com/servicecatalog/latest/dg/API_DescribeProductAsAdmin.html

The primary key for this table is **_cq_id**.
The following field is used to calculate the value of `_cq_id`: **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn|`utf8`|
|tags|`json`|
|budgets|`json`|
|product_view_detail|`json`|
|provisioning_artifact_summaries|`json`|
|tag_options|`json`|