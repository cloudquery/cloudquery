
# Table: gcp_sql_instance_settings_ip_config_authorized_networks
An entry for an Access Control list
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|instance_cq_id|uuid|Unique ID of gcp_sql_instances table (FK)|
|instance_name|text||
|expiration_time|text|The time when this access control entry expires in RFC 3339 format, for example *2012-11-15T16:19:00094Z*|
|kind|text|This is always *sql#aclEntry*|
|name|text|A label to identify this entry|
|value|text|The allowlisted value for the access control list|
