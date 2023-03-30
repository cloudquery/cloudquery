# Table: aws_redshift_data_shares

This table shows data for Redshift Data Shares.

https://docs.aws.amazon.com/redshift/latest/APIReference/API_DataShare.html

The primary key for this table is **_cq_id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id (PK)|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|allow_publicly_accessible_consumers|Bool|
|data_share_arn|String|
|data_share_associations|JSON|
|managed_by|String|
|producer_arn|String|