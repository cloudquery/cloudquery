# Table: aws_servicecatalog_portfolios

This table shows data for AWS Service Catalog Portfolios.

https://docs.aws.amazon.com/servicecatalog/latest/dg/API_PortfolioDetail.html

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
|description|`utf8`|
|display_name|`utf8`|
|id|`utf8`|
|provider_name|`utf8`|