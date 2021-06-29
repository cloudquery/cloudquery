
# Table: aws_directconnect_connections
Information about a Direct Connect Connection
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|aws_device_v2|text|The Direct Connect endpoint on which the physical connection terminates.|
|bandwidth|text|The bandwidth of the connection.|
|connection_id|text|The ID of the connection.|
|connection_name|text|The name of the connection.|
|connection_state|text|The state of the connection. Possible values are: ordering, requested, pending, available, down, deleting, deleted, rejected, unknown|
|encryption_mode|text|The MAC Security (MACsec) connection encryption mode. The valid values are: no_encrypt, should_encrypt, and must_encrypt.|
|has_logical_redundancy|text|Indicates whether the connection supports a secondary BGP peer in the same address family (IPv4/IPv6). Valid values are: yes, no, unknown|
|jumbo_frame_capable|boolean|Indicates whether jumbo frames (9001 MTU) are supported.|
|lag_id|text|The ID of the LAG.|
|loa_issue_time|timestamp without time zone|The time of the most recent call to DescribeLoa for this connection.|
|location|text|The location of the connection.|
|mac_sec_capable|boolean|Indicates whether the connection supports MAC Security (MACsec).|
|owner_account|text|The ID of the AWS account that owns the connection.|
|partner_name|text|The name of the AWS Direct Connect service provider associated with the connection.|
|port_encryption_status|text|The MAC Security (MACsec) port link status of the connection. The valid values are Encryption Up, which means that there is an active Connection Key Name, or Encryption Down.|
|provider_name|text|The name of the service provider associated with the connection.|
|tags|jsonb|The tags associated with the connection.|
|vlan|integer|The ID of the VLAN.|
