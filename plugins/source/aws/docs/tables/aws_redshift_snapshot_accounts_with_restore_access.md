
# Table: aws_redshift_snapshot_accounts_with_restore_access
Describes an AWS customer account authorized to restore a snapshot.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|snapshot_cq_id|uuid|Unique CloudQuery ID of aws_redshift_snapshots table (FK)|
|account_alias|text|The identifier of an AWS support account authorized to restore a snapshot|
|account_id|text|The identifier of an AWS customer account authorized to restore a snapshot.|
