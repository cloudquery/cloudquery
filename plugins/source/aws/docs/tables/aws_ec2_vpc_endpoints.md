# Table: aws_ec2_vpc_endpoints

This table shows data for Amazon Elastic Compute Cloud (EC2) VPC Endpoints.

https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_VpcEndpoint.html

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
|creation_timestamp|`timestamp[us, tz=UTC]`|
|dns_entries|`json`|
|dns_options|`json`|
|groups|`json`|
|ip_address_type|`utf8`|
|last_error|`json`|
|network_interface_ids|`list<item: utf8, nullable>`|
|owner_id|`utf8`|
|policy_document|`utf8`|
|private_dns_enabled|`bool`|
|requester_managed|`bool`|
|route_table_ids|`list<item: utf8, nullable>`|
|service_name|`utf8`|
|state|`utf8`|
|subnet_ids|`list<item: utf8, nullable>`|
|vpc_endpoint_id|`utf8`|
|vpc_endpoint_type|`utf8`|
|vpc_id|`utf8`|

## Example Queries

These SQL queries are sampled from CloudQuery policies and are compatible with PostgreSQL.

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


