# Table: azure_network_interfaces

https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/services/network/mgmt/2020-11-01/network#Interface

The primary key for this table is **id**.



## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|subscription_id|String|
|extended_location|JSON|
|virtual_machine|JSON|
|network_security_group|JSON|
|private_endpoint|JSON|
|ip_configurations|JSON|
|tap_configurations|JSON|
|dns_settings|JSON|
|mac_address|String|
|primary|Bool|
|enable_accelerated_networking|Bool|
|enable_ip_forwarding|Bool|
|hosted_workloads|StringArray|
|dscp_configuration|JSON|
|resource_guid|String|
|provisioning_state|String|
|nic_type|String|
|private_link_service|JSON|
|migration_phase|String|
|etag|String|
|id (PK)|String|
|name|String|
|type|String|
|location|String|
|tags|JSON|