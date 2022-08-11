
# Table: azure_storage_account_private_endpoint_connections
Azure storage account private endpoint connection
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_cq_id|uuid|Unique ID of azure_storage_accounts table (FK)|
|private_endpoint_id|text|The ARM identifier for Private Endpoint|
|private_link_service_connection_state_status|text|Indicates whether the connection has been Approved/Rejected/Removed by the owner of the service Possible values include: 'Pending', 'Approved', 'Rejected'|
|private_link_service_connection_state_description|text|The reason for approval/rejection of the connection|
|private_link_service_connection_state_action_required|text|A message indicating if changes on the service provider require any updates on the consumer|
|provisioning_state|text|The provisioning state of the private endpoint connection resource Possible values include: 'PrivateEndpointConnectionProvisioningStateSucceeded', 'PrivateEndpointConnectionProvisioningStateCreating', 'PrivateEndpointConnectionProvisioningStateDeleting', 'PrivateEndpointConnectionProvisioningStateFailed'|
|id|text|Fully qualified resource ID for the resource Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}|
|name|text|The name of the resource|
|type|text|The type of the resource Eg "MicrosoftCompute/virtualMachines" or "MicrosoftStorage/storageAccounts"|
