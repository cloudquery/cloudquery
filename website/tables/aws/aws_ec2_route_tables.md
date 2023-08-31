# Table: aws_ec2_route_tables

This table shows data for Amazon Elastic Compute Cloud (EC2) Route Tables.

https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_RouteTable.html

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
|associations|`json`|
|owner_id|`utf8`|
|propagating_vgws|`json`|
|route_table_id|`utf8`|
|routes|`json`|
|vpc_id|`utf8`|

## Example Queries

These SQL queries are sampled from CloudQuery policies and are compatible with PostgreSQL.

### Find all ec2 instances that have unrestricted access to the internet with a wide open security group and routing

```sql
-- Find all AWS instances that are in a subnet that includes a catchall route
SELECT
  'Find all ec2 instances that have unrestricted access to the internet with a wide open security group and routing'
    AS title,
  account_id,
  instance_id AS resource_id,
  'fail' AS status
FROM
  aws_ec2_instances
WHERE
  subnet_id
  IN (
      SELECT
        a->>'SubnetId'
      FROM
        aws_ec2_route_tables AS t,
        jsonb_array_elements(t.associations) AS a,
        jsonb_array_elements(t.routes) AS r
      WHERE
        r->>'DestinationCidrBlock' = '0.0.0.0/0'
        OR r->>'DestinationIpv6CidrBlock' = '::/0'
    )
  AND instance_id
    IN (
        SELECT
          instance_id
        FROM
          aws_ec2_instances,
          jsonb_array_elements(security_groups) AS sg
          INNER JOIN view_aws_security_group_egress_rules ON id = sg->>'GroupId'
        WHERE
          ip = '0.0.0.0/0' OR ip6 = '::/0'
      );
```

### Unused route table

```sql
SELECT
  'Unused route table' AS title,
  account_id,
  arn AS resource_id,
  'fail' AS status
FROM
  aws_ec2_route_tables
WHERE
  COALESCE(jsonb_array_length(associations), 0) = 0;
```

### Find all lambda functions that have unrestricted access to the internet

```sql
SELECT
  DISTINCT
  'Find all lambda functions that have unrestricted access to the internet'
    AS title,
  account_id,
  arn AS resource_id,
  'fail' AS status
FROM
  aws_lambda_functions,
  jsonb_array_elements_text(configuration->'VpcConfig'->'SecurityGroupIds')
    AS sgs,
  jsonb_array_elements_text(configuration->'VpcConfig'->' SubnetIds') AS sns
WHERE
  sns
  IN (
      SELECT
        a->>'SubnetId'
      FROM
        public.aws_ec2_route_tables,
        jsonb_array_elements(associations) AS a,
        jsonb_array_elements(routes) AS r
      WHERE
        r->>'DestinationCidrBlock' = '0.0.0.0/0'
        OR r->>'DestinationIpv6CidrBlock' = '::/0'
    )
  AND sgs
    IN (
        SELECT
          id
        FROM
          view_aws_security_group_egress_rules
        WHERE
          ip = '0.0.0.0/0' OR ip6 = '::/0'
      )
UNION
  SELECT
    DISTINCT
    'Find all lambda functions that have unrestricted access to the internet'
      AS title,
    account_id,
    arn AS resource_id,
    'fail' AS status
  FROM
    aws_lambda_functions
  WHERE
    (configuration->'VpcConfig'->>'VpcId') IS NULL;
```


