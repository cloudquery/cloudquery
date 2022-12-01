# Table: azure_compute_disks

https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2021-03-01/compute#Disk

The primary key for this table is **id**.



## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|subscription_id|String|
|managed_by|String|
|managed_by_extended|StringArray|
|sku|JSON|
|zones|StringArray|
|extended_location|JSON|
|time_created|Timestamp|
|os_type|String|
|hyper_v_generation|String|
|purchase_plan|JSON|
|creation_data|JSON|
|disk_size_gb|Int|
|disk_size_bytes|Int|
|unique_id|String|
|encryption_settings_collection|JSON|
|provisioning_state|String|
|disk_iops_read_write|Int|
|disk_m_bps_read_write|Int|
|disk_iops_read_only|Int|
|disk_m_bps_read_only|Int|
|disk_state|String|
|encryption|JSON|
|max_shares|Int|
|share_info|JSON|
|network_access_policy|String|
|disk_access_id|String|
|tier|String|
|bursting_enabled|Bool|
|property_updates_in_progress|JSON|
|supports_hibernation|Bool|
|security_profile|JSON|
|id (PK)|String|
|name|String|
|type|String|
|location|String|
|tags|JSON|