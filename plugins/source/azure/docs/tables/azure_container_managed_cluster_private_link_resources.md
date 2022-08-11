
# Table: azure_container_managed_cluster_private_link_resources
PrivateLinkResource a private link resource
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|managed_cluster_cq_id|uuid|Unique CloudQuery ID of azure_container_managed_clusters table (FK)|
|id|text|The ID of the private link resource|
|name|text|The name of the private link resource|
|type|text|The resource type|
|group_id|text|The group ID of the resource|
|required_members|text[]|RequiredMembers of the resource|
|private_link_service_id|text|The private link service ID of the resource, this field is exposed only to NRP internally|
