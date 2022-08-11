
# Table: azure_mariadb_server_configurations
MariaDB server configuration
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|server_cq_id|uuid|Unique ID of azure_mariadb_servers table (FK)|
|id|text|Fully qualified resource ID for the resource.|
|name|text|The name of the resource.|
|type|text|The type of the resource.|
|value|text|Value of the configuration.|
|description|text|Description of the configuration.|
|default_value|text|Default value of the configuration.|
|data_type|text|Data type of the configuration.|
|allowed_values|text|Allowed values of the configuration.|
|source|text|Source of the configuration.|
