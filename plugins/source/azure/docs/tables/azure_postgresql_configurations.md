# Table: azure_postgresql_configurations

https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/services/postgresql/mgmt/2020-01-01/postgresql#Configuration

The primary key for this table is **id**.

## Relations
This table depends on [azure_postgresql_servers](azure_postgresql_servers.md).


## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|subscription_id|String|
|postgresql_server_id|String|
|value|String|
|description|String|
|default_value|String|
|data_type|String|
|allowed_values|String|
|source|String|
|id (PK)|String|
|name|String|
|type|String|