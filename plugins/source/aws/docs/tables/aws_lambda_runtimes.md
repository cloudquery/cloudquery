# Table: aws_lambda_runtimes

This table shows data for AWS Lambda Runtimes.

https://docs.aws.amazon.com/lambda/latest/dg/lambda-runtimes.html

The primary key for this table is **_cq_id**.
The following field is used to calculate the value of `_cq_id`: **name**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|name|`utf8`|