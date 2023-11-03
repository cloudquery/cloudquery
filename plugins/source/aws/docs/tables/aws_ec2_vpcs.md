# Table: aws_ec2_vpcs

This table shows data for Amazon Elastic Compute Cloud (EC2) VPCs.

https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_Vpc.html

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
|cidr_block|`utf8`|
|cidr_block_association_set|`json`|
|dhcp_options_id|`utf8`|
|instance_tenancy|`utf8`|
|ipv6_cidr_block_association_set|`json`|
|is_default|`bool`|
|owner_id|`utf8`|
|state|`utf8`|
|vpc_id|`utf8`|

## Example Queries

These SQL queries are sampled from CloudQuery policies and are compatible with PostgreSQL.

### VPC flow logging should be enabled in all VPCs

```sql
SELECT
  'VPC flow logging should be enabled in all VPCs' AS title,
  aws_ec2_vpcs.account_id,
  aws_ec2_vpcs.arn,
  CASE
  WHEN aws_ec2_flow_logs.resource_id IS NULL THEN 'fail'
  ELSE 'pass'
  END
FROM
  aws_ec2_vpcs
  LEFT JOIN aws_ec2_flow_logs ON
      aws_ec2_vpcs.vpc_id = aws_ec2_flow_logs.resource_id;
```

### Amazon EC2 should be configured to use VPC endpoints that are created for the Amazon EC2 service

```sql
WITH
  endpoints
    AS (
      SELECT
        vpc_endpoint_id
      FROM
        aws_ec2_vpc_endpoints
      WHERE
        vpc_endpoint_type = 'Interface'
        AND service_name ~ concat('com.amazonaws.', region, '.ec2')
    )
SELECT
  'Amazon EC2 should be configured to use VPC endpoints that are created for the Amazon EC2 service'
    AS title,
  account_id,
  vpc_id AS resource_id,
  CASE
  WHEN endpoints.vpc_endpoint_id IS NULL THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  aws_ec2_vpcs
  LEFT JOIN endpoints ON aws_ec2_vpcs.vpc_id = endpoints.vpc_endpoint_id;
```


