# Table: gcp_cloudidentity_device_users

https://cloud.google.com/identity/docs/reference/rest/v1/devices.deviceUsers#DeviceUser

The composite primary key for this table is (**project_id**, **name**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|project_id (PK)|String|
|compromised_state|String|
|create_time|String|
|first_sync_time|String|
|language_code|String|
|last_sync_time|String|
|management_state|String|
|name (PK)|String|
|password_state|String|
|user_agent|String|
|user_email|String|