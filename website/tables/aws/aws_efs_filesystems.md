# Table: aws_efs_filesystems

This table shows data for Amazon Elastic File System (EFS) Filesystems.

https://docs.aws.amazon.com/efs/latest/ug/API_FileSystemDescription.html

The primary key for this table is **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn (PK)|`utf8`|
|backup_policy_status|`utf8`|
|tags|`json`|
|creation_time|`timestamp[us, tz=UTC]`|
|creation_token|`utf8`|
|file_system_id|`utf8`|
|life_cycle_state|`utf8`|
|number_of_mount_targets|`int64`|
|owner_id|`utf8`|
|performance_mode|`utf8`|
|size_in_bytes|`json`|
|availability_zone_id|`utf8`|
|availability_zone_name|`utf8`|
|encrypted|`bool`|
|file_system_arn|`utf8`|
|kms_key_id|`utf8`|
|name|`utf8`|
|provisioned_throughput_in_mibps|`float64`|
|throughput_mode|`utf8`|

## Example Queries

These SQL queries are sampled from CloudQuery policies and are compatible with PostgreSQL.

### Amazon EFS volumes should be in backup plans

```sql
SELECT
  'Amazon EFS volumes should be in backup plans' AS title,
  account_id,
  arn AS resource_id,
  CASE
  WHEN backup_policy_status IS DISTINCT FROM 'ENABLED' THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  aws_efs_filesystems;
```

### Unused EFS filesystem

```sql
SELECT
  'Unused EFS filesystem' AS title,
  account_id,
  arn AS resource_id,
  'fail' AS status
FROM
  aws_efs_filesystems
WHERE
  number_of_mount_targets = 0;
```

### Amazon EFS should be configured to encrypt file data at rest using AWS KMS

```sql
SELECT
  'Amazon EFS should be configured to encrypt file data at rest using AWS KMS'
    AS title,
  account_id,
  arn AS resource_id,
  CASE
  WHEN encrypted IS NOT true OR kms_key_id IS NULL THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  aws_efs_filesystems;
```


