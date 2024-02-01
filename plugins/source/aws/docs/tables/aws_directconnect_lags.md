# Table: aws_directconnect_lags

This table shows data for AWS Direct Connect Lags.

https://docs.aws.amazon.com/directconnect/latest/APIReference/API_Lag.html

The primary key for this table is **_cq_id**.
The following fields are used to calculate the value of `_cq_id`: (**request_account_id**, **arn**).

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
|allows_hosted_connections|`bool`|
|aws_device|`utf8`|
|aws_device_v2|`utf8`|
|aws_logical_device_id|`utf8`|
|connections|`json`|
|connections_bandwidth|`utf8`|
|encryption_mode|`utf8`|
|has_logical_redundancy|`utf8`|
|jumbo_frame_capable|`bool`|
|lag_id|`utf8`|
|lag_name|`utf8`|
|lag_state|`utf8`|
|location|`utf8`|
|mac_sec_capable|`bool`|
|mac_sec_keys|`json`|
|minimum_links|`int64`|
|number_of_connections|`int64`|
|owner_account|`utf8`|
|provider_name|`utf8`|
|region|`utf8`|