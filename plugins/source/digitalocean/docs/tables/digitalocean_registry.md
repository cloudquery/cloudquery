
# Table: digitalocean_registry
Registry represents a registry.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|name|text|A globally unique name for the container registry. Must be lowercase and be composed only of numbers, letters and `-`, up to a limit of 63 characters.|
|storage_usage_bytes|bigint|The amount of storage used in the registry in bytes.|
|storage_usage_bytes_updated_at|timestamp without time zone|The time at which the storage usage was updated. Storage usage is calculated asynchronously, and may not immediately reflect pushes to the registry.|
|created_at|timestamp without time zone|A time value given in ISO8601 combined date and time format that represents when the registry was created.|
