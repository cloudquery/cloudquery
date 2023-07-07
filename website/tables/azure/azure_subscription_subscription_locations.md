# Table: azure_subscription_subscription_locations

This table shows data for Azure Subscription Subscription Locations.

https://learn.microsoft.com/en-us/rest/api/resources/subscriptions/list-locations?tabs=HTTP#location

The primary key for this table is **id**.

## Relations

This table depends on [azure_subscription_subscriptions](azure_subscription_subscriptions).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|subscription_id|`utf8`|
|latitude|`utf8`|
|longitude|`utf8`|
|metadata|`json`|
|display_name|`utf8`|
|id (PK)|`utf8`|
|name|`utf8`|
|regional_display_name|`utf8`|
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


