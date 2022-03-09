
# Table: azure_search_service_shared_private_link_resources
SharedPrivateLinkResource describes a Shared Private Link Resource managed by the Azure Cognitive Search service.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|service_cq_id|uuid|Unique CloudQuery ID of azure_search_services table (FK)|
|private_link_resource_id|text|The resource id of the resource the shared private link resource is for.|
|group_id|text|The group id from the provider of resource the shared private link resource is for.|
|request_message|text|The request message for requesting approval of the shared private link resource.|
|resource_region|text|Optional|
|status|text|Status of the shared private link resource|
|provisioning_state|text|The provisioning state of the shared private link resource|
|id|text|Fully qualified resource ID for the resource|
|name|text|The name of the resource|
|type|text|The type of the resource|
