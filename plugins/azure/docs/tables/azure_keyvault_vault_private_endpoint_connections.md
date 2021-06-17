
# Table: azure_keyvault_vault_private_endpoint_connections
Azure ketvault vault endpoint connection
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|vault_id|uuid|Unique ID of azure_keyvault_vaults table (FK)|
|private_endpoint_id|text|Full identifier of the private endpoint resource|
|private_link_service_connection_state_status|text|Indicates whether the connection has been approved, rejected or removed by the key vault owner Possible values include: 'PrivateEndpointServiceConnectionStatusPending', 'PrivateEndpointServiceConnectionStatusApproved', 'PrivateEndpointServiceConnectionStatusRejected', 'PrivateEndpointServiceConnectionStatusDisconnected'|
|private_link_service_connection_state_description|text|The reason for approval or rejection|
|private_link_service_connection_state_action_required|text|A message indicating if changes on the service provider require any updates on the consumer|
|provisioning_state|text|Provisioning state of the private endpoint connection Possible values include: 'Succeeded', 'Creating', 'Updating', 'Deleting', 'Failed', 'Disconnected'|
