
# Table: azure_sql_managed_instance_private_endpoint_connections
ManagedInstancePecProperty a private endpoint connection under a managed instance
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|managed_instance_cq_id|uuid|Unique CloudQuery ID of azure_sql_managed_instances table (FK)|
|id|text|Resource ID|
|private_endpoint_id|text|Resource id of the private endpoint|
|private_link_service_connection_state_status|text|The private link service connection status|
|private_link_service_connection_state_description|text|The private link service connection description|
|private_link_service_connection_state_actions_required|text|The private link service connection description|
|provisioning_state|text|State of the Private Endpoint Connection|
