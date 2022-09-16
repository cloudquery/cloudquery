
# Table: aws_directconnect_lag_mac_sec_keys
The MAC Security (MACsec) security keys associated with the LAG.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|lag_cq_id|uuid|Unique CloudQuery ID of aws_directconnect_lags table (FK)|
|lag_id|text|The ID of the LAG.|
|ckn|text|The Connection Key Name (CKN) for the MAC Security secret key.|
|secret_arn|text|The Amazon Resource Name (ARN) of the MAC Security (MACsec) secret key.|
|start_on|text|The date that the MAC Security (MACsec) secret key takes effect. The value is displayed in UTC format.|
|state|text|The state of the MAC Security secret key. The possible values are: associating, associated, disassociating, disassociated|
