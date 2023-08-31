# Table: aws_glacier_data_retrieval_policies

This table shows data for Glacier Data Retrieval Policies.

https://docs.aws.amazon.com/amazonglacier/latest/dev/api-GetDataRetrievalPolicy.html

The composite primary key for this table is (**account_id**, **region**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|region (PK)|`utf8`|
|rules|`json`|