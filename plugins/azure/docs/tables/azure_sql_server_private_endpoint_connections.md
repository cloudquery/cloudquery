
# Table: azure_sql_server_private_endpoint_connections
List of private endpoint connections on a server
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|server_cq_id|uuid|Unique ID of azure_sql_servers table (FK)|
|id|text|Resource ID|
|private_endpoint_id|text|Resource id of the private endpoint|
|private_link_service_connection_state_status|text|The private link service connection status Possible values include: 'Approved', 'Pending', 'Rejected', 'Disconnected'|
|private_link_service_connection_state_description|text|The private link service connection description|
|private_link_service_connection_state_actions_required|text|The actions required for private link service connection Possible values include: 'PrivateLinkServiceConnectionStateActionsRequireNone'|
|provisioning_state|text|State of the private endpoint connection Possible values include: 'PrivateEndpointProvisioningStateApproving', 'PrivateEndpointProvisioningStateReady', 'PrivateEndpointProvisioningStateDropping', 'PrivateEndpointProvisioningStateFailed', 'PrivateEndpointProvisioningStateRejecting'|
