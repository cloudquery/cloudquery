# Table: azure_security_settings

This table shows data for Azure Security Settings.

https://learn.microsoft.com/en-us/rest/api/defenderforcloud/settings/list?tabs=HTTP#settingslist

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|subscription_id|`utf8`|
|kind|`utf8`|
|id (PK)|`utf8`|
|name|`utf8`|
|type|`utf8`|

## Example Queries

These SQL queries are sampled from CloudQuery policies and are compatible with PostgreSQL.

### Ensure that Microsoft Cloud App Security (MCAS) integration with Security Center is selected (Automatic)

```sql
SELECT
  'Ensure that Microsoft Cloud App Security (MCAS) integration with Security Center is selected (Automatic)'
    AS title,
  subscription_id,
  id,
  CASE
  WHEN enabled = true THEN 'pass'
  ELSE 'fail'
  END
FROM
  azure_security_settings AS ass
WHERE
  name = 'MCAS';
```


