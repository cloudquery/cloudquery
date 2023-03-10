# Table: aws_glacier_data_retrieval_policies

https://docs.aws.amazon.com/amazonglacier/latest/dev/api-GetDataRetrievalPolicy.html

The composite primary key for this table is (**account_id**, **region**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id (PK)|String|
|region (PK)|String|
|rules|JSON|