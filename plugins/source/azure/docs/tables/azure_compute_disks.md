# Table: azure_compute_disks

https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v4#Disk

The primary key for this table is **id**.



## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|subscription_id|String|
|location|String|
|extended_location|JSON|
|creation_data|JSON|
|bursting_enabled|Bool|
|completion_percent|Float|
|data_access_auth_mode|String|
|disk_access_id|String|
|disk_iops_read_only|Int|
|disk_iops_read_write|Int|
|disk_m_bps_read_only|Int|
|disk_m_bps_read_write|Int|
|disk_size_gb|Int|
|encryption|JSON|
|encryption_settings_collection|JSON|
|hyper_v_generation|String|
|max_shares|Int|
|network_access_policy|String|
|os_type|String|
|optimized_for_frequent_attach|Bool|
|public_network_access|String|
|purchase_plan|JSON|
|security_profile|JSON|
|supported_capabilities|JSON|
|supports_hibernation|Bool|
|tier|String|
|bursting_enabled_time|Timestamp|
|disk_size_bytes|Int|
|disk_state|String|
|property_updates_in_progress|JSON|
|provisioning_state|String|
|share_info|JSON|
|time_created|Timestamp|
|unique_id|String|
|sku|JSON|
|tags|JSON|
|zones|StringArray|
|id (PK)|String|
|managed_by|String|
|managed_by_extended|StringArray|
|name|String|
|type|String|