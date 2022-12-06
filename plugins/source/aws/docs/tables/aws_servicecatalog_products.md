# Table: aws_servicecatalog_products

https://docs.aws.amazon.com/servicecatalog/latest/dg/API_ProductViewDetail.html

The primary key for this table is **arn**.



## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|arn (PK)|String|
|tags|JSON|
|created_time|Timestamp|
|product_view_summary|JSON|
|source_connection|JSON|
|status|String|