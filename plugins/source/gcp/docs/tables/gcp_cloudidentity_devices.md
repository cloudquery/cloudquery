# Table: gcp_cloudidentity_devices

https://cloud.google.com/identity/docs/reference/rest/v1/devices#Device

The composite primary key for this table is (**project_id**, **name**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|project_id (PK)|String|
|android_specific_attributes|JSON|
|asset_tag|String|
|baseband_version|String|
|bootloader_version|String|
|brand|String|
|build_number|String|
|compromised_state|String|
|create_time|String|
|device_id|String|
|device_type|String|
|enabled_developer_options|Bool|
|enabled_usb_debugging|Bool|
|encryption_state|String|
|imei|String|
|kernel_version|String|
|last_sync_time|String|
|management_state|String|
|manufacturer|String|
|meid|String|
|model|String|
|name (PK)|String|
|network_operator|String|
|os_version|String|
|other_accounts|StringArray|
|owner_type|String|
|release_version|String|
|security_patch_time|String|
|serial_number|String|
|wifi_mac_addresses|StringArray|