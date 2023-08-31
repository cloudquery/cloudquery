# Table: aws_ec2_transit_gateway_attachments

This table shows data for Amazon Elastic Compute Cloud (EC2) Transit Gateway Attachments.

https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_TransitGatewayAttachment.html

The primary key for this table is **_cq_id**.

## Relations

This table depends on [aws_ec2_transit_gateways](aws_ec2_transit_gateways).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|transit_gateway_arn|`utf8`|
|tags|`json`|
|association|`json`|
|creation_time|`timestamp[us, tz=UTC]`|
|resource_id|`utf8`|
|resource_owner_id|`utf8`|
|resource_type|`utf8`|
|state|`utf8`|
|transit_gateway_attachment_id|`utf8`|
|transit_gateway_id|`utf8`|
|transit_gateway_owner_id|`utf8`|

## Example Queries

These SQL queries are sampled from CloudQuery policies and are compatible with PostgreSQL.

### Unused transit gateway

```sql
WITH
  attachment
    AS (
      SELECT
        DISTINCT transit_gateway_arn
      FROM
        aws_ec2_transit_gateway_attachments
    )
SELECT
  'Unused transit gateway' AS title,
  gateway.account_id,
  gateway.arn AS resource_id,
  'fail' AS status
FROM
  aws_ec2_transit_gateways AS gateway
  LEFT JOIN attachment ON attachment.transit_gateway_arn = gateway.arn
WHERE
  attachment.transit_gateway_arn IS NULL;
```


