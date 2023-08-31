# Table: azure_network_watchers

This table shows data for Azure Network Watchers.

https://learn.microsoft.com/en-us/rest/api/network-watcher/network-watchers/list-all?tabs=HTTP#networkwatcher

The primary key for this table is **id**.

## Relations

The following tables depend on azure_network_watchers:
  - [azure_network_watcher_flow_logs](azure_network_watcher_flow_logs)

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

### Network Watcher should be enable

```sql
SELECT
  'Network Watcher should be enable' AS title,
  l.subscription_id,
  l.id,
  CASE
  WHEN anw._cq_id IS NULL THEN 'fail'
  ELSE 'pass'
  END
FROM
  azure_subscription_subscription_locations AS l
  LEFT JOIN azure_network_watchers AS anw ON l.name = anw.location;
```


