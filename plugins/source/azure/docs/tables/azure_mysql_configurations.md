# Table: azure_mysql_configurations


The primary key for this table is **id**.

## Relations
This table depends on [`azure_mysql_servers`](azure_mysql_servers.md).

## Columns
| Name          | Type          |
| ------------- | ------------- |
|subscription_id|String|
|mysql_server_id|String|
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