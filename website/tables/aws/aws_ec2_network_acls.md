# Table: aws_ec2_network_acls

This table shows data for Amazon Elastic Compute Cloud (EC2) Network ACLs.

https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_NetworkAcl.html

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
|entries|`json`|
|is_default|`bool`|
|network_acl_id|`utf8`|
|owner_id|`utf8`|
|vpc_id|`utf8`|

## Example Queries

These SQL queries are sampled from CloudQuery policies and are compatible with PostgreSQL.

### Unused network access control list

```sql
SELECT
  'Unused network access control list' AS title,
  account_id,
  arn AS resource_id,
  'fail' AS status
FROM
  aws_ec2_network_acls
WHERE
  COALESCE(jsonb_array_length(associations), 0) = 0;
```

### Unused network access control lists should be removed

```sql
WITH
  results
    AS (
      SELECT
        DISTINCT
        account_id,
        network_acl_id AS resource_id,
        CASE
        WHEN (a->>'NetworkAclAssociationId') IS NULL THEN 'pass'
        ELSE 'fail'
        END
          AS status
      FROM
        aws_ec2_network_acls
        LEFT JOIN jsonb_array_elements(aws_ec2_network_acls.associations)
            AS a ON true
    )
SELECT
  'Unused network access control lists should be removed' AS title,
  account_id,
  resource_id,
  status
FROM
  results;
```


