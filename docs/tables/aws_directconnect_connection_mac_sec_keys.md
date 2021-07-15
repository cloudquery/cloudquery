
# Table: aws_directconnect_connection_mac_sec_keys
The MAC Security (MACsec) security keys associated with the connection.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|connection_cq_id|uuid|Unique CloudQuery ID of aws_directconnect_connections table (FK)|
|connection_id|text|The ID of the connection.|
|ckn|text|The Connection Key Name (CKN) for the MAC Security secret key.|
|secret_arn|text|The Amazon Resource Name (ARN) of the MAC Security (MACsec) secret key.|
|start_on|text|The date that the MAC Security (MACsec) secret key takes effect. The value is displayed in UTC format.|
|state|text|The state of the MAC Security secret key. The possible values are: associating, associated, disassociating, disassociated|
