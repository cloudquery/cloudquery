
# Table: aws_directconnect_lags
Information about Direct Connect Link Aggregation Group (LAG)
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|allows_hosted_connections|boolean|Indicates whether the LAG can host other connections.|
|aws_device_v2|text|The AWS Direct Connect endpoint that hosts the LAG.|
|connection_ids|text[]|The list of IDs of Direct Connect Connections bundled by the LAG|
|connections_bandwidth|text|The individual bandwidth of the physical connections bundled by the LAG.|
|encryption_mode|text|The LAG MAC Security (MACsec) encryption mode.|
|has_logical_redundancy|text|Indicates whether the LAG supports a secondary BGP peer in the same address family (IPv4/IPv6).|
|jumbo_frame_capable|boolean|Indicates whether jumbo frames (9001 MTU) are supported.|
|id|text|The ID of the LAG.|
|name|text|The name of the LAG.|
|state|text|The state of the LAG. Possible values are: requested, pending, available, down, deleting, deleted, unknown|
|location|text|The location of the LAG.|
|mac_sec_capable|boolean|Indicates whether the LAG supports MAC Security (MACsec).|
|minimum_links|integer|The minimum number of physical dedicated connections that must be operational for the LAG itself to be operational.|
|number_of_connections|integer|The number of physical dedicated connections bundled by the LAG, up to a maximum of 10.|
|owner_account|text|The ID of the AWS account that owns the LAG.|
|provider_name|text|The name of the service provider associated with the LAG.|
|tags|jsonb|The tags associated with the LAG.|
