# Table: aws_directconnect_connections

https://docs.aws.amazon.com/directconnect/latest/APIReference/API_Connection.html

The composite primary key for this table is (**arn**, **id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|arn (PK)|String|
|id (PK)|String|
|aws_device|String|
|aws_device_v2|String|
|aws_logical_device_id|String|
|bandwidth|String|
|connection_id|String|
|connection_name|String|
|connection_state|String|
|encryption_mode|String|
|has_logical_redundancy|String|
|jumbo_frame_capable|Bool|
|lag_id|String|
|loa_issue_time|Timestamp|
|location|String|
|mac_sec_capable|Bool|
|mac_sec_keys|JSON|
|owner_account|String|
|partner_name|String|
|port_encryption_status|String|
|provider_name|String|
|tags|JSON|
|vlan|Int|