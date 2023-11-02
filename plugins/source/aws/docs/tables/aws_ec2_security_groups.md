# Table: aws_ec2_security_groups

This table shows data for Amazon Elastic Compute Cloud (EC2) Security Groups.

https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_SecurityGroup.html

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
|description|`utf8`|
|group_id|`utf8`|
|group_name|`utf8`|
|ip_permissions|`json`|
|ip_permissions_egress|`json`|
|owner_id|`utf8`|
|vpc_id|`utf8`|

## Example Queries

These SQL queries are sampled from CloudQuery policies and are compatible with PostgreSQL.

### The VPC default security group should not allow inbound and outbound traffic

```sql
SELECT
  'The VPC default security group should not allow inbound and outbound traffic'
    AS title,
  account_id,
  arn,
  CASE
  WHEN group_name = 'default'
  AND (
      jsonb_array_length(ip_permissions) > 0
      OR jsonb_array_length(ip_permissions_egress) > 0
    )
  THEN 'fail'
  ELSE 'pass'
  END
FROM
  aws_ec2_security_groups;
```

### Security group is not currently in use so it should be deleted

```sql
WITH
  interface_groups
    AS (
      SELECT
        DISTINCT g->>'GroupId' AS id
      FROM
        aws_ec2_instances AS i,
        jsonb_array_elements(network_interfaces) AS a,
        jsonb_array_elements(a->'Groups') AS g
    )
SELECT
  'security group is not currently in use so it should be deleted' AS title,
  account_id,
  arn AS resource_id,
  CASE WHEN interface_groups.id IS NULL THEN 'fail' ELSE 'pass' END AS status
FROM
  aws_ec2_security_groups
  LEFT JOIN interface_groups ON
      aws_ec2_security_groups.group_id = interface_groups.id;
```

### Unused EC2 security group

```sql
WITH
  interface_groups
    AS (
      SELECT
        DISTINCT a->>'GroupId' AS group_id
      FROM
        aws_ec2_instances, jsonb_array_elements(security_groups) AS a
    )
SELECT
  'Unused EC2 security group' AS title,
  sg.account_id,
  sg.arn AS resource_id,
  'fail' AS status
FROM
  aws_ec2_security_groups AS sg
  LEFT JOIN interface_groups ON interface_groups.group_id = sg.group_id
WHERE
  interface_groups.group_id IS NULL;
```


