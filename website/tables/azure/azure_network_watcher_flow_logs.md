# Table: azure_network_watcher_flow_logs

This table shows data for Azure Network Watcher Flow Logs.

https://learn.microsoft.com/en-us/rest/api/network-watcher/flow-logs/list?tabs=HTTP#definitions

The primary key for this table is **id**.

## Relations

This table depends on [azure_network_watchers](azure_network_watchers).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|subscription_id|`utf8`|
|id (PK)|`utf8`|
|location|`utf8`|
|properties|`json`|
|tags|`json`|
|etag|`utf8`|
|name|`utf8`|
|type|`utf8`|

## Example Queries

These SQL queries are sampled from CloudQuery policies and are compatible with PostgreSQL.

### Ensure that Network Security Group Flow Log retention period is ''greater than 90 days''

```sql
SELECT
  e'Ensure that Network Security Group Flow Log retention period is \'greater than 90 days\''
    AS title,
  subscription_id AS subscription_id,
  id AS resource_id,
  CASE
  WHEN (properties->'retentionPolicy'->>'days')::INT8 >= 90 THEN 'pass'
  ELSE 'fail'
  END
    AS status
FROM
  azure_network_watcher_flow_logs;
```


