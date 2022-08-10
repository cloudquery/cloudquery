
# Table: digitalocean_database_firewall_rules

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|database_cq_id|uuid|Unique CloudQuery ID of digitalocean_databases table (FK)|
|uuid|text|A unique ID for the firewall rule itself.|
|cluster_uuid|text|A unique ID for the database cluster to which the rule is applied.|
|type|text|The type of resource that the firewall rule allows to access the database cluster.|
|value|text|The ID of the specific resource, the name of a tag applied to a group of resources, or the IP address that the firewall rule allows to access the database cluster.|
|created_at|timestamp without time zone|A time value given in ISO8601 combined date and time format that represents when the firewall rule was created.|
