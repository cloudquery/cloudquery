# Table: aws_directconnect_connections

This table shows data for AWS Direct Connect Connections.

https://docs.aws.amazon.com/directconnect/latest/APIReference/API_Connection.html

The primary key for this table is **_cq_id**.
The following fields are used to calculate the value of `_cq_id`: (**request_account_id**, **arn**, **id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|request_account_id|`utf8`|
|request_region|`utf8`|
|arn|`utf8`|
|id|`utf8`|
|tags|`json`|
|aws_device|`utf8`|
|aws_device_v2|`utf8`|
|aws_logical_device_id|`utf8`|
|bandwidth|`utf8`|
|connection_id|`utf8`|
|connection_name|`utf8`|
|connection_state|`utf8`|
|encryption_mode|`utf8`|
|has_logical_redundancy|`utf8`|
|jumbo_frame_capable|`bool`|
|lag_id|`utf8`|
|loa_issue_time|`timestamp[us, tz=UTC]`|
|location|`utf8`|
|mac_sec_capable|`bool`|
|mac_sec_keys|`json`|
|owner_account|`utf8`|
|partner_name|`utf8`|
|port_encryption_status|`utf8`|
|provider_name|`utf8`|
|region|`utf8`|
|vlan|`int64`|