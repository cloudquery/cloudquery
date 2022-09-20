
# Table: azure_batch_account_private_endpoint_connections
PrivateEndpointConnection contains information about a private link resource
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_cq_id|uuid|Unique CloudQuery ID of azure_batch_accounts table (FK)|
|provisioning_state|text|Possible values include: 'PrivateEndpointConnectionProvisioningStateSucceeded', 'PrivateEndpointConnectionProvisioningStateUpdating', 'PrivateEndpointConnectionProvisioningStateFailed'|
|private_endpoint_id|text|The resource id of the private endpoint resource from Microsoft.Network provider.|
|private_link_connection_status|text|Possible values include: 'PrivateLinkServiceConnectionStatusApproved', 'PrivateLinkServiceConnectionStatusPending', 'PrivateLinkServiceConnectionStatusRejected', 'PrivateLinkServiceConnectionStatusDisconnected'|
|private_link_connection_description|text|The description for the private link service connection state.|
|private_link_connection_action_required|text|A description of any extra actions that may be required.|
|id|text|The ID of the resource|
|name|text|The name of the resource|
|type|text|The type of the resource|
|etag|text|The ETag of the resource, used for concurrency statements|
