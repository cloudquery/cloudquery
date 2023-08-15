# Table: aws_ec2_ebs_snapshot_attributes

This table shows data for Amazon Elastic Compute Cloud (EC2) Amazon Elastic Block Store (EBS) Snapshot Attributes.

https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeSnapshotAttribute.html

The primary key for this table is **snapshot_arn**.

## Relations

This table depends on [aws_ec2_ebs_snapshots](aws_ec2_ebs_snapshots).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|snapshot_arn (PK)|`utf8`|
|create_volume_permissions|`json`|
|product_codes|`json`|
|snapshot_id|`utf8`|

## Example Queries

These SQL queries are sampled from CloudQuery policies and are compatible with PostgreSQL.

### Amazon EBS snapshots should not be public, determined by the ability to be restorable by anyone

```sql
WITH
  snapshot_access_groups
    AS (
      SELECT
        account_id,
        region,
        snapshot_id,
        jsonb_array_elements(create_volume_permissions)->>'Group' AS group,
        jsonb_array_elements(create_volume_permissions)->>'UserId' AS user_id
      FROM
        aws_ec2_ebs_snapshot_attributes
    )
SELECT
  DISTINCT
  'Amazon EBS snapshots should not be public, determined by the ability to be restorable by anyone'
    AS title,
  account_id,
  snapshot_id AS resource_id,
  CASE
  WHEN "group" = 'all' OR user_id IS DISTINCT FROM '' THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  snapshot_access_groups;
```


