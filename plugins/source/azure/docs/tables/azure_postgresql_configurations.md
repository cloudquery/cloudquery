# Table: azure_postgresql_configurations


The primary key for this table is **id**.

## Relations
This table depends on [`azure_postgresql_servers`](azure_postgresql_servers.md).

## Columns
| Name          | Type          |
| ------------- | ------------- |
|subscription_id|String|
|postgresql_server_id|UUID|
|value|String|
|description|String|
|default_value|String|
|data_type|String|
|allowed_values|String|
|source|String|
|id (PK)|String|
|name|String|
|type|String|
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|