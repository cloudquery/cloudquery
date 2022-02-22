
# Table: azure_container_registry_replications
Replication an object that represents a replication for a container registry
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|registry_cq_id|uuid|Unique CloudQuery ID of azure_container_registries table (FK)|
|provisioning_state|text|The provisioning state of the replication at the time the operation was called|
|status|text|The short label for the status|
|status_message|text|The detailed message for the status, including alerts and error messages|
|status_timestamp|timestamp without time zone|The timestamp when the status was changed to the current value|
|id|text|The resource ID|
|name|text|The name of the resource|
|type|text|The type of the resource|
|location|text|The location of the resource|
|tags|jsonb|The tags of the resource|
