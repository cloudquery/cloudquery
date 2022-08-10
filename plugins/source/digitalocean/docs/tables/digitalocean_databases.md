
# Table: digitalocean_databases
Database represents a DigitalOcean managed database product
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|id|text|A unique ID that can be used to identify and reference a database cluster.|
|name|text|A unique, human-readable name referring to a database cluster.|
|engine|text|A slug representing the database engine used for the cluster. The possible values are: "pg" for PostgreSQL, "mysql" for MySQL, "redis" for Redis, and "mongodb" for MongoDB.|
|version|text|A string representing the version of the database engine in use for the cluster.|
|connection_uri|text|A connection string in the format accepted by the `psql` command. This is provided as a convenience and should be able to be constructed by the other attributes.|
|connection_database|text|The name of the default database.|
|connection_host|text|The FQDN pointing to the database cluster's current primary node.|
|connection_port|bigint|The port on which the database cluster is listening.|
|connection_user|text|The default user for the database.|
|connection_password|text|The randomly generated password for the default user.|
|connection_ssl|boolean|A boolean value indicating if the connection should be made over SSL.|
|private_connection_uri|text|A connection string in the format accepted by the `psql` command. This is provided as a convenience and should be able to be constructed by the other attributes.|
|private_connection_database|text|The name of the default database.|
|private_connection_host|text|The FQDN pointing to the database cluster's current primary node.|
|private_connection_port|bigint|The port on which the database cluster is listening.|
|private_connection_user|text|The default user for the database.|
|private_connection_password|text|The randomly generated password for the default user.|
|private_connection_ssl|boolean|A boolean value indicating if the connection should be made over SSL.|
|num_nodes|bigint|The number of nodes in the database cluster.|
|size_slug|text||
|db_names|text[]|An array of strings containing the names of databases created in the database cluster.|
|region_slug|text||
|status|text|A string representing the current status of the database cluster.|
|maintenance_window_day|text|The day of the week on which to apply maintenance updates.|
|maintenance_window_hour|text|The hour in UTC at which maintenance updates will be applied in 24 hour format.|
|maintenance_window_pending|boolean|A boolean value indicating whether any maintenance is scheduled to be performed in the next window.|
|maintenance_window_description|text[]|A list of strings, each containing information about a pending maintenance update.|
|created_at|timestamp without time zone|A time value given in ISO8601 combined date and time format that represents when the database cluster was created.|
|private_network_uuid|text|A string specifying the UUID of the VPC to which the database cluster will be assigned. If excluded, the cluster when creating a new database cluster, it will be assigned to your account's default VPC for the region.|
|tags|text[]|An array of tags that have been applied to the database cluster.|
