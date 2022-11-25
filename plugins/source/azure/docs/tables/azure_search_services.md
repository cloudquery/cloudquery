# Table: azure_search_services

https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/services/search/mgmt/2020-08-01/search#Service

The primary key for this table is **id**.



## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|subscription_id|String|
|replica_count|Int|
|partition_count|Int|
|hosting_mode|String|
|public_network_access|String|
|status|String|
|status_details|String|
|provisioning_state|String|
|network_rule_set|JSON|
|private_endpoint_connections|JSON|
|shared_private_link_resources|JSON|
|sku|JSON|
|identity|JSON|
|tags|JSON|
|location|String|
|id (PK)|String|
|name|String|
|type|String|