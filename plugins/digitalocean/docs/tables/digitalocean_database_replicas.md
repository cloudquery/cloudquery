
# Table: digitalocean_database_replicas
DatabaseReplica represents a read-only replica of a particular database
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|database_cq_id|uuid|Unique CloudQuery ID of digitalocean_databases table (FK)|
|name|text|The name to give the read-only replicating|
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
|region|text|A slug identifier for the region where the read-only replica will be located. If excluded, the replica will be placed in the same region as the cluster.|
|status|text|A string representing the current status of the database cluster.|
|created_at|timestamp without time zone|A time value given in ISO8601 combined date and time format that represents when the database cluster was created.|
|private_network_uuid|text|A string specifying the UUID of the VPC to which the read-only replica will be assigned. If excluded, the replica will be assigned to your account's default VPC for the region.|
|tags|text[]|A flat array of tag names as strings to apply to the read-only replica after it is created. Tag names can either be existing or new tags.|
