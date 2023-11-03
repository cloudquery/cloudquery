# Table: aws_config_configuration_recorders

This table shows data for Config Configuration Recorders.

https://docs.aws.amazon.com/config/latest/APIReference/API_ConfigurationRecorder.html

The primary key for this table is **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn (PK)|`utf8`|
|name|`utf8`|
|recording_group|`json`|
|role_arn|`utf8`|
|status_last_error_code|`utf8`|
|status_last_error_message|`utf8`|
|status_last_start_time|`timestamp[us, tz=UTC]`|
|status_last_status|`utf8`|
|status_last_status_change_time|`timestamp[us, tz=UTC]`|
|status_last_stop_time|`timestamp[us, tz=UTC]`|
|status_recording|`bool`|

## Example Queries

These SQL queries are sampled from CloudQuery policies and are compatible with PostgreSQL.

### AWS Config should be enabled

```sql
SELECT
  'AWS Config should be enabled' AS title,
  account_id,
  arn AS resource_id,
  CASE
  WHEN (recording_group->>'IncludeGlobalResourceTypes')::BOOL IS NOT true
  OR (recording_group->>'AllSupported')::BOOL IS NOT true
  OR status_recording IS NOT true
  OR status_last_status IS DISTINCT FROM 'SUCCESS'
  THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  aws_config_configuration_recorders;
```


