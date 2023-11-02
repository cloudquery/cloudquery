# Table: aws_ec2_internet_gateways

This table shows data for Amazon Elastic Compute Cloud (EC2) Internet Gateways.

https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_InternetGateway.html

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
|attachments|`json`|
|internet_gateway_id|`utf8`|
|owner_id|`utf8`|

## Example Queries

These SQL queries are sampled from CloudQuery policies and are compatible with PostgreSQL.

### Unused internet gateway

```sql
SELECT
  'Unused internet gateway' AS title,
  account_id,
  arn AS resource_id,
  'fail' AS status
FROM
  aws_ec2_internet_gateways
WHERE
  COALESCE(jsonb_array_length(attachments), 0) = 0;
```


