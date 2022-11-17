# Table: azure_network_security_groups

https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v2#SecurityGroup

The primary key for this table is **id**.



## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|subscription_id|String|
|id (PK)|String|
|location|String|
|flush_connection|Bool|
|security_rules|JSON|
|default_security_rules|JSON|
|flow_logs|JSON|
|network_interfaces|JSON|
|provisioning_state|String|
|resource_guid|String|
|subnets|JSON|
|tags|JSON|
|etag|String|
|name|String|
|type|String|