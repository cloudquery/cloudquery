# Table: aws_servicecatalog_products

This table shows data for AWS Service Catalog Products.

https://docs.aws.amazon.com/servicecatalog/latest/dg/API_ProductViewDetail.html

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
|product_arn|`utf8`|
|product_view_summary|`json`|
|source_connection|`json`|
|status|`utf8`|