
# Table: aws_qldb_ledgers

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|tags|jsonb|The tags associated with the pipeline.|
|arn|text|The Amazon Resource Name (ARN) for the ledger.|
|creation_date_time|timestamp without time zone|The date and time, in epoch time format, when the ledger was created|
|deletion_protection|boolean|The flag that prevents a ledger from being deleted by any user|
|encryption_status|text|The current state of encryption at rest for the ledger|
|kms_key_arn|text|The Amazon Resource Name (ARN) of the customer managed KMS key that the ledger uses for encryption at rest|
|inaccessible_kms_key_date_time|timestamp without time zone|The date and time, in epoch time format, when the KMS key first became inaccessible, in the case of an error|
|name|text|The name of the ledger.|
|permissions_mode|text|The permissions mode of the ledger.|
|state|text|The current status of the ledger.|
