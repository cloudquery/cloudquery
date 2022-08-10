
# Table: aws_glue_connections
Defines a connection to a data source
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|arn|text|ARN of the resource|
|account_id|text|The AWS Account ID of the resource|
|region|text|The AWS Region of the resource|
|connection_properties|jsonb|Key-value pairs that define parameters for the connection|
|connection_type|text|The type of the connection|
|creation_time|timestamp without time zone|The time that this connection definition was created|
|description|text|The description of the connection|
|last_updated_by|text|The user, group, or role that last updated this connection definition|
|last_updated_time|timestamp without time zone|The last time that this connection definition was updated|
|match_criteria|text[]|A list of criteria that can be used in selecting this connection|
|name|text|The name of the connection definition|
|availability_zone|text|The connection's Availability Zone|
|security_group_id_list|text[]|The security group ID list used by the connection|
|subnet_id|text|The subnet ID used by the connection|
