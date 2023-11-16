# Table: aws_servicecatalog_portfolios

This table shows data for AWS Service Catalog Portfolios.

https://docs.aws.amazon.com/servicecatalog/latest/dg/API_DescribePortfolio.html

The primary key for this table is **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn (PK)|`utf8`|
|tags|`json`|
|budgets|`json`|
|portfolio_detail|`json`|
|tag_options|`json`|