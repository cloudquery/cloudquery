# Table: aws_lightsail_disks

This table shows data for Lightsail Disks.

https://docs.aws.amazon.com/lightsail/2016-11-28/api-reference/API_Disk.html

The primary key for this table is **arn**.

## Relations

The following tables depend on aws_lightsail_disks:
  - [aws_lightsail_disk_snapshots](aws_lightsail_disk_snapshots)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn (PK)|`utf8`|
|tags|`json`|
|add_ons|`json`|
|attached_to|`utf8`|
|attachment_state|`utf8`|
|auto_mount_status|`utf8`|
|created_at|`timestamp[us, tz=UTC]`|
|gb_in_use|`int64`|
|iops|`int64`|
|is_attached|`bool`|
|is_system_disk|`bool`|
|location|`json`|
|name|`utf8`|
|path|`utf8`|
|resource_type|`utf8`|
|size_in_gb|`int64`|
|state|`utf8`|
|support_code|`utf8`|

## Example Queries

These SQL queries are sampled from CloudQuery policies and are compatible with PostgreSQL.

### Unused Lightsail disks

```sql
SELECT
  'Unused Lightsail disks' AS title,
  account_id,
  arn AS resource_id,
  'fail' AS status
FROM
  aws_lightsail_disks
WHERE
  is_attached = false;
```


