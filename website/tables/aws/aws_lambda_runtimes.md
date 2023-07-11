# Table: aws_lambda_runtimes

This table shows data for AWS Lambda Runtimes.

https://docs.aws.amazon.com/lambda/latest/dg/lambda-runtimes.html

The primary key for this table is **name**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|name (PK)|`utf8`|

## Example Queries

These SQL queries are sampled from CloudQuery policies and are compatible with PostgreSQL.

### Lambda functions should use supported runtimes

```sql
SELECT
  'Lambda functions should use supported runtimes' AS title,
  f.account_id,
  f.arn AS resource_id,
  CASE WHEN r.name IS NULL THEN 'fail' ELSE 'pass' END AS status
FROM
  aws_lambda_functions AS f
  LEFT JOIN aws_lambda_runtimes AS r ON r.name = f.configuration->>'Runtime'
WHERE
  f.configuration->>'PackageType' != 'Image';
```


